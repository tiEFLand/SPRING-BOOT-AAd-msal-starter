
package utils

import (
	"errors"
	"time"
)

type Lock struct {
	chLock chan struct{}
}

func (l *Lock) Lock() {
	l.chLock <- struct{}{}
}

func (l *Lock) Unlock() {
	select {
	case <-l.chLock:
		{
		}
	default:
		{
			panic("double Unlock")
		}
	}
}

func (l *Lock) TryLock(timeout time.Duration) error {
	t := time.After(timeout)
	select {
	case l.chLock <- struct{}{}:
		return nil
	case <-t:
		return errors.New("timeout")
	}
}

func GenLock() Lock {
	return Lock{
		chLock: make(chan struct{}, 1),
	}
}