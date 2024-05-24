// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	internalEndpointDiscovery "github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds or removes replicas in the specified global table. The global table must
// already exist to be able to use this operation. Any replica to be added must be
// empty, have the same name as the global table, have the same key schema, have
// DynamoDB Streams enabled, and have the same provisioned and maximum write
// capacity units.
//
// For global tables, this operation only applies to global tables using Version
// 2019.11.21 (Current version), as it provides greater flexibility, higher
// efficiency and consumes less write capacity than 2017.11.29 (Legacy). To
// determine which version you are using, see [Determining the version]. To update existing global tables
// from version 2017.11.29 (Legacy) to version 2019.11.21 (Current), see [Updating global tables].
//
// For global tables, this operation only applies to global tables using Version
// 2019.11.21 (Current version). If you are using global tables [Version 2019.11.21]you can use [UpdateTable]
// instead.
//
// Although you can use UpdateGlobalTable to add replicas and remove replicas in a
// single request, for simplicity we recommend that you issue separate requests for
// adding or removing replicas.
//
// If global secondary indexes are specified, then the following conditions must
// also be met:
//
//   - The global secondary indexes must have the same name.
//
//   - The global secondary indexes must have the same hash key and sort key (if
//     present).
//
//   - The global secondary indexes must have the same provisioned and maximum
//     write capacity units.
//
// [Updating global tables]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/V2globaltables_upgrade.html
// [UpdateTable]: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_UpdateTable.html
// [Version 2019.11.21]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GlobalTables.html
// [Determining the version]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.DetermineVersion.html
func (c *Client) UpdateGlobalTable(ctx context.Context, params *UpdateGlobalTableInput, optFns ...func(*Options)) (*UpdateGlobalTableOutput, error) {
	if params == nil {
		params = &UpdateGlobalTableInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateGlobalTable", params, optFns, c.addOperationUpdateGlobalTableMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateGlobalTableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateGlobalTableInput struct {

	// The global table name.
	//
	// This member is required.
	GlobalTableName *string

	// A list of Regions that should be added or removed from the global table.
	//
	// This member is required.
	ReplicaUpdates []types.ReplicaUpdate

	noSmithyDocumentSerde
}

type UpdateGlobalTableOutput struct {

	// Contains the details of the global table.
	GlobalTableDescription *types.GlobalTableDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateGlobalTableMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateGlobalTable{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateGlobalTable{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateGlobalTable"); err != nil {
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
	if err = addOpUpdateGlobalTableDiscoverEndpointMiddleware(stack, options, c); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateGlobalTableValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateGlobalTable(options.Region), middleware.Before); err != nil {
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
	if err = addValidateResponseChecksum(stack, options); err != nil {
		return err
	}
	if err = addAcceptEncodingGzip(stack, options); err != nil {
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

func addOpUpdateGlobalTableDiscoverEndpointMiddleware(stack *middleware.Stack, o Options, c *Client) error {
	return stack.Finalize.Insert(&internalEndpointDiscovery.DiscoverEndpoint{
		Options: []func(*internalEndpointDiscovery.DiscoverEndpointOptions){
			func(opt *internalEndpointDiscovery.DiscoverEndpointOptions) {
				opt.DisableHTTPS = o.EndpointOptions.DisableHTTPS
				opt.Logger = o.Logger
			},
		},
		DiscoverOperation:            c.fetchOpUpdateGlobalTableDiscoverEndpoint,
		EndpointDiscoveryEnableState: o.EndpointDiscovery.EnableEndpointDiscovery,
		EndpointDiscoveryRequired:    false,
		Region:                       o.Region,
	}, "ResolveEndpointV2", middleware.After)
}

func (c *Client) fetchOpUpdateGlobalTableDiscoverEndpoint(ctx context.Context, region string, optFns ...func(*internalEndpointDiscovery.DiscoverEndpointOptions)) (internalEndpointDiscovery.WeightedAddress, error) {
	input := getOperationInput(ctx)
	in, ok := input.(*UpdateGlobalTableInput)
	if !ok {
		return internalEndpointDiscovery.WeightedAddress{}, fmt.Errorf("unknown input type %T", input)
	}
	_ = in

	identifierMap := make(map[string]string, 0)
	identifierMap["sdk#Region"] = region

	key := fmt.Sprintf("DynamoDB.%v", identifierMap)

	if v, ok := c.endpointCache.Get(key); ok {
		return v, nil
	}

	discoveryOperationInput := &DescribeEndpointsInput{}

	opt := internalEndpointDiscovery.DiscoverEndpointOptions{}
	for _, fn := range optFns {
		fn(&opt)
	}

	go c.handleEndpointDiscoveryFromService(ctx, discoveryOperationInput, region, key, opt)
	return internalEndpointDiscovery.WeightedAddress{}, nil
}

func newServiceMetadataMiddleware_opUpdateGlobalTable(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateGlobalTable",
	}
}
