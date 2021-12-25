/*
 * Connector Service Fleet Manager Private APIs
 *
 * Connector Service Fleet Manager apis that are used by internal services.
 *
 * API version: 0.0.3
 * Contact: foo@bar
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package private

// ConnectorDeployment Holds the deployment configuration of a connector
type ConnectorDeployment struct {
	Id       string                           `json:"id,omitempty"`
	Kind     string                           `json:"kind,omitempty"`
	Href     string                           `json:"href,omitempty"`
	Metadata ConnectorDeploymentAllOfMetadata `json:"metadata,omitempty"`
	Spec     ConnectorDeploymentSpec          `json:"spec,omitempty"`
	Status   ConnectorDeploymentStatus        `json:"status,omitempty"`
}
