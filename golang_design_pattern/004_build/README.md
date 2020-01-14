创建者模式
===
子类的内容会有所不同，但是其组合在一起的算法或执行逻辑相对稳定时，使用

```
type Builder interface {
    XXX1()
    XXX2()
    XXX3()
}

type Diretory struct{
    builder Builder
}
func (d *Diretory) Construct(){
    d.duilder.XXX1()
    d.duilder.XXX2()
    d.duilder.XXX3()
    // todo 内部做着关于 Builder 的逻辑或算法
}

func NewDiretory (build Builder) *Diretory{}

```