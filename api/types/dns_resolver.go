package types

type DnsResolver interface {
	SRVResolve(host string) (bool, string, uint16)
}
