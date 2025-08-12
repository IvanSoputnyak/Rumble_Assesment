package transactions

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetTransactionsHandler(source string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		transactions, err := LoadTransactions(source)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to get transactions: %v", err), http.StatusInternalServerError)
			return
		}
		result := MaskAll(transactions)

		b, err := json.Marshal(result)
		if err != nil {
			http.Error(w, "encode error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(b)
	})
}

func GetTransactionsOrderedHandler(source string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		transactions, err := LoadTransactions(source)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to get transactions: %v", err), http.StatusInternalServerError)
			return
		}
		result := MaskAll(OrderByPostedTimestampDesc(transactions))

		b, err := json.Marshal(result)
		if err != nil {
			http.Error(w, "encode error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(b)
	})
}
