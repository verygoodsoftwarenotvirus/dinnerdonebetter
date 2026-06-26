package api

import (
	"net/http"

	"github.com/dinnerdonebetter/dinnerdonebetter/backend/internal/domain/auth"
	paymentswebhook "github.com/dinnerdonebetter/dinnerdonebetter/backend/internal/services/payments/http"

	"github.com/primandproper/platform-go/encoding"
	"github.com/primandproper/platform-go/healthcheck"
	"github.com/primandproper/platform-go/observability/logging"
	"github.com/primandproper/platform-go/observability/metrics"
	"github.com/primandproper/platform-go/observability/tracing"
	"github.com/primandproper/platform-go/routing"
	routingcfg "github.com/primandproper/platform-go/routing/config"
	"github.com/primandproper/platform-go/version"
)

func ProvideAPIRouter(
	routingConfig routingcfg.Config,
	logger logging.Logger,
	tracerProvider tracing.TracerProvider,
	metricsProvider metrics.Provider,
	authService auth.AuthDataService,
	paymentsWebhookHandler *paymentswebhook.WebhookHandler,
	healthRegistry healthcheck.Registry,
) (routing.Router, error) {
	router, err := routingConfig.ProvideRouter(logger, tracerProvider, metricsProvider)
	if err != nil {
		return nil, err
	}

	encoder := encoding.ProvideServerEncoderDecoder(logger, tracerProvider, encoding.ContentTypeJSON)

	router.Route("/_ops_", func(metaRouter routing.Router) {
		// Expose a liveness check on /live
		metaRouter.Get("/live", func(res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(http.StatusOK)
		})

		// Expose a readiness check on /ready
		metaRouter.Get("/ready", func(res http.ResponseWriter, req *http.Request) {
			result := healthRegistry.CheckAll(req.Context())
			status := http.StatusOK
			if result.Status != healthcheck.StatusUp {
				status = http.StatusServiceUnavailable
			}
			encoder.EncodeResponseWithStatus(req.Context(), res, result, status)
		})

		metaRouter.Get("/version", func(res http.ResponseWriter, req *http.Request) {
			encoder.EncodeResponseWithStatus(req.Context(), res, version.Get(), http.StatusOK)
		})
	})

	router.Route("/oauth2", func(userRouter routing.Router) {
		userRouter.Get("/authorize", authService.AuthorizeHandler)
		userRouter.Post("/token", authService.TokenHandler)
		userRouter.Post("/revoke", authService.RevokeHandler)
	})

	router.Route("/api/payments/webhooks", func(paymentsRouter routing.Router) {
		paymentsRouter.Post("/{provider}", paymentsWebhookHandler.Handle)
	})

	return router, nil
}
