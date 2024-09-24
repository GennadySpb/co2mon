// Code generated by sdkgen. DO NOT EDIT.

package smartcaptcha

import (
	"context"

	"google.golang.org/grpc"
)

// SmartCaptcha provides access to "smartcaptcha" component of Yandex.Cloud
type SmartCaptcha struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// NewSmartCaptcha creates instance of SmartCaptcha
func NewSmartCaptcha(g func(ctx context.Context) (*grpc.ClientConn, error)) *SmartCaptcha {
	return &SmartCaptcha{g}
}

// Captcha gets CaptchaService client
func (s *SmartCaptcha) Captcha() *CaptchaServiceClient {
	return &CaptchaServiceClient{getConn: s.getConn}
}
