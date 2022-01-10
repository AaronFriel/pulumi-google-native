// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new reCAPTCHA Enterprise key.
type Key struct {
	pulumi.CustomResourceState

	// Settings for keys that can be used by Android apps.
	AndroidSettings GoogleCloudRecaptchaenterpriseV1AndroidKeySettingsResponseOutput `pulumi:"androidSettings"`
	// The timestamp corresponding to the creation of this Key.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Human-readable display name of this key. Modifiable by user.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Settings for keys that can be used by iOS apps.
	IosSettings GoogleCloudRecaptchaenterpriseV1IOSKeySettingsResponseOutput `pulumi:"iosSettings"`
	// See Creating and managing labels.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The resource name for the Key in the format "projects/{project}/keys/{key}".
	Name pulumi.StringOutput `pulumi:"name"`
	// Options for user acceptance testing.
	TestingOptions GoogleCloudRecaptchaenterpriseV1TestingOptionsResponseOutput `pulumi:"testingOptions"`
	// Settings for WAF
	WafSettings GoogleCloudRecaptchaenterpriseV1WafSettingsResponseOutput `pulumi:"wafSettings"`
	// Settings for keys that can be used by websites.
	WebSettings GoogleCloudRecaptchaenterpriseV1WebKeySettingsResponseOutput `pulumi:"webSettings"`
}

// NewKey registers a new resource with the given unique name, arguments, and options.
func NewKey(ctx *pulumi.Context,
	name string, args *KeyArgs, opts ...pulumi.ResourceOption) (*Key, error) {
	if args == nil {
		args = &KeyArgs{}
	}

	var resource Key
	err := ctx.RegisterResource("google-native:recaptchaenterprise/v1:Key", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKey gets an existing Key resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyState, opts ...pulumi.ResourceOption) (*Key, error) {
	var resource Key
	err := ctx.ReadResource("google-native:recaptchaenterprise/v1:Key", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Key resources.
type keyState struct {
}

type KeyState struct {
}

func (KeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyState)(nil)).Elem()
}

type keyArgs struct {
	// Settings for keys that can be used by Android apps.
	AndroidSettings *GoogleCloudRecaptchaenterpriseV1AndroidKeySettings `pulumi:"androidSettings"`
	// The timestamp corresponding to the creation of this Key.
	CreateTime *string `pulumi:"createTime"`
	// Human-readable display name of this key. Modifiable by user.
	DisplayName *string `pulumi:"displayName"`
	// Settings for keys that can be used by iOS apps.
	IosSettings *GoogleCloudRecaptchaenterpriseV1IOSKeySettings `pulumi:"iosSettings"`
	// See Creating and managing labels.
	Labels map[string]string `pulumi:"labels"`
	// The resource name for the Key in the format "projects/{project}/keys/{key}".
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Options for user acceptance testing.
	TestingOptions *GoogleCloudRecaptchaenterpriseV1TestingOptions `pulumi:"testingOptions"`
	// Settings for WAF
	WafSettings *GoogleCloudRecaptchaenterpriseV1WafSettings `pulumi:"wafSettings"`
	// Settings for keys that can be used by websites.
	WebSettings *GoogleCloudRecaptchaenterpriseV1WebKeySettings `pulumi:"webSettings"`
}

// The set of arguments for constructing a Key resource.
type KeyArgs struct {
	// Settings for keys that can be used by Android apps.
	AndroidSettings GoogleCloudRecaptchaenterpriseV1AndroidKeySettingsPtrInput
	// The timestamp corresponding to the creation of this Key.
	CreateTime pulumi.StringPtrInput
	// Human-readable display name of this key. Modifiable by user.
	DisplayName pulumi.StringPtrInput
	// Settings for keys that can be used by iOS apps.
	IosSettings GoogleCloudRecaptchaenterpriseV1IOSKeySettingsPtrInput
	// See Creating and managing labels.
	Labels pulumi.StringMapInput
	// The resource name for the Key in the format "projects/{project}/keys/{key}".
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Options for user acceptance testing.
	TestingOptions GoogleCloudRecaptchaenterpriseV1TestingOptionsPtrInput
	// Settings for WAF
	WafSettings GoogleCloudRecaptchaenterpriseV1WafSettingsPtrInput
	// Settings for keys that can be used by websites.
	WebSettings GoogleCloudRecaptchaenterpriseV1WebKeySettingsPtrInput
}

func (KeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyArgs)(nil)).Elem()
}

type KeyInput interface {
	pulumi.Input

	ToKeyOutput() KeyOutput
	ToKeyOutputWithContext(ctx context.Context) KeyOutput
}

func (*Key) ElementType() reflect.Type {
	return reflect.TypeOf((*Key)(nil))
}

func (i *Key) ToKeyOutput() KeyOutput {
	return i.ToKeyOutputWithContext(context.Background())
}

func (i *Key) ToKeyOutputWithContext(ctx context.Context) KeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyOutput)
}

type KeyOutput struct{ *pulumi.OutputState }

func (KeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Key)(nil))
}

func (o KeyOutput) ToKeyOutput() KeyOutput {
	return o
}

func (o KeyOutput) ToKeyOutputWithContext(ctx context.Context) KeyOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KeyInput)(nil)).Elem(), &Key{})
	pulumi.RegisterOutputType(KeyOutput{})
}