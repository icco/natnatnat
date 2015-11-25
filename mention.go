package main

import (
	"net/url"
	"time"

	"golang.org/x/net/context"

	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

type Mention struct {
	Source   url.URL
	Target   url.URL
	Recieved time.Time
	Verified bool
}

func NewMention(c context.Context, source string, target string) (*Mention, error) {
	var k *datastore.Key
	var e *Mention

	k = GetMention(c, source, target)
	if k == nil {
		k = datastore.NewIncompleteKey(c, "Mention", nil)
		e = new(Mention)
		src, err := url.Parse(source)
		if err != nil {
			return nil, err
		}

		trg, err := url.Parse(target)
		if err != nil {
			return nil, err
		}

		e.Source = *src
		e.Target = *trg
		e.Recieved = time.Now()
		e.Verified = false

		k, err = datastore.Put(c, k, e)
		if err != nil {
			return nil, err
		}
	} else {
		if err := datastore.Get(c, k, &e); err != nil {
			return nil, err
		}
	}

	return e, nil
}

func MentionExists(c context.Context, source string, target string) bool {
	return GetMention(c, source, target) != nil
}

func GetMention(c context.Context, source string, target string) *datastore.Key {
	q := datastore.NewQuery("Mention").Filter("source =", source).Filter("target =", target).KeysOnly()
	k, err := q.Run(c).Next(nil)
	if err != nil {
		log.Infof(c, "Error getting Mention (%s => %s): %v", source, target, err)
		return nil
	}
	return k
}
