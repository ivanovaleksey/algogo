package rwlock

import "sync"

type RWLock struct {
	numReaders int
	rlock      sync.Locker

	wLock sync.Locker
}

func NewRWLock() *RWLock {
	return &RWLock{
		rlock: &sync.Mutex{},
		wLock: &sync.Mutex{},
	}
}

func (rw *RWLock) RLock() {
	rw.rlock.Lock()

	if rw.numReaders > 0 {
		rw.numReaders++
		rw.rlock.Unlock()
		return
	}

	rw.wLock.Lock()

	rw.numReaders++
	rw.rlock.Unlock()
}

func (rw *RWLock) RUnlock() {
	rw.rlock.Lock()

	if rw.numReaders > 1 {
		rw.numReaders--
		rw.rlock.Unlock()
		return
	}

	rw.wLock.Unlock()

	rw.numReaders--
	rw.rlock.Unlock()
}

func (rw *RWLock) Lock() {
	rw.wLock.Lock()
}

func (rw *RWLock) Unlock() {
	rw.wLock.Unlock()
}
