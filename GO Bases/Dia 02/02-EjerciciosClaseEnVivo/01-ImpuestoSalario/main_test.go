package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testingData struct {
	salary                    float64
	taxRateExpected           float64
	percentageTaxRateExpected int
	taxRate                   float64
	percentageTaxRate         int
}

func TestImpuestoSalario(t *testing.T) {
	//Arrange.
	var testsData = []testingData{
		{
			salary:                    30000,
			taxRateExpected:           0,
			percentageTaxRateExpected: 0,
		},
		{
			salary:                    70000,
			taxRateExpected:           11900,
			percentageTaxRateExpected: 17,
		},
		{
			salary:                    190000,
			taxRateExpected:           51300,
			percentageTaxRateExpected: 27,
		},
	}
	//Act.
	for i, testData := range testsData {
		testsData[i].taxRate, testsData[i].percentageTaxRate = ImpuestoSalario(testData.salary)
	}
	//Assert.
	for _, testData := range testsData {
		//Validamos que el impuesto devuelto, se igual al esperado
		assert.Equalf(t, testData.taxRateExpected, testData.taxRate,
			"Para el salario de $%.2f se esperaba un impuesto de (%.2f) pero la funcion retorno (%.2f)",
			testData.salary, testData.taxRateExpected, testData.taxRate)
		//Validamos que el porcentaje del impuesto devuelto, se igual al esperado
		assert.Equalf(t, testData.percentageTaxRateExpected, testData.percentageTaxRate,
			"Para el salario de $%.2f se esperaba un porcentaje de impuesto de (%d) pero la funcion retorno (%d)",
			testData.salary, testData.percentageTaxRateExpected, testData.percentageTaxRate)
	}

	assert.Equal(t, 1, 1)
}
