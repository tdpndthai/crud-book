package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

var books []Book //danh sách book

func main() {
	router := mux.NewRouter()

	books = append(books,
		Book{ID: "1", Title: "Golang pointers", Author: "Mr. Golang", Year: "2010"},
		Book{ID: "2", Title: "Goroutines", Author: "Mr. Goroutine", Year: "2011"},
		Book{ID: "3", Title: "Golang routers", Author: "Mr. Router", Year: "2012"},
		Book{ID: "4", Title: "Golang concurrency", Author: "Mr. Currency", Year: "2013"},
		Book{ID: "5", Title: "Golang good parts", Author: "Mr. Good", Year: "2014"})

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/book/{id}", getBook).Methods("GET")
	router.HandleFunc("/addbook", addBook).Methods("POST")
	router.HandleFunc("/updateBook", updateBook).Methods("PUT")
	router.HandleFunc("/delbook/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4560", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books) //chuyển đổi thành file json
	// log.Fatal("get all books")
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) //return về 1 map(tập hợp key,value)--cụ thể ở đây trả về map[id:3 ],Vars trả về 1 biến định tuyến
	//fmt.Print(params)
	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(&book) //trỏ tới địa chỉ của book tìm được
			//fmt.Println(book)
			//fmt.Println(&book)
		}
	}
	// log.Fatal("get book by id")
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	//giải mã body json
	json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(w).Encode(&books)
	fmt.Print(books)
	//{1 Golang pointers Mr. Golang 2010} {2 Goroutines Mr. Goroutine 2011} {3 Golang routers Mr. Router 2012} {4 Golang concurrency Mr. Currency 2013} {5 Golang good parts Mr. Good 2014}]
	// log.Fatal("add new book")
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	json.NewDecoder(r.Body).Decode(&book) //decode giải mã đọc giá trị json
	for i, item := range books {
		if item.ID == book.ID {
			books[i] = book
		}
	}
	json.NewEncoder(w).Encode(&books)
	//log.Fatal("update book")
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i,item := range books{
		if item.ID == params["id"] {
			books = append(books[:i],books[i+1:]...) //books[:i] tạo mới slices
		}else {
			fmt.Println("không tìm thấy book")
		}
	}
	json.NewEncoder(w).Encode(books)
	
	//log.Fatal("delete book by id")
}
