package utils

import (
    "crypto/tls"
    "errors"
    "fmt"
    "github.com/cobra-base/cobra-go/glog"
    "io"
    "net/http"
    "net/url"
    "os"
    "time"
)

const httpReqTimeout = time.Second * 5

func HttpGet(webUrl string) ([]byte, error) {
    req, err := http.NewRequest("GET", webUrl, nil)
    if err != nil {
        return nil, err
    }

    tr := &http.Transport{}
    tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

    proxyUrl := os.Getenv("HTTPS_PROXY")
    if len(proxyUrl) > 0 {
        proxy, err := url.Parse("http://127.0.0.1:7897")
        if err != nil {
            return nil, err
        }
        tr.Proxy = http.ProxyURL(proxy)
    }

    client := &http.Client{
        Transport: tr,
        Timeout:   httpReqTimeout,
    }
    stopWatch := NewStopWatch()

    rsp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer rsp.Body.Close()

    stopWatch.Stop()
    spendTime := stopWatch.Duration()

    statusCode := rsp.StatusCode
    statusMessage := rsp.Status

    glog.Debugw("http get", "statusCode", statusCode, "statusMessage", statusMessage,
        "spendTime(Ms)", spendTime, "webUrl", webUrl)

    if statusCode != 200 {
        return nil, errors.New(fmt.Sprintf("%s,%s", statusMessage, webUrl))
    }

    body, err := io.ReadAll(rsp.Body)
    if err != nil {
        return nil, err
    }

    return body, err
}
