package sl

import (
	"fmt"
	"log/slog"
)

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}

func Op(defaultOp, funcName string) slog.Attr {
	return slog.Attr{
		Key:   "op",
		Value: slog.StringValue(fmt.Sprintf("%s.%s", defaultOp, funcName)),
	}
}
