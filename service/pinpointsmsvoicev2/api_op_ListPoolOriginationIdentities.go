// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all associated origination identities in your pool. If you specify
// filters, the output includes information for only those origination identities
// that meet the filter criteria.
func (c *Client) ListPoolOriginationIdentities(ctx context.Context, params *ListPoolOriginationIdentitiesInput, optFns ...func(*Options)) (*ListPoolOriginationIdentitiesOutput, error) {
	if params == nil {
		params = &ListPoolOriginationIdentitiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPoolOriginationIdentities", params, optFns, c.addOperationListPoolOriginationIdentitiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPoolOriginationIdentitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPoolOriginationIdentitiesInput struct {

	// The unique identifier for the pool. This value can be either the PoolId or
	// PoolArn.
	//
	// This member is required.
	PoolId *string

	// An array of PoolOriginationIdentitiesFilter objects to filter the results..
	Filters []types.PoolOriginationIdentitiesFilter

	// The maximum number of results to return per each request.
	MaxResults *int32

	// The token to be used for the next set of paginated results. You don't need to
	// supply a value for this field in the initial request.
	NextToken *string

	noSmithyDocumentSerde
}

type ListPoolOriginationIdentitiesOutput struct {

	// The token to be used for the next set of paginated results. If this field is
	// empty then there are no more results.
	NextToken *string

	// An array of any OriginationIdentityMetadata objects.
	OriginationIdentities []types.OriginationIdentityMetadata

	// The Amazon Resource Name (ARN) for the pool.
	PoolArn *string

	// The unique PoolId of the pool.
	PoolId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPoolOriginationIdentitiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListPoolOriginationIdentities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListPoolOriginationIdentities{}, middleware.After)
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
	if err = addOpListPoolOriginationIdentitiesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPoolOriginationIdentities(options.Region), middleware.Before); err != nil {
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

// ListPoolOriginationIdentitiesAPIClient is a client that implements the
// ListPoolOriginationIdentities operation.
type ListPoolOriginationIdentitiesAPIClient interface {
	ListPoolOriginationIdentities(context.Context, *ListPoolOriginationIdentitiesInput, ...func(*Options)) (*ListPoolOriginationIdentitiesOutput, error)
}

var _ ListPoolOriginationIdentitiesAPIClient = (*Client)(nil)

// ListPoolOriginationIdentitiesPaginatorOptions is the paginator options for
// ListPoolOriginationIdentities
type ListPoolOriginationIdentitiesPaginatorOptions struct {
	// The maximum number of results to return per each request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPoolOriginationIdentitiesPaginator is a paginator for
// ListPoolOriginationIdentities
type ListPoolOriginationIdentitiesPaginator struct {
	options   ListPoolOriginationIdentitiesPaginatorOptions
	client    ListPoolOriginationIdentitiesAPIClient
	params    *ListPoolOriginationIdentitiesInput
	nextToken *string
	firstPage bool
}

// NewListPoolOriginationIdentitiesPaginator returns a new
// ListPoolOriginationIdentitiesPaginator
func NewListPoolOriginationIdentitiesPaginator(client ListPoolOriginationIdentitiesAPIClient, params *ListPoolOriginationIdentitiesInput, optFns ...func(*ListPoolOriginationIdentitiesPaginatorOptions)) *ListPoolOriginationIdentitiesPaginator {
	if params == nil {
		params = &ListPoolOriginationIdentitiesInput{}
	}

	options := ListPoolOriginationIdentitiesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPoolOriginationIdentitiesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPoolOriginationIdentitiesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPoolOriginationIdentities page.
func (p *ListPoolOriginationIdentitiesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPoolOriginationIdentitiesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListPoolOriginationIdentities(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListPoolOriginationIdentities(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms-voice",
		OperationName: "ListPoolOriginationIdentities",
	}
}