package env

import(
    "strings"
    "testing"
)

func TestParceEnv(t *testing.T) {
    //t.Log(key)
    key , value := parceEnv("")
    if key != "" || value != ""  {
       t.Errorf("if line is '' return ''  ''")
    }

    key , value = parceEnv("test")
    if key != "" || value != "" {
       t.Errorf("if the line is only a word return ''")
    }

    key , value = parceEnv("#test comment")
    if key != "" || value != "" {
       t.Errorf("if the line init with # return ''")
    }

    key , value = parceEnv("VAR=TEST")
    if key != "VAR" || value != "TEST" {
       t.Errorf("key = value")
    }
}

func TestIntiEnv(t *testing.T) {
    err := IntiEnv(IntiEnvParams{})

    if err != nil {
       t.Errorf("An Problem init IntiEnv")
    }

    var valueFromEnv = GetEnv("TEST")
    if  strings.Compare(valueFromEnv, "UNIT") != 0 {
        t.Errorf("Error: i cant get value from test .env")
    }

    err = IntiEnv(IntiEnvParams{EnvPath: "fail"})

    if err == nil {
        t.Errorf("IntiEnv with wrong params not fail")
    }

}

func TestGetEnv(t *testing.T) {
    setEnv("TestGetEnv", "SAME_VALUE")
    testVal := GetEnv("TestGetEnv")
    if strings.Compare(testVal, "SAME_VALUE") != 0 {
        t.Errorf(`Error: not setting value o not get value error from
            ENVIRONMENT VAR`)
    }
}
