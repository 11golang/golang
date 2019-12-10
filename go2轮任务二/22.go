package main

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	_ "log"
	"net/http"
	_ "text/template"
)
type re struct{
	name string
	text string
	num  int
}

func check(err error){
if err !=nil{
	panic(err)
}
}

func flu(w http.ResponseWriter,r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "fluid.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			panic(err)
			return
		}
		fmt.Fprintf(w, "submit success")
		Name := r.FormValue("name")
		Text := r.FormValue("text")
		db, err := sql.Open("mysql", "root:root@/test")
		defer db.Close()

		if err != nil {
			fmt.Println(err)
			return
		}else{
		fmt.Println("success")
	}
		result, err := db.Exec("INSERT  INTO re(name,text) VALUES(?,?)",Name, Text)
		check(err)
		numes, err := result.LastInsertId()
		fmt.Println(numes)
	default:
		fmt.Fprintf(w, "err")
	}
}
func Query(w http.ResponseWriter,r *http.Request) {
	db,err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	defer db.Close()
	check(err)

	rows, err := db.Query("SELECT * FROM re  WHERE num >= (SELECT FLOOR( MAX(num) * RAND()) FROM re ) ORDER BY num LIMIT 1")
	check(err)

	for rows.Next(){
		var fetch re
		err = rows.Scan(&fetch.name,&fetch.text,&fetch.num)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		fmt.Fprintf(w,"name=%s\n text=%s",fetch.name,fetch.text)
		fmt.Println(fetch.name,fetch.text)
	}

}
func Delete(w http.ResponseWriter,r *http.Request){
	db,err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	defer db.Close()
	check(err)
	Num := r.FormValue("Num")
	res, err := db.Exec("DELETE FROM re WHERE num=?", Num)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	fmt.Fprintf(w,"delete successful!")
}
func square(w http.ResponseWriter,r *http.Request){
	switch r.Method{
	case "GET":
		http.ServeFile(w,r,"square.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
		fmt.Println(err)
		return
	}
		fmt.Fprintf(w,"scu")
	}
}
func main(){

		http.HandleFunc("/", flu)
		http.HandleFunc("/get",Query)
		http.HandleFunc("/delete",Delete)
		http.HandleFunc("/square.html",square)
		fmt.Println("working")
		http.ListenAndServe("localhost:8080", nil)


		}
