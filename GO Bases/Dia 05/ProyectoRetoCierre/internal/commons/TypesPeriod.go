package commons

import (
	"fmt"
	"time"
)

const (
	MADRUGADA = "A" // (00:00 - 06:59)
	MANANA    = "B" // (07:00 - 11:59)
	TARDE     = "C" // (13:00 - 19:59)
	NOCHE     = "D" // (20:00 - 23:59)
)

type OperationPeriod func() (startTime, endTime time.Time)

type TypesPeriod struct {
}

func (t TypesPeriod) OrchestratorPeriod(typePeriod string) (operationPeriod OperationPeriod, err error) {
	switch typePeriod {
	case MADRUGADA:
		operationPeriod = t.GetPeriodMadrugada
	case MANANA:
		operationPeriod = t.GetPeriodManana
	case TARDE:
		operationPeriod = t.GetPeriodTarde
	case NOCHE:
		operationPeriod = t.GetPeriodNoche
	default:
		err = fmt.Errorf("tipo de periodo no valido")
	}
	return
}

func (t TypesPeriod) GetPeriodMadrugada() (startTime, endTime time.Time) {
	startTime, _ = time.Parse("15:04", "00:00")
	endTime, _ = time.Parse("15:04", "06:59")
	return
}
func (t TypesPeriod) GetPeriodManana() (startTime, endTime time.Time) {
	startTime, _ = time.Parse("15:04", "07:00")
	endTime, _ = time.Parse("15:04", "11:59")
	return
}

func (t TypesPeriod) GetPeriodTarde() (startTime, endTime time.Time) {
	startTime, _ = time.Parse("15:04", "13:00")
	endTime, _ = time.Parse("15:04", "19:59")
	return
}

func (t TypesPeriod) GetPeriodNoche() (startTime, endTime time.Time) {
	startTime, _ = time.Parse("15:04", "20:00")
	endTime, _ = time.Parse("15:04", "23:59")
	return
}
