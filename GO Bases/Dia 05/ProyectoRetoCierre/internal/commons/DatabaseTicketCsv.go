package commons

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type DatabaseTicketCsv struct {
	IDatabaseTicket
	File *os.File
}

func (d *DatabaseTicketCsv) GetTickets() (tickets []Ticket, err error) {
	defer func() {
		d.File.Close()
		existPanic := recover()
		if existPanic != nil {
			fmt.Println("Metodo GetTickets [DatabaseTicketCsv]: Finalizo mal.")
			return
		}
	}()

	d.File, err = os.Open("tickets.csv")
	if err != nil {
		err = fmt.Errorf("el archivo indicado no fue encontrado o está dañado")
		return
	}

	scanner := bufio.NewScanner(d.File)
	for scanner.Scan() {
		fileString := scanner.Text()
		values := strings.Split(fileString, ",")
		id, err1 := strconv.Atoi(values[0])
		price, err2 := strconv.ParseFloat(values[5], 64)
		flightTime, err3 := time.Parse("15:04", values[4])

		if err1 != nil {
			err = err1
			return
		}
		if err2 != nil {
			err = err2
			return
		}

		if err3 != nil {
			err = err3
			return
		}

		var ticket Ticket = Ticket{
			Id:                 id,
			Name:               values[1],
			Email:              values[2],
			DestinationCountry: values[3],
			FlightTime:         flightTime,
			Price:              price,
		}
		//Agregamos el ticket al array
		tickets = append(tickets, ticket)
	}

	return
}
