// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// GCPFailureDomainApplyConfiguration represents an declarative configuration of the GCPFailureDomain type for use
// with apply.
type GCPFailureDomainApplyConfiguration struct {
	Zone *string `json:"zone,omitempty"`
}

// GCPFailureDomainApplyConfiguration constructs an declarative configuration of the GCPFailureDomain type for use with
// apply.
func GCPFailureDomain() *GCPFailureDomainApplyConfiguration {
	return &GCPFailureDomainApplyConfiguration{}
}

// WithZone sets the Zone field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Zone field is set to the value of the last call.
func (b *GCPFailureDomainApplyConfiguration) WithZone(value string) *GCPFailureDomainApplyConfiguration {
	b.Zone = &value
	return b
}
