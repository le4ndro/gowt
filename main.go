package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Tool struct
type Tool struct {
	Id       int
	Name     string
	Category string
	URL      string
	Rating   int
	Notes    string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := os.Getenv("DATABASE_USERNAME")
	dbPass := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")
	dbServer := os.Getenv("DATABASE_SERVER")
	dbPort := os.Getenv("DATABASE_PORT")
	log.Println("Database host: " + dbServer)
	log.Println("Database port: " + dbPort)
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbServer+":"+dbPort+")/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("templates/*"))

//Index handler
func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM tools ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	tool := Tool{}
	res := []Tool{}

	for selDB.Next() {
		var id, rating int
		var name, category, url, notes string
		err := selDB.Scan(&id, &name, &category, &url, &rating, &notes)
		if err != nil {
			panic(err.Error())
		}
		log.Println("Listing Row: Id " + string(id) + " | name " + name + " | category " + category + " | url " + url + " | rating " + string(rating) + " | notes " + notes)

		tool.Id = id
		tool.Name = name
		tool.Category = category
		tool.URL = url
		tool.Rating = rating
		tool.Notes = notes
		res = append(res, tool)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

//Show handler
func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM tools WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}

	tool := Tool{}

	for selDB.Next() {
		var id, rating int
		var name, category, url, notes string
		err := selDB.Scan(&id, &name, &category, &url, &rating, &notes)
		if err != nil {
			panic(err.Error())
		}

		log.Println("Showing Row: Id " + string(id) + " | name " + name + " | category " + category + " | url " + url + " | rating " + string(rating) + " | notes " + notes)

		tool.Id = id
		tool.Name = name
		tool.Category = category
		tool.URL = url
		tool.Rating = rating
		tool.Notes = notes
	}
	tmpl.ExecuteTemplate(w, "Show", tool)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM tools WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}

	tool := Tool{}

	for selDB.Next() {
		var id, rating int
		var name, category, url, notes string
		err := selDB.Scan(&id, &name, &category, &url, &rating, &notes)
		if err != nil {
			panic(err.Error())
		}

		tool.Id = id
		tool.Name = name
		tool.Category = category
		tool.URL = url
		tool.Rating = rating
		tool.Notes = notes
	}

	tmpl.ExecuteTemplate(w, "Edit", tool)
	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		category := r.FormValue("category")
		url := r.FormValue("url")
		rating := r.FormValue("rating")
		notes := r.FormValue("notes")
		insForm, err := db.Prepare("INSERT INTO tools (name, category, url, rating, notes) VALUES (?, ?, ?, ?, ?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, category, url, rating, notes)
		log.Println("Insert Data: name " + name + " | category " + category + " | url " + url + " | rating " + rating + " | notes " + notes)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		category := r.FormValue("category")
		url := r.FormValue("url")
		rating := r.FormValue("rating")
		notes := r.FormValue("notes")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE tools SET name=?, category=?, url=?, rating=?, notes=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, category, url, rating, notes, id)
		log.Println("UPDATE Data: name " + name + " | category " + category + " | url " + url + " | rating " + rating + " | notes " + notes)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	tool := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM tools WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(tool)
	log.Println("DELETE " + tool)
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8080", nil)
}
