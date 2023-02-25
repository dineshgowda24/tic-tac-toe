/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/dineshgowda24/tic-tac-toe/cli"
	"github.com/dineshgowda24/tic-tac-toe/source/server"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tic-tac-toe",
	Short: "Tic Tac Toe written in Golang",
	Long:  `Tic Tac Toe which can be played over the network`,
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run in server mode",
	Long: `Application will start accepting TCP connections.
Default port will be 8080.`,
	Aliases: []string{"serve"},
	Run: func(cmd *cobra.Command, args []string) {
		srv := server.New()
		srv.Serve()
	},
}

var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "Run in cli mode",
	Long:  `Application will allow you to play in cli.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Cli()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(cliCmd)
}
