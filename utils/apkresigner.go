package utils

import (
	"fmt"
)

// APKResigner is the interface for anything used to resign apks
type APKResigner interface {
	Resign(apk string) error
}

type apkresigner struct {
	keystore   string
	keystorepw string
	shell      Shell
}

// NewAPKResigner creates a new instance of APKResigner
func NewAPKResigner(shell Shell) APKResigner {
	return &apkresigner{
		keystore:   "/put/something/here/for/now", // TODO: pull this from config
		keystorepw: "putsomethingherefornow",      // TODO: pull this from config
		shell:      shell,
	}
}

func (a apkresigner) Resign(apk string) error {
	_, err := a.shell.Exec("java", "-jar", "bin/uber-apk-signer-1.1.0.jar", "-apks", apk)
	if err != nil {
		fmt.Errorf("Error resigning apk: %v", err)
	}
	return nil
}
