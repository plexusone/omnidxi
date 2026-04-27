// Package omnidxi provides a batteries-included Digital Experience Intelligence
// (DXI) client with support for multiple analytics providers.
//
// This package re-exports core types from omnidxi-core and provides convenience
// functions for working with multiple providers simultaneously.
//
// For individual provider usage, import the provider packages directly:
//
//	import "github.com/plexusone/omni-amplitude/omnidxi"
//	import "github.com/plexusone/omni-mixpanel/omnidxi"
package omnidxi

import (
	"context"
	"sync"

	core "github.com/plexusone/omnidxi-core"
	amplitude "github.com/plexusone/omni-amplitude/omnidxi"
	mixpanel "github.com/plexusone/omni-mixpanel/omnidxi"
)

// Re-export core types for convenience.
type (
	Tracker      = core.Tracker
	TrackerInfo  = core.TrackerInfo
	Event        = core.Event
	EventType    = core.EventType
	EventContext = core.EventContext
	User         = core.User
	UserTraits   = core.UserTraits
	Group        = core.Group
	Alias        = core.Alias
	Config       = core.Config
	Option       = core.Option
)

// Re-export event types.
const (
	EventTypePageView    = core.EventTypePageView
	EventTypePageLeave   = core.EventTypePageLeave
	EventTypeUIClick     = core.EventTypeUIClick
	EventTypeUIInput     = core.EventTypeUIInput
	EventTypeUIScroll    = core.EventTypeUIScroll
	EventTypeUISubmit    = core.EventTypeUISubmit
	EventTypeStateChange = core.EventTypeStateChange
	EventTypeAPIRequest  = core.EventTypeAPIRequest
	EventTypeAPIResponse = core.EventTypeAPIResponse
	EventTypeJourneyStep = core.EventTypeJourneyStep
	EventTypeError       = core.EventTypeError
	EventTypePerformance = core.EventTypePerformance
	EventTypeCustom      = core.EventTypeCustom
)

// Re-export constructors.
var (
	NewEvent   = core.NewEvent
	NewUser    = core.NewUser
	NewGroup   = core.NewGroup
	NewAlias   = core.NewAlias
	NewConfig  = core.NewConfig
	WithAPIKey = core.WithAPIKey
	WithLogger = core.WithLogger
	WithDebug  = core.WithDebug
)

// Re-export errors.
var (
	ErrDisabled         = core.ErrDisabled
	ErrNoAPIKey         = core.ErrNoAPIKey
	ErrInvalidUserID    = core.ErrInvalidUserID
	ErrInvalidEventName = core.ErrInvalidEventName
	ErrFlushFailed      = core.ErrFlushFailed
	ErrClosed           = core.ErrClosed
)

// Provider constructors for convenience.
var (
	// NewAmplitudeTracker creates an Amplitude tracker.
	NewAmplitudeTracker = amplitude.New

	// NewMixpanelTracker creates a Mixpanel tracker.
	NewMixpanelTracker = mixpanel.New
)

// MultiTracker sends events to multiple providers simultaneously.
type MultiTracker struct {
	trackers []core.Tracker
	mu       sync.RWMutex
}

// NewMultiTracker creates a new MultiTracker with the given trackers.
func NewMultiTracker(trackers ...core.Tracker) *MultiTracker {
	return &MultiTracker{
		trackers: trackers,
	}
}

// Add adds a tracker to the multi-tracker.
func (m *MultiTracker) Add(t core.Tracker) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.trackers = append(m.trackers, t)
}

// Track sends an event to all trackers.
func (m *MultiTracker) Track(ctx context.Context, event core.Event) error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var firstErr error
	for _, t := range m.trackers {
		if err := t.Track(ctx, event); err != nil && firstErr == nil {
			firstErr = err
		}
	}
	return firstErr
}

// Identify sends user identification to all trackers.
func (m *MultiTracker) Identify(ctx context.Context, user core.User) error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var firstErr error
	for _, t := range m.trackers {
		if err := t.Identify(ctx, user); err != nil && firstErr == nil {
			firstErr = err
		}
	}
	return firstErr
}

// Group sends group association to all trackers.
func (m *MultiTracker) Group(ctx context.Context, group core.Group) error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var firstErr error
	for _, t := range m.trackers {
		if err := t.Group(ctx, group); err != nil && firstErr == nil {
			firstErr = err
		}
	}
	return firstErr
}

// Alias sends alias to all trackers.
func (m *MultiTracker) Alias(ctx context.Context, alias core.Alias) error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var firstErr error
	for _, t := range m.trackers {
		if err := t.Alias(ctx, alias); err != nil && firstErr == nil {
			firstErr = err
		}
	}
	return firstErr
}

// Flush flushes all trackers.
func (m *MultiTracker) Flush(ctx context.Context) error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var firstErr error
	for _, t := range m.trackers {
		if err := t.Flush(ctx); err != nil && firstErr == nil {
			firstErr = err
		}
	}
	return firstErr
}

// Close closes all trackers.
func (m *MultiTracker) Close() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	var firstErr error
	for _, t := range m.trackers {
		if err := t.Close(); err != nil && firstErr == nil {
			firstErr = err
		}
	}
	m.trackers = nil
	return firstErr
}

// Ensure MultiTracker implements Tracker.
var _ core.Tracker = (*MultiTracker)(nil)
