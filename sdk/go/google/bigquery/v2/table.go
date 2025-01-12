// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new, empty table in the dataset.
// Auto-naming is currently not supported for this resource.
type Table struct {
	pulumi.CustomResourceState

	// [Beta] Clustering specification for the table. Must be specified with partitioning, data in the table will be first partitioned and subsequently clustered.
	Clustering ClusteringResponseOutput `pulumi:"clustering"`
	// The time when this table was created, in milliseconds since the epoch.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// The default collation of the table.
	DefaultCollation pulumi.StringOutput `pulumi:"defaultCollation"`
	// [Optional] A user-friendly description of this table.
	Description pulumi.StringOutput `pulumi:"description"`
	// Custom encryption configuration (e.g., Cloud KMS keys).
	EncryptionConfiguration EncryptionConfigurationResponseOutput `pulumi:"encryptionConfiguration"`
	// A hash of the table metadata. Used to ensure there were no concurrent modifications to the resource when attempting an update. Not guaranteed to change when the table contents or the fields numRows, numBytes, numLongTermBytes or lastModifiedTime change.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// [Optional] The time when this table expires, in milliseconds since the epoch. If not present, the table will persist indefinitely. Expired tables will be deleted and their storage reclaimed. The defaultTableExpirationMs property of the encapsulating dataset can be used to set a default expirationTime on newly created tables.
	ExpirationTime pulumi.StringOutput `pulumi:"expirationTime"`
	// [Optional] Describes the data format, location, and other properties of a table stored outside of BigQuery. By defining these properties, the data source can then be queried as if it were a standard BigQuery table.
	ExternalDataConfiguration ExternalDataConfigurationResponseOutput `pulumi:"externalDataConfiguration"`
	// [Optional] A descriptive name for this table.
	FriendlyName pulumi.StringOutput `pulumi:"friendlyName"`
	// The type of the resource.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The labels associated with this table. You can use these to organize and group your tables. Label keys and values can be no longer than 63 characters, can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. Label values are optional. Label keys must start with a letter and each label in the list must have a different key.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The time when this table was last modified, in milliseconds since the epoch.
	LastModifiedTime pulumi.StringOutput `pulumi:"lastModifiedTime"`
	// The geographic location where the table resides. This value is inherited from the dataset.
	Location pulumi.StringOutput `pulumi:"location"`
	// [Optional] Materialized view definition.
	MaterializedView MaterializedViewDefinitionResponseOutput `pulumi:"materializedView"`
	// [Output-only, Beta] Present iff this table represents a ML model. Describes the training information for the model, and it is required to run 'PREDICT' queries.
	Model ModelDefinitionResponseOutput `pulumi:"model"`
	// The size of this table in bytes, excluding any data in the streaming buffer.
	NumBytes pulumi.StringOutput `pulumi:"numBytes"`
	// The number of bytes in the table that are considered "long-term storage".
	NumLongTermBytes pulumi.StringOutput `pulumi:"numLongTermBytes"`
	// [TrustedTester] The physical size of this table in bytes, excluding any data in the streaming buffer. This includes compression and storage used for time travel.
	NumPhysicalBytes pulumi.StringOutput `pulumi:"numPhysicalBytes"`
	// The number of rows of data in this table, excluding any data in the streaming buffer.
	NumRows pulumi.StringOutput `pulumi:"numRows"`
	// [TrustedTester] Range partitioning specification for this table. Only one of timePartitioning and rangePartitioning should be specified.
	RangePartitioning RangePartitioningResponseOutput `pulumi:"rangePartitioning"`
	// [Optional] If set to true, queries over this table require a partition filter that can be used for partition elimination to be specified.
	RequirePartitionFilter pulumi.BoolOutput `pulumi:"requirePartitionFilter"`
	// [Optional] Describes the schema of this table.
	Schema TableSchemaResponseOutput `pulumi:"schema"`
	// A URL that can be used to access this resource again.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Snapshot definition.
	SnapshotDefinition SnapshotDefinitionResponseOutput `pulumi:"snapshotDefinition"`
	// Contains information regarding this table's streaming buffer, if one is present. This field will be absent if the table is not being streamed to or if there is no data in the streaming buffer.
	StreamingBuffer StreamingbufferResponseOutput `pulumi:"streamingBuffer"`
	// [Required] Reference describing the ID of this table.
	TableReference TableReferenceResponseOutput `pulumi:"tableReference"`
	// Time-based partitioning specification for this table. Only one of timePartitioning and rangePartitioning should be specified.
	TimePartitioning TimePartitioningResponseOutput `pulumi:"timePartitioning"`
	// Describes the table type. The following values are supported: TABLE: A normal BigQuery table. VIEW: A virtual table defined by a SQL query. SNAPSHOT: An immutable, read-only table that is a copy of another table. [TrustedTester] MATERIALIZED_VIEW: SQL query whose result is persisted. EXTERNAL: A table that references data stored in an external storage system, such as Google Cloud Storage. The default value is TABLE.
	Type pulumi.StringOutput `pulumi:"type"`
	// [Optional] The view definition.
	View ViewDefinitionResponseOutput `pulumi:"view"`
}

// NewTable registers a new resource with the given unique name, arguments, and options.
func NewTable(ctx *pulumi.Context,
	name string, args *TableArgs, opts ...pulumi.ResourceOption) (*Table, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatasetId == nil {
		return nil, errors.New("invalid value for required argument 'DatasetId'")
	}
	var resource Table
	err := ctx.RegisterResource("google-native:bigquery/v2:Table", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTable gets an existing Table resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TableState, opts ...pulumi.ResourceOption) (*Table, error) {
	var resource Table
	err := ctx.ReadResource("google-native:bigquery/v2:Table", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Table resources.
type tableState struct {
}

type TableState struct {
}

func (TableState) ElementType() reflect.Type {
	return reflect.TypeOf((*tableState)(nil)).Elem()
}

type tableArgs struct {
	// [Beta] Clustering specification for the table. Must be specified with partitioning, data in the table will be first partitioned and subsequently clustered.
	Clustering *Clustering `pulumi:"clustering"`
	DatasetId  string      `pulumi:"datasetId"`
	// [Optional] A user-friendly description of this table.
	Description *string `pulumi:"description"`
	// Custom encryption configuration (e.g., Cloud KMS keys).
	EncryptionConfiguration *EncryptionConfiguration `pulumi:"encryptionConfiguration"`
	// [Optional] The time when this table expires, in milliseconds since the epoch. If not present, the table will persist indefinitely. Expired tables will be deleted and their storage reclaimed. The defaultTableExpirationMs property of the encapsulating dataset can be used to set a default expirationTime on newly created tables.
	ExpirationTime *string `pulumi:"expirationTime"`
	// [Optional] Describes the data format, location, and other properties of a table stored outside of BigQuery. By defining these properties, the data source can then be queried as if it were a standard BigQuery table.
	ExternalDataConfiguration *ExternalDataConfiguration `pulumi:"externalDataConfiguration"`
	// [Optional] A descriptive name for this table.
	FriendlyName *string `pulumi:"friendlyName"`
	// The labels associated with this table. You can use these to organize and group your tables. Label keys and values can be no longer than 63 characters, can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. Label values are optional. Label keys must start with a letter and each label in the list must have a different key.
	Labels map[string]string `pulumi:"labels"`
	// [Optional] Materialized view definition.
	MaterializedView *MaterializedViewDefinition `pulumi:"materializedView"`
	// [Output-only, Beta] Present iff this table represents a ML model. Describes the training information for the model, and it is required to run 'PREDICT' queries.
	Model   *ModelDefinition `pulumi:"model"`
	Project *string          `pulumi:"project"`
	// [TrustedTester] Range partitioning specification for this table. Only one of timePartitioning and rangePartitioning should be specified.
	RangePartitioning *RangePartitioning `pulumi:"rangePartitioning"`
	// [Optional] If set to true, queries over this table require a partition filter that can be used for partition elimination to be specified.
	RequirePartitionFilter *bool `pulumi:"requirePartitionFilter"`
	// [Optional] Describes the schema of this table.
	Schema *TableSchema `pulumi:"schema"`
	// [Required] Reference describing the ID of this table.
	TableReference *TableReference `pulumi:"tableReference"`
	// Time-based partitioning specification for this table. Only one of timePartitioning and rangePartitioning should be specified.
	TimePartitioning *TimePartitioning `pulumi:"timePartitioning"`
	// [Optional] The view definition.
	View *ViewDefinition `pulumi:"view"`
}

// The set of arguments for constructing a Table resource.
type TableArgs struct {
	// [Beta] Clustering specification for the table. Must be specified with partitioning, data in the table will be first partitioned and subsequently clustered.
	Clustering ClusteringPtrInput
	DatasetId  pulumi.StringInput
	// [Optional] A user-friendly description of this table.
	Description pulumi.StringPtrInput
	// Custom encryption configuration (e.g., Cloud KMS keys).
	EncryptionConfiguration EncryptionConfigurationPtrInput
	// [Optional] The time when this table expires, in milliseconds since the epoch. If not present, the table will persist indefinitely. Expired tables will be deleted and their storage reclaimed. The defaultTableExpirationMs property of the encapsulating dataset can be used to set a default expirationTime on newly created tables.
	ExpirationTime pulumi.StringPtrInput
	// [Optional] Describes the data format, location, and other properties of a table stored outside of BigQuery. By defining these properties, the data source can then be queried as if it were a standard BigQuery table.
	ExternalDataConfiguration ExternalDataConfigurationPtrInput
	// [Optional] A descriptive name for this table.
	FriendlyName pulumi.StringPtrInput
	// The labels associated with this table. You can use these to organize and group your tables. Label keys and values can be no longer than 63 characters, can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. Label values are optional. Label keys must start with a letter and each label in the list must have a different key.
	Labels pulumi.StringMapInput
	// [Optional] Materialized view definition.
	MaterializedView MaterializedViewDefinitionPtrInput
	// [Output-only, Beta] Present iff this table represents a ML model. Describes the training information for the model, and it is required to run 'PREDICT' queries.
	Model   ModelDefinitionPtrInput
	Project pulumi.StringPtrInput
	// [TrustedTester] Range partitioning specification for this table. Only one of timePartitioning and rangePartitioning should be specified.
	RangePartitioning RangePartitioningPtrInput
	// [Optional] If set to true, queries over this table require a partition filter that can be used for partition elimination to be specified.
	RequirePartitionFilter pulumi.BoolPtrInput
	// [Optional] Describes the schema of this table.
	Schema TableSchemaPtrInput
	// [Required] Reference describing the ID of this table.
	TableReference TableReferencePtrInput
	// Time-based partitioning specification for this table. Only one of timePartitioning and rangePartitioning should be specified.
	TimePartitioning TimePartitioningPtrInput
	// [Optional] The view definition.
	View ViewDefinitionPtrInput
}

func (TableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tableArgs)(nil)).Elem()
}

type TableInput interface {
	pulumi.Input

	ToTableOutput() TableOutput
	ToTableOutputWithContext(ctx context.Context) TableOutput
}

func (*Table) ElementType() reflect.Type {
	return reflect.TypeOf((**Table)(nil)).Elem()
}

func (i *Table) ToTableOutput() TableOutput {
	return i.ToTableOutputWithContext(context.Background())
}

func (i *Table) ToTableOutputWithContext(ctx context.Context) TableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableOutput)
}

type TableOutput struct{ *pulumi.OutputState }

func (TableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Table)(nil)).Elem()
}

func (o TableOutput) ToTableOutput() TableOutput {
	return o
}

func (o TableOutput) ToTableOutputWithContext(ctx context.Context) TableOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TableInput)(nil)).Elem(), &Table{})
	pulumi.RegisterOutputType(TableOutput{})
}
