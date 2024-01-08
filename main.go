package main

import (
	"log"

	server "github.com/AdityaMayukhSom/ruskin/server"
	store "github.com/AdityaMayukhSom/ruskin/store"
)

func main() {
	serverconf := &server.ServerConfig{
		ListenAddr:   ":3000",
		StoreFactory: store.NewMemoryStoreFactory(nil),
	}

	server, err := server.NewServer(serverconf)
	if err != nil {
		log.Fatal(err)
	}

	err = server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
