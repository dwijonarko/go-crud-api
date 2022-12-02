package test

import (
	"go-crud-api/controllers"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
}

func NewClient(httpClient *http.Client, baseURL string) Client {
	return Client{
		httpClient: httpClient,
		baseURL:    baseURL,
	}
}

func TestGetProduct(t *testing.T) {
	router := http.NewServeMux()
	router.HandleFunc("/api/products", controllers.GetProducts)

	svr := httptest.NewServer(router)
	defer svr.Close()
	request, _ := http.NewRequest("GET", "api/productas", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
