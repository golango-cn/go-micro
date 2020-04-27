package broadcast

type Watcher interface {
	Next() (interface{}, error)
	Stop() error
}

type BroadCast interface {
	Send(interface{}) error
	Watch() Watcher
}

var broadCast BroadCast

func Init(b BroadCast) {
	broadCast = b
}

func GetBroadCast() BroadCast {
	return broadCast
}
