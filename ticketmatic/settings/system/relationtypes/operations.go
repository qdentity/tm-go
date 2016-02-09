package relationtypes

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// List results
type List struct {
	// Result data
	Data []*ticketmatic.RelationType `json:"data"`
}

// Get a list of relation types
func Getlist(client *ticketmatic.Client, params *ticketmatic.RelationTypeQuery) (*List, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/system/relationtypes")
	if params != nil {
		r.AddParameter("includearchived", params.Includearchived)
		r.AddParameter("lastupdatesince", params.Lastupdatesince)
		r.AddParameter("filter", params.Filter)
	}

	var obj *List
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Get a single relation type
func Get(client *ticketmatic.Client, id int64) (*ticketmatic.RelationType, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/system/relationtypes/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.RelationType
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create a new relation type
func Create(client *ticketmatic.Client, data *ticketmatic.RelationType) (*ticketmatic.RelationType, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/system/relationtypes")
	r.Body(data)

	var obj *ticketmatic.RelationType
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Modify an existing relation type
func Update(client *ticketmatic.Client, id int64, data *ticketmatic.RelationType) (*ticketmatic.RelationType, error) {
	r := client.NewRequest("PUT", "/{accountname}/settings/system/relationtypes/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj *ticketmatic.RelationType
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Remove a relation type
//
// Relation types are archivable: this call won't actually delete the object from
// the database. Instead, it will mark the object as archived, which means it won't
// show up anymore in most places.
//
// Most object types are archivable and can't be deleted: this is needed to ensure
// consistency of historical data.
func Delete(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("DELETE", "/{accountname}/settings/system/relationtypes/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}
