package models

import (
	"appengine"
	"appengine/datastore"
)

func AllTags(c appengine.Context) map[string]int {
	q := datastore.NewQuery("Entry").Project("Tags")
	t := q.Run(c)
	m := make(map[string]int, 0)
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

		for _, v := range e.Tags {
			m[v] += 1
		}
	}

	return m
}
