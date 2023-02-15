package monitor_key

import (
	"github.com/go-vgo/robotgo"
)

type KeyCallbackFun func(string)

var key_callback KeyCallbackFun

func SetKey(key string) {

	go func() {
		for true {
			ok := robotgo.AddEvents(key)
			if ok {
				if key_callback != nil {
					key_callback(key)
				}
			}
		}
	}()
}

func SetCallback(cb KeyCallbackFun) {
	key_callback = cb
}
