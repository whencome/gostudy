package main

import (
	"fmt"
	"sync"
	"time"
)

// 方式一：程序启动时创建实例
// 普通的单例模式，对象在程序启动时初始化，后续只使用此对象
var fooInst *Foo = newFoo()

// 方式二：延迟加载，调用时初始化
// 升级版，在调用时才进行初始化，实现延迟加载
var once sync.Once
var barInst *Bar

type Foo struct {
	timeStamp int64
}

// 创建实例，此方法每次都会实例化对象
// 只是通过在应用运行时初始化的全局变量实现单例模式
func newFoo() *Foo {
	fmt.Printf("%s ------- initialize Foo instance --------", time.Now().Format("2006-01-02 15:04:05"))
	return &Foo{
		timeStamp: time.Now().Unix(),
	}
}

func (f *Foo) Print() {
	fmt.Printf("[Foo] Time Stamp => %d\n", f.timeStamp)
}

type Bar struct {
	timeStamp int64
}

// 单例模式：调用时初始化
// 使用 sync.Once 保证实例只被初始化一次
func NewBar() *Bar {
	if barInst != nil {
		return barInst
	}
	once.Do(func() {
		fmt.Printf("%s ------- initialize Bar instance --------\n", time.Now().Format("2006-01-02 15:04:05"))
		barInst = &Bar{
			timeStamp: time.Now().Unix(),
		}
	})
	return barInst
}

func (b *Bar) Print() {
	fmt.Printf("[Bar] Time Stamp => %d\n", b.timeStamp)
}

func main() {
	fmt.Println("start app...")
	time.Sleep(time.Second * 3)
	for i := 0; i < 10; i++ {
		go func() {
			fooInst.Print()
		}()
		go func() {
			bi := NewBar()
			bi.Print()
		}()
	}
	time.Sleep(time.Second * 3)
	fmt.Println("exit app...")
}
