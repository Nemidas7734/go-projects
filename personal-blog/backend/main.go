package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"
)

type Article struct {

    Title string `json:"title"`
    Content string `json:"content"`
    Date time.Time `json:"date"`

}


var articlesDir = "articles"

func loadArticles() ([]Article, error) {

    var articles [] Article
    files, err := os.ReadDir(articlesDir)

    if err != nil {
        return nil, err
    }

    for _, file := range files {
        if file.IsDir() {
            continue
        }
        
        filePath := fmt.Sprintf("%s/%s", articlesDir, file.Name())
        data, err := os.ReadFile(filePath)
        
        if err != nil {
            return nil, err
        }

        var article Article
        err = json.Unmarshal(data, &article)
        if err != nil {
            return nil, err
        }
        articles = append(articles, article)
    }
    return articles, nil
}


func homeHandler(w http.ResponseWriter, r *http.Request) {

    articles, err := loadArticles()
    if err != nil {
        http.Error(w, "Error loading articles", http.StatusInternalServerError)
        return
    }

    tmpl, err := template.New("home").ParseFiles("frontend/templates/index.html")
    if err != nil {
        http.Error(w, "Error loading template", http.StatusInternalServerError)
        return
    }

    tmpl.Execute(w, articles)
}



func main() { 

    http.HandleFunc("/", homeHandler)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Error Loading Server", err)
    }
}