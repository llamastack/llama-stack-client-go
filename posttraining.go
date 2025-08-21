// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"net/http"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
	"github.com/llamastack/llama-stack-client-go/shared/constant"
)

// PostTrainingService contains methods and other services that help with
// interacting with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPostTrainingService] method instead.
type PostTrainingService struct {
	Options []option.RequestOption
	Job     PostTrainingJobService
}

// NewPostTrainingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPostTrainingService(opts ...option.RequestOption) (r PostTrainingService) {
	r = PostTrainingService{}
	r.Options = opts
	r.Job = NewPostTrainingJobService(opts...)
	return
}

// Run preference optimization of a model.
func (r *PostTrainingService) PreferenceOptimize(ctx context.Context, body PostTrainingPreferenceOptimizeParams, opts ...option.RequestOption) (res *PostTrainingJob, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/post-training/preference-optimize"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Run supervised fine-tuning of a model.
func (r *PostTrainingService) SupervisedFineTune(ctx context.Context, body PostTrainingSupervisedFineTuneParams, opts ...option.RequestOption) (res *PostTrainingJob, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/post-training/supervised-fine-tune"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func AlgorithmConfigParamOfQat(groupSize int64, quantizerName string) AlgorithmConfigUnionParam {
	var qat AlgorithmConfigQatParam
	qat.GroupSize = groupSize
	qat.QuantizerName = quantizerName
	return AlgorithmConfigUnionParam{OfQat: &qat}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AlgorithmConfigUnionParam struct {
	OfLoRa *AlgorithmConfigLoRaParam `json:",omitzero,inline"`
	OfQat  *AlgorithmConfigQatParam  `json:",omitzero,inline"`
	paramUnion
}

func (u AlgorithmConfigUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfLoRa, u.OfQat)
}
func (u *AlgorithmConfigUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AlgorithmConfigUnionParam) asAny() any {
	if !param.IsOmitted(u.OfLoRa) {
		return u.OfLoRa
	} else if !param.IsOmitted(u.OfQat) {
		return u.OfQat
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AlgorithmConfigUnionParam) GetAlpha() *int64 {
	if vt := u.OfLoRa; vt != nil {
		return &vt.Alpha
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AlgorithmConfigUnionParam) GetApplyLoraToMlp() *bool {
	if vt := u.OfLoRa; vt != nil {
		return &vt.ApplyLoraToMlp
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AlgorithmConfigUnionParam) GetApplyLoraToOutput() *bool {
	if vt := u.OfLoRa; vt != nil {
		return &vt.ApplyLoraToOutput
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AlgorithmConfigUnionParam) GetLoraAttnModules() []string {
	if vt := u.OfLoRa; vt != nil {
		return vt.LoraAttnModules
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AlgorithmConfigUnionParam) GetRank() *int64 {
	if vt := u.OfLoRa; vt != nil {
		return &vt.Rank
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AlgorithmConfigUnionParam) GetQuantizeBase() *bool {
	if vt := u.OfLoRa; vt != nil && vt.QuantizeBase.Valid() {
		return &vt.QuantizeBase.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AlgorithmConfigUnionParam) GetUseDora() *bool {
	if vt := u.OfLoRa; vt != nil && vt.UseDora.Valid() {
		return &vt.UseDora.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AlgorithmConfigUnionParam) GetGroupSize() *int64 {
	if vt := u.OfQat; vt != nil {
		return &vt.GroupSize
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AlgorithmConfigUnionParam) GetQuantizerName() *string {
	if vt := u.OfQat; vt != nil {
		return &vt.QuantizerName
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AlgorithmConfigUnionParam) GetType() *string {
	if vt := u.OfLoRa; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfQat; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[AlgorithmConfigUnionParam](
		"type",
		apijson.Discriminator[AlgorithmConfigLoRaParam]("LoRA"),
		apijson.Discriminator[AlgorithmConfigQatParam]("QAT"),
	)
}

// Configuration for Low-Rank Adaptation (LoRA) fine-tuning.
//
// The properties Alpha, ApplyLoraToMlp, ApplyLoraToOutput, LoraAttnModules, Rank,
// Type are required.
type AlgorithmConfigLoRaParam struct {
	// LoRA scaling parameter that controls adaptation strength
	Alpha int64 `json:"alpha,required"`
	// Whether to apply LoRA to MLP layers
	ApplyLoraToMlp bool `json:"apply_lora_to_mlp,required"`
	// Whether to apply LoRA to output projection layers
	ApplyLoraToOutput bool `json:"apply_lora_to_output,required"`
	// List of attention module names to apply LoRA to
	LoraAttnModules []string `json:"lora_attn_modules,omitzero,required"`
	// Rank of the LoRA adaptation (lower rank = fewer parameters)
	Rank int64 `json:"rank,required"`
	// (Optional) Whether to quantize the base model weights
	QuantizeBase param.Opt[bool] `json:"quantize_base,omitzero"`
	// (Optional) Whether to use DoRA (Weight-Decomposed Low-Rank Adaptation)
	UseDora param.Opt[bool] `json:"use_dora,omitzero"`
	// Algorithm type identifier, always "LoRA"
	//
	// This field can be elided, and will marshal its zero value as "LoRA".
	Type constant.LoRa `json:"type,required"`
	paramObj
}

func (r AlgorithmConfigLoRaParam) MarshalJSON() (data []byte, err error) {
	type shadow AlgorithmConfigLoRaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlgorithmConfigLoRaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for Quantization-Aware Training (QAT) fine-tuning.
//
// The properties GroupSize, QuantizerName, Type are required.
type AlgorithmConfigQatParam struct {
	// Size of groups for grouped quantization
	GroupSize int64 `json:"group_size,required"`
	// Name of the quantization algorithm to use
	QuantizerName string `json:"quantizer_name,required"`
	// Algorithm type identifier, always "QAT"
	//
	// This field can be elided, and will marshal its zero value as "QAT".
	Type constant.Qat `json:"type,required"`
	paramObj
}

func (r AlgorithmConfigQatParam) MarshalJSON() (data []byte, err error) {
	type shadow AlgorithmConfigQatParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlgorithmConfigQatParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostTrainingJob struct {
	JobUuid string `json:"job_uuid,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		JobUuid     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PostTrainingJob) RawJSON() string { return r.JSON.raw }
func (r *PostTrainingJob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PostTrainingPreferenceOptimizeParams struct {
	// The algorithm configuration.
	AlgorithmConfig PostTrainingPreferenceOptimizeParamsAlgorithmConfig `json:"algorithm_config,omitzero,required"`
	// The model to fine-tune.
	FinetunedModel string `json:"finetuned_model,required"`
	// The hyperparam search configuration.
	HyperparamSearchConfig map[string]PostTrainingPreferenceOptimizeParamsHyperparamSearchConfigUnion `json:"hyperparam_search_config,omitzero,required"`
	// The UUID of the job to create.
	JobUuid string `json:"job_uuid,required"`
	// The logger configuration.
	LoggerConfig map[string]PostTrainingPreferenceOptimizeParamsLoggerConfigUnion `json:"logger_config,omitzero,required"`
	// The training configuration.
	TrainingConfig PostTrainingPreferenceOptimizeParamsTrainingConfig `json:"training_config,omitzero,required"`
	paramObj
}

func (r PostTrainingPreferenceOptimizeParams) MarshalJSON() (data []byte, err error) {
	type shadow PostTrainingPreferenceOptimizeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PostTrainingPreferenceOptimizeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The algorithm configuration.
//
// The properties Beta, LossType are required.
type PostTrainingPreferenceOptimizeParamsAlgorithmConfig struct {
	// Temperature parameter for the DPO loss
	Beta float64 `json:"beta,required"`
	// The type of loss function to use for DPO
	//
	// Any of "sigmoid", "hinge", "ipo", "kto_pair".
	LossType string `json:"loss_type,omitzero,required"`
	paramObj
}

func (r PostTrainingPreferenceOptimizeParamsAlgorithmConfig) MarshalJSON() (data []byte, err error) {
	type shadow PostTrainingPreferenceOptimizeParamsAlgorithmConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PostTrainingPreferenceOptimizeParamsAlgorithmConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PostTrainingPreferenceOptimizeParamsAlgorithmConfig](
		"loss_type", "sigmoid", "hinge", "ipo", "kto_pair",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PostTrainingPreferenceOptimizeParamsHyperparamSearchConfigUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u PostTrainingPreferenceOptimizeParamsHyperparamSearchConfigUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *PostTrainingPreferenceOptimizeParamsHyperparamSearchConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PostTrainingPreferenceOptimizeParamsHyperparamSearchConfigUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PostTrainingPreferenceOptimizeParamsLoggerConfigUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u PostTrainingPreferenceOptimizeParamsLoggerConfigUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *PostTrainingPreferenceOptimizeParamsLoggerConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PostTrainingPreferenceOptimizeParamsLoggerConfigUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	}
	return nil
}

// The training configuration.
//
// The properties GradientAccumulationSteps, MaxStepsPerEpoch, NEpochs are
// required.
type PostTrainingPreferenceOptimizeParamsTrainingConfig struct {
	// Number of steps to accumulate gradients before updating
	GradientAccumulationSteps int64 `json:"gradient_accumulation_steps,required"`
	// Maximum number of steps to run per epoch
	MaxStepsPerEpoch int64 `json:"max_steps_per_epoch,required"`
	// Number of training epochs to run
	NEpochs int64 `json:"n_epochs,required"`
	// (Optional) Data type for model parameters (bf16, fp16, fp32)
	Dtype param.Opt[string] `json:"dtype,omitzero"`
	// (Optional) Maximum number of validation steps per epoch
	MaxValidationSteps param.Opt[int64] `json:"max_validation_steps,omitzero"`
	// (Optional) Configuration for data loading and formatting
	DataConfig PostTrainingPreferenceOptimizeParamsTrainingConfigDataConfig `json:"data_config,omitzero"`
	// (Optional) Configuration for memory and compute optimizations
	EfficiencyConfig PostTrainingPreferenceOptimizeParamsTrainingConfigEfficiencyConfig `json:"efficiency_config,omitzero"`
	// (Optional) Configuration for the optimization algorithm
	OptimizerConfig PostTrainingPreferenceOptimizeParamsTrainingConfigOptimizerConfig `json:"optimizer_config,omitzero"`
	paramObj
}

func (r PostTrainingPreferenceOptimizeParamsTrainingConfig) MarshalJSON() (data []byte, err error) {
	type shadow PostTrainingPreferenceOptimizeParamsTrainingConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PostTrainingPreferenceOptimizeParamsTrainingConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) Configuration for data loading and formatting
//
// The properties BatchSize, DataFormat, DatasetID, Shuffle are required.
type PostTrainingPreferenceOptimizeParamsTrainingConfigDataConfig struct {
	// Number of samples per training batch
	BatchSize int64 `json:"batch_size,required"`
	// Format of the dataset (instruct or dialog)
	//
	// Any of "instruct", "dialog".
	DataFormat string `json:"data_format,omitzero,required"`
	// Unique identifier for the training dataset
	DatasetID string `json:"dataset_id,required"`
	// Whether to shuffle the dataset during training
	Shuffle bool `json:"shuffle,required"`
	// (Optional) Whether to pack multiple samples into a single sequence for
	// efficiency
	Packed param.Opt[bool] `json:"packed,omitzero"`
	// (Optional) Whether to compute loss on input tokens as well as output tokens
	TrainOnInput param.Opt[bool] `json:"train_on_input,omitzero"`
	// (Optional) Unique identifier for the validation dataset
	ValidationDatasetID param.Opt[string] `json:"validation_dataset_id,omitzero"`
	paramObj
}

func (r PostTrainingPreferenceOptimizeParamsTrainingConfigDataConfig) MarshalJSON() (data []byte, err error) {
	type shadow PostTrainingPreferenceOptimizeParamsTrainingConfigDataConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PostTrainingPreferenceOptimizeParamsTrainingConfigDataConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PostTrainingPreferenceOptimizeParamsTrainingConfigDataConfig](
		"data_format", "instruct", "dialog",
	)
}

// (Optional) Configuration for memory and compute optimizations
type PostTrainingPreferenceOptimizeParamsTrainingConfigEfficiencyConfig struct {
	// (Optional) Whether to use activation checkpointing to reduce memory usage
	EnableActivationCheckpointing param.Opt[bool] `json:"enable_activation_checkpointing,omitzero"`
	// (Optional) Whether to offload activations to CPU to save GPU memory
	EnableActivationOffloading param.Opt[bool] `json:"enable_activation_offloading,omitzero"`
	// (Optional) Whether to offload FSDP parameters to CPU
	FsdpCPUOffload param.Opt[bool] `json:"fsdp_cpu_offload,omitzero"`
	// (Optional) Whether to use memory-efficient FSDP wrapping
	MemoryEfficientFsdpWrap param.Opt[bool] `json:"memory_efficient_fsdp_wrap,omitzero"`
	paramObj
}

func (r PostTrainingPreferenceOptimizeParamsTrainingConfigEfficiencyConfig) MarshalJSON() (data []byte, err error) {
	type shadow PostTrainingPreferenceOptimizeParamsTrainingConfigEfficiencyConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PostTrainingPreferenceOptimizeParamsTrainingConfigEfficiencyConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) Configuration for the optimization algorithm
//
// The properties Lr, NumWarmupSteps, OptimizerType, WeightDecay are required.
type PostTrainingPreferenceOptimizeParamsTrainingConfigOptimizerConfig struct {
	// Learning rate for the optimizer
	Lr float64 `json:"lr,required"`
	// Number of steps for learning rate warmup
	NumWarmupSteps int64 `json:"num_warmup_steps,required"`
	// Type of optimizer to use (adam, adamw, or sgd)
	//
	// Any of "adam", "adamw", "sgd".
	OptimizerType string `json:"optimizer_type,omitzero,required"`
	// Weight decay coefficient for regularization
	WeightDecay float64 `json:"weight_decay,required"`
	paramObj
}

func (r PostTrainingPreferenceOptimizeParamsTrainingConfigOptimizerConfig) MarshalJSON() (data []byte, err error) {
	type shadow PostTrainingPreferenceOptimizeParamsTrainingConfigOptimizerConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PostTrainingPreferenceOptimizeParamsTrainingConfigOptimizerConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PostTrainingPreferenceOptimizeParamsTrainingConfigOptimizerConfig](
		"optimizer_type", "adam", "adamw", "sgd",
	)
}

type PostTrainingSupervisedFineTuneParams struct {
	// The hyperparam search configuration.
	HyperparamSearchConfig map[string]PostTrainingSupervisedFineTuneParamsHyperparamSearchConfigUnion `json:"hyperparam_search_config,omitzero,required"`
	// The UUID of the job to create.
	JobUuid string `json:"job_uuid,required"`
	// The logger configuration.
	LoggerConfig map[string]PostTrainingSupervisedFineTuneParamsLoggerConfigUnion `json:"logger_config,omitzero,required"`
	// The training configuration.
	TrainingConfig PostTrainingSupervisedFineTuneParamsTrainingConfig `json:"training_config,omitzero,required"`
	// The directory to save checkpoint(s) to.
	CheckpointDir param.Opt[string] `json:"checkpoint_dir,omitzero"`
	// The model to fine-tune.
	Model param.Opt[string] `json:"model,omitzero"`
	// The algorithm configuration.
	AlgorithmConfig AlgorithmConfigUnionParam `json:"algorithm_config,omitzero"`
	paramObj
}

func (r PostTrainingSupervisedFineTuneParams) MarshalJSON() (data []byte, err error) {
	type shadow PostTrainingSupervisedFineTuneParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PostTrainingSupervisedFineTuneParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PostTrainingSupervisedFineTuneParamsHyperparamSearchConfigUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u PostTrainingSupervisedFineTuneParamsHyperparamSearchConfigUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *PostTrainingSupervisedFineTuneParamsHyperparamSearchConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PostTrainingSupervisedFineTuneParamsHyperparamSearchConfigUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PostTrainingSupervisedFineTuneParamsLoggerConfigUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u PostTrainingSupervisedFineTuneParamsLoggerConfigUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *PostTrainingSupervisedFineTuneParamsLoggerConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PostTrainingSupervisedFineTuneParamsLoggerConfigUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	}
	return nil
}

// The training configuration.
//
// The properties GradientAccumulationSteps, MaxStepsPerEpoch, NEpochs are
// required.
type PostTrainingSupervisedFineTuneParamsTrainingConfig struct {
	// Number of steps to accumulate gradients before updating
	GradientAccumulationSteps int64 `json:"gradient_accumulation_steps,required"`
	// Maximum number of steps to run per epoch
	MaxStepsPerEpoch int64 `json:"max_steps_per_epoch,required"`
	// Number of training epochs to run
	NEpochs int64 `json:"n_epochs,required"`
	// (Optional) Data type for model parameters (bf16, fp16, fp32)
	Dtype param.Opt[string] `json:"dtype,omitzero"`
	// (Optional) Maximum number of validation steps per epoch
	MaxValidationSteps param.Opt[int64] `json:"max_validation_steps,omitzero"`
	// (Optional) Configuration for data loading and formatting
	DataConfig PostTrainingSupervisedFineTuneParamsTrainingConfigDataConfig `json:"data_config,omitzero"`
	// (Optional) Configuration for memory and compute optimizations
	EfficiencyConfig PostTrainingSupervisedFineTuneParamsTrainingConfigEfficiencyConfig `json:"efficiency_config,omitzero"`
	// (Optional) Configuration for the optimization algorithm
	OptimizerConfig PostTrainingSupervisedFineTuneParamsTrainingConfigOptimizerConfig `json:"optimizer_config,omitzero"`
	paramObj
}

func (r PostTrainingSupervisedFineTuneParamsTrainingConfig) MarshalJSON() (data []byte, err error) {
	type shadow PostTrainingSupervisedFineTuneParamsTrainingConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PostTrainingSupervisedFineTuneParamsTrainingConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) Configuration for data loading and formatting
//
// The properties BatchSize, DataFormat, DatasetID, Shuffle are required.
type PostTrainingSupervisedFineTuneParamsTrainingConfigDataConfig struct {
	// Number of samples per training batch
	BatchSize int64 `json:"batch_size,required"`
	// Format of the dataset (instruct or dialog)
	//
	// Any of "instruct", "dialog".
	DataFormat string `json:"data_format,omitzero,required"`
	// Unique identifier for the training dataset
	DatasetID string `json:"dataset_id,required"`
	// Whether to shuffle the dataset during training
	Shuffle bool `json:"shuffle,required"`
	// (Optional) Whether to pack multiple samples into a single sequence for
	// efficiency
	Packed param.Opt[bool] `json:"packed,omitzero"`
	// (Optional) Whether to compute loss on input tokens as well as output tokens
	TrainOnInput param.Opt[bool] `json:"train_on_input,omitzero"`
	// (Optional) Unique identifier for the validation dataset
	ValidationDatasetID param.Opt[string] `json:"validation_dataset_id,omitzero"`
	paramObj
}

func (r PostTrainingSupervisedFineTuneParamsTrainingConfigDataConfig) MarshalJSON() (data []byte, err error) {
	type shadow PostTrainingSupervisedFineTuneParamsTrainingConfigDataConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PostTrainingSupervisedFineTuneParamsTrainingConfigDataConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PostTrainingSupervisedFineTuneParamsTrainingConfigDataConfig](
		"data_format", "instruct", "dialog",
	)
}

// (Optional) Configuration for memory and compute optimizations
type PostTrainingSupervisedFineTuneParamsTrainingConfigEfficiencyConfig struct {
	// (Optional) Whether to use activation checkpointing to reduce memory usage
	EnableActivationCheckpointing param.Opt[bool] `json:"enable_activation_checkpointing,omitzero"`
	// (Optional) Whether to offload activations to CPU to save GPU memory
	EnableActivationOffloading param.Opt[bool] `json:"enable_activation_offloading,omitzero"`
	// (Optional) Whether to offload FSDP parameters to CPU
	FsdpCPUOffload param.Opt[bool] `json:"fsdp_cpu_offload,omitzero"`
	// (Optional) Whether to use memory-efficient FSDP wrapping
	MemoryEfficientFsdpWrap param.Opt[bool] `json:"memory_efficient_fsdp_wrap,omitzero"`
	paramObj
}

func (r PostTrainingSupervisedFineTuneParamsTrainingConfigEfficiencyConfig) MarshalJSON() (data []byte, err error) {
	type shadow PostTrainingSupervisedFineTuneParamsTrainingConfigEfficiencyConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PostTrainingSupervisedFineTuneParamsTrainingConfigEfficiencyConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) Configuration for the optimization algorithm
//
// The properties Lr, NumWarmupSteps, OptimizerType, WeightDecay are required.
type PostTrainingSupervisedFineTuneParamsTrainingConfigOptimizerConfig struct {
	// Learning rate for the optimizer
	Lr float64 `json:"lr,required"`
	// Number of steps for learning rate warmup
	NumWarmupSteps int64 `json:"num_warmup_steps,required"`
	// Type of optimizer to use (adam, adamw, or sgd)
	//
	// Any of "adam", "adamw", "sgd".
	OptimizerType string `json:"optimizer_type,omitzero,required"`
	// Weight decay coefficient for regularization
	WeightDecay float64 `json:"weight_decay,required"`
	paramObj
}

func (r PostTrainingSupervisedFineTuneParamsTrainingConfigOptimizerConfig) MarshalJSON() (data []byte, err error) {
	type shadow PostTrainingSupervisedFineTuneParamsTrainingConfigOptimizerConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PostTrainingSupervisedFineTuneParamsTrainingConfigOptimizerConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PostTrainingSupervisedFineTuneParamsTrainingConfigOptimizerConfig](
		"optimizer_type", "adam", "adamw", "sgd",
	)
}
