package singleton

import (
	"fmt"
	"sync"
)

type singleton struct {
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		fmt.Println("New instance")
		instance = &singleton{}
	})
	fmt.Println("Call instance")
	return instance
}
