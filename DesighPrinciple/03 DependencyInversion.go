package DesighPrinciple

import "fmt"

// 依赖倒置原则

// 交通工具接口
type Drive interface {
	drive()
}

type Bike struct{}

func (b *Bike) drive() {
	fmt.Println("drive by Bike!")
}

type Car struct{}

func (c *Car) drive() {
	fmt.Println("drive by Car!")
}

// Person类
type Person struct {
	drive Drive
}

func NewPerson() *Person {
	return &Person{
		drive: new(Bike),
	}
}
func (p *Person) DriveTool() {
	p.drive.drive()
}

/*
Person中使用接口Drive来定义交通工具，出行依赖的是交通工具接口，而不是具体的交通工具。
这样，可以自由选择交通工具，只要交通工具实现该接口，并将其传递给Person类中的Drive接口。
好处：可扩展性好，以后添加其他交通工具不影响代码实现，即抽象不依赖于细节。
*/

// 控制反转
// 反转：在没有使用框架之前，程序员自己控制整个程序的执行，当使用框架之后，整个程序的执行可以通过框架来控制；
// 框架提供了可扩展的代码骨架，用来组装对象/管理整个执行流程，程序只需关注扩展点，就可以利用框架来驱动整个程序流程的执行。

func NewPersonx(drive Drive) *Person {
	return &Person{
		drive: drive,
	}
}
