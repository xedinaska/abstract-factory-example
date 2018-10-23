package main

import (
	"git.oxagile.com/mmgfusppointment/factory/app"
	"git.oxagile.com/mmgfusppointment/factory/factory"
	"os"
)

func main() {
	pms := os.Getenv("pms")

	var f factory.Synchronizer
	if pms == "dentrix" {
		f = &factory.DentrixSynchronizer{}
	} else {
		f = &factory.SoftdentSynchronizer{}
	}

	a := app.New(f)
	a.Sync()
}