package utils

import (
	"fmt"
)

const (
	_nativeInjectTemplate = `import lief;libnative = lief.parse("%s");libnative.add_library("%s");libnative.write("%s")`
)

// NativeLibInjector is the interface for anything that can be used to inject a payload library into another library
// to make it a dependency
type NativeLibInjector interface {
	InjectIntoLibrary(payload string, target string) error
}

type liefnativelibinjector struct {
	shell Shell
}

// NewNativeLibInjector creates a new instance of NativeLibInjector
func NewNativeLibInjector(shell Shell) NativeLibInjector {
	return &liefnativelibinjector{
		shell: shell,
	}
}

// InjectIntoLibrary injects payload library into the target library as a dependency using lief
func (n liefnativelibinjector) InjectIntoLibrary(payload string, target string) error {
	pscript := fmt.Sprintf(_nativeInjectTemplate, target, payload, target)
	_, err := n.shell.Exec("python", "-c", pscript)
	if err != nil {
		return fmt.Errorf("Error doing lief native injection: %v", err)
	}

	return nil
}
