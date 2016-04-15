package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/hnakamur/delayedticker"
)

func main() {
	log.Printf("start")
	ticker := delayedticker.NewDelayedTicker(time.Duration(5)*time.Second, time.Duration(2)*time.Second)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	for {
		select {
		case t := <-ticker.C:
			log.Printf("received from ticker: t=%v", t)
		case <-c:
			log.Print("received signal")
			ticker.Stop()
			time.Sleep(time.Second)
			return
		}
	}
}
