package factory_method

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func compute(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}

func TestOperatorFactory(t *testing.T) {
	var (
		factory OperatorFactory
	)
	factory = MinusOperatorFactory{}
	assert.Equal(t,compute(factory,2,1),1)

	factory = PlusOperatorFactory{}
	assert.Equal(t,compute(factory,1,2),3)
}

