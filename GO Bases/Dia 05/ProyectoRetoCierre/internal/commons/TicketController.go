package commons

import (
	"fmt"
	"strings"
)

type TicketController struct {
	Tickets []Ticket
	err     error
}

func (t *TicketController) Init(databaseTicket IDatabaseTicket) (err error) {
	defer func() {
		existPanic := recover()
		if existPanic != nil {
			fmt.Println("Metodo Init: Finalizo mal.")
			return
		}
	}()
	t.Tickets, err = databaseTicket.GetTickets()
	if err != nil {
		return
	}

	return
}

func (t *TicketController) GetTotalTickets(destination string) (totalTickets int, err error) {
	if t.err != nil {
		err = fmt.Errorf("no se puede utilizar el metodo GetTotalTickets debido a que el metodo Init no se ejecuto como se esperaba. error: %w", t.err)
		return
	}

	for _, ticket := range t.Tickets {
		if strings.EqualFold(strings.ToUpper(ticket.DestinationCountry), strings.ToUpper(destination)) {
			totalTickets++
		}
	}

	return
}

func (t *TicketController) GetCountByPeriod(typePeriod string) (totalPeople int, err error) {
	if t.err != nil {
		err = fmt.Errorf("no se puede utilizar el metodo GetTotalTickets debido a que el metodo Init no se ejecuto como se esperaba. error: %w", t.err)
		return
	}
	typesPeriod := TypesPeriod{}
	operationType, err := typesPeriod.OrchestratorPeriod(typePeriod)
	if err != nil {
		return
	}
	startTime, endTime := operationType()

	for _, ticket := range t.Tickets {
		if (ticket.FlightTime.Before(endTime) && ticket.FlightTime.After(startTime)) ||
			ticket.FlightTime.Equal(startTime) ||
			ticket.FlightTime.Equal(endTime) {
			totalPeople++
		}
	}

	return
}
func (t *TicketController) GetAvarageDestination(destination string) (avarage float64, err error) {
	if t.err != nil {
		err = fmt.Errorf("no se puede utilizar el metodo GetTotalTickets debido a que el metodo Init no se ejecuto como se esperaba. error: %w", t.err)
		return
	}
	totalTickes, err := t.GetTotalTickets(destination)
	if err != nil {
		return
	}
	avarage = float64(totalTickes) / float64(len(t.Tickets))
	return
}
