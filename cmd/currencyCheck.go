/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"VierCode/handlers"
	"github.com/spf13/cobra"
	"io"
	"os"
)

// currencyCheckCmd represents the currencyCheck command
var currencyCheckCmd = &cobra.Command{
	Use:   "currencyCheck",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		x, err := cmd.Flags().GetInt("x")
		if err != nil {
			return err
		}

		y, err := cmd.Flags().GetInt("y")
		if err != nil {
			return err
		}
		file, err := os.Create("./data/log.txt")
		if err != nil {
			return err
		}
		defer file.Close()

		w := io.MultiWriter(os.Stdout, file)

		h := handlers.NewHost(w)

		err = h.CurrencyCheck(x, y)
		if err != nil {
			return err
		}
		return nil

	},
}

func init() {
	rootCmd.AddCommand(currencyCheckCmd)

	currencyCheckCmd.Flags().Int("x", 1, "x")
	currencyCheckCmd.Flags().Int("y", 1, "y")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// currencyCheckCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// currencyCheckCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
