package types

// An DnsResolver represent an object that can resolve SRV record
type DnsResolver interface {
	SRVResolve(host string) (bool, string, uint16) // Resolve an SRV Record and return if found, host and port
}
