package pluralize_test

import (
	"testing"

	"github.com/hgsigner/stringfy/pluralize"
	"github.com/stretchr/testify/assert"
)

func Test_Plural(t *testing.T) {
	a := assert.New(t)

	tests := []struct {
		count     int
		singular  string
		addPlural bool
		plural    string
		out       string
	}{
		{
			count:    1,
			singular: "ball",
			out:      "1 ball",
		},
		{
			count:     2,
			singular:  "ball",
			addPlural: true,
			plural:    "ballys",
			out:       "2 ballys",
		},
		{
			count:    2,
			singular: "boat",
			out:      "2 boats",
		},
		{
			count:    2,
			singular: "axis",
			out:      "2 axes",
		},
		{
			count:    2,
			singular: "Testis",
			out:      "2 testes",
		},
		{
			count:    2,
			singular: "octopus",
			out:      "2 octopi",
		},
		{
			count:    2,
			singular: "virus",
			out:      "2 viri",
		},
		{
			count:    2,
			singular: "viri",
			out:      "2 viri",
		},
		{
			count:    2,
			singular: "octopi",
			out:      "2 octopi",
		},
		{
			count:    2,
			singular: "bus",
			out:      "2 buses",
		},
		{
			count:    2,
			singular: "buffalo",
			out:      "2 buffaloes",
		},
		{
			count:    2,
			singular: "tomato",
			out:      "2 tomatoes",
		},
		{
			count:    2,
			singular: "zoocytium",
			out:      "2 zoocytia",
		},
		{
			count:    2,
			singular: "zoocytia",
			out:      "2 zoocytia",
		},
		{
			count:    2,
			singular: "analysis",
			out:      "2 analyses",
		},
		{
			count:    2,
			singular: "wolf",
			out:      "2 wolves",
		},
		{
			count:    2,
			singular: "thief",
			out:      "2 thieves",
		},
		{
			count:    2,
			singular: "hive",
			out:      "2 hives",
		},
		{
			count:    2,
			singular: "quality",
			out:      "2 qualities",
		},
		{
			count:    2,
			singular: "fox",
			out:      "2 foxes",
		},
		{
			count:    2,
			singular: "dash",
			out:      "2 dashes",
		},
		{
			count:    2,
			singular: "pass",
			out:      "2 passes",
		},
		{
			count:    2,
			singular: "vouch",
			out:      "2 vouches",
		},
		{
			count:    2,
			singular: "quiz",
			out:      "2 quizes",
		},
	}

	for _, t := range tests {
		p := pluralize.New()
		if t.addPlural {
			p.Options(pluralize.AddPlural(t.plural))
		}
		pf := p.Perform(t.count, t.singular)
		a.Equal(t.out, pf)
	}
}
