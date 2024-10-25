package library

import (
	"errors"
	"log"
)

var (
	ErrFailRequests = errors.New("error to send request")
	Log             = log.Fatalln
)
