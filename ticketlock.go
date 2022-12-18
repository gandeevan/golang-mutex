package main

import (
    "github.com/sirupsen/logrus"
    "runtime"
    "sync/atomic"
)

// Limitations:
// Doesn't check if the lock is held by this thread before attempting to unlock.
// Wasted context switches

type Ticketlock struct {
    currTicket int64
    nextTicket int64
}

func NewTicketlock() *Ticketlock {
    var lock Ticketlock
    lock.currTicket = 0
    lock.nextTicket = -1
    return &lock
}

func (m *Ticketlock) Lock() {
    myTicket := atomic.AddInt64(&m.nextTicket, 1)
    for {
        if myTicket != m.currTicket {
            runtime.Gosched()
        } else {
            break
        }
    }
    logrus.Info("Mutex locked")
}

func (m *Ticketlock) Unlock() {
    atomic.AddInt64(&m.currTicket, 1)
    logrus.Infof("Mutex unlocked")
}