/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/erda-project/hepa-operator/pkg/kong"
)

var (
	deleteInput string
	objName     string
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: Delete,
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.PersistentFlags().StringVarP(&deleteInput, "input", "i", "", "the delete list file")
	deleteCmd.PersistentFlags().StringVar(&objName, "obj-name", "", "the object name: service, route, plugin")
	var persistentRequires = []string{
		"input",
		"obj-name",
	}
	for _, item := range persistentRequires {
		if err := deleteCmd.MarkPersistentFlagRequired(item); err != nil {
			logrus.Fatalln(err)
		}
	}
}

func Delete(cmd *cobra.Command, args []string) {
	data, err := os.ReadFile(deleteInput)
	if err != nil {
		logrus.Fatalln(err)
	}
	var list []deleteItem
	if err = json.Unmarshal(data, &list); err != nil {
		logrus.Fatalln(err)
	}
	client, err := kong.NewClient(GetOpenapi())
	if err != nil {
		logrus.Fatalln(err)
	}
	for _, item := range list {
		if item.ID == "" {
			logrus.Fatalln("invalid id, delete nothing")
		}
	}
	for i, item := range list {
		switch objName {
		case "service", "services":
			logrus.Infof("delete service %v\t%s\n", i, item.ID)
			if err := client.ServiceDeleter().Delete(item.ID); err != nil {
				logrus.Fatalln(err)
			}
		case "route", "routes":
			logrus.Infof("delete route %v\t%s\n", i, item.ID)
			if item.Service != nil {
				logrus.Infof("\tdelete service %v\t%s\n", i, item.Service.ID)
				if err := client.ServiceDeleter().Delete(item.Service.ID); err != nil {
					logrus.Fatalln(err)
				}
			}
			if err := client.RouteDeleter().Delete(item.ID); err != nil {
				logrus.Fatalln(err)
			}
		case "plugin", "plugins":
			logrus.Infof("delete plugin %v\t%s\n", i, item.ID)
			if err := client.PluginDeleter().Delete(item.ID); err != nil {
				logrus.Fatalln(err)
			}
		default:
			logrus.Fatalf("unknown kong object name")
		}
	}
}

type deleteItem struct {
	ID      string      `json:"id,omitempty"`
	Service *deleteItem `json:"service,omitempty"`
}
