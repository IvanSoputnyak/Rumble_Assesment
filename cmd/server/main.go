package main

import (
	"flag"
	"log"
	"net/http"

	"assesment/internal/transactions"
)

func main() {
	var source string
	flag.StringVar(&source, "transactions", "", "URL or filepath to transactions JSON")
	flag.Parse()
	if source == "" {
		log.Fatal("missing required flag: --transactions")
	}

	mux := http.NewServeMux()
	mux.Handle("/transactions", transactions.GetTransactionsHandler(source))
	mux.Handle("/transactions/ordered", transactions.GetTransactionsOrderedHandler(source))

	log.Println("listening on :8000")
	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatal(err)
	}
}
