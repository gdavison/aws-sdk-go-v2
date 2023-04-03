// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a product. A delegated admin is authorized to invoke this command. The
// user or role that performs this operation must have the
// cloudformation:GetTemplate IAM policy permission. This policy permission is
// required when using the ImportFromPhysicalId template source in the information
// data section.
func (c *Client) CreateProduct(ctx context.Context, params *CreateProductInput, optFns ...func(*Options)) (*CreateProductOutput, error) {
	if params == nil {
		params = &CreateProductInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateProduct", params, optFns, c.addOperationCreateProductMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateProductOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProductInput struct {

	// A unique identifier that you provide to ensure idempotency. If multiple requests
	// differ only by the idempotency token, the same response is returned for each
	// repeated request.
	//
	// This member is required.
	IdempotencyToken *string

	// The name of the product.
	//
	// This member is required.
	Name *string

	// The owner of the product.
	//
	// This member is required.
	Owner *string

	// The type of product.
	//
	// This member is required.
	ProductType types.ProductType

	// The language code.
	//
	// * jp - Japanese
	//
	// * zh - Chinese
	AcceptLanguage *string

	// The description of the product.
	Description *string

	// The distributor of the product.
	Distributor *string

	// The configuration of the provisioning artifact.
	ProvisioningArtifactParameters *types.ProvisioningArtifactProperties

	// Specifies connection details for the created product and syncs the product to
	// the connection source artifact. This automatically manages the product's
	// artifacts based on changes to the source. The SourceConnection parameter
	// consists of the following sub-fields.
	//
	// * Type
	//
	// * ConnectionParamters
	SourceConnection *types.SourceConnection

	// The support information about the product.
	SupportDescription *string

	// The contact email for product support.
	SupportEmail *string

	// The contact URL for product support. ^https?:\/\// / is the pattern used to
	// validate SupportUrl.
	SupportUrl *string

	// One or more tags.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateProductOutput struct {

	// Information about the product view.
	ProductViewDetail *types.ProductViewDetail

	// Information about the provisioning artifact.
	ProvisioningArtifactDetail *types.ProvisioningArtifactDetail

	// Information about the tags associated with the product.
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateProductMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateProduct{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateProduct{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opCreateProductMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateProductValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProduct(options.Region), middleware.Before); err != nil {
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
	return nil
}

type idempotencyToken_initializeOpCreateProduct struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateProduct) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateProduct) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateProductInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateProductInput ")
	}

	if input.IdempotencyToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.IdempotencyToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateProductMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateProduct{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateProduct(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "CreateProduct",
	}
}
