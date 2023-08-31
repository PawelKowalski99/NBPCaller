/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"VierCode/cmd"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	slog.SetDefault(logger)

	cmd.Execute()

}
