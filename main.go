package main

import "github.com/mjrdev/api-go/cmd/api"

func main() {
	s := api.Start()
	api.Run(s)
}
