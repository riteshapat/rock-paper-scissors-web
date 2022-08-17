package main

import (
	"encoding/json"
	"html/template"
	"log"
	"myapp/rps"
	"net/http"
	"strconv"
)

func myWeb(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "index.html")

	//first ways
	// html := `<strong>Hello World!</strong>`
	// w.Header().Set("Content-type", "text/html")
	// fmt.Fprint(w, html)
}

func playRound(w http.ResponseWriter, r *http.Request) {

	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.PlayRound(playerChoice)

	//winner, computervalues, result := rps.PlayRound(1)
	//result := rps.PlayRound(1)

	out, err := json.MarshalIndent(result, "", "   ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(out)

	//log.Println(winner, computervalues, result)

}

func main() {

	http.HandleFunc("/", myWeb)
	http.HandleFunc("/play", playRound)

	log.Println("starting we browser on port 8080")
	http.ListenAndServe(":8080", nil)
}

func renderTemplate(w http.ResponseWriter, page string) {

	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}

}
