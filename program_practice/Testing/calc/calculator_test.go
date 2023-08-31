package calc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {

	//Arrange
	num1 := 2.0
	num2 := 3.0
	expectedResult := 5.0

	//Action
	obtainedResult := Add(num1, num2)

	//Assert
	assert.Equal(t, expectedResult, obtainedResult)

	/* con testing nativo:
	if obtainedResult != expectedResult {
			t.Fatal("Add did not return 5 as expected")
		}
	*/
}
