package main
import (
    "net/http"
    "html/template"
)


var templates = template.Must(template.ParseFiles("chart.html"))

func charts(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "chart", nil)
}

func renderTemplate(w http.ResponseWriter, templateName string, data interface{}) {
    //handle js files
    http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
    http.Handle("/data/", http.StripPrefix("/data/", http.FileServer(http.Dir("data"))))

    err := templates.ExecuteTemplate(w, templateName + ".html", data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}


