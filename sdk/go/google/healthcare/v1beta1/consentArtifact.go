// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new Consent artifact in the parent consent store.
type ConsentArtifact struct {
	pulumi.CustomResourceState

	// Optional. Screenshots, PDFs, or other binary information documenting the user's consent.
	ConsentContentScreenshots ImageResponseArrayOutput `pulumi:"consentContentScreenshots"`
	// Optional. An string indicating the version of the consent information shown to the user.
	ConsentContentVersion pulumi.StringOutput `pulumi:"consentContentVersion"`
	// Optional. A signature from a guardian.
	GuardianSignature SignatureResponseOutput `pulumi:"guardianSignature"`
	// Optional. Metadata associated with the Consent artifact. For example, the consent locale or user agent version.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// Resource name of the Consent artifact, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/consentArtifacts/{consent_artifact_id}`. Cannot be changed after creation.
	Name pulumi.StringOutput `pulumi:"name"`
	// User's UUID provided by the client.
	UserId pulumi.StringOutput `pulumi:"userId"`
	// Optional. User's signature.
	UserSignature SignatureResponseOutput `pulumi:"userSignature"`
	// Optional. A signature from a witness.
	WitnessSignature SignatureResponseOutput `pulumi:"witnessSignature"`
}

// NewConsentArtifact registers a new resource with the given unique name, arguments, and options.
func NewConsentArtifact(ctx *pulumi.Context,
	name string, args *ConsentArtifactArgs, opts ...pulumi.ResourceOption) (*ConsentArtifact, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConsentStoreId == nil {
		return nil, errors.New("invalid value for required argument 'ConsentStoreId'")
	}
	if args.DatasetId == nil {
		return nil, errors.New("invalid value for required argument 'DatasetId'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	var resource ConsentArtifact
	err := ctx.RegisterResource("google-native:healthcare/v1beta1:ConsentArtifact", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConsentArtifact gets an existing ConsentArtifact resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConsentArtifact(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConsentArtifactState, opts ...pulumi.ResourceOption) (*ConsentArtifact, error) {
	var resource ConsentArtifact
	err := ctx.ReadResource("google-native:healthcare/v1beta1:ConsentArtifact", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConsentArtifact resources.
type consentArtifactState struct {
}

type ConsentArtifactState struct {
}

func (ConsentArtifactState) ElementType() reflect.Type {
	return reflect.TypeOf((*consentArtifactState)(nil)).Elem()
}

type consentArtifactArgs struct {
	// Optional. Screenshots, PDFs, or other binary information documenting the user's consent.
	ConsentContentScreenshots []Image `pulumi:"consentContentScreenshots"`
	// Optional. An string indicating the version of the consent information shown to the user.
	ConsentContentVersion *string `pulumi:"consentContentVersion"`
	ConsentStoreId        string  `pulumi:"consentStoreId"`
	DatasetId             string  `pulumi:"datasetId"`
	// Optional. A signature from a guardian.
	GuardianSignature *Signature `pulumi:"guardianSignature"`
	Location          *string    `pulumi:"location"`
	// Optional. Metadata associated with the Consent artifact. For example, the consent locale or user agent version.
	Metadata map[string]string `pulumi:"metadata"`
	// Resource name of the Consent artifact, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/consentArtifacts/{consent_artifact_id}`. Cannot be changed after creation.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// User's UUID provided by the client.
	UserId string `pulumi:"userId"`
	// Optional. User's signature.
	UserSignature *Signature `pulumi:"userSignature"`
	// Optional. A signature from a witness.
	WitnessSignature *Signature `pulumi:"witnessSignature"`
}

// The set of arguments for constructing a ConsentArtifact resource.
type ConsentArtifactArgs struct {
	// Optional. Screenshots, PDFs, or other binary information documenting the user's consent.
	ConsentContentScreenshots ImageArrayInput
	// Optional. An string indicating the version of the consent information shown to the user.
	ConsentContentVersion pulumi.StringPtrInput
	ConsentStoreId        pulumi.StringInput
	DatasetId             pulumi.StringInput
	// Optional. A signature from a guardian.
	GuardianSignature SignaturePtrInput
	Location          pulumi.StringPtrInput
	// Optional. Metadata associated with the Consent artifact. For example, the consent locale or user agent version.
	Metadata pulumi.StringMapInput
	// Resource name of the Consent artifact, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/consentArtifacts/{consent_artifact_id}`. Cannot be changed after creation.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// User's UUID provided by the client.
	UserId pulumi.StringInput
	// Optional. User's signature.
	UserSignature SignaturePtrInput
	// Optional. A signature from a witness.
	WitnessSignature SignaturePtrInput
}

func (ConsentArtifactArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*consentArtifactArgs)(nil)).Elem()
}

type ConsentArtifactInput interface {
	pulumi.Input

	ToConsentArtifactOutput() ConsentArtifactOutput
	ToConsentArtifactOutputWithContext(ctx context.Context) ConsentArtifactOutput
}

func (*ConsentArtifact) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsentArtifact)(nil)).Elem()
}

func (i *ConsentArtifact) ToConsentArtifactOutput() ConsentArtifactOutput {
	return i.ToConsentArtifactOutputWithContext(context.Background())
}

func (i *ConsentArtifact) ToConsentArtifactOutputWithContext(ctx context.Context) ConsentArtifactOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsentArtifactOutput)
}

type ConsentArtifactOutput struct{ *pulumi.OutputState }

func (ConsentArtifactOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsentArtifact)(nil)).Elem()
}

func (o ConsentArtifactOutput) ToConsentArtifactOutput() ConsentArtifactOutput {
	return o
}

func (o ConsentArtifactOutput) ToConsentArtifactOutputWithContext(ctx context.Context) ConsentArtifactOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConsentArtifactInput)(nil)).Elem(), &ConsentArtifact{})
	pulumi.RegisterOutputType(ConsentArtifactOutput{})
}
