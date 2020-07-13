package main

import (
    "fmt"
    "github.com/LeonelSoriano/my-stock-framework/src/infrastructure"
    "github.com/LeonelSoriano/my-stock-framework/src/base/env"
)

func main() {

    env.IntiEnv(env.IntiEnvParams{})
    resp,err := infrastructure.DoGet(infrastructure.HttpReqData{
        BaseUrl: "http://dummy.restapiexample.com",
        Path: "/api/v1/employees",
    })


    fmt.Println(err)
    fmt.Println(resp.Body)
}
