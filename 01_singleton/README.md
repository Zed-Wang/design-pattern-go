## Go语言设计模式01——单例模式

单例模式（singleton）是一种创建型模式，它提供了一种创建对象的最佳方式。这种模式涉及到一个单一的类，该类负责创建自己的对象，同时**确保只有单个对象被创建**。这个类提供了一种访问其唯一的对象的方式，可以直接访问，不需要实例化该类的对象

- 单例类只能有一个实例
- 单例类必须自己创建自己的唯一实例
- 单例类必须给所有其他对象提供这一实例

### 优点

1. 避免反复创建重复类，减少了内存的开销
2. 避免对资源的多重占用

### 缺点

1. 没有接口，不能继承，与单一职责原则冲突，一个类应该只关心内部逻辑，而不关心外面怎么样来实例化。

### 实现方式1——懒汉式

在第一次调用实例时创建对象

```go
package singleton

type Singleton struct{}

var lazySingleton *Singleton

func GetLazyInstance() *Singleton {
	if lazySingleton == nil {
		lazySingleton = &Singleton{}
	}
	return lazySingleton
}
```

存在线程安全问题

### 实现方式2——饿汉式

在包加载的时候创建单例对象

```go
package singleton

type Singleton struct{}

var singleton *Singleton

func init() {
	singleton = &Singleton{}
}

func GetInstance() *Singleton {
	return singleton
}
```

比懒汉式更加安全，但是如果该实例没有被用到，则浪费空间且减缓启动速度

### 实现方式3——双重检查机制

```go
package singleton

import "sync"

type Singleton struct{}

var (
   lazySingletonDCL *Singleton
   once             sync.Once
)

// Double Check Lock
func GetLazyDCLInstance() *Singleton {
   once.Do(func() {
      lazySingletonDCL = &Singleton{}
   })
   return lazySingletonDCL
}
```

`sync.Once.Do(f func())`确保 f 在整个程序中只被执行一次，不论 f 是否被更换

