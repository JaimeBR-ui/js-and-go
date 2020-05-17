// from https://tutorialedge.net/golang/creating-restful-api-with-golang/
package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
)

type Article struct {
    Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

type Articles [] Article

func allArticles(w http.ResponseWriter, r *http.Request)  {
    articles := Articles{
        Article{Title:"Test Title", Desc: "Test Description",
        Content: "Hello world."},
    }
    fmt.Println("Endpoint Hit: All Articles Endpoint")
    json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Homepage Endpoint Hit")
    http.Redirect(w, r, "/index.html", http.StatusSeeOther)
}

func handleRequests() {
    http.Handle("/", http.FileServer(http.Dir("./static")))
    http.HandleFunc("/articles", allArticles)
    log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
    handleRequests();
}
