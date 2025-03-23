package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/babukre1/go-bookstore/pkg/models"
	"github.com/babukre1/go-bookstore/pkg/utils"
)

var NewBook models.Book



func CreateBook(w http.ResponseWriter, r *http.Request){
	
	Book := &models.Book{}
	
	utils.ParseBody(r, Book)

	b:= Book.CreateBook()

	res, _ := json.Marshal(b)

	w.WriteHeader(http.StatusOK)

	w.Write(res)


}


func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks:=models.GetAllBooks()

	res,_ := json.Marshal(newBooks)

	w.Header().Set("Content-Type","pkglication/json")

	w.WriteHeader(http.StatusOK)

	w.Write(res)
}



func GetBookById(w http.ResponseWriter, r *http.Request){

}

func UpdateBook(w http.ResponseWriter, r *http.Request){

}
func DeleteBook(w http.ResponseWriter, r *http.Request){

}