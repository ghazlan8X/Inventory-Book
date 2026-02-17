package app

import (
	"BelajarGolang5/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) handler {
	return handler{DB: db}
}

func (h *handler) GetBooks(c *gin.Context) {
	var books []models.Books

	h.DB.Find(&books)
	c.HTML(http.StatusOK, "book.index", gin.H{
		"PageTitle": "Home Page",
		"payload":   books,
	})
}

func (h *handler) GetBookById(c *gin.Context) {
	BookId := c.Param("id")
	var books models.Books

	if h.DB.Find(&books, "id=?", BookId).RecordNotFound() {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.HTML(http.StatusOK, "book.detail", gin.H{
		"PageTitle": books.Title,
		"payload":   books,
		"auth":      c.Query("auth"),
	})
}

func (h *handler) AddBook(c *gin.Context) {
	c.HTML(http.StatusOK, "book.form", gin.H{
		"PageTitle": "Add Book",
		"auth":      c.Query("auth"),
	})
}

func (h *handler) PostBook(c *gin.Context) {
	var book models.Books

	c.ShouldBind(&book)
	h.DB.Create(&book)

	c.Redirect(http.StatusSeeOther, "/books")
}

func (h *handler) UpdateBook(c *gin.Context) {
	var book models.Books

	bookId := c.Param("id")
	if h.DB.Find(&book, "id=?", bookId).RecordNotFound() {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": "Not Found",
		})
		return
	}

	c.HTML(http.StatusOK, "book.form", gin.H{
		"PageTitle": "Add Book",
		"payload":   book,
		"auth":      c.Query("auth"),
	})

}

func (h *handler) PutBook(c *gin.Context) {
	var book models.Books

	bookId := c.Param("id")
	if h.DB.Find(&book, "id=?", bookId).RecordNotFound() {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": "Not Found",
		})
		return
	}

	var ReqBook = book
	c.ShouldBind(&ReqBook)

	h.DB.Model(&book).Where("id=?", bookId).Update(&ReqBook)

	c.Redirect(http.StatusSeeOther, "/books")
}

func (h *handler) DeleteBook(c *gin.Context) {
	var book models.Books
	bookId := c.Param("id")

	h.DB.Where("id=?", bookId).Delete(&book)

	c.Redirect(http.StatusSeeOther, "/books")
}
