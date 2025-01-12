// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a conversation.
type Conversation struct {
	pulumi.CustomResourceState

	// An opaque, user-specified string representing the human agent who handled the conversation.
	AgentId pulumi.StringOutput `pulumi:"agentId"`
	// Call-specific metadata.
	CallMetadata GoogleCloudContactcenterinsightsV1ConversationCallMetadataResponseOutput `pulumi:"callMetadata"`
	// The time at which the conversation was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The source of the audio and transcription for the conversation.
	DataSource GoogleCloudContactcenterinsightsV1ConversationDataSourceResponseOutput `pulumi:"dataSource"`
	// All the matched Dialogflow intents in the call. The key corresponds to a Dialogflow intent, format: projects/{project}/agent/{agent}/intents/{intent}
	DialogflowIntents pulumi.StringMapOutput `pulumi:"dialogflowIntents"`
	// The duration of the conversation.
	Duration pulumi.StringOutput `pulumi:"duration"`
	// The time at which this conversation should expire. After this time, the conversation data and any associated analyses will be deleted.
	ExpireTime pulumi.StringOutput `pulumi:"expireTime"`
	// A map for the user to specify any custom fields. A maximum of 20 labels per conversation is allowed, with a maximum of 256 characters per entry.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// A user-specified language code for the conversation.
	LanguageCode pulumi.StringOutput `pulumi:"languageCode"`
	// The conversation's latest analysis, if one exists.
	LatestAnalysis GoogleCloudContactcenterinsightsV1AnalysisResponseOutput `pulumi:"latestAnalysis"`
	// Immutable. The conversation medium, if unspecified will default to PHONE_CALL.
	Medium pulumi.StringOutput `pulumi:"medium"`
	// Immutable. The resource name of the conversation. Format: projects/{project}/locations/{location}/conversations/{conversation}
	Name pulumi.StringOutput `pulumi:"name"`
	// Obfuscated user ID which the customer sent to us.
	ObfuscatedUserId pulumi.StringOutput `pulumi:"obfuscatedUserId"`
	// The annotations that were generated during the customer and agent interaction.
	RuntimeAnnotations GoogleCloudContactcenterinsightsV1RuntimeAnnotationResponseArrayOutput `pulumi:"runtimeAnnotations"`
	// The time at which the conversation started.
	StartTime pulumi.StringOutput `pulumi:"startTime"`
	// The conversation transcript.
	Transcript GoogleCloudContactcenterinsightsV1ConversationTranscriptResponseOutput `pulumi:"transcript"`
	// Input only. The TTL for this resource. If specified, then this TTL will be used to calculate the expire time.
	Ttl pulumi.StringOutput `pulumi:"ttl"`
	// The number of turns in the conversation.
	TurnCount pulumi.IntOutput `pulumi:"turnCount"`
	// The most recent time at which the conversation was updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewConversation registers a new resource with the given unique name, arguments, and options.
func NewConversation(ctx *pulumi.Context,
	name string, args *ConversationArgs, opts ...pulumi.ResourceOption) (*Conversation, error) {
	if args == nil {
		args = &ConversationArgs{}
	}

	var resource Conversation
	err := ctx.RegisterResource("google-native:contactcenterinsights/v1:Conversation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConversation gets an existing Conversation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConversation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConversationState, opts ...pulumi.ResourceOption) (*Conversation, error) {
	var resource Conversation
	err := ctx.ReadResource("google-native:contactcenterinsights/v1:Conversation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Conversation resources.
type conversationState struct {
}

type ConversationState struct {
}

func (ConversationState) ElementType() reflect.Type {
	return reflect.TypeOf((*conversationState)(nil)).Elem()
}

type conversationArgs struct {
	// An opaque, user-specified string representing the human agent who handled the conversation.
	AgentId *string `pulumi:"agentId"`
	// Call-specific metadata.
	CallMetadata   *GoogleCloudContactcenterinsightsV1ConversationCallMetadata `pulumi:"callMetadata"`
	ConversationId *string                                                     `pulumi:"conversationId"`
	// The source of the audio and transcription for the conversation.
	DataSource *GoogleCloudContactcenterinsightsV1ConversationDataSource `pulumi:"dataSource"`
	// The time at which this conversation should expire. After this time, the conversation data and any associated analyses will be deleted.
	ExpireTime *string `pulumi:"expireTime"`
	// A map for the user to specify any custom fields. A maximum of 20 labels per conversation is allowed, with a maximum of 256 characters per entry.
	Labels map[string]string `pulumi:"labels"`
	// A user-specified language code for the conversation.
	LanguageCode *string `pulumi:"languageCode"`
	Location     *string `pulumi:"location"`
	// Immutable. The conversation medium, if unspecified will default to PHONE_CALL.
	Medium *ConversationMedium `pulumi:"medium"`
	// Immutable. The resource name of the conversation. Format: projects/{project}/locations/{location}/conversations/{conversation}
	Name *string `pulumi:"name"`
	// Obfuscated user ID which the customer sent to us.
	ObfuscatedUserId *string `pulumi:"obfuscatedUserId"`
	Project          *string `pulumi:"project"`
	// The time at which the conversation started.
	StartTime *string `pulumi:"startTime"`
	// Input only. The TTL for this resource. If specified, then this TTL will be used to calculate the expire time.
	Ttl *string `pulumi:"ttl"`
}

// The set of arguments for constructing a Conversation resource.
type ConversationArgs struct {
	// An opaque, user-specified string representing the human agent who handled the conversation.
	AgentId pulumi.StringPtrInput
	// Call-specific metadata.
	CallMetadata   GoogleCloudContactcenterinsightsV1ConversationCallMetadataPtrInput
	ConversationId pulumi.StringPtrInput
	// The source of the audio and transcription for the conversation.
	DataSource GoogleCloudContactcenterinsightsV1ConversationDataSourcePtrInput
	// The time at which this conversation should expire. After this time, the conversation data and any associated analyses will be deleted.
	ExpireTime pulumi.StringPtrInput
	// A map for the user to specify any custom fields. A maximum of 20 labels per conversation is allowed, with a maximum of 256 characters per entry.
	Labels pulumi.StringMapInput
	// A user-specified language code for the conversation.
	LanguageCode pulumi.StringPtrInput
	Location     pulumi.StringPtrInput
	// Immutable. The conversation medium, if unspecified will default to PHONE_CALL.
	Medium ConversationMediumPtrInput
	// Immutable. The resource name of the conversation. Format: projects/{project}/locations/{location}/conversations/{conversation}
	Name pulumi.StringPtrInput
	// Obfuscated user ID which the customer sent to us.
	ObfuscatedUserId pulumi.StringPtrInput
	Project          pulumi.StringPtrInput
	// The time at which the conversation started.
	StartTime pulumi.StringPtrInput
	// Input only. The TTL for this resource. If specified, then this TTL will be used to calculate the expire time.
	Ttl pulumi.StringPtrInput
}

func (ConversationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*conversationArgs)(nil)).Elem()
}

type ConversationInput interface {
	pulumi.Input

	ToConversationOutput() ConversationOutput
	ToConversationOutputWithContext(ctx context.Context) ConversationOutput
}

func (*Conversation) ElementType() reflect.Type {
	return reflect.TypeOf((**Conversation)(nil)).Elem()
}

func (i *Conversation) ToConversationOutput() ConversationOutput {
	return i.ToConversationOutputWithContext(context.Background())
}

func (i *Conversation) ToConversationOutputWithContext(ctx context.Context) ConversationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConversationOutput)
}

type ConversationOutput struct{ *pulumi.OutputState }

func (ConversationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Conversation)(nil)).Elem()
}

func (o ConversationOutput) ToConversationOutput() ConversationOutput {
	return o
}

func (o ConversationOutput) ToConversationOutputWithContext(ctx context.Context) ConversationOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConversationInput)(nil)).Elem(), &Conversation{})
	pulumi.RegisterOutputType(ConversationOutput{})
}
