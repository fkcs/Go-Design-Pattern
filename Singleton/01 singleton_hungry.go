package Singleton

//  方式一：饿汉方式
//  是线程安全，在程序加载阶段，即申请一段空间，但是当程序用不到该对象实例，则浪费一定空间
type SingletonA struct{}

var instanceA *SingletonA

func init() {
	instanceA = &SingletonA{}
}

func GetInstanceA() *SingletonA {
	return instanceA
}
