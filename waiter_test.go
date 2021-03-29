package SignalWaiter

import (
	"fmt"
	"syscall"
	"testing"
	"time"
)

func Test_Wait(t *testing.T) {
	fmt.Println("### LOG ### Main: Start testing")

	go send_signal_thread()

	// SignalWaiter.Wait()
	Wait()

	fmt.Println("### LOG ### Main: Caught singal:", GetSignal())

	// SignalWaiter.Close()
	Close()
}

func send_signal_thread() {
	// in 5 seconds
	fmt.Println("### LOG ### Thread: send sighup in 5 seconds")
	time.Sleep(5 * time.Second)
	// SignalWaiter.Close()
	SendSignal(syscall.SIGHUP)
}
