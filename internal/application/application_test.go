package application_test

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"net/http/httptest"

	"github.com/Gromov2009/calc_go/internal/application"
)

func TestRequestHandlerSuccessSimpleCase(t *testing.T) {
	expected := `{ "result": "11.000000" }`
	bodyReader := strings.NewReader(`{"expression": "5+6"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bodyReader)

	w := httptest.NewRecorder()
	application.CalcHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if string(data) != expected {
		t.Errorf("Expected %s but got %s", expected, string(data))
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("wrong status code")
	}

	//if strings.TrimSpace(rr.Body.String()) != tc.expectedBody {
}

func TestRequestHandlerSuccessCase(t *testing.T) {
	expected := `{ "result": "-52.000000" }`
	bodyReader := strings.NewReader(`{"expression": "33/11-5*(5+6)"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bodyReader)

	w := httptest.NewRecorder()
	application.CalcHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if string(data) != expected {
		t.Errorf("Expected %s but got %s", expected, string(data))
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("wrong status code")
	}

	//if strings.TrimSpace(rr.Body.String()) != tc.expectedBody {
}

func TestRequestHandlerEmptyBodyCase(t *testing.T) {

	bodyReader := strings.NewReader("")
	req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bodyReader)
	w := httptest.NewRecorder()
	application.CalcHandler(w, req)
	res := w.Result()

	if res.StatusCode != http.StatusUnprocessableEntity {
		t.Fatalf("request is invalid but Status %s was obtained", res.Status)
	}

}

func TestRequestHandlerBadRequestCase(t *testing.T) {

	bodyReader := strings.NewReader(`{"expression": "(5+6"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bodyReader)

	w := httptest.NewRecorder()
	application.CalcHandler(w, req)

	res := w.Result()
	fmt.Println(res)
	if res.StatusCode != http.StatusUnprocessableEntity {
		t.Fatalf("request is invalid but Status %s was obtained", res.Status)
	}

}

func TestRequestHandlerBadSymbolsCase(t *testing.T) {

	bodyReader := strings.NewReader(`{"expression": "3*(5+qwer6)"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bodyReader)

	w := httptest.NewRecorder()
	application.CalcHandler(w, req)

	res := w.Result()
	fmt.Println(res)
	if res.StatusCode != http.StatusUnprocessableEntity {
		t.Fatalf("request is invalid but Status %s was obtained", res.Status)
	}

}

func TestRequestHandlerDivByErrorCase(t *testing.T) {

	bodyReader := strings.NewReader(`{"expression": "5+6/0"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bodyReader)

	w := httptest.NewRecorder()

	application.CalcHandler(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusUnprocessableEntity {
		t.Fatalf("request is invalid but Status %s was obtained", res.Status)
	}

}
