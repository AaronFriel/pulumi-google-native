// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta4

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an SSL certificate and returns it along with the private key and server certificate authority. The new certificate will not be usable until the instance is restarted.
// Auto-naming is currently not supported for this resource.
type SslCert struct {
	pulumi.CustomResourceState

	// PEM representation.
	Cert pulumi.StringOutput `pulumi:"cert"`
	// Serial number, as extracted from the certificate.
	CertSerialNumber pulumi.StringOutput `pulumi:"certSerialNumber"`
	// User supplied name. Constrained to [a-zA-Z.-_ ]+.
	CommonName pulumi.StringOutput `pulumi:"commonName"`
	// The time when the certificate was created in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example **2012-11-15T16:19:00.094Z**.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The time when the certificate expires in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example **2012-11-15T16:19:00.094Z**.
	ExpirationTime pulumi.StringOutput `pulumi:"expirationTime"`
	// Name of the database instance.
	Instance pulumi.StringOutput `pulumi:"instance"`
	// This is always **sql#sslCert**.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The URI of this resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Sha1 Fingerprint.
	Sha1Fingerprint pulumi.StringOutput `pulumi:"sha1Fingerprint"`
}

// NewSslCert registers a new resource with the given unique name, arguments, and options.
func NewSslCert(ctx *pulumi.Context,
	name string, args *SslCertArgs, opts ...pulumi.ResourceOption) (*SslCert, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Instance == nil {
		return nil, errors.New("invalid value for required argument 'Instance'")
	}
	var resource SslCert
	err := ctx.RegisterResource("google-native:sqladmin/v1beta4:SslCert", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSslCert gets an existing SslCert resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSslCert(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SslCertState, opts ...pulumi.ResourceOption) (*SslCert, error) {
	var resource SslCert
	err := ctx.ReadResource("google-native:sqladmin/v1beta4:SslCert", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SslCert resources.
type sslCertState struct {
}

type SslCertState struct {
}

func (SslCertState) ElementType() reflect.Type {
	return reflect.TypeOf((*sslCertState)(nil)).Elem()
}

type sslCertArgs struct {
	// User supplied name. Must be a distinct name from the other certificates for this instance.
	CommonName *string `pulumi:"commonName"`
	Instance   string  `pulumi:"instance"`
	Project    *string `pulumi:"project"`
}

// The set of arguments for constructing a SslCert resource.
type SslCertArgs struct {
	// User supplied name. Must be a distinct name from the other certificates for this instance.
	CommonName pulumi.StringPtrInput
	Instance   pulumi.StringInput
	Project    pulumi.StringPtrInput
}

func (SslCertArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sslCertArgs)(nil)).Elem()
}

type SslCertInput interface {
	pulumi.Input

	ToSslCertOutput() SslCertOutput
	ToSslCertOutputWithContext(ctx context.Context) SslCertOutput
}

func (*SslCert) ElementType() reflect.Type {
	return reflect.TypeOf((**SslCert)(nil)).Elem()
}

func (i *SslCert) ToSslCertOutput() SslCertOutput {
	return i.ToSslCertOutputWithContext(context.Background())
}

func (i *SslCert) ToSslCertOutputWithContext(ctx context.Context) SslCertOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslCertOutput)
}

type SslCertOutput struct{ *pulumi.OutputState }

func (SslCertOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SslCert)(nil)).Elem()
}

func (o SslCertOutput) ToSslCertOutput() SslCertOutput {
	return o
}

func (o SslCertOutput) ToSslCertOutputWithContext(ctx context.Context) SslCertOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SslCertInput)(nil)).Elem(), &SslCert{})
	pulumi.RegisterOutputType(SslCertOutput{})
}
