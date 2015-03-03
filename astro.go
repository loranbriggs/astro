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
  "github.com/PuerkitoBio/goquery"
  "net/url"
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
  idparm    := vars["id"]
  id, notInt := strconv.Atoi(idparm)
	conn, _ := sql.Open("firebirdsql", url)
	defer conn.Close()
  var query string
  if notInt == nil {
    // we have an id
    query = fmt.Sprintf("SELECT * FROM %s WHERE id=%d", table, id)
  } else {
    // else list them all
    query = fmt.Sprintf("SELECT * FROM %s", table)
  }
	a, _ := gosqljson.QueryDbToMapJson(conn, "lower", query, 0, 3)
	fmt.Fprintf(w, "%s", a)
	//conn.QueryRow("SELECT Count(*) FROM STARS").Scan(&n)
    //fmt.Fprintf(w, "Found %v entries in %s", n, table)
}

func getImages(w http.ResponseWriter, r *http.Request) {
  vars  := mux.Vars(r)
  url := vars["url"]
  fmt.Println(url)

  doc, err := goquery.NewDocument(url) 
  if err != nil {
    log.Fatal(err)
  }

  var src1 string
  //var src2 string

  doc.Find("img").Each(func(i int, s *goquery.Selection) {
    if i == 1 {
      var srca, _ = s.Attr("src")
      src1 = srca
    }
    /*
    if i == 2 {
      var srcb, _ = s.Attr("src")
      src2 = srcb
    }
    */
  })
  fmt.Fprintf(w, "%s", url.QueryEscape(src1))
}

func main() {
  r := mux.NewRouter().StrictSlash(true)
  r.HandleFunc("/", home)
  r.HandleFunc("/{table}", querydb)
  r.HandleFunc("/{table}/{id}", querydb)
  r.HandleFunc("/getimages/{url}" getImages)
  http.Handle("/", r)
  fmt.Println("listing at 4444....")
  http.ListenAndServe(":4444", nil)
}
