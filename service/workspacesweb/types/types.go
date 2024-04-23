// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// The browser settings resource that can be associated with a web portal. Once
// associated with a web portal, browser settings control how the browser will
// behave once a user starts a streaming session for the web portal.
type BrowserSettings struct {

	// The ARN of the browser settings.
	//
	// This member is required.
	BrowserSettingsArn *string

	// The additional encryption context of the browser settings.
	AdditionalEncryptionContext map[string]string

	// A list of web portal ARNs that this browser settings is associated with.
	AssociatedPortalArns []string

	// A JSON string containing Chrome Enterprise policies that will be applied to all
	// streaming sessions.
	BrowserPolicy *string

	// The customer managed key used to encrypt sensitive information in the browser
	// settings.
	CustomerManagedKey *string

	noSmithyDocumentSerde
}

// The summary for browser settings.
type BrowserSettingsSummary struct {

	// The ARN of the browser settings.
	//
	// This member is required.
	BrowserSettingsArn *string

	noSmithyDocumentSerde
}

// The certificate.
type Certificate struct {

	// The body of the certificate.
	Body []byte

	// The entity that issued the certificate.
	Issuer *string

	// The certificate is not valid after this date.
	NotValidAfter *time.Time

	// The certificate is not valid before this date.
	NotValidBefore *time.Time

	// The entity the certificate belongs to.
	Subject *string

	// A hexadecimal identifier for the certificate.
	Thumbprint *string

	noSmithyDocumentSerde
}

// The summary of the certificate.
type CertificateSummary struct {

	// The entity that issued the certificate.
	Issuer *string

	// The certificate is not valid after this date.
	NotValidAfter *time.Time

	// The certificate is not valid before this date.
	NotValidBefore *time.Time

	// The entity the certificate belongs to.
	Subject *string

	// A hexadecimal identifier for the certificate.
	Thumbprint *string

	noSmithyDocumentSerde
}

// Specifies a single cookie or set of cookies in an end user's browser.
type CookieSpecification struct {

	// The domain of the cookie.
	//
	// This member is required.
	Domain *string

	// The name of the cookie.
	Name *string

	// The path of the cookie.
	Path *string

	noSmithyDocumentSerde
}

// The configuration that specifies which cookies should be synchronized from the
// end user's local browser to the remote browser.
type CookieSynchronizationConfiguration struct {

	// The list of cookie specifications that are allowed to be synchronized to the
	// remote browser.
	//
	// This member is required.
	Allowlist []CookieSpecification

	// The list of cookie specifications that are blocked from being synchronized to
	// the remote browser.
	Blocklist []CookieSpecification

	noSmithyDocumentSerde
}

// The identity provider.
type IdentityProvider struct {

	// The ARN of the identity provider.
	//
	// This member is required.
	IdentityProviderArn *string

	// The identity provider details. The following list describes the provider detail
	// keys for each identity provider type.
	//   - For Google and Login with Amazon:
	//   - client_id
	//   - client_secret
	//   - authorize_scopes
	//   - For Facebook:
	//   - client_id
	//   - client_secret
	//   - authorize_scopes
	//   - api_version
	//   - For Sign in with Apple:
	//   - client_id
	//   - team_id
	//   - key_id
	//   - private_key
	//   - authorize_scopes
	//   - For OIDC providers:
	//   - client_id
	//   - client_secret
	//   - attributes_request_method
	//   - oidc_issuer
	//   - authorize_scopes
	//   - authorize_url if not available from discovery URL specified by oidc_issuer
	//   key
	//   - token_url if not available from discovery URL specified by oidc_issuer key
	//   - attributes_url if not available from discovery URL specified by oidc_issuer
	//   key
	//   - jwks_uri if not available from discovery URL specified by oidc_issuer key
	//   - For SAML providers:
	//   - MetadataFile OR MetadataURL
	//   - IDPSignout (boolean) optional
	//   - IDPInit (boolean) optional
	//   - RequestSigningAlgorithm (string) optional - Only accepts rsa-sha256
	//   - EncryptedResponses (boolean) optional
	IdentityProviderDetails map[string]string

	// The identity provider name.
	IdentityProviderName *string

	// The identity provider type.
	IdentityProviderType IdentityProviderType

	noSmithyDocumentSerde
}

// The summary of the identity provider.
type IdentityProviderSummary struct {

	// The ARN of the identity provider.
	//
	// This member is required.
	IdentityProviderArn *string

	// The identity provider name.
	IdentityProviderName *string

	// The identity provider type.
	IdentityProviderType IdentityProviderType

	noSmithyDocumentSerde
}

// The IP access settings resource that can be associated with a web portal.
type IpAccessSettings struct {

	// The ARN of the IP access settings resource.
	//
	// This member is required.
	IpAccessSettingsArn *string

	// The additional encryption context of the IP access settings.
	AdditionalEncryptionContext map[string]string

	// A list of web portal ARNs that this IP access settings resource is associated
	// with.
	AssociatedPortalArns []string

	// The creation date timestamp of the IP access settings.
	CreationDate *time.Time

	// The customer managed key used to encrypt sensitive information in the IP access
	// settings.
	CustomerManagedKey *string

	// The description of the IP access settings.
	Description *string

	// The display name of the IP access settings.
	DisplayName *string

	// The IP rules of the IP access settings.
	IpRules []IpRule

	noSmithyDocumentSerde
}

// The summary of IP access settings.
type IpAccessSettingsSummary struct {

	// The ARN of IP access settings.
	//
	// This member is required.
	IpAccessSettingsArn *string

	// The creation date timestamp of the IP access settings.
	CreationDate *time.Time

	// The description of the IP access settings.
	Description *string

	// The display name of the IP access settings.
	DisplayName *string

	noSmithyDocumentSerde
}

// The IP rules of the IP access settings.
type IpRule struct {

	// The IP range of the IP rule.
	//
	// This member is required.
	IpRange *string

	// The description of the IP rule.
	Description *string

	noSmithyDocumentSerde
}

// A network settings resource that can be associated with a web portal. Once
// associated with a web portal, network settings define how streaming instances
// will connect with your specified VPC.
type NetworkSettings struct {

	// The ARN of the network settings.
	//
	// This member is required.
	NetworkSettingsArn *string

	// A list of web portal ARNs that this network settings is associated with.
	AssociatedPortalArns []string

	// One or more security groups used to control access from streaming instances to
	// your VPC.
	SecurityGroupIds []string

	// The subnets in which network interfaces are created to connect streaming
	// instances to your VPC. At least two of these subnets must be in different
	// availability zones.
	SubnetIds []string

	// The VPC that streaming instances will connect to.
	VpcId *string

	noSmithyDocumentSerde
}

// The summary of network settings.
type NetworkSettingsSummary struct {

	// The ARN of the network settings.
	//
	// This member is required.
	NetworkSettingsArn *string

	// The VPC ID of the network settings.
	VpcId *string

	noSmithyDocumentSerde
}

// The web portal.
type Portal struct {

	// The ARN of the web portal.
	//
	// This member is required.
	PortalArn *string

	// The additional encryption context of the portal.
	AdditionalEncryptionContext map[string]string

	// The type of authentication integration points used when signing into the web
	// portal. Defaults to Standard . Standard web portals are authenticated directly
	// through your identity provider. You need to call CreateIdentityProvider to
	// integrate your identity provider with your web portal. User and group access to
	// your web portal is controlled through your identity provider. IAM Identity
	// Center web portals are authenticated through IAM Identity Center (successor to
	// Single Sign-On). Identity sources (including external identity provider
	// integration), plus user and group access to your web portal, can be configured
	// in the IAM Identity Center.
	AuthenticationType AuthenticationType

	// The ARN of the browser settings that is associated with this web portal.
	BrowserSettingsArn *string

	// The browser that users see when using a streaming session.
	BrowserType BrowserType

	// The creation date of the web portal.
	CreationDate *time.Time

	// The customer managed key used to encrypt sensitive information in the portal.
	CustomerManagedKey *string

	// The name of the web portal.
	DisplayName *string

	// The type and resources of the underlying instance.
	InstanceType InstanceType

	// The ARN of the IP access settings.
	IpAccessSettingsArn *string

	// The maximum number of concurrent sessions for the portal.
	MaxConcurrentSessions *int32

	// The ARN of the network settings that is associated with the web portal.
	NetworkSettingsArn *string

	// The endpoint URL of the web portal that users access in order to start
	// streaming sessions.
	PortalEndpoint *string

	// The status of the web portal.
	PortalStatus PortalStatus

	// The renderer that is used in streaming sessions.
	RendererType RendererType

	// A message that explains why the web portal is in its current status.
	StatusReason *string

	// The ARN of the trust store that is associated with the web portal.
	TrustStoreArn *string

	// The ARN of the user access logging settings that is associated with the web
	// portal.
	UserAccessLoggingSettingsArn *string

	// The ARN of the user settings that is associated with the web portal.
	UserSettingsArn *string

	noSmithyDocumentSerde
}

// The summary of the portal.
type PortalSummary struct {

	// The ARN of the web portal.
	//
	// This member is required.
	PortalArn *string

	// The type of authentication integration points used when signing into the web
	// portal. Defaults to Standard . Standard web portals are authenticated directly
	// through your identity provider. You need to call CreateIdentityProvider to
	// integrate your identity provider with your web portal. User and group access to
	// your web portal is controlled through your identity provider. IAM Identity
	// Center web portals are authenticated through IAM Identity Center (successor to
	// Single Sign-On). Identity sources (including external identity provider
	// integration), plus user and group access to your web portal, can be configured
	// in the IAM Identity Center.
	AuthenticationType AuthenticationType

	// The ARN of the browser settings that is associated with the web portal.
	BrowserSettingsArn *string

	// The browser type of the web portal.
	BrowserType BrowserType

	// The creation date of the web portal.
	CreationDate *time.Time

	// The name of the web portal.
	DisplayName *string

	// The type and resources of the underlying instance.
	InstanceType InstanceType

	// The ARN of the IP access settings.
	IpAccessSettingsArn *string

	// The maximum number of concurrent sessions for the portal.
	MaxConcurrentSessions *int32

	// The ARN of the network settings that is associated with the web portal.
	NetworkSettingsArn *string

	// The endpoint URL of the web portal that users access in order to start
	// streaming sessions.
	PortalEndpoint *string

	// The status of the web portal.
	PortalStatus PortalStatus

	// The renderer that is used in streaming sessions.
	RendererType RendererType

	// The ARN of the trust that is associated with this web portal.
	TrustStoreArn *string

	// The ARN of the user access logging settings that is associated with the web
	// portal.
	UserAccessLoggingSettingsArn *string

	// The ARN of the user settings that is associated with the web portal.
	UserSettingsArn *string

	noSmithyDocumentSerde
}

// The tag.
type Tag struct {

	// The key of the tag.
	//
	// This member is required.
	Key *string

	// The value of the tag
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// A trust store that can be associated with a web portal. A trust store contains
// certificate authority (CA) certificates. Once associated with a web portal, the
// browser in a streaming session will recognize certificates that have been issued
// using any of the CAs in the trust store. If your organization has internal
// websites that use certificates issued by private CAs, you should add the private
// CA certificate to the trust store.
type TrustStore struct {

	// The ARN of the trust store.
	//
	// This member is required.
	TrustStoreArn *string

	// A list of web portal ARNs that this trust store is associated with.
	AssociatedPortalArns []string

	noSmithyDocumentSerde
}

// The summary of the trust store.
type TrustStoreSummary struct {

	// The ARN of the trust store.
	TrustStoreArn *string

	noSmithyDocumentSerde
}

// A user access logging settings resource that can be associated with a web
// portal.
type UserAccessLoggingSettings struct {

	// The ARN of the user access logging settings.
	//
	// This member is required.
	UserAccessLoggingSettingsArn *string

	// A list of web portal ARNs that this user access logging settings is associated
	// with.
	AssociatedPortalArns []string

	// The ARN of the Kinesis stream.
	KinesisStreamArn *string

	noSmithyDocumentSerde
}

// The summary of user access logging settings.
type UserAccessLoggingSettingsSummary struct {

	// The ARN of the user access logging settings.
	//
	// This member is required.
	UserAccessLoggingSettingsArn *string

	// The ARN of the Kinesis stream.
	KinesisStreamArn *string

	noSmithyDocumentSerde
}

// A user settings resource that can be associated with a web portal. Once
// associated with a web portal, user settings control how users can transfer data
// between a streaming session and the their local devices.
type UserSettings struct {

	// The ARN of the user settings.
	//
	// This member is required.
	UserSettingsArn *string

	// The additional encryption context of the user settings.
	AdditionalEncryptionContext map[string]string

	// A list of web portal ARNs that this user settings is associated with.
	AssociatedPortalArns []string

	// The configuration that specifies which cookies should be synchronized from the
	// end user's local browser to the remote browser.
	CookieSynchronizationConfiguration *CookieSynchronizationConfiguration

	// Specifies whether the user can copy text from the streaming session to the
	// local device.
	CopyAllowed EnabledType

	// The customer managed key used to encrypt sensitive information in the user
	// settings.
	CustomerManagedKey *string

	// The amount of time that a streaming session remains active after users
	// disconnect.
	DisconnectTimeoutInMinutes *int32

	// Specifies whether the user can download files from the streaming session to the
	// local device.
	DownloadAllowed EnabledType

	// The amount of time that users can be idle (inactive) before they are
	// disconnected from their streaming session and the disconnect timeout interval
	// begins.
	IdleDisconnectTimeoutInMinutes *int32

	// Specifies whether the user can paste text from the local device to the
	// streaming session.
	PasteAllowed EnabledType

	// Specifies whether the user can print to the local device.
	PrintAllowed EnabledType

	// Specifies whether the user can upload files from the local device to the
	// streaming session.
	UploadAllowed EnabledType

	noSmithyDocumentSerde
}

// The summary of user settings.
type UserSettingsSummary struct {

	// The ARN of the user settings.
	//
	// This member is required.
	UserSettingsArn *string

	// The configuration that specifies which cookies should be synchronized from the
	// end user's local browser to the remote browser.
	CookieSynchronizationConfiguration *CookieSynchronizationConfiguration

	// Specifies whether the user can copy text from the streaming session to the
	// local device.
	CopyAllowed EnabledType

	// The amount of time that a streaming session remains active after users
	// disconnect.
	DisconnectTimeoutInMinutes *int32

	// Specifies whether the user can download files from the streaming session to the
	// local device.
	DownloadAllowed EnabledType

	// The amount of time that users can be idle (inactive) before they are
	// disconnected from their streaming session and the disconnect timeout interval
	// begins.
	IdleDisconnectTimeoutInMinutes *int32

	// Specifies whether the user can paste text from the local device to the
	// streaming session.
	PasteAllowed EnabledType

	// Specifies whether the user can print to the local device.
	PrintAllowed EnabledType

	// Specifies whether the user can upload files from the local device to the
	// streaming session.
	UploadAllowed EnabledType

	noSmithyDocumentSerde
}

// Information about a field passed inside a request that resulted in an exception.
type ValidationExceptionField struct {

	// The message describing why the field failed validation.
	//
	// This member is required.
	Message *string

	// The name of the field that failed validation.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
