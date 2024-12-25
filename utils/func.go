package utils

import (
	"github.com/cobra-base/cobra-go/glog"
	"runtime/debug"
)

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
