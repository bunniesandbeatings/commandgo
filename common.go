package commandgo

import "log"

type HandlesErrors struct {
	ErrorHandler func(error)
}

func NewHandlesErrors() HandlesErrors {
	return HandlesErrors{
		ErrorHandler: func(err error) {
			log.Panic(err)
		},
	}
}
