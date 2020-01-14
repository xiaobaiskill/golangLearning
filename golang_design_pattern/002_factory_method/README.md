工厂方法模式
===
通过方法产生实例
产生实例的方法就是工厂方法

```
type AAA struct{}

type factory struct{}
func (factory) Create()AAA{}
```

