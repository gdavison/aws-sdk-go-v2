// Code generated by smithy-go-codegen DO NOT EDIT.

package budgets

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/budgets/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a budget action.
func (c *Client) CreateBudgetAction(ctx context.Context, params *CreateBudgetActionInput, optFns ...func(*Options)) (*CreateBudgetActionOutput, error) {
	if params == nil {
		params = &CreateBudgetActionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBudgetAction", params, optFns, c.addOperationCreateBudgetActionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBudgetActionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBudgetActionInput struct {

	// The account ID of the user. It's a 12-digit number.
	//
	// This member is required.
	AccountId *string

	// The trigger threshold of the action.
	//
	// This member is required.
	ActionThreshold *types.ActionThreshold

	// The type of action. This defines the type of tasks that can be carried out by
	// this action. This field also determines the format for definition.
	//
	// This member is required.
	ActionType types.ActionType

	// This specifies if the action needs manual or automatic approval.
	//
	// This member is required.
	ApprovalModel types.ApprovalModel

	// A string that represents the budget name. The ":" and "\" characters, and the
	// "/action/" substring, aren't allowed.
	//
	// This member is required.
	BudgetName *string

	// Specifies all of the type-specific parameters.
	//
	// This member is required.
	Definition *types.Definition

	// The role passed for action execution and reversion. Roles and actions must be
	// in the same account.
	//
	// This member is required.
	ExecutionRoleArn *string

	// The type of a notification. It must be ACTUAL or FORECASTED.
	//
	// This member is required.
	NotificationType types.NotificationType

	// A list of subscribers.
	//
	// This member is required.
	Subscribers []types.Subscriber

	// An optional list of tags to associate with the specified budget action. Each
	// tag consists of a key and a value, and each key must be unique for the resource.
	ResourceTags []types.ResourceTag

	noSmithyDocumentSerde
}

type CreateBudgetActionOutput struct {

	// The account ID of the user. It's a 12-digit number.
	//
	// This member is required.
	AccountId *string

	// A system-generated universally unique identifier (UUID) for the action.
	//
	// This member is required.
	ActionId *string

	// A string that represents the budget name. The ":" and "\" characters, and the
	// "/action/" substring, aren't allowed.
	//
	// This member is required.
	BudgetName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateBudgetActionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateBudgetAction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateBudgetAction{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateBudgetAction"); err != nil {
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
	if err = addOpCreateBudgetActionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBudgetAction(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateBudgetAction(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateBudgetAction",
	}
}
