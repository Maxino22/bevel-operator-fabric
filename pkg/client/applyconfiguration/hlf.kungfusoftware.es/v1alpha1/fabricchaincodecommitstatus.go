/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kfsoftware/hlf-operator/pkg/apis/hlf.kungfusoftware.es/v1alpha1"
	status "github.com/kfsoftware/hlf-operator/pkg/status"
)

// FabricChaincodeCommitStatusApplyConfiguration represents a declarative configuration of the FabricChaincodeCommitStatus type for use
// with apply.
type FabricChaincodeCommitStatusApplyConfiguration struct {
	Conditions    *status.Conditions         `json:"conditions,omitempty"`
	Message       *string                    `json:"message,omitempty"`
	Status        *v1alpha1.DeploymentStatus `json:"status,omitempty"`
	TransactionID *string                    `json:"transactionID,omitempty"`
}

// FabricChaincodeCommitStatusApplyConfiguration constructs a declarative configuration of the FabricChaincodeCommitStatus type for use with
// apply.
func FabricChaincodeCommitStatus() *FabricChaincodeCommitStatusApplyConfiguration {
	return &FabricChaincodeCommitStatusApplyConfiguration{}
}

// WithConditions sets the Conditions field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Conditions field is set to the value of the last call.
func (b *FabricChaincodeCommitStatusApplyConfiguration) WithConditions(value status.Conditions) *FabricChaincodeCommitStatusApplyConfiguration {
	b.Conditions = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *FabricChaincodeCommitStatusApplyConfiguration) WithMessage(value string) *FabricChaincodeCommitStatusApplyConfiguration {
	b.Message = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *FabricChaincodeCommitStatusApplyConfiguration) WithStatus(value v1alpha1.DeploymentStatus) *FabricChaincodeCommitStatusApplyConfiguration {
	b.Status = &value
	return b
}

// WithTransactionID sets the TransactionID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TransactionID field is set to the value of the last call.
func (b *FabricChaincodeCommitStatusApplyConfiguration) WithTransactionID(value string) *FabricChaincodeCommitStatusApplyConfiguration {
	b.TransactionID = &value
	return b
}