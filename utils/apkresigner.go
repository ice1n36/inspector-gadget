package utils

// APKResigner is the interface for anything used to resign apks
type APKResigner interface {
	Resign(apk string)
}

type apkresigner struct {
	keystore   string
	keystorepw string
}

// NewAPKResigner creates a new instance of APKResigner
func NewAPKResigner() APKResigner {
	return &apkresigner{
		keystore:   "/put/something/here/for/now", // TODO: pull this from config
		keystorepw: "putsomethingherefornow",      // TODO: pull this from config
	}
}

func (a apkresigner) Resign(apk string) {
	//TODO
}
