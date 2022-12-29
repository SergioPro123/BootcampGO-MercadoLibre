package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testingData struct {
	notes           []float64
	avarage         float64
	avarageExpected float64
}

func TestCalcularNotas(t *testing.T) {
	//Arrange.
	var testsData = []testingData{
		{
			notes:           []float64{2.3, 2.3, 2.3},
			avarageExpected: 2.3,
		},
		{
			notes:           []float64{6, 6, 6.5, 5.5},
			avarageExpected: 6.0,
		},
		{
			notes: []float64{3.4, 5.7, 9.7, 3.45},

			avarageExpected: 5.5625,
		},
	}
	//Act.
	for i, testData := range testsData {
		testsData[i].avarage = calcularNotas(testData.notes...)
	}
	//Assert.
	for _, testData := range testsData {
		//Validamos que el promedio devuelto corresponda con el esperado
		assert.Equalf(t, testData.avarageExpected, testData.avarage,
			"Para las notas %.2f se esperaba un promedio de (%.2f) pero la funcion retorno (%.2f)",
			testData.notes, testData.avarageExpected, testData.avarage)
	}

	assert.Equal(t, 1, 1)
}
