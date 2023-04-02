package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var randomMutex sync.Mutex

func RandomQuote() {
	randomMutex.Lock()
	defer randomMutex.Unlock()

	rand.Seed(time.Now().Unix())
	quotes := []string{
		"Be the change you wish to see in the world.",
		"Success is not final, failure is not fatal: It is the courage to continue that counts.",
		"Believe you can and you're halfway there.",
		"Try to be a rainbow in someone's cloud.",
		"In the end, it's not the years in your life that count. It's the life in your years.",
	}
	fmt.Println("Random Quote:", quotes[rand.Intn(len(quotes))])
}
