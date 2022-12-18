package main

import (
    "testing"
)

func TestTicketLock(t *testing.T) {
    testMutexImplementation(t, NewTicketlock())
}