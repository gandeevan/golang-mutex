package main

const (
    StateLocked   = 1
    StateUnlocked = 0
)

type Mutex interface {
    Lock()
    Unlock()
}
