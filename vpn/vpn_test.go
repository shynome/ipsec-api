package vpn

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var x = sync.Mutex{}
var y = 0

func s() {
	x.Lock()
	defer x.Unlock()
	t := <-time.After(2 * time.Second)
	y++
	fmt.Printf("s %v %v\n", y, t)
}

func TestSync(t *testing.T) {
	i := 0
	for {
		go s()
		i++
		fmt.Printf("core %v \n", i)
		<-time.After(time.Second)
	}
}
