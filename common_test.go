package main

import (
    "github.com/stretchr/testify/require"
    "sync"
    "testing"
)
func testMutexImplementation(t *testing.T, s Mutex) {
    var counter = 0
    var parallelism = 50
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
