/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/erda-project/hepa-operator/pkg/kong"
)

const (
	__contains__ = "__contains__"
)

// listPluginsCmd represents the listPlugins command
var listPluginsCmd = &cobra.Command{
	Use:   "plugins",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: ListPluginsCmd,
}

func init() {
	listCmd.AddCommand(listPluginsCmd)
}

func ListPluginsCmd(cmd *cobra.Command, args []string) {
	client, err := kong.NewClient(GetOpenapi())
	if err != nil {
		logrus.Fatalln(err)
	}
	var (
		pager       = client.PluginPager()
		filterFuncs = getPluginsFilterFuncs()
		list        = getPluginsList(pager, filterFuncs)
	)
	if len(list) == 0 {
		logrus.Infoln("plugin not found")
		return
	}
	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		logrus.Fatalln(err)
	}
	var out io.Writer = os.Stdout
	if GetOutput() != "" {
		file, err := os.OpenFile(GetOutput(), os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
		if err != nil {
			logrus.Fatalln(err)
		}
		out = file
		defer file.Close()
	}
	if _, err = fmt.Fprintln(out, string(data)); err != nil {
		logrus.Fatalln(err)
	}

	logrus.Infoln(len(list), "kong objects found")
}

type pluginsFilterFunc func(plugin kong.Plugin) bool

func getPluginsFilterFuncs() (filterFuncs []pluginsFilterFunc) {
	for _, condition := range GetFilter() {
		var f pluginsFilterFunc
		var expr string
		var exprF func(string, string) bool
		switch {
		case strings.Contains(condition, "="):
			expr = "="
			exprF = strings.EqualFold
		case strings.Contains(condition, __contains__):
			expr = __contains__
			exprF = strings.Contains
		default:
			logrus.Fatalf("unknown expr: %s", condition)
		}

		index := strings.Index(condition, expr)
		switch key, value := condition[:index], condition[index+len(expr):]; key {
		case "id":
			f = func(plugin kong.Plugin) bool {
				return exprF(plugin.Id, value)
			}
		case "name":
			f = func(plugin kong.Plugin) bool {
				return exprF(plugin.Name, value)
			}
		case "route.id":
			f = func(plugin kong.Plugin) bool {
				return plugin.Route != nil && exprF(plugin.Route.ID, value)
			}
		case "service.id":
			f = func(plugin kong.Plugin) bool {
				return plugin.Service != nil && exprF(plugin.Service.ID, value)
			}
		case "consumer.id":
			f = func(plugin kong.Plugin) bool {
				return plugin.Consumer != nil && exprF(plugin.Consumer.ID, value)
			}
		case "consumer.username":
			f = func(plugin kong.Plugin) bool {
				return plugin.Consumer != nil && exprF(plugin.Consumer.Username, value)
			}
		case "protocols":
			f = func(plugin kong.Plugin) bool {
				for _, protocol := range plugin.Protocols {
					if ok := exprF(protocol, value); !ok {
						return false
					}
				}
				return true
			}
		case "enabled":
			f = func(plugin kong.Plugin) bool {
				return exprF(fmt.Sprint(plugin.Enabled), value)
			}
		case "tags":
			f = func(plugin kong.Plugin) bool {
				for _, tag := range plugin.Tags {
					if ok := exprF(tag, value); !ok {
						return false
					}
				}
				return true
			}
		default:
			if strings.HasPrefix(key, "config.") {
				f = func(plugin kong.Plugin) bool {
					if len(plugin.Config) == 0 {
						return false
					}
					v, ok := plugin.Config[strings.TrimPrefix(key, "config.")]
					if !ok {
						return false
					}
					return exprF(fmt.Sprint(v), value)
				}
			} else {
				logrus.Fatalf("unknown key %s", key)
			}
		}
		if f != nil {
			filterFuncs = append(filterFuncs, f)
		}
	}

	return
}

func getPluginsList(pager *kong.PluginPager, filterFuncs []pluginsFilterFunc) (list []kong.Plugin) {
	for len(list) < GetMaxSize() {
		page, ok, err := pager.Next()
		if err != nil {
			logrus.Fatalln(err)
		}
		if !ok || len(page.Data) == 0 {
			return
		}
	RangePlugins:
		for j := 0; j < len(page.Data) && len(list) < GetMaxSize(); j++ {
			for _, f := range filterFuncs {
				if ok := f(page.Data[j]); !ok {
					continue RangePlugins
				}
			}
			list = append(list, page.Data[j])
		}
	}
	return
}
