package utils

import (
	"os/exec"
)

type Shell interface {
	Exec(name string, arg ...string) ([]byte, error)
}

type shellImpl struct {
}

// NewShell creates a new instance of shellImpl
func NewShell() Shell {
	return &shellImpl{}
}

// Exec is the implementation of Shell's Exec command using exec.Command
func (s shellImpl) Exec(name string, arg ...string) ([]byte, error) {
	out, err := exec.Command(name, arg...).Output()
	if err != nil {
		return nil, err
	}
	return out, nil
}
