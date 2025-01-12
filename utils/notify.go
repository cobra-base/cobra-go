package utils

import (
    "fmt"
    "github.com/cobra-base/cobra-go/glog"
    "net/http"
    "sync"
    "time"
)

func SmsNotify(proj string, content string) {
    SafeGo(func() {
        url := fmt.Sprintf("http://110.40.228.192:51001/smsNotify?proj=%s&content=%s", proj, content)
        _, err := http.Get(url)
        if err != nil {
            glog.Errorw("sms notify except", "proj", proj, "content", content)
        }
    })
}

var smsBooks sync.Map

func SmsLimit(proj string, content string, interval int64) {
    if interval > 0 {
        var last int64
        k := fmt.Sprintf("%s_%s", proj, content)
        v, ok := smsBooks.Load(k)
        if ok {
            last = v.(int64)
        }
        now := time.Now().Unix()
        if last+interval > now {
            return
        }
        smsBooks.Store(k, now)
    }

    SmsNotify(proj, content)
}
