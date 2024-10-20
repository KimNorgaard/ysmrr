package main

import (
	"time"

	"github.com/chelnak/ysmrr"
	"github.com/chelnak/ysmrr/pkg/colors"
)

type App struct {
	spinnerManager ysmrr.SpinnerManager[colors.Color]
}

func (a *App) doMoreWork() {
	s := a.spinnerManager.AddSpinner("Loading...DoMoreWork()")
	time.Sleep(time.Second * 5)
	s.CompleteWithMessage("Done")
}

func (a *App) doWork() {
	s := a.spinnerManager.AddSpinner("Loading...DoWork()")
	time.Sleep(time.Second * 5)
	s.CompleteWithMessage("It's Done")

	a.doMoreWork()
}

func (a *App) Run() {
	a.spinnerManager.Start()
	s := a.spinnerManager.AddSpinner("Loading...Run()")
	a.doWork()
	s.CompleteWithMessage("Loading...Run() Done")
}

func (a *App) Stop() {
	defer a.spinnerManager.Stop()
}

func NewApp() *App {
	return &App{
		spinnerManager: ysmrr.NewSpinnerManager[colors.Color](),
	}
}

func main() {
	app := NewApp()
	app.Run()
	app.Stop()
}
