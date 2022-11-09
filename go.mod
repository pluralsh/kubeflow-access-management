module github.com/kubeflow/kubeflow/components/access-management

go 1.16

require (
	github.com/gorilla/mux v1.8.0
	github.com/pluralsh/kubeflow-profile-controller v0.0.0-20220909092542-a99665faeef0
	github.com/prometheus/client_golang v1.12.1
	github.com/sirupsen/logrus v1.8.1
	istio.io/api v0.0.0-20220418200313-202d5d40987b
	istio.io/client-go v1.13.3
	k8s.io/api v0.23.5
	k8s.io/apimachinery v0.23.5
	k8s.io/client-go v0.23.5
	sigs.k8s.io/controller-runtime v0.11.2
)
