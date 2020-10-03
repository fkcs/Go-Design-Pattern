package Singleton

import (
	"sync"
	"sync/atomic"
)

// 方式二：懒汉方式
// 非线程安全，在获取对象实例时，如果实例为空则创建，但是如果是多线程调用，则会创建多个实例
type SingletonLazy struct{}

var instanceB *SingletonLazy

func GetInstanceLaze() *SingletonLazy {
	if instanceB == nil {
		instanceB = &SingletonLazy{}
	}
	return instanceB
}

// 方式三：双重检查机制，sync.Once()
// 3.1  加锁避免资源竞争导致数据不一致问题，但是每次请求创建实例时都会加锁，导致性能降低
type SingletonC struct{}

var instanceC *SingletonC
var mx sync.Mutex

func GetInstanceC() *SingletonC {
	mx.Lock()
	defer mx.Unlock()
	if instanceC == nil {
		instanceC = &SingletonC{}
	}
	return instanceC
}

// 3.2  sync.Once，底层实现是通过原子操作+锁
type SingletonD struct{}

var instanceD *SingletonD
var once sync.Once

func GetInstanceD() *SingletonD {
	once.Do(func() {
		instanceD = &SingletonD{}
	})
	return instanceD
}

type SingletonF struct{}

var instanceF *SingletonF
var tmp uint32

func GetInstanceF() *SingletonF {
	if atomic.LoadUint32(&tmp) == 1 {
		return instanceF
	}
	mx.Lock()
	defer mx.Unlock()
	if tmp == 0 {
		instanceF = &SingletonF{}
		atomic.StoreUint32(&tmp, 1)
	}
	return instanceF
}
