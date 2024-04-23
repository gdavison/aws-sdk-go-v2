// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Use to specify a automatic model evaluation job. The
// EvaluationDatasetMetricConfig object is used to specify the prompt datasets,
// task type, and metric names.
type AutomatedEvaluationConfig struct {

	// Specifies the required elements for an automatic model evaluation job.
	//
	// This member is required.
	DatasetMetricConfigs []EvaluationDatasetMetricConfig

	noSmithyDocumentSerde
}

// CloudWatch logging configuration.
type CloudWatchConfig struct {

	// The log group name.
	//
	// This member is required.
	LogGroupName *string

	// The role Amazon Resource Name (ARN).
	//
	// This member is required.
	RoleArn *string

	// S3 configuration for delivering a large amount of data.
	LargeDataDeliveryS3Config *S3Config

	noSmithyDocumentSerde
}

// Summary information for a custom model.
type CustomModelSummary struct {

	// The base model Amazon Resource Name (ARN).
	//
	// This member is required.
	BaseModelArn *string

	// The base model name.
	//
	// This member is required.
	BaseModelName *string

	// Creation time of the model.
	//
	// This member is required.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) of the custom model.
	//
	// This member is required.
	ModelArn *string

	// The name of the custom model.
	//
	// This member is required.
	ModelName *string

	// Specifies whether to carry out continued pre-training of a model or whether to
	// fine-tune it. For more information, see Custom models (https://docs.aws.amazon.com/bedrock/latest/userguide/custom-models.html)
	// .
	CustomizationType CustomizationType

	noSmithyDocumentSerde
}

// Contains the ARN of the Amazon Bedrock models specified in your model
// evaluation job. Each Amazon Bedrock model supports different inferenceParams .
// To learn more about supported inference parameters for Amazon Bedrock models,
// see Inference parameters for foundation models (https://docs.aws.amazon.com/bedrock/latest/userguide/model-evaluation-prompt-datasets-custom.html)
// . The inferenceParams are specified using JSON. To successfully insert JSON as
// string make sure that all quotations are properly escaped. For example,
// "temperature":"0.25" key value pair would need to be formatted as
// \"temperature\":\"0.25\" to successfully accepted in the request.
type EvaluationBedrockModel struct {

	// Each Amazon Bedrock support different inference parameters that change how the
	// model behaves during inference.
	//
	// This member is required.
	InferenceParams *string

	// The ARN of the Amazon Bedrock model specified.
	//
	// This member is required.
	ModelIdentifier *string

	noSmithyDocumentSerde
}

// Used to specify either a AutomatedEvaluationConfig or HumanEvaluationConfig
// object.
//
// The following types satisfy this interface:
//
//	EvaluationConfigMemberAutomated
//	EvaluationConfigMemberHuman
type EvaluationConfig interface {
	isEvaluationConfig()
}

// Used to specify an automated model evaluation job. See AutomatedEvaluationConfig
// to view the required parameters.
type EvaluationConfigMemberAutomated struct {
	Value AutomatedEvaluationConfig

	noSmithyDocumentSerde
}

func (*EvaluationConfigMemberAutomated) isEvaluationConfig() {}

// Used to specify a model evaluation job that uses human workers.See
// HumanEvaluationConfig to view the required parameters.
type EvaluationConfigMemberHuman struct {
	Value HumanEvaluationConfig

	noSmithyDocumentSerde
}

func (*EvaluationConfigMemberHuman) isEvaluationConfig() {}

// Used to specify the name of a built-in prompt dataset and optionally, the
// Amazon S3 bucket where a custom prompt dataset is saved.
type EvaluationDataset struct {

	// Used to specify supported built-in prompt datasets. Valid values are
	// Builtin.Bold , Builtin.BoolQ , Builtin.NaturalQuestions , Builtin.Gigaword ,
	// Builtin.RealToxicityPrompts , Builtin.TriviaQa , Builtin.T-Rex ,
	// Builtin.WomensEcommerceClothingReviews and Builtin.Wikitext2 .
	//
	// This member is required.
	Name *string

	// For custom prompt datasets, you must specify the location in Amazon S3 where
	// the prompt dataset is saved.
	DatasetLocation EvaluationDatasetLocation

	noSmithyDocumentSerde
}

// The location in Amazon S3 where your prompt dataset is stored.
//
// The following types satisfy this interface:
//
//	EvaluationDatasetLocationMemberS3Uri
type EvaluationDatasetLocation interface {
	isEvaluationDatasetLocation()
}

// The S3 URI of the S3 bucket specified in the job.
type EvaluationDatasetLocationMemberS3Uri struct {
	Value string

	noSmithyDocumentSerde
}

func (*EvaluationDatasetLocationMemberS3Uri) isEvaluationDatasetLocation() {}

// Defines the built-in prompt datasets, built-in metric names and custom metric
// names, and the task type.
type EvaluationDatasetMetricConfig struct {

	// Specifies the prompt dataset.
	//
	// This member is required.
	Dataset *EvaluationDataset

	// The names of the metrics used. For automated model evaluation jobs valid values
	// are "Builtin.Accuracy" , "Builtin.Robustness" , and "Builtin.Toxicity" . In
	// human-based model evaluation jobs the array of strings must match the name
	// parameter specified in HumanEvaluationCustomMetric .
	//
	// This member is required.
	MetricNames []string

	// The task type you want the model to carry out.
	//
	// This member is required.
	TaskType EvaluationTaskType

	noSmithyDocumentSerde
}

// Used to define the models you want used in your model evaluation job. Automated
// model evaluation jobs support only a single model. In a human-based model
// evaluation job, your annotator can compare the responses for up to two different
// models.
//
// The following types satisfy this interface:
//
//	EvaluationInferenceConfigMemberModels
type EvaluationInferenceConfig interface {
	isEvaluationInferenceConfig()
}

// Used to specify the models.
type EvaluationInferenceConfigMemberModels struct {
	Value []EvaluationModelConfig

	noSmithyDocumentSerde
}

func (*EvaluationInferenceConfigMemberModels) isEvaluationInferenceConfig() {}

// Defines the models used in the model evaluation job.
//
// The following types satisfy this interface:
//
//	EvaluationModelConfigMemberBedrockModel
type EvaluationModelConfig interface {
	isEvaluationModelConfig()
}

// Defines the Amazon Bedrock model and inference parameters you want used.
type EvaluationModelConfigMemberBedrockModel struct {
	Value EvaluationBedrockModel

	noSmithyDocumentSerde
}

func (*EvaluationModelConfigMemberBedrockModel) isEvaluationModelConfig() {}

// The Amazon S3 location where the results of your model evaluation job are saved.
type EvaluationOutputDataConfig struct {

	// The Amazon S3 URI where the results of model evaluation job are saved.
	//
	// This member is required.
	S3Uri *string

	noSmithyDocumentSerde
}

// A summary of the model evaluation job.
type EvaluationSummary struct {

	// When the model evaluation job was created.
	//
	// This member is required.
	CreationTime *time.Time

	// What task type was used in the model evaluation job.
	//
	// This member is required.
	EvaluationTaskTypes []EvaluationTaskType

	// The Amazon Resource Name (ARN) of the model evaluation job.
	//
	// This member is required.
	JobArn *string

	// The name of the model evaluation job.
	//
	// This member is required.
	JobName *string

	// The type, either human or automatic, of model evaluation job.
	//
	// This member is required.
	JobType EvaluationJobType

	// The Amazon Resource Names (ARNs) of the model(s) used in the model evaluation
	// job.
	//
	// This member is required.
	ModelIdentifiers []string

	// The current status of the model evaluation job.
	//
	// This member is required.
	Status EvaluationJobStatus

	noSmithyDocumentSerde
}

// Information about a foundation model.
type FoundationModelDetails struct {

	// The model Amazon Resource Name (ARN).
	//
	// This member is required.
	ModelArn *string

	// The model identifier.
	//
	// This member is required.
	ModelId *string

	// The customization that the model supports.
	CustomizationsSupported []ModelCustomization

	// The inference types that the model supports.
	InferenceTypesSupported []InferenceType

	// The input modalities that the model supports.
	InputModalities []ModelModality

	// Contains details about whether a model version is available or deprecated
	ModelLifecycle *FoundationModelLifecycle

	// The model name.
	ModelName *string

	// The output modalities that the model supports.
	OutputModalities []ModelModality

	// The model's provider name.
	ProviderName *string

	// Indicates whether the model supports streaming.
	ResponseStreamingSupported *bool

	noSmithyDocumentSerde
}

// Details about whether a model version is available or deprecated.
type FoundationModelLifecycle struct {

	// Specifies whether a model version is available ( ACTIVE ) or deprecated ( LEGACY
	// .
	//
	// This member is required.
	Status FoundationModelLifecycleStatus

	noSmithyDocumentSerde
}

// Summary information for a foundation model.
type FoundationModelSummary struct {

	// The Amazon Resource Name (ARN) of the foundation model.
	//
	// This member is required.
	ModelArn *string

	// The model ID of the foundation model.
	//
	// This member is required.
	ModelId *string

	// Whether the model supports fine-tuning or continual pre-training.
	CustomizationsSupported []ModelCustomization

	// The inference types that the model supports.
	InferenceTypesSupported []InferenceType

	// The input modalities that the model supports.
	InputModalities []ModelModality

	// Contains details about whether a model version is available or deprecated.
	ModelLifecycle *FoundationModelLifecycle

	// The name of the model.
	ModelName *string

	// The output modalities that the model supports.
	OutputModalities []ModelModality

	// The model's provider name.
	ProviderName *string

	// Indicates whether the model supports streaming.
	ResponseStreamingSupported *bool

	noSmithyDocumentSerde
}

// Contains filter strengths for harmful content. Guardrails support the following
// content filters to detect and filter harmful user inputs and FM-generated
// outputs.
//   - Hate – Describes language or a statement that discriminates, criticizes,
//     insults, denounces, or dehumanizes a person or group on the basis of an identity
//     (such as race, ethnicity, gender, religion, sexual orientation, ability, and
//     national origin).
//   - Insults – Describes language or a statement that includes demeaning,
//     humiliating, mocking, insulting, or belittling language. This type of language
//     is also labeled as bullying.
//   - Sexual – Describes language or a statement that indicates sexual interest,
//     activity, or arousal using direct or indirect references to body parts, physical
//     traits, or sex.
//   - Violence – Describes language or a statement that includes glorification of
//     or threats to inflict physical pain, hurt, or injury toward a person, group or
//     thing.
//
// Content filtering depends on the confidence classification of user inputs and
// FM responses across each of the four harmful categories. All input and output
// statements are classified into one of four confidence levels (NONE, LOW, MEDIUM,
// HIGH) for each harmful category. For example, if a statement is classified as
// Hate with HIGH confidence, the likelihood of the statement representing hateful
// content is high. A single statement can be classified across multiple categories
// with varying confidence levels. For example, a single statement can be
// classified as Hate with HIGH confidence, Insults with LOW confidence, Sexual
// with NONE confidence, and Violence with MEDIUM confidence. For more information,
// see Guardrails content filters (https://docs.aws.amazon.com/bedrock/latest/userguide/guardrails-filters.html)
// . This data type is used in the following API operations:
//   - GetGuardrail response body (https://docs.aws.amazon.com/bedrock/latest/APIReference/API_GetGuardrail.html#API_GetGuardrail_ResponseSyntax)
type GuardrailContentFilter struct {

	// The strength of the content filter to apply to prompts. As you increase the
	// filter strength, the likelihood of filtering harmful content increases and the
	// probability of seeing harmful content in your application reduces.
	//
	// This member is required.
	InputStrength GuardrailFilterStrength

	// The strength of the content filter to apply to model responses. As you increase
	// the filter strength, the likelihood of filtering harmful content increases and
	// the probability of seeing harmful content in your application reduces.
	//
	// This member is required.
	OutputStrength GuardrailFilterStrength

	// The harmful category that the content filter is applied to.
	//
	// This member is required.
	Type GuardrailContentFilterType

	noSmithyDocumentSerde
}

// Contains filter strengths for harmful content. Guardrails support the following
// content filters to detect and filter harmful user inputs and FM-generated
// outputs.
//   - Hate – Describes language or a statement that discriminates, criticizes,
//     insults, denounces, or dehumanizes a person or group on the basis of an identity
//     (such as race, ethnicity, gender, religion, sexual orientation, ability, and
//     national origin).
//   - Insults – Describes language or a statement that includes demeaning,
//     humiliating, mocking, insulting, or belittling language. This type of language
//     is also labeled as bullying.
//   - Sexual – Describes language or a statement that indicates sexual interest,
//     activity, or arousal using direct or indirect references to body parts, physical
//     traits, or sex.
//   - Violence – Describes language or a statement that includes glorification of
//     or threats to inflict physical pain, hurt, or injury toward a person, group or
//     thing.
//
// Content filtering depends on the confidence classification of user inputs and
// FM responses across each of the four harmful categories. All input and output
// statements are classified into one of four confidence levels (NONE, LOW, MEDIUM,
// HIGH) for each harmful category. For example, if a statement is classified as
// Hate with HIGH confidence, the likelihood of the statement representing hateful
// content is high. A single statement can be classified across multiple categories
// with varying confidence levels. For example, a single statement can be
// classified as Hate with HIGH confidence, Insults with LOW confidence, Sexual
// with NONE confidence, and Violence with MEDIUM confidence. For more information,
// see Guardrails content filters (https://docs.aws.amazon.com/bedrock/latest/userguide/guardrails-filters.html)
// . This data type is used in the following API operations:
//   - CreateGuardrail request body (https://docs.aws.amazon.com/bedrock/latest/APIReference/API_CreateGuardrail.html#API_CreateGuardrail_RequestSyntax)
//   - UpdateGuardrail request body (https://docs.aws.amazon.com/bedrock/latest/APIReference/API_UpdateGuardrail.html#API_UpdateGuardrail_RequestSyntax)
type GuardrailContentFilterConfig struct {

	// The strength of the content filter to apply to prompts. As you increase the
	// filter strength, the likelihood of filtering harmful content increases and the
	// probability of seeing harmful content in your application reduces.
	//
	// This member is required.
	InputStrength GuardrailFilterStrength

	// The strength of the content filter to apply to model responses. As you increase
	// the filter strength, the likelihood of filtering harmful content increases and
	// the probability of seeing harmful content in your application reduces.
	//
	// This member is required.
	OutputStrength GuardrailFilterStrength

	// The harmful category that the content filter is applied to.
	//
	// This member is required.
	Type GuardrailContentFilterType

	noSmithyDocumentSerde
}

// Contains details about how to handle harmful content. This data type is used in
// the following API operations:
//   - GetGuardrail response body (https://docs.aws.amazon.com/bedrock/latest/APIReference/API_GetGuardrail.html#API_GetGuardrail_ResponseSyntax)
type GuardrailContentPolicy struct {

	// Contains the type of the content filter and how strongly it should apply to
	// prompts and model responses.
	Filters []GuardrailContentFilter

	noSmithyDocumentSerde
}

// Contains details about how to handle harmful content. This data type is used in
// the following API operations:
//   - CreateGuardrail request body (https://docs.aws.amazon.com/bedrock/latest/APIReference/API_CreateGuardrail.html#API_CreateGuardrail_RequestSyntax)
//   - UpdateGuardrail request body (https://docs.aws.amazon.com/bedrock/latest/APIReference/API_UpdateGuardrail.html#API_UpdateGuardrail_RequestSyntax)
type GuardrailContentPolicyConfig struct {

	// Contains the type of the content filter and how strongly it should apply to
	// prompts and model responses.
	//
	// This member is required.
	FiltersConfig []GuardrailContentFilterConfig

	noSmithyDocumentSerde
}

// The managed word list that was configured for the guardrail. (This is a list of
// words that are pre-defined and managed by Guardrails only.)
type GuardrailManagedWords struct {

	// ManagedWords$type The managed word type that was configured for the guardrail.
	// (For now, we only offer profanity word list)
	//
	// This member is required.
	Type GuardrailManagedWordsType

	noSmithyDocumentSerde
}

// The managed word list to configure for the guardrail.
type GuardrailManagedWordsConfig struct {

	// The managed word type to configure for the guardrail.
	//
	// This member is required.
	Type GuardrailManagedWordsType

	noSmithyDocumentSerde
}

// The PII entity configured for the guardrail.
type GuardrailPiiEntity struct {

	// The configured guardrail action when PII entity is detected.
	//
	// This member is required.
	Action GuardrailSensitiveInformationAction

	// The type of PII entity. For example, Social Security Number.
	//
	// This member is required.
	Type GuardrailPiiEntityType

	noSmithyDocumentSerde
}

// The PII entity to configure for the guardrail.
type GuardrailPiiEntityConfig struct {

	// Configure guardrail action when the PII entity is detected.
	//
	// This member is required.
	Action GuardrailSensitiveInformationAction

	// Configure guardrail type when the PII entity is detected.
	//
	// This member is required.
	Type GuardrailPiiEntityType

	noSmithyDocumentSerde
}

// The regular expression configured for the guardrail.
type GuardrailRegex struct {

	// The action taken when a match to the regular expression is detected.
	//
	// This member is required.
	Action GuardrailSensitiveInformationAction

	// The name of the regular expression for the guardrail.
	//
	// This member is required.
	Name *string

	// The pattern of the regular expression configured for the guardrail.
	//
	// This member is required.
	Pattern *string

	// The description of the regular expression for the guardrail.
	Description *string

	noSmithyDocumentSerde
}

// The regular expression to configure for the guardrail.
type GuardrailRegexConfig struct {

	// The guardrail action to configure when matching regular expression is detected.
	//
	// This member is required.
	Action GuardrailSensitiveInformationAction

	// The name of the regular expression to configure for the guardrail.
	//
	// This member is required.
	Name *string

	// The regular expression pattern to configure for the guardrail.
	//
	// This member is required.
	Pattern *string

	// The description of the regular expression to configure for the guardrail.
	Description *string

	noSmithyDocumentSerde
}

// Contains details about PII entities and regular expressions configured for the
// guardrail.
type GuardrailSensitiveInformationPolicy struct {

	// The list of PII entities configured for the guardrail.
	PiiEntities []GuardrailPiiEntity

	// The list of regular expressions configured for the guardrail.
	Regexes []GuardrailRegex

	noSmithyDocumentSerde
}

// Contains details about PII entities and regular expressions to configure for
// the guardrail.
type GuardrailSensitiveInformationPolicyConfig struct {

	// A list of PII entities to configure to the guardrail.
	PiiEntitiesConfig []GuardrailPiiEntityConfig

	// A list of regular expressions to configure to the guardrail.
	RegexesConfig []GuardrailRegexConfig

	noSmithyDocumentSerde
}

// Contains details about a guardrail. This data type is used in the following API
// operations:
//   - ListGuardrails response body (https://docs.aws.amazon.com/bedrock/latest/APIReference/API_ListGuardrails.html#API_ListGuardrails_ResponseSyntax)
type GuardrailSummary struct {

	// The ARN of the guardrail.
	//
	// This member is required.
	Arn *string

	// The date and time at which the guardrail was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The unique identifier of the guardrail.
	//
	// This member is required.
	Id *string

	// The name of the guardrail.
	//
	// This member is required.
	Name *string

	// The status of the guardrail.
	//
	// This member is required.
	Status GuardrailStatus

	// The date and time at which the guardrail was last updated.
	//
	// This member is required.
	UpdatedAt *time.Time

	// The version of the guardrail.
	//
	// This member is required.
	Version *string

	// A description of the guardrail.
	Description *string

	noSmithyDocumentSerde
}

// Details about topics for the guardrail to identify and deny. This data type is
// used in the following API operations:
//   - GetGuardrail response body (https://docs.aws.amazon.com/bedrock/latest/APIReference/API_GetGuardrail.html#API_GetGuardrail_ResponseSyntax)
type GuardrailTopic struct {

	// A definition of the topic to deny.
	//
	// This member is required.
	Definition *string

	// The name of the topic to deny.
	//
	// This member is required.
	Name *string

	// A list of prompts, each of which is an example of a prompt that can be
	// categorized as belonging to the topic.
	Examples []string

	// Specifies to deny the topic.
	Type GuardrailTopicType

	noSmithyDocumentSerde
}

// Details about topics for the guardrail to identify and deny. This data type is
// used in the following API operations:
//   - CreateGuardrail request body (https://docs.aws.amazon.com/bedrock/latest/APIReference/API_CreateGuardrail.html#API_CreateGuardrail_RequestSyntax)
//   - UpdateGuardrail request body (https://docs.aws.amazon.com/bedrock/latest/APIReference/API_UpdateGuardrail.html#API_UpdateGuardrail_RequestSyntax)
type GuardrailTopicConfig struct {

	// A definition of the topic to deny.
	//
	// This member is required.
	Definition *string

	// The name of the topic to deny.
	//
	// This member is required.
	Name *string

	// Specifies to deny the topic.
	//
	// This member is required.
	Type GuardrailTopicType

	// A list of prompts, each of which is an example of a prompt that can be
	// categorized as belonging to the topic.
	Examples []string

	noSmithyDocumentSerde
}

// Contains details about topics that the guardrail should identify and deny. This
// data type is used in the following API operations:
//   - GetGuardrail response body (https://docs.aws.amazon.com/bedrock/latest/APIReference/API_GetGuardrail.html#API_GetGuardrail_ResponseSyntax)
type GuardrailTopicPolicy struct {

	// A list of policies related to topics that the guardrail should deny.
	//
	// This member is required.
	Topics []GuardrailTopic

	noSmithyDocumentSerde
}

// Contains details about topics that the guardrail should identify and deny. This
// data type is used in the following API operations:
//   - CreateGuardrail request body (https://docs.aws.amazon.com/bedrock/latest/APIReference/API_CreateGuardrail.html#API_CreateGuardrail_RequestSyntax)
//   - UpdateGuardrail request body (https://docs.aws.amazon.com/bedrock/latest/APIReference/API_UpdateGuardrail.html#API_UpdateGuardrail_RequestSyntax)
type GuardrailTopicPolicyConfig struct {

	// A list of policies related to topics that the guardrail should deny.
	//
	// This member is required.
	TopicsConfig []GuardrailTopicConfig

	noSmithyDocumentSerde
}

// A word configured for the guardrail.
type GuardrailWord struct {

	// Text of the word configured for the guardrail to block.
	//
	// This member is required.
	Text *string

	noSmithyDocumentSerde
}

// A word to configure for the guardrail.
type GuardrailWordConfig struct {

	// Text of the word configured for the guardrail to block.
	//
	// This member is required.
	Text *string

	noSmithyDocumentSerde
}

// Contains details about the word policy configured for the guardrail.
type GuardrailWordPolicy struct {

	// A list of managed words configured for the guardrail.
	ManagedWordLists []GuardrailManagedWords

	// A list of words configured for the guardrail.
	Words []GuardrailWord

	noSmithyDocumentSerde
}

// Contains details about the word policy to configured for the guardrail.
type GuardrailWordPolicyConfig struct {

	// A list of managed words to configure for the guardrail.
	ManagedWordListsConfig []GuardrailManagedWordsConfig

	// A list of words to configure for the guardrail.
	WordsConfig []GuardrailWordConfig

	noSmithyDocumentSerde
}

// Specifies the custom metrics, how tasks will be rated, the flow definition ARN,
// and your custom prompt datasets. Model evaluation jobs use human workers only
// support the use of custom prompt datasets. To learn more about custom prompt
// datasets and the required format, see Custom prompt datasets (https://docs.aws.amazon.com/bedrock/latest/userguide/model-evaluation-prompt-datasets-custom.html)
// . When you create custom metrics in HumanEvaluationCustomMetric you must
// specify the metric's name . The list of names specified in the
// HumanEvaluationCustomMetric array, must match the metricNames array of strings
// specified in EvaluationDatasetMetricConfig . For example, if in the
// HumanEvaluationCustomMetric array your specified the names "accuracy",
// "toxicity", "readability" as custom metrics then the metricNames array would
// need to look like the following ["accuracy", "toxicity", "readability"] in
// EvaluationDatasetMetricConfig .
type HumanEvaluationConfig struct {

	// Use to specify the metrics, task, and prompt dataset to be used in your model
	// evaluation job.
	//
	// This member is required.
	DatasetMetricConfigs []EvaluationDatasetMetricConfig

	// A HumanEvaluationCustomMetric object. It contains the names the metrics, how
	// the metrics are to be evaluated, an optional description.
	CustomMetrics []HumanEvaluationCustomMetric

	// The parameters of the human workflow.
	HumanWorkflowConfig *HumanWorkflowConfig

	noSmithyDocumentSerde
}

// In a model evaluation job that uses human workers you must define the name of
// the metric, and how you want that metric rated ratingMethod , and an optional
// description of the metric.
type HumanEvaluationCustomMetric struct {

	// The name of the metric. Your human evaluators will see this name in the
	// evaluation UI.
	//
	// This member is required.
	Name *string

	// Choose how you want your human workers to evaluation your model. Valid values
	// for rating methods are ThumbsUpDown , IndividualLikertScale ,
	// ComparisonLikertScale , ComparisonChoice , and ComparisonRank
	//
	// This member is required.
	RatingMethod *string

	// An optional description of the metric. Use this parameter to provide more
	// details about the metric.
	Description *string

	noSmithyDocumentSerde
}

// Contains SageMakerFlowDefinition object. The object is used to specify the
// prompt dataset, task type, rating method and metric names.
type HumanWorkflowConfig struct {

	// The Amazon Resource Number (ARN) for the flow definition
	//
	// This member is required.
	FlowDefinitionArn *string

	// Instructions for the flow definition
	Instructions *string

	noSmithyDocumentSerde
}

// Configuration fields for invocation logging.
type LoggingConfig struct {

	// CloudWatch logging configuration.
	CloudWatchConfig *CloudWatchConfig

	// Set to include embeddings data in the log delivery.
	EmbeddingDataDeliveryEnabled *bool

	// Set to include image data in the log delivery.
	ImageDataDeliveryEnabled *bool

	// S3 configuration for storing log data.
	S3Config *S3Config

	// Set to include text data in the log delivery.
	TextDataDeliveryEnabled *bool

	noSmithyDocumentSerde
}

// Information about one customization job
type ModelCustomizationJobSummary struct {

	// Amazon Resource Name (ARN) of the base model.
	//
	// This member is required.
	BaseModelArn *string

	// Creation time of the custom model.
	//
	// This member is required.
	CreationTime *time.Time

	// Amazon Resource Name (ARN) of the customization job.
	//
	// This member is required.
	JobArn *string

	// Name of the customization job.
	//
	// This member is required.
	JobName *string

	// Status of the customization job.
	//
	// This member is required.
	Status ModelCustomizationJobStatus

	// Amazon Resource Name (ARN) of the custom model.
	CustomModelArn *string

	// Name of the custom model.
	CustomModelName *string

	// Specifies whether to carry out continued pre-training of a model or whether to
	// fine-tune it. For more information, see Custom models (https://docs.aws.amazon.com/bedrock/latest/userguide/custom-models.html)
	// .
	CustomizationType CustomizationType

	// Time that the customization job ended.
	EndTime *time.Time

	// Time that the customization job was last modified.
	LastModifiedTime *time.Time

	noSmithyDocumentSerde
}

// S3 Location of the output data.
type OutputDataConfig struct {

	// The S3 URI where the output data is stored.
	//
	// This member is required.
	S3Uri *string

	noSmithyDocumentSerde
}

// A summary of information about a Provisioned Throughput. This data type is used
// in the following API operations:
//   - ListProvisionedThroughputs response (https://docs.aws.amazon.com/bedrock/latest/APIReference/API_ListProvisionedModelThroughputs.html#API_ListProvisionedModelThroughputs_ResponseSyntax)
type ProvisionedModelSummary struct {

	// The time that the Provisioned Throughput was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) of the model requested to be associated to this
	// Provisioned Throughput. This value differs from the modelArn if updating hasn't
	// completed.
	//
	// This member is required.
	DesiredModelArn *string

	// The number of model units that was requested to be allocated to the Provisioned
	// Throughput.
	//
	// This member is required.
	DesiredModelUnits *int32

	// The Amazon Resource Name (ARN) of the base model for which the Provisioned
	// Throughput was created, or of the base model that the custom model for which the
	// Provisioned Throughput was created was customized.
	//
	// This member is required.
	FoundationModelArn *string

	// The time that the Provisioned Throughput was last modified.
	//
	// This member is required.
	LastModifiedTime *time.Time

	// The Amazon Resource Name (ARN) of the model associated with the Provisioned
	// Throughput.
	//
	// This member is required.
	ModelArn *string

	// The number of model units allocated to the Provisioned Throughput.
	//
	// This member is required.
	ModelUnits *int32

	// The Amazon Resource Name (ARN) of the Provisioned Throughput.
	//
	// This member is required.
	ProvisionedModelArn *string

	// The name of the Provisioned Throughput.
	//
	// This member is required.
	ProvisionedModelName *string

	// The status of the Provisioned Throughput.
	//
	// This member is required.
	Status ProvisionedModelStatus

	// The duration for which the Provisioned Throughput was committed.
	CommitmentDuration CommitmentDuration

	// The timestamp for when the commitment term of the Provisioned Throughput
	// expires.
	CommitmentExpirationTime *time.Time

	noSmithyDocumentSerde
}

// S3 configuration for storing log data.
type S3Config struct {

	// S3 bucket name.
	//
	// This member is required.
	BucketName *string

	// S3 prefix.
	KeyPrefix *string

	noSmithyDocumentSerde
}

// Definition of the key/value pair for a tag.
type Tag struct {

	// Key for the tag.
	//
	// This member is required.
	Key *string

	// Value for the tag.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// S3 Location of the training data.
type TrainingDataConfig struct {

	// The S3 URI where the training data is stored.
	//
	// This member is required.
	S3Uri *string

	noSmithyDocumentSerde
}

// Metrics associated with the custom job.
type TrainingMetrics struct {

	// Loss metric associated with the custom job.
	TrainingLoss *float32

	noSmithyDocumentSerde
}

// Array of up to 10 validators.
type ValidationDataConfig struct {

	// Information about the validators.
	//
	// This member is required.
	Validators []Validator

	noSmithyDocumentSerde
}

// Information about a validator.
type Validator struct {

	// The S3 URI where the validation data is stored.
	//
	// This member is required.
	S3Uri *string

	noSmithyDocumentSerde
}

// The metric for the validator.
type ValidatorMetric struct {

	// The validation loss associated with this validator.
	ValidationLoss *float32

	noSmithyDocumentSerde
}

// VPC configuration.
type VpcConfig struct {

	// VPC configuration security group Ids.
	//
	// This member is required.
	SecurityGroupIds []string

	// VPC configuration subnets.
	//
	// This member is required.
	SubnetIds []string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte

	noSmithyDocumentSerde
}

func (*UnknownUnionMember) isEvaluationConfig()          {}
func (*UnknownUnionMember) isEvaluationDatasetLocation() {}
func (*UnknownUnionMember) isEvaluationInferenceConfig() {}
func (*UnknownUnionMember) isEvaluationModelConfig()     {}
