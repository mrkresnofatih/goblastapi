package main

import (
	"goblastapi-example/controllers"
	"sync"
)

func main() {
	var runState sync.WaitGroup
	controllers.InitServer(&runState)
	runState.Wait()
}
