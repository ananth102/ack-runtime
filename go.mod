module github.com/aws-controllers-k8s/runtime

go 1.14

require (
	github.com/aws/aws-sdk-go v1.37.10
	github.com/go-logr/logr v0.1.0
	github.com/google/go-cmp v0.3.1
	github.com/jaypipes/envutil v1.0.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.1.0
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.5.1
	go.uber.org/zap v1.10.0
	k8s.io/api v0.18.2
	k8s.io/apimachinery v0.18.6
	k8s.io/client-go v0.18.2
	sigs.k8s.io/controller-runtime v0.6.0
)
