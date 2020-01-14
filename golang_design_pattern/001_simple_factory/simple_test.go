package simplefactory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAPI1(t *testing.T) {
	api := NewAPi(1)
	assert.Equal(t,api.Say("jmz"),"HI, jmz")
}

func TestAPI2(t *testing.T){
	api := NewAPi(2)
	assert.Equal(t, api.Say("jmz"),"Hello, jmz")
}