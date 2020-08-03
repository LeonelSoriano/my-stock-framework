package infrastructure

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/LeonelSoriano/my-stock-framework/src/base/env"
    "strings"
)

func TestDo(t *testing.T) {

    env.IntiEnv(env.IntiEnvParams{EnvPath: "../base/env/"})

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

    SetHttpStrategy(HttpSecurity {
        Strategy: AlpacaStrategy{},
    })



    resp, _ = Do(HttpReqData{BaseUrl: srv.URL, Path: "/test-strategy",
        Method: http.MethodGet})

    if resp.StatusCode != 200 {
        t.Error("Test response http get expected 200 straygy fail", resp.StatusCode)
    }
}


func serverMock() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/test-request", usersMock)
	handler.HandleFunc("/test-strategy", strategyMock)

	srv := httptest.NewServer(handler)
	return srv
}

func usersMock(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    _, _ = w.Write([]byte(`{"test": 1}`))
}

func strategyMock(w http.ResponseWriter, r *http.Request) {
    if strings.Compare(r.Header.Get("APCA-API-KEY-ID"), "test-id") == 0 &&
        strings.Compare(r.Header.Get("APCA-API-SECRET-KEY"), "test-secret") ==
            0 {
        w.WriteHeader(http.StatusOK)
    } else {
        w.WriteHeader(http.StatusInternalServerError)
    }
    w.Header().Set("Content-Type", "application/json")
    _, _ = w.Write([]byte(`{"strategy": 1}`))
}
