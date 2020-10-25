package gadget

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/spf13/afero"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var (
	_libSearch  = "lib/.*\\.so"
	_arm64Seach = "arm64-v8a"
)

// Injector is the interface for things that inject gadgets
type Injector interface {
	InjectGadget(appPath string, gadgetPath string) error
}

type injector struct {
	fs afero.Fs
}

// NewInjector creates a new injector instance
func NewInjector(fs afero.Fs) Injector {
	return &injector{
		fs: fs,
	}
}

func (i *injector) InjectGadget(appPath string, gadgetPath string) error {
	fmt.Printf("Injecting Gadget @ %s into %s...\n", gadgetPath, appPath)

	if !strings.HasSuffix(appPath, ".apk") {
		return errors.New("currently, this only supports apks")
	}

	//tmpDir := afero.GetTempDir(i.fs, "")

	// get list of native libraries
	re := regexp.MustCompile(_libSearch)
	listOfAllContents, err := getListOfAllContents(appPath)
	if err != nil {
		return err
	}
	listOfLibs := re.FindAll([]byte(listOfAllContents), -1)

	// ask user if they would like to inject into a native library (if there are native libraries)
	var nativeLibInjection bool
	if len(listOfLibs) > 0 {
		fmt.Print("Inject into a native library? [Y/n]: ")
		reader := bufio.NewReader(os.Stdin)
		response, err := reader.ReadString('\n')
		if err != nil {
			errors.New("Unable to read user input")
		}
		response = strings.ToLower(strings.TrimSpace(response))
		if len(response) == 0 || response == "y" || response == "yes" {
			nativeLibInjection = true
		} else if response == "n" || response == "no" {
			nativeLibInjection = false
		} else {
			return errors.New("Unknown input")
		}

	}

	if nativeLibInjection {
		var nativeLibToInjectTo string
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
		fmt.Printf("lib choosen: %s\n", nativeLibToInjectTo)

		// apktool d -rs -o output/${packagename} downloads/apks/${packagename}/${apkname}

		// add libfridagadget.so and config (as libfridagegadget.config.so) to lib/arm64-v8a directory
		// add gadget to native library
	} else {
		// add System.LoadLibary

		errors.New("Dex injection is not yet supported")
	}

	return nil
}

func getListOfAllContents(apkPath string) (string, error) {
	out, err := exec.Command("unzip", "-l", apkPath).Output()
	if err != nil {
		return "", errors.New("Unable to get contents of apk")
	}
	return string(out), nil
}
