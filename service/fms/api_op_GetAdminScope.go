// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/fms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about the specified account's administrative scope. The
// administrative scope defines the resources that an Firewall Manager
// administrator can manage.
func (c *Client) GetAdminScope(ctx context.Context, params *GetAdminScopeInput, optFns ...func(*Options)) (*GetAdminScopeOutput, error) {
	if params == nil {
		params = &GetAdminScopeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAdminScope", params, optFns, c.addOperationGetAdminScopeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAdminScopeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAdminScopeInput struct {

	// The administrator account that you want to get the details for.
	//
	// This member is required.
	AdminAccount *string

	noSmithyDocumentSerde
}

type GetAdminScopeOutput struct {

	// Contains details about the administrative scope of the requested account.
	AdminScope *types.AdminScope

	// The current status of the request to onboard a member account as an Firewall
	// Manager administrator.
	//   - ONBOARDING - The account is onboarding to Firewall Manager as an
	//   administrator.
	//   - ONBOARDING_COMPLETE - Firewall Manager The account is onboarded to Firewall
	//   Manager as an administrator, and can perform actions on the resources defined in
	//   their AdminScope .
	//   - OFFBOARDING - The account is being removed as an Firewall Manager
	//   administrator.
	//   - OFFBOARDING_COMPLETE - The account has been removed as an Firewall Manager
	//   administrator.
	Status types.OrganizationStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAdminScopeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetAdminScope{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetAdminScope{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetAdminScope"); err != nil {
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
	if err = addOpGetAdminScopeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAdminScope(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetAdminScope(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetAdminScope",
	}
}
