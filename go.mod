module github.com/kubeflow/kubeflow/components/access-management

go 1.16

require (
	github.com/gorilla/mux v1.8.0
	github.com/pluralsh/kubeflow-profile-controller v0.0.0-20210809114424-ef7e43acb86e
	github.com/prometheus/client_golang v1.12.2
	github.com/sirupsen/logrus v1.8.1
	istio.io/api v0.0.0-20210713195055-3a340e4f154e
	istio.io/client-go v1.10.3
	k8s.io/api v0.25.0
	k8s.io/apimachinery v0.25.0
	k8s.io/client-go v0.25.0
	sigs.k8s.io/controller-runtime v0.13.1
)
