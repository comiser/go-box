package context

import (
	"github.com/codegangsta/inject"
)

type Context interface {
	Next()
	Use(...Middleware) Context
	Provide(...interface{}) Context
}

type context struct {
	injector inject.Injector
	lists    []Middleware
	index    int
}

func New() Context {
	c := &context{}
	c.injector = inject.New()
	return c.Provide(c)
}

func (c *context) Next() {
	c.index++
	if c.index <= len(c.lists) {
		mw := c.lists[c.index-1]
		_, err := c.injector.Invoke(mw)
		if err != nil {
			// TODO
			panic(err)
		}
	}
}

func (c *context) Use(mw ...Middleware) Context {
	c.lists = append(c.lists, mw...)
	return c
}

func (c *context) Provide(values ...interface{}) Context {
	for _, value := range values {
		c.injector.Map(value)
	}
	return c
}
