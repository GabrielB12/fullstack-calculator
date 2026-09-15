package main

import (
	"encoding/json"
	"net/http"
)

type CalcRequest struct {
	A float64 `json:"a"`
	B float64 `json:"b,omitempty"`
}

type CalcResponse struct {
	Result float64 `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func writeJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, ErrorResponse{Error: message})
}

func decodeRequest(w http.ResponseWriter, r *http.Request) (CalcRequest, bool) {
	var req CalcRequest
	if r.Body == nil {
		writeError(w, http.StatusBadRequest, "request body is required")
		return req, false
	}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body: "+err.Error())
		return req, false
	}
	return req, true
}

func methodGuard(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "only POST is allowed")
		return false
	}
	return true
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	if !methodGuard(w, r) {
		return
	}
	req, ok := decodeRequest(w, r)
	if !ok {
		return
	}
	writeJSON(w, http.StatusOK, CalcResponse{Result: Add(req.A, req.B)})
}

func subtractHandler(w http.ResponseWriter, r *http.Request) {
	if !methodGuard(w, r) {
		return
	}
	req, ok := decodeRequest(w, r)
	if !ok {
		return
	}
	writeJSON(w, http.StatusOK, CalcResponse{Result: Subtract(req.A, req.B)})
}

func multiplyHandler(w http.ResponseWriter, r *http.Request) {
	if !methodGuard(w, r) {
		return
	}
	req, ok := decodeRequest(w, r)
	if !ok {
		return
	}
	writeJSON(w, http.StatusOK, CalcResponse{Result: Multiply(req.A, req.B)})
}

func divideHandler(w http.ResponseWriter, r *http.Request) {
	if !methodGuard(w, r) {
		return
	}
	req, ok := decodeRequest(w, r)
	if !ok {
		return
	}
	result, err := Divide(req.A, req.B)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, CalcResponse{Result: result})
}

func powerHandler(w http.ResponseWriter, r *http.Request) {
	if !methodGuard(w, r) {
		return
	}
	req, ok := decodeRequest(w, r)
	if !ok {
		return
	}
	writeJSON(w, http.StatusOK, CalcResponse{Result: Power(req.A, req.B)})
}

func sqrtHandler(w http.ResponseWriter, r *http.Request) {
	if !methodGuard(w, r) {
		return
	}
	req, ok := decodeRequest(w, r)
	if !ok {
		return
	}
	result, err := SquareRoot(req.A)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, CalcResponse{Result: result})
}

func percentageHandler(w http.ResponseWriter, r *http.Request) {
	if !methodGuard(w, r) {
		return
	}
	req, ok := decodeRequest(w, r)
	if !ok {
		return
	}
	result, err := Percentage(req.A, req.B)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, CalcResponse{Result: result})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}