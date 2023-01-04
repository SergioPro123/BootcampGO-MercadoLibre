package commons

type IDatabaseTicket interface {
	GetTickets() ([]Ticket, error)
}
