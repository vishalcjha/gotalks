package main

import "fmt"

type intQueue interface {
	get() int
	put(x int)
}

type basicIntQueue struct {
	queue []int
}

func (b *basicIntQueue) get() int {
	firstElem := b.queue[0]
	b.queue = b.queue[1:]
	return firstElem
}

func (b *basicIntQueue) put(p int) {
	b.queue = append(b.queue, p)
}

type doublingIntQueue struct {
	queue intQueue
}

func (d *doublingIntQueue) get() int {
	return d.queue.get()
}

func (d *doublingIntQueue) put(p int) {
	d.queue.put(2 * p)
}

type positiveIntQueue struct {
	queue intQueue
}

func (p *positiveIntQueue) get() int {
	return p.queue.get()
}

func (p *positiveIntQueue) put(x int) {
	if x > 0 {
		p.queue.put(x)
	}
}

func main() {
	q1 := basicIntQueue{}
	q1.put(10)
	q1.put(20)
	q1.put(-10)

	p2 := positiveIntQueue{
		queue: &doublingIntQueue{
			queue: &basicIntQueue{},
		},
	}

	p2.put(-10)
	p2.put(10)
	p2.put(20)

	fmt.Println(p2.get())
	fmt.Println(p2.get())
}
