package twitter

import (
    "fmt"
    "net/http"
)

func GetFinalUrl(shortUrl string) (string, error) {
    client := &http.Client{
        CheckRedirect: func(req *http.Request, via []*http.Request) error {
            return http.ErrUseLastResponse
        },
    }

    rsp, err := client.Get(shortUrl)
    if err != nil {
        return "", err
    }

    defer rsp.Body.Close()

    finalUrl := rsp.Header.Get("Location")
    if len(finalUrl) == 0 {
        return "", fmt.Errorf("not found location header,short url %s", shortUrl)
    }

    return finalUrl, nil
}
