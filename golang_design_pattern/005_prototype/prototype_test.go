package prototype

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Type1 struct{
	name string
}

func (t *Type1) Clone() Cloneable{
	tc := *t
	return &tc
}


var manage *Prototypemanager

func init(){
	manage = NewPrototypeManager()
	manage.Set("t1",&Type1{"jmz"})
}


func TestClone(t *testing.T){
	t1 := manage.Get("t1")
	t2 := t1.Clone()
	assert.Equal(t,t1,t2)
}

func TestCloneFormmanager(t *testing.T){
	t1 := manage.Get("t1")

	tc1 := t1.Clone()
	tc2 := t1.Clone()

	t2 := tc1.(*Type1)
	assert.Equal(t,t2.name,"jmz")
	t3 := tc2.(*Type1)
	t3.name = "jjjj"

	tc3 := t1.Clone()
	t4 := tc3.(*Type1)
	assert.NotEqual(t,t4.name,"jjj")
}