package infrastructure

import (
    "net/http"
    "time"
    "log"
    "io/ioutil"
    "strings"
    "encoding/json"
)


/*type httpData struct{

}*/


// default header for all http request
var default_header = map[string]string{
    "Content-type": "application/json",
}

// data for http request
type HttpReqData struct {
    BaseUrl string
    Path string
    Headers http.Header
    ParamUrl  map[string]string // parameters in URL example: local?test=1&other=2
    Method string
    DisableStrategy bool
}

// data response http request
type RespHttp struct {
    Body map[string]interface{}
    Headers  http.Header
    StatusCode int
}

// Call Http request get
func Do(data HttpReqData) (RespHttp, error) {

    if (data.Headers == nil) {
        data.Headers = http.Header{}
    }

    if (httpStrategy != HttpSecurity{} &&
        data.DisableStrategy == false) {
        httpStrategy.Strategy.Add(&data)
    }

    timeOut := time.Duration(5 * time.Second)

    var respHttp RespHttp // wrapper response http

    client := http.Client{
        Timeout: timeOut,
    }

    var url strings.Builder

    url.WriteString(data.BaseUrl)
    url.WriteString(data.Path)

    if data.Method == "" {
        data.Method = http.MethodGet
    }

    request, err := http.NewRequest(data.Method, url.String(), nil)

    for k, v := range default_header {
        request.Header.Set(k, v)
    }
    log.Println(data.Headers)
    request.Header = data.Headers


    /*if err != nil {
        return RespHttp{}, err
    }*/

    resp, err := client.Do(request)


    if err != nil {
        return RespHttp{}, err
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        return RespHttp{}, err
    }


    var result map[string]interface{}
    json.Unmarshal(body,&result)

    respHttp.Body = result

    respHttp.Headers = resp.Header
    respHttp.StatusCode = resp.StatusCode

    //resolver
    //fmt.Println(string(body))
    //var result map[string]interface{}
    //json.Unmarshal(body,&result)
    //hola := result["data"].([]interface{})
    //fmt.Println(hola[0].(map[string]interface{})["employee_salary"])


    // ver tipo de dato
    //xType := fmt.Sprintf("%T", hola[0])
    //fmt.Println(xType)

    return respHttp, nil
}
