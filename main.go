package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"rockraperscissors/RockPaperScissors"
)

// homePage & playRound r handlers, evry web handler always 1)takes an http.ResponseWriter & 2)it takes a pointer to an http.Request
func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

func playRound(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.PlayRound(playerChoice) //calling the func PlayRound from rps.go

	out, err := json.MarshalIndent(result, "", "     ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func main() {
	//roots
	http.HandleFunc("/play", playRound)
	http.HandleFunc("/", homePage)

	log.Println("Starting web server on port 8080   -->   http://127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}

func renderTemplate(w http.ResponseWriter, page string) {
	//parse our html file n2 da necessary format & send dat 2da browser
	//t: for template, err: checks for an error
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}

	//this func executes our template
	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}