package infrastructure

import (
    "net/http"
    "net/http/httptest"
    "testing"
    _"fmt"
)

func TestDoGet(t *testing.T) {
	srv := serverMock()
	defer srv.Close()

    _, err := Do(HttpReqData{BaseUrl: "fail",
        Path: "/bad-request", Method: http.MethodGet})

    if err == nil {
        t.Error("its bad url problen ", err)
    }

    resp, _ := Do(HttpReqData{BaseUrl: srv.URL, Path: "/test-request",
        Method: http.MethodGet})

    if resp.StatusCode != 200 {
        t.Error("Test response http get expected 200", resp.StatusCode)
    }
}


func serverMock() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/test-request", usersMock)
	srv := httptest.NewServer(handler)
	return srv
}

func usersMock(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    _, _ = w.Write([]byte(`{"test": 1}`))
}
