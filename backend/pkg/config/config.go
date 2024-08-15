package config

import (
    "log"
	"context"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
)

func LoadConfig() aws.Config {
    cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
    if err != nil {
        log.Fatalf("unable to load SDK config, %v", err)
    }
    return cfg
}