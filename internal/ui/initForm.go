package ui

import (
	"github.com/charmbracelet/huh/spinner"
)

var InitSpinner spinner.Spinner = *spinner.New().
	Title("Initializing...")

var FetchSpinner spinner.Spinner = *spinner.New().
	Title("Fetching all installed apps...")
