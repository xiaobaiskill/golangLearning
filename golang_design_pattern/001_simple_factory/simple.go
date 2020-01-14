package simplefactory

import "fmt"

type API interface {
	Say(name string) string
}

func NewAPi(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}
	return nil
}

type hiAPI struct{}

func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("HI, %s", name)
}

type helloAPI struct{}

func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
