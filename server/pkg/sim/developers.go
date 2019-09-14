package main

import (
	"fmt"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type DeveloperTeam struct {
	start            time.Time
	behavior         actor.Behavior
	marketingTeamPID *actor.PID
	timeToBuild      time.Duration
}
type ProductionLaunch struct {
}

func (state *DeveloperTeam) Receive(context actor.Context) {
	state.behavior.Receive(context)
}
func (state *DeveloperTeam) Finished(context actor.Context) {

}
func (state *DeveloperTeam) Building(context actor.Context) {
	switch msg := context.Message().(type) {
	case CurrentTime:
		if msg.Sub(state.start) > state.timeToBuild {
			fmt.Println("Product finished")
			state.behavior.Become(state.Finished)
			context.Send(state.marketingTeamPID, ProductionLaunch{})
		}
	}
}

func NewDeveloperTeam(start time.Time, marketingTeamPID *actor.PID, timeToBuild time.Duration) func() actor.Actor {
	fn := func() actor.Actor {
		act := &DeveloperTeam{
			timeToBuild:      timeToBuild,
			marketingTeamPID: marketingTeamPID,
			start:            start,
			behavior:         actor.NewBehavior(),
		}
		act.behavior.Become(act.Building)
		return act
	}
	return fn
}
