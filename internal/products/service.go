package products

import "context"

type Service interface {
	ListProducts(ctx context.Context) (string, error)
}

type svc struct {
	// repository

}

// In here we are returning svc struct but in func signature we are giving interface
func NewService() Service {
	return &svc{}
}

func (svc *svc) ListProducts(ctx context.Context) (string, error) {

	return "", nil
}
