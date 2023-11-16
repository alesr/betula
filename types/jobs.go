package types

import "time"

// If you make something drastic to this file, reflect the changes in Adding a new job.md

type JobCategory string

const (
	NotifyAboutMyRepost   JobCategory = "notify about my repost"
	VerifyTheirRepost     JobCategory = "verify their repost"
	ReceiveUnrepost       JobCategory = "receive unrepost"
	NotifyAboutMyUnrepost JobCategory = "notify about my unrepost"
)

// Job is a task for Betula to do later.
type Job struct {
	// ID is a unique identifier for the Job. You get it when reading from the database. Do not set it when issuing a new job.
	ID       int64
	Category JobCategory
	Due      time.Time
	// Payload is some data.
	Payload any
}