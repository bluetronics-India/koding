// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package kmsiface provides an interface for the AWS Key Management Service.
package kmsiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kms"
)

// KMSAPI is the interface type for kms.KMS.
type KMSAPI interface {
	CreateAliasRequest(*kms.CreateAliasInput) (*aws.Request, *kms.CreateAliasOutput)

	CreateAlias(*kms.CreateAliasInput) (*kms.CreateAliasOutput, error)

	CreateGrantRequest(*kms.CreateGrantInput) (*aws.Request, *kms.CreateGrantOutput)

	CreateGrant(*kms.CreateGrantInput) (*kms.CreateGrantOutput, error)

	CreateKeyRequest(*kms.CreateKeyInput) (*aws.Request, *kms.CreateKeyOutput)

	CreateKey(*kms.CreateKeyInput) (*kms.CreateKeyOutput, error)

	DecryptRequest(*kms.DecryptInput) (*aws.Request, *kms.DecryptOutput)

	Decrypt(*kms.DecryptInput) (*kms.DecryptOutput, error)

	DeleteAliasRequest(*kms.DeleteAliasInput) (*aws.Request, *kms.DeleteAliasOutput)

	DeleteAlias(*kms.DeleteAliasInput) (*kms.DeleteAliasOutput, error)

	DescribeKeyRequest(*kms.DescribeKeyInput) (*aws.Request, *kms.DescribeKeyOutput)

	DescribeKey(*kms.DescribeKeyInput) (*kms.DescribeKeyOutput, error)

	DisableKeyRequest(*kms.DisableKeyInput) (*aws.Request, *kms.DisableKeyOutput)

	DisableKey(*kms.DisableKeyInput) (*kms.DisableKeyOutput, error)

	DisableKeyRotationRequest(*kms.DisableKeyRotationInput) (*aws.Request, *kms.DisableKeyRotationOutput)

	DisableKeyRotation(*kms.DisableKeyRotationInput) (*kms.DisableKeyRotationOutput, error)

	EnableKeyRequest(*kms.EnableKeyInput) (*aws.Request, *kms.EnableKeyOutput)

	EnableKey(*kms.EnableKeyInput) (*kms.EnableKeyOutput, error)

	EnableKeyRotationRequest(*kms.EnableKeyRotationInput) (*aws.Request, *kms.EnableKeyRotationOutput)

	EnableKeyRotation(*kms.EnableKeyRotationInput) (*kms.EnableKeyRotationOutput, error)

	EncryptRequest(*kms.EncryptInput) (*aws.Request, *kms.EncryptOutput)

	Encrypt(*kms.EncryptInput) (*kms.EncryptOutput, error)

	GenerateDataKeyRequest(*kms.GenerateDataKeyInput) (*aws.Request, *kms.GenerateDataKeyOutput)

	GenerateDataKey(*kms.GenerateDataKeyInput) (*kms.GenerateDataKeyOutput, error)

	GenerateDataKeyWithoutPlaintextRequest(*kms.GenerateDataKeyWithoutPlaintextInput) (*aws.Request, *kms.GenerateDataKeyWithoutPlaintextOutput)

	GenerateDataKeyWithoutPlaintext(*kms.GenerateDataKeyWithoutPlaintextInput) (*kms.GenerateDataKeyWithoutPlaintextOutput, error)

	GenerateRandomRequest(*kms.GenerateRandomInput) (*aws.Request, *kms.GenerateRandomOutput)

	GenerateRandom(*kms.GenerateRandomInput) (*kms.GenerateRandomOutput, error)

	GetKeyPolicyRequest(*kms.GetKeyPolicyInput) (*aws.Request, *kms.GetKeyPolicyOutput)

	GetKeyPolicy(*kms.GetKeyPolicyInput) (*kms.GetKeyPolicyOutput, error)

	GetKeyRotationStatusRequest(*kms.GetKeyRotationStatusInput) (*aws.Request, *kms.GetKeyRotationStatusOutput)

	GetKeyRotationStatus(*kms.GetKeyRotationStatusInput) (*kms.GetKeyRotationStatusOutput, error)

	ListAliasesRequest(*kms.ListAliasesInput) (*aws.Request, *kms.ListAliasesOutput)

	ListAliases(*kms.ListAliasesInput) (*kms.ListAliasesOutput, error)

	ListAliasesPages(*kms.ListAliasesInput, func(*kms.ListAliasesOutput, bool) bool) error

	ListGrantsRequest(*kms.ListGrantsInput) (*aws.Request, *kms.ListGrantsOutput)

	ListGrants(*kms.ListGrantsInput) (*kms.ListGrantsOutput, error)

	ListGrantsPages(*kms.ListGrantsInput, func(*kms.ListGrantsOutput, bool) bool) error

	ListKeyPoliciesRequest(*kms.ListKeyPoliciesInput) (*aws.Request, *kms.ListKeyPoliciesOutput)

	ListKeyPolicies(*kms.ListKeyPoliciesInput) (*kms.ListKeyPoliciesOutput, error)

	ListKeyPoliciesPages(*kms.ListKeyPoliciesInput, func(*kms.ListKeyPoliciesOutput, bool) bool) error

	ListKeysRequest(*kms.ListKeysInput) (*aws.Request, *kms.ListKeysOutput)

	ListKeys(*kms.ListKeysInput) (*kms.ListKeysOutput, error)

	ListKeysPages(*kms.ListKeysInput, func(*kms.ListKeysOutput, bool) bool) error

	PutKeyPolicyRequest(*kms.PutKeyPolicyInput) (*aws.Request, *kms.PutKeyPolicyOutput)

	PutKeyPolicy(*kms.PutKeyPolicyInput) (*kms.PutKeyPolicyOutput, error)

	ReEncryptRequest(*kms.ReEncryptInput) (*aws.Request, *kms.ReEncryptOutput)

	ReEncrypt(*kms.ReEncryptInput) (*kms.ReEncryptOutput, error)

	RetireGrantRequest(*kms.RetireGrantInput) (*aws.Request, *kms.RetireGrantOutput)

	RetireGrant(*kms.RetireGrantInput) (*kms.RetireGrantOutput, error)

	RevokeGrantRequest(*kms.RevokeGrantInput) (*aws.Request, *kms.RevokeGrantOutput)

	RevokeGrant(*kms.RevokeGrantInput) (*kms.RevokeGrantOutput, error)

	UpdateAliasRequest(*kms.UpdateAliasInput) (*aws.Request, *kms.UpdateAliasOutput)

	UpdateAlias(*kms.UpdateAliasInput) (*kms.UpdateAliasOutput, error)

	UpdateKeyDescriptionRequest(*kms.UpdateKeyDescriptionInput) (*aws.Request, *kms.UpdateKeyDescriptionOutput)

	UpdateKeyDescription(*kms.UpdateKeyDescriptionInput) (*kms.UpdateKeyDescriptionOutput, error)
}