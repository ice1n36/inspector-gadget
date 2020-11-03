package utils

import (
	"fmt"
	"github.com/spf13/afero"
)

// APKTool is an interface for all apktool functionality
type APKTool interface {
	Decode(apk string, force, nores, nosrc bool) error
	Build(output string, force bool) error
	CopyToDecoded(src, dst string) error
	GetIntermediateDir() string
}

type apktoolcmdlineimpl struct {
	fs              afero.Fs
	shell           Shell
	intermediatedir string
}

// NewAPKTool creates a new instance of a command line implementation of APKTool
func NewAPKTool(fs afero.Fs, shell Shell) (APKTool, error) {
	tmpDir, err := afero.TempDir(fs, afero.GetTempDir(fs, ""), "")
	if err != nil {
		return nil, fmt.Errorf("Could not create temp dir: %v", err)
	}
	return &apktoolcmdlineimpl{
		fs:              fs,
		shell:           shell,
		intermediatedir: tmpDir + "/out/",
	}, nil
}

// Decode is a command line/shell implementation of 'apktool d'
func (a apktoolcmdlineimpl) Decode(apk string, force, nores, nosrc bool) error {
	_, err := a.shell.Exec("apktool", "d", "-rs", "-o", a.intermediatedir, apk)
	if err != nil {
		return fmt.Errorf("Error extracting apk %v", err)
	}
	return nil
}

// Build is a command line/shell implementation of 'apktool b'
func (a apktoolcmdlineimpl) Build(output string, force bool) error {
	_, err := a.shell.Exec("apktool", "b", a.intermediatedir, "-o", output)
	if err != nil {
		return fmt.Errorf("Error building apk %v", err)
	}
	return nil
}

// CopyToDecoded is a command line/shell implementation of cp that copies artifacts
// into the decoded intermediate dir
func (a apktoolcmdlineimpl) CopyToDecoded(src, dst string) error {
	_, err := a.shell.Exec("cp", src, a.intermediatedir+dst)
	if err != nil {
		return fmt.Errorf("Error copying %s artifact into decoded intermediate dir: %v", src, err)
	}
	return nil
}

// GetIntermediateDir simply returns the intermediate dir used
func (a apktoolcmdlineimpl) GetIntermediateDir() string {
	return a.intermediatedir
}
