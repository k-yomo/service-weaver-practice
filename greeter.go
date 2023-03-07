package main

import (
	"context"
	"fmt"
	"github.com/ServiceWeaver/weaver"
)

type greeterOptions struct {
	Greeting string
}

type Greeter interface {
	Greet(ctx context.Context, name string) (string, error)
}

type greeter struct {
	weaver.Implements[Greeter]
	weaver.WithConfig[greeterOptions]
}

func (g *greeter) Greet(ctx context.Context, name string) (string, error) {
	greeting := g.Config().Greeting
	if greeting == "" {
		greeting = "Hello"
	}
	return fmt.Sprintf("%s, %s!", greeting, name), nil
}
