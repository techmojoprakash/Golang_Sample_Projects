package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"stocksapi/models"
)

type StockHandlerImpl struct {
	Service models.StockService
}

func NewStockHandler(ser models.StockService) models.StockHandler {
	return &StockHandlerImpl{
		Service: ser,
	}
}

func (h *StockHandlerImpl) ListStocksHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("ListStocksHandler: Fetching all stocks")
	w.Header().Set("Content-Type", "application/json")

	getStocks, err := h.Service.ListStocks()
	if err != nil {
		log.Println("ListStocksHandler:Unable to fetch the stocks", err)
		respondWithError(w, http.StatusInternalServerError, "Unable to fetch stocks")
	}
	encode_err := json.NewEncoder(w).Encode(getStocks)
	if encode_err != nil {
		log.Println("ListStocksHandler:Unable to encode the stocks", err)
		respondWithError(w, http.StatusInternalServerError, "Unable to encode the stocks")
	}
}

func (h *StockHandlerImpl) GetStockHandler(w http.ResponseWriter, r *http.Request) {

}
func (h *StockHandlerImpl) CreateStockHandler(w http.ResponseWriter, r *http.Request) {

}
func (h *StockHandlerImpl) UpdateStockHandler(w http.ResponseWriter, r *http.Request) {

}
func (h *StockHandlerImpl) DeleteStockHandler(w http.ResponseWriter, r *http.Request) {

}

func respondWithError(w http.ResponseWriter, statuscode int, error_msg string) {
	w.WriteHeader(statuscode)
	err := json.NewEncoder(w).Encode(map[string]string{
		"error": error_msg,
	})
	if err != nil {
		log.Printf("respondWithError:Unable to encode error message")
	}
}
