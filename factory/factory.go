package factory

import (
	"git.oxagile.com/mmgfusppointment/factory/dentrix"
	"git.oxagile.com/mmgfusppointment/factory/softdent"
)

//Product
type Loader interface {
	Load() error
}

//Factory
type Synchronizer interface {
	CreatePatientsLoader() Loader
	CreateAppointmentsLoader() Loader
}

type DentrixSynchronizer struct {}

func (s *DentrixSynchronizer) CreatePatientsLoader() Loader {
	return &dentrix.PatientsLoader{}
}

func (s *DentrixSynchronizer) CreateAppointmentsLoader() Loader {
	return &dentrix.AppointmentsLoader{}
}

type SoftdentSynchronizer struct {}

func (s *SoftdentSynchronizer) CreatePatientsLoader() Loader {
	return &softdent.PatientsLoader{}
}

func (s *SoftdentSynchronizer) CreateAppointmentsLoader() Loader {
	return &softdent.AppointmentsLoader{}
}