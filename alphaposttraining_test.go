// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/llamastack/llama-stack-client-go"
	"github.com/llamastack/llama-stack-client-go/internal/testutil"
	"github.com/llamastack/llama-stack-client-go/option"
)

func TestAlphaPostTrainingPreferenceOptimizeWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Alpha.PostTraining.PreferenceOptimize(context.TODO(), llamastackclient.AlphaPostTrainingPreferenceOptimizeParams{
		AlgorithmConfig: llamastackclient.AlphaPostTrainingPreferenceOptimizeParamsAlgorithmConfig{
			Beta:     0,
			LossType: "sigmoid",
		},
		FinetunedModel: "finetuned_model",
		HyperparamSearchConfig: map[string]llamastackclient.AlphaPostTrainingPreferenceOptimizeParamsHyperparamSearchConfigUnion{
			"foo": {
				OfBool: llamastackclient.Bool(true),
			},
		},
		JobUuid: "job_uuid",
		LoggerConfig: map[string]llamastackclient.AlphaPostTrainingPreferenceOptimizeParamsLoggerConfigUnion{
			"foo": {
				OfBool: llamastackclient.Bool(true),
			},
		},
		TrainingConfig: llamastackclient.AlphaPostTrainingPreferenceOptimizeParamsTrainingConfig{
			GradientAccumulationSteps: 0,
			MaxStepsPerEpoch:          0,
			NEpochs:                   0,
			DataConfig: llamastackclient.AlphaPostTrainingPreferenceOptimizeParamsTrainingConfigDataConfig{
				BatchSize:           0,
				DataFormat:          "instruct",
				DatasetID:           "dataset_id",
				Shuffle:             true,
				Packed:              llamastackclient.Bool(true),
				TrainOnInput:        llamastackclient.Bool(true),
				ValidationDatasetID: llamastackclient.String("validation_dataset_id"),
			},
			Dtype: llamastackclient.String("dtype"),
			EfficiencyConfig: llamastackclient.AlphaPostTrainingPreferenceOptimizeParamsTrainingConfigEfficiencyConfig{
				EnableActivationCheckpointing: llamastackclient.Bool(true),
				EnableActivationOffloading:    llamastackclient.Bool(true),
				FsdpCPUOffload:                llamastackclient.Bool(true),
				MemoryEfficientFsdpWrap:       llamastackclient.Bool(true),
			},
			MaxValidationSteps: llamastackclient.Int(0),
			OptimizerConfig: llamastackclient.AlphaPostTrainingPreferenceOptimizeParamsTrainingConfigOptimizerConfig{
				Lr:             0,
				NumWarmupSteps: 0,
				OptimizerType:  "adam",
				WeightDecay:    0,
			},
		},
	})
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAlphaPostTrainingSupervisedFineTuneWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Alpha.PostTraining.SupervisedFineTune(context.TODO(), llamastackclient.AlphaPostTrainingSupervisedFineTuneParams{
		HyperparamSearchConfig: map[string]llamastackclient.AlphaPostTrainingSupervisedFineTuneParamsHyperparamSearchConfigUnion{
			"foo": {
				OfBool: llamastackclient.Bool(true),
			},
		},
		JobUuid: "job_uuid",
		LoggerConfig: map[string]llamastackclient.AlphaPostTrainingSupervisedFineTuneParamsLoggerConfigUnion{
			"foo": {
				OfBool: llamastackclient.Bool(true),
			},
		},
		TrainingConfig: llamastackclient.AlphaPostTrainingSupervisedFineTuneParamsTrainingConfig{
			GradientAccumulationSteps: 0,
			MaxStepsPerEpoch:          0,
			NEpochs:                   0,
			DataConfig: llamastackclient.AlphaPostTrainingSupervisedFineTuneParamsTrainingConfigDataConfig{
				BatchSize:           0,
				DataFormat:          "instruct",
				DatasetID:           "dataset_id",
				Shuffle:             true,
				Packed:              llamastackclient.Bool(true),
				TrainOnInput:        llamastackclient.Bool(true),
				ValidationDatasetID: llamastackclient.String("validation_dataset_id"),
			},
			Dtype: llamastackclient.String("dtype"),
			EfficiencyConfig: llamastackclient.AlphaPostTrainingSupervisedFineTuneParamsTrainingConfigEfficiencyConfig{
				EnableActivationCheckpointing: llamastackclient.Bool(true),
				EnableActivationOffloading:    llamastackclient.Bool(true),
				FsdpCPUOffload:                llamastackclient.Bool(true),
				MemoryEfficientFsdpWrap:       llamastackclient.Bool(true),
			},
			MaxValidationSteps: llamastackclient.Int(0),
			OptimizerConfig: llamastackclient.AlphaPostTrainingSupervisedFineTuneParamsTrainingConfigOptimizerConfig{
				Lr:             0,
				NumWarmupSteps: 0,
				OptimizerType:  "adam",
				WeightDecay:    0,
			},
		},
		AlgorithmConfig: llamastackclient.AlgorithmConfigUnionParam{
			OfLoRa: &llamastackclient.AlgorithmConfigLoRaParam{
				Alpha:             0,
				ApplyLoraToMlp:    true,
				ApplyLoraToOutput: true,
				LoraAttnModules:   []string{"string"},
				Rank:              0,
				QuantizeBase:      llamastackclient.Bool(true),
				UseDora:           llamastackclient.Bool(true),
			},
		},
		CheckpointDir: llamastackclient.String("checkpoint_dir"),
		Model:         llamastackclient.String("model"),
	})
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
