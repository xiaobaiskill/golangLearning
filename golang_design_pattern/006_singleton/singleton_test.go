package singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetInstance(t *testing.T) {
	s := GetInstance()
	s.SetName("jmz")

	s1 := GetInstance()
	assert.Equal(t,s1.GetName(),"jmz")
}
