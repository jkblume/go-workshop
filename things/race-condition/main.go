package main

import "time"


type T struct {
	sharedInt   int
	unusedValue int
}

func (t T) runReader() {
	for {
		var val int = t.sharedInt
		if val%10 == 0 {
			t.unusedValue = t.unusedValue + 1
		}
	}
}

func (t T) runWriter() {
	for {
		t.sharedInt = t.sharedInt + 1
	}
}

func main() {
	t := T{}
	go t.runReader()
	go t.runWriter()
	time.Sleep(1 * time.Second)
}
