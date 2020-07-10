// Set values to ENVIRONMENT var from  .env file.

package env

import (
    "bufio"
    "os"
    "path/filepath"
    "strings"
)

// Default name form enviroment vars
var envPath string = ".env"

// param to init Env vars
type IntiEnvParams struct {
    EnvPath string // path to .env file
}

// init and add var to Enviroment
//
// circumstances:
//    - if line is "" return ("", "")
//    - if line is "#" is comment then return ("", "")
//    - if line is only key return ("", "")
//    - error when (not found env file)
func IntiEnv(params IntiEnvParams) error {

    if params.EnvPath != "" {
        envPath = params.EnvPath
    }
    absPath, _ := filepath.Abs(envPath)
    file, err := os.Open(absPath)

   if err != nil {
        return err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        key, value := parceEnv(scanner.Text())
        if key != "" {
            os.Setenv(key, value)
        }
    }
    return nil
}

// parce line to line env file
func parceEnv(line string) (string, string) {
    if line == "" { // is void line
        return "", ""
    }

    if line[0] == '#' { // is comment
        return "", ""
    }

    split := strings.SplitN(line, "=", 2)

    if len(split) != 2 { // have only key
        return "", ""
    }
    return split[0],split[1] // is ok
}



func setEnv(key string, val string) {
    os.Setenv(key, val)
}

func GetEnv(key string) string {
    return os.Getenv(key)
}
