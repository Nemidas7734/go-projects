package main

import (
    "encoding/json"
    "fmt"
    "html/template"
    "net/http"
    "os"
    "time"
    "strings"
)

type Article struct {
    ID      string    `json:"id"`
    Title   string    `json:"title"`
    Content string    `json:"content"`
    Date    time.Time `json:"date"`
}

var articlesDir = "articles"

var templateFuncs = template.FuncMap{
    "formatDate": func(t time.Time) string {
        return t.Format("January 2, 2006")
    },
}

// Update getArticle to work with article number
func getArticle(id string) (*Article, error) {
    filename := fmt.Sprintf("article%s.json", id)
    filePath := fmt.Sprintf("%s/%s", articlesDir, filename)
    
    data, err := os.ReadFile(filePath)
    if err != nil {
        return nil, err
    }

    var article Article
    err = json.Unmarshal(data, &article)
    if err != nil {
        return nil, err
    }
    
    article.ID = id // Set the ID from filename
    return &article, nil
}

func articleHandler(w http.ResponseWriter, r *http.Request) {
    // Extract ID from URL
    id := strings.TrimPrefix(r.URL.Path, "/article/")
    
    article, err := getArticle(id)
    if err != nil {
        http.Error(w, "Article not found", http.StatusNotFound)
        return
    }

    tmpl, err := template.New("article.html").
        Funcs(templateFuncs).
        ParseFiles("templates/article.html")
    
    if err != nil {
        http.Error(w, "Error loading template", http.StatusInternalServerError)
        return
    }

    err = tmpl.Execute(w, article)
    if err != nil {
        http.Error(w, "Error executing template", http.StatusInternalServerError)
        return
    }
}

func loadArticles() ([]Article, error) {
    var articles []Article
    
    if err := os.MkdirAll(articlesDir, 0755); err != nil {
        return nil, err
    }
    
    files, err := os.ReadDir(articlesDir)
    if err != nil {
        return nil, err
    }

    for _, file := range files {
        if file.IsDir() {
            continue
        }
        
        // Extract article number from filename
        id := strings.TrimPrefix(strings.TrimSuffix(file.Name(), ".json"), "article")
        
        filePath := fmt.Sprintf("%s/%s", articlesDir, file.Name())
        data, err := os.ReadFile(filePath)
        if err != nil {
            continue
        }

        var article Article
        err = json.Unmarshal(data, &article)
        if err != nil {
            continue
        }
        
        article.ID = id // Set the ID from filename
        articles = append(articles, article)
    }
    return articles, nil
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    articles, err := loadArticles()
    if err != nil {
        http.Error(w, "Error loading articles", http.StatusInternalServerError)
        return
    }

    tmpl, err := template.New("home.html").
        Funcs(templateFuncs).
        ParseFiles("templates/home.html")
    
    if err != nil {
        http.Error(w, "Error loading template", http.StatusInternalServerError)
        return
    }

    err = tmpl.Execute(w, articles)
    if err != nil {
        http.Error(w, "Error executing template", http.StatusInternalServerError)
        return
    }
}

func main() {
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/article/", articleHandler)
    
    fmt.Println("Server starting on http://localhost:8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Error Loading Server:", err)
    }
}