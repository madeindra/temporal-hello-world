package app

import (
	"context"
	"fmt"
)

func ComposeGreeting(ctx context.Context, name string) (string, error) {
	greetings := fmt.Sprintf("Hello %s", name)
	return greetings, nil
}
