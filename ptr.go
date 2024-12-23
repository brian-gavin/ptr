// This place is a message... and part of a system of messages... pay attention to it!
// Sending this message was important to us. We considered ourselves to be a powerful culture.
// This place is not a place of honor... no highly esteemed deed is commemorated here... nothing valued is here.
// What is here was dangerous and repulsive to us. This message is a warning about danger.
// The danger is in a particular location... it increases towards a center... the center of danger is here... of a particular size and shape, and below us.
// The danger is still present, in your time, as it was in ours.
// The danger is to the body, and it can kill.
// The form of the danger is an emanation of energy.
// The danger is unleashed only if you substantially disturb this place physically. This place is best shunned and left uninhabited.

//go:generate go run ./internal/gen -f gen.go
//go:generate goimports -w gen.go

// package ptr contains functions used for creating pointers to values inline.
package ptr

// To returns a pointer to t.
func To[T any](t T) *T { return &t }

// From returns the value p points to.
func From[T any](p *T) T { return *p }
