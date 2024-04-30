// Code generated by smithy-go-codegen DO NOT EDIT.

package omics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/omics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the resource shares associated with an account. Use the filter
// parameter to retrieve a specific subset of the shares.
func (c *Client) ListShares(ctx context.Context, params *ListSharesInput, optFns ...func(*Options)) (*ListSharesOutput, error) {
	if params == nil {
		params = &ListSharesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListShares", params, optFns, c.addOperationListSharesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSharesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSharesInput struct {

	// The account that owns the resource shares.
	//
	// This member is required.
	ResourceOwner types.ResourceOwner

	// Attributes that you use to filter for a specific subset of resource shares.
	Filter *types.Filter

	// The maximum number of shares to return in one page of results.
	MaxResults *int32

	// Next token returned in the response of a previous ListReadSetUploadPartsRequest
	// call. Used to get the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSharesOutput struct {

	// The shares available and their metadata details.
	//
	// This member is required.
	Shares []types.ShareDetails

	// Next token returned in the response of a previous ListSharesResponse call. Used
	// to get the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSharesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListShares{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListShares{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListShares"); err != nil {
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
	if err = addEndpointPrefix_opListSharesMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListSharesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListShares(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListSharesMiddleware struct {
}

func (*endpointPrefix_opListSharesMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListSharesMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "analytics-" + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opListSharesMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opListSharesMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// ListSharesAPIClient is a client that implements the ListShares operation.
type ListSharesAPIClient interface {
	ListShares(context.Context, *ListSharesInput, ...func(*Options)) (*ListSharesOutput, error)
}

var _ ListSharesAPIClient = (*Client)(nil)

// ListSharesPaginatorOptions is the paginator options for ListShares
type ListSharesPaginatorOptions struct {
	// The maximum number of shares to return in one page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSharesPaginator is a paginator for ListShares
type ListSharesPaginator struct {
	options   ListSharesPaginatorOptions
	client    ListSharesAPIClient
	params    *ListSharesInput
	nextToken *string
	firstPage bool
}

// NewListSharesPaginator returns a new ListSharesPaginator
func NewListSharesPaginator(client ListSharesAPIClient, params *ListSharesInput, optFns ...func(*ListSharesPaginatorOptions)) *ListSharesPaginator {
	if params == nil {
		params = &ListSharesInput{}
	}

	options := ListSharesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSharesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSharesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListShares page.
func (p *ListSharesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSharesOutput, error) {
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

	result, err := p.client.ListShares(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListShares(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListShares",
	}
}
