package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Home struct {
	Name        string
	Age         int
	Occupation  string
	Description string
}

type Random struct {
	Quote string
}

type Greeting struct {
	Time      string
	DayOfWeek string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		bio := Home{
			Name:        "Donovan Yam",
			Age:         18,
			Occupation:  "IT Student",
			Description: "I am a IT Student able to do basic web development.",
		}

		template, err := template.ParseFiles("Home.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = template.Execute(w, bio)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	quotes := []string{
		"Be the change you wish to see in the world.",
		"Success is not final, failure is not fatal: It is the courage to continue that counts.",
		"Believe you can and you're halfway there.",
		"Try to be a rainbow in someone's cloud.",
		"In the end, it's not the years in your life that count. It's the life in your years.",
	}

	http.HandleFunc("/quote", func(w http.ResponseWriter, r *http.Request) {
		rand.Seed(time.Now().Unix())
		index := rand.Intn(len(quotes))

		quote := quotes[index]
		qotd := Random{
			Quote: quote,
		}
		template, err := template.ParseFiles("Random.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = template.Execute(w, qotd)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/greeting", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		dayOfWeek := now.Weekday().String()

		greeting := Greeting{
			Time:      now.Format(time.Kitchen),
			DayOfWeek: dayOfWeek,
		}

		template, err := template.ParseFiles("greeting.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = template.Execute(w, greeting)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
