package grifts

import (
	"github.com/gofrs/uuid"
	"github.com/markbates/grift/grift"
)

var db = make(map[uuid.UUID]models.user)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		// Add DB seeding stuff here
		return nil
	})

})
