package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testingData struct {
	minutesWorked int
	category      string
	salaryExpect  float64
	salary        float64
}

func TestCalcularSalario(t *testing.T) {
	//Arrange.
	var minutesWorked int = 1200
	var testsData = []testingData{
		{
			minutesWorked: minutesWorked,
			category:      "A",
			salaryExpect:  90000.0,
		},
		{
			minutesWorked: minutesWorked,
			category:      "B",
			salaryExpect:  36000.0,
		},
		{
			minutesWorked: minutesWorked,
			category:      "C",
			salaryExpect:  20000.0,
		},
		{
			minutesWorked: minutesWorked,
			category:      "XXXXXX",
			salaryExpect:  0,
		},
	}
	//Act.
	for i, testData := range testsData {
		testsData[i].salary = CalcularSalario(testData.minutesWorked, testData.category)
	}
	//Assert.
	for _, testData := range testsData {
		//Validamos que el salario devuelto corresponda con el esperado
		assert.Equalf(t, testData.salaryExpect, testData.salary,
			"Para la categoria %s con %d minutos trabajados se esperaba un salario de (%.5f) pero la funcion retorno (%.5f)",
			testData.category, minutesWorked, testData.salaryExpect, testData.salary)
	}

	assert.Equal(t, 1, 1)
}
