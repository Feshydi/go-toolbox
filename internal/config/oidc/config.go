package oidc_config

import (
	"context"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
)

func MustCreateProvider(issuer string, timeout time.Duration) *oidc.Provider {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	provider, err := oidc.NewProvider(ctx, issuer)
	if err != nil {
		panic("failed to construct a Provider: " + err.Error())
	}

	return provider
}
