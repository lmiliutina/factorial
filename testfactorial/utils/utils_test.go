package utils

import "testing"

func TestFactorial(t *testing.T) {
    cases := []struct {
        in, want int
    }{
        {0, 1},
        {1, 1},
        {5, 120},
        {6, 720},
    }
    for _, c := range cases {
        got := Factorial(c.in)
        if got != c.want {
            t.Errorf("Factorial(%d) == %d, want %d", c.in, got, c.want)
        }
    }
}
