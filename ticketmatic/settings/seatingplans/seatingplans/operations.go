package seatingplans

import (
	"bytes"

	"github.com/ticketmatic/tm-go/ticketmatic"
)

// List results
type List struct {
	// Result data
	Data []*ticketmatic.SeatingPlan `json:"data"`

	// The total number of results that are available without considering limit and offset, useful for paging.
	NbrOfResults int `json:"nbrofresults"`
}

// Get a list of seating plans
func Getlist(client *ticketmatic.Client, params *ticketmatic.SeatingPlanQuery) (*List, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/seatingplans/seatingplans", "json")
	if params != nil {
		r.AddParameter("filter", params.Filter)
		r.AddParameter("includearchived", params.Includearchived)
		r.AddParameter("lastupdatesince", params.Lastupdatesince)
	}

	var obj *List
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Get a single seating plan
func Get(client *ticketmatic.Client, id int64) (*ticketmatic.SeatingPlan, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/seatingplans/seatingplans/{id}", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj *ticketmatic.SeatingPlan
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Create a new seating plan
func Create(client *ticketmatic.Client, data *ticketmatic.SeatingPlan) (*ticketmatic.SeatingPlan, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/seatingplans/seatingplans", "json")
	r.Body(data, "json")

	var obj *ticketmatic.SeatingPlan
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Modify an existing seating plan
func Update(client *ticketmatic.Client, id int64, data *ticketmatic.SeatingPlan) (*ticketmatic.SeatingPlan, error) {
	r := client.NewRequest("PUT", "/{accountname}/settings/seatingplans/seatingplans/{id}", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data, "json")

	var obj *ticketmatic.SeatingPlan
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Remove a seating plan
//
// Seating plans are archivable: this call won't actually delete the object from
// the database. Instead, it will mark the object as archived, which means it won't
// show up anymore in most places.
//
// Most object types are archivable and can't be deleted: this is needed to ensure
// consistency of historical data.
func Delete(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("DELETE", "/{accountname}/settings/seatingplans/seatingplans/{id}", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}

// Get the svg for a specific zone
//
// Returns the svg for this specific zone.
func Getsvg(client *ticketmatic.Client, id int64, zoneid string) (*bytes.Buffer, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/seatingplans/seatingplans/{id}/svg/{zoneid}", "svg")
	r.UrlParameters(map[string]interface{}{
		"id":     id,
		"zoneid": zoneid,
	})

	var obj bytes.Buffer
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// Update the svg for a specific zone
//
// Updates the svg for this specific zone.
func Savesvg(client *ticketmatic.Client, id int64, zoneid string, data string) (*bytes.Buffer, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/seatingplans/seatingplans/{id}/svg/{zoneid}", "svg")
	r.UrlParameters(map[string]interface{}{
		"id":     id,
		"zoneid": zoneid,
	})
	r.Body(data, "svg")

	var obj bytes.Buffer
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// Get the lock templates
//
// Returns the lock templates for this seating plan.
func Getlocktemplates(client *ticketmatic.Client, id int64) ([]*ticketmatic.LockTemplate, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/seatingplans/seatingplans/{id}/locktemplates", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj []*ticketmatic.LockTemplate
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Save the lock templates
//
// Save the lock templates for this seating plan.
func Savelocktemplates(client *ticketmatic.Client, id int64, data []*ticketmatic.LockTemplate) ([]*ticketmatic.LockTemplate, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/seatingplans/seatingplans/{id}/locktemplates", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data, "json")

	var obj []*ticketmatic.LockTemplate
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Get the seat description templates
//
// Returns the seat description templates for this seating plan.
func Getseatdescriptiontemplates(client *ticketmatic.Client, id int64) ([]*ticketmatic.SeatDescriptionTemplate, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/seatingplans/seatingplans/{id}/seatdescriptiontemplates", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	var obj []*ticketmatic.SeatDescriptionTemplate
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Save the seat description templates
//
// Save the seat description templates for this seating plan.
func Saveseatdescriptiontemplates(client *ticketmatic.Client, id int64, data []*ticketmatic.SeatDescriptionTemplate) ([]*ticketmatic.SeatDescriptionTemplate, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/seatingplans/seatingplans/{id}/seatdescriptiontemplates", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})
	r.Body(data, "json")

	var obj []*ticketmatic.SeatDescriptionTemplate
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Get the logical plan for a specific zone
//
// Returns the logical plan for this specific zone.
func Getlogicalplan(client *ticketmatic.Client, id int64, zoneid string) (*ticketmatic.LogicalPlan, error) {
	r := client.NewRequest("GET", "/{accountname}/settings/seatingplans/seatingplans/{id}/logicalplan/{zoneid}", "json")
	r.UrlParameters(map[string]interface{}{
		"id":     id,
		"zoneid": zoneid,
	})

	var obj *ticketmatic.LogicalPlan
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Update the logical plan for this zone.
//
// Updates the logical plan for this zone.
func Savelogicalplan(client *ticketmatic.Client, id int64, zoneid string, data *ticketmatic.LogicalPlan) (*ticketmatic.LogicalPlan, error) {
	r := client.NewRequest("POST", "/{accountname}/settings/seatingplans/seatingplans/{id}/logicalplan/{zoneid}", "json")
	r.UrlParameters(map[string]interface{}{
		"id":     id,
		"zoneid": zoneid,
	})
	r.Body(data, "json")

	var obj *ticketmatic.LogicalPlan
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Purge seatingplan
func Purge(client *ticketmatic.Client, id int64) error {
	r := client.NewRequest("PUT", "/{accountname}/settings/seatingplans/seatingplans/{id}/purge", "json")
	r.UrlParameters(map[string]interface{}{
		"id": id,
	})

	return r.Run(nil)
}
