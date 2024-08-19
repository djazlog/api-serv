package interceptor

import (
	"context"
	"google.golang.org/grpc"
)

type Validator interface {
	Validate() error
}

func ValidateInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	if v, ok := req.(Validator); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}
	return handler(ctx, req)
}
