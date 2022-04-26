package main

import (
	"log"
	"sync"
)

type Singleton struct {

}

var (
	instance *Singleton
	mu sync.Mutex
)

func Instance() *Singleton {
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		instance = &Singleton{}
	}
	return instance
}

func main() {
	log.Println("singleton");


}
