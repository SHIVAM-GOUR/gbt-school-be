package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/SHIVAM-GOUR/gbt-school-be/internal/models"
	"github.com/go-chi/chi/v5"
	"github.com/jinzhu/gorm"
)

type ClassHandler struct {
	DB *gorm.DB
}

func (h *ClassHandler) Create(w http.ResponseWriter, r *http.Request) {
	var class models.Class
	if err := json.NewDecoder(r.Body).Decode(&class); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.DB.Create(&class).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(class)
}

func (h *ClassHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var class models.Class
	if err := h.DB.First(&class, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(class)
}

func (h *ClassHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var class models.Class
	if err := h.DB.First(&class, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	var update models.Class
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	class.Name = update.Name
	if err := h.DB.Save(&class).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(class)
}

func (h *ClassHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := h.DB.Delete(&models.Class{}, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *ClassHandler) List(w http.ResponseWriter, r *http.Request) {
	var classes []models.Class
	if err := h.DB.Find(&classes).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(classes)
}
