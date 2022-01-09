package errormap

import "sync"

var errorMap sync.Map

func Init(keys ...string) {
	for _, k := range keys {
		ch := make(chan error, 1024)
		errorMap.Store(k, ch)
	}
}

func List(key string) []error {
	ch := mapChan(key)
	if ch == nil {
		return nil
	}

	var list []error

out:
	for {
		select {
		case err := <-ch:
			list = append(list, err)
		default:
			break out
		}
	}

	return list
}

func mapChan(key string) chan error {
	value, ok := errorMap.Load(key)
	if !ok {
		return nil
	}

	return value.(chan error)
}

func Store(key string, err error) {
	ch := mapChan(key)
	if ch == nil {
		return
	}

	select {
	case ch <- err:
	default:
		break
	}
}
