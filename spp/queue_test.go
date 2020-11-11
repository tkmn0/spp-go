package spp

import (
	"math/rand"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	arr := rand.Perm(3000)

	q := NewQueue()
	dequeued := make(chan bool)

	r := make([]int, 0)

	go func() {
	loop:
		for {
			n := q.Dequeue()
			if n != nil {
				r = append(r, n.(int))
			}
			if len(arr) == len(r) {
				dequeued <- true
				break loop
			}
		}
	}()

	for _, n := range arr {
		q.Enqueue(n)
	}

	<-dequeued

	for i, n := range arr {
		if n != r[i] {
			t.Errorf("queued value miss matched %v, %v", n, r[i])
		}
	}
}
