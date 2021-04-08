// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new `BuildTrigger`. This API is experimental.
type Trigger struct {
	pulumi.CustomResourceState

	// Contents of the build template.
	Build BuildResponseOutput `pulumi:"build"`
	// Time when the trigger was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Human-readable description of this trigger.
	Description pulumi.StringOutput `pulumi:"description"`
	// If true, the trigger will never automatically execute a build.
	Disabled pulumi.BoolOutput `pulumi:"disabled"`
	// Path, from the source root, to the build configuration file (i.e. cloudbuild.yaml).
	Filename pulumi.StringOutput `pulumi:"filename"`
	// GitHubEventsConfig describes the configuration of a trigger that creates a build whenever a GitHub event is received. Mutually exclusive with `trigger_template`.
	Github GitHubEventsConfigResponseOutput `pulumi:"github"`
	// ignored_files and included_files are file glob matches using https://golang.org/pkg/path/filepath/#Match extended with support for "**". If ignored_files and changed files are both empty, then they are not used to determine whether or not to trigger a build. If ignored_files is not empty, then we ignore any files that match any of the ignored_file globs. If the change has no files that are outside of the ignored_files globs, then we do not trigger a build.
	IgnoredFiles pulumi.StringArrayOutput `pulumi:"ignoredFiles"`
	// If any of the files altered in the commit pass the ignored_files filter and included_files is empty, then as far as this filter is concerned, we should trigger the build. If any of the files altered in the commit pass the ignored_files filter and included_files is not empty, then we make sure that at least one of those files matches a included_files glob. If not, then we do not trigger a build.
	IncludedFiles pulumi.StringArrayOutput `pulumi:"includedFiles"`
	// User-assigned name of the trigger. Must be unique within the project. Trigger names must meet the following requirements: + They must contain only alphanumeric characters and dashes. + They can be 1-64 characters long. + They must begin and end with an alphanumeric character.
	Name pulumi.StringOutput `pulumi:"name"`
	// Substitutions for Build resource. The keys must match the following regular expression: `^_[A-Z0-9_]+$`.
	Substitutions pulumi.StringMapOutput `pulumi:"substitutions"`
	// Tags for annotation of a `BuildTrigger`
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Template describing the types of source changes to trigger a build. Branch and tag names in trigger templates are interpreted as regular expressions. Any branch or tag change that matches that regular expression will trigger a build. Mutually exclusive with `github`.
	TriggerTemplate RepoSourceResponseOutput `pulumi:"triggerTemplate"`
}

// NewTrigger registers a new resource with the given unique name, arguments, and options.
func NewTrigger(ctx *pulumi.Context,
	name string, args *TriggerArgs, opts ...pulumi.ResourceOption) (*Trigger, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.TriggerId == nil {
		return nil, errors.New("invalid value for required argument 'TriggerId'")
	}
	var resource Trigger
	err := ctx.RegisterResource("gcp-native:cloudbuild/v1:Trigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrigger gets an existing Trigger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TriggerState, opts ...pulumi.ResourceOption) (*Trigger, error) {
	var resource Trigger
	err := ctx.ReadResource("gcp-native:cloudbuild/v1:Trigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Trigger resources.
type triggerState struct {
	// Contents of the build template.
	Build *BuildResponse `pulumi:"build"`
	// Time when the trigger was created.
	CreateTime *string `pulumi:"createTime"`
	// Human-readable description of this trigger.
	Description *string `pulumi:"description"`
	// If true, the trigger will never automatically execute a build.
	Disabled *bool `pulumi:"disabled"`
	// Path, from the source root, to the build configuration file (i.e. cloudbuild.yaml).
	Filename *string `pulumi:"filename"`
	// GitHubEventsConfig describes the configuration of a trigger that creates a build whenever a GitHub event is received. Mutually exclusive with `trigger_template`.
	Github *GitHubEventsConfigResponse `pulumi:"github"`
	// ignored_files and included_files are file glob matches using https://golang.org/pkg/path/filepath/#Match extended with support for "**". If ignored_files and changed files are both empty, then they are not used to determine whether or not to trigger a build. If ignored_files is not empty, then we ignore any files that match any of the ignored_file globs. If the change has no files that are outside of the ignored_files globs, then we do not trigger a build.
	IgnoredFiles []string `pulumi:"ignoredFiles"`
	// If any of the files altered in the commit pass the ignored_files filter and included_files is empty, then as far as this filter is concerned, we should trigger the build. If any of the files altered in the commit pass the ignored_files filter and included_files is not empty, then we make sure that at least one of those files matches a included_files glob. If not, then we do not trigger a build.
	IncludedFiles []string `pulumi:"includedFiles"`
	// User-assigned name of the trigger. Must be unique within the project. Trigger names must meet the following requirements: + They must contain only alphanumeric characters and dashes. + They can be 1-64 characters long. + They must begin and end with an alphanumeric character.
	Name *string `pulumi:"name"`
	// Substitutions for Build resource. The keys must match the following regular expression: `^_[A-Z0-9_]+$`.
	Substitutions map[string]string `pulumi:"substitutions"`
	// Tags for annotation of a `BuildTrigger`
	Tags []string `pulumi:"tags"`
	// Template describing the types of source changes to trigger a build. Branch and tag names in trigger templates are interpreted as regular expressions. Any branch or tag change that matches that regular expression will trigger a build. Mutually exclusive with `github`.
	TriggerTemplate *RepoSourceResponse `pulumi:"triggerTemplate"`
}

type TriggerState struct {
	// Contents of the build template.
	Build BuildResponsePtrInput
	// Time when the trigger was created.
	CreateTime pulumi.StringPtrInput
	// Human-readable description of this trigger.
	Description pulumi.StringPtrInput
	// If true, the trigger will never automatically execute a build.
	Disabled pulumi.BoolPtrInput
	// Path, from the source root, to the build configuration file (i.e. cloudbuild.yaml).
	Filename pulumi.StringPtrInput
	// GitHubEventsConfig describes the configuration of a trigger that creates a build whenever a GitHub event is received. Mutually exclusive with `trigger_template`.
	Github GitHubEventsConfigResponsePtrInput
	// ignored_files and included_files are file glob matches using https://golang.org/pkg/path/filepath/#Match extended with support for "**". If ignored_files and changed files are both empty, then they are not used to determine whether or not to trigger a build. If ignored_files is not empty, then we ignore any files that match any of the ignored_file globs. If the change has no files that are outside of the ignored_files globs, then we do not trigger a build.
	IgnoredFiles pulumi.StringArrayInput
	// If any of the files altered in the commit pass the ignored_files filter and included_files is empty, then as far as this filter is concerned, we should trigger the build. If any of the files altered in the commit pass the ignored_files filter and included_files is not empty, then we make sure that at least one of those files matches a included_files glob. If not, then we do not trigger a build.
	IncludedFiles pulumi.StringArrayInput
	// User-assigned name of the trigger. Must be unique within the project. Trigger names must meet the following requirements: + They must contain only alphanumeric characters and dashes. + They can be 1-64 characters long. + They must begin and end with an alphanumeric character.
	Name pulumi.StringPtrInput
	// Substitutions for Build resource. The keys must match the following regular expression: `^_[A-Z0-9_]+$`.
	Substitutions pulumi.StringMapInput
	// Tags for annotation of a `BuildTrigger`
	Tags pulumi.StringArrayInput
	// Template describing the types of source changes to trigger a build. Branch and tag names in trigger templates are interpreted as regular expressions. Any branch or tag change that matches that regular expression will trigger a build. Mutually exclusive with `github`.
	TriggerTemplate RepoSourceResponsePtrInput
}

func (TriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerState)(nil)).Elem()
}

type triggerArgs struct {
	// Contents of the build template.
	Build *BuildType `pulumi:"build"`
	// Human-readable description of this trigger.
	Description *string `pulumi:"description"`
	// If true, the trigger will never automatically execute a build.
	Disabled *bool `pulumi:"disabled"`
	// Path, from the source root, to the build configuration file (i.e. cloudbuild.yaml).
	Filename *string `pulumi:"filename"`
	// GitHubEventsConfig describes the configuration of a trigger that creates a build whenever a GitHub event is received. Mutually exclusive with `trigger_template`.
	Github *GitHubEventsConfig `pulumi:"github"`
	// ignored_files and included_files are file glob matches using https://golang.org/pkg/path/filepath/#Match extended with support for "**". If ignored_files and changed files are both empty, then they are not used to determine whether or not to trigger a build. If ignored_files is not empty, then we ignore any files that match any of the ignored_file globs. If the change has no files that are outside of the ignored_files globs, then we do not trigger a build.
	IgnoredFiles []string `pulumi:"ignoredFiles"`
	// If any of the files altered in the commit pass the ignored_files filter and included_files is empty, then as far as this filter is concerned, we should trigger the build. If any of the files altered in the commit pass the ignored_files filter and included_files is not empty, then we make sure that at least one of those files matches a included_files glob. If not, then we do not trigger a build.
	IncludedFiles []string `pulumi:"includedFiles"`
	// User-assigned name of the trigger. Must be unique within the project. Trigger names must meet the following requirements: + They must contain only alphanumeric characters and dashes. + They can be 1-64 characters long. + They must begin and end with an alphanumeric character.
	Name      *string `pulumi:"name"`
	ProjectId string  `pulumi:"projectId"`
	// Substitutions for Build resource. The keys must match the following regular expression: `^_[A-Z0-9_]+$`.
	Substitutions map[string]string `pulumi:"substitutions"`
	// Tags for annotation of a `BuildTrigger`
	Tags      []string `pulumi:"tags"`
	TriggerId string   `pulumi:"triggerId"`
	// Template describing the types of source changes to trigger a build. Branch and tag names in trigger templates are interpreted as regular expressions. Any branch or tag change that matches that regular expression will trigger a build. Mutually exclusive with `github`.
	TriggerTemplate *RepoSource `pulumi:"triggerTemplate"`
}

// The set of arguments for constructing a Trigger resource.
type TriggerArgs struct {
	// Contents of the build template.
	Build BuildTypePtrInput
	// Human-readable description of this trigger.
	Description pulumi.StringPtrInput
	// If true, the trigger will never automatically execute a build.
	Disabled pulumi.BoolPtrInput
	// Path, from the source root, to the build configuration file (i.e. cloudbuild.yaml).
	Filename pulumi.StringPtrInput
	// GitHubEventsConfig describes the configuration of a trigger that creates a build whenever a GitHub event is received. Mutually exclusive with `trigger_template`.
	Github GitHubEventsConfigPtrInput
	// ignored_files and included_files are file glob matches using https://golang.org/pkg/path/filepath/#Match extended with support for "**". If ignored_files and changed files are both empty, then they are not used to determine whether or not to trigger a build. If ignored_files is not empty, then we ignore any files that match any of the ignored_file globs. If the change has no files that are outside of the ignored_files globs, then we do not trigger a build.
	IgnoredFiles pulumi.StringArrayInput
	// If any of the files altered in the commit pass the ignored_files filter and included_files is empty, then as far as this filter is concerned, we should trigger the build. If any of the files altered in the commit pass the ignored_files filter and included_files is not empty, then we make sure that at least one of those files matches a included_files glob. If not, then we do not trigger a build.
	IncludedFiles pulumi.StringArrayInput
	// User-assigned name of the trigger. Must be unique within the project. Trigger names must meet the following requirements: + They must contain only alphanumeric characters and dashes. + They can be 1-64 characters long. + They must begin and end with an alphanumeric character.
	Name      pulumi.StringPtrInput
	ProjectId pulumi.StringInput
	// Substitutions for Build resource. The keys must match the following regular expression: `^_[A-Z0-9_]+$`.
	Substitutions pulumi.StringMapInput
	// Tags for annotation of a `BuildTrigger`
	Tags      pulumi.StringArrayInput
	TriggerId pulumi.StringInput
	// Template describing the types of source changes to trigger a build. Branch and tag names in trigger templates are interpreted as regular expressions. Any branch or tag change that matches that regular expression will trigger a build. Mutually exclusive with `github`.
	TriggerTemplate RepoSourcePtrInput
}

func (TriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerArgs)(nil)).Elem()
}

type TriggerInput interface {
	pulumi.Input

	ToTriggerOutput() TriggerOutput
	ToTriggerOutputWithContext(ctx context.Context) TriggerOutput
}

func (*Trigger) ElementType() reflect.Type {
	return reflect.TypeOf((*Trigger)(nil))
}

func (i *Trigger) ToTriggerOutput() TriggerOutput {
	return i.ToTriggerOutputWithContext(context.Background())
}

func (i *Trigger) ToTriggerOutputWithContext(ctx context.Context) TriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerOutput)
}

type TriggerOutput struct {
	*pulumi.OutputState
}

func (TriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Trigger)(nil))
}

func (o TriggerOutput) ToTriggerOutput() TriggerOutput {
	return o
}

func (o TriggerOutput) ToTriggerOutputWithContext(ctx context.Context) TriggerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TriggerOutput{})
}