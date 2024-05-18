package mygopkg

import "testing"

func TestAdd(t *testing.T) {
    if Add(2, 3) != 5 {
        t.Error("Expected 2 + 3 to equal 5")
    }
}

func TestMultiply(t *testing.T) {
    if Multiply(2, 3) != 6 {
        t.Error("Expected 2 * 3 to equal 6")
    }
}
