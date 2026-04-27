module github.com/plexusone/omnidxi

go 1.22

require github.com/plexusone/omnidxi-core v0.1.0

require github.com/google/uuid v1.6.0 // indirect

// Provider adapters (uncomment as they become available)
// require github.com/plexusone/omni-amplitude/omnidxi v0.1.0
// require github.com/plexusone/omni-mixpanel/omnidxi v0.1.0

// Local development - remove before release
replace github.com/plexusone/omnidxi-core => ../omnidxi-core
