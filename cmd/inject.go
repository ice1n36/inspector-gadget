package cmd

import (
	"fmt"
	"github.com/ice1n36/inspector-gadget/gadget"
	"github.com/ice1n36/inspector-gadget/utils"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"os"
)

var (
	appPath    = ""
	outputPath = ""
	gadgetPath = ""
	configPath = ""

	interactiveMode    = false
	nativeLibInjection = false
	injectCmd          = &cobra.Command{
		Use:   "inject",
		Short: "inject frida-gadget in app",
		Long:  `injects frida-gadget and configuration into your apk (and soon to come ipa)`,
		Run: func(cmd *cobra.Command, args []string) {
			fs := afero.NewOsFs()
			shell := utils.NewShell()
			i := gadget.NewInjector(fs, shell, interactiveMode, nativeLibInjection)
			err := i.InjectGadget(appPath, outputPath, gadgetPath, configPath)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
)

func init() {
	injectCmd.Flags().StringVarP(&appPath, "app", "a", "", "Path to APK/IPA/etc (required)")
	injectCmd.MarkFlagRequired("app")
	injectCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Output path (required)")
	injectCmd.MarkFlagRequired("output")
	injectCmd.Flags().StringVarP(&gadgetPath, "gadget", "g", "", "Path to frida-gadget to inject (required)")
	injectCmd.MarkFlagRequired("gadget")
	injectCmd.Flags().StringVarP(&configPath, "config", "c", "", "Path to frida-gadget config to inject (required)")
	injectCmd.MarkFlagRequired("config")
	injectCmd.Flags().BoolVarP(&interactiveMode, "interactive", "i", false, "Go through this in interactive mode")
	injectCmd.Flags().BoolVarP(&nativeLibInjection, "nativeinject", "n", false, "Inject into native library")
}
