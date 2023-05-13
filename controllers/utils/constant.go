package utils

import "time"

const (
	// DefaultBackOff is the max backoff period, exported for the e2e test
	DefaultBackOff = 10 * time.Second
	// MaxBackOff is the max backoff period, exported for the e2e test
	MaxBackOff = 360 * time.Second
	// DefaultBackoffLimit for Jobs
	DefaultBackoffLimit = 0
	// DefaultRestartPolicy for Jobs
	DefaultRestartPolicy = "Never"
	// ZeroTTL
	ZeroTTL = 0
	// TTLDjustment
	TTLDjustment = 10
)

/*
GetTTL adjusts backoff in way that Polyaxon has timet to finalize the operation
*/
func GetTTL(ttlSecondsAfterFinished *int32) *int32 {
	jobTtlSecondsAfterFinished := ttlSecondsAfterFinished
	if ttlSecondsAfterFinished == nil {
		defaultTtl := int32(ZeroTTL + TTLDjustment)
		jobTtlSecondsAfterFinished = &defaultTtl
	} else {
		newTtl := int32(*ttlSecondsAfterFinished + TTLDjustment)
		jobTtlSecondsAfterFinished = &newTtl
	}
	return jobTtlSecondsAfterFinished
}

/*
GetBackoffLimit utils function to handle default case
*/
func GetBackoffLimit(backoffLimit *int32) *int32 {
	jobBackoffLimit := backoffLimit
	if backoffLimit == nil {
		defaultBackoffLimit := int32(DefaultBackoffLimit)
		jobBackoffLimit = &defaultBackoffLimit
	}
	return jobBackoffLimit
}
