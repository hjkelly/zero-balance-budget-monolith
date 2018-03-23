package categories

import (
	"github.com/hjkelly/zbbapi/common"
	"github.com/hjkelly/zbbapi/models"
	uuid "github.com/satori/go.uuid"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func UpdateID(id string, input models.Category) (*models.Category, error) {
	ds := newDatastore()

	// Make sure the one we're updating exists.
	current := models.Category{}
	err := ds.C().Find(bson.M{
		"_id": uuid.FromStringOrNil(id),
	}).One(&current)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, common.NotFoundErr
		} else {
			return nil, err
		}
	}

	// Validate the input and use it to update the current data.
	err = validate(input)
	if err != nil {
		return nil, err
	}
	result := getUpdated(current, input)
	result.SetModificationTimestamp()

	// Update the database with our new result.
	err = ds.C().UpdateId(result.ID, result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
