package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ServiceWeaver/weaver"
)

func main() {
	// Get a network listener on address "localhost:12345".
	root := weaver.Init(context.Background())

	greeter, err := weaver.Get[Greeter](root)
	if err != nil {
		log.Fatal(err)
	}

	reverser, err := weaver.Get[Reverser](root)
	if err != nil {
		log.Fatal(err)
	}

	opts := weaver.ListenerOptions{LocalAddress: "localhost:12345"}
	lis, err := root.Listener("hello", opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("hello listener available on %v\n", lis)

	// Serve the /hello endpoint.
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		name := r.URL.Query().Get("name")
		reverse := r.URL.Query().Get("reverse")

		if reverse == "true" {
			var err error
			name, err = reverser.Reverse(ctx, name)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		greeting, err := greeter.Greet(ctx, name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, greeting)
	})
	http.Serve(lis, nil)
}
