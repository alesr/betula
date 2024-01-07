package activities

import (
	"encoding/json"
	"fmt"
	"git.sr.ht/~bouncepaw/betula/settings"
	"git.sr.ht/~bouncepaw/betula/stricks"
)

// NewAccept wraps the acceptedActivity in an Accept activity.
// The @context of the wrapped activity is deleted.
func NewAccept(acceptedActivity dict) ([]byte, error) {
	delete(acceptedActivity, "@context")
	return json.Marshal(dict{
		"@context": atContext,
		"id":       fmt.Sprintf("%s/temp/%s", settings.SiteURL(), stricks.RandomWhatever()),
		"type":     "Accept",
		"actor":    betulaActor,
		"object":   acceptedActivity,
	})
}

type AcceptReport struct {
	ActorID  string
	ObjectID string
	Object   dict
}

func guessAccept(activity dict) (any, error) {
	report := AcceptReport{
		ActorID:  getIDSomehow(activity, "actor"),
		ObjectID: getIDSomehow(activity, "object"),
	}
	if report.ActorID == "" {
		return nil, ErrNoActor
	}
	if report.ObjectID == "" {
		return nil, ErrNoObject
	}
	if obj, ok := activity["object"]; ok {
		switch v := obj.(type) {
		case dict:
			report.Object = v
		}
	}

	return report, nil
}
