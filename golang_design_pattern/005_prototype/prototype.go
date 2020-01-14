package prototype


type Cloneable interface{
	Clone() Cloneable
}

type Prototypemanager struct{
	prototypes map[string]Cloneable
}

func (p *Prototypemanager) Get(name string)Cloneable{
	return p.prototypes[name]
}

func (p *Prototypemanager) Set(name string,prototype Cloneable){
	p.prototypes[name] = prototype
}

func NewPrototypeManager() *Prototypemanager{
	return &Prototypemanager{prototypes:make(map[string]Cloneable)}
}
