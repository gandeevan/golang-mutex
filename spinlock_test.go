package main

import (
    "github.com/stretchr/testify/require"
    "sync"
    "testing"
)

func TestSpinlockBasic(t *testing.T) {
    s := NewSpinlock()
    s.Lock()
    s.Unlock()
}

func TestMutualExclusion(t *testing.T) {
    var counter = 0
    var parallelism = 50
    var s = NewSpinlock()
    var wg sync.WaitGroup

    criticalSection := func() {
        s.Lock()
        counter++
        s.Unlock()
        wg.Done()
    }

    for i := 0; i<parallelism; i++ {
        wg.Add(1)
        go criticalSection()
    }


    wg.Wait()
    require.Equal(t, counter, parallelism)
}