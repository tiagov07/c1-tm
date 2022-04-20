package calculator

// import testing package
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	//data initialization
	num1 := 3
	num2 := 5
	wantedResult := 8

	result := Add(num1, num2)

	//answer validation
	// if result != wantedResult {
	// 	t.Errorf(`wrong answer, the function give us %v and the result should be %v`, result, wantedResult)
	// }

	//validate using testify
	assert.Equal(t, wantedResult, result, "must be equal")
}
