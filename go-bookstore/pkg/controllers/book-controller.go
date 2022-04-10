package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/anishgowda21/go-projects/go-bookstore/pkg/models"
	"github.com/anishgowda21/go-projects/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)


var NewBook models.Book

func GetBook(w http.ResponseWriter,r *http.Request){
	newBooks:=models.GetAllBooks()
	res,_ := json.Marshal(newBooks)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func GetBookByID(w http.ResponseWriter,r *http.Request){
	vars:= mux.Vars(r)
	bookId := vars["bookId"]
	ID, err:= strconv.ParseInt(bookId,0,0)
	if err!=nil{
		fmt.Println("Error while parsing")
	}
	bookDeails,_ :=models.GetBookByID(ID)
	res,_ := json.Marshal(bookDeails)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res) 
}


func Createbook(w http.ResponseWriter, r *http.Request){
	cb := &models.Book{}
	utils.ParseBody(r,cb)
	b:=cb.CreateBook()
	res,_ := json.Marshal(b)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars:= mux.Vars(r)
	bookId := vars["bookId"]
	ID, err:= strconv.ParseInt(bookId,0,0)
	if err!=nil{
		fmt.Println("Error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	ub := &models.Book{}
	utils.ParseBody(r,ub)
	vars:= mux.Vars(r)
	bookId := vars["bookId"]
	ID, err:= strconv.ParseInt(bookId,0,0)
	if err!=nil{
		fmt.Println("Error while parsing")
	}
	bookDetails,db := models.GetBookByID(ID)
	if ub.Name != ""{
		bookDetails.Name = ub.Name
	}
	if ub.Author != ""{
		bookDetails.Author = ub.Author
	}
	if ub.Publication != ""{
		bookDetails.Publication = ub.Publication
	}
	db.Save(&bookDetails)
	res,_ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

