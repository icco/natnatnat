package models

import (
	"appengine"
	"appengine/datastore"
)

func AllTags(c appengine.Context) {
	q := datastore.NewQuery("Entry").Project("Tags")
	t := q.Run(c)
	for {
		var e Entry
		_, err := t.Next(&e)
		if err == datastore.Done {
			break
		}
		if err != nil {
			c.Errorf("Running query: %v", err)
			break
		}
		c.Infof("Got Tags: %v", e.Tags)
	}
}
