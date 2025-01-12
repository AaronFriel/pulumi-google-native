// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create security settings in the specified location.
type SecuritySetting struct {
	pulumi.CustomResourceState

	// [DLP](https://cloud.google.com/dlp/docs) deidentify template name. Use this template to define de-identification configuration for the content. The `DLP De-identify Templates Reader` role is needed on the Dialogflow service identity service account (has the form `service-PROJECT_NUMBER@gcp-sa-dialogflow.iam.gserviceaccount.com`) for your agent's project. If empty, Dialogflow replaces sensitive info with `[redacted]` text. The template name will have one of the following formats: `projects//locations//deidentifyTemplates/` OR `organizations//locations//deidentifyTemplates/` Note: `deidentify_template` must be located in the same region as the `SecuritySettings`.
	DeidentifyTemplate pulumi.StringOutput `pulumi:"deidentifyTemplate"`
	// The human-readable name of the security settings, unique within the location.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Controls conversation exporting settings to Insights after conversation is completed. If retention_strategy is set to REMOVE_AFTER_CONVERSATION, Insights export is disabled no matter what you configure here.
	InsightsExportSettings GoogleCloudDialogflowCxV3SecuritySettingsInsightsExportSettingsResponseOutput `pulumi:"insightsExportSettings"`
	// [DLP](https://cloud.google.com/dlp/docs) inspect template name. Use this template to define inspect base settings. The `DLP Inspect Templates Reader` role is needed on the Dialogflow service identity service account (has the form `service-PROJECT_NUMBER@gcp-sa-dialogflow.iam.gserviceaccount.com`) for your agent's project. If empty, we use the default DLP inspect config. The template name will have one of the following formats: `projects//locations//inspectTemplates/` OR `organizations//locations//inspectTemplates/` Note: `inspect_template` must be located in the same region as the `SecuritySettings`.
	InspectTemplate pulumi.StringOutput `pulumi:"inspectTemplate"`
	// Resource name of the settings. Required for the SecuritySettingsService.UpdateSecuritySettings method. SecuritySettingsService.CreateSecuritySettings populates the name automatically. Format: `projects//locations//securitySettings/`.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of types of data to remove when retention settings triggers purge.
	PurgeDataTypes pulumi.StringArrayOutput `pulumi:"purgeDataTypes"`
	// Defines the data for which Dialogflow applies redaction. Dialogflow does not redact data that it does not have access to – for example, Cloud logging.
	RedactionScope pulumi.StringOutput `pulumi:"redactionScope"`
	// Strategy that defines how we do redaction.
	RedactionStrategy pulumi.StringOutput `pulumi:"redactionStrategy"`
	// Retains data in interaction logging for the specified number of days. This does not apply to Cloud logging, which is owned by the user - not Dialogflow. User must set a value lower than Dialogflow's default 365d TTL. Setting a value higher than that has no effect. A missing value or setting to 0 also means we use Dialogflow's default TTL. Note: Interaction logging is a limited access feature. Talk to your Google representative to check availability for you.
	RetentionWindowDays pulumi.IntOutput `pulumi:"retentionWindowDays"`
}

// NewSecuritySetting registers a new resource with the given unique name, arguments, and options.
func NewSecuritySetting(ctx *pulumi.Context,
	name string, args *SecuritySettingArgs, opts ...pulumi.ResourceOption) (*SecuritySetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	var resource SecuritySetting
	err := ctx.RegisterResource("google-native:dialogflow/v3:SecuritySetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecuritySetting gets an existing SecuritySetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecuritySetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecuritySettingState, opts ...pulumi.ResourceOption) (*SecuritySetting, error) {
	var resource SecuritySetting
	err := ctx.ReadResource("google-native:dialogflow/v3:SecuritySetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecuritySetting resources.
type securitySettingState struct {
}

type SecuritySettingState struct {
}

func (SecuritySettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*securitySettingState)(nil)).Elem()
}

type securitySettingArgs struct {
	// [DLP](https://cloud.google.com/dlp/docs) deidentify template name. Use this template to define de-identification configuration for the content. The `DLP De-identify Templates Reader` role is needed on the Dialogflow service identity service account (has the form `service-PROJECT_NUMBER@gcp-sa-dialogflow.iam.gserviceaccount.com`) for your agent's project. If empty, Dialogflow replaces sensitive info with `[redacted]` text. The template name will have one of the following formats: `projects//locations//deidentifyTemplates/` OR `organizations//locations//deidentifyTemplates/` Note: `deidentify_template` must be located in the same region as the `SecuritySettings`.
	DeidentifyTemplate *string `pulumi:"deidentifyTemplate"`
	// The human-readable name of the security settings, unique within the location.
	DisplayName string `pulumi:"displayName"`
	// Controls conversation exporting settings to Insights after conversation is completed. If retention_strategy is set to REMOVE_AFTER_CONVERSATION, Insights export is disabled no matter what you configure here.
	InsightsExportSettings *GoogleCloudDialogflowCxV3SecuritySettingsInsightsExportSettings `pulumi:"insightsExportSettings"`
	// [DLP](https://cloud.google.com/dlp/docs) inspect template name. Use this template to define inspect base settings. The `DLP Inspect Templates Reader` role is needed on the Dialogflow service identity service account (has the form `service-PROJECT_NUMBER@gcp-sa-dialogflow.iam.gserviceaccount.com`) for your agent's project. If empty, we use the default DLP inspect config. The template name will have one of the following formats: `projects//locations//inspectTemplates/` OR `organizations//locations//inspectTemplates/` Note: `inspect_template` must be located in the same region as the `SecuritySettings`.
	InspectTemplate *string `pulumi:"inspectTemplate"`
	Location        *string `pulumi:"location"`
	// Resource name of the settings. Required for the SecuritySettingsService.UpdateSecuritySettings method. SecuritySettingsService.CreateSecuritySettings populates the name automatically. Format: `projects//locations//securitySettings/`.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// List of types of data to remove when retention settings triggers purge.
	PurgeDataTypes []SecuritySettingPurgeDataTypesItem `pulumi:"purgeDataTypes"`
	// Defines the data for which Dialogflow applies redaction. Dialogflow does not redact data that it does not have access to – for example, Cloud logging.
	RedactionScope *SecuritySettingRedactionScope `pulumi:"redactionScope"`
	// Strategy that defines how we do redaction.
	RedactionStrategy *SecuritySettingRedactionStrategy `pulumi:"redactionStrategy"`
	// Retains data in interaction logging for the specified number of days. This does not apply to Cloud logging, which is owned by the user - not Dialogflow. User must set a value lower than Dialogflow's default 365d TTL. Setting a value higher than that has no effect. A missing value or setting to 0 also means we use Dialogflow's default TTL. Note: Interaction logging is a limited access feature. Talk to your Google representative to check availability for you.
	RetentionWindowDays *int `pulumi:"retentionWindowDays"`
}

// The set of arguments for constructing a SecuritySetting resource.
type SecuritySettingArgs struct {
	// [DLP](https://cloud.google.com/dlp/docs) deidentify template name. Use this template to define de-identification configuration for the content. The `DLP De-identify Templates Reader` role is needed on the Dialogflow service identity service account (has the form `service-PROJECT_NUMBER@gcp-sa-dialogflow.iam.gserviceaccount.com`) for your agent's project. If empty, Dialogflow replaces sensitive info with `[redacted]` text. The template name will have one of the following formats: `projects//locations//deidentifyTemplates/` OR `organizations//locations//deidentifyTemplates/` Note: `deidentify_template` must be located in the same region as the `SecuritySettings`.
	DeidentifyTemplate pulumi.StringPtrInput
	// The human-readable name of the security settings, unique within the location.
	DisplayName pulumi.StringInput
	// Controls conversation exporting settings to Insights after conversation is completed. If retention_strategy is set to REMOVE_AFTER_CONVERSATION, Insights export is disabled no matter what you configure here.
	InsightsExportSettings GoogleCloudDialogflowCxV3SecuritySettingsInsightsExportSettingsPtrInput
	// [DLP](https://cloud.google.com/dlp/docs) inspect template name. Use this template to define inspect base settings. The `DLP Inspect Templates Reader` role is needed on the Dialogflow service identity service account (has the form `service-PROJECT_NUMBER@gcp-sa-dialogflow.iam.gserviceaccount.com`) for your agent's project. If empty, we use the default DLP inspect config. The template name will have one of the following formats: `projects//locations//inspectTemplates/` OR `organizations//locations//inspectTemplates/` Note: `inspect_template` must be located in the same region as the `SecuritySettings`.
	InspectTemplate pulumi.StringPtrInput
	Location        pulumi.StringPtrInput
	// Resource name of the settings. Required for the SecuritySettingsService.UpdateSecuritySettings method. SecuritySettingsService.CreateSecuritySettings populates the name automatically. Format: `projects//locations//securitySettings/`.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// List of types of data to remove when retention settings triggers purge.
	PurgeDataTypes SecuritySettingPurgeDataTypesItemArrayInput
	// Defines the data for which Dialogflow applies redaction. Dialogflow does not redact data that it does not have access to – for example, Cloud logging.
	RedactionScope SecuritySettingRedactionScopePtrInput
	// Strategy that defines how we do redaction.
	RedactionStrategy SecuritySettingRedactionStrategyPtrInput
	// Retains data in interaction logging for the specified number of days. This does not apply to Cloud logging, which is owned by the user - not Dialogflow. User must set a value lower than Dialogflow's default 365d TTL. Setting a value higher than that has no effect. A missing value or setting to 0 also means we use Dialogflow's default TTL. Note: Interaction logging is a limited access feature. Talk to your Google representative to check availability for you.
	RetentionWindowDays pulumi.IntPtrInput
}

func (SecuritySettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securitySettingArgs)(nil)).Elem()
}

type SecuritySettingInput interface {
	pulumi.Input

	ToSecuritySettingOutput() SecuritySettingOutput
	ToSecuritySettingOutputWithContext(ctx context.Context) SecuritySettingOutput
}

func (*SecuritySetting) ElementType() reflect.Type {
	return reflect.TypeOf((**SecuritySetting)(nil)).Elem()
}

func (i *SecuritySetting) ToSecuritySettingOutput() SecuritySettingOutput {
	return i.ToSecuritySettingOutputWithContext(context.Background())
}

func (i *SecuritySetting) ToSecuritySettingOutputWithContext(ctx context.Context) SecuritySettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecuritySettingOutput)
}

type SecuritySettingOutput struct{ *pulumi.OutputState }

func (SecuritySettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecuritySetting)(nil)).Elem()
}

func (o SecuritySettingOutput) ToSecuritySettingOutput() SecuritySettingOutput {
	return o
}

func (o SecuritySettingOutput) ToSecuritySettingOutputWithContext(ctx context.Context) SecuritySettingOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecuritySettingInput)(nil)).Elem(), &SecuritySetting{})
	pulumi.RegisterOutputType(SecuritySettingOutput{})
}
