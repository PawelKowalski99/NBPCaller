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

// checkHostCmd represents the checkHost command
var checkHostCmd = &cobra.Command{
	Use:   "checkHost",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		host, err := cmd.Flags().GetString("host")
		if err != nil {
			return err
		}

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

		err = h.CheckHost(host, x, y)
		if err != nil {
			return err
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(checkHostCmd)

	checkHostCmd.Flags().String("host", "", "host")
	_ = checkHostCmd.MarkFlagRequired("host")
	checkHostCmd.Flags().Int("x", 1, "x")
	checkHostCmd.Flags().Int("y", 1, "y")
}
