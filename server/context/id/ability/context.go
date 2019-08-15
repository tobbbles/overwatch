package ability

import (
	"context"
	"errors"
	"strconv"
)

const abilityIDKey = "abilityIDContextKey"

// NewContext takes an ability ID and places it into the returned context
func NewContext(parent context.Context, abilityID string) (context.Context, error) {
	id, err := strconv.Atoi(abilityID)
	if err != nil {
		return nil, err
	}

	return context.WithValue(parent, abilityIDKey, id), nil
}

// FromContext retrieves the ability_id from the provided context
func FromContext(ctx context.Context) (int, error) {
	val := ctx.Value(abilityIDKey)

	if ability, ok := val.(int); ok {
		return ability, nil
	}

	return 0, ErrNoAbilityIDInContext
}

var (
	ErrNoAbilityIDInContext = errors.New("no ability id found in given context")
)
