/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// arguments
var (
	filter  []string
	output  string
	maxSize int
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch objName {
		case "services", "service":
			ListServiceCmd(cmd, args)
		case "routes", "route":
			ListRoutesCmd(cmd, args)
		case "plugins", "plugin":
			ListPluginsCmd(cmd, args)
		default:
			logrus.Fatalln("Please specify --obj-name or use subcommand")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.PersistentFlags().StringSliceVar(&filter, "filter", nil, "filter condition. e.g. --filter host=https://sample.io --filter name__like__my-service")
	listCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "output json file")
	listCmd.PersistentFlags().StringVar(&objName, "obj-name", "", "kong object type name: services, routes, plugins")
	listCmd.PersistentFlags().IntVar(&maxSize, "max-size", 1000, "Maximum number of kong objects to query")
}

func GetFilter() []string {
	return filter
}

func GetOutput() string {
	return output
}

func GetMaxSize() int {
	return maxSize
}
