// Copyright (c) Meta Platforms, Inc. and affiliates.
// All rights reserved.
//
// This source code is licensed under the terms described in the LICENSE file in
// the root directory of this source tree.
//
// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"net/http"
	"slices"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
)

// AlphaPostTrainingService contains methods and other services that help with
// interacting with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAlphaPostTrainingService] method instead.
type AlphaPostTrainingService struct {
	Options []option.RequestOption
	Job     AlphaPostTrainingJobService
}

// NewAlphaPostTrainingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAlphaPostTrainingService(opts ...option.RequestOption) (r AlphaPostTrainingService) {
	r = AlphaPostTrainingService{}
	r.Options = opts
	r.Job = NewAlphaPostTrainingJobService(opts...)
	return
}

// Run preference optimization of a model.
func (r *AlphaPostTrainingService) PreferenceOptimize(ctx context.Context, body AlphaPostTrainingPreferenceOptimizeParams, opts ...option.RequestOption) (res *PostTrainingJob, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1alpha/post-training/preference-optimize"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Run supervised fine-tuning of a model.
func (r *AlphaPostTrainingService) SupervisedFineTune(ctx context.Context, body AlphaPostTrainingSupervisedFineTuneParams, opts ...option.RequestOption) (res *PostTrainingJob, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1alpha/post-training/supervised-fine-tune"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
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

type AlphaPostTrainingPreferenceOptimizeParams struct {
	// Configuration for Direct Preference Optimization (DPO) alignment.
	AlgorithmConfig        AlphaPostTrainingPreferenceOptimizeParamsAlgorithmConfig `json:"algorithm_config,omitzero,required"`
	FinetunedModel         string                                                   `json:"finetuned_model,required"`
	HyperparamSearchConfig map[string]any                                           `json:"hyperparam_search_config,omitzero,required"`
	JobUuid                string                                                   `json:"job_uuid,required"`
	LoggerConfig           map[string]any                                           `json:"logger_config,omitzero,required"`
	// Comprehensive configuration for the training process.
	TrainingConfig AlphaPostTrainingPreferenceOptimizeParamsTrainingConfig `json:"training_config,omitzero,required"`
	paramObj
}

func (r AlphaPostTrainingPreferenceOptimizeParams) MarshalJSON() (data []byte, err error) {
	type shadow AlphaPostTrainingPreferenceOptimizeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaPostTrainingPreferenceOptimizeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for Direct Preference Optimization (DPO) alignment.
//
// The property Beta is required.
type AlphaPostTrainingPreferenceOptimizeParamsAlgorithmConfig struct {
	Beta float64 `json:"beta,required"`
	// Any of "sigmoid", "hinge", "ipo", "kto_pair".
	LossType string `json:"loss_type,omitzero"`
	paramObj
}

func (r AlphaPostTrainingPreferenceOptimizeParamsAlgorithmConfig) MarshalJSON() (data []byte, err error) {
	type shadow AlphaPostTrainingPreferenceOptimizeParamsAlgorithmConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaPostTrainingPreferenceOptimizeParamsAlgorithmConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AlphaPostTrainingPreferenceOptimizeParamsAlgorithmConfig](
		"loss_type", "sigmoid", "hinge", "ipo", "kto_pair",
	)
}

// Comprehensive configuration for the training process.
//
// The property NEpochs is required.
type AlphaPostTrainingPreferenceOptimizeParamsTrainingConfig struct {
	NEpochs                   int64             `json:"n_epochs,required"`
	Dtype                     param.Opt[string] `json:"dtype,omitzero"`
	MaxValidationSteps        param.Opt[int64]  `json:"max_validation_steps,omitzero"`
	GradientAccumulationSteps param.Opt[int64]  `json:"gradient_accumulation_steps,omitzero"`
	MaxStepsPerEpoch          param.Opt[int64]  `json:"max_steps_per_epoch,omitzero"`
	// Configuration for training data and data loading.
	DataConfig AlphaPostTrainingPreferenceOptimizeParamsTrainingConfigDataConfig `json:"data_config,omitzero"`
	// Configuration for memory and compute efficiency optimizations.
	EfficiencyConfig AlphaPostTrainingPreferenceOptimizeParamsTrainingConfigEfficiencyConfig `json:"efficiency_config,omitzero"`
	// Configuration parameters for the optimization algorithm.
	OptimizerConfig AlphaPostTrainingPreferenceOptimizeParamsTrainingConfigOptimizerConfig `json:"optimizer_config,omitzero"`
	paramObj
}

func (r AlphaPostTrainingPreferenceOptimizeParamsTrainingConfig) MarshalJSON() (data []byte, err error) {
	type shadow AlphaPostTrainingPreferenceOptimizeParamsTrainingConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaPostTrainingPreferenceOptimizeParamsTrainingConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for training data and data loading.
//
// The properties BatchSize, DataFormat, DatasetID, Shuffle are required.
type AlphaPostTrainingPreferenceOptimizeParamsTrainingConfigDataConfig struct {
	BatchSize int64 `json:"batch_size,required"`
	// Format of the training dataset.
	//
	// Any of "instruct", "dialog".
	DataFormat          string            `json:"data_format,omitzero,required"`
	DatasetID           string            `json:"dataset_id,required"`
	Shuffle             bool              `json:"shuffle,required"`
	Packed              param.Opt[bool]   `json:"packed,omitzero"`
	TrainOnInput        param.Opt[bool]   `json:"train_on_input,omitzero"`
	ValidationDatasetID param.Opt[string] `json:"validation_dataset_id,omitzero"`
	paramObj
}

func (r AlphaPostTrainingPreferenceOptimizeParamsTrainingConfigDataConfig) MarshalJSON() (data []byte, err error) {
	type shadow AlphaPostTrainingPreferenceOptimizeParamsTrainingConfigDataConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaPostTrainingPreferenceOptimizeParamsTrainingConfigDataConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AlphaPostTrainingPreferenceOptimizeParamsTrainingConfigDataConfig](
		"data_format", "instruct", "dialog",
	)
}

// Configuration for memory and compute efficiency optimizations.
type AlphaPostTrainingPreferenceOptimizeParamsTrainingConfigEfficiencyConfig struct {
	EnableActivationCheckpointing param.Opt[bool] `json:"enable_activation_checkpointing,omitzero"`
	EnableActivationOffloading    param.Opt[bool] `json:"enable_activation_offloading,omitzero"`
	FsdpCPUOffload                param.Opt[bool] `json:"fsdp_cpu_offload,omitzero"`
	MemoryEfficientFsdpWrap       param.Opt[bool] `json:"memory_efficient_fsdp_wrap,omitzero"`
	paramObj
}

func (r AlphaPostTrainingPreferenceOptimizeParamsTrainingConfigEfficiencyConfig) MarshalJSON() (data []byte, err error) {
	type shadow AlphaPostTrainingPreferenceOptimizeParamsTrainingConfigEfficiencyConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaPostTrainingPreferenceOptimizeParamsTrainingConfigEfficiencyConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration parameters for the optimization algorithm.
//
// The properties Lr, NumWarmupSteps, OptimizerType, WeightDecay are required.
type AlphaPostTrainingPreferenceOptimizeParamsTrainingConfigOptimizerConfig struct {
	Lr             float64 `json:"lr,required"`
	NumWarmupSteps int64   `json:"num_warmup_steps,required"`
	// Available optimizer algorithms for training.
	//
	// Any of "adam", "adamw", "sgd".
	OptimizerType string  `json:"optimizer_type,omitzero,required"`
	WeightDecay   float64 `json:"weight_decay,required"`
	paramObj
}

func (r AlphaPostTrainingPreferenceOptimizeParamsTrainingConfigOptimizerConfig) MarshalJSON() (data []byte, err error) {
	type shadow AlphaPostTrainingPreferenceOptimizeParamsTrainingConfigOptimizerConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaPostTrainingPreferenceOptimizeParamsTrainingConfigOptimizerConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AlphaPostTrainingPreferenceOptimizeParamsTrainingConfigOptimizerConfig](
		"optimizer_type", "adam", "adamw", "sgd",
	)
}

type AlphaPostTrainingSupervisedFineTuneParams struct {
	HyperparamSearchConfig map[string]any `json:"hyperparam_search_config,omitzero,required"`
	JobUuid                string         `json:"job_uuid,required"`
	LoggerConfig           map[string]any `json:"logger_config,omitzero,required"`
	// Comprehensive configuration for the training process.
	TrainingConfig AlphaPostTrainingSupervisedFineTuneParamsTrainingConfig `json:"training_config,omitzero,required"`
	CheckpointDir  param.Opt[string]                                       `json:"checkpoint_dir,omitzero"`
	// Model descriptor for training if not in provider config`
	Model param.Opt[string] `json:"model,omitzero"`
	// Configuration for Low-Rank Adaptation (LoRA) fine-tuning.
	AlgorithmConfig AlphaPostTrainingSupervisedFineTuneParamsAlgorithmConfigUnion `json:"algorithm_config,omitzero"`
	paramObj
}

func (r AlphaPostTrainingSupervisedFineTuneParams) MarshalJSON() (data []byte, err error) {
	type shadow AlphaPostTrainingSupervisedFineTuneParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaPostTrainingSupervisedFineTuneParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Comprehensive configuration for the training process.
//
// The property NEpochs is required.
type AlphaPostTrainingSupervisedFineTuneParamsTrainingConfig struct {
	NEpochs                   int64             `json:"n_epochs,required"`
	Dtype                     param.Opt[string] `json:"dtype,omitzero"`
	MaxValidationSteps        param.Opt[int64]  `json:"max_validation_steps,omitzero"`
	GradientAccumulationSteps param.Opt[int64]  `json:"gradient_accumulation_steps,omitzero"`
	MaxStepsPerEpoch          param.Opt[int64]  `json:"max_steps_per_epoch,omitzero"`
	// Configuration for training data and data loading.
	DataConfig AlphaPostTrainingSupervisedFineTuneParamsTrainingConfigDataConfig `json:"data_config,omitzero"`
	// Configuration for memory and compute efficiency optimizations.
	EfficiencyConfig AlphaPostTrainingSupervisedFineTuneParamsTrainingConfigEfficiencyConfig `json:"efficiency_config,omitzero"`
	// Configuration parameters for the optimization algorithm.
	OptimizerConfig AlphaPostTrainingSupervisedFineTuneParamsTrainingConfigOptimizerConfig `json:"optimizer_config,omitzero"`
	paramObj
}

func (r AlphaPostTrainingSupervisedFineTuneParamsTrainingConfig) MarshalJSON() (data []byte, err error) {
	type shadow AlphaPostTrainingSupervisedFineTuneParamsTrainingConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaPostTrainingSupervisedFineTuneParamsTrainingConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for training data and data loading.
//
// The properties BatchSize, DataFormat, DatasetID, Shuffle are required.
type AlphaPostTrainingSupervisedFineTuneParamsTrainingConfigDataConfig struct {
	BatchSize int64 `json:"batch_size,required"`
	// Format of the training dataset.
	//
	// Any of "instruct", "dialog".
	DataFormat          string            `json:"data_format,omitzero,required"`
	DatasetID           string            `json:"dataset_id,required"`
	Shuffle             bool              `json:"shuffle,required"`
	Packed              param.Opt[bool]   `json:"packed,omitzero"`
	TrainOnInput        param.Opt[bool]   `json:"train_on_input,omitzero"`
	ValidationDatasetID param.Opt[string] `json:"validation_dataset_id,omitzero"`
	paramObj
}

func (r AlphaPostTrainingSupervisedFineTuneParamsTrainingConfigDataConfig) MarshalJSON() (data []byte, err error) {
	type shadow AlphaPostTrainingSupervisedFineTuneParamsTrainingConfigDataConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaPostTrainingSupervisedFineTuneParamsTrainingConfigDataConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AlphaPostTrainingSupervisedFineTuneParamsTrainingConfigDataConfig](
		"data_format", "instruct", "dialog",
	)
}

// Configuration for memory and compute efficiency optimizations.
type AlphaPostTrainingSupervisedFineTuneParamsTrainingConfigEfficiencyConfig struct {
	EnableActivationCheckpointing param.Opt[bool] `json:"enable_activation_checkpointing,omitzero"`
	EnableActivationOffloading    param.Opt[bool] `json:"enable_activation_offloading,omitzero"`
	FsdpCPUOffload                param.Opt[bool] `json:"fsdp_cpu_offload,omitzero"`
	MemoryEfficientFsdpWrap       param.Opt[bool] `json:"memory_efficient_fsdp_wrap,omitzero"`
	paramObj
}

func (r AlphaPostTrainingSupervisedFineTuneParamsTrainingConfigEfficiencyConfig) MarshalJSON() (data []byte, err error) {
	type shadow AlphaPostTrainingSupervisedFineTuneParamsTrainingConfigEfficiencyConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaPostTrainingSupervisedFineTuneParamsTrainingConfigEfficiencyConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration parameters for the optimization algorithm.
//
// The properties Lr, NumWarmupSteps, OptimizerType, WeightDecay are required.
type AlphaPostTrainingSupervisedFineTuneParamsTrainingConfigOptimizerConfig struct {
	Lr             float64 `json:"lr,required"`
	NumWarmupSteps int64   `json:"num_warmup_steps,required"`
	// Available optimizer algorithms for training.
	//
	// Any of "adam", "adamw", "sgd".
	OptimizerType string  `json:"optimizer_type,omitzero,required"`
	WeightDecay   float64 `json:"weight_decay,required"`
	paramObj
}

func (r AlphaPostTrainingSupervisedFineTuneParamsTrainingConfigOptimizerConfig) MarshalJSON() (data []byte, err error) {
	type shadow AlphaPostTrainingSupervisedFineTuneParamsTrainingConfigOptimizerConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaPostTrainingSupervisedFineTuneParamsTrainingConfigOptimizerConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AlphaPostTrainingSupervisedFineTuneParamsTrainingConfigOptimizerConfig](
		"optimizer_type", "adam", "adamw", "sgd",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AlphaPostTrainingSupervisedFineTuneParamsAlgorithmConfigUnion struct {
	OfLoRa *AlphaPostTrainingSupervisedFineTuneParamsAlgorithmConfigLoRa `json:",omitzero,inline"`
	OfQat  *AlphaPostTrainingSupervisedFineTuneParamsAlgorithmConfigQat  `json:",omitzero,inline"`
	paramUnion
}

func (u AlphaPostTrainingSupervisedFineTuneParamsAlgorithmConfigUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfLoRa, u.OfQat)
}
func (u *AlphaPostTrainingSupervisedFineTuneParamsAlgorithmConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AlphaPostTrainingSupervisedFineTuneParamsAlgorithmConfigUnion) asAny() any {
	if !param.IsOmitted(u.OfLoRa) {
		return u.OfLoRa
	} else if !param.IsOmitted(u.OfQat) {
		return u.OfQat
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AlphaPostTrainingSupervisedFineTuneParamsAlgorithmConfigUnion) GetAlpha() *int64 {
	if vt := u.OfLoRa; vt != nil {
		return &vt.Alpha
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AlphaPostTrainingSupervisedFineTuneParamsAlgorithmConfigUnion) GetApplyLoraToMlp() *bool {
	if vt := u.OfLoRa; vt != nil {
		return &vt.ApplyLoraToMlp
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AlphaPostTrainingSupervisedFineTuneParamsAlgorithmConfigUnion) GetApplyLoraToOutput() *bool {
	if vt := u.OfLoRa; vt != nil {
		return &vt.ApplyLoraToOutput
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AlphaPostTrainingSupervisedFineTuneParamsAlgorithmConfigUnion) GetLoraAttnModules() []string {
	if vt := u.OfLoRa; vt != nil {
		return vt.LoraAttnModules
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AlphaPostTrainingSupervisedFineTuneParamsAlgorithmConfigUnion) GetRank() *int64 {
	if vt := u.OfLoRa; vt != nil {
		return &vt.Rank
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AlphaPostTrainingSupervisedFineTuneParamsAlgorithmConfigUnion) GetQuantizeBase() *bool {
	if vt := u.OfLoRa; vt != nil && vt.QuantizeBase.Valid() {
		return &vt.QuantizeBase.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AlphaPostTrainingSupervisedFineTuneParamsAlgorithmConfigUnion) GetUseDora() *bool {
	if vt := u.OfLoRa; vt != nil && vt.UseDora.Valid() {
		return &vt.UseDora.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AlphaPostTrainingSupervisedFineTuneParamsAlgorithmConfigUnion) GetGroupSize() *int64 {
	if vt := u.OfQat; vt != nil {
		return &vt.GroupSize
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AlphaPostTrainingSupervisedFineTuneParamsAlgorithmConfigUnion) GetQuantizerName() *string {
	if vt := u.OfQat; vt != nil {
		return &vt.QuantizerName
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AlphaPostTrainingSupervisedFineTuneParamsAlgorithmConfigUnion) GetType() *string {
	if vt := u.OfLoRa; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfQat; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[AlphaPostTrainingSupervisedFineTuneParamsAlgorithmConfigUnion](
		"type",
		apijson.Discriminator[AlphaPostTrainingSupervisedFineTuneParamsAlgorithmConfigLoRa]("LoRA"),
		apijson.Discriminator[AlphaPostTrainingSupervisedFineTuneParamsAlgorithmConfigQat]("QAT"),
	)
}

// Configuration for Low-Rank Adaptation (LoRA) fine-tuning.
//
// The properties Alpha, ApplyLoraToMlp, ApplyLoraToOutput, LoraAttnModules, Rank
// are required.
type AlphaPostTrainingSupervisedFineTuneParamsAlgorithmConfigLoRa struct {
	Alpha             int64           `json:"alpha,required"`
	ApplyLoraToMlp    bool            `json:"apply_lora_to_mlp,required"`
	ApplyLoraToOutput bool            `json:"apply_lora_to_output,required"`
	LoraAttnModules   []string        `json:"lora_attn_modules,omitzero,required"`
	Rank              int64           `json:"rank,required"`
	QuantizeBase      param.Opt[bool] `json:"quantize_base,omitzero"`
	UseDora           param.Opt[bool] `json:"use_dora,omitzero"`
	// Any of "LoRA".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r AlphaPostTrainingSupervisedFineTuneParamsAlgorithmConfigLoRa) MarshalJSON() (data []byte, err error) {
	type shadow AlphaPostTrainingSupervisedFineTuneParamsAlgorithmConfigLoRa
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaPostTrainingSupervisedFineTuneParamsAlgorithmConfigLoRa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AlphaPostTrainingSupervisedFineTuneParamsAlgorithmConfigLoRa](
		"type", "LoRA",
	)
}

// Configuration for Quantization-Aware Training (QAT) fine-tuning.
//
// The properties GroupSize, QuantizerName are required.
type AlphaPostTrainingSupervisedFineTuneParamsAlgorithmConfigQat struct {
	GroupSize     int64  `json:"group_size,required"`
	QuantizerName string `json:"quantizer_name,required"`
	// Any of "QAT".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r AlphaPostTrainingSupervisedFineTuneParamsAlgorithmConfigQat) MarshalJSON() (data []byte, err error) {
	type shadow AlphaPostTrainingSupervisedFineTuneParamsAlgorithmConfigQat
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaPostTrainingSupervisedFineTuneParamsAlgorithmConfigQat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AlphaPostTrainingSupervisedFineTuneParamsAlgorithmConfigQat](
		"type", "QAT",
	)
}
