package platform

import "context"

// S is a service
type S interface {
	Run(context.Context) error
}
