package validation

import "github.com/grafana/loki/v3/pkg/runtime"

// Limits is an alias to runtime.Limits to avoid rewriting all imports at once
// Deprecated: Use runtime.Limits instead
type Limits = runtime.Limits

// TenantLimits is an alias to runtime.Limits to avoid rewriting all imports at once
// Deprecated: Use runtime.Overrides instead
type TenantLimits = runtime.TenantLimits

// Overrides is an alias to runtime.Limits to avoid rewriting all imports at once
// Deprecated: Use runtime.Overrides instead
type Overrides = runtime.Overrides
