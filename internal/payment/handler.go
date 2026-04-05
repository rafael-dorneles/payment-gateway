package payment

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/rafael-dorneles/payment-gateway/internal/models"
)

type PaymentHandle struct {
	service PaymentService
}

func NewPaymentHandle(s PaymentService) *PaymentHandle {
	return &PaymentHandle{
		service: s,
	}
}

func (h *PaymentHandle) Create(w http.ResponseWriter, r *http.Request) {

	var record models.PaymentRequest

	if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
		http.Error(w, "Json inválido", http.StatusBadRequest)
		return
	}

	transaction, err := h.service.Create(record)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transaction)
}

func (h *PaymentHandle) UpdateStatus(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")
	id, _ := uuid.Parse((idParam))

	var body struct {
		Status string `json : status`
	}
	json.NewDecoder(r.Body).Decode(&body)

	err := h.service.UpdateStatus(id, body.Status)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (h *PaymentHandle) GetById(c *gin.Context) {

	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "ID de transacao invalido"})
	}

	transaction, err := h.service.GetById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pagamento não encontrado"})
		return
	}

	c.JSON(http.StatusOK, transaction)

}
