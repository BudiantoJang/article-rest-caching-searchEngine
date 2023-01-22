package port

import "context"

type Validation interface {
	ValidateRequest(ctx context.Context, req interface{}) error
}
