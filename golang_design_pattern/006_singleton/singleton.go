package singleton

import "sync"

type Class struct {
	name string
}

func (s *Class) SetName(name string){
	s.name = name
}
func (s *Class) GetName()string{
	return s.name
}


var singleton *Class

var once sync.Once

// getInstance 用于获取单列模型对象
func GetInstance() *Class{
	once.Do(func() {
		singleton = &Class{}
	})
	return singleton
}

