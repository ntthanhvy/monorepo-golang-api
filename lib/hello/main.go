package hello

import (
  "fmt"
)

// Greet greets the given name.
func Greet(audience string) string {
  return fmt.Sprintf("Hello, %s!", audience)
}