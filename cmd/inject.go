package cmd

import (
	"fmt"
	"github.com/ice1n36/inspector-gadget/gadget"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"os"
)

var (
	appPath    = ""
	gadgetPath = ""
	injectCmd  = &cobra.Command{
		Use:   "inject",
		Short: "inject frida-gadget in app",
		Long:  `injects frida-gadget and configuration into your apk (and soon to come ipa)`,
		Run: func(cmd *cobra.Command, args []string) {
			fs := afero.NewOsFs()
			err := gadget.NewInjector(fs).InjectGadget(appPath, gadgetPath)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
)

func init() {
	injectCmd.Flags().StringVarP(&appPath, "app", "a", "", "Path to APK/IPA/etc (requied)")
	injectCmd.MarkFlagRequired("app")
	injectCmd.Flags().StringVarP(&gadgetPath, "gadget", "g", "", "Path to frida-gadget to inject (requied)")
	injectCmd.MarkFlagRequired("gadget")
}
