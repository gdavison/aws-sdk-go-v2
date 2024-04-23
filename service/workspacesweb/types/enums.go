// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AuthenticationType string

// Enum values for AuthenticationType
const (
	AuthenticationTypeStandard          AuthenticationType = "Standard"
	AuthenticationTypeIamIdentityCenter AuthenticationType = "IAM_Identity_Center"
)

// Values returns all known values for AuthenticationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AuthenticationType) Values() []AuthenticationType {
	return []AuthenticationType{
		"Standard",
		"IAM_Identity_Center",
	}
}

type BrowserType string

// Enum values for BrowserType
const (
	BrowserTypeChrome BrowserType = "Chrome"
)

// Values returns all known values for BrowserType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (BrowserType) Values() []BrowserType {
	return []BrowserType{
		"Chrome",
	}
}

type EnabledType string

// Enum values for EnabledType
const (
	EnabledTypeDisabled EnabledType = "Disabled"
	EnabledTypeEnabled  EnabledType = "Enabled"
)

// Values returns all known values for EnabledType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (EnabledType) Values() []EnabledType {
	return []EnabledType{
		"Disabled",
		"Enabled",
	}
}

type IdentityProviderType string

// Enum values for IdentityProviderType
const (
	IdentityProviderTypeSaml            IdentityProviderType = "SAML"
	IdentityProviderTypeFacebook        IdentityProviderType = "Facebook"
	IdentityProviderTypeGoogle          IdentityProviderType = "Google"
	IdentityProviderTypeLoginWithAmazon IdentityProviderType = "LoginWithAmazon"
	IdentityProviderTypeSignInWithApple IdentityProviderType = "SignInWithApple"
	IdentityProviderTypeOidc            IdentityProviderType = "OIDC"
)

// Values returns all known values for IdentityProviderType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (IdentityProviderType) Values() []IdentityProviderType {
	return []IdentityProviderType{
		"SAML",
		"Facebook",
		"Google",
		"LoginWithAmazon",
		"SignInWithApple",
		"OIDC",
	}
}

type InstanceType string

// Enum values for InstanceType
const (
	InstanceTypeStandardRegular InstanceType = "standard.regular"
	InstanceTypeStandardLarge   InstanceType = "standard.large"
	InstanceTypeStandardXlarge  InstanceType = "standard.xlarge"
)

// Values returns all known values for InstanceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (InstanceType) Values() []InstanceType {
	return []InstanceType{
		"standard.regular",
		"standard.large",
		"standard.xlarge",
	}
}

type PortalStatus string

// Enum values for PortalStatus
const (
	PortalStatusIncomplete PortalStatus = "Incomplete"
	PortalStatusPending    PortalStatus = "Pending"
	PortalStatusActive     PortalStatus = "Active"
)

// Values returns all known values for PortalStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PortalStatus) Values() []PortalStatus {
	return []PortalStatus{
		"Incomplete",
		"Pending",
		"Active",
	}
}

type RendererType string

// Enum values for RendererType
const (
	RendererTypeAppstream RendererType = "AppStream"
)

// Values returns all known values for RendererType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RendererType) Values() []RendererType {
	return []RendererType{
		"AppStream",
	}
}

type ValidationExceptionReason string

// Enum values for ValidationExceptionReason
const (
	ValidationExceptionReasonUnknownOperation      ValidationExceptionReason = "unknownOperation"
	ValidationExceptionReasonCannotParse           ValidationExceptionReason = "cannotParse"
	ValidationExceptionReasonFieldValidationFailed ValidationExceptionReason = "fieldValidationFailed"
	ValidationExceptionReasonOther                 ValidationExceptionReason = "other"
)

// Values returns all known values for ValidationExceptionReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ValidationExceptionReason) Values() []ValidationExceptionReason {
	return []ValidationExceptionReason{
		"unknownOperation",
		"cannotParse",
		"fieldValidationFailed",
		"other",
	}
}
