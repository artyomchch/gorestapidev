package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Part1 struct {
	StaticID string `json:"staticid"`
	Root     string `json:"root"`
	Model    string `json:"model"`
	System   string `json:"system"`
	IMEI     string `json:"imei"`
}

type Part2 struct {
	ApkId         string `json:"apkId"`
	ApkName       string `json:"apk_name"`
	ApkFullName   string `json:"apkFullName"`
	ApkVersion    string `json:"apk_version"`
	ApkPath       string `json:"apk_path"`
	ApkPermission string `json:"apk_permission"`
}

type Output struct {
	ApkId         string `json:"apkId"`
	ApkName       string `json:"apk_name"`
	ApkFullName   string `json:"apkFullName"`
	ApkVersion    string `json:"apk_version"`
	ApkProtection string `json:"apk_protection"`
	ApkTitle      string `json:"apk_title"`
}

var books []Book
var parts1 []Part1
var parts2 []Part2

//var parts3 []Part3
var outputs []Output

func main() {
	r := mux.NewRouter()
	//books = append(books, Book{ID: "1", Title: "Война и Мир", Author: &Author{Firstname: "Лев", Lastname: "Толстой"}})
	//books = append(books, Book{ID: "2", Title: "Преступление и наказание", Author: &Author{Firstname: "Фёдор", Lastname: "Достоевский"}})
	//parts1 = append(parts1, Part1{StaticID: "1", Root: "1", Model: "Xiaomi Mi A2 Lite", System: "Android 9", IMEI: "12435364575656234"})
	//	r.HandleFunc("books/{id}", updateBook).Methods("PUT")
	//	r.HandleFunc("books/{id}", deleteBook).Methods("DELETE")

	//r.HandleFunc("/parts1", getParts1).Methods("GET")
	//r.HandleFunc("/parts1/{id}", getPart1).Methods("GET")
	//r.HandleFunc("/parts1", createPart1).Methods("POST")
	//r.HandleFunc("/parts2", getParts2).Methods("GET")
	//r.HandleFunc("/parts2/{id}", getPart2).Methods("GET")
	//r.HandleFunc("/parts2", createPart2).Methods("POST")
	//r.HandleFunc("/outputs", getOutputs).Methods("GET")
	//r.HandleFunc("/outputs/{id}", getOutput).Methods("GET")
	//r.HandleFunc("/outputs", createOutput).Methods("POST")

	//log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, r) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}

	//r := grace.Serve(":" + port, context.ClearHandler(http.DefaultServeMux))
}

/////////////PART1/////////////

func getParts1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(parts1)
}

func getPart1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range parts1 {
		if item.StaticID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Part1{})
}

func createPart1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var part1 Part1
	_ = json.NewDecoder(r.Body).Decode(&part1)
	part1.StaticID = strconv.Itoa(rand.Intn(1000000))
	parts1 = append(parts1, part1)
	json.NewEncoder(w).Encode(part1)
}

/////////////PART2/////////////

func getParts2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(parts2)

}

func getPart2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range parts2 {
		if item.ApkId == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Part2{})
}

func createPart2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var part2 Part2
	_ = json.NewDecoder(r.Body).Decode(&part2)
	part2.ApkId = strconv.Itoa(rand.Intn(1000000))
	parts2 = append(parts2, part2)
	json.NewEncoder(w).Encode(part2)
}

/////////////OUTPUT/////////////

func getOutputs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(outputs)
}

func getOutput(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range outputs {
		if item.ApkId == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Output{})
}

func createOutput(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var output Output
	_ = json.NewDecoder(r.Body).Decode(&output)
	output.ApkId = strconv.Itoa(rand.Intn(1000000))
	outputs = append(outputs, output)
	json.NewEncoder(w).Encode(output)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}
