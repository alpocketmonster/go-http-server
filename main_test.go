package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHomepageHandler(t *testing.T) {
	router, _ := setupRouter()

	request, _ := http.NewRequest("GET", "/auth", nil)
	request.Header.Set("Content-Type", "application/vnd.kafka.avro.v2+json")
	request.Header.Set("X-Original-Uri", "000-0.sap-erp.db.operations.orders05.0")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)

	//responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
}

func BenchmarkTestHomepageHandler(b *testing.B) {
	router, _ := setupRouter()
	request, _ := http.NewRequest("GET", "/auth", nil)
	request.Header.Set("Content-Type", "application/vnd.kafka.avro.v2+json")
	request.Header.Set("X-Original-Uri", "000-0.sap-erp.db.operations.orders05.0")
	w := httptest.NewRecorder()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		router.ServeHTTP(w, request)
	}
}

func TestControllerWrongContentType(t *testing.T) {
	router, _ := setupRouter()
	request, _ := http.NewRequest("GET", "/auth", nil)
	request.Header.Set("X-Original-Uri", "000-0.sap-erp.db.operations.orders05.0")
	request.Header.Set("Content-Type", "wrong")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestControllerWrongURL(t *testing.T) {
	router, _ := setupRouter()
	request, _ := http.NewRequest("GET", "/auth", nil)
	request.Header.Set("Content-Type", "application/vnd.kafka.avro.v2+json")
	request.Header.Set("X-Original-Uri", "wrong")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
