// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new DeliveryPipeline in a given project and location.
// Auto-naming is currently not supported for this resource.
type DeliveryPipeline struct {
	pulumi.CustomResourceState

	// User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
	Annotations pulumi.StringMapOutput `pulumi:"annotations"`
	// Information around the state of the Delivery Pipeline.
	Condition PipelineConditionResponseOutput `pulumi:"condition"`
	// Time at which the pipeline was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Description of the `DeliveryPipeline`. Max length is 255 characters.
	Description pulumi.StringOutput `pulumi:"description"`
	// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be <= 128 bytes.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Optional. Name of the `DeliveryPipeline`. Format is projects/{project}/ locations/{location}/deliveryPipelines/a-z{0,62}.
	Name pulumi.StringOutput `pulumi:"name"`
	// SerialPipeline defines a sequential set of stages for a `DeliveryPipeline`.
	SerialPipeline SerialPipelineResponseOutput `pulumi:"serialPipeline"`
	// Unique identifier of the `DeliveryPipeline`.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// Most recent time at which the pipeline was updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewDeliveryPipeline registers a new resource with the given unique name, arguments, and options.
func NewDeliveryPipeline(ctx *pulumi.Context,
	name string, args *DeliveryPipelineArgs, opts ...pulumi.ResourceOption) (*DeliveryPipeline, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeliveryPipelineId == nil {
		return nil, errors.New("invalid value for required argument 'DeliveryPipelineId'")
	}
	var resource DeliveryPipeline
	err := ctx.RegisterResource("google-native:clouddeploy/v1:DeliveryPipeline", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeliveryPipeline gets an existing DeliveryPipeline resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeliveryPipeline(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeliveryPipelineState, opts ...pulumi.ResourceOption) (*DeliveryPipeline, error) {
	var resource DeliveryPipeline
	err := ctx.ReadResource("google-native:clouddeploy/v1:DeliveryPipeline", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeliveryPipeline resources.
type deliveryPipelineState struct {
}

type DeliveryPipelineState struct {
}

func (DeliveryPipelineState) ElementType() reflect.Type {
	return reflect.TypeOf((*deliveryPipelineState)(nil)).Elem()
}

type deliveryPipelineArgs struct {
	// User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
	Annotations        map[string]string `pulumi:"annotations"`
	DeliveryPipelineId string            `pulumi:"deliveryPipelineId"`
	// Description of the `DeliveryPipeline`. Max length is 255 characters.
	Description *string `pulumi:"description"`
	// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag *string `pulumi:"etag"`
	// Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be <= 128 bytes.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// Optional. Name of the `DeliveryPipeline`. Format is projects/{project}/ locations/{location}/deliveryPipelines/a-z{0,62}.
	Name      *string `pulumi:"name"`
	Project   *string `pulumi:"project"`
	RequestId *string `pulumi:"requestId"`
	// SerialPipeline defines a sequential set of stages for a `DeliveryPipeline`.
	SerialPipeline *SerialPipeline `pulumi:"serialPipeline"`
	ValidateOnly   *string         `pulumi:"validateOnly"`
}

// The set of arguments for constructing a DeliveryPipeline resource.
type DeliveryPipelineArgs struct {
	// User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
	Annotations        pulumi.StringMapInput
	DeliveryPipelineId pulumi.StringInput
	// Description of the `DeliveryPipeline`. Max length is 255 characters.
	Description pulumi.StringPtrInput
	// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag pulumi.StringPtrInput
	// Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be <= 128 bytes.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// Optional. Name of the `DeliveryPipeline`. Format is projects/{project}/ locations/{location}/deliveryPipelines/a-z{0,62}.
	Name      pulumi.StringPtrInput
	Project   pulumi.StringPtrInput
	RequestId pulumi.StringPtrInput
	// SerialPipeline defines a sequential set of stages for a `DeliveryPipeline`.
	SerialPipeline SerialPipelinePtrInput
	ValidateOnly   pulumi.StringPtrInput
}

func (DeliveryPipelineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deliveryPipelineArgs)(nil)).Elem()
}

type DeliveryPipelineInput interface {
	pulumi.Input

	ToDeliveryPipelineOutput() DeliveryPipelineOutput
	ToDeliveryPipelineOutputWithContext(ctx context.Context) DeliveryPipelineOutput
}

func (*DeliveryPipeline) ElementType() reflect.Type {
	return reflect.TypeOf((**DeliveryPipeline)(nil)).Elem()
}

func (i *DeliveryPipeline) ToDeliveryPipelineOutput() DeliveryPipelineOutput {
	return i.ToDeliveryPipelineOutputWithContext(context.Background())
}

func (i *DeliveryPipeline) ToDeliveryPipelineOutputWithContext(ctx context.Context) DeliveryPipelineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryPipelineOutput)
}

type DeliveryPipelineOutput struct{ *pulumi.OutputState }

func (DeliveryPipelineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeliveryPipeline)(nil)).Elem()
}

func (o DeliveryPipelineOutput) ToDeliveryPipelineOutput() DeliveryPipelineOutput {
	return o
}

func (o DeliveryPipelineOutput) ToDeliveryPipelineOutputWithContext(ctx context.Context) DeliveryPipelineOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeliveryPipelineInput)(nil)).Elem(), &DeliveryPipeline{})
	pulumi.RegisterOutputType(DeliveryPipelineOutput{})
}
