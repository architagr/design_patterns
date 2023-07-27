package singletonVideoExample

import "sync"

type ICounter interface {
	Increment() int
	Decrement() int
	GetValue() int
}

type counter struct {
	value int
}

var counterInstance ICounter
var once sync.Once

func (cntr *counter) Increment() int {
	cntr.value++
	return cntr.value
}

func (cntr *counter) Decrement() int {
	cntr.value--
	return cntr.value
}
func (cntr *counter) GetValue() int {
	return cntr.value
}
func CreateCounter() ICounter {

	once.Do(func() {
		counterInstance = &counter{
			value: 0,
		}
	})
	return counterInstance
}
