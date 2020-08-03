package infrastructure

import (
    "github.com/LeonelSoriano/my-stock-framework/src/base/env"
)

var httpStrategy HttpSecurity

type HttpSecurityStrategy interface {
    Add(data *HttpReqData)
}

type HttpSecurity struct {
    GlobalUrlBackend string // url for server request token etc
    Strategy HttpSecurityStrategy
}

type AlpacaStrategy struct{}

func (a AlpacaStrategy) Add(data *HttpReqData) {
    data.Headers.Add("APCA-API-KEY-ID", env.GetEnv("alpacaApiKeyId"))
    data.Headers.Set("APCA-API-SECRET-KEY", env.GetEnv("alpacaApySecret"))
}

func SetHttpStrategy(s HttpSecurity) {
    httpStrategy = s
}
