package main

import (
    "net/http"
    "log"
    //"os"
    //"strings"
    //"io/ioutil"
    "fmt"
    //"html/template"
    //"runtime"
    "github.com/sirMackk/templates_ago"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

type Record struct {
   Title, Body, Fs_path string 
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
    //rec := make(map[string]interface{})
    //rec["Name"] = "Main Index"
    //rec["Records"] = []string{"this", "is", "sparta"}
    //recs := make([]map[string]string)
    recs := make([]Record, 0)
    rows, err := db.Query("select title, body, fs_path from Records limit 10")
    //cols, err := rows.Columns()
    if err != nil {
        fmt.Println(52)
        fmt.Println(err)
        panic(err)
    }
    defer rows.Close()
    for rows.Next() {
        //vals := make([]interface{}, len(cols))
        //rows.Scan(vals...)
        var title, body, fs_path string
        rows.Scan(&title, &body, &fs_path)
        //recs = append(&Record{Title: string(vals[0]), Body: vals[1], Fs_path: vals[2]})
        recs = append(recs, Record{Title: title, Body: body, Fs_path: fs_path})
    }


    err = templates.Execute(w, recs)
    if err != nil { panic(err) }
}

func SimplePage(w http.ResponseWriter, req *http.Request) {
    fmt.Fprint(w, "Simple routing test")
}

var db *sql.DB

func setupDatabase() {
    var err error
    db, err = sql.Open("sqlite3", "./test_db.db")
    if err != nil { panic(err) }

    creation := `
    create table if not exists Records(id integer not null primary key, title text, body text, fs_path text);
    create table if not exists Users(id integer not null primary key, username text, password text, role text);
    `

    _, err = db.Exec(creation)
    if err != nil {
        fmt.Println("85")
        panic(err)
    }

    tx, err := db.Begin()
    if err != nil { 
        fmt.Println("91")
        panic(err)
    }

    stmt, err := tx.Prepare("insert into Records(title, body, fs_path) values(?, ?, ?)")
    if err != nil { 
      fmt.Println(97)
      fmt.Println(err)
    }
    defer stmt.Close()

    for i := 0; i < 10; i++ {
        _, err = stmt.Exec(fmt.Sprintf("Title %d", i), "Body", "./")
        if err != nil {
            fmt.Println("stmt error")
        }
    }

    tx.Commit()
}


var templates = templates_ago.NewTemplates()

func init() {
    setupDatabase()
    templates_ago.LoadTemplates("templates/", templates)
}

func main() {
    defer db.Close()
    http.HandleFunc("/", HandleWrapper(Index))
    http.HandleFunc("/other/", HandleWrapper(SimplePage))

    err := http.ListenAndServe("localhost:3000", nil)
    if err != nil { panic(err) }
}




