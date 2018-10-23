package app

import (
	"git.oxagile.com/mmgfusppointment/factory/factory"
	"log"
)

type Application struct {
	patientsLoader factory.Loader
	appointmentsLoader factory.Loader
}

func New(f factory.Synchronizer) *Application {
	app := &Application{}
	app.patientsLoader = f.CreatePatientsLoader()
	app.appointmentsLoader = f.CreateAppointmentsLoader()

	return app
}

func (a *Application) Sync() {
	if err := a.patientsLoader.Load(); err != nil {
		log.Printf("[ERROR]failed to load patients: `%s`", err.Error())
	}

	if err := a.appointmentsLoader.Load(); err != nil {
		log.Printf("[ERROR]failed to load appointments: `%s`", err.Error())
	}
}