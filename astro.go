package main

import (
    "fmt"
	"html/template"
    "net/http"
	"database/sql"
    _ "github.com/nakagami/firebirdsql"
	"github.com/elgs/gosqljson"
)

func home(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/home.html")
	var s = "hi"
	t.Execute(w, s)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	var table = r.URL.Path[len("/view/"):]
	//var n int
	conn, _ := sql.Open("firebirdsql", "firebird:firebird@localhost/var/lib/firebird/2.5/data/astro.fdb")
	defer conn.Close()
	var query = fmt.Sprintf("SELECT * FROM %s", table)
	a, _ := gosqljson.QueryDbToMapJson(conn, "lower", query, 0, 3)
	fmt.Fprintf(w, "Json: %s", a)
	//conn.QueryRow("SELECT Count(*) FROM STARS").Scan(&n)
    //fmt.Fprintf(w, "Found %v entries in %s", n, table)
}

func main() {
	http.HandleFunc("/", home)
    http.HandleFunc("/view/", viewHandler)
    http.ListenAndServe(":4444", nil)
}
