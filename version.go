package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version = "master"
	date    = "1970-01-01T12:00:00-0000"
	commit  = "abcd1234"
)

func init() {
	rootCmd.AddCommand(newVersionCommand())
}

func newVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show the sensu-agent version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s version %s, build %s, built %s\n",
				checkName,
				version,
				commit,
				date,
			)
		},
	}

	return cmd
}
