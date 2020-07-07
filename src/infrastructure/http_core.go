package infrastructure

import (
    "net/http"
    "time"
    "log"
    "io/ioutil"
)

func Sum(x int, y int) int {
    return x + y
}

/*type httpData struct{

}*/

type HttpReqData struct {
    BaseUrl string
    path string
    Headers  map[string]string
    ParamUrl  map[string]string
    body string
}



func DoGet() {
    timeOut := time.Duration(5 * time.Second)

    client := http.Client{
        Timeout: timeOut,
    }

    url := "http://dummy.restapiexample.com/api/v1/employees"

    request, err := http.NewRequest("GET", url, nil)
    request.Header.Set("Content-type", "application/json")

    if err != nil {
        log.Fatalln(err)
    }

    resp, err := client.Do(request)

    if err != nil {
        log.Fatalln(err)
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        log.Fatalln(err)
    }

    log.Println(string(body))
    //log.Println(resp.Status)
}
