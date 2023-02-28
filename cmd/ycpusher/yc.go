package main

import (
	"context"
	"fmt"
	"net"
	"strings"
	"time"

	ycsdk "github.com/yandex-cloud/go-sdk"
	"github.com/yandex-cloud/go-sdk/iamkey"
)

func Credentials() (ycsdk.Credentials, error) {
	if ServiceAccountKeyFile != "" && Token != "" {
		return nil, fmt.Errorf("only one of 'token' or 'sa-key-file' should be specified")
	}

	if ServiceAccountKeyFile != "" {
		key, err := iamkey.ReadFromJSONFile(ServiceAccountKeyFile)
		if err != nil {
			return nil, err
		}
		return ycsdk.ServiceAccountKey(key)
	}

	if Token != "" {
		if strings.HasPrefix(Token, "t1.") && strings.Count(Token, ".") == 2 {
			return ycsdk.NewIAMTokenCredentials(Token), nil
		}
		return ycsdk.OAuthToken(Token), nil
	}

	if sa := ycsdk.InstanceServiceAccount(); checkServiceAccountAvailable(context.Background(), sa) {
		fmt.Println("Trying to get Instance Service Account.")
		return sa, nil
	}

	return nil, fmt.Errorf("one of 'token' or 'sa-key-file' should be specified; if you are inside compute instance, you can attach service account to it in order to authenticate via instance service account")
}

func checkServiceAccountAvailable(ctx context.Context, sa ycsdk.NonExchangeableCredentials) bool {
	dialer := net.Dialer{Timeout: 50 * time.Millisecond}
	conn, err := dialer.Dial("tcp", net.JoinHostPort(ycsdk.InstanceMetadataAddr, "80"))
	if err != nil {
		return false
	}
	_ = conn.Close()
	_, err = sa.IAMToken(ctx)
	return err == nil
}
