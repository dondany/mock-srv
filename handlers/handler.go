package handlers

import (
	"encoding/json"
	"mocksrv/db"
	"net/http"
	"strconv"
)

type Handler struct {
	Collection string
	DB         *db.Database
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	result, err := h.DB.FindAll(h.Collection, query)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	response, err := json.MarshalIndent(result, "", " ")
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseFloat(idStr, 64)
	if err != nil {
		http.Error(w, "Wrong id format", http.StatusBadRequest)
		return
	}

	result, err := h.DB.FindById(h.Collection, id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	response, err := json.MarshalIndent(result, "", " ")
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	var object map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&object)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.DB.Insert(h.Collection, object)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseFloat(idStr, 64)
	if err != nil {
		http.Error(w, "Wrong id format", http.StatusBadRequest)
		return
	}

	var object map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&object)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updated, err := h.DB.Update(h.Collection, id, object)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response, err := json.MarshalIndent(updated, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(response)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseFloat(idStr, 64)
	if err != nil {
		http.Error(w, "Wrong id format", http.StatusBadRequest)
		return
	}

	err = h.DB.Delete(h.Collection, id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
