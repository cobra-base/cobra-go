package utils

import (
    "github.com/cobra-base/cobra-go/glog"
    "math/rand"
    "runtime/debug"
)

func GenerateStrId(n int) string {

    letters := []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func SafeGo(callable func()) {
    go func() {
        defer func() {
            if err := recover(); err != nil {
                glog.Errorw("Panic", "err", err, "stack", string(debug.Stack()))
            }
        }()

        callable()
    }()
}
