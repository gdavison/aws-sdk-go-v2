// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ActivationOverrideBehavior string

// Enum values for ActivationOverrideBehavior
const (
	ActivationOverrideBehaviorDistributedGrantsOnly      ActivationOverrideBehavior = "DISTRIBUTED_GRANTS_ONLY"
	ActivationOverrideBehaviorAllGrantsPermittedByIssuer ActivationOverrideBehavior = "ALL_GRANTS_PERMITTED_BY_ISSUER"
)

// Values returns all known values for ActivationOverrideBehavior. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ActivationOverrideBehavior) Values() []ActivationOverrideBehavior {
	return []ActivationOverrideBehavior{
		"DISTRIBUTED_GRANTS_ONLY",
		"ALL_GRANTS_PERMITTED_BY_ISSUER",
	}
}

type AllowedOperation string

// Enum values for AllowedOperation
const (
	AllowedOperationCreateGrant              AllowedOperation = "CreateGrant"
	AllowedOperationCheckoutLicense          AllowedOperation = "CheckoutLicense"
	AllowedOperationCheckoutBorrowLicense    AllowedOperation = "CheckoutBorrowLicense"
	AllowedOperationCheckInLicense           AllowedOperation = "CheckInLicense"
	AllowedOperationExtendConsumptionLicense AllowedOperation = "ExtendConsumptionLicense"
	AllowedOperationListPurchasedLicenses    AllowedOperation = "ListPurchasedLicenses"
	AllowedOperationCreateToken              AllowedOperation = "CreateToken"
)

// Values returns all known values for AllowedOperation. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AllowedOperation) Values() []AllowedOperation {
	return []AllowedOperation{
		"CreateGrant",
		"CheckoutLicense",
		"CheckoutBorrowLicense",
		"CheckInLicense",
		"ExtendConsumptionLicense",
		"ListPurchasedLicenses",
		"CreateToken",
	}
}

type CheckoutType string

// Enum values for CheckoutType
const (
	CheckoutTypeProvisional CheckoutType = "PROVISIONAL"
	CheckoutTypePerpetual   CheckoutType = "PERPETUAL"
)

// Values returns all known values for CheckoutType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (CheckoutType) Values() []CheckoutType {
	return []CheckoutType{
		"PROVISIONAL",
		"PERPETUAL",
	}
}

type DigitalSignatureMethod string

// Enum values for DigitalSignatureMethod
const (
	DigitalSignatureMethodJwtPs384 DigitalSignatureMethod = "JWT_PS384"
)

// Values returns all known values for DigitalSignatureMethod. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DigitalSignatureMethod) Values() []DigitalSignatureMethod {
	return []DigitalSignatureMethod{
		"JWT_PS384",
	}
}

type EntitlementDataUnit string

// Enum values for EntitlementDataUnit
const (
	EntitlementDataUnitCount              EntitlementDataUnit = "Count"
	EntitlementDataUnitNone               EntitlementDataUnit = "None"
	EntitlementDataUnitSeconds            EntitlementDataUnit = "Seconds"
	EntitlementDataUnitMicroseconds       EntitlementDataUnit = "Microseconds"
	EntitlementDataUnitMilliseconds       EntitlementDataUnit = "Milliseconds"
	EntitlementDataUnitBytes              EntitlementDataUnit = "Bytes"
	EntitlementDataUnitKilobytes          EntitlementDataUnit = "Kilobytes"
	EntitlementDataUnitMegabytes          EntitlementDataUnit = "Megabytes"
	EntitlementDataUnitGigabytes          EntitlementDataUnit = "Gigabytes"
	EntitlementDataUnitTerabytes          EntitlementDataUnit = "Terabytes"
	EntitlementDataUnitBits               EntitlementDataUnit = "Bits"
	EntitlementDataUnitKilobits           EntitlementDataUnit = "Kilobits"
	EntitlementDataUnitMegabits           EntitlementDataUnit = "Megabits"
	EntitlementDataUnitGigabits           EntitlementDataUnit = "Gigabits"
	EntitlementDataUnitTerabits           EntitlementDataUnit = "Terabits"
	EntitlementDataUnitPercent            EntitlementDataUnit = "Percent"
	EntitlementDataUnitBytesPerSecond     EntitlementDataUnit = "Bytes/Second"
	EntitlementDataUnitKilobytesPerSecond EntitlementDataUnit = "Kilobytes/Second"
	EntitlementDataUnitMegabytesPerSecond EntitlementDataUnit = "Megabytes/Second"
	EntitlementDataUnitGigabytesPerSecond EntitlementDataUnit = "Gigabytes/Second"
	EntitlementDataUnitTerabytesPerSecond EntitlementDataUnit = "Terabytes/Second"
	EntitlementDataUnitBitsPerSecond      EntitlementDataUnit = "Bits/Second"
	EntitlementDataUnitKilobitsPerSecond  EntitlementDataUnit = "Kilobits/Second"
	EntitlementDataUnitMegabitsPerSecond  EntitlementDataUnit = "Megabits/Second"
	EntitlementDataUnitGigabitsPerSecond  EntitlementDataUnit = "Gigabits/Second"
	EntitlementDataUnitTerabitsPerSecond  EntitlementDataUnit = "Terabits/Second"
	EntitlementDataUnitCountPerSecond     EntitlementDataUnit = "Count/Second"
)

// Values returns all known values for EntitlementDataUnit. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EntitlementDataUnit) Values() []EntitlementDataUnit {
	return []EntitlementDataUnit{
		"Count",
		"None",
		"Seconds",
		"Microseconds",
		"Milliseconds",
		"Bytes",
		"Kilobytes",
		"Megabytes",
		"Gigabytes",
		"Terabytes",
		"Bits",
		"Kilobits",
		"Megabits",
		"Gigabits",
		"Terabits",
		"Percent",
		"Bytes/Second",
		"Kilobytes/Second",
		"Megabytes/Second",
		"Gigabytes/Second",
		"Terabytes/Second",
		"Bits/Second",
		"Kilobits/Second",
		"Megabits/Second",
		"Gigabits/Second",
		"Terabits/Second",
		"Count/Second",
	}
}

type EntitlementUnit string

// Enum values for EntitlementUnit
const (
	EntitlementUnitCount              EntitlementUnit = "Count"
	EntitlementUnitNone               EntitlementUnit = "None"
	EntitlementUnitSeconds            EntitlementUnit = "Seconds"
	EntitlementUnitMicroseconds       EntitlementUnit = "Microseconds"
	EntitlementUnitMilliseconds       EntitlementUnit = "Milliseconds"
	EntitlementUnitBytes              EntitlementUnit = "Bytes"
	EntitlementUnitKilobytes          EntitlementUnit = "Kilobytes"
	EntitlementUnitMegabytes          EntitlementUnit = "Megabytes"
	EntitlementUnitGigabytes          EntitlementUnit = "Gigabytes"
	EntitlementUnitTerabytes          EntitlementUnit = "Terabytes"
	EntitlementUnitBits               EntitlementUnit = "Bits"
	EntitlementUnitKilobits           EntitlementUnit = "Kilobits"
	EntitlementUnitMegabits           EntitlementUnit = "Megabits"
	EntitlementUnitGigabits           EntitlementUnit = "Gigabits"
	EntitlementUnitTerabits           EntitlementUnit = "Terabits"
	EntitlementUnitPercent            EntitlementUnit = "Percent"
	EntitlementUnitBytesPerSecond     EntitlementUnit = "Bytes/Second"
	EntitlementUnitKilobytesPerSecond EntitlementUnit = "Kilobytes/Second"
	EntitlementUnitMegabytesPerSecond EntitlementUnit = "Megabytes/Second"
	EntitlementUnitGigabytesPerSecond EntitlementUnit = "Gigabytes/Second"
	EntitlementUnitTerabytesPerSecond EntitlementUnit = "Terabytes/Second"
	EntitlementUnitBitsPerSecond      EntitlementUnit = "Bits/Second"
	EntitlementUnitKilobitsPerSecond  EntitlementUnit = "Kilobits/Second"
	EntitlementUnitMegabitsPerSecond  EntitlementUnit = "Megabits/Second"
	EntitlementUnitGigabitsPerSecond  EntitlementUnit = "Gigabits/Second"
	EntitlementUnitTerabitsPerSecond  EntitlementUnit = "Terabits/Second"
	EntitlementUnitCountPerSecond     EntitlementUnit = "Count/Second"
)

// Values returns all known values for EntitlementUnit. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EntitlementUnit) Values() []EntitlementUnit {
	return []EntitlementUnit{
		"Count",
		"None",
		"Seconds",
		"Microseconds",
		"Milliseconds",
		"Bytes",
		"Kilobytes",
		"Megabytes",
		"Gigabytes",
		"Terabytes",
		"Bits",
		"Kilobits",
		"Megabits",
		"Gigabits",
		"Terabits",
		"Percent",
		"Bytes/Second",
		"Kilobytes/Second",
		"Megabytes/Second",
		"Gigabytes/Second",
		"Terabytes/Second",
		"Bits/Second",
		"Kilobits/Second",
		"Megabits/Second",
		"Gigabits/Second",
		"Terabits/Second",
		"Count/Second",
	}
}

type GrantStatus string

// Enum values for GrantStatus
const (
	GrantStatusPendingWorkflow   GrantStatus = "PENDING_WORKFLOW"
	GrantStatusPendingAccept     GrantStatus = "PENDING_ACCEPT"
	GrantStatusRejected          GrantStatus = "REJECTED"
	GrantStatusActive            GrantStatus = "ACTIVE"
	GrantStatusFailedWorkflow    GrantStatus = "FAILED_WORKFLOW"
	GrantStatusDeleted           GrantStatus = "DELETED"
	GrantStatusPendingDelete     GrantStatus = "PENDING_DELETE"
	GrantStatusDisabled          GrantStatus = "DISABLED"
	GrantStatusWorkflowCompleted GrantStatus = "WORKFLOW_COMPLETED"
)

// Values returns all known values for GrantStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (GrantStatus) Values() []GrantStatus {
	return []GrantStatus{
		"PENDING_WORKFLOW",
		"PENDING_ACCEPT",
		"REJECTED",
		"ACTIVE",
		"FAILED_WORKFLOW",
		"DELETED",
		"PENDING_DELETE",
		"DISABLED",
		"WORKFLOW_COMPLETED",
	}
}

type InventoryFilterCondition string

// Enum values for InventoryFilterCondition
const (
	InventoryFilterConditionEquals     InventoryFilterCondition = "EQUALS"
	InventoryFilterConditionNotEquals  InventoryFilterCondition = "NOT_EQUALS"
	InventoryFilterConditionBeginsWith InventoryFilterCondition = "BEGINS_WITH"
	InventoryFilterConditionContains   InventoryFilterCondition = "CONTAINS"
)

// Values returns all known values for InventoryFilterCondition. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (InventoryFilterCondition) Values() []InventoryFilterCondition {
	return []InventoryFilterCondition{
		"EQUALS",
		"NOT_EQUALS",
		"BEGINS_WITH",
		"CONTAINS",
	}
}

type LicenseConfigurationStatus string

// Enum values for LicenseConfigurationStatus
const (
	LicenseConfigurationStatusAvailable LicenseConfigurationStatus = "AVAILABLE"
	LicenseConfigurationStatusDisabled  LicenseConfigurationStatus = "DISABLED"
)

// Values returns all known values for LicenseConfigurationStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (LicenseConfigurationStatus) Values() []LicenseConfigurationStatus {
	return []LicenseConfigurationStatus{
		"AVAILABLE",
		"DISABLED",
	}
}

type LicenseConversionTaskStatus string

// Enum values for LicenseConversionTaskStatus
const (
	LicenseConversionTaskStatusInProgress LicenseConversionTaskStatus = "IN_PROGRESS"
	LicenseConversionTaskStatusSucceeded  LicenseConversionTaskStatus = "SUCCEEDED"
	LicenseConversionTaskStatusFailed     LicenseConversionTaskStatus = "FAILED"
)

// Values returns all known values for LicenseConversionTaskStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (LicenseConversionTaskStatus) Values() []LicenseConversionTaskStatus {
	return []LicenseConversionTaskStatus{
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
	}
}

type LicenseCountingType string

// Enum values for LicenseCountingType
const (
	LicenseCountingTypeVcpu     LicenseCountingType = "vCPU"
	LicenseCountingTypeInstance LicenseCountingType = "Instance"
	LicenseCountingTypeCore     LicenseCountingType = "Core"
	LicenseCountingTypeSocket   LicenseCountingType = "Socket"
)

// Values returns all known values for LicenseCountingType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (LicenseCountingType) Values() []LicenseCountingType {
	return []LicenseCountingType{
		"vCPU",
		"Instance",
		"Core",
		"Socket",
	}
}

type LicenseDeletionStatus string

// Enum values for LicenseDeletionStatus
const (
	LicenseDeletionStatusPendingDelete LicenseDeletionStatus = "PENDING_DELETE"
	LicenseDeletionStatusDeleted       LicenseDeletionStatus = "DELETED"
)

// Values returns all known values for LicenseDeletionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (LicenseDeletionStatus) Values() []LicenseDeletionStatus {
	return []LicenseDeletionStatus{
		"PENDING_DELETE",
		"DELETED",
	}
}

type LicenseStatus string

// Enum values for LicenseStatus
const (
	LicenseStatusAvailable        LicenseStatus = "AVAILABLE"
	LicenseStatusPendingAvailable LicenseStatus = "PENDING_AVAILABLE"
	LicenseStatusDeactivated      LicenseStatus = "DEACTIVATED"
	LicenseStatusSuspended        LicenseStatus = "SUSPENDED"
	LicenseStatusExpired          LicenseStatus = "EXPIRED"
	LicenseStatusPendingDelete    LicenseStatus = "PENDING_DELETE"
	LicenseStatusDeleted          LicenseStatus = "DELETED"
)

// Values returns all known values for LicenseStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (LicenseStatus) Values() []LicenseStatus {
	return []LicenseStatus{
		"AVAILABLE",
		"PENDING_AVAILABLE",
		"DEACTIVATED",
		"SUSPENDED",
		"EXPIRED",
		"PENDING_DELETE",
		"DELETED",
	}
}

type ReceivedStatus string

// Enum values for ReceivedStatus
const (
	ReceivedStatusPendingWorkflow   ReceivedStatus = "PENDING_WORKFLOW"
	ReceivedStatusPendingAccept     ReceivedStatus = "PENDING_ACCEPT"
	ReceivedStatusRejected          ReceivedStatus = "REJECTED"
	ReceivedStatusActive            ReceivedStatus = "ACTIVE"
	ReceivedStatusFailedWorkflow    ReceivedStatus = "FAILED_WORKFLOW"
	ReceivedStatusDeleted           ReceivedStatus = "DELETED"
	ReceivedStatusDisabled          ReceivedStatus = "DISABLED"
	ReceivedStatusWorkflowCompleted ReceivedStatus = "WORKFLOW_COMPLETED"
)

// Values returns all known values for ReceivedStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ReceivedStatus) Values() []ReceivedStatus {
	return []ReceivedStatus{
		"PENDING_WORKFLOW",
		"PENDING_ACCEPT",
		"REJECTED",
		"ACTIVE",
		"FAILED_WORKFLOW",
		"DELETED",
		"DISABLED",
		"WORKFLOW_COMPLETED",
	}
}

type RenewType string

// Enum values for RenewType
const (
	RenewTypeNone    RenewType = "None"
	RenewTypeWeekly  RenewType = "Weekly"
	RenewTypeMonthly RenewType = "Monthly"
)

// Values returns all known values for RenewType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (RenewType) Values() []RenewType {
	return []RenewType{
		"None",
		"Weekly",
		"Monthly",
	}
}

type ReportFrequencyType string

// Enum values for ReportFrequencyType
const (
	ReportFrequencyTypeDay   ReportFrequencyType = "DAY"
	ReportFrequencyTypeWeek  ReportFrequencyType = "WEEK"
	ReportFrequencyTypeMonth ReportFrequencyType = "MONTH"
)

// Values returns all known values for ReportFrequencyType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ReportFrequencyType) Values() []ReportFrequencyType {
	return []ReportFrequencyType{
		"DAY",
		"WEEK",
		"MONTH",
	}
}

type ReportType string

// Enum values for ReportType
const (
	ReportTypeLicenseConfigurationSummaryReport ReportType = "LicenseConfigurationSummaryReport"
	ReportTypeLicenseConfigurationUsageReport   ReportType = "LicenseConfigurationUsageReport"
)

// Values returns all known values for ReportType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ReportType) Values() []ReportType {
	return []ReportType{
		"LicenseConfigurationSummaryReport",
		"LicenseConfigurationUsageReport",
	}
}

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypeEc2Instance                   ResourceType = "EC2_INSTANCE"
	ResourceTypeEc2Host                       ResourceType = "EC2_HOST"
	ResourceTypeEc2Ami                        ResourceType = "EC2_AMI"
	ResourceTypeRds                           ResourceType = "RDS"
	ResourceTypeSystemsManagerManagedInstance ResourceType = "SYSTEMS_MANAGER_MANAGED_INSTANCE"
)

// Values returns all known values for ResourceType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ResourceType) Values() []ResourceType {
	return []ResourceType{
		"EC2_INSTANCE",
		"EC2_HOST",
		"EC2_AMI",
		"RDS",
		"SYSTEMS_MANAGER_MANAGED_INSTANCE",
	}
}

type TokenType string

// Enum values for TokenType
const (
	TokenTypeRefreshToken TokenType = "REFRESH_TOKEN"
)

// Values returns all known values for TokenType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (TokenType) Values() []TokenType {
	return []TokenType{
		"REFRESH_TOKEN",
	}
}
