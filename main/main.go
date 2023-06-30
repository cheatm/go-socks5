package main

import (
	"fmt"
	"os"
	"socks5"
)

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		value = defaultValue
	}
	return value
}

func main() {
	addr := getEnv("BIND_ADDR", "0.0.0.0:1080")
	network := getEnv("NETWORK", "tcp")

	// Create a SOCKS5 server
	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("server listen %s://%s", network, addr)
	// Create SOCKS5 proxy on localhost port 8000
	if err := server.ListenAndServe(network, addr); err != nil {
		panic(err)
	}
}
