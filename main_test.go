package main
import "testing"
func TestAdd(t *testing.T) {
  result := Add(2, 3)
  expect := 5
  if result != expected {
    t.Errorf("Add(2, 3) = %d; want %d", result, expected)
      }
  }
