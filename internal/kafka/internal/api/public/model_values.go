/*
 * Kafka Management API
 *
 * Kafka Management API is a REST API to manage Kafka instances
 *
 * API version: 1.11.0
 * Contact: rhosak-support@redhat.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package public

// Values struct for Values
type Values struct {
	Timestamp int64   `json:"timestamp,omitempty"`
	Value     float64 `json:"value"`
}
