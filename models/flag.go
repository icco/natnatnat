package models

import (
	"appengine"
	"appengine/datastore"
)

type Flag struct {
	Name  string
	Value string
}

func SetFlag(c appengine.Context, flag string, value string) error {
	e := new(Flag)
	e.Name = flag
	e.Value = value

	k := datastore.NewKey(c, "Flag", flag, 0, nil)
	_, err := datastore.Put(c, k, e)
	if err == nil {
		c.Infof("Wrote %+v", e)
	} else {
		c.Warningf("Error writing entry: %v", e)
	}
	return err
}

func GetFlag(c appengine.Context, flag string) (string, error) {
	var retrieved Flag
	k := datastore.NewKey(c, "Flag", flag, 0, nil)
	err := datastore.Get(c, k, &retrieved)
	if err != nil {
		return "", err
	}

	return retrieved.Value, nil
}
