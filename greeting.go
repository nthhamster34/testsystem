package main

import (
	"fmt"
	"sync"
	"time"
)

var greetingMutex sync.Mutex

func Greeting() {
	greetingMutex.Lock()
	defer greetingMutex.Unlock()

	fmt.Println("Hello, it's Friday!")
	fmt.Println("The current time is:", time.Now().Format("15:04:05"))
}
