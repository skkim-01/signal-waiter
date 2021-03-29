package SignalWaiter

import (
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"time"
)

type waitObject struct {
	SigResult       os.Signal
	SigEventChannel chan os.Signal
	counter         uint64
}

var instance *waitObject
var once sync.Once

func getInstance() *waitObject {
	once.Do(func() {
		instance = &waitObject{}
	})
	return instance
}

// GetSignal: which signal is caught....
func GetSignal() os.Signal {
	return getInstance().SigResult
}

// SendSignal: other thread can signal
func SendSignal(sig os.Signal) {
	getInstance().SigEventChannel <- sig
}

func Close() {
	i := getInstance()
	i.SigResult = nil
	close(i.SigEventChannel)
}

// Wait: wait signal every 1 second
func Wait(argc ...os.Signal) {
	getInstance().SigEventChannel = make(chan os.Signal, 1)
	// atomic int = 1
	atomic.AddUint64(&getInstance().counter, 1)
	go getInstance().sigHandleThread(argc...)

_START_SECTION:
	select {
	case <-time.After(time.Second):
		if 0 == atomic.LoadUint64(&getInstance().counter) {
			goto _EXIT
		}
		goto _START_SECTION
	}
_EXIT:
}

// sigHandleThread: signal wait thread
func (w *waitObject) sigHandleThread(slWaitSignals ...os.Signal) {
	signal.Notify(w.SigEventChannel, slWaitSignals...)
	sig := <-w.SigEventChannel
	w.SigResult = sig
	atomic.StoreUint64(&getInstance().counter, 0)
}
