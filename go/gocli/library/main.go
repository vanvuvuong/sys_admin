package library

import (
	"errors"
	"fmt"
	"time"

	"github.com/op/go-logging"
)

var (
	ErrWhenSendRequest     = errors.New("error to send request")
	Pl                     = fmt.Println
	Pf                     = fmt.Printf
	Sf                     = fmt.Sprintf
	Err                    = fmt.Errorf
	log                    = logging.MustGetLogger("library")
	RetryCount             = 5
	RetryWait              = 5 * time.Second
	RequestTimeout         = 15 * time.Second
	DownloadRequestTimeout = 30 * time.Minute
	UploadRequestTimeout   = 30 * time.Minute
)
