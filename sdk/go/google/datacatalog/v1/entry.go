// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an entry. You can create entries only with 'FILESET', 'CLUSTER', 'DATA_STREAM', or custom types. Data Catalog automatically creates entries with other types during metadata ingestion from integrated systems. You must enable the Data Catalog API in the project identified by the `parent` parameter. For more information, see [Data Catalog resource project](https://cloud.google.com/data-catalog/docs/concepts/resource-project). An entry group can have a maximum of 100,000 entries.
// Auto-naming is currently not supported for this resource.
type Entry struct {
	pulumi.CustomResourceState

	// Specification for a group of BigQuery tables with the `[prefix]YYYYMMDD` name pattern. For more information, see [Introduction to partitioned tables] (https://cloud.google.com/bigquery/docs/partitioned-tables#partitioning_versus_sharding).
	BigqueryDateShardedSpec GoogleCloudDatacatalogV1BigQueryDateShardedSpecResponseOutput `pulumi:"bigqueryDateShardedSpec"`
	// Specification that applies to a BigQuery table. Valid only for entries with the `TABLE` type.
	BigqueryTableSpec GoogleCloudDatacatalogV1BigQueryTableSpecResponseOutput `pulumi:"bigqueryTableSpec"`
	// Business Context of the entry.
	BusinessContext GoogleCloudDatacatalogV1BusinessContextResponseOutput `pulumi:"businessContext"`
	// Physical location of the entry.
	DataSource GoogleCloudDatacatalogV1DataSourceResponseOutput `pulumi:"dataSource"`
	// Specification that applies to a data source connection. Valid only for entries with the `DATA_SOURCE_CONNECTION` type.
	DataSourceConnectionSpec GoogleCloudDatacatalogV1DataSourceConnectionSpecResponseOutput `pulumi:"dataSourceConnectionSpec"`
	// Specification that applies to a table resource. Valid only for entries with the `TABLE` type.
	DatabaseTableSpec GoogleCloudDatacatalogV1DatabaseTableSpecResponseOutput `pulumi:"databaseTableSpec"`
	// Entry description that can consist of several sentences or paragraphs that describe entry contents. The description must not contain Unicode non-characters as well as C0 and C1 control codes except tabs (HT), new lines (LF), carriage returns (CR), and page breaks (FF). The maximum size is 2000 bytes when encoded in UTF-8. Default value is an empty string.
	Description pulumi.StringOutput `pulumi:"description"`
	// Display name of an entry. The name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), and can't start or end with spaces. The maximum size is 200 bytes when encoded in UTF-8. Default value is an empty string.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Fully qualified name (FQN) of the resource. Set automatically for entries representing resources from synced systems. Settable only during creation and read-only afterwards. Can be used for search and lookup of the entries. FQNs take two forms: * For non-regionalized resources: `{SYSTEM}:{PROJECT}.{PATH_TO_RESOURCE_SEPARATED_WITH_DOTS}` * For regionalized resources: `{SYSTEM}:{PROJECT}.{LOCATION_ID}.{PATH_TO_RESOURCE_SEPARATED_WITH_DOTS}` Example for a DPMS table: `dataproc_metastore:{PROJECT_ID}.{LOCATION_ID}.{INSTANCE_ID}.{DATABASE_ID}.{TABLE_ID}`
	FullyQualifiedName pulumi.StringOutput `pulumi:"fullyQualifiedName"`
	// Specification that applies to a Cloud Storage fileset. Valid only for entries with the `FILESET` type.
	GcsFilesetSpec GoogleCloudDatacatalogV1GcsFilesetSpecResponseOutput `pulumi:"gcsFilesetSpec"`
	// Indicates the entry's source system that Data Catalog integrates with, such as BigQuery, Pub/Sub, or Dataproc Metastore.
	IntegratedSystem pulumi.StringOutput `pulumi:"integratedSystem"`
	// Cloud labels attached to the entry. In Data Catalog, you can create and modify labels attached only to custom entries. Synced entries have unmodifiable labels that come from the source system.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The resource this metadata entry refers to. For Google Cloud Platform resources, `linked_resource` is the [Full Resource Name] (https://cloud.google.com/apis/design/resource_names#full_resource_name). For example, the `linked_resource` for a table resource from BigQuery is: `//bigquery.googleapis.com/projects/{PROJECT_ID}/datasets/{DATASET_ID}/tables/{TABLE_ID}` Output only when the entry is one of the types in the `EntryType` enum. For entries with a `user_specified_type`, this field is optional and defaults to an empty string. The resource string must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), periods (.), colons (:), slashes (/), dashes (-), and hashes (#). The maximum size is 200 bytes when encoded in UTF-8.
	LinkedResource pulumi.StringOutput `pulumi:"linkedResource"`
	// The resource name of an entry in URL format. Note: The entry itself and its child resources might not be stored in the location specified in its name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Additional information related to the entry. Private to the current user.
	PersonalDetails GoogleCloudDatacatalogV1PersonalDetailsResponseOutput `pulumi:"personalDetails"`
	// Specification that applies to a user-defined function or procedure. Valid only for entries with the `ROUTINE` type.
	RoutineSpec GoogleCloudDatacatalogV1RoutineSpecResponseOutput `pulumi:"routineSpec"`
	// Schema of the entry. An entry might not have any schema attached to it.
	Schema GoogleCloudDatacatalogV1SchemaResponseOutput `pulumi:"schema"`
	// Timestamps from the underlying resource, not from the Data Catalog entry. Output only when the entry has a type listed in the `EntryType` enum. For entries with `user_specified_type`, this field is optional and defaults to an empty timestamp.
	SourceSystemTimestamps GoogleCloudDatacatalogV1SystemTimestampsResponseOutput `pulumi:"sourceSystemTimestamps"`
	// The type of the entry. Only used for entries with types listed in the `EntryType` enum. Currently, only `FILESET` enum value is allowed. All other entries created in Data Catalog must use the `user_specified_type`.
	Type pulumi.StringOutput `pulumi:"type"`
	// Resource usage statistics.
	UsageSignal GoogleCloudDatacatalogV1UsageSignalResponseOutput `pulumi:"usageSignal"`
	// Indicates the entry's source system that Data Catalog doesn't automatically integrate with. The `user_specified_system` string has the following limitations: * Is case insensitive. * Must begin with a letter or underscore. * Can only contain letters, numbers, and underscores. * Must be at least 1 character and at most 64 characters long.
	UserSpecifiedSystem pulumi.StringOutput `pulumi:"userSpecifiedSystem"`
	// Custom entry type that doesn't match any of the values allowed for input and listed in the `EntryType` enum. When creating an entry, first check the type values in the enum. If there are no appropriate types for the new entry, provide a custom value, for example, `my_special_type`. The `user_specified_type` string has the following limitations: * Is case insensitive. * Must begin with a letter or underscore. * Can only contain letters, numbers, and underscores. * Must be at least 1 character and at most 64 characters long.
	UserSpecifiedType pulumi.StringOutput `pulumi:"userSpecifiedType"`
}

// NewEntry registers a new resource with the given unique name, arguments, and options.
func NewEntry(ctx *pulumi.Context,
	name string, args *EntryArgs, opts ...pulumi.ResourceOption) (*Entry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EntryGroupId == nil {
		return nil, errors.New("invalid value for required argument 'EntryGroupId'")
	}
	if args.EntryId == nil {
		return nil, errors.New("invalid value for required argument 'EntryId'")
	}
	var resource Entry
	err := ctx.RegisterResource("google-native:datacatalog/v1:Entry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEntry gets an existing Entry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEntry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EntryState, opts ...pulumi.ResourceOption) (*Entry, error) {
	var resource Entry
	err := ctx.ReadResource("google-native:datacatalog/v1:Entry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Entry resources.
type entryState struct {
}

type EntryState struct {
}

func (EntryState) ElementType() reflect.Type {
	return reflect.TypeOf((*entryState)(nil)).Elem()
}

type entryArgs struct {
	// Specification for a group of BigQuery tables with the `[prefix]YYYYMMDD` name pattern. For more information, see [Introduction to partitioned tables] (https://cloud.google.com/bigquery/docs/partitioned-tables#partitioning_versus_sharding).
	BigqueryDateShardedSpec *GoogleCloudDatacatalogV1BigQueryDateShardedSpec `pulumi:"bigqueryDateShardedSpec"`
	// Specification that applies to a BigQuery table. Valid only for entries with the `TABLE` type.
	BigqueryTableSpec *GoogleCloudDatacatalogV1BigQueryTableSpec `pulumi:"bigqueryTableSpec"`
	// Business Context of the entry.
	BusinessContext *GoogleCloudDatacatalogV1BusinessContext `pulumi:"businessContext"`
	// Specification that applies to a data source connection. Valid only for entries with the `DATA_SOURCE_CONNECTION` type.
	DataSourceConnectionSpec *GoogleCloudDatacatalogV1DataSourceConnectionSpec `pulumi:"dataSourceConnectionSpec"`
	// Specification that applies to a table resource. Valid only for entries with the `TABLE` type.
	DatabaseTableSpec *GoogleCloudDatacatalogV1DatabaseTableSpec `pulumi:"databaseTableSpec"`
	// Entry description that can consist of several sentences or paragraphs that describe entry contents. The description must not contain Unicode non-characters as well as C0 and C1 control codes except tabs (HT), new lines (LF), carriage returns (CR), and page breaks (FF). The maximum size is 2000 bytes when encoded in UTF-8. Default value is an empty string.
	Description *string `pulumi:"description"`
	// Display name of an entry. The name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), and can't start or end with spaces. The maximum size is 200 bytes when encoded in UTF-8. Default value is an empty string.
	DisplayName  *string `pulumi:"displayName"`
	EntryGroupId string  `pulumi:"entryGroupId"`
	EntryId      string  `pulumi:"entryId"`
	// Fully qualified name (FQN) of the resource. Set automatically for entries representing resources from synced systems. Settable only during creation and read-only afterwards. Can be used for search and lookup of the entries. FQNs take two forms: * For non-regionalized resources: `{SYSTEM}:{PROJECT}.{PATH_TO_RESOURCE_SEPARATED_WITH_DOTS}` * For regionalized resources: `{SYSTEM}:{PROJECT}.{LOCATION_ID}.{PATH_TO_RESOURCE_SEPARATED_WITH_DOTS}` Example for a DPMS table: `dataproc_metastore:{PROJECT_ID}.{LOCATION_ID}.{INSTANCE_ID}.{DATABASE_ID}.{TABLE_ID}`
	FullyQualifiedName *string `pulumi:"fullyQualifiedName"`
	// Specification that applies to a Cloud Storage fileset. Valid only for entries with the `FILESET` type.
	GcsFilesetSpec *GoogleCloudDatacatalogV1GcsFilesetSpec `pulumi:"gcsFilesetSpec"`
	// Cloud labels attached to the entry. In Data Catalog, you can create and modify labels attached only to custom entries. Synced entries have unmodifiable labels that come from the source system.
	Labels map[string]string `pulumi:"labels"`
	// The resource this metadata entry refers to. For Google Cloud Platform resources, `linked_resource` is the [Full Resource Name] (https://cloud.google.com/apis/design/resource_names#full_resource_name). For example, the `linked_resource` for a table resource from BigQuery is: `//bigquery.googleapis.com/projects/{PROJECT_ID}/datasets/{DATASET_ID}/tables/{TABLE_ID}` Output only when the entry is one of the types in the `EntryType` enum. For entries with a `user_specified_type`, this field is optional and defaults to an empty string. The resource string must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), periods (.), colons (:), slashes (/), dashes (-), and hashes (#). The maximum size is 200 bytes when encoded in UTF-8.
	LinkedResource *string `pulumi:"linkedResource"`
	Location       *string `pulumi:"location"`
	Project        *string `pulumi:"project"`
	// Specification that applies to a user-defined function or procedure. Valid only for entries with the `ROUTINE` type.
	RoutineSpec *GoogleCloudDatacatalogV1RoutineSpec `pulumi:"routineSpec"`
	// Schema of the entry. An entry might not have any schema attached to it.
	Schema *GoogleCloudDatacatalogV1Schema `pulumi:"schema"`
	// Timestamps from the underlying resource, not from the Data Catalog entry. Output only when the entry has a type listed in the `EntryType` enum. For entries with `user_specified_type`, this field is optional and defaults to an empty timestamp.
	SourceSystemTimestamps *GoogleCloudDatacatalogV1SystemTimestamps `pulumi:"sourceSystemTimestamps"`
	// The type of the entry. Only used for entries with types listed in the `EntryType` enum. Currently, only `FILESET` enum value is allowed. All other entries created in Data Catalog must use the `user_specified_type`.
	Type *EntryType `pulumi:"type"`
	// Indicates the entry's source system that Data Catalog doesn't automatically integrate with. The `user_specified_system` string has the following limitations: * Is case insensitive. * Must begin with a letter or underscore. * Can only contain letters, numbers, and underscores. * Must be at least 1 character and at most 64 characters long.
	UserSpecifiedSystem *string `pulumi:"userSpecifiedSystem"`
	// Custom entry type that doesn't match any of the values allowed for input and listed in the `EntryType` enum. When creating an entry, first check the type values in the enum. If there are no appropriate types for the new entry, provide a custom value, for example, `my_special_type`. The `user_specified_type` string has the following limitations: * Is case insensitive. * Must begin with a letter or underscore. * Can only contain letters, numbers, and underscores. * Must be at least 1 character and at most 64 characters long.
	UserSpecifiedType *string `pulumi:"userSpecifiedType"`
}

// The set of arguments for constructing a Entry resource.
type EntryArgs struct {
	// Specification for a group of BigQuery tables with the `[prefix]YYYYMMDD` name pattern. For more information, see [Introduction to partitioned tables] (https://cloud.google.com/bigquery/docs/partitioned-tables#partitioning_versus_sharding).
	BigqueryDateShardedSpec GoogleCloudDatacatalogV1BigQueryDateShardedSpecPtrInput
	// Specification that applies to a BigQuery table. Valid only for entries with the `TABLE` type.
	BigqueryTableSpec GoogleCloudDatacatalogV1BigQueryTableSpecPtrInput
	// Business Context of the entry.
	BusinessContext GoogleCloudDatacatalogV1BusinessContextPtrInput
	// Specification that applies to a data source connection. Valid only for entries with the `DATA_SOURCE_CONNECTION` type.
	DataSourceConnectionSpec GoogleCloudDatacatalogV1DataSourceConnectionSpecPtrInput
	// Specification that applies to a table resource. Valid only for entries with the `TABLE` type.
	DatabaseTableSpec GoogleCloudDatacatalogV1DatabaseTableSpecPtrInput
	// Entry description that can consist of several sentences or paragraphs that describe entry contents. The description must not contain Unicode non-characters as well as C0 and C1 control codes except tabs (HT), new lines (LF), carriage returns (CR), and page breaks (FF). The maximum size is 2000 bytes when encoded in UTF-8. Default value is an empty string.
	Description pulumi.StringPtrInput
	// Display name of an entry. The name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), and can't start or end with spaces. The maximum size is 200 bytes when encoded in UTF-8. Default value is an empty string.
	DisplayName  pulumi.StringPtrInput
	EntryGroupId pulumi.StringInput
	EntryId      pulumi.StringInput
	// Fully qualified name (FQN) of the resource. Set automatically for entries representing resources from synced systems. Settable only during creation and read-only afterwards. Can be used for search and lookup of the entries. FQNs take two forms: * For non-regionalized resources: `{SYSTEM}:{PROJECT}.{PATH_TO_RESOURCE_SEPARATED_WITH_DOTS}` * For regionalized resources: `{SYSTEM}:{PROJECT}.{LOCATION_ID}.{PATH_TO_RESOURCE_SEPARATED_WITH_DOTS}` Example for a DPMS table: `dataproc_metastore:{PROJECT_ID}.{LOCATION_ID}.{INSTANCE_ID}.{DATABASE_ID}.{TABLE_ID}`
	FullyQualifiedName pulumi.StringPtrInput
	// Specification that applies to a Cloud Storage fileset. Valid only for entries with the `FILESET` type.
	GcsFilesetSpec GoogleCloudDatacatalogV1GcsFilesetSpecPtrInput
	// Cloud labels attached to the entry. In Data Catalog, you can create and modify labels attached only to custom entries. Synced entries have unmodifiable labels that come from the source system.
	Labels pulumi.StringMapInput
	// The resource this metadata entry refers to. For Google Cloud Platform resources, `linked_resource` is the [Full Resource Name] (https://cloud.google.com/apis/design/resource_names#full_resource_name). For example, the `linked_resource` for a table resource from BigQuery is: `//bigquery.googleapis.com/projects/{PROJECT_ID}/datasets/{DATASET_ID}/tables/{TABLE_ID}` Output only when the entry is one of the types in the `EntryType` enum. For entries with a `user_specified_type`, this field is optional and defaults to an empty string. The resource string must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), periods (.), colons (:), slashes (/), dashes (-), and hashes (#). The maximum size is 200 bytes when encoded in UTF-8.
	LinkedResource pulumi.StringPtrInput
	Location       pulumi.StringPtrInput
	Project        pulumi.StringPtrInput
	// Specification that applies to a user-defined function or procedure. Valid only for entries with the `ROUTINE` type.
	RoutineSpec GoogleCloudDatacatalogV1RoutineSpecPtrInput
	// Schema of the entry. An entry might not have any schema attached to it.
	Schema GoogleCloudDatacatalogV1SchemaPtrInput
	// Timestamps from the underlying resource, not from the Data Catalog entry. Output only when the entry has a type listed in the `EntryType` enum. For entries with `user_specified_type`, this field is optional and defaults to an empty timestamp.
	SourceSystemTimestamps GoogleCloudDatacatalogV1SystemTimestampsPtrInput
	// The type of the entry. Only used for entries with types listed in the `EntryType` enum. Currently, only `FILESET` enum value is allowed. All other entries created in Data Catalog must use the `user_specified_type`.
	Type EntryTypePtrInput
	// Indicates the entry's source system that Data Catalog doesn't automatically integrate with. The `user_specified_system` string has the following limitations: * Is case insensitive. * Must begin with a letter or underscore. * Can only contain letters, numbers, and underscores. * Must be at least 1 character and at most 64 characters long.
	UserSpecifiedSystem pulumi.StringPtrInput
	// Custom entry type that doesn't match any of the values allowed for input and listed in the `EntryType` enum. When creating an entry, first check the type values in the enum. If there are no appropriate types for the new entry, provide a custom value, for example, `my_special_type`. The `user_specified_type` string has the following limitations: * Is case insensitive. * Must begin with a letter or underscore. * Can only contain letters, numbers, and underscores. * Must be at least 1 character and at most 64 characters long.
	UserSpecifiedType pulumi.StringPtrInput
}

func (EntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*entryArgs)(nil)).Elem()
}

type EntryInput interface {
	pulumi.Input

	ToEntryOutput() EntryOutput
	ToEntryOutputWithContext(ctx context.Context) EntryOutput
}

func (*Entry) ElementType() reflect.Type {
	return reflect.TypeOf((*Entry)(nil))
}

func (i *Entry) ToEntryOutput() EntryOutput {
	return i.ToEntryOutputWithContext(context.Background())
}

func (i *Entry) ToEntryOutputWithContext(ctx context.Context) EntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntryOutput)
}

type EntryOutput struct{ *pulumi.OutputState }

func (EntryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Entry)(nil))
}

func (o EntryOutput) ToEntryOutput() EntryOutput {
	return o
}

func (o EntryOutput) ToEntryOutputWithContext(ctx context.Context) EntryOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EntryInput)(nil)).Elem(), &Entry{})
	pulumi.RegisterOutputType(EntryOutput{})
}