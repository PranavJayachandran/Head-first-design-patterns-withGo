package main

import (
	"fmt"
	"time"
)

type subject interface {
	subscribe(s chan string)
	unsubscribe(s chan string)
	notify(message string)
}
type dataSource struct {
	subs []chan<- string
}

func (d *dataSource) subscribe(s chan string) {

	d.subs = append(d.subs, s)
}
func (d *dataSource) unsubscribe(s chan string) {
	for i, sub := range d.subs {
		if sub == s {
			d.subs = append(d.subs[:i], d.subs[i+1:]...)
		}
	}
}
func (d *dataSource) notify(message string) {
	for _, subs := range d.subs {
		subs <- message
	}
}

func main() {
	var d dataSource = dataSource{}
	go func() {
		for {
			time.Sleep(1 * time.Second)
			d.notify(time.Now().String())
		}
	}()
	go func() {
		ch := make(chan string)
		d.subscribe(ch)
		for {
			fmt.Println("First", <-ch)
		}
	}()
	go func() {
		ch := make(chan string)
		d.subscribe(ch)
		for {
			fmt.Println(<-ch)
		}
	}()
	fmt.Scanln()
}
