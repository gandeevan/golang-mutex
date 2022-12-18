package main

import (
    "testing"
)

func TestYieldLock(t *testing.T) {
    testMutexImplementation(t, NewYieldlock())
}