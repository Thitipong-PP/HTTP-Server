package server

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// Data Model
type Server struct {
	ID int `json:"id"`
	Name string `json:"name"`
	IP string `json:"ip"`
	Status bool `json:"status"`
}

// Mock Database
var serverList []Server // Slice save server list
var nxtID int = 1 // Next id

// Handler
func ServerHandler(w http.ResponseWriter, r *http.Request) {
	// MIME Type
	w.Header().Set("Content-Type", "application/json")

	// Get param
	param := strings.TrimPrefix(r.URL.Path, "/server/")
	id, _ := strconv.Atoi(param)

	// fmt.Println(param)
	// fmt.Println(id)

	// Choose method
	switch r.Method {
		case http.MethodGet :
			if id > 0 {
				// Get server at id
				getWithId(w, id)
			}else {
				// Get all server
				getAll(w)
			}
		case http.MethodPost :
			insert(w, r)
		case http.MethodDelete :
			delete(w, id)
		case http.MethodPut :
			update(w, r, id)
		default :
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
	}
}

// Search
func search (id int) int {
	left := 0
	right := len(serverList)-1
	for left<=right {
		mid := (left + right)/2
		
		// Found
		if id == serverList[mid].ID {
			return mid
		}else if id > serverList[mid].ID {
			left = mid+1
		}else {
			right = mid-1
		}
	}
	return -1
}

// Get
func getAll (w http.ResponseWriter) {
	if err := json.NewEncoder(w).Encode(serverList); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getWithId (w http.ResponseWriter, id int) {
	// Search id in server list
	idx := search(id)
	if idx != -1 {
		if err := json.NewEncoder(w).Encode(serverList[idx]); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}else {
		http.Error(w, "ID not found", http.StatusNotFound)
	}
}

// Post
func insert (w http.ResponseWriter, r *http.Request) {
	// Get body request
	var req Server
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req.ID = nxtID
	nxtID++

	// Insert to server list
	serverList = append(serverList, req)

	// Response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(req)
}

// Delete
func delete (w http.ResponseWriter, id int) {
	// Search id in server list
	idx := search(id)
	if idx != -1 {
		serverList = append(serverList[:idx], serverList[idx+1:]...)
		w.WriteHeader(http.StatusNoContent)
	}else {
		http.Error(w, "ID not found", http.StatusNotFound)
	}
}

// Update
func update (w http.ResponseWriter, r *http.Request, id int) {
	// Get body request
	var req Server
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req.ID = id

	idx := search(id)
	if idx != -1 {
		serverList[idx] = req
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(serverList[idx])
	}else {
		http.Error(w, "ID not found", http.StatusNotFound)
	}
}