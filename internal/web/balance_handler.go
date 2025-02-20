package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/usecase/get_balance"
	"github.com/go-chi/chi/v5"
)

type WebBalanceHandler struct {
	GetBalanceUseCase get_balance.GetBalanceUseCase
}

func NewWebBalanceHandler(getBalanceUseCase get_balance.GetBalanceUseCase) *WebBalanceHandler {
	return &WebBalanceHandler{
		GetBalanceUseCase: getBalanceUseCase,
	}
}

func (h *WebBalanceHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	account_id := chi.URLParam(r, "account_id")

	output, err := h.GetBalanceUseCase.Execute(account_id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println("dkdasçpld223")
		fmt.Println(err)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}
