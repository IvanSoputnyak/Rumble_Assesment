package transactions

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

const rootJSONPath = "../../transactions.json"

func readTransactionsFromRoot(t *testing.T) []Transaction {
	t.Helper()
	f, err := os.Open(rootJSONPath)
	if err != nil {
		t.Fatalf("open %s: %v", rootJSONPath, err)
	}
	defer f.Close()
	var txs []Transaction
	if err := json.NewDecoder(f).Decode(&txs); err != nil {
		t.Fatalf("decode %s: %v", rootJSONPath, err)
	}
	return txs
}

func TestGetTransactionsHandlerFromFile_Success(t *testing.T) {
	orig := readTransactionsFromRoot(t)

	h := GetTransactionsHandler(rootJSONPath)

	req := httptest.NewRequest(http.MethodGet, "/transactions", nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("status=%d, want %d; body=%s", rr.Code, http.StatusOK, rr.Body.String())
	}
	if ct := rr.Header().Get("Content-Type"); ct != "application/json" {
		t.Fatalf("Content-Type=%q, want application/json", ct)
	}

	var out []TransactionMasked
	if err := json.Unmarshal(rr.Body.Bytes(), &out); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(out) != len(orig) {
		t.Fatalf("len=%d, want %d (as in %s)", len(out), len(orig), rootJSONPath)
	}
}

func TestGetTransactionsHandler_LoadError(t *testing.T) {
	t.Parallel()

	bad := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "boom", http.StatusInternalServerError)
	}))
	defer bad.Close()

	h := GetTransactionsHandler(bad.URL)

	req := httptest.NewRequest(http.MethodGet, "/transactions", nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Fatalf("status=%d, want %d", rr.Code, http.StatusInternalServerError)
	}
}

func TestGetTransactionsOrderedHandler_SortsAndMasks(t *testing.T) {
	t.Parallel()

	h := GetTransactionsOrderedHandler(rootJSONPath)

	req := httptest.NewRequest(http.MethodGet, "/transactions/ordered", nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("status=%d, want %d; body=%s", rr.Code, http.StatusOK, rr.Body.String())
	}

	var out []TransactionMasked
	if err := json.Unmarshal(rr.Body.Bytes(), &out); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(out) < 2 {
		t.Skip("need at least two transactions in transactions.json")
	}

	t0, e0 := time.Parse(time.RFC3339, out[0].PostedTimeStamp)
	t1, e1 := time.Parse(time.RFC3339, out[1].PostedTimeStamp)

	switch {
	case e0 == nil && e1 == nil:
		if t0.Before(t1) {
			t.Fatalf("top-2 not in desc order: %s < %s", t0, t1)
		}
	case e0 == nil && e1 != nil:
		// we are good here, invalid timestamp we keep below
	case e0 != nil && e1 == nil:
		t.Fatalf("first has invalid timestamp but second is valid; invalid must come after valid")
	default:
		// both invalid
	}
}

func TestGetTransactionsOrderedHandler_LoadError(t *testing.T) {
	t.Parallel()

	bad := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "boom", http.StatusInternalServerError)
	}))
	defer bad.Close()

	h := GetTransactionsOrderedHandler(bad.URL)

	req := httptest.NewRequest(http.MethodGet, "/transactions/ordered", nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Fatalf("status=%d, want %d", rr.Code, http.StatusInternalServerError)
	}
}
