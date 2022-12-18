package main

import (
    "github.com/sirupsen/logrus"
    "sync/atomic"
)

// Limitations:
// Doesn't check if the lock is held by this thread before attempting to unlock.
// Doesn't ensure fairness (FIFO)
// Busy looping wastes CPU cycles


type Spinlock struct {
    state int64
}

func NewSpinlock() *Spinlock {
    var lock Spinlock
    lock.state = StateUnlocked
    return &lock
}

func (m *Spinlock) Lock() {
    for {
        if atomic.CompareAndSwapInt64(&m.state, StateUnlocked, StateLocked) {
            break
        }
    }
    logrus.Info("Mutex locked by routine ")
}

func (m *Spinlock) Unlock() {
    atomic.StoreInt64(&m.state, StateUnlocked)
    logrus.Infof("Mutex unlocked by routine")
}