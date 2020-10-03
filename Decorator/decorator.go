package Decorator

import "fmt"

// 定义对象接口，可以给这些对象东条添加职责
type Component interface {
	operation()
}

// 定义一个对象，可以给这个对象添加一些职责
type ConcreteComponent struct{}

func (c *ConcreteComponent) operation() {
	fmt.Println("我是被装饰的对象")
}

// 拥有一个指向Component对象的引用，并定义一个与Component接口一致的接口
// 将请求转发给它的Component对象，并有可能在转发请求前后执行一些附加的动作
type Decorator struct {
	component Component
}

func NewDecorator(component Component) *Decorator {
	return &Decorator{
		component: component,
	}
}
func (d *Decorator) operation() {
	if d.component != nil {
		d.component.operation()
	}
}

// 被装饰者,向组件添加职责
type ConCreteDecoratorA struct {
	decorator *Decorator
}

func NewConCreteDecoratorA(component Component) *ConCreteDecoratorA {
	return &ConCreteDecoratorA{
		decorator: &Decorator{
			component: component,
		},
	}
}
func (c *ConCreteDecoratorA) operation() {
	c.decorator.operation()
	c.decorate()
}
func (c *ConCreteDecoratorA) decorate() {
	fmt.Println("装饰操作A")
}
