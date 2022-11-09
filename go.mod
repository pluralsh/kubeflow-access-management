module github.com/kubeflow/kubeflow/components/access-management

go 1.16

require (
	github.com/gorilla/mux v1.8.0
	github.com/pluralsh/kubeflow-profile-controller v0.0.0-20210809114424-ef7e43acb86e
	github.com/prometheus/client_golang v1.11.0
	github.com/sirupsen/logrus v1.8.1
	istio.io/api v0.0.0-20221013011440-bc935762d2b9
	istio.io/client-go v1.15.3
	k8s.io/api v0.24.1
	k8s.io/apimachinery v0.24.1
	k8s.io/client-go v0.23.5
	sigs.k8s.io/controller-runtime v0.9.2
)
