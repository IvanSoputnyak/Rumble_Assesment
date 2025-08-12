package transactions

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

const maxChunkSize int64 = 5 * 1024 * 1024 // 5 MB

func loadFileByURL(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("failed to fetch file: %s", resp.Status)
	}

	return struct {
		io.Reader
		io.Closer
	}{
		Reader: io.LimitReader(resp.Body, maxChunkSize),
		Closer: resp.Body,
	}, nil
}

func loadFileLocal(path string) (io.ReadCloser, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return struct {
		io.Reader
		io.Closer
	}{
		Reader: io.LimitReader(file, maxChunkSize),
		Closer: file,
	}, nil
}

func LoadTransactions(source string) ([]Transaction, error) {
	var reader io.ReadCloser
	u, err := url.Parse(source)
	if err == nil && (u.Scheme == "http" || u.Scheme == "https") {
		reader, err = loadFileByURL(source)
	} else {
		reader, err = loadFileLocal(source)
	}
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	decoder := json.NewDecoder(reader)
	var transactions []Transaction
	err = decoder.Decode(&transactions)
	if err != nil {
		return nil, fmt.Errorf("failed to decode transaction")
	}

	return transactions, nil
}
