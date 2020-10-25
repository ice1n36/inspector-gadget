package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	AppPath   = ""
	injectCmd = &cobra.Command{
		Use:   "inject",
		Short: "inject frida-gadget in app",
		Long:  `injects frida-gadget and configuration into your apk (and soon to come ipa)`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			fmt.Printf("Injecting to %s\n", AppPath)
		},
	}
)

func init() {
	injectCmd.Flags().StringVarP(&AppPath, "app", "a", "", "Path to APK/IPA/etc (requied)")
	injectCmd.MarkFlagRequired("app")
}
