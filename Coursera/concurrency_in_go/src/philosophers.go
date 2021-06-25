package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	id      int
	left    *Chopstick
	right   *Chopstick
	channel chan bool
}

type Chopstick struct {
	sync.Mutex
}

func CreateChopsticks() []*Chopstick {
	chopsticks := make([]*Chopstick, 0)

	for i := 0; i < 5; i++ {
		chopsticks = append(chopsticks, new(Chopstick))
	}

	return chopsticks
}

func CreatePhilosohpers(chopsticks []*Chopstick) []Philosopher {
	philosophers := make([]Philosopher, 0)
	for i := 0; i < 5; i++ {
		philosophers = append(philosophers, Philosopher{id: i + 1, left: chopsticks[i], right: chopsticks[(i+1)%5], channel: make(chan bool)})
	}

	return philosophers
}

func Host(channel chan Philosopher) {
	for philosopher := range channel {
		philosopher.left.Lock()
		philosopher.right.Lock()
		philosopher.channel <- true
	}
}

func (philosopher Philosopher) Eat(channel chan Philosopher, wait_group *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		channel <- philosopher
		<-philosopher.channel

		fmt.Println("starting to eat", philosopher.id)
		time.Sleep(time.Millisecond)
		fmt.Println("finishing eating", philosopher.id)

		philosopher.left.Unlock()
		philosopher.right.Unlock()
	}

	wait_group.Done()
}

func main() {
	var wait_group sync.WaitGroup
	channel := make(chan Philosopher, 2)

	chopsticks := CreateChopsticks()
	philosophers := CreatePhilosohpers(chopsticks)

	go Host(channel)

	for _, philosopher := range philosophers {
		go philosopher.Eat(channel, &wait_group)
		wait_group.Add(1)
	}

	wait_group.Wait()
}
