package build

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDirector(t *testing.T) {
	var director *Director

	builder1 := &Builder1{}
	director = NewDirector(builder1)
	director.Construct()
	assert.Equal(t,builder1.GetResult(),"123")

	builder2 := &Builder2{}
	director = NewDirector(builder2)
	director.Construct()
	assert.Equal(t,builder2.GetResult(),6)
}
