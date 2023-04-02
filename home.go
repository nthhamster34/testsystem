package main

import (
	"fmt"
	"sync"
)

var homeMutex sync.Mutex

func Home() {
	homeMutex.Lock()
	defer homeMutex.Unlock()
	fmt.Println("Welcome to my home!")
	fmt.Println("My name is John.")
	fmt.Println("I am a software developer.")
	fmt.Println("I love programming and learning new things.")
}
