package library

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrWhenSendRequest = errors.New("error to send request")
	Log                = log.Default().Fatalln
	Pl                 = fmt.Println
	Pf                 = fmt.Printf
	Sf                 = fmt.Sprintf
	Err                = fmt.Errorf
)
