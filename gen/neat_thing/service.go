// Code generated by goa v3.0.2, DO NOT EDIT.
//
// neatThing service
//
// Command:
// $ goa gen github.com/taothit/one-neat-thing-today/design

package neatthing

import (
	"context"
)

// Allows access to discover neat things.
type Service interface {
	// NeatThingToday implements neatThingToday.
	NeatThingToday(context.Context) (res *NeatThing, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "neatThing"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"neatThingToday"}

// NeatThing is the result type of the neatThing service neatThingToday method.
type NeatThing struct {
	// The neat thing
	Name *string
	// What the neat thing is
	Definition *string
	// Illustrative link for the neat thing
	Link *string
	// When this was a neat thing
	Date         *string
	Bibliography []string
}