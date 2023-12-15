package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	ycsdk "github.com/yandex-cloud/go-sdk"
	"go.uber.org/zap"
)

var log *zap.Logger

func init() {
	initLogger()
}

func initLogger() {
	config := zap.NewProductionConfig()
	config.DisableCaller = true
	config.Level.SetLevel(zap.InfoLevel)
	log, _ = config.Build()
}

func Handler(ctx context.Context, request interface{}) error {
	if err := SendToBudgetQueue(ctx); err != nil {
		return err
	}
	if err := StopComputeInstances(ctx); err != nil {
		return err
	}

	return nil
}

func SendToBudgetQueue(ctx context.Context) error {
	queue, err := newQueueClient(ctx)
	if err != nil {
		log.Error("Message Queue setup failed", zap.Error(err))
		return err
	}

	queueURL := os.Getenv("BUDGET_QUEUE_URL")
	message := "Budget trigger was triggered!"
	_, err = queue.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    &queueURL,
		MessageBody: &message,
	})
	if err != nil {
		log.Error("sending to Message Queue failed", zap.Error(err))
		return err

	}

	return nil
}

func newQueueClient(ctx context.Context) (*sqs.Client, error) {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:           "https://message-queue.api.cloud.yandex.net",
			SigningRegion: "ru-central1",
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithEndpointResolverWithOptions(customResolver),
	)
	if err != nil {
		return nil, err
	}

	return sqs.NewFromConfig(cfg), nil
}

func StopComputeInstances(ctx context.Context) error {
	folderID := requireEnvStr("FOLDER_ID")
	tag := requireEnvStr("TAG")

	sdk, err := ycsdk.Build(ctx, ycsdk.Config{
		Credentials: ycsdk.InstanceServiceAccount(),
	})
	if err != nil {
		log.Error("sdk setup failed", zap.Error(err))
		return err
	}
	var instances []*compute.Instance
	pageToken := ""
	for {
		listInstancesResponse, err := sdk.Compute().Instance().List(ctx, &compute.ListInstancesRequest{
			FolderId:  folderID,
			PageToken: pageToken,
		})
		if err != nil {
			log.Error("failed to list instances", zap.Error(err))
			return err
		}
		instances = append(instances, listInstancesResponse.GetInstances()...)
		if pageToken = listInstancesResponse.GetNextPageToken(); len(pageToken) == 0 {
			break
		}
	}

	stoppedCount := 0
	defer func() {
		log.Info(fmt.Sprintf("successfully stopped %d instances", stoppedCount))
	}()
	for _, i := range instances {
		log.Info("iam in " + tag)
		labels := i.Labels
		if v, ok := labels[tag]; ok && v == "true" && i.Status == compute.Instance_RUNNING {
			log.Info(fmt.Sprintf("stopping instance %s", i.GetId()))
			_, err := stopComputeInstance(ctx, sdk, i.GetId())
			if err != nil {
				log.Error(fmt.Sprintf("failed to stop instance %s", i.GetId()), zap.Error(err))
				return err
			}
			log.Info("iam in 2")
			stoppedCount++
		}
	}
	return nil
}

func stopComputeInstance(ctx context.Context, sdk *ycsdk.SDK, id string) (*operation.Operation, error) {
	return sdk.Compute().Instance().Stop(ctx, &compute.StopInstanceRequest{
		InstanceId: id,
	})
}

func requireEnvStr(envName string) string {
	res := os.Getenv(envName)
	if len(res) == 0 {
		log.Panic(fmt.Sprintf("env var is empty: %s", envName))
	}
	return res
}
