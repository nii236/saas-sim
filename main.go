package main

import (
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type CurrentTime struct {
	time.Time
}

func main() {
	currentTime := time.Date(2019, 9, 1, 0, 0, 0, 0, time.Local)
	timeToBuild := 1 * time.Hour
	rootContext := actor.EmptyRootContext
	customerPIDs := []*actor.PID{}
	for i := 0; i < 10000; i++ {
		customer := actor.PropsFromProducer(NewCustomer)
		customerPIDs = append(customerPIDs, rootContext.Spawn(customer))
	}

	marketingTeamPID := rootContext.Spawn(actor.PropsFromProducer(NewMarketingTeam(customerPIDs)))
	developerTeamPID := rootContext.Spawn(actor.PropsFromProducer(NewDeveloperTeam(currentTime, marketingTeamPID, timeToBuild)))

	t := time.NewTicker(10 * time.Millisecond)
	for {
		select {
		case <-t.C:
			msg := CurrentTime{currentTime}
			for _, pid := range customerPIDs {
				rootContext.Send(pid, msg)
			}

			rootContext.Send(marketingTeamPID, msg)
			rootContext.Send(developerTeamPID, msg)

			currentTime = currentTime.Add(1 * time.Hour)
		}
	}
}
