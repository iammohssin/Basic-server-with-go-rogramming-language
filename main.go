package main
import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

type Item struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

var items []Item

func getItems(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(items)
}

func getItem(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r) // Get route parameters
    for _, item := range items {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(Item{})
}

func createItem(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var newItem Item
    _ = json.NewDecoder(r.Body).Decode(&newItem)
    newItem.ID = fmt.Sprintf("%d", len(items)+1)
    items = append(items, newItem)
    json.NewEncoder(w).Encode(newItem)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range items {
        if item.ID == params["id"] {
            var updatedItem Item
            _ = json.NewDecoder(r.Body).Decode(&updatedItem)
            updatedItem.ID = params["id"]
            items[index] = updatedItem
            json.NewEncoder(w).Encode(updatedItem)
            return
        }
    }
    json.NewEncoder(w).Encode(Item{})
}

func main() {
    r := mux.NewRouter()
    items = append(items, Item{ID: "1", Name: "Item 1"})
    items = append(items, Item{ID: "2", Name: "Item 2"})

    r.HandleFunc("/items", getItems).Methods("GET")
    r.HandleFunc("/items/{id}", getItem).Methods("GET")
    r.HandleFunc("/items", createItem).Methods("POST")
    r.HandleFunc("/items/{id}", updateItem).Methods("PUT")
	fmt.Print("Server running")
    log.Fatal(http.ListenAndServe(":8001", r))

}
