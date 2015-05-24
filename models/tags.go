package models

import (
	"golang.org/x/net/context"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func AllTags(c context.Context) map[string]int {
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

func (a *Alias) Save(c context.Context) error {
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

func AllAliases(c context.Context) (*[]Alias, error) {
	return Aliases(c, -1)
}

func AliasMap(c context.Context) map[string]string {
	m := make(map[string]string, 0)
	aliases, err := AllAliases(c)
	if err != nil {
		c.Errorf("Error building Alias Map: %+v", err)
		return m
	}

	for _, v := range *aliases {
		m[v.Name] = v.Tag
	}

	return m
}

func Aliases(c context.Context, limit int) (*[]Alias, error) {
	q := datastore.NewQuery("Alias")

	if limit > 0 {
		q = q.Limit(limit)
	}

	aliases := new([]Alias)
	_, err := q.GetAll(c, aliases)
	return aliases, err
}

func GetAlias(c context.Context, tag string) (bool, string) {
	var retrieved Alias
	k := datastore.NewKey(c, "Alias", tag, 0, nil)
	err := datastore.Get(c, k, &retrieved)
	if err != nil {
		return false, ""
	}

	return true, retrieved.Tag
}

func GetTagAliases(c context.Context, tag string) *[]string {
	aliases := new([]Alias)
	ret := make([]string, 0)
	q := datastore.NewQuery("Alias").Filter("Tag =", tag)
	_, err := q.GetAll(c, aliases)
	if err != nil {
		return &ret
	}

	for _, v := range *aliases {
		ret = append(ret, v.Name)
	}

	return &ret
}
