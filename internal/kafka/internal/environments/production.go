package environments

import "github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/environments"

func NewProductionEnvLoader() environments.EnvLoader {
	return environments.SimpleEnvLoader{
		"v":                                 "1",
		"ocm-debug":                         "false",
		"enable-ocm-mock":                   "false",
		"enable-sentry":                     "true",
		"enable-deny-list":                  "true",
		"enable-access-list":                "false",
		"max-allowed-instances":             "1",
		"mas-sso-realm":                     "rhoas",
		"mas-sso-base-url":                  "https://identity.api.openshift.com",
		"enable-kafka-external-certificate": "true",
		"aws-cluster-compute-machine-type":  "m5.2xlarge",
	}
}
