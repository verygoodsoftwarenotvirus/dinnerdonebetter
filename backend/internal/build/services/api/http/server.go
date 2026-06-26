package api

import (
	"github.com/dinnerdonebetter/dinnerdonebetter/backend/internal/domain/auth"
	paymentswebhook "github.com/dinnerdonebetter/dinnerdonebetter/backend/internal/services/payments/http"

	"github.com/primandproper/platform-go/healthcheck"
	"github.com/primandproper/platform-go/observability/logging"
	"github.com/primandproper/platform-go/observability/metrics"
	"github.com/primandproper/platform-go/observability/tracing"
	"github.com/primandproper/platform-go/routing"
	routingcfg "github.com/primandproper/platform-go/routing/config"

	"github.com/samber/do/v2"
)

// RegisterAPIRouter registers the API router provider with the injector.
func RegisterAPIRouter(i do.Injector) {
	do.Provide[routing.Router](i, func(i do.Injector) (routing.Router, error) {
		return ProvideAPIRouter(
			*do.MustInvoke[*routingcfg.Config](i),
			do.MustInvoke[logging.Logger](i),
			do.MustInvoke[tracing.TracerProvider](i),
			do.MustInvoke[metrics.Provider](i),
			do.MustInvoke[auth.AuthDataService](i),
			do.MustInvoke[*paymentswebhook.WebhookHandler](i),
			do.MustInvoke[healthcheck.Registry](i),
		)
	})
}
