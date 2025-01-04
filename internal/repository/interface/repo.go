package repositoryinterface

import (
	"context"
	"orderly/internal/domain/dto"
)

type Repository interface {
	GetCredential(ctx context.Context, username string, role string) (*dto.Credentials, error)
}
