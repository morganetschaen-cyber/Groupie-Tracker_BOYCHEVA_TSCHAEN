package main

import(
	"net/http"
	"log"
)

type Artist struct {
	Id            int    `json:"id"`
    Image         string `json:"image"`
	Name          string `json:"name"`
    CreationDate  int    `json:"creationDate"`
	FirstAlbum    string `json:"firstAlbum"`
	Locations     []string `json:"locations"`
	ConcertDates  []string `json:"concertDates"`
}

func main() {
    http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static")),
		),
	)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/index.html")
	})

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}