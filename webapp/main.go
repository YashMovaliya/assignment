package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	Id         int    `json:id`
	Name       string `json:name`
	City       string `json:city`
	Department string `json:department`
	Salary     int    `json:salary`
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "golang"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	//new template engine

	router.GET("/", func(ctx *gin.Context) {
		//render only file, must full name with extension
		db := dbConn()
		selDB, err := db.Query("SELECT * FROM Employee ORDER BY id DESC")
		if err != nil {
			panic(err.Error())
		}
		emp := Employee{}
		res := []Employee{}
		for selDB.Next() {
			var id, salary int
			var name, city, department string
			err = selDB.Scan(&id, &name, &city, &department, &salary)
			if err != nil {
				panic(err.Error())
			}
			emp.Id = id
			emp.Name = name
			emp.City = city
			emp.Department = department
			emp.Salary = salary
			res = append(res, emp)
		}
		ctx.HTML(http.StatusOK, "page.html", gin.H{"title": "Page file title!!", "a": res})
	})

	router.GET("/add", func(ctx *gin.Context) {
		//render only file, must full name with extension
		ctx.HTML(http.StatusOK, "add.html", gin.H{"title": "Page file title!!"})
	})

	router.GET("/submit", func(ctx *gin.Context) {
		//render only file, must full name with extension
		var name, city, department string
		var salary string
		name = ctx.Query("name")
		city = ctx.Query("price")
		department = ctx.Query("quality")
		salary = ctx.Query("salary")

		db := dbConn()
		insForm, err := db.Prepare("INSERT INTO Employee(name, city, department, salary) VALUES(?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, city, department, salary)
		ctx.HTML(http.StatusOK, "page.html", gin.H{"title": "Page file title!!"})
	})

	router.Run(":8080")
}
