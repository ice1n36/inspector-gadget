package gadget

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/ice1n36/inspector-gadget/utils"
	"github.com/spf13/afero"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	_libSearch          = "lib/.*\\.so"
	_arm64Seach         = "arm64-v8a"
	_fridaGadgetLibName = "libfridagadget.so"
)

// Injector is the interface for things that inject gadgets
type Injector interface {
	InjectGadget(appPath string, outputPath string, gadgetPath string, configPath string) error
}

type injector struct {
	fs                 afero.Fs
	shell              utils.Shell
	interactiveMode    bool
	nativeLibInjection bool
	apkresigner        utils.APKResigner
	nativeLibInjector  utils.NativeLibInjector
}

// NewInjector creates a new injector instance
func NewInjector(fs afero.Fs, shell utils.Shell, interactiveMode bool, nativeLibInjection bool) Injector {
	return &injector{
		fs:                 fs,
		shell:              shell,
		interactiveMode:    interactiveMode,
		nativeLibInjection: nativeLibInjection,
		apkresigner:        utils.NewAPKResigner(shell),
		nativeLibInjector:  utils.NewNativeLibInjector(shell),
	}
}

func (i *injector) InjectGadget(appPath string, outputPath string, gadgetPath string, configPath string) error {
	fmt.Printf("Injecting Gadget @ %s into %s...\n", gadgetPath, appPath)

	if !strings.HasSuffix(appPath, ".apk") {
		return errors.New("currently, this only supports apks")
	}

	// get list of native libraries
	listOfLibs, err := i.getListOfAllNativeLibs(appPath)
	if err != nil {
		return errors.New("error getting list of native libs")
	}

	var nativeLibToInjectTo string
	if i.interactiveMode {
		// ask user if they would like to inject into a native library (if there are native libraries)
		if len(listOfLibs) > 0 {
			fmt.Print("Inject into a native library? [Y/n]: ")
			reader := bufio.NewReader(os.Stdin)
			response, err := reader.ReadString('\n')
			if err != nil {
				errors.New("Unable to read user input")
			}
			response = strings.ToLower(strings.TrimSpace(response))
			if len(response) == 0 || response == "y" || response == "yes" {
				i.nativeLibInjection = true
			} else if response == "n" || response == "no" {
				i.nativeLibInjection = false
			} else {
				return errors.New("Unknown input")
			}

		}
	}

	if i.nativeLibInjection {
		if i.interactiveMode {
			for i, lib := range listOfLibs {
				fmt.Printf("[%d] %s\n", i, string(lib))
			}
			fmt.Println("Which library would you like to inject into?")
			var chosenLib string
			fmt.Scanln(&chosenLib)
			var index int
			if chosenLib == "" {
				index = 0
			} else {
				var err error
				index, err = strconv.Atoi(chosenLib)
				if err != nil {
					errors.New("Unknown library choice")
				}
			}
			nativeLibToInjectTo = string(listOfLibs[index])
		} else {
			nativeLibToInjectTo = string(listOfLibs[0])
		}
		fmt.Printf("injecting into native library: %s...\n", nativeLibToInjectTo)

		apktool, err := utils.NewAPKTool(i.fs, i.shell)
		if err != nil {
			return fmt.Errorf("APKTool creation failed: %v", err)
		}
		// apktool d -rs -o output/${packagename} downloads/apks/${packagename}/${apkname}
		err = apktool.Decode(appPath, false, true, true)
		if err != nil {
			return fmt.Errorf("Error extracting apk %v", err)
		}

		// add libfridagadget.so and config (as libfridagegadget.config.so) to lib/arm64-v8a directory
		err = apktool.CopyToDecoded(gadgetPath, "lib/arm64-v8a/libfridagadget.so")
		if err != nil {
			return fmt.Errorf("Error embedding frida gadget %v", err)
		}
		err = apktool.CopyToDecoded(configPath, "lib/arm64-v8a/libfridagadget.config.so")
		if err != nil {
			return fmt.Errorf("Error embedding frida gadget config %v", err)
		}

		i.nativeLibInjector.InjectIntoLibrary(
			_fridaGadgetLibName,
			apktool.GetIntermediateDir()+nativeLibToInjectTo)

		err = apktool.Build(outputPath, false)
		if err != nil {
			return fmt.Errorf("Error building apk %v", err)
		}

	} else {
		// add System.LoadLibary

		// apktool d -r -o output/${packagename} downloads/apks/${packagename}/${apkname}

		return errors.New("Dex injection is not yet supported")
	}
	i.apkresigner.Resign(outputPath)

	fmt.Println("Complete!  APK: " + outputPath)
	return nil
}

func (i injector) getListOfAllNativeLibs(apkPath string) ([][]byte, error) {
	re := regexp.MustCompile(_libSearch)
	listOfAllContents, err := i.getListOfAllContents(apkPath)
	if err != nil {
		return nil, err
	}
	return re.FindAll([]byte(listOfAllContents), -1), nil
}

func (i injector) getListOfAllContents(apkPath string) (string, error) {
	out, err := i.shell.Exec("unzip", "-l", apkPath)
	if err != nil {
		return "", errors.New("Unable to get contents of apk")
	}
	return string(out), nil
}
