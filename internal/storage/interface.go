package storage

import "context"

type Repo interface {
	AddMessage(ctx context.Context, message string, param string) error
	TakeLastMessages(ctx context.Context, param string) ([]string, error)
}
