// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create a rate plan that is associated with an API product in an organization. Using rate plans, API product owners can monetize their API products by configuring one or more of the following: - Billing frequency - Initial setup fees for using an API product - Payment funding model (postpaid only) - Fixed recurring or consumption-based charges for using an API product - Revenue sharing with developer partners An API product can have multiple rate plans associated with it but *only one* rate plan can be active at any point of time. **Note: From the developer's perspective, they purchase API products not rate plans.
// Auto-naming is currently not supported for this resource.
type RatePlan struct {
	pulumi.CustomResourceState

	// Name of the API product that the rate plan is associated with.
	Apiproduct pulumi.StringOutput `pulumi:"apiproduct"`
	// Frequency at which the customer will be billed.
	BillingPeriod pulumi.StringOutput `pulumi:"billingPeriod"`
	// API call volume ranges and the fees charged when the total number of API calls is within a given range. The method used to calculate the final fee depends on the selected pricing model. For example, if the pricing model is `STAIRSTEP` and the ranges are defined as follows: ```{ "start": 1, "end": 100, "fee": 75 }, { "start": 101, "end": 200, "fee": 100 }, }``` Then the following fees would be charged based on the total number of API calls (assuming the currency selected is `USD`): * 1 call costs $75 * 50 calls cost $75 * 150 calls cost $100 The number of API calls cannot exceed 200.
	ConsumptionPricingRates GoogleCloudApigeeV1RateRangeResponseArrayOutput `pulumi:"consumptionPricingRates"`
	// Pricing model used for consumption-based charges.
	ConsumptionPricingType pulumi.StringOutput `pulumi:"consumptionPricingType"`
	// Time that the rate plan was created in milliseconds since epoch.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Currency to be used for billing. Consists of a three-letter code as defined by the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) standard.
	CurrencyCode pulumi.StringOutput `pulumi:"currencyCode"`
	// Description of the rate plan.
	Description pulumi.StringOutput `pulumi:"description"`
	// Display name of the rate plan.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Time when the rate plan will expire in milliseconds since epoch. Set to 0 or `null` to indicate that the rate plan should never expire.
	EndTime pulumi.StringOutput `pulumi:"endTime"`
	// Frequency at which the fixed fee is charged.
	FixedFeeFrequency pulumi.IntOutput `pulumi:"fixedFeeFrequency"`
	// Fixed amount that is charged at a defined interval and billed in advance of use of the API product. The fee will be prorated for the first billing period.
	FixedRecurringFee GoogleTypeMoneyResponseOutput `pulumi:"fixedRecurringFee"`
	// Time the rate plan was last modified in milliseconds since epoch.
	LastModifiedAt pulumi.StringOutput `pulumi:"lastModifiedAt"`
	// Name of the rate plan.
	Name pulumi.StringOutput `pulumi:"name"`
	// Details of the revenue sharing model.
	RevenueShareRates GoogleCloudApigeeV1RevenueShareRangeResponseArrayOutput `pulumi:"revenueShareRates"`
	// Method used to calculate the revenue that is shared with developers.
	RevenueShareType pulumi.StringOutput `pulumi:"revenueShareType"`
	// Initial, one-time fee paid when purchasing the API product.
	SetupFee GoogleTypeMoneyResponseOutput `pulumi:"setupFee"`
	// Time when the rate plan becomes active in milliseconds since epoch.
	StartTime pulumi.StringOutput `pulumi:"startTime"`
	// Current state of the rate plan (draft or published).
	State pulumi.StringOutput `pulumi:"state"`
}

// NewRatePlan registers a new resource with the given unique name, arguments, and options.
func NewRatePlan(ctx *pulumi.Context,
	name string, args *RatePlanArgs, opts ...pulumi.ResourceOption) (*RatePlan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiproductId == nil {
		return nil, errors.New("invalid value for required argument 'ApiproductId'")
	}
	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	var resource RatePlan
	err := ctx.RegisterResource("google-native:apigee/v1:RatePlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRatePlan gets an existing RatePlan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRatePlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RatePlanState, opts ...pulumi.ResourceOption) (*RatePlan, error) {
	var resource RatePlan
	err := ctx.ReadResource("google-native:apigee/v1:RatePlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RatePlan resources.
type ratePlanState struct {
}

type RatePlanState struct {
}

func (RatePlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*ratePlanState)(nil)).Elem()
}

type ratePlanArgs struct {
	// Name of the API product that the rate plan is associated with.
	Apiproduct   *string `pulumi:"apiproduct"`
	ApiproductId string  `pulumi:"apiproductId"`
	// Frequency at which the customer will be billed.
	BillingPeriod *RatePlanBillingPeriod `pulumi:"billingPeriod"`
	// API call volume ranges and the fees charged when the total number of API calls is within a given range. The method used to calculate the final fee depends on the selected pricing model. For example, if the pricing model is `STAIRSTEP` and the ranges are defined as follows: ```{ "start": 1, "end": 100, "fee": 75 }, { "start": 101, "end": 200, "fee": 100 }, }``` Then the following fees would be charged based on the total number of API calls (assuming the currency selected is `USD`): * 1 call costs $75 * 50 calls cost $75 * 150 calls cost $100 The number of API calls cannot exceed 200.
	ConsumptionPricingRates []GoogleCloudApigeeV1RateRange `pulumi:"consumptionPricingRates"`
	// Pricing model used for consumption-based charges.
	ConsumptionPricingType *RatePlanConsumptionPricingType `pulumi:"consumptionPricingType"`
	// Currency to be used for billing. Consists of a three-letter code as defined by the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) standard.
	CurrencyCode *string `pulumi:"currencyCode"`
	// Description of the rate plan.
	Description *string `pulumi:"description"`
	// Display name of the rate plan.
	DisplayName *string `pulumi:"displayName"`
	// Time when the rate plan will expire in milliseconds since epoch. Set to 0 or `null` to indicate that the rate plan should never expire.
	EndTime *string `pulumi:"endTime"`
	// Frequency at which the fixed fee is charged.
	FixedFeeFrequency *int `pulumi:"fixedFeeFrequency"`
	// Fixed amount that is charged at a defined interval and billed in advance of use of the API product. The fee will be prorated for the first billing period.
	FixedRecurringFee *GoogleTypeMoney `pulumi:"fixedRecurringFee"`
	OrganizationId    string           `pulumi:"organizationId"`
	// Details of the revenue sharing model.
	RevenueShareRates []GoogleCloudApigeeV1RevenueShareRange `pulumi:"revenueShareRates"`
	// Method used to calculate the revenue that is shared with developers.
	RevenueShareType *RatePlanRevenueShareType `pulumi:"revenueShareType"`
	// Initial, one-time fee paid when purchasing the API product.
	SetupFee *GoogleTypeMoney `pulumi:"setupFee"`
	// Time when the rate plan becomes active in milliseconds since epoch.
	StartTime *string `pulumi:"startTime"`
	// Current state of the rate plan (draft or published).
	State *RatePlanStateEnum `pulumi:"state"`
}

// The set of arguments for constructing a RatePlan resource.
type RatePlanArgs struct {
	// Name of the API product that the rate plan is associated with.
	Apiproduct   pulumi.StringPtrInput
	ApiproductId pulumi.StringInput
	// Frequency at which the customer will be billed.
	BillingPeriod RatePlanBillingPeriodPtrInput
	// API call volume ranges and the fees charged when the total number of API calls is within a given range. The method used to calculate the final fee depends on the selected pricing model. For example, if the pricing model is `STAIRSTEP` and the ranges are defined as follows: ```{ "start": 1, "end": 100, "fee": 75 }, { "start": 101, "end": 200, "fee": 100 }, }``` Then the following fees would be charged based on the total number of API calls (assuming the currency selected is `USD`): * 1 call costs $75 * 50 calls cost $75 * 150 calls cost $100 The number of API calls cannot exceed 200.
	ConsumptionPricingRates GoogleCloudApigeeV1RateRangeArrayInput
	// Pricing model used for consumption-based charges.
	ConsumptionPricingType RatePlanConsumptionPricingTypePtrInput
	// Currency to be used for billing. Consists of a three-letter code as defined by the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) standard.
	CurrencyCode pulumi.StringPtrInput
	// Description of the rate plan.
	Description pulumi.StringPtrInput
	// Display name of the rate plan.
	DisplayName pulumi.StringPtrInput
	// Time when the rate plan will expire in milliseconds since epoch. Set to 0 or `null` to indicate that the rate plan should never expire.
	EndTime pulumi.StringPtrInput
	// Frequency at which the fixed fee is charged.
	FixedFeeFrequency pulumi.IntPtrInput
	// Fixed amount that is charged at a defined interval and billed in advance of use of the API product. The fee will be prorated for the first billing period.
	FixedRecurringFee GoogleTypeMoneyPtrInput
	OrganizationId    pulumi.StringInput
	// Details of the revenue sharing model.
	RevenueShareRates GoogleCloudApigeeV1RevenueShareRangeArrayInput
	// Method used to calculate the revenue that is shared with developers.
	RevenueShareType RatePlanRevenueShareTypePtrInput
	// Initial, one-time fee paid when purchasing the API product.
	SetupFee GoogleTypeMoneyPtrInput
	// Time when the rate plan becomes active in milliseconds since epoch.
	StartTime pulumi.StringPtrInput
	// Current state of the rate plan (draft or published).
	State RatePlanStateEnumPtrInput
}

func (RatePlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ratePlanArgs)(nil)).Elem()
}

type RatePlanInput interface {
	pulumi.Input

	ToRatePlanOutput() RatePlanOutput
	ToRatePlanOutputWithContext(ctx context.Context) RatePlanOutput
}

func (*RatePlan) ElementType() reflect.Type {
	return reflect.TypeOf((**RatePlan)(nil)).Elem()
}

func (i *RatePlan) ToRatePlanOutput() RatePlanOutput {
	return i.ToRatePlanOutputWithContext(context.Background())
}

func (i *RatePlan) ToRatePlanOutputWithContext(ctx context.Context) RatePlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RatePlanOutput)
}

type RatePlanOutput struct{ *pulumi.OutputState }

func (RatePlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RatePlan)(nil)).Elem()
}

func (o RatePlanOutput) ToRatePlanOutput() RatePlanOutput {
	return o
}

func (o RatePlanOutput) ToRatePlanOutputWithContext(ctx context.Context) RatePlanOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RatePlanInput)(nil)).Elem(), &RatePlan{})
	pulumi.RegisterOutputType(RatePlanOutput{})
}
