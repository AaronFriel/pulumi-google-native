// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Immutable. The type of data this data collector will collect.
type DataCollectorType pulumi.String

const (
	// For future compatibility.
	DataCollectorTypeTypeUnspecified = DataCollectorType("TYPE_UNSPECIFIED")
	// For integer values.
	DataCollectorTypeInteger = DataCollectorType("INTEGER")
	// For float values.
	DataCollectorTypeFloat = DataCollectorType("FLOAT")
	// For string values.
	DataCollectorTypeString = DataCollectorType("STRING")
	// For boolean values.
	DataCollectorTypeBoolean = DataCollectorType("BOOLEAN")
	// For datetime values.
	DataCollectorTypeDatetime = DataCollectorType("DATETIME")
)

func (DataCollectorType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e DataCollectorType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataCollectorType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataCollectorType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DataCollectorType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Sampler of distributed tracing. OFF is the default value.
type GoogleCloudApigeeV1TraceSamplingConfigSampler pulumi.String

const (
	// Sampler unspecified.
	GoogleCloudApigeeV1TraceSamplingConfigSamplerSamplerUnspecified = GoogleCloudApigeeV1TraceSamplingConfigSampler("SAMPLER_UNSPECIFIED")
	// OFF means distributed trace is disabled, or the sampling probability is 0.
	GoogleCloudApigeeV1TraceSamplingConfigSamplerOff = GoogleCloudApigeeV1TraceSamplingConfigSampler("OFF")
	// PROBABILITY means traces are captured on a probability that defined by sampling_rate. The sampling rate is limited to 0 to 0.5 when this is set.
	GoogleCloudApigeeV1TraceSamplingConfigSamplerProbability = GoogleCloudApigeeV1TraceSamplingConfigSampler("PROBABILITY")
)

func (GoogleCloudApigeeV1TraceSamplingConfigSampler) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e GoogleCloudApigeeV1TraceSamplingConfigSampler) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudApigeeV1TraceSamplingConfigSampler) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudApigeeV1TraceSamplingConfigSampler) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GoogleCloudApigeeV1TraceSamplingConfigSampler) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The log type that this config enables.
type GoogleIamV1AuditLogConfigLogType pulumi.String

const (
	// Default case. Should never be this.
	GoogleIamV1AuditLogConfigLogTypeLogTypeUnspecified = GoogleIamV1AuditLogConfigLogType("LOG_TYPE_UNSPECIFIED")
	// Admin reads. Example: CloudIAM getIamPolicy
	GoogleIamV1AuditLogConfigLogTypeAdminRead = GoogleIamV1AuditLogConfigLogType("ADMIN_READ")
	// Data writes. Example: CloudSQL Users create
	GoogleIamV1AuditLogConfigLogTypeDataWrite = GoogleIamV1AuditLogConfigLogType("DATA_WRITE")
	// Data reads. Example: CloudSQL Users list
	GoogleIamV1AuditLogConfigLogTypeDataRead = GoogleIamV1AuditLogConfigLogType("DATA_READ")
)

func (GoogleIamV1AuditLogConfigLogType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e GoogleIamV1AuditLogConfigLogType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleIamV1AuditLogConfigLogType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleIamV1AuditLogConfigLogType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GoogleIamV1AuditLogConfigLogType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Optional. Size of the CIDR block range that will be reserved by the instance. PAID organizations support `SLASH_16` to `SLASH_20` and defaults to `SLASH_16`. Evaluation organizations support only `SLASH_23`.
type InstancePeeringCidrRange pulumi.String

const (
	// Range not specified.
	InstancePeeringCidrRangeCidrRangeUnspecified = InstancePeeringCidrRange("CIDR_RANGE_UNSPECIFIED")
	// `/16` CIDR range.
	InstancePeeringCidrRangeSlash16 = InstancePeeringCidrRange("SLASH_16")
	// `/17` CIDR range.
	InstancePeeringCidrRangeSlash17 = InstancePeeringCidrRange("SLASH_17")
	// `/18` CIDR range.
	InstancePeeringCidrRangeSlash18 = InstancePeeringCidrRange("SLASH_18")
	// `/19` CIDR range.
	InstancePeeringCidrRangeSlash19 = InstancePeeringCidrRange("SLASH_19")
	// `/20` CIDR range.
	InstancePeeringCidrRangeSlash20 = InstancePeeringCidrRange("SLASH_20")
	// `/23` CIDR range. Supported for evaluation only.
	InstancePeeringCidrRangeSlash23 = InstancePeeringCidrRange("SLASH_23")
)

func (InstancePeeringCidrRange) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e InstancePeeringCidrRange) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e InstancePeeringCidrRange) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e InstancePeeringCidrRange) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e InstancePeeringCidrRange) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Billing type of the Apigee organization. See [Apigee pricing](https://cloud.google.com/apigee/pricing).
type OrganizationBillingType pulumi.String

const (
	// Billing type not specified.
	OrganizationBillingTypeBillingTypeUnspecified = OrganizationBillingType("BILLING_TYPE_UNSPECIFIED")
	// A pre-paid subscription to Apigee.
	OrganizationBillingTypeSubscription = OrganizationBillingType("SUBSCRIPTION")
	// Free and limited access to Apigee for evaluation purposes only. only.
	OrganizationBillingTypeEvaluation = OrganizationBillingType("EVALUATION")
)

func (OrganizationBillingType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e OrganizationBillingType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OrganizationBillingType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OrganizationBillingType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OrganizationBillingType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Required. Runtime type of the Apigee organization based on the Apigee subscription purchased.
type OrganizationRuntimeType pulumi.String

const (
	// Runtime type not specified.
	OrganizationRuntimeTypeRuntimeTypeUnspecified = OrganizationRuntimeType("RUNTIME_TYPE_UNSPECIFIED")
	// Google-managed Apigee runtime.
	OrganizationRuntimeTypeCloud = OrganizationRuntimeType("CLOUD")
	// User-managed Apigee hybrid runtime.
	OrganizationRuntimeTypeHybrid = OrganizationRuntimeType("HYBRID")
)

func (OrganizationRuntimeType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e OrganizationRuntimeType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OrganizationRuntimeType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OrganizationRuntimeType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OrganizationRuntimeType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Not used by Apigee.
type OrganizationType pulumi.String

const (
	// Subscription type not specified.
	OrganizationTypeTypeUnspecified = OrganizationType("TYPE_UNSPECIFIED")
	// Subscription to Apigee is free, limited, and used for evaluation purposes only.
	OrganizationTypeTypeTrial = OrganizationType("TYPE_TRIAL")
	// Full subscription to Apigee has been purchased. See [Apigee pricing](https://cloud.google.com/apigee/pricing/).
	OrganizationTypeTypePaid = OrganizationType("TYPE_PAID")
	// For internal users only.
	OrganizationTypeTypeInternal = OrganizationType("TYPE_INTERNAL")
)

func (OrganizationType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e OrganizationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OrganizationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OrganizationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OrganizationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Frequency at which the customer will be billed.
type RatePlanBillingPeriod pulumi.String

const (
	// Billing period not specified.
	RatePlanBillingPeriodBillingPeriodUnspecified = RatePlanBillingPeriod("BILLING_PERIOD_UNSPECIFIED")
	// Weekly billing period. **Note**: Not supported by Apigee at this time.
	RatePlanBillingPeriodWeekly = RatePlanBillingPeriod("WEEKLY")
	// Monthly billing period.
	RatePlanBillingPeriodMonthly = RatePlanBillingPeriod("MONTHLY")
)

func (RatePlanBillingPeriod) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e RatePlanBillingPeriod) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RatePlanBillingPeriod) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RatePlanBillingPeriod) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RatePlanBillingPeriod) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Pricing model used for consumption-based charges.
type RatePlanConsumptionPricingType pulumi.String

const (
	// Pricing model not specified. This is the default.
	RatePlanConsumptionPricingTypeConsumptionPricingTypeUnspecified = RatePlanConsumptionPricingType("CONSUMPTION_PRICING_TYPE_UNSPECIFIED")
	// Fixed rate charged for each API call.
	RatePlanConsumptionPricingTypeFixedPerUnit = RatePlanConsumptionPricingType("FIXED_PER_UNIT")
	// Variable rate charged based on the total volume of API calls. Example: * 1-100 calls cost $2 per call * 101-200 calls cost $1.50 per call * 201-300 calls cost $1 per call * Total price for 50 calls: 50 x $2 = $100 * Total price for 150 calls: 150 x $1.5 = $225 * Total price for 250 calls: 250 x $1 = $250. **Note**: Not supported by Apigee at this time.
	RatePlanConsumptionPricingTypeBanded = RatePlanConsumptionPricingType("BANDED")
	// Variable rate charged for each API call based on price tiers. Example: * 1-100 calls cost $2 per call * 101-200 calls cost $1.50 per call * 201-300 calls cost $1 per call * Total price for 50 calls: 50 x $2 = $100 * Total price for 150 calls: 100 x $2 + 50 x $1.5 = $275 * Total price for 250 calls: 100 x $2 + 100 x $1.5 + 50 x $1 = $400. **Note**: Not supported by Apigee at this time.
	RatePlanConsumptionPricingTypeTiered = RatePlanConsumptionPricingType("TIERED")
	// Flat rate charged for a bundle of API calls whether or not the entire bundle is used. Example: * 1-100 calls cost $75 flat fee * 101-200 calls cost $100 flat free * 201-300 calls cost $150 flat fee * Total price for 1 call: $75 * Total price for 50 calls: $75 * Total price for 150 calls: $100 * Total price for 250 calls: $150. **Note**: Not supported by Apigee at this time.
	RatePlanConsumptionPricingTypeStairstep = RatePlanConsumptionPricingType("STAIRSTEP")
)

func (RatePlanConsumptionPricingType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e RatePlanConsumptionPricingType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RatePlanConsumptionPricingType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RatePlanConsumptionPricingType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RatePlanConsumptionPricingType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Flag that specifies the billing account type, prepaid or postpaid.
type RatePlanPaymentFundingModel pulumi.String

const (
	// Billing account type not specified.
	RatePlanPaymentFundingModelPaymentFundingModelUnspecified = RatePlanPaymentFundingModel("PAYMENT_FUNDING_MODEL_UNSPECIFIED")
	// Prepaid billing account type. Developer pays in advance for the use of your API products. Funds are deducted from their prepaid account balance. **Note**: Not supported by Apigee at this time.
	RatePlanPaymentFundingModelPrepaid = RatePlanPaymentFundingModel("PREPAID")
	// Postpaid billing account type. Developer is billed through an invoice after using your API products.
	RatePlanPaymentFundingModelPostpaid = RatePlanPaymentFundingModel("POSTPAID")
)

func (RatePlanPaymentFundingModel) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e RatePlanPaymentFundingModel) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RatePlanPaymentFundingModel) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RatePlanPaymentFundingModel) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RatePlanPaymentFundingModel) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Method used to calculate the revenue that is shared with developers.
type RatePlanRevenueShareType pulumi.String

const (
	// Revenue share type is not specified.
	RatePlanRevenueShareTypeRevenueShareTypeUnspecified = RatePlanRevenueShareType("REVENUE_SHARE_TYPE_UNSPECIFIED")
	// Fixed percentage of the total revenue will be shared. The percentage to be shared can be configured by the API provider.
	RatePlanRevenueShareTypeFixed = RatePlanRevenueShareType("FIXED")
	// Amount of revenue shared depends on the number of API calls. The API call volume ranges and the revenue share percentage for each volume can be configured by the API provider. **Note**: Not supported by Apigee at this time.
	RatePlanRevenueShareTypeVolumeBanded = RatePlanRevenueShareType("VOLUME_BANDED")
)

func (RatePlanRevenueShareType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e RatePlanRevenueShareType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RatePlanRevenueShareType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RatePlanRevenueShareType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RatePlanRevenueShareType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Current state of the rate plan (draft or published).
type RatePlanStateEnum pulumi.String

const (
	// State of the rate plan is not specified.
	RatePlanStateEnumStateUnspecified = RatePlanStateEnum("STATE_UNSPECIFIED")
	// Rate plan is in draft mode and only visible to API providers.
	RatePlanStateEnumDraft = RatePlanStateEnum("DRAFT")
	// Rate plan is published and will become visible to developers for the configured duration (between `startTime` and `endTime`).
	RatePlanStateEnumPublished = RatePlanStateEnum("PUBLISHED")
)

func (RatePlanStateEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e RatePlanStateEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RatePlanStateEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RatePlanStateEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RatePlanStateEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}