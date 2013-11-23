package main

import (
    "net/http"
    "log"
    "os"
    "strings"
    "io/ioutil"
    "fmt"
    "html/template"
    "runtime"
)

type Record struct {
  Name string
  Id int
}

type HandleFunc func(w http.ResponseWriter, req *http.Request)

func logPanic(function HandleFunc) HandleFunc {
  return func(w http.ResponseWriter, req *http.Request) {
    defer func() {
      if r := recover(); r != nil {
        log.Println(r)
      }
    }()
    function(w, req)
  }
}

func addHeader(function HandleFunc) HandleFunc {
  return func(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    function(w, req)
  }
}

func HandleWrapper(function HandleFunc) HandleFunc {
  return logPanic(addHeader(function))
}

func Index(w http.ResponseWriter, req *http.Request) {
    //rec := &Record{Id: 1, Name: "Bob"}
    rec := make(map[string]interface{})
    rec["Name"] = "Main Index"
    rec["Records"] = []string{"this", "is", "sparta"}
    fmt.Println(getFuncName())
    //err := templates.ExecuteTemplate(w, "index.html", rec)
    //err := templates["index.html"].Execute(w, rec)
    
    err := templates[getFuncName() + ".html"].ExecuteTemplate(w, "base", rec)
    if err != nil { panic(err) }
    //t, _ := template.ParseFiles("index.html")
    //t.Execute(w, rec)
}

func SimplePage(w http.ResponseWriter, req *http.Request) {
    fmt.Fprint(w, "Simple routing test")
}

func getFuncName() string {
    pc, _, _, ok := runtime.Caller(1)
    if !ok { return "Unknown" }
    fnName := runtime.FuncForPC(pc)
    if fnName == nil { return "Anonymous" }
    fnNameParts := strings.Split(fnName.Name(), ".")
    return strings.ToLower(fnNameParts[len(fnNameParts)-1])
}



func setupTemplates() {
    tplPath := "templates/"
    if _, err := os.Stat("templates/base.html"); os.IsNotExist(err) {
        panic("Missing templates/base.html file")
    }
    files, err := ioutil.ReadDir(tplPath)
    if err != nil { panic(err) }
    for _, file := range files {
        if file.Name() == "base.html" { continue }
        fmt.Println(file.Name())
        templates[file.Name()] = template.Must(template.ParseFiles(tplPath + file.Name(), "base.html"))
    }
}

//var templates map[string]*template.Template
var templates = make(map[string]*template.Template)

func init() {
    setupTemplates()
}

func main() {
    http.HandleFunc("/", HandleWrapper(Index))
    http.HandleFunc("/other/", HandleWrapper(SimplePage))

    err := http.ListenAndServe("localhost:3000", nil)
    if err != nil { panic(err) }
}




