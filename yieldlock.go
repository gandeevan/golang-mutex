package main

import (
    "github.com/sirupsen/logrus"
    "runtime"
    "sync/atomic"
)

// Limitations:
// Doesn't check if the lock is held by this thread before attempting to unlock.
// Doesn't ensure fairness (FIFO)
// Wasted context switches

type Yieldlock struct {
    state int64
}

func NewYieldlock() *Yieldlock {
    var lock Yieldlock
    lock.state = StateUnlocked
    return &lock
}

func (m *Yieldlock) Lock() {
    for {
        if !atomic.CompareAndSwapInt64(&m.state, StateUnlocked, StateLocked) {
            runtime.Gosched()
        }
        break
    }
    logrus.Info("Mutex locked")
}

func (m *Yieldlock) Unlock() {
    atomic.StoreInt64(&m.state, StateUnlocked)
    logrus.Infof("Mutex unlocked")
}