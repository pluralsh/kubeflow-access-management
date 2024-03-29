/*
 * Kubeflow Auth
 *
 * Access Management API.
 *
 * API version: 0.1.0
 * Contact: kubeflow-engineering@google.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"flag"
	"net/http"

	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes/scheme"

	"github.com/kubeflow/kubeflow/components/access-management/kfam"
	profile "github.com/kubeflow/kubeflow/components/access-management/pkg/apis/kubeflow/v2alpha1"

	istioSecurityClient "istio.io/client-go/pkg/apis/security/v1beta1"
)

// kfam API assume coming request will contain user id in request header.
// set this parameter to specify header key containing user id.
const USERIDHEADER = "userid-header"

// set this parameter to specify header value prefix (if any) before user id.
const USERIDPREFIX = "userid-prefix"

// set cluster admin user id here.
const CLUSTERADMIN = "cluster-admin"

func main() {
	log.Printf("Server started")
	var userIdHeader string
	var userIdPrefix string
	var clusterAdmin string
	flag.StringVar(&userIdHeader, USERIDHEADER, "x-goog-authenticated-user-email", "Key of request header containing user id")
	flag.StringVar(&userIdPrefix, USERIDPREFIX, "accounts.google.com:", "Request header user id common prefix")
	flag.StringVar(&clusterAdmin, CLUSTERADMIN, "", "cluster admin")
	flag.Parse()

	profile.AddToScheme(scheme.Scheme)
	istioSecurityClient.AddToScheme(scheme.Scheme)

	profileClient, err := kfam.NewKfamClient(userIdHeader, userIdPrefix, clusterAdmin)
	if err != nil {
		log.Print(err)
		panic(err)
	}

	router := kfam.NewRouter(profileClient)

	log.Fatal(http.ListenAndServe(":8082", router))
}
