package main

import (
	"time"

	"github.com/chelnak/ysmrr"
	"github.com/chelnak/ysmrr/pkg/colors"
)

func main() {
	// Create a new spinner manager
	sm := ysmrr.NewSpinnerManager[colors.Color]()

	// Set up our spinner
	downloading := sm.AddSpinner("Downloading...")

	// Start the spinners that have been added to the group
	sm.Start()
	defer sm.Stop()

	// Set downloading to complete
	time.Sleep(2 * time.Second)
	downloading.Complete()
}
