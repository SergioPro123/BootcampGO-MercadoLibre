package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type orquestadorOperacionesStruct struct {
	notes           []float64
	min             float64
	max             float64
	average         float64
	err             error
	minExpected     float64
	maxExpected     float64
	averageExpected float64
}

func TestOrquestadorOperaciones(t *testing.T) {
	//Arrange.
	var testsData = []orquestadorOperacionesStruct{
		{
			notes:           []float64{2.4, 4.5, 7.8, 9.8, 1.2},
			minExpected:     1.2,
			maxExpected:     9.8,
			averageExpected: 5.14,
		},
		{
			notes:           []float64{3.6, 2.9, 0.2, 87.4, 55.2},
			minExpected:     0.2,
			maxExpected:     87.4,
			averageExpected: 29.860000000000003,
		},
	}
	//Act.
	for i, testData := range testsData {

		var operations = []string{minimo, promedio, maximo, "error"}
		//Recorremos cada operación

		for _, operation := range operations {
			operationFunction, err := OrquestadorOperaciones(operation)

			testsData[i].err = err
			result := operationFunction(testData.notes...)

			switch operation {
			case minimo:
				testsData[i].min = result
			case promedio:
				testsData[i].average = result
			case maximo:
				testsData[i].max = result
			}
		}
	}
	//Assert.
	for _, testData := range testsData {
		//Validamos que el valor minimo corresponda con el esperado
		assert.Equalf(t, testData.minExpected, testData.min,
			"Para las notas %f el valor minimo se esperaba de (%.5f) pero la funcion retorno (%.5f)",
			testData.notes, testData.minExpected, testData.min)
		//Validamos que el valor maximo corresponda con el esperado
		assert.Equalf(t, testData.maxExpected, testData.max,
			"Para las notas %f el valor maximo se esperaba de (%.5f) pero la funcion retorno (%.5f)",
			testData.notes, testData.maxExpected, testData.max)
		//Validamos que el valor promedio corresponda con el esperado
		assert.Equalf(t, testData.averageExpected, testData.average,
			"Para las notas %f el valor promedio se esperaba de (%.5f) pero la funcion retorno (%.5f)",
			testData.notes, testData.averageExpected, testData.average)
		//Validamos que el valor error corresponda con el esperado
	}

	assert.Equal(t, 1, 1)
}
