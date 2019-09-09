package main

import (
	"math/rand"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type Hello struct{ Who string }
type Customer struct {
	behavior actor.Behavior
}

type Signup struct {
}
type DeleteAccount struct {
}

func (state *Customer) Receive(context actor.Context) {
	state.behavior.Receive(context)
}

func (state *Customer) Nonsubscriber(context actor.Context) {
	switch context.Message().(type) {

	case PostLaunchCallToAction:
		if rand.Intn(100) > 10 {
			context.Respond(Signup{})
			state.behavior.Become(state.Subscriber)
		}
	case PrelaunchCallToAction:
		if rand.Intn(100) > 5 {
			context.Respond(Signup{})
			state.behavior.Become(state.Subscriber)
		}
	}
}

func (state *Customer) Subscriber(context actor.Context) {
	switch msg := context.Message().(type) {
	case CurrentTime:
		if msg.Time.Weekday() == time.Friday {

		}

	}
}

func NewCustomer() actor.Actor {
	act := &Customer{
		behavior: actor.NewBehavior(),
	}
	act.behavior.Become(act.Nonsubscriber)
	return act
}
