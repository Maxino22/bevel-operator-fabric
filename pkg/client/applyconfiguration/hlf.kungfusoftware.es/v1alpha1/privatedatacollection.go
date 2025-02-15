/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// PrivateDataCollectionApplyConfiguration represents a declarative configuration of the PrivateDataCollection type for use
// with apply.
type PrivateDataCollectionApplyConfiguration struct {
	Name              *string                                                   `json:"name,omitempty"`
	Policy            *string                                                   `json:"policy,omitempty"`
	RequiredPeerCount *int32                                                    `json:"requiredPeerCount,omitempty"`
	MaxPeerCount      *int32                                                    `json:"maxPeerCount,omitempty"`
	BlockToLive       *uint64                                                   `json:"blockToLive,omitempty"`
	MemberOnlyRead    *bool                                                     `json:"memberOnlyRead,omitempty"`
	MemberOnlyWrite   *bool                                                     `json:"memberOnlyWrite,omitempty"`
	EndorsementPolicy *PrivateDataCollectionEndorsementPolicyApplyConfiguration `json:"endorsementPolicy,omitempty"`
}

// PrivateDataCollectionApplyConfiguration constructs a declarative configuration of the PrivateDataCollection type for use with
// apply.
func PrivateDataCollection() *PrivateDataCollectionApplyConfiguration {
	return &PrivateDataCollectionApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *PrivateDataCollectionApplyConfiguration) WithName(value string) *PrivateDataCollectionApplyConfiguration {
	b.Name = &value
	return b
}

// WithPolicy sets the Policy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Policy field is set to the value of the last call.
func (b *PrivateDataCollectionApplyConfiguration) WithPolicy(value string) *PrivateDataCollectionApplyConfiguration {
	b.Policy = &value
	return b
}

// WithRequiredPeerCount sets the RequiredPeerCount field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RequiredPeerCount field is set to the value of the last call.
func (b *PrivateDataCollectionApplyConfiguration) WithRequiredPeerCount(value int32) *PrivateDataCollectionApplyConfiguration {
	b.RequiredPeerCount = &value
	return b
}

// WithMaxPeerCount sets the MaxPeerCount field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MaxPeerCount field is set to the value of the last call.
func (b *PrivateDataCollectionApplyConfiguration) WithMaxPeerCount(value int32) *PrivateDataCollectionApplyConfiguration {
	b.MaxPeerCount = &value
	return b
}

// WithBlockToLive sets the BlockToLive field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BlockToLive field is set to the value of the last call.
func (b *PrivateDataCollectionApplyConfiguration) WithBlockToLive(value uint64) *PrivateDataCollectionApplyConfiguration {
	b.BlockToLive = &value
	return b
}

// WithMemberOnlyRead sets the MemberOnlyRead field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MemberOnlyRead field is set to the value of the last call.
func (b *PrivateDataCollectionApplyConfiguration) WithMemberOnlyRead(value bool) *PrivateDataCollectionApplyConfiguration {
	b.MemberOnlyRead = &value
	return b
}

// WithMemberOnlyWrite sets the MemberOnlyWrite field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MemberOnlyWrite field is set to the value of the last call.
func (b *PrivateDataCollectionApplyConfiguration) WithMemberOnlyWrite(value bool) *PrivateDataCollectionApplyConfiguration {
	b.MemberOnlyWrite = &value
	return b
}

// WithEndorsementPolicy sets the EndorsementPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EndorsementPolicy field is set to the value of the last call.
func (b *PrivateDataCollectionApplyConfiguration) WithEndorsementPolicy(value *PrivateDataCollectionEndorsementPolicyApplyConfiguration) *PrivateDataCollectionApplyConfiguration {
	b.EndorsementPolicy = value
	return b
}
