// Code generated by smithy-go-codegen DO NOT EDIT.

package backupgateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disassociates a backup gateway from the specified server. After the
// disassociation process finishes, the gateway can no longer access the virtual
// machines on the server.
func (c *Client) DisassociateGatewayFromServer(ctx context.Context, params *DisassociateGatewayFromServerInput, optFns ...func(*Options)) (*DisassociateGatewayFromServerOutput, error) {
	if params == nil {
		params = &DisassociateGatewayFromServerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateGatewayFromServer", params, optFns, c.addOperationDisassociateGatewayFromServerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateGatewayFromServerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateGatewayFromServerInput struct {

	// The Amazon Resource Name (ARN) of the gateway to disassociate.
	//
	// This member is required.
	GatewayArn *string

	noSmithyDocumentSerde
}

type DisassociateGatewayFromServerOutput struct {

	// The Amazon Resource Name (ARN) of the gateway you disassociated.
	GatewayArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateGatewayFromServerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDisassociateGatewayFromServer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDisassociateGatewayFromServer{}, middleware.After)
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
	if err = addOpDisassociateGatewayFromServerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateGatewayFromServer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateGatewayFromServer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup-gateway",
		OperationName: "DisassociateGatewayFromServer",
	}
}