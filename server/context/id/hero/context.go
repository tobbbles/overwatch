package hero

import (
	"context"
	"errors"
	"strconv"
)

const heroIDKey = "heroIDContextKey"

// NewContext takes an hero ID and places it into the returned context
func NewContext(parent context.Context, heroID string) (context.Context, error) {
	id, err := strconv.Atoi(heroID)
	if err != nil {
		return nil, err
	}

	return context.WithValue(parent, heroIDKey, id), nil
}

// FromContext retrieves the hero_id from the provided context
func FromContext(ctx context.Context) (int, error) {
	val := ctx.Value(heroIDKey)

	if hero, ok := val.(int); ok {
		return hero, nil
	}

	return 0, ErrNoAbilityIDInContext
}

var (
	ErrNoAbilityIDInContext = errors.New("no hero id found in given context")
)
