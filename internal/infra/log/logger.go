package log

import (
	"os"

	"github.com/go-kit/kit/log"
)

func NewLogger() (logger log.Logger) {
	w := log.NewSyncWriter(os.Stderr)
	logger = log.NewJSONLogger(w)
	logger = log.With(logger, "timestamp", log.DefaultTimestampUTC)
	logger = log.With(logger, "called from", log.DefaultCaller)
	return
}
