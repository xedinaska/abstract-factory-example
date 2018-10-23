package dentrix

import "log"

type PatientsLoader struct {}
func (pl *PatientsLoader) Load() error {
	log.Printf("loading dentrix patients..done")
	return nil
}

type AppointmentsLoader struct{}
func (pl *AppointmentsLoader) Load() error {
	log.Printf("loading dentrix appointments..done")
	return nil
}