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
	Id   string  `json:"id"`
	Apps []*Apps `json:"apps"`
}

type Apps struct {
	ApkId int    `json:"apkId"`
	App   []*App `json:"app"`
}

type App struct {
	ApkName       string   `json:"nameApp"`
	ApkFullName   string   `json:"packageName"`
	ApkVersion    string   `json:"versionApp"`
	ApkPath       string   `json:"pathApp"`
	ApkPermission []string `json:"permissionArray"`
}

type Output struct {
	ApkId         string `json:"apkId"`
	ApkName       string `json:"apk_name"`
	ApkFullName   string `json:"apkFullName"`
	ApkVersion    string `json:"apk_version"`
	ApkProtection string `json:"apk_protection"`
	ApkTitle      string `json:"apk_title"`
}

type DynaApps struct {
	ApkId                                      string        `json:"apkId"`
	ApkName                                    string        `json:"apk_name"`
	AndroidAccountsAccount                     []string      `json:"androidAccountsAccount"`
	AndroidServiceVoiceVoiceInteractionSession []string      `json:"androidServiceVoiceVoiceInteractionSession"`
	AndroidTelephonyPhoneStateListener         []string      `json:" androidTelephonyPhoneStateListener"`
	AndroidViewInputmethodBaseInputConnection  []string      `json:"androidViewInputmethodBaseInputConnection"`
	JavaLangReflectMethod                      []string      `json:"javaLangReflectMethod"`
	JavaIoFile                                 []string      `json:"javaIoFile"`
	JavaNetUri                                 []*JavaNetUri `json:"javaNetUri"`
}

type JavaNetUri struct {
	Time []string `json:"time"`
	Url  string   `json:"url"`
}

var books []Book

// part 1 - get id device
var parts1 []Part1

// part 2 static analyze
var parts2 []Part2
var apps []Apps
var app []App

// part3 dynamic analyze
var dynaApps []DynaApps
var javaNetUri []JavaNetUri

//var parts3 []Part3
//output
var outputs []Output

var appInt int = 0
var deviceInt int = 0

func main() {

	r := mux.NewRouter()
	//books = append(books, Book{ID: "1", Title: "Война и Мир", Author: &Author{Firstname: "Лев", Lastname: "Толстой"}})
	//books = append(books, Book{ID: "2", Title: "Преступление и наказание", Author: &Author{Firstname: "Фёдор", Lastname: "Достоевский"}})
	//parts1 = append(parts1, Part1{StaticID: "1", Root: "1", Model: "Xiaomi Mi A2 Lite", System: "Android 9", IMEI: "12435364575656234"})
	//	r.HandleFunc("books/{id}", updateBook).Methods("PUT")
	//	r.HandleFunc("books/{id}", deleteBook).Methods("DELETE")

	r.HandleFunc("/devices", getParts1).Methods("GET")
	r.HandleFunc("/device/{id}", getPart1).Methods("GET")
	r.HandleFunc("/devices", createPart1).Methods("POST")
	r.HandleFunc("/device/{id}/apps", getAppsOfDevice).Methods("GET")
	r.HandleFunc("/device/{id}/app/{id_app}", getCurrentIdOfApp).Methods("GET")
	r.HandleFunc("/appdyna", createCurrentIdOfApp).Methods("POST")
	r.HandleFunc("/apps", getParts2).Methods("GET")
	r.HandleFunc("/apps/{id}", getPart2).Methods("GET")
	r.HandleFunc("/apps", createPart2).Methods("POST")
	r.HandleFunc("/outputs", getOutputs).Methods("GET")
	r.HandleFunc("/outputs/{id}", getOutput).Methods("GET")
	r.HandleFunc("/outputs", createOutput).Methods("POST")

	//log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
	//log.Fatal(http.ListenAndServe(":8000", r))

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
	deviceInt++
	w.Header().Set("Content-Type", "application/json")
	var part1 Part1
	_ = json.NewDecoder(r.Body).Decode(&part1)
	part1.StaticID = strconv.Itoa(deviceInt)
	parts1 = append(parts1, part1)
	json.NewEncoder(w).Encode(part1)
}

//////////////////APPSofID////////
func getAppsOfDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range parts2 {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Part2{})
}

/////////CURRENT ID OF APP  ///////////////
func getCurrentIdOfApp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range dynaApps {
		if item.ApkId == params["id_app"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&DynaApps{})
}

func createCurrentIdOfApp(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var dynaApp DynaApps
	_ = json.NewDecoder(r.Body).Decode(&dynaApp)
	//part2.Id = strconv.Itoa(deviceInt)
	dynaApps = append(dynaApps, dynaApp)
	json.NewEncoder(w).Encode(dynaApp)
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
		if item.Id == params["id"] {
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
	part2.Id = strconv.Itoa(deviceInt)
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
