package models

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"time"
)

type Mention struct {
	Source   string
	Target   string
	Recieved time.Time
	Verified bool
}

func NewMention(c appengine.Context, source string, target string) *Mention {
	var k *datastore.Key
	var e *Mention

	k = GetMention(c, source, target)
	if k == nil {
		k = datastore.NewIncompleteKey(c, "Mention", nil)
		e = new(Mention)
		e.Source = source
		e.target = target
		e.Recieved = time.Now()
		e.Verified = false
	}

	return e
}

func MentionExists(c appengine.Context, source string, target string) bool {
	return GetMention(c, source, target) != nil
}

func GetMention(c appengine.Context, source string, target string) *datastore.Key {
	q := datastore.NewQuery("Mention").Filter("source =", source).Filter("target =", target).KeysOnly()
	k, err := q.Run(c).Next(nil)
	if err != nil {
		c.Infof("Error getting Mention (%s => %s): %v", source, target, err)
		return nil
	}
	return k
}

func (e *Entry) Save(c appengine.Context) error {
	var k *datastore.Key
	if e.HasId() {
		id, _ := MaxId(c)
		e.Id = id + 1
		k = datastore.NewIncompleteKey(c, "Entry", nil)
	} else {
		k = datastore.NewKey(c, "Entry", fmt.Sprintf("%d", e.Id), 0, nil)
	}

	_, err := datastore.Put(c, k, e)
	if err == nil {
		c.Infof("Wrote %+v", e)
	} else {
		c.Warningf("Error writing entry: %v", e)
	}
	return err
}
