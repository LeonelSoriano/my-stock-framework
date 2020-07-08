package main

import (
    "fmt"
    _"github.com/LeonelSoriano/my-stock-framework/src/infrastructure"
    "github.com/LeonelSoriano/my-stock-framework/src/base/env"
)

func main() {

    env.IntiEnv(env.IntiEnvParams{})
    //infrastructure.DoGet()
    fmt.Println(env.GetEnv("urlDatabase"))
}
