package segment

import (
	"errors"
	"fmt"
)

// All errors that segment could return.
var (
	ErrPartSizeInvalid     = errors.New("part size invalid")
	ErrPartIntersected     = errors.New("part intersected")
	ErrSegmentNotInitiated = errors.New("segment not initiated")
	ErrSegmentPartsEmpty   = errors.New("segment Parts are empty")
	ErrSegmentNotFulfilled = errors.New("segment not fulfilled")
)

// Error represents error related to endpoint.
type Error struct {
	Op  string
	Err error

	Segment *Segment
	Part    *Part
}

func (e *Error) Error() string {
	if e.Part == nil {
		return fmt.Sprintf("%s: %s: %s", e.Op, e.Segment, e.Err.Error())
	}
	return fmt.Sprintf("%s: %s, %s: %s", e.Op, e.Segment, e.Part, e.Err.Error())
}

// Unwrap implements xerrors.Wrapper
func (e *Error) Unwrap() error {
	return e.Err
}