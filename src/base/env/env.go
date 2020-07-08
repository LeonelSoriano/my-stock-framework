package env

import (
    "bufio"
    _"fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
)

var envPath string = ".env"

func setEnv(key string, val string) {
    os.Setenv(key, val)
}

func GetEnv(key string) string {
    return os.Getenv(key)
}

type IntiEnvParams struct {
    EnvPath string
}
func IntiEnv(params IntiEnvParams) {
    //fmt.Println(os.Getenv("PATH"))

    if params.EnvPath != "" {
        envPath = params.EnvPath
    }

    absPath, _ := filepath.Abs(envPath)
    file, err := os.Open(absPath)

   if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        key, value := parceEnv(scanner.Text())
        if key != "" {
            os.Setenv(key, value)
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func parceEnv(line string) (string, string) {
    if line == "" {
        return "", ""
    }

    if line[0] == '#' {
        return "", ""
    }

    split := strings.SplitN(line, "=", 2)

    if len(split) != 2 {
        return "", ""
    }

    return split[0],split[1]
}
