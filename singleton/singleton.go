package singleton

import (
	"fmt"
	"sync"
)

type Singleton struct {
	Num int
}

var instance *Singleton
var once = &sync.Once{}

func GetInstance() *Singleton {
	if instance == nil {
		instance = &Singleton{}
		// do sth
		instance.Num = 1
		fmt.Print("do sth when init Singleton\n")
	}
	return instance
}

func GetInstanceWithOnce() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
		// do sth
		instance.Num = 1
		fmt.Print("do sth when init Singleton\n")
	})
	return instance
}

func (s *Singleton) GetNum() int {
	fmt.Print("Singleton Operation\n")
	return s.Num
}