package handlers

import (
	"bytes"
	"net/http/httptest"
	"testing"
)

func BenchmarkServiceHandler(b *testing.B) {

	handler := NewServiceHandler(NewServiceSubHandler())

	for i := 0; i < b.N; i++ {
		r := httptest.NewRequest("POST", "/product", bytes.NewReader([]byte(`{"product":"Soap"}`)))
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, r)
	}
}