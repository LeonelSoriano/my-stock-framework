// Set values to ENVIRONMENT var from  .env file.

package env

import (
    "bufio"
    "os"
    "path/filepath"
    "strings"
    "errors"
)

// Default name form environment vars
var envPath string = ".env"

// parameter to init env vars
type IntiEnvParams struct {
    EnvPath string // path to .env file
}

// init and add var to Environment
//
// circumstances:
//    - if line is "" return ("", "")
//    - if line is "#" is comment then return ("", "")
//    - if line is only key return ("", "")
//    - error when (not found env file)
func IntiEnv(params IntiEnvParams) error {

    if params.EnvPath != "" {
        envPath = params.EnvPath + ".env"
    }

    fileInfo, err := os.Stat(envPath)

    if err != nil {
      return err
    }

    if fileInfo.IsDir() {
        return errors.New(".env cant be folder")
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
    return split[0],split[1] // is OK
}



func setEnv(key string, val string) {
    os.Setenv(key, val)
}

func GetEnv(key string) string {
    return os.Getenv(key)
}
