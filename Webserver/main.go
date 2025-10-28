package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"webser/dbopr"

	"gorm.io/gorm"
)

var (
	items = make(map[string]Item)
	mu    sync.Mutex // Mutex to protect concurrent access to 'items' map
)

type Item struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func createItem(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// mu.Lock()
	// defer mu.Unlock()
	// if _, exists := items[newItem.ID]; exists {
	// 	http.Error(w, "Item with this ID already exists", http.StatusConflict)
	// 	return
	// }
	// items[newItem.ID] = newItem

	db := dbopr.DbConn()
	res := db.Create(&newItem)
	log.Printf("Created item: %+v\n", newItem)
	log.Println("res ", res)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)
}

func getAllItems(w http.ResponseWriter, r *http.Request) {
	// mu.Lock()
	// defer mu.Unlock()
	// all := make([]Item, 0, len(items))
	// for _, item := range items {
	// 	all = append(all, item)
	// }

	var listItem []Item
	db := dbopr.DbConn()
	db.Raw("select * from items").Scan(&listItem)
	//db.Scan(&listItem)

	for _, val := range listItem {
		fmt.Println("Val ", val)
	}
	json.NewEncoder(w).Encode(listItem)
}

func getItem(w http.ResponseWriter, r *http.Request, id string) {

	var itemByID Item
	db := dbopr.DbConn()

	db.First(&itemByID, id)

	log.Printf("Item by ID: %+v\n", itemByID)

	// mu.Lock()
	// defer mu.Unlock()
	// item, found := items[id]
	// if !found {
	// 	http.Error(w, "Item not found", http.StatusNotFound)
	// 	return
	// }

	if itemByID == (Item{}) {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(itemByID)
}

func updateItem(w http.ResponseWriter, r *http.Request, id string) {
	var updatedItem Item
	err := json.NewDecoder(r.Body).Decode(&updatedItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := dbopr.DbConn()
	var tmpItem Item

	//db.Model(&Item{}).Where("id=?", id).Updates(updatedItem)
	db.First(&tmpItem, id)
	db.Model(&tmpItem).Updates(Item{Name: updatedItem.Name, Price: updatedItem.Price})

	// mu.Lock()
	// defer mu.Unlock()
	// if _, found := items[id]; !found {
	// 	http.Error(w, "Item not found", http.StatusNotFound)
	// 	return
	// }
	// updatedItem.ID = id // Ensure the ID from the URL is used
	// items[id] = updatedItem
	log.Println("Record updated")
	json.NewEncoder(w).Encode(updatedItem)
}

func deleteItem(w http.ResponseWriter, r *http.Request, id string) {
	// mu.Lock()
	// defer mu.Unlock()
	// if _, found := items[id]; !found {
	// 	http.Error(w, "Item not found", http.StatusNotFound)
	// 	return
	// }
	// delete(items, id)
	db := dbopr.DbConn()

	db.Delete(&Item{}, id)
	log.Println("Record deleted")

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	http.HandleFunc("/items", itemsHandler)
	http.HandleFunc("/items/", itemHandler) // For specific item operations

	db := dbopr.DbConn()
	dbconn(db)
	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func dbconn(db *gorm.DB) {
	/* defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Printf("Error getting underlying *sql.DB: %v", err)
			return
		}
		if err := sqlDB.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}()

		newItem := Item{ID: "100", Name: "Siva", Price: 70}
		db.Create(&newItem)
		log.Printf("Created item: %+v\n", newItem)

	*/

	// Query a user by ID
	var itemByID Item
	db.First(&itemByID, 100)
	log.Printf("Item by ID: %+v\n", itemByID)

}

// itemsHandler handles requests to /items (GET for all, POST for create)
func itemsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAllItems(w, r)
	case "POST":
		createItem(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// itemHandler handles requests to /items/{id} (GET, PUT, DELETE)
func itemHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/items/"):] // Extract ID from URL

	switch r.Method {
	case "GET":
		getItem(w, r, id)
	case "PUT":
		updateItem(w, r, id)
	case "DELETE":
		deleteItem(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
