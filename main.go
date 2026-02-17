package main

import (
	apps "BelajarGolang5/app"
	"BelajarGolang5/auth"
	"BelajarGolang5/db"
	"BelajarGolang5/middleware"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	conn := db.InitDB()

	router := gin.Default()
	// router.LoadHTMLGlob("resource/**/*")
	router.LoadHTMLGlob("resource/**/*.html")

	handler := apps.New(conn)
	authH := auth.New(conn)

	// Home
	router.GET("/", auth.HomeHandler)

	// Login
	router.GET("/login", auth.LoginGetHandler)
	router.POST("/login", authH.LoginPostHandler)

	// logout
	router.GET("/logout", auth.Logout)

	// get all books
	router.GET("/books", middleware.AuthValid, handler.GetBooks)

	// detail book
	router.GET("/book/:id", middleware.AuthValid, handler.GetBookById)

	// get and post book
	router.GET("/addBook", middleware.AuthValid, handler.AddBook)
	router.POST("/book", middleware.AuthValid, handler.PostBook)

	// edit and put
	router.GET("/updateBook/:id", middleware.AuthValid, handler.UpdateBook)
	router.POST("/updateBook/:id", middleware.AuthValid, handler.PutBook)

	// delete
	router.POST("/deleteBook/:id", middleware.AuthValid, handler.DeleteBook)

	router.Run()
}
