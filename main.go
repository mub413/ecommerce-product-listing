package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID          int     `json:""`
	Title       string  `json:"title"`
	Description string  `json:"description "`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

var productList []Product

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I am Mohi Uddin, I'm a new go learner")
}
func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		http.Error(w, "Please give me GET request", 400)
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(productList)

}

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "Please give me POST request", 400)
		return
	}
	var newProduct Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid json", 400)
	}

	newProduct.ID = len(productList) + 1

	productList = append(productList, newProduct)

	w.WriteHeader(201)
	encoder := json.NewEncoder(w)
	encoder.Encode(newProduct)

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/products", getProducts)
	mux.HandleFunc("/create", createProduct)
	fmt.Println("Server running on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is red",
		Price:       100,
		ImgUrl:      "https://i.chaldn.com/_mpimage/komola-orange-imported-50-gm-1-kg?src=https%3A%2F%2Feggyolk.chaldal.com%2Fapi%2FPicture%2FRaw%3FpictureId%3D64292&q=best&v=1&m=1600",
	}
	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is Green",
		Price:       150,
		ImgUrl:      "https://assets.clevelandclinic.org/transform/LargeFeatureImage/cd71f4bd-81d4-45d8-a450-74df78e4477a/Apples-184940975-770x533-1_jpg",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is Yellow",
		Price:       50,
		ImgUrl:      "https://www.orchardfood.co.za/cdn/shop/files/Bananas-2.png?v=1753697464&width=1946",
	}
	prd4 := Product{
		ID:          4,
		Title:       "Lichu",
		Description: "Lichu is red",
		Price:       10,
		ImgUrl:      "https://media.gettyimages.com/id/566454679/photo/lychee-fruits.jpg?s=1024x1024&w=gi&k=20&c=mb13OLKoGK9QZYHi3j5nMgBfkuhSMue19wMRCrLWy5g=",
	}

	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
	productList = append(productList, prd4)

}
