package main

import (
  "strconv"
  //"strings"
  "fmt"
	"html/template"
    "net/http"
	"database/sql"
  _ "github.com/nakagami/firebirdsql"
	"github.com/elgs/gosqljson"
  "github.com/gorilla/mux"
)

var username = "sysdba"
var password = "masterkey"
var path     = "/Users/loran/firebird/testdb/test.fdb"
//var path     = "C:/Users/toram_000/firebird/testdb/test.fdb";
var url      = fmt.Sprintf("%s:%s@localhost%s", username, password, path)

func home(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/home.html")
	var s = "hi"
	t.Execute(w, s)
}

func querydb(w http.ResponseWriter, r *http.Request) {
  vars  := mux.Vars(r)
  table := vars["table"]
	conn, _ := sql.Open("firebirdsql", url)
	defer conn.Close()
  query := fmt.Sprintf("SELECT * FROM %s", table)
	a, _ := gosqljson.QueryDbToMapJson(conn, "lower", query, 0, 3)
	fmt.Fprintf(w, "%s", a)
	//conn.QueryRow("SELECT Count(*) FROM STARS").Scan(&n)
    //fmt.Fprintf(w, "Found %v entries in %s", n, table)
}

func queryind(w http.ResponseWriter, r *http.Request) {
  vars  := mux.Vars(r)
  table := vars["table"]
  idparm    := vars["id"]
  id, _ := strconv.Atoi(idparm)
  conn, _ := sql.Open("firebirdsql", url)
  defer conn.Close()
  var query string
  var child string
  var attr  string
  switch table {
    case "planets":
      child = "moons"
      attr  = "parent"
      break
    case "stars":
      child = "planets"
      attr  = "parent"
      break
    case "solar_systems":
      child = "stars"
      attr  = "parent"
      break
    case "galaxies":
      child = "solar_systems"
      attr  = "parent"
    default:
      child = table
      attr  = "id"
  }
  query = fmt.Sprintf("SELECT * FROM %s WHERE %s=%d", child, attr, id)
  a, _ := gosqljson.QueryDbToMapJson(conn, "lower", query, 0, 3)
  fmt.Fprintf(w, "%s", a)
  //conn.QueryRow("SELECT Count(*) FROM STARS").Scan(&n)
    //fmt.Fprintf(w, "Found %v entries in %s", n, table)
}

func main() {
  r := mux.NewRouter().StrictSlash(true)
  r.HandleFunc("/", home)
  r.HandleFunc("/{table}", querydb)
  r.HandleFunc("/{table}/{id}", queryind)
  http.Handle("/", r)
  fmt.Println("listing at 4444....")
  http.ListenAndServe(":4444", nil)
}
