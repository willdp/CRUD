package DBService

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"password"

	_ "github.com/go-sql-driver/mysql"
)

type Order struct {
	Id     int
	Name   string
	Food   string
	Adress string
	Price  float64
}

func DbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := password.DBpassword
	dbName := "orders"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err)
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := DbConn()
	selDB, err := db.Query("SELECT * FROM orders ORDER BY id DESC")

	if err != nil {
		panic(err)
	}

	ord := Order{}
	res := []Order{}

	for selDB.Next() {
		var id int
		var name, food, adress string
		var price float64
		err = selDB.Scan(&id, &name, &food, &adress, &price)
		if err != nil {
			panic(err)
		}
		ord.Id = id
		ord.Name = name
		ord.Food = food
		ord.Adress = adress
		ord.Price = price
		res = append(res, ord)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := DbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM orders WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	ord := Order{}
	for selDB.Next() {
		var id int
		var name, food, adress string
		var price float64
		err = selDB.Scan(&id, &name, &food, &adress, &price)
		if err != nil {
			panic(err.Error())
		}
		ord.Id = id
		ord.Name = name
		ord.Food = food
		ord.Adress = adress
		ord.Price = price
	}
	tmpl.ExecuteTemplate(w, "Show", ord)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := DbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM orders WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	ord := Order{}
	for selDB.Next() {
		var id int
		var name, food, adress string
		var price float64
		err = selDB.Scan(&id, &name, &food, &adress, &price)
		if err != nil {
			panic(err.Error())
		}
		ord.Id = id
		ord.Name = name
		ord.Food = food
		ord.Adress = adress
		ord.Price = price
	}
	tmpl.ExecuteTemplate(w, "Edit", ord)
	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := DbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		food := r.FormValue("food")
		adress := r.FormValue("adress")
		price := r.FormValue("price")
		insForm, err := db.Prepare("INSERT INTO orders(name, food, adress, price) VALUES(?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, food, adress, price)
		log.Println("Update: Name " + name + " | Food " + food + " | Adress " + adress + " | price " + price)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := DbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		food := r.FormValue("food")
		adress := r.FormValue("adress")
		price := r.FormValue("price")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE orders SET name=?, food=?, adress=?, price=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, food, adress, price, id)
		log.Println("UPDATE: Name: " + name + " | Food: " + food + " | Adress" + adress + " | price" + price)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := DbConn()
	ord := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM orders WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(ord)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}
