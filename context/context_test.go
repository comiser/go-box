package context

import (
	"testing"
)

func TestContext(t *testing.T) {
	c := New()
	c.Use(func(c Context) {
		println("1")
		c.Next()
		println("2")
	})
	c.Use(func(c Context) {
		defer func() {
			recover()
		}()
		c.Next()
	})
	c.Use(func() {
		println("3")
		panic("")
	})

	c.Provide(1)
	c.Next()
}
