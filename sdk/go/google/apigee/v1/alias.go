// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an alias from a key/certificate pair. The structure of the request is controlled by the `format` query parameter: - `keycertfile` - Separate PEM-encoded key and certificate files are uploaded. Set `Content-Type: multipart/form-data` and include the `keyFile`, `certFile`, and `password` (if keys are encrypted) fields in the request body. If uploading to a truststore, omit `keyFile`. - `pkcs12` - A PKCS12 file is uploaded. Set `Content-Type: multipart/form-data`, provide the file in the `file` field, and include the `password` field if the file is encrypted in the request body. - `selfsignedcert` - A new private key and certificate are generated. Set `Content-Type: application/json` and include CertificateGenerationSpec in the request body.
// Auto-naming is currently not supported for this resource.
type Alias struct {
	pulumi.CustomResourceState

	// Resource ID for this alias. Values must match the regular expression `[^/]{1,255}`.
	Alias pulumi.StringOutput `pulumi:"alias"`
	// Chain of certificates under this alias.
	CertsInfo GoogleCloudApigeeV1CertificateResponseOutput `pulumi:"certsInfo"`
	// Type of alias.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAlias registers a new resource with the given unique name, arguments, and options.
func NewAlias(ctx *pulumi.Context,
	name string, args *AliasArgs, opts ...pulumi.ResourceOption) (*Alias, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvironmentId == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentId'")
	}
	if args.Format == nil {
		return nil, errors.New("invalid value for required argument 'Format'")
	}
	if args.KeystoreId == nil {
		return nil, errors.New("invalid value for required argument 'KeystoreId'")
	}
	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	var resource Alias
	err := ctx.RegisterResource("google-native:apigee/v1:Alias", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlias gets an existing Alias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AliasState, opts ...pulumi.ResourceOption) (*Alias, error) {
	var resource Alias
	err := ctx.ReadResource("google-native:apigee/v1:Alias", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Alias resources.
type aliasState struct {
}

type AliasState struct {
}

func (AliasState) ElementType() reflect.Type {
	return reflect.TypeOf((*aliasState)(nil)).Elem()
}

type aliasArgs struct {
	Alias *string `pulumi:"alias"`
	// The HTTP Content-Type header value specifying the content type of the body.
	ContentType *string `pulumi:"contentType"`
	// The HTTP request/response body as raw binary.
	Data          *string `pulumi:"data"`
	EnvironmentId string  `pulumi:"environmentId"`
	// Application specific response metadata. Must be set in the first response for streaming APIs.
	Extensions              []map[string]string `pulumi:"extensions"`
	Format                  string              `pulumi:"format"`
	IgnoreExpiryValidation  *string             `pulumi:"ignoreExpiryValidation"`
	IgnoreNewlineValidation *string             `pulumi:"ignoreNewlineValidation"`
	KeystoreId              string              `pulumi:"keystoreId"`
	OrganizationId          string              `pulumi:"organizationId"`
}

// The set of arguments for constructing a Alias resource.
type AliasArgs struct {
	Alias pulumi.StringPtrInput
	// The HTTP Content-Type header value specifying the content type of the body.
	ContentType pulumi.StringPtrInput
	// The HTTP request/response body as raw binary.
	Data          pulumi.StringPtrInput
	EnvironmentId pulumi.StringInput
	// Application specific response metadata. Must be set in the first response for streaming APIs.
	Extensions              pulumi.StringMapArrayInput
	Format                  pulumi.StringInput
	IgnoreExpiryValidation  pulumi.StringPtrInput
	IgnoreNewlineValidation pulumi.StringPtrInput
	KeystoreId              pulumi.StringInput
	OrganizationId          pulumi.StringInput
}

func (AliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aliasArgs)(nil)).Elem()
}

type AliasInput interface {
	pulumi.Input

	ToAliasOutput() AliasOutput
	ToAliasOutputWithContext(ctx context.Context) AliasOutput
}

func (*Alias) ElementType() reflect.Type {
	return reflect.TypeOf((**Alias)(nil)).Elem()
}

func (i *Alias) ToAliasOutput() AliasOutput {
	return i.ToAliasOutputWithContext(context.Background())
}

func (i *Alias) ToAliasOutputWithContext(ctx context.Context) AliasOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasOutput)
}

type AliasOutput struct{ *pulumi.OutputState }

func (AliasOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Alias)(nil)).Elem()
}

func (o AliasOutput) ToAliasOutput() AliasOutput {
	return o
}

func (o AliasOutput) ToAliasOutputWithContext(ctx context.Context) AliasOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AliasInput)(nil)).Elem(), &Alias{})
	pulumi.RegisterOutputType(AliasOutput{})
}
