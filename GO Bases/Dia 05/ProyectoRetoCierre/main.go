package main

import (
	"fmt"
	"log"

	"github.com/bootcamp-go/desafio-go-bases/internal/commons"
)

func main() {
	destino := "argentina"
	//Instanciamos la clase que se encarga de obtener los datos del CSV
	var databaseTicket commons.IDatabaseTicket = &commons.DatabaseTicketCsv{}

	ticketController := commons.TicketController{}
	err := ticketController.Init(databaseTicket)
	if err != nil {
		fmt.Println("Ocurrio un error: ", err.Error())
		log.Fatal(err)
	}

	//Obtenemos la cantidad de personas que viajan a un pais determinado
	totalTickes, err := ticketController.GetTotalTickets(destino)
	if err != nil {
		fmt.Println("Ocurrio un error: ", err.Error())
		log.Fatal(err)
	}
	fmt.Printf("Cantidad de personas que viajan a %s : %d\n", destino, totalTickes)

	//Obtenemos el numero de personas que viaja en cada periodo
	fmt.Println("--------------- Personas según el periodo ---------------")
	periods := []string{commons.MADRUGADA, commons.MANANA, commons.TARDE, commons.NOCHE}
	for _, period := range periods {
		totalPeople, err := ticketController.GetCountByPeriod(period)
		if err != nil {
			fmt.Printf("Error para la jornada '%s': %s \n", period, err.Error())
			continue
		}
		fmt.Printf("Personas en jornada '%s' : %d \n", period, totalPeople)
	}
	//Calculamos el promedio de personas que viajan a un pais determinado
	fmt.Println("--------------- Promedio en un un pais determinado ---------------")
	avarage, err := ticketController.GetAvarageDestination(destino)
	if err != nil {
		fmt.Println("Ocurrio un error: ", err.Error())
		log.Fatal(err)
	}
	fmt.Printf("Promedio de personas en %s : %.2f \n", destino, avarage)

}
