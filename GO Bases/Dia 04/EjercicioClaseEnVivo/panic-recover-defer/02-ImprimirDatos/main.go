package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// ------------- Structures -------------
type Customer struct {
	DNI    int
	Name   string
	Docket int // *Legajo*
	Phone  string
	Home   string
}
type CustomerController struct {
	Customers   []Customer
	File        *os.File
	ErrorStatus error
}

// ------------- Methods -------------
func (c *CustomerController) Init(fileName string) (err error) {
	defer func() {
		c.File.Close()
		existPanic := recover()
		if existPanic != nil {
			fmt.Println(existPanic)
			fmt.Println("Error: ", err)
		}

		fmt.Println("ejecución finalizada")
	}()

	c.File, err = os.Open(fileName)
	c.ErrorStatus = err
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}

	var fileString string
	scanner := bufio.NewScanner(c.File)
	for scanner.Scan() {
		fileString += scanner.Text()
	}

	err = json.Unmarshal([]byte(fileString), &c.Customers)
	if err != nil {
		panic("La estructura JSON dentro del archivo txt es invalido.")
	}

	return
}
func (c CustomerController) PrintCustomers() {
	for i, customer := range c.Customers {
		fmt.Printf("------------------ Customer #%d ------------------ \n", i+1)
		fmt.Println("DNI   : ", customer.DNI)
		fmt.Println("Name  : ", customer.Name)
		fmt.Println("Docket: ", customer.Docket)
		fmt.Println("Phone : ", customer.Phone)
		fmt.Println("Home  : ", customer.Home)
	}
}

func main() {
	var customerController CustomerController = CustomerController{}
	customerController.Init("customers.txt")
	customerController.PrintCustomers()

}
