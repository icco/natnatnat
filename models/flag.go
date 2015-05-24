package models

import (
	"golang.org/x/net/context"

	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

type Flag struct {
	Name  string
	Value string
}

func SetFlag(c context.Context, flag string, value string) error {
	e := new(Flag)
	e.Name = flag
	e.Value = value

	k := datastore.NewKey(c, "Flag", flag, 0, nil)
	_, err := datastore.Put(c, k, e)
	if err == nil {
		log.Infof(c, "Wrote %+v", e)
	} else {
		log.Warningf(c, "Error writing entry: %v", e)
	}
	return err
}

func GetFlag(c context.Context, flag string) (string, error) {
	var retrieved Flag
	k := datastore.NewKey(c, "Flag", flag, 0, nil)
	err := datastore.Get(c, k, &retrieved)
	if err != nil {
		return "", err
	}

	return retrieved.Value, nil
}

func GetFlagLogError(c context.Context, flag string) string {
	ret, err := GetFlag(c, flag)
	if err != nil {
		log.Warningf(c, "Error getting flag '%s': %v", flag, err)
		return ""
	}

	return ret
}

func WriteVersionKey(c context.Context) error {
	return SetFlag(c, "VERSION", "1.0.1")
}
