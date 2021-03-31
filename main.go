package main

import(
	"log"
	"net/http"
	"encoding/json"
	"strconv"
	"math/rand"

	"github.com/gorilla/mux"
)

type Book struct{
	ID 		string 	`json:"id"`
	Isbn 	string 	`json:"isbn"`
	Title 	string 	`json:"title`
	Author 	*Author	`json:"author`
}

type Author struct{
	Fullname 	string `json:"fullname"`
	Country 	string `json:"country"`
}

//Init books vas as slice Book Struct
var books []Book

//Get All Books
func getBooks (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)

}
//Get single book
func getBook (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get Params
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}
//Create a New Book
func createBook (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(&book)
}
//Update Book
func updateBook (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get Params
	for index, item := range books{
		if item.ID == params["id"]{
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(&book)			
			return
		}	
	}
	json.NewEncoder(w).Encode(books)
}
//Delete Book
func deleteBook (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get Params
	for index, item := range books{
		if item.ID == params["id"]{
			books = append(books[:index], books[index+1:]...)
			break
		}	
	}
	json.NewEncoder(w).Encode(books)	
}
func main(){
	//Init Router
	r := mux.NewRouter()

	//mock data - implement DB
	book1 := Book{
		ID: "1",
		Isbn: "12345",
		Title: "Book one",
		Author: &Author{
			Fullname: "John Doe",
			Country: "USA",
		},
	}
	books = append(books, book1)
	books = append(books, Book{ID: "2", Isbn: "98765", Title: "Book two", Author: &Author{
		Fullname: "Steve Smith", Country: "England"}})

	//Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))
}
