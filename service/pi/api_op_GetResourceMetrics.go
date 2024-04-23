// Code generated by smithy-go-codegen DO NOT EDIT.

package pi

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/pi/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieve Performance Insights metrics for a set of data sources over a time
// period. You can provide specific dimension groups and dimensions, and provide
// filtering criteria for each group. You must specify an aggregate function for
// each metric. Each response element returns a maximum of 500 bytes. For larger
// elements, such as SQL statements, only the first 500 bytes are returned.
func (c *Client) GetResourceMetrics(ctx context.Context, params *GetResourceMetricsInput, optFns ...func(*Options)) (*GetResourceMetricsOutput, error) {
	if params == nil {
		params = &GetResourceMetricsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetResourceMetrics", params, optFns, c.addOperationGetResourceMetricsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetResourceMetricsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetResourceMetricsInput struct {

	// The date and time specifying the end of the requested time series query range.
	// The value specified is exclusive. Thus, the command returns data points less
	// than (but not equal to) EndTime . The value for EndTime must be later than the
	// value for StartTime .
	//
	// This member is required.
	EndTime *time.Time

	// An immutable identifier for a data source that is unique for an Amazon Web
	// Services Region. Performance Insights gathers metrics from this data source. In
	// the console, the identifier is shown as ResourceID. When you call
	// DescribeDBInstances , the identifier is returned as DbiResourceId . To use a DB
	// instance as a data source, specify its DbiResourceId value. For example,
	// specify db-ABCDEFGHIJKLMNOPQRSTU1VW2X .
	//
	// This member is required.
	Identifier *string

	// An array of one or more queries to perform. Each query must specify a
	// Performance Insights metric and specify an aggregate function, and you can
	// provide filtering criteria. You must append the aggregate function to the
	// metric. For example, to find the average for the metric db.load you must use
	// db.load.avg . Valid values for aggregate functions include .avg , .min , .max ,
	// and .sum .
	//
	// This member is required.
	MetricQueries []types.MetricQuery

	// The Amazon Web Services service for which Performance Insights returns metrics.
	// Valid values are as follows:
	//   - RDS
	//   - DOCDB
	//
	// This member is required.
	ServiceType types.ServiceType

	// The date and time specifying the beginning of the requested time series query
	// range. You can't specify a StartTime that is earlier than 7 days ago. By
	// default, Performance Insights has 7 days of retention, but you can extend this
	// range up to 2 years. The value specified is inclusive. Thus, the command returns
	// data points equal to or greater than StartTime . The value for StartTime must
	// be earlier than the value for EndTime .
	//
	// This member is required.
	StartTime *time.Time

	// The maximum number of items to return in the response. If more items exist than
	// the specified MaxRecords value, a pagination token is included in the response
	// so that the remaining results can be retrieved.
	MaxResults *int32

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the token, up to the
	// value specified by MaxRecords .
	NextToken *string

	// The returned timestamp which is the start or end time of the time periods. The
	// default value is END_TIME .
	PeriodAlignment types.PeriodAlignment

	// The granularity, in seconds, of the data points returned from Performance
	// Insights. A period can be as short as one second, or as long as one day (86400
	// seconds). Valid values are:
	//   - 1 (one second)
	//   - 60 (one minute)
	//   - 300 (five minutes)
	//   - 3600 (one hour)
	//   - 86400 (twenty-four hours)
	// If you don't specify PeriodInSeconds , then Performance Insights will choose a
	// value for you, with a goal of returning roughly 100-200 data points in the
	// response.
	PeriodInSeconds *int32

	noSmithyDocumentSerde
}

type GetResourceMetricsOutput struct {

	// The end time for the returned metrics, after alignment to a granular boundary
	// (as specified by PeriodInSeconds ). AlignedEndTime will be greater than or
	// equal to the value of the user-specified Endtime .
	AlignedEndTime *time.Time

	// The start time for the returned metrics, after alignment to a granular boundary
	// (as specified by PeriodInSeconds ). AlignedStartTime will be less than or equal
	// to the value of the user-specified StartTime .
	AlignedStartTime *time.Time

	// An immutable identifier for a data source that is unique for an Amazon Web
	// Services Region. Performance Insights gathers metrics from this data source. In
	// the console, the identifier is shown as ResourceID. When you call
	// DescribeDBInstances , the identifier is returned as DbiResourceId .
	Identifier *string

	// An array of metric results, where each array element contains all of the data
	// points for a particular dimension.
	MetricList []types.MetricKeyDataPoints

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the token, up to the
	// value specified by MaxRecords .
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetResourceMetricsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetResourceMetrics{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetResourceMetrics{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetResourceMetrics"); err != nil {
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
	if err = addOpGetResourceMetricsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetResourceMetrics(options.Region), middleware.Before); err != nil {
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

// GetResourceMetricsAPIClient is a client that implements the GetResourceMetrics
// operation.
type GetResourceMetricsAPIClient interface {
	GetResourceMetrics(context.Context, *GetResourceMetricsInput, ...func(*Options)) (*GetResourceMetricsOutput, error)
}

var _ GetResourceMetricsAPIClient = (*Client)(nil)

// GetResourceMetricsPaginatorOptions is the paginator options for
// GetResourceMetrics
type GetResourceMetricsPaginatorOptions struct {
	// The maximum number of items to return in the response. If more items exist than
	// the specified MaxRecords value, a pagination token is included in the response
	// so that the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetResourceMetricsPaginator is a paginator for GetResourceMetrics
type GetResourceMetricsPaginator struct {
	options   GetResourceMetricsPaginatorOptions
	client    GetResourceMetricsAPIClient
	params    *GetResourceMetricsInput
	nextToken *string
	firstPage bool
}

// NewGetResourceMetricsPaginator returns a new GetResourceMetricsPaginator
func NewGetResourceMetricsPaginator(client GetResourceMetricsAPIClient, params *GetResourceMetricsInput, optFns ...func(*GetResourceMetricsPaginatorOptions)) *GetResourceMetricsPaginator {
	if params == nil {
		params = &GetResourceMetricsInput{}
	}

	options := GetResourceMetricsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetResourceMetricsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetResourceMetricsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetResourceMetrics page.
func (p *GetResourceMetricsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetResourceMetricsOutput, error) {
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

	result, err := p.client.GetResourceMetrics(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetResourceMetrics(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetResourceMetrics",
	}
}
