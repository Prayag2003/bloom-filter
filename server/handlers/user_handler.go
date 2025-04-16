package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/Prayag2003/bloom-filter/bloom"
	"github.com/Prayag2003/bloom-filter/models"
	"github.com/Prayag2003/bloom-filter/storage"
)

type UserHandler struct {
	Filter  bloom.Filter
	Storage storage.Storage
	mu      sync.Mutex
}

func NewUserHandler(f bloom.Filter, s storage.Storage) *UserHandler {
	return &UserHandler{Filter: f, Storage: s}
}

func (h *UserHandler) LoadUsernames() {
	usernames, err := h.Storage.Load()
	if err != nil {
		log.Fatalf("❌ Failed to load usernames: %v", err)
	}
	for _, name := range usernames {
		h.Filter.Add(name)
	}
	log.Println("✅ Bloom filter initialized with usernames")
}

func (h *UserHandler) CheckUsername(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.UsernameRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	isAvailable := !h.Filter.Check(req.Username)
	log.Printf("🔍 Checked: %s → available: %v", req.Username, isAvailable)

	json.NewEncoder(w).Encode(struct {
		Available bool `json:"available"`
	}{Available: isAvailable})
}

func (h *UserHandler) RegisterUsername(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.UsernameRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	if h.Filter.Check(req.Username) {
		http.Error(w, "Username already taken", http.StatusConflict)
		log.Printf("❌ Username taken: %s", req.Username)
		return
	}

	h.Filter.Add(req.Username)
	if err := h.Storage.Save(req.Username); err != nil {
		http.Error(w, "Failed to save", http.StatusInternalServerError)
		return
	}
	log.Printf("✅ Registered: %s", req.Username)
	w.Write([]byte("Registered successfully"))
}
