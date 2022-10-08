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

// listRoutesCmd represents the listRoutes command
var listRoutesCmd = &cobra.Command{
	Use:   "routes",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: ListRoutesCmd,
}

func init() {
	listCmd.AddCommand(listRoutesCmd)
}

func ListRoutesCmd(cmd *cobra.Command, args []string) {
	client, err := kong.NewClient(GetOpenapi())
	if err != nil {
		logrus.Fatalln(err)
	}
	var (
		pager       = client.RoutePager()
		filterFuncs = getRoutesFilterFuncs()
		list        = getRoutesList(pager, filterFuncs)
	)
	if len(list) == 0 {
		logrus.Infoln("route not found")
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

type routesFilterFunc func(route kong.Route) bool

func getRoutesFilterFuncs() (filterFuncs []routesFilterFunc) {
	for _, condition := range GetFilter() {
		var f routesFilterFunc
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
			f = func(route kong.Route) bool {
				return exprF(route.Id, value)
			}
		case "name":
			f = func(route kong.Route) bool {
				return exprF(route.Name, value)
			}
		case "hosts":
			f = func(route kong.Route) bool {
				for _, host := range route.Hosts {
					if exprF(host, value) {
						return true
					}
				}
				return false
			}
		case "paths":
			f = func(route kong.Route) bool {
				for _, pth := range route.Paths {
					if exprF(pth, value) {
						return true
					}
				}
				return false
			}
		case "tags":
			f = func(route kong.Route) bool {
				for _, tag := range route.Tags {
					if exprF(tag, value) {
						return true
					}
				}
				return false
			}
		case "service.id":
			f = func(route kong.Route) bool {
				return exprF(route.Service.Id, value)
			}
		default:
			logrus.Fatalf("unknown key %s", key)
		}
		if f != nil {
			filterFuncs = append(filterFuncs, f)
		}
	}
	return
}

func getRoutesList(pager *kong.RoutePager, filterFuncs []routesFilterFunc) (list []kong.Route) {
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
