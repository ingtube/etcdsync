package main

import (
	"log"
	"os"
	"time"

	"github.com/ingtube/etcdsync"
	"go.etcd.io/etcd/client"
)

func main() {
	cfg := client.Config{
		Endpoints:               []string{"http://127.0.0.1:2379"},
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		return
	}
	m, err := etcdsync.New("/mylock", 10, c)
	if m == nil || err != nil {
		log.Printf("etcdsync.New failed")
		return
	}
	m.SetDebugLogger(os.Stdout)
	err = m.Lock()
	if err != nil {
		log.Printf("etcdsync.Lock failed")
	} else {
		log.Printf("etcdsync.Lock OK")
	}

	log.Printf("Get the lock. Do something here.")

	err = m.Unlock()
	if err != nil {
		log.Printf("etcdsync.Unlock failed")
	} else {
		log.Printf("etcdsync.Unlock OK")
	}
}
