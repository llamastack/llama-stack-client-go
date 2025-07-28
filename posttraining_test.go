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

func TestPostTrainingPreferenceOptimizeWithOptionalParams(t *testing.T) {
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
	_, err := client.PostTraining.PreferenceOptimize(context.TODO(), llamastackclient.PostTrainingPreferenceOptimizeParams{
		AlgorithmConfig: llamastackclient.PostTrainingPreferenceOptimizeParamsAlgorithmConfig{
			Beta:     0,
			LossType: "sigmoid",
		},
		FinetunedModel: "finetuned_model",
		HyperparamSearchConfig: map[string]llamastackclient.PostTrainingPreferenceOptimizeParamsHyperparamSearchConfigUnion{
			"foo": {
				OfBool: llamastackclient.Bool(true),
			},
		},
		JobUuid: "job_uuid",
		LoggerConfig: map[string]llamastackclient.PostTrainingPreferenceOptimizeParamsLoggerConfigUnion{
			"foo": {
				OfBool: llamastackclient.Bool(true),
			},
		},
		TrainingConfig: llamastackclient.PostTrainingPreferenceOptimizeParamsTrainingConfig{
			GradientAccumulationSteps: 0,
			MaxStepsPerEpoch:          0,
			NEpochs:                   0,
			DataConfig: llamastackclient.PostTrainingPreferenceOptimizeParamsTrainingConfigDataConfig{
				BatchSize:           0,
				DataFormat:          "instruct",
				DatasetID:           "dataset_id",
				Shuffle:             true,
				Packed:              llamastackclient.Bool(true),
				TrainOnInput:        llamastackclient.Bool(true),
				ValidationDatasetID: llamastackclient.String("validation_dataset_id"),
			},
			Dtype: llamastackclient.String("dtype"),
			EfficiencyConfig: llamastackclient.PostTrainingPreferenceOptimizeParamsTrainingConfigEfficiencyConfig{
				EnableActivationCheckpointing: llamastackclient.Bool(true),
				EnableActivationOffloading:    llamastackclient.Bool(true),
				FsdpCPUOffload:                llamastackclient.Bool(true),
				MemoryEfficientFsdpWrap:       llamastackclient.Bool(true),
			},
			MaxValidationSteps: llamastackclient.Int(0),
			OptimizerConfig: llamastackclient.PostTrainingPreferenceOptimizeParamsTrainingConfigOptimizerConfig{
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

func TestPostTrainingSupervisedFineTuneWithOptionalParams(t *testing.T) {
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
	_, err := client.PostTraining.SupervisedFineTune(context.TODO(), llamastackclient.PostTrainingSupervisedFineTuneParams{
		HyperparamSearchConfig: map[string]llamastackclient.PostTrainingSupervisedFineTuneParamsHyperparamSearchConfigUnion{
			"foo": {
				OfBool: llamastackclient.Bool(true),
			},
		},
		JobUuid: "job_uuid",
		LoggerConfig: map[string]llamastackclient.PostTrainingSupervisedFineTuneParamsLoggerConfigUnion{
			"foo": {
				OfBool: llamastackclient.Bool(true),
			},
		},
		TrainingConfig: llamastackclient.PostTrainingSupervisedFineTuneParamsTrainingConfig{
			GradientAccumulationSteps: 0,
			MaxStepsPerEpoch:          0,
			NEpochs:                   0,
			DataConfig: llamastackclient.PostTrainingSupervisedFineTuneParamsTrainingConfigDataConfig{
				BatchSize:           0,
				DataFormat:          "instruct",
				DatasetID:           "dataset_id",
				Shuffle:             true,
				Packed:              llamastackclient.Bool(true),
				TrainOnInput:        llamastackclient.Bool(true),
				ValidationDatasetID: llamastackclient.String("validation_dataset_id"),
			},
			Dtype: llamastackclient.String("dtype"),
			EfficiencyConfig: llamastackclient.PostTrainingSupervisedFineTuneParamsTrainingConfigEfficiencyConfig{
				EnableActivationCheckpointing: llamastackclient.Bool(true),
				EnableActivationOffloading:    llamastackclient.Bool(true),
				FsdpCPUOffload:                llamastackclient.Bool(true),
				MemoryEfficientFsdpWrap:       llamastackclient.Bool(true),
			},
			MaxValidationSteps: llamastackclient.Int(0),
			OptimizerConfig: llamastackclient.PostTrainingSupervisedFineTuneParamsTrainingConfigOptimizerConfig{
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
