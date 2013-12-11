package main

import (
    //"log"
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "os"
    "encoding/json"
)

var STATIC_PATH = "static/"

func ServeStatic(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    asset_path := fmt.Sprintf("%s%s", STATIC_PATH, vars["asset"])
    if _, err := os.Stat(asset_path); os.IsNotExist(err) {
        http.Error(w, "Not Found", 404)
    } else {
        http.ServeFile(w, r, asset_path)
    }
}

func homePage(w http.ResponseWriter, r *http.Request) {
    page := `
    <html data-ng-app="app">
    <head>
    <script src="static/angular.min.js"></script>
    <script src="static/angular-route.min.js"></script>
    <script src="static/app.js"></script>
    <script src="static/router.js"></script>
    <script src="static/factories.js"></script>
    <script src="static/controllers.js"></script>
    <title>Testing</title>
    </head>
    <body data-ng-controller="exampleCtrl">
    <h1>Assets and json testing</h1>
    <div data-ng-view></div>
    <ul>
      <li data-ng-repeat="item in Items">
        {{ item }}
      </li>
    </ul>
    <form id="daform" data-ng-submit="sendForm()">
      <input type="text" data-ng-model="Name">
      <input type="text" data-ng-model="Message">
      <input type="submit" value="CLICK">
    </form>
    <p>{{ Name | uppercase }}</p>
    <p>{{ Message }}</p>
      
    </body>
    </html>
    `
    fmt.Fprint(w, page)
}

type Message struct {
    Name, Message string
}

func ServeData(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        msg := Message{Name: "BOOMBLE", Message: "Tajny"}
        b, _ := json.Marshal(msg)

        fmt.Println(string(b))
        fmt.Fprint(w, string(b))
    case "POST":
        defer r.Body.Close()
        bajts := make([]byte, 512)
        var msg Message
        n, _ := r.Body.Read(bajts)
        fmt.Println(string(bajts))
        err := json.Unmarshal(bajts[:n], &msg)
        if err != nil {
          fmt.Println(err)
        }
        fmt.Println(msg)
        http.Redirect(w, r, "/", 301)
    }
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/", homePage)
    router.HandleFunc("/static/{asset}", ServeStatic)
    router.HandleFunc("/data", ServeData)

    http.Handle("/", router)
    err := http.ListenAndServe(":3000", nil)
    if err != nil { panic(err) }
}


