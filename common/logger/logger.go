package logger

import (
	"context"
	"fmt"
)

func Error(ctx context.Context, message string, attributes ...interface{}) {
	fmt.Println(message)
}
