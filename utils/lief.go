package utils

// NativeLibInjector is the interface for anything that can be used to inject a payload library into another library
// to make it a dependency
type NativeLibInjector interface {
	InjectIntoLibrary(payload string, target string) error
}

type liefnativelibinjector struct {
}

// NewNativeLibInjector creates a new instance of NativeLibInjector
func NewNativeLibInjector() NativeLibInjector {
	return &liefnativelibinjector{}
}

// InjectIntoLibrary injects payload library into the target library as a dependency using lief
func (n liefnativelibinjector) InjectIntoLibrary(payload string, target string) error {
	// TODO
	return nil
}
