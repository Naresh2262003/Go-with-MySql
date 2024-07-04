package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Naresh2262003/pkg/models"
	"github.com/Naresh2262003/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing")
	}

	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// func CreateBook(w http.ResponseWriter, r *http.Request) {
// 	CreateBook := &models.Book{}
// 	utils.ParseBody(r, CreateBook)
// 	b := CreateBook.CreateBook()
// 	res, _ := json.Marshal(b)
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

func CreateBook(w http.ResponseWriter, r *http.Request) {
	newBook := &models.Book{}
	utils.ParseBody(r, newBook)
	createdBook := newBook.CreateBook()
	res, _ := json.Marshal(createdBook)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// func DeleteBook(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	bookId := vars["bookId"]
// 	ID, err := strconv.ParseInt(bookId, 0, 0)
// 	if err != nil {
// 		fmt.Println("error while parsing")
// 	}
// 	book := models.DeleteBook(ID)
// 	res, _ := json.Marshal(book)
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	book, err := models.DeleteBook(ID)
	if err != nil {
		fmt.Println("error while deleting book:", err)
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var UpdateBook = &models.Book{}
	utils.ParseBody(r, UpdateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	booksDetails, db := models.GetBookById(ID)

	// booksDetails.Id = ID

	if UpdateBook.Name != "" {
		booksDetails.Name = UpdateBook.Name
	}
	if UpdateBook.Author != "" {
		booksDetails.Author = UpdateBook.Author
	}
	if UpdateBook.Publication != "" {
		booksDetails.Publication = UpdateBook.Publication
	}

	db.Save(&booksDetails)
	res, _ := json.Marshal(booksDetails)
	w.Header().Set("COntent-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
