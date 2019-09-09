package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/guptarohit/asciigraph"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type CustomerSubscriber struct {
	*actor.PID
	time.Time
}
type MarketingTeam struct {
	currentTime     time.Time
	lastPostCTA     time.Time
	lastPreCTA      time.Time
	behavior        actor.Behavior
	customerPIDs    []*actor.PID
	subscribers     []*CustomerSubscriber
	points          Points
	productLaunched bool
}
type PostLaunchCallToAction struct {
}
type PrelaunchCallToAction struct {
}

func (state *MarketingTeam) Receive(context actor.Context) {
	state.behavior.Receive(context)
}

func (state *MarketingTeam) Idle(context actor.Context) {
	switch context.Message().(type) {
	case ProductionLaunch:
		state.productLaunched = true
		state.behavior.Become(state.Running)
	}
}

func (state *MarketingTeam) Running(context actor.Context) {
	switch msg := context.Message().(type) {
	case DeleteAccount:
		for i, sub := range state.subscribers {
			if sub.PID == context.Sender() {
				state.subscribers = append(state.subscribers[:i], state.subscribers[i:]...)
			}
		}
	case Signup:
		state.subscribers = append(state.subscribers, &CustomerSubscriber{context.Sender(), state.currentTime})
		AddPoint(state.points, state.currentTime)
	case CurrentTime:
		state.currentTime = msg.Time
		if !state.productLaunched {
			return
		}
		if msg.Time.Sub(state.lastPostCTA).Hours() > 24 {
			state.lastPostCTA = msg.Time
			for _, pid := range state.customerPIDs {
				context.Request(pid, PostLaunchCallToAction{})
			}
		}
		// cmd := exec.Command("clear")
		// cmd.Stdout = os.Stdout
		// cmd.Run()
		Plot(state.points)
	}
}

type XTick string

type YValue int

type Points map[XTick]YValue

func AddPoint(points map[XTick]YValue, now time.Time) {
	val, ok := points[XTick(now.Format("20060102"))]
	if !ok {
		points[XTick(now.Format("20060102"))] = 1
		return
	}
	points[XTick(now.Format("20060102"))] = val + 1
}

func Plot(points Points) {
	var keys []string
	for k := range points {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)

	sorted := []YValue{}
	for _, k := range keys {
		sorted = append(sorted, points[XTick(k)])
	}
	data := []float64{}
	for _, y := range sorted {
		data = append(data, float64(y))
	}
	if len(data) < 2 {
		return
	}
	fmt.Println(len(data))
	graph := asciigraph.Plot(data)

	fmt.Println(graph)
}

func NewMarketingTeam(customerPIDs []*actor.PID) func() actor.Actor {
	fn := func() actor.Actor {

		act := &MarketingTeam{
			customerPIDs:    customerPIDs,
			subscribers:     []*CustomerSubscriber{},
			lastPostCTA:     time.Now(),
			lastPreCTA:      time.Now(),
			behavior:        actor.NewBehavior(),
			productLaunched: false,
			points:          map[XTick]YValue{},
		}
		act.behavior.Become(act.Idle)
		return act
	}
	return fn
}
