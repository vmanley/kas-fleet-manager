/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage Kafka instances.
 *
 * API version: 1.7.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package public

// MetricsRangeQueryList struct for MetricsRangeQueryList
type MetricsRangeQueryList struct {
	Kind  string       `json:"kind,omitempty"`
	Id    string       `json:"id,omitempty"`
	Items []RangeQuery `json:"items,omitempty"`
}
