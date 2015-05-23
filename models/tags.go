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

type Alias struct {
	Name string `json:"alias"`
	Tag  string `json:"tag"`
}

func NewAlias(alias string, tag string) *Alias {
	a := new(Alias)

	a.Name = alias
	a.Tag = tag

	return a
}

func (a *Alias) Save(c appengine.Context) error {
	k := datastore.NewKey(c, "Alias", a.Name, 0, nil)
	k2, err := datastore.Put(c, k, a)
	if err == nil {
		c.Infof("Wrote %+v", a)
		c.Infof("Old key: %+v; New Key: %+v", k, k2)
	} else {
		c.Warningf("Error writing Alias: %v", a)
	}
	return err
}

func AllAliases(c appengine.Context) (*[]Alias, error) {
	return Aliases(c, -1)
}

func AliasMap(c appengine.Context) map[string]string {
	m := make(map[string]int, 0)
	aliases, err := AllAliases(c)
	if err != nil {
		c.Error(err.Error())
		return m
	}

	for k, v := range aliases {
		m[v.Name] = v.Tag
	}

	return m
}

func Aliases(c appengine.Context, limit int) (*[]Alias, error) {
	q := datastore.NewQuery("Alias")

	if limit > 0 {
		q = q.Limit(limit)
	}

	aliases := new([]Alias)
	_, err := q.GetAll(c, aliases)
	return aliases, err
}
