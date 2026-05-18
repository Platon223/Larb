package ui

import (
	"github.com/briandowns/spinner"
	"time"
)

// Defines a spinner for every command

func WithSpinner(msg string, cmd func() error) error {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Suffix = " " + msg
	s.Color("green", "bold")
	s.Start()
	err := cmd()
	s.Stop()

	return err
}
