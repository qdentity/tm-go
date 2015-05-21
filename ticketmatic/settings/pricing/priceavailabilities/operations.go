package priceavailabilities

import (
	"github.com/ticketmatic/tm-go/ticketmatic"
)

// List results
type List struct {
	// Result data
	Data []*ticketmatic.PriceAvailability `json:"data"`
}

// Get a list of price availabilities
func Getlist(client *ticketmatic.Client, params *ticketmatic.PriceAvailabilityQuery) (*List, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/pricing/priceavailabilities")
	r.AddParameter("includearchived", params.Includearchived)
	r.AddParameter("lastupdatesince", params.Lastupdatesince)
	r.AddParameter("filter", params.Filter)

	var obj *List
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Get a single price availability
func Get(client *ticketmatic.Client, id int64) (*ticketmatic.PriceAvailability, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/pricing/priceavailabilities/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.PriceAvailability
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create a new price availability
func Create(client *ticketmatic.Client, data *ticketmatic.PriceAvailability) (*ticketmatic.PriceAvailability, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/pricing/priceavailabilities")
	r.Body(data)

	var obj *ticketmatic.PriceAvailability
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Modify an existing price availability
func Update(client *ticketmatic.Client, id int64, data *ticketmatic.PriceAvailability) (*ticketmatic.PriceAvailability, error) {
	r := client.NewRequest("PUT", "/{accountname}/settings/pricing/priceavailabilities/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data)

	var obj *ticketmatic.PriceAvailability
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Remove a price availability
//
// Price availabilities are archivable: this call won't actually delete the object
// from the database. Instead, it will mark the object as archived, which means it
// won't show up anymore in most places.
//
// Most object types are archivable and can't be deleted: this is needed to ensure
// consistency of historical data.
func Delete(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("DELETE", "/{accountname}/settings/pricing/priceavailabilities/{id}")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}
