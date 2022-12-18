package main

import (
    "testing"
)

func TestSpinlock(t *testing.T) {
    testMutexImplementation(t, NewSpinlock())
}