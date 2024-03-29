package main

import (
	"context"
	"fmt"

	"github.com/DFXLuna/apiserver/internal"
)

func main() {
	errch := make(chan error)
	go internal.Run(context.Background(), errch)

	err := <-errch
	fmt.Printf("server shutdown: %v", err)
}
