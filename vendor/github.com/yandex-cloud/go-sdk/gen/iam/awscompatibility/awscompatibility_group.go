// Code generated by sdkgen. DO NOT EDIT.

package awscompatibility

import (
	"context"

	"google.golang.org/grpc"
)

// AWSCompatibility provides access to "awscompatibility" component of Yandex.Cloud
type AWSCompatibility struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// NewAWSCompatibility creates instance of AWSCompatibility
func NewAWSCompatibility(g func(ctx context.Context) (*grpc.ClientConn, error)) *AWSCompatibility {
	return &AWSCompatibility{g}
}

// AccessKey gets AccessKeyService client
func (a *AWSCompatibility) AccessKey() *AccessKeyServiceClient {
	return &AccessKeyServiceClient{getConn: a.getConn}
}
