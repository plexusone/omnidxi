# OmniDXI

[![Go CI][go-ci-svg]][go-ci-url]
[![Go Lint][go-lint-svg]][go-lint-url]
[![Go SAST][go-sast-svg]][go-sast-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![Visualization][viz-svg]][viz-url]
[![License][license-svg]][license-url]

 [go-ci-svg]: https://github.com/plexusone/omnidxi/actions/workflows/go-ci.yaml/badge.svg?branch=main
 [go-ci-url]: https://github.com/plexusone/omnidxi/actions/workflows/go-ci.yaml
 [go-lint-svg]: https://github.com/plexusone/omnidxi/actions/workflows/go-lint.yaml/badge.svg?branch=main
 [go-lint-url]: https://github.com/plexusone/omnidxi/actions/workflows/go-lint.yaml
 [go-sast-svg]: https://github.com/plexusone/omnidxi/actions/workflows/go-sast-codeql.yaml/badge.svg?branch=main
 [go-sast-url]: https://github.com/plexusone/omnidxi/actions/workflows/go-sast-codeql.yaml
 [goreport-svg]: https://goreportcard.com/badge/github.com/plexusone/omnidxi
 [goreport-url]: https://goreportcard.com/report/github.com/plexusone/omnidxi
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/plexusone/omnidxi
 [docs-godoc-url]: https://pkg.go.dev/github.com/plexusone/omnidxi
 [docs-mkdoc-svg]: https://img.shields.io/badge/docs-guide-blue.svg
 [docs-mkdoc-url]: https://plexusone.github.io/omnidxi
 [viz-svg]: https://img.shields.io/badge/repo-visualization-blue.svg
 [viz-url]: https://mango-dune-07a8b7110.1.azurestaticapps.net/?repo=plexusone%2Fomnidxi
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/plexusone/omnidxi/blob/main/LICENSE

Batteries-included Digital Experience Intelligence (DXI) client for Go.

## Overview

`omnidxi` provides a unified interface for product analytics platforms like Amplitude, Mixpanel, Heap, and Pendo. It re-exports types from `omnidxi-core` and includes utilities for working with multiple providers.

## Installation

```bash
go get github.com/plexusone/omnidxi
```

## Quick Start

```go
package main

import (
    "context"

    "github.com/plexusone/omnidxi"
    "github.com/plexusone/omni-amplitude/omnidxi/amplitude"
    "github.com/plexusone/omni-mixpanel/omnidxi/mixpanel"
)

func main() {
    ctx := context.Background()

    // Create provider trackers
    amp := amplitude.New(omnidxi.WithAPIKey("amp-key"))
    mix := mixpanel.New(omnidxi.WithAPIKey("mix-token"))

    // Combine into multi-tracker
    tracker := omnidxi.NewMultiTracker(amp, mix)
    defer tracker.Close()

    // Track events to all providers
    event := omnidxi.NewEvent(omnidxi.EventTypePageView, "Home Viewed").
        WithUserID("user_123").
        WithProperty("source", "direct")

    tracker.Track(ctx, event)

    // Identify user across all providers
    user := omnidxi.NewUser("user_123").
        WithTraits(omnidxi.UserTraits{
            Email: "user@example.com",
            Name:  "Jane Doe",
        })

    tracker.Identify(ctx, user)

    // Flush before exit
    tracker.Flush(ctx)
}
```

## Multi-Tracker

Send events to multiple providers simultaneously:

```go
tracker := omnidxi.NewMultiTracker(amplitude, mixpanel, heap)

// All providers receive the event
tracker.Track(ctx, event)
```

## No-op Tracker

For testing or when tracking should be disabled:

```go
tracker := omnidxi.NewNoopTracker()
```

## Provider Packages

| Provider | Package | Status |
|----------|---------|--------|
| Amplitude | `github.com/plexusone/omni-amplitude/omnidxi` | Planned |
| Mixpanel | `github.com/plexusone/omni-mixpanel/omnidxi` | Planned |
| Heap | `github.com/plexusone/omni-heap/omnidxi` | Future |
| Pendo | `github.com/plexusone/omni-pendo/omnidxi` | Future |

## Architecture

```
┌─────────────────────────────────────────────────────┐
│                    Application                       │
└─────────────────────────┬───────────────────────────┘
                          │
                          ▼
┌─────────────────────────────────────────────────────┐
│                     omnidxi                          │
│              (batteries-included)                    │
│                                                      │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐  │
│  │ MultiTracker│  │ NoopTracker │  │  Re-exports │  │
│  └─────────────┘  └─────────────┘  └─────────────┘  │
└─────────────────────────┬───────────────────────────┘
                          │
                          ▼
┌─────────────────────────────────────────────────────┐
│                   omnidxi-core                       │
│           (interfaces, types, schema)                │
└─────────────────────────┬───────────────────────────┘
                          │
        ┌─────────────────┼─────────────────┐
        ▼                 ▼                 ▼
┌───────────────┐ ┌───────────────┐ ┌───────────────┐
│omni-amplitude │ │ omni-mixpanel │ │   omni-heap   │
│   /omnidxi    │ │   /omnidxi    │ │   /omnidxi    │
└───────────────┘ └───────────────┘ └───────────────┘
        │                 │                 │
        ▼                 ▼                 ▼
┌───────────────┐ ┌───────────────┐ ┌───────────────┐
│ analytics-go  │ │ mixpanel-go   │ │   heap SDK    │
│  (official)   │ │  (official)   │ │  (official)   │
└───────────────┘ └───────────────┘ └───────────────┘
```

## Related Projects

- [omnidxi-core](https://github.com/plexusone/omnidxi-core) - Core interfaces
- [ProductGraph](https://github.com/plexusone/productgraph) - Product intelligence platform
- [OmniObserve](https://github.com/plexusone/omniobserve) - AI/ML observability

## License

MIT
