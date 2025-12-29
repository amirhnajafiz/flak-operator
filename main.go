package main

import (
	"net/http"

	"github.com/amirhnajafiz/flap-operator/include/cmd"
	"github.com/amirhnajafiz/flap-operator/include/configs"
	"github.com/amirhnajafiz/flap-operator/include/handlers"
	"github.com/amirhnajafiz/flap-operator/include/telemetry/logging"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
)

var (
	// required variables for working with kubernetes runtime objects
	scheme = runtime.NewScheme()
	codecs = serializer.NewCodecFactory(scheme)
)

func init() {
	_ = corev1.AddToScheme(scheme)
}

func main() {
	// load config parameters
	cfg := configs.LoadConfigs()

	// set logrus logger
	logging.SetLogger(cfg.Logger.Level, cfg.Logger.JSON)

	// register http handlers
	http.HandleFunc("/health", handlers.Health)
	http.HandleFunc("/mutate", handlers.MutatePods(codecs))

	// listens to clear text http unless TLS env var is set to "true"
	if cfg.TLS.Enable {
		cmd.ServeHTTPS(cfg.TLS.CertPath, cfg.TLS.KeyPath)
	} else {
		cmd.ServeHTTP()
	}
}
