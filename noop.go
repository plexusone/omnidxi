package omnidxi

import (
	"context"

	core "github.com/plexusone/omnidxi-core"
)

// NoopTracker is a tracker that does nothing.
// Useful for testing or when tracking should be disabled.
type NoopTracker struct{}

// NewNoopTracker creates a new no-op tracker.
func NewNoopTracker() *NoopTracker {
	return &NoopTracker{}
}

func (n *NoopTracker) Track(ctx context.Context, event core.Event) error {
	return nil
}

func (n *NoopTracker) Identify(ctx context.Context, user core.User) error {
	return nil
}

func (n *NoopTracker) Group(ctx context.Context, group core.Group) error {
	return nil
}

func (n *NoopTracker) Alias(ctx context.Context, alias core.Alias) error {
	return nil
}

func (n *NoopTracker) Flush(ctx context.Context) error {
	return nil
}

func (n *NoopTracker) Close() error {
	return nil
}

// Ensure NoopTracker implements Tracker.
var _ core.Tracker = (*NoopTracker)(nil)
