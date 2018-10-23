package softdent

import "log"

type PatientsLoader struct {}

func (pl *PatientsLoader) Load() error {
	log.Printf("loading softdent patients..done")
	return nil
}

type AppointmentsLoader struct{}

func (pl *AppointmentsLoader) Load() error {
	log.Printf("loading softdent appointments..done")
	return nil
}