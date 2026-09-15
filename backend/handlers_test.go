package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func doRequest(handler http.HandlerFunc, body interface{}) *httptest.ResponseRecorder {
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(b))
	w := httptest.NewRecorder()
	handler(w, req)
	return w
}

func TestAddHandler(t *testing.T) {
	w := doRequest(addHandler, CalcRequest{A: 2, B: 3})
	var resp CalcResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	if w.Code != http.StatusOK || resp.Result != 5 {
		t.Errorf("expected 200/5, got %d/%v", w.Code, resp.Result)
	}
}

func TestDivideHandlerByZero(t *testing.T) {
	w := doRequest(divideHandler, CalcRequest{A: 10, B: 0})
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

func TestDivideHandlerInvalidJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte("not json")))
	w := httptest.NewRecorder()
	divideHandler(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

func TestAddHandlerWrongMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	addHandler(w, req)
	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected 405, got %d", w.Code)
	}
}
