package database

import "errors"

type Watcher struct {
	exit    chan interface{}
	updated chan interface{}
}

func (w Watcher) Next() (interface{}, error) {
	select {
	case <-w.exit:
		return nil, errors.New("Watch stoped")
	case v := <-w.updated:
		return v, nil
	}
}

func (w Watcher) Stop() error {
	select {
	case <-w.exit:
	default:
		close(w.exit)
	}
	return nil
}
