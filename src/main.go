package main

import (
    _"fmt"
    infra "github.com/LeonelSoriano/my-stock-framework/src/infrastructure"
    "github.com/LeonelSoriano/my-stock-framework/src/base/env"
    "net/http"
)

func main() {

    env.IntiEnv(env.IntiEnvParams{})

    //t := HttpSecurity{GlobalUrlBackend: "algo",Strategy: AlpacaStrategy{} };


    infra.SetHttpStrategy(infra.HttpSecurity{
        Strategy: infra.AlpacaStrategy{},
        GlobalUrlBackend: env.GetEnv("alpacaEndPoint"),
    })

    infra.Do(infra.HttpReqData{
        BaseUrl: "http://dummy.restapiexample.com",
        Path: "/api/v1/employees",
        Method: http.MethodGet,
    })


    //fmt.Println(err)
    //fmt.Println(resp.Body)
}
