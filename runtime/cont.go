package runtime

import (
	"os"
	"syscall"
)

// Cont is an interface for things that can be run in a Thread.  Implementations
// of Cont a typically an invocation of a Lua function or an invocation of a Go
// function.
//
// TODO: document the methods.
type Cont interface {
	Push(*Runtime, Value)
	PushEtc(*Runtime, []Value)
	RunInThread(*Thread) (Cont, *Error)
	Next() Cont
	Parent() Cont
	DebugInfo() *DebugInfo
}

// Push is a convenience method that pushes a number of values to the
// continuation c.
func (r *Runtime) Push(c Cont, vals ...Value) {
	// TODO: should that consume CPU?
	for _, v := range vals {
		c.Push(r, v)
	}
}

func (r *Runtime) Push1(c Cont, v Value) {
	c.Push(r, v)
}

func (r *Runtime) PushIoError(c Cont, ioErr error) *Error {
	// It is not specified in the Lua docs, but the test suite expects an
	// errno as the third value returned in case of an error.  I'm not sure
	// the conversion to syscall.Errno is future-proof though?

	var err error
	switch tErr := ioErr.(type) {
	case *os.PathError:
		err = tErr.Unwrap()
	case *os.LinkError:
		err = tErr.Unwrap()
	default:
		return NewErrorE(ioErr)
	}
	r.Push1(c, NilValue)
	r.Push1(c, StringValue(ioErr.Error()))
	if errno, ok := err.(syscall.Errno); ok {
		r.Push1(c, IntValue(int64(errno)))
	}
	return nil
}

func (r *Runtime) ProcessIoError(c Cont, ioErr error) (Cont, *Error) {
	if err := r.PushIoError(c, ioErr); err != nil {
		return nil, err
	}
	return c, nil
}
