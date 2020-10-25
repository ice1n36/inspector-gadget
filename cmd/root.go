package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "inspector-gadget",
		Short: "inspector-gadget is a frida-gadget injection tool",
		Long:  `A tool that injects frida-gadget and configuration into your apk (and soon to come ipa)`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}
)

func init() {
	rootCmd.AddCommand(injectCmd)
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
