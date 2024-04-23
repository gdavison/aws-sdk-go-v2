// Code generated by smithy-go-codegen DO NOT EDIT.

package workspacesweb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/workspacesweb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a user settings resource that can be associated with a web portal. Once
// associated with a web portal, user settings control how users can transfer data
// between a streaming session and the their local devices.
func (c *Client) CreateUserSettings(ctx context.Context, params *CreateUserSettingsInput, optFns ...func(*Options)) (*CreateUserSettingsOutput, error) {
	if params == nil {
		params = &CreateUserSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateUserSettings", params, optFns, c.addOperationCreateUserSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateUserSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateUserSettingsInput struct {

	// Specifies whether the user can copy text from the streaming session to the
	// local device.
	//
	// This member is required.
	CopyAllowed types.EnabledType

	// Specifies whether the user can download files from the streaming session to the
	// local device.
	//
	// This member is required.
	DownloadAllowed types.EnabledType

	// Specifies whether the user can paste text from the local device to the
	// streaming session.
	//
	// This member is required.
	PasteAllowed types.EnabledType

	// Specifies whether the user can print to the local device.
	//
	// This member is required.
	PrintAllowed types.EnabledType

	// Specifies whether the user can upload files from the local device to the
	// streaming session.
	//
	// This member is required.
	UploadAllowed types.EnabledType

	// The additional encryption context of the user settings.
	AdditionalEncryptionContext map[string]string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. Idempotency ensures that an API request completes only once.
	// With an idempotent request, if the original request completes successfully,
	// subsequent retries with the same client token returns the result from the
	// original successful request. If you do not specify a client token, one is
	// automatically generated by the Amazon Web Services SDK.
	ClientToken *string

	// The configuration that specifies which cookies should be synchronized from the
	// end user's local browser to the remote browser.
	CookieSynchronizationConfiguration *types.CookieSynchronizationConfiguration

	// The customer managed key used to encrypt sensitive information in the user
	// settings.
	CustomerManagedKey *string

	// The amount of time that a streaming session remains active after users
	// disconnect.
	DisconnectTimeoutInMinutes *int32

	// The amount of time that users can be idle (inactive) before they are
	// disconnected from their streaming session and the disconnect timeout interval
	// begins.
	IdleDisconnectTimeoutInMinutes *int32

	// The tags to add to the user settings resource. A tag is a key-value pair.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateUserSettingsOutput struct {

	// The ARN of the user settings.
	//
	// This member is required.
	UserSettingsArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateUserSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateUserSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateUserSettings{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateUserSettings"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opCreateUserSettingsMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateUserSettingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateUserSettings(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

type idempotencyToken_initializeOpCreateUserSettings struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateUserSettings) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateUserSettings) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateUserSettingsInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateUserSettingsInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateUserSettingsMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateUserSettings{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateUserSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateUserSettings",
	}
}
