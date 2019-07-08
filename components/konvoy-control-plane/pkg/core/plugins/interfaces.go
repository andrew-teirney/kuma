package plugins

import (
	core_discovery "github.com/Kong/konvoy/components/konvoy-control-plane/pkg/core/discovery"
	core_store "github.com/Kong/konvoy/components/konvoy-control-plane/pkg/core/resources/store"
	core_runtime "github.com/Kong/konvoy/components/konvoy-control-plane/pkg/core/runtime"
)

type Plugin interface{}

type PluginConfig interface{}

type PluginContext = core_runtime.BuilderContext

type MutablePluginContext = core_runtime.Builder

// BootstrapPlugin is responsible for environment-specific initialization at start up,
// e.g. Kubernetes-specific part of configuration.
// Unlike other plugins, can mutate plugin context directly.
type BootstrapPlugin interface {
	Plugin
	Bootstrap(*MutablePluginContext, PluginConfig) error
}

// ResourceStorePlugin is responsible for instantiating a particular ResourceStore.
type ResourceStorePlugin interface {
	Plugin
	NewResourceStore(PluginContext, PluginConfig) (core_store.ResourceStore, error)
}

// DiscoveryPlugin is responsible for instantiating a particular DiscoverySource.
type DiscoveryPlugin interface {
	Plugin
	NewDiscoverySource(PluginContext, PluginConfig) (core_discovery.DiscoverySource, error)
}