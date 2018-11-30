package loader

import (
	"github.com/ipsn/go-ipfs/plugin"
	pluginbadgerds "github.com/ipsn/go-ipfs/plugin/plugins/badgerds"
	pluginflatfs "github.com/ipsn/go-ipfs/plugin/plugins/flatfs"
	pluginipldgit "github.com/ipsn/go-ipfs/plugin/plugins/git"
	pluginlevelds "github.com/ipsn/go-ipfs/plugin/plugins/levelds"
)

// DO NOT EDIT THIS FILE
// This file is being generated as part of plugin build process
// To change it, modify the plugin/loader/preload.sh

var preloadPlugins = []plugin.Plugin{
	pluginipldgit.Plugins[0],
	pluginbadgerds.Plugins[0],
	pluginflatfs.Plugins[0],
	pluginlevelds.Plugins[0],
}
