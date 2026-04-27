# Release Notes - v0.1.0

**Release Date:** 2026-04-27

## Overview

Initial release of omnidxi, the batteries-included DXI (Digital Experience Intelligence) client for Go.

This package provides a high-level client that re-exports all core types from omnidxi-core, bundles Amplitude and Mixpanel adapters, and adds the MultiTracker for sending events to multiple analytics providers simultaneously.

## Highlights

- Batteries-included DXI client for Go
- MultiTracker for sending events to multiple providers simultaneously
- Bundled Amplitude and Mixpanel provider adapters

## What's Included

### Bundled Providers

Convenience constructors for creating provider trackers:

```go
import "github.com/plexusone/omnidxi"

// Create Amplitude tracker
amp := omnidxi.NewAmplitudeTracker(omnidxi.WithAPIKey("amp-key"))

// Create Mixpanel tracker
mp := omnidxi.NewMixpanelTracker(omnidxi.WithAPIKey("mp-token"))
```

### MultiTracker

Send events to multiple DXI providers with a single call:

```go
import "github.com/plexusone/omnidxi"

// Create provider trackers
amp := omnidxi.NewAmplitudeTracker(omnidxi.WithAPIKey("amp-key"))
mp := omnidxi.NewMixpanelTracker(omnidxi.WithAPIKey("mp-token"))

// Combine into MultiTracker
tracker := omnidxi.NewMultiTracker(amp, mp)

// Single call sends to both providers
tracker.Track(ctx, omnidxi.Event{
    Type:   omnidxi.EventTypeUIClick,
    Name:   "button_clicked",
    UserID: "user-123",
})
```

Features:

- Thread-safe `Add()` method for runtime provider addition
- Aggregated error handling across all providers
- Implements the standard `Tracker` interface

### NoopTracker

A no-operation tracker for testing and disabled scenarios:

```go
tracker := omnidxi.NewNoopTracker()
// All methods succeed without side effects
```

### Re-exported Types

All core types from omnidxi-core are re-exported for convenience:

- `Tracker` interface
- `Event`, `EventType`, `EventContext`
- `User`, `Group`, `Alias`, `UserTraits`
- `Config` and functional options
- Error types

## Requirements

- Go 1.22 or later

## Dependencies

- `github.com/plexusone/omnidxi-core` v0.1.0
- `github.com/plexusone/omni-amplitude` v0.1.0
- `github.com/plexusone/omni-mixpanel` v0.1.0

## Installation

```bash
go get github.com/plexusone/omnidxi@v0.1.0
```

## Related Packages

- [omnidxi-core](https://github.com/plexusone/omnidxi-core) - Core interfaces and types
- [omni-amplitude](https://github.com/plexusone/omni-amplitude) - Amplitude adapter (standalone)
- [omni-mixpanel](https://github.com/plexusone/omni-mixpanel) - Mixpanel adapter (standalone)
