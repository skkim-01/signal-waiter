# signal-waiter

### Package Name : SignalWaiter

#### abstract
*os signal waiter using by singleton struct*

#### type definitions
##### waitObject struct
```go
  type waitObject struct {
    SigResult       os.Signal
    SigEventChannel chan os.Signal
    counter         uint64
  }
```  
  
#### APIs
##### Wait(): wait with signals, if null, wait all signals
##### https://golang.org/pkg/os/signal/#Notify
```go
  func Wait(argc ...os.Signal)
```

##### GetSignal() : which signal is caught
```go
  func GetSignal() os.Signal {
```

##### SendSignal : other thread can send signal internally
```go
  func SendSignal(sig os.Signal)
```
