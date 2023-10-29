package main

import ()

func main() {

	server := NewAPIServer(":3000")
	server.Run()
}
