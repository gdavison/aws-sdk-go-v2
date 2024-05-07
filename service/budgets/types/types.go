// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// A budget action resource.
type Action struct {

	// A system-generated universally unique identifier (UUID) for the action.
	//
	// This member is required.
	ActionId *string

	// The trigger threshold of the action.
	//
	// This member is required.
	ActionThreshold *ActionThreshold

	// The type of action. This defines the type of tasks that can be carried out by
	// this action. This field also determines the format for definition.
	//
	// This member is required.
	ActionType ActionType

	// This specifies if the action needs manual or automatic approval.
	//
	// This member is required.
	ApprovalModel ApprovalModel

	// A string that represents the budget name. The ":" and "\" characters, and the
	// "/action/" substring, aren't allowed.
	//
	// This member is required.
	BudgetName *string

	// Where you specify all of the type-specific parameters.
	//
	// This member is required.
	Definition *Definition

	// The role passed for action execution and reversion. Roles and actions must be
	// in the same account.
	//
	// This member is required.
	ExecutionRoleArn *string

	// The type of a notification. It must be ACTUAL or FORECASTED.
	//
	// This member is required.
	NotificationType NotificationType

	// The status of the action.
	//
	// This member is required.
	Status ActionStatus

	// A list of subscribers.
	//
	// This member is required.
	Subscribers []Subscriber

	noSmithyDocumentSerde
}

// The historical records for a budget action.
type ActionHistory struct {

	// The description of the details for the event.
	//
	// This member is required.
	ActionHistoryDetails *ActionHistoryDetails

	// This distinguishes between whether the events are triggered by the user or are
	// generated by the system.
	//
	// This member is required.
	EventType EventType

	// The status of action at the time of the event.
	//
	// This member is required.
	Status ActionStatus

	// A generic time stamp. In Java, it's transformed to a Date object.
	//
	// This member is required.
	Timestamp *time.Time

	noSmithyDocumentSerde
}

// The description of the details for the event.
type ActionHistoryDetails struct {

	// The budget action resource.
	//
	// This member is required.
	Action *Action

	// A generic string.
	//
	// This member is required.
	Message *string

	noSmithyDocumentSerde
}

// The trigger threshold of the action.
type ActionThreshold struct {

	// The type of threshold for a notification.
	//
	// This member is required.
	ActionThresholdType ThresholdType

	// The threshold of a notification.
	//
	// This member is required.
	ActionThresholdValue float64

	noSmithyDocumentSerde
}

// The parameters that determine the budget amount for an auto-adjusting budget.
type AutoAdjustData struct {

	// The string that defines whether your budget auto-adjusts based on historical or
	// forecasted data.
	//
	// This member is required.
	AutoAdjustType AutoAdjustType

	// The parameters that define or describe the historical data that your
	// auto-adjusting budget is based on.
	HistoricalOptions *HistoricalOptions

	// The last time that your budget was auto-adjusted.
	LastAutoAdjustTime *time.Time

	noSmithyDocumentSerde
}

// Represents the output of the CreateBudget operation. The content consists of
// the detailed metadata and data file information, and the current status of the
// budget object. This is the Amazon Resource Name (ARN) pattern for a budget:
// arn:aws:budgets::AccountId:budget/budgetName
type Budget struct {

	// The name of a budget. The name must be unique within an account. The : and \
	// characters, and the "/action/" substring, aren't allowed in BudgetName .
	//
	// This member is required.
	BudgetName *string

	// Specifies whether this budget tracks costs, usage, RI utilization, RI coverage,
	// Savings Plans utilization, or Savings Plans coverage.
	//
	// This member is required.
	BudgetType BudgetType

	// The length of time until a budget resets the actual and forecasted spend.
	//
	// This member is required.
	TimeUnit TimeUnit

	// The parameters that determine the budget amount for an auto-adjusting budget.
	AutoAdjustData *AutoAdjustData

	// The total amount of cost, usage, RI utilization, RI coverage, Savings Plans
	// utilization, or Savings Plans coverage that you want to track with your budget.
	// BudgetLimit is required for cost or usage budgets, but optional for RI or
	// Savings Plans utilization or coverage budgets. RI and Savings Plans utilization
	// or coverage budgets default to 100 . This is the only valid value for RI or
	// Savings Plans utilization or coverage budgets. You can't use BudgetLimit with
	// PlannedBudgetLimits for CreateBudget and UpdateBudget actions.
	BudgetLimit *Spend

	// The actual and forecasted cost or usage that the budget tracks.
	CalculatedSpend *CalculatedSpend

	// The cost filters, such as Region , Service , LinkedAccount , Tag , or
	// CostCategory , that are applied to a budget. Amazon Web Services Budgets
	// supports the following services as a Service filter for RI budgets:
	//   - Amazon EC2
	//   - Amazon Redshift
	//   - Amazon Relational Database Service
	//   - Amazon ElastiCache
	//   - Amazon OpenSearch Service
	CostFilters map[string][]string

	// The types of costs that are included in this COST budget. USAGE , RI_UTILIZATION
	// , RI_COVERAGE , SAVINGS_PLANS_UTILIZATION , and SAVINGS_PLANS_COVERAGE budgets
	// do not have CostTypes .
	CostTypes *CostTypes

	// The last time that you updated this budget.
	LastUpdatedTime *time.Time

	// A map containing multiple BudgetLimit , including current or future limits.
	// PlannedBudgetLimits is available for cost or usage budget and supports both
	// monthly and quarterly TimeUnit . For monthly budgets, provide 12 months of
	// PlannedBudgetLimits values. This must start from the current month and include
	// the next 11 months. The key is the start of the month, UTC in epoch seconds.
	// For quarterly budgets, provide four quarters of PlannedBudgetLimits value
	// entries in standard calendar quarter increments. This must start from the
	// current quarter and include the next three quarters. The key is the start of
	// the quarter, UTC in epoch seconds. If the planned budget expires before 12
	// months for monthly or four quarters for quarterly, provide the
	// PlannedBudgetLimits values only for the remaining periods. If the budget begins
	// at a date in the future, provide PlannedBudgetLimits values from the start date
	// of the budget. After all of the BudgetLimit values in PlannedBudgetLimits are
	// used, the budget continues to use the last limit as the BudgetLimit . At that
	// point, the planned budget provides the same experience as a fixed budget.
	// DescribeBudget and DescribeBudgets response along with PlannedBudgetLimits also
	// contain BudgetLimit representing the current month or quarter limit present in
	// PlannedBudgetLimits . This only applies to budgets that are created with
	// PlannedBudgetLimits . Budgets that are created without PlannedBudgetLimits only
	// contain BudgetLimit . They don't contain PlannedBudgetLimits .
	PlannedBudgetLimits map[string]Spend

	// The period of time that's covered by a budget. You setthe start date and end
	// date. The start date must come before the end date. The end date must come
	// before 06/15/87 00:00 UTC . If you create your budget and don't specify a start
	// date, Amazon Web Services defaults to the start of your chosen time period
	// (DAILY, MONTHLY, QUARTERLY, or ANNUALLY). For example, if you created your
	// budget on January 24, 2018, chose DAILY , and didn't set a start date, Amazon
	// Web Services set your start date to 01/24/18 00:00 UTC . If you chose MONTHLY ,
	// Amazon Web Services set your start date to 01/01/18 00:00 UTC . If you didn't
	// specify an end date, Amazon Web Services set your end date to 06/15/87 00:00 UTC
	// . The defaults are the same for the Billing and Cost Management console and the
	// API. You can change either date with the UpdateBudget operation. After the end
	// date, Amazon Web Services deletes the budget and all the associated
	// notifications and subscribers.
	TimePeriod *TimePeriod

	noSmithyDocumentSerde
}

// The amount of cost or usage that you created the budget for, compared to your
// actual costs or usage.
type BudgetedAndActualAmounts struct {

	// Your actual costs or usage for a budget period.
	ActualAmount *Spend

	// The amount of cost or usage that you created the budget for.
	BudgetedAmount *Spend

	// The time period that's covered by this budget comparison.
	TimePeriod *TimePeriod

	noSmithyDocumentSerde
}

// The budget name and associated notifications for an account.
type BudgetNotificationsForAccount struct {

	// A string that represents the budget name. The ":" and "\" characters, and the
	// "/action/" substring, aren't allowed.
	BudgetName *string

	// A list of notifications.
	Notifications []Notification

	noSmithyDocumentSerde
}

// A history of the state of a budget at the end of the budget's specified time
// period.
type BudgetPerformanceHistory struct {

	// A string that represents the budget name. The ":" and "\" characters, and the
	// "/action/" substring, aren't allowed.
	BudgetName *string

	// The type of a budget. It must be one of the following types: COST , USAGE ,
	// RI_UTILIZATION , RI_COVERAGE , SAVINGS_PLANS_UTILIZATION , or
	// SAVINGS_PLANS_COVERAGE .
	BudgetType BudgetType

	// A list of amounts of cost or usage that you created budgets for, which are
	// compared to your actual costs or usage.
	BudgetedAndActualAmountsList []BudgetedAndActualAmounts

	// The history of the cost filters for a budget during the specified time period.
	CostFilters map[string][]string

	// The history of the cost types for a budget during the specified time period.
	CostTypes *CostTypes

	// The time unit of the budget, such as MONTHLY or QUARTERLY.
	TimeUnit TimeUnit

	noSmithyDocumentSerde
}

// The spend objects that are associated with this budget. The actualSpend tracks
// how much you've used, cost, usage, RI units, or Savings Plans units and the
// forecastedSpend tracks how much that you're predicted to spend based on your
// historical usage profile. For example, if it's the 20th of the month and you
// have spent 50 dollars on Amazon EC2, your actualSpend is 50 USD , and your
// forecastedSpend is 75 USD .
type CalculatedSpend struct {

	// The amount of cost, usage, RI units, or Savings Plans units that you used.
	//
	// This member is required.
	ActualSpend *Spend

	// The amount of cost, usage, RI units, or Savings Plans units that you're
	// forecasted to use.
	ForecastedSpend *Spend

	noSmithyDocumentSerde
}

// The types of cost that are included in a COST budget, such as tax and
// subscriptions. USAGE , RI_UTILIZATION , RI_COVERAGE , SAVINGS_PLANS_UTILIZATION
// , and SAVINGS_PLANS_COVERAGE budgets don't have CostTypes .
type CostTypes struct {

	// Specifies whether a budget includes credits. The default value is true .
	IncludeCredit *bool

	// Specifies whether a budget includes discounts. The default value is true .
	IncludeDiscount *bool

	// Specifies whether a budget includes non-RI subscription costs. The default
	// value is true .
	IncludeOtherSubscription *bool

	// Specifies whether a budget includes recurring fees such as monthly RI fees. The
	// default value is true .
	IncludeRecurring *bool

	// Specifies whether a budget includes refunds. The default value is true .
	IncludeRefund *bool

	// Specifies whether a budget includes subscriptions. The default value is true .
	IncludeSubscription *bool

	// Specifies whether a budget includes support subscription fees. The default
	// value is true .
	IncludeSupport *bool

	// Specifies whether a budget includes taxes. The default value is true .
	IncludeTax *bool

	// Specifies whether a budget includes upfront RI costs. The default value is true .
	IncludeUpfront *bool

	// Specifies whether a budget uses the amortized rate. The default value is false .
	UseAmortized *bool

	// Specifies whether a budget uses a blended rate. The default value is false .
	UseBlended *bool

	noSmithyDocumentSerde
}

// Specifies all of the type-specific parameters.
type Definition struct {

	// The Identity and Access Management (IAM) action definition details.
	IamActionDefinition *IamActionDefinition

	// The service control policies (SCPs) action definition details.
	ScpActionDefinition *ScpActionDefinition

	// The Amazon Web Services Systems Manager (SSM) action definition details.
	SsmActionDefinition *SsmActionDefinition

	noSmithyDocumentSerde
}

// The parameters that define or describe the historical data that your
// auto-adjusting budget is based on.
type HistoricalOptions struct {

	// The number of budget periods included in the moving-average calculation that
	// determines your auto-adjusted budget amount. The maximum value depends on the
	// TimeUnit granularity of the budget:
	//   - For the DAILY granularity, the maximum value is 60 .
	//   - For the MONTHLY granularity, the maximum value is 12 .
	//   - For the QUARTERLY granularity, the maximum value is 4 .
	//   - For the ANNUALLY granularity, the maximum value is 1 .
	//
	// This member is required.
	BudgetAdjustmentPeriod *int32

	// The integer that describes how many budget periods in your
	// BudgetAdjustmentPeriod are included in the calculation of your current
	// BudgetLimit . If the first budget period in your BudgetAdjustmentPeriod has no
	// cost data, then that budget period isn’t included in the average that determines
	// your budget limit. For example, if you set BudgetAdjustmentPeriod as 4
	// quarters, but your account had no cost data in the first quarter, then only the
	// last three quarters are included in the calculation. In this scenario,
	// LookBackAvailablePeriods returns 3 . You can’t set your own
	// LookBackAvailablePeriods . The value is automatically calculated from the
	// BudgetAdjustmentPeriod and your historical cost data.
	LookBackAvailablePeriods *int32

	noSmithyDocumentSerde
}

// The Identity and Access Management (IAM) action definition details.
type IamActionDefinition struct {

	// The Amazon Resource Name (ARN) of the policy to be attached.
	//
	// This member is required.
	PolicyArn *string

	// A list of groups to be attached. There must be at least one group.
	Groups []string

	// A list of roles to be attached. There must be at least one role.
	Roles []string

	// A list of users to be attached. There must be at least one user.
	Users []string

	noSmithyDocumentSerde
}

// A notification that's associated with a budget. A budget can have up to ten
// notifications. Each notification must have at least one subscriber. A
// notification can have one SNS subscriber and up to 10 email subscribers, for a
// total of 11 subscribers. For example, if you have a budget for 200 dollars and
// you want to be notified when you go over 160 dollars, create a notification with
// the following parameters:
//   - A notificationType of ACTUAL
//   - A thresholdType of PERCENTAGE
//   - A comparisonOperator of GREATER_THAN
//   - A notification threshold of 80
type Notification struct {

	// The comparison that's used for this notification.
	//
	// This member is required.
	ComparisonOperator ComparisonOperator

	// Specifies whether the notification is for how much you have spent ( ACTUAL ) or
	// for how much that you're forecasted to spend ( FORECASTED ).
	//
	// This member is required.
	NotificationType NotificationType

	// The threshold that's associated with a notification. Thresholds are always a
	// percentage, and many customers find value being alerted between 50% - 200% of
	// the budgeted amount. The maximum limit for your threshold is 1,000,000% above
	// the budgeted amount.
	//
	// This member is required.
	Threshold float64

	// Specifies whether this notification is in alarm. If a budget notification is in
	// the ALARM state, you passed the set threshold for the budget.
	NotificationState NotificationState

	// The type of threshold for a notification. For ABSOLUTE_VALUE thresholds, Amazon
	// Web Services notifies you when you go over or are forecasted to go over your
	// total cost threshold. For PERCENTAGE thresholds, Amazon Web Services notifies
	// you when you go over or are forecasted to go over a certain percentage of your
	// forecasted spend. For example, if you have a budget for 200 dollars and you have
	// a PERCENTAGE threshold of 80%, Amazon Web Services notifies you when you go
	// over 160 dollars.
	ThresholdType ThresholdType

	noSmithyDocumentSerde
}

// A notification with subscribers. A notification can have one SNS subscriber and
// up to 10 email subscribers, for a total of 11 subscribers.
type NotificationWithSubscribers struct {

	// The notification that's associated with a budget.
	//
	// This member is required.
	Notification *Notification

	// A list of subscribers who are subscribed to this notification.
	//
	// This member is required.
	Subscribers []Subscriber

	noSmithyDocumentSerde
}

// The tag structure that contains a tag key and value.
type ResourceTag struct {

	// The key that's associated with the tag.
	//
	// This member is required.
	Key *string

	// The value that's associated with the tag.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// The service control policies (SCP) action definition details.
type ScpActionDefinition struct {

	// The policy ID attached.
	//
	// This member is required.
	PolicyId *string

	// A list of target IDs.
	//
	// This member is required.
	TargetIds []string

	noSmithyDocumentSerde
}

// The amount of cost or usage that's measured for a budget. Cost example: A Spend
// for 3 USD of costs has the following parameters:
//   - An Amount of 3
//   - A Unit of USD
//
// Usage example: A Spend for 3 GB of S3 usage has the following parameters:
//   - An Amount of 3
//   - A Unit of GB
type Spend struct {

	// The cost or usage amount that's associated with a budget forecast, actual
	// spend, or budget threshold.
	//
	// This member is required.
	Amount *string

	// The unit of measurement that's used for the budget forecast, actual spend, or
	// budget threshold.
	//
	// This member is required.
	Unit *string

	noSmithyDocumentSerde
}

// The Amazon Web Services Systems Manager (SSM) action definition details.
type SsmActionDefinition struct {

	// The action subType.
	//
	// This member is required.
	ActionSubType ActionSubType

	// The EC2 and RDS instance IDs.
	//
	// This member is required.
	InstanceIds []string

	// The Region to run the SSM document.
	//
	// This member is required.
	Region *string

	noSmithyDocumentSerde
}

// The subscriber to a budget notification. The subscriber consists of a
// subscription type and either an Amazon SNS topic or an email address. For
// example, an email subscriber has the following parameters:
//   - A subscriptionType of EMAIL
//   - An address of example@example.com
type Subscriber struct {

	// The address that Amazon Web Services sends budget notifications to, either an
	// SNS topic or an email. When you create a subscriber, the value of Address can't
	// contain line breaks.
	//
	// This member is required.
	Address *string

	// The type of notification that Amazon Web Services sends to a subscriber.
	//
	// This member is required.
	SubscriptionType SubscriptionType

	noSmithyDocumentSerde
}

// The period of time that's covered by a budget. The period has a start date and
// an end date. The start date must come before the end date. There are no
// restrictions on the end date.
type TimePeriod struct {

	// The end date for a budget. If you didn't specify an end date, Amazon Web
	// Services set your end date to 06/15/87 00:00 UTC . The defaults are the same for
	// the Billing and Cost Management console and the API. After the end date, Amazon
	// Web Services deletes the budget and all the associated notifications and
	// subscribers. You can change your end date with the UpdateBudget operation.
	End *time.Time

	// The start date for a budget. If you created your budget and didn't specify a
	// start date, Amazon Web Services defaults to the start of your chosen time period
	// (DAILY, MONTHLY, QUARTERLY, or ANNUALLY). For example, if you created your
	// budget on January 24, 2018, chose DAILY , and didn't set a start date, Amazon
	// Web Services set your start date to 01/24/18 00:00 UTC . If you chose MONTHLY ,
	// Amazon Web Services set your start date to 01/01/18 00:00 UTC . The defaults are
	// the same for the Billing and Cost Management console and the API. You can change
	// your start date with the UpdateBudget operation.
	Start *time.Time

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
