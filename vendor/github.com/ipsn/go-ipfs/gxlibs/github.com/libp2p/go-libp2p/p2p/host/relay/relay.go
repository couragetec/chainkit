package relay

import (
	"context"

	basic "github.com/ipsn/go-ipfs/gxlibs/github.com/libp2p/go-libp2p/p2p/host/basic"

	discovery "github.com/ipsn/go-ipfs/gxlibs/github.com/libp2p/go-libp2p-discovery"
	host "github.com/ipsn/go-ipfs/gxlibs/github.com/libp2p/go-libp2p-host"
	ma "github.com/ipsn/go-ipfs/gxlibs/github.com/multiformats/go-multiaddr"
)

// RelayHost is a Host that provides Relay services.
type RelayHost struct {
	*basic.BasicHost
	advertise discovery.Advertiser
	addrsF    basic.AddrsFactory
}

// New constructs a new RelayHost
func NewRelayHost(ctx context.Context, bhost *basic.BasicHost, advertise discovery.Advertiser) *RelayHost {
	h := &RelayHost{
		BasicHost: bhost,
		addrsF:    bhost.AddrsFactory,
		advertise: advertise,
	}
	bhost.AddrsFactory = h.hostAddrs
	discovery.Advertise(ctx, advertise, RelayRendezvous)
	return h
}

func (h *RelayHost) hostAddrs(addrs []ma.Multiaddr) []ma.Multiaddr {
	return filterUnspecificRelay(h.addrsF(addrs))
}

var _ host.Host = (*RelayHost)(nil)
