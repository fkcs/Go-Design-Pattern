package Decorator

import "testing"

func TestDecorator(t *testing.T) {
	var component Component
	tmp := new(ConcreteComponent)
	component = NewConCreteDecoratorA(tmp)
	component.operation()
}
