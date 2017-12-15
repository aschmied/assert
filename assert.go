package assert

import "testing"

func Int(t *testing.T, actual, expected int) {
    if actual != expected {
        t.Errorf("Expected %d but got %d", expected, actual)
    }
}

func String(t *testing.T, actual, expected string) {
    if actual != expected {
        t.Errorf("Expected %s but got %s", expected, actual)
    }
}
