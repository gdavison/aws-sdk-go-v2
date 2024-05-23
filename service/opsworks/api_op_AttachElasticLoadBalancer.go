// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Attaches an Elastic Load Balancing load balancer to a specified layer. OpsWorks
// Stacks does not support Application Load Balancer. You can only use Classic Load
// Balancer with OpsWorks Stacks. For more information, see [Elastic Load Balancing].
//
// You must create the Elastic Load Balancing instance separately, by using the
// Elastic Load Balancing console, API, or CLI. For more information, see the [Elastic Load Balancing Developer Guide].
//
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack, or an attached policy that explicitly grants
// permissions. For more information on user permissions, see [Managing User Permissions].
//
// [Elastic Load Balancing]: https://docs.aws.amazon.com/opsworks/latest/userguide/layers-elb.html
// [Elastic Load Balancing Developer Guide]: https://docs.aws.amazon.com/ElasticLoadBalancing/latest/DeveloperGuide/Welcome.html
// [Managing User Permissions]: https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html
func (c *Client) AttachElasticLoadBalancer(ctx context.Context, params *AttachElasticLoadBalancerInput, optFns ...func(*Options)) (*AttachElasticLoadBalancerOutput, error) {
	if params == nil {
		params = &AttachElasticLoadBalancerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AttachElasticLoadBalancer", params, optFns, c.addOperationAttachElasticLoadBalancerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AttachElasticLoadBalancerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AttachElasticLoadBalancerInput struct {

	// The Elastic Load Balancing instance's name.
	//
	// This member is required.
	ElasticLoadBalancerName *string

	// The ID of the layer to which the Elastic Load Balancing instance is to be
	// attached.
	//
	// This member is required.
	LayerId *string

	noSmithyDocumentSerde
}

type AttachElasticLoadBalancerOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAttachElasticLoadBalancerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAttachElasticLoadBalancer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAttachElasticLoadBalancer{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AttachElasticLoadBalancer"); err != nil {
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
	if err = addOpAttachElasticLoadBalancerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAttachElasticLoadBalancer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAttachElasticLoadBalancer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AttachElasticLoadBalancer",
	}
}
