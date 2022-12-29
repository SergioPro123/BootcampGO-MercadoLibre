package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
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
	Customers    []Customer
	FileName     string
	File         *os.File
	ErrorsStatus []error
}
type ErrorCustomer struct {
	Message string
	Code    int
}

// ------------- Methods -------------
func (e ErrorCustomer) Error() string {
	return fmt.Sprintf("Error '%s' with code %d", e.Message, e.Code)
}

func (c Customer) ValidateDates() (validate bool, err []error) {
	//Realizamos validaciones
	if c.DNI == 0 {
		err = append(err, ErrorCustomer{
			Message: "DNI cannot be zero.",
			Code:    1,
		})
	}
	if c.Name == "" {
		err = append(err, ErrorCustomer{
			Message: "Name cannot be empty.",
			Code:    2,
		})
	}
	if c.Docket == 0 {
		err = append(err, ErrorCustomer{
			Message: "Docket cannot be zero.",
			Code:    3,
		})
	}
	if c.DNI == 0 {
		err = append(err, ErrorCustomer{
			Message: "DNI cannot be zero.",
			Code:    4,
		})
	}
	if c.Phone == "" {
		err = append(err, ErrorCustomer{
			Message: "Phone cannot be empty.",
			Code:    5,
		})
	}
	if c.Home == "" {
		err = append(err, ErrorCustomer{
			Message: "Home cannot be empty.",
			Code:    6,
		})
	}

	if err == nil {
		validate = true
	}

	return
}

func (c *CustomerController) Init(fileName string) (err error) {
	defer func() {
		c.File.Close()
		existPanic := recover()
		if existPanic != nil {
			fmt.Println("Metodo Init: Finalizo mal.")
			return
		}
	}()
	c.FileName = fileName

	c.File, err = os.Open(c.FileName)
	if err != nil {
		c.ErrorsStatus = append(c.ErrorsStatus, err)
		panic("el archivo indicado no fue encontrado o está dañado")
	}

	var fileString string
	scanner := bufio.NewScanner(c.File)
	for scanner.Scan() {
		fileString += scanner.Text()
	}

	err = json.Unmarshal([]byte(fileString), &c.Customers)

	if err != nil {
		c.ErrorsStatus = append(c.ErrorsStatus, err)
		panic("La estructura JSON dentro del archivo txt es invalido.")
	}
	return
}
func (c CustomerController) ValidateExistCustomer(customer Customer) (exist bool) {
	//Validamos si el cliente ya existe, con la propiedad "DNI"
	for _, customerLoop := range c.Customers {
		exist = customerLoop.DNI == customer.DNI
		if exist {
			exist = true
			break
		}
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
func (c CustomerController) PrintErrors() {
	fmt.Println("------------------ Errores CustomerController  ------------------")
	for _, errorStatus := range c.ErrorsStatus {
		fmt.Println(errorStatus.Error())
	}
}
func (c *CustomerController) WriteCustomer(customer Customer) (err error) {
	defer func() {
		c.File.Close()
		existPanic := recover()
		if existPanic != nil {
			fmt.Println(existPanic)
		}
	}()

	existCustomer := c.ValidateExistCustomer(customer)
	if existCustomer {
		c.ErrorsStatus = append(c.ErrorsStatus, ErrorCustomer{
			Message: "Client already exists.",
			Code:    21,
		})
		panic(c.ErrorsStatus[len(c.ErrorsStatus)-1])
	}

	validateDates, errs := customer.ValidateDates()
	if !validateDates {
		c.ErrorsStatus = append(c.ErrorsStatus, errs...)
		panic(c.ErrorsStatus[len(c.ErrorsStatus)-1])
	}

	//Escribimos el cliente en el fichero
	expectedCustomers := append(c.Customers, customer)

	c.File, err = os.OpenFile(c.FileName, os.O_WRONLY, os.ModeAppend)
	if err != nil {
		c.ErrorsStatus = append(c.ErrorsStatus, err)
		panic("el archivo indicado no fue encontrado o está dañado")
	}

	customersMarshal, err := json.Marshal(expectedCustomers)
	if err != nil {
		c.ErrorsStatus = append(c.ErrorsStatus, err)
		panic("Hubo un error al realizar Marshal sobre la estructura de clientes.")
	}

	fmt.Println(string(customersMarshal))

	w := bufio.NewWriter(c.File)
	w.WriteString(string(customersMarshal))

	err = w.Flush()
	if err != nil {
		c.ErrorsStatus = append(c.ErrorsStatus, err)
		panic("Hubo un error al escribir los clientes.")
	}

	c.Customers = append(c.Customers, customer)
	return
}

func main() {
	defer func() {
		fmt.Println("Fin de la ejecución.")
	}()
	var customerController CustomerController = CustomerController{}
	err := customerController.Init("customers.txt")
	if err != nil {
		println(err.Error())
		log.Fatal()
	}
	//Leemos todos los clientes
	customerController.PrintCustomers()
	//Escribimos un cliente valido
	customerController.WriteCustomer(Customer{
		DNI:    5,
		Name:   "",
		Docket: 0,
		Phone:  "+57 333333",
		Home:   "Calle de prueba",
	})

	customerController.PrintErrors()

}
