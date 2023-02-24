package storm

// #include <StormLib.h>
import "C"
import "fmt"

type StormError struct {
	Code    uint32
	Message string
}

// Implementation of the error interface.
func (err *StormError) Error() string {
	return err.Message
}

func newStormError(code uint32, message string) *StormError {
	return &StormError{
		Code:    code,
		Message: fmt.Sprintf("storm: %s (code: %d)", message, code),
	}
}
