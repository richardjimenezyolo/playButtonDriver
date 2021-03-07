package main

import (
	"fmt"
	"time"

	"github.com/micmonay/keybd_event"
	hook "github.com/robotn/gohook"
)

func main() {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic("Error when initializing keyboard")
	}
	hook.Register(hook.KeyDown, []string{"ctr", "shift", "space"}, func(e hook.Event) {
		fmt.Println(e)
		kb.SetKeys(keybd_event.VK_MEDIA_PLAY_PAUSE)
		kb.Press()
		time.Sleep(1 * time.Second)
		kb.Release()
	})
	s := hook.Start()
	defer hook.End()

	<-hook.Process(s)

}
