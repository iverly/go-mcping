package dns

import (
	"context"
	"net"
)

type resolver struct {
	internalResolver *net.Resolver
}

// Create a new resolver with the DNS server of Cloudflare (1.1.1.1)
// The DNS server can be changed by using resolver#SetInternalResolver
func NewResolver() *resolver {
	return &resolver{
		internalResolver: &net.Resolver{
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				return (&net.Dialer{}).DialContext(ctx, network, "1.1.1.1:53")
			},
		},
	}
}

// Set internal resolver of an resolver
// By default, the resolver using the Cloudflare DNS server (1.1.1.1)
//
// Example
//	&net.Resolver{
//		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
//			return (&net.Dialer{}).DialContext(ctx, network, "1.1.1.1:53")
//		},
//	},
func (r *resolver) SetInternalResolver(internalResolver *net.Resolver) {
	r.internalResolver = internalResolver
}

func (r *resolver) SRVResolve(host string) (bool, string, uint16) {
	_, srvs, err := r.internalResolver.LookupSRV(
		context.Background(), "minecraft", "tcp", host,
	)
	if err != nil || len(srvs) == 0 {
		return false, "", 0
	}
	return true, srvs[0].Target[:len(srvs[0].Target)-1], srvs[0].Port
}
