package ticketmatic

import (
	"encoding/json"
	"strings"
)

// An account parameter defines general behavior of your account
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/AccountParameter).
type AccountParameter struct {
	// The name of the account parameter
	Key string `json:"key"`

	// Value
	Value interface{} `json:"value,omitempty"`
}

// Result when adding tickets
// (https://www.ticketmatic.com/docs/api/orders/addtickets) or products
// (https://www.ticketmatic.com/docs/api/orders/addproducts) to an order
// (https://www.ticketmatic.com/docs/api/types/Order).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/AddItemsResult).
type AddItemsResult struct {
	// Ids of the items that were added
	Ids []int64 `json:"ids"`

	// The modified order
	Order *Order `json:"order,omitempty"`
}

// Request data used to add a payment
// (https://www.ticketmatic.com/docs/api/orders/addpayments) to an order
// (https://www.ticketmatic.com/docs/api/types/Order).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/AddPayments).
type AddPayments struct {
	// Amount for the payment
	Amount float64 `json:"amount"`

	// Id of the payment method to be used for the payment
	Paymentmethodid int64 `json:"paymentmethodid"`
}

// Request data used to add products
// (https://www.ticketmatic.com/docs/api/orders/addproducts) to an order
// (https://www.ticketmatic.com/docs/api/types/Order).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/AddProducts).
type AddProducts struct {
	// Product information
	Products []*CreateProduct `json:"products"`
}

// Request data used to refund a payment
// (https://www.ticketmatic.com/docs/api/orders/addrefunds) for an order
// (https://www.ticketmatic.com/docs/api/types/Order).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/AddRefunds).
type AddRefunds struct {
	// Amount that needs to be refunded
	Amount float64 `json:"amount"`

	// Id of the payment that needs to be refunded
	Paymentid int64 `json:"paymentid"`
}

// Request data used to add tickets
// (https://www.ticketmatic.com/docs/api/orders/addtickets) to an order
// (https://www.ticketmatic.com/docs/api/types/Order). The amount of tickets that
// can be added is limited to 50 per call.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/AddTickets).
type AddTickets struct {
	// Ticket information
	Tickets []*CreateTicket `json:"tickets"`
}

// Parameters used to create voucher codes
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/AddVoucherCodes).
type AddVoucherCodes struct {
	// Value of the voucher
	Amount float64 `json:"amount,omitempty"`

	// List of voucher codes, can also (optionally) contain expiry timestamps.
	Codes []*VoucherCode `json:"codes"`

	// Number of codes to create
	Count int64 `json:"count"`
}

// Address, used for physical deliveries and contact details.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/Address).
type Address struct {
	// Address ID
	//
	// Note: Only available when used for a contact
	// (https://www.ticketmatic.com/docs/api/types/Contact) address.
	Id int64 `json:"id"`

	// Address type ID
	//
	// Note: Only available when used for a contact
	// (https://www.ticketmatic.com/docs/api/types/Contact) address.
	Typeid int64 `json:"typeid"`

	// Addressee
	//
	// Note: Only available when used as an order
	// (https://www.ticketmatic.com/docs/api/types/Order) deliver address.
	Addressee string `json:"addressee"`

	// City
	City string `json:"city"`

	// Country name (based on typeid, returned as a convenience).
	//
	// Note: Only available when used for a contact
	// (https://www.ticketmatic.com/docs/api/types/Contact) address.
	Country string `json:"country"`

	// ISO 3166-1 alpha-2 country code
	// (http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
	Countrycode string `json:"countrycode"`

	// Contact this address belongs to
	//
	// Note: Only available when used for a contact
	// (https://www.ticketmatic.com/docs/api/types/Contact) address.
	Customerid int64 `json:"customerid"`

	// State
	State string `json:"state"`

	// Street field 1 (typically the street name)
	Street1 string `json:"street1"`

	// Street field 2 (typically the number)
	Street2 string `json:"street2"`

	// Street field 3 (sometimes used for box numbers or suffixes)
	Street3 string `json:"street3"`

	// Address type (based on typeid, returned as a convenience).
	//
	// Note: Only available when used for a contact
	// (https://www.ticketmatic.com/docs/api/types/Contact) address.
	Type string `json:"type"`

	// Zip code
	Zip string `json:"zip"`
}

// Batch operations performed on contacts
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/BatchContactOperation).
type BatchContactOperation struct {
	// Apply operation to all contacts except for the supplied IDs
	Excludeids []int64 `json:"excludeids"`

	// Restrict operation to supplied IDs
	Ids []int64 `json:"ids"`

	// Operation to perform
	Operation string `json:"operation"`

	// Operation-specific parameters
	Parameters *BatchContactParameters `json:"parameters,omitempty"`
}

// Parameters for batch operations performed on contacts
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/BatchContactParameters).
type BatchContactParameters struct {
	// Selection name
	Name string `json:"name"`

	// Relation type IDs
	Ids []int64 `json:"ids"`
}

// A single contact.
//
// More info: see the get operation
// (https://www.ticketmatic.com/docs/api/contacts/get) and the contacts endpoint
// (https://www.ticketmatic.com/docs/api/contacts).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/Contact).
type Contact struct {
	// Contact ID
	//
	// Note: Ignored when creating a new contact.
	//
	// Note: Ignored when updating an existing contact.
	Id int64 `json:"id"`

	// Account type.
	//
	// Indicates the authentication type supported for this contact (used when
	// authentication is enabled in web sales).
	//
	// Possible values:
	//
	// * 0: No authentication
	//
	// * 1901: Password authentication
	//
	// * 1902: Facebook
	//
	// * 1903: Google
	//
	// * 1904: Twitter
	//
	// Note: Ignored when creating a new contact.
	//
	// Note: Ignored when updating an existing contact.
	AccountType int64 `json:"account_type,omitempty"`

	// Addresses
	Addresses []*Address `json:"addresses"`

	// Birth date
	Birthdate Time `json:"birthdate"`

	// Company
	Company string `json:"company"`

	// Customer title ID (also determines the gender of the contact)
	Customertitleid int64 `json:"customertitleid"`

	// E-mail address
	Email string `json:"email"`

	// First name
	Firstname string `json:"firstname"`

	// Language (ISO 639-1 code (http://en.wikipedia.org/wiki/List_of_ISO_639-1_codes))
	Languagecode string `json:"languagecode"`

	// Last name
	Lastname string `json:"lastname"`

	// Middle name
	Middlename string `json:"middlename"`

	// Job function
	Organizationfunction string `json:"organizationfunction"`

	// Phone numbers
	Phonenumbers []*Phonenumber `json:"phonenumbers"`

	// Relation type IDs
	Relationtypes []int64 `json:"relationtypes"`

	// Sex
	Sex string `json:"sex"`

	// Contact status
	//
	// Possible values:
	//
	// * deleted: Contact has been deleted
	//
	// * incomplete: Contact misses crucial account information
	//
	// * (blank): Normal contact
	//
	// Note: Ignored when creating a new contact.
	//
	// Note: Ignored when updating an existing contact.
	Status string `json:"status"`

	// Whether or not this contact is subscribed in the e-mail marketing integration
	//
	// Note: Ignored when creating a new contact.
	//
	// Note: Ignored when updating an existing contact.
	Subscribed bool `json:"subscribed"`

	// VAT Number (for organizations)
	Vatnumber string `json:"vatnumber"`

	// Whether or not this contact has been deleted
	//
	// Note: Ignored when creating a new contact.
	//
	// Note: Ignored when updating an existing contact.
	Isdeleted bool `json:"isdeleted"`

	// Created timestamp
	//
	// Note: Ignored when creating a new contact.
	//
	// Note: Ignored when updating an existing contact.
	Createdts Time `json:"createdts"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new contact.
	//
	// Note: Ignored when updating an existing contact.
	Lastupdatets Time `json:"lastupdatets"`

	// Custom fields
	CustomFields map[string]interface{} `json:"-"`
}

// Custom unmarshaller with support for custom fields
func (o *Contact) UnmarshalJSON(data []byte) error {
	// Alias the type, to avoid calling UnmarshalJSON. Unpack it.
	type tmp Contact
	var obj tmp
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return err
	}

	*o = Contact(obj)

	// Unpack it again, this time to a map, so we can pull out the custom fields.
	var raw map[string]interface{}
	err = json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	o.CustomFields = make(map[string]interface{})
	for key, val := range raw {
		if strings.HasPrefix(key, "c_") {
			o.CustomFields[key[2:]] = val
		}
	}

	// Note: We're doing a double JSON decode here, I'd love to get rid of it
	// but I'm not sure how we can do this easily. Suggestions welcome:
	// developers@ticketmatic.com!

	return nil
}

// Custom marshaller with support for custom fields
func (o *Contact) MarshalJSON() ([]byte, error) {
	// Use a custom type to avoid the custom marshaller, marshal the data.
	type tmp struct {
		Id                   int64          `json:"id,omitempty"`
		AccountType          int64          `json:"account_type,omitempty"`
		Addresses            []*Address     `json:"addresses,omitempty"`
		Birthdate            Time           `json:"birthdate,omitempty"`
		Company              string         `json:"company,omitempty"`
		Customertitleid      int64          `json:"customertitleid,omitempty"`
		Email                string         `json:"email,omitempty"`
		Firstname            string         `json:"firstname,omitempty"`
		Languagecode         string         `json:"languagecode,omitempty"`
		Lastname             string         `json:"lastname,omitempty"`
		Middlename           string         `json:"middlename,omitempty"`
		Organizationfunction string         `json:"organizationfunction,omitempty"`
		Phonenumbers         []*Phonenumber `json:"phonenumbers,omitempty"`
		Relationtypes        []int64        `json:"relationtypes,omitempty"`
		Sex                  string         `json:"sex,omitempty"`
		Status               string         `json:"status,omitempty"`
		Subscribed           bool           `json:"subscribed,omitempty"`
		Vatnumber            string         `json:"vatnumber,omitempty"`
		Isdeleted            bool           `json:"isdeleted,omitempty"`
		Createdts            Time           `json:"createdts,omitempty"`
		Lastupdatets         Time           `json:"lastupdatets,omitempty"`
	}

	obj := tmp{
		Id:                   o.Id,
		AccountType:          o.AccountType,
		Addresses:            o.Addresses,
		Birthdate:            o.Birthdate,
		Company:              o.Company,
		Customertitleid:      o.Customertitleid,
		Email:                o.Email,
		Firstname:            o.Firstname,
		Languagecode:         o.Languagecode,
		Lastname:             o.Lastname,
		Middlename:           o.Middlename,
		Organizationfunction: o.Organizationfunction,
		Phonenumbers:         o.Phonenumbers,
		Relationtypes:        o.Relationtypes,
		Sex:                  o.Sex,
		Status:               o.Status,
		Subscribed:           o.Subscribed,
		Vatnumber:            o.Vatnumber,
		Isdeleted:            o.Isdeleted,
		Createdts:            o.Createdts,
		Lastupdatets:         o.Lastupdatets,
	}
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	// Unpack it again, to get the wire representation
	var raw map[string]interface{}
	err = json.Unmarshal(data, &raw)
	if err != nil {
		return nil, err
	}

	for key, val := range o.CustomFields {
		raw["c_"+key] = val
	}

	// Pack it again
	return json.Marshal(raw)

	// Note: Like UnmarshalJSON, this is quite crazy. But it works beautifully.
	// Know a way to do this better? Get in touch!
}

// A single contact address type.
//
// More info: see the get operation
// (https://www.ticketmatic.com/docs/api/settings/system/contactaddresstypes/get)
// and the contact address types endpoint
// (https://www.ticketmatic.com/docs/api/settings/system/contactaddresstypes).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/ContactAddressType).
type ContactAddressType struct {
	// Unique ID
	//
	// Note: Ignored when creating a new contact address type.
	//
	// Note: Ignored when updating an existing contact address type.
	Id int64 `json:"id"`

	// Name of the address type
	Name string `json:"name"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new contact address type.
	//
	// Note: Ignored when updating an existing contact address type.
	Isarchived bool `json:"isarchived"`

	// Created timestamp
	//
	// Note: Ignored when creating a new contact address type.
	//
	// Note: Ignored when updating an existing contact address type.
	Createdts Time `json:"createdts"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new contact address type.
	//
	// Note: Ignored when updating an existing contact address type.
	Lastupdatets Time `json:"lastupdatets"`
}

// Set of parameters used to filter contact address types.
//
// More info: see contact address type
// (https://www.ticketmatic.com/docs/api/types/ContactAddressType), the getlist
// operation
// (https://www.ticketmatic.com/docs/api/settings/system/contactaddresstypes/getlist)
// and the contact address types endpoint
// (https://www.ticketmatic.com/docs/api/settings/system/contactaddresstypes).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/ContactAddressTypeQuery).
type ContactAddressTypeQuery struct {
	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`

	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`
}

// Contact field is a list of field that are asked upon registration
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/ContactField).
type ContactField struct {
	// Id of this ContactField
	Id int64 `json:"id"`

	// Name of this contactfield
	Name string `json:"name"`

	// Caption of this contactfield
	Caption string `json:"caption"`
}

// Optional alternative methods to retrieve a contact
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/ContactGetQuery).
type ContactGetQuery struct {
	// Contact e-mail address
	Email string `json:"email,omitempty"`
}

// Filter parameters to fetch a list of contacts
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/ContactQuery).
type ContactQuery struct {
	// A SQL query that returns contact IDs
	//
	// Can be used to do arbitrary filtering. See the database documentation for
	// contact (https://www.ticketmatic.com/docs/db/contact) for more information.
	Filter string `json:"filter,omitempty"`

	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// Only include contacts that have been updated since the given timestamp.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Limit results to at most the given amount of contacts.
	Limit int64 `json:"limit,omitempty"`

	// Skip the first X contacts.
	Offset int64 `json:"offset,omitempty"`

	// Order by the given field.
	//
	// Supported values: name, lastupdatets, createdts.
	Orderby string `json:"orderby,omitempty"`

	// Output format.
	//
	// Possible values:
	//
	// * ids: Only fill the ID field
	//
	// * minimal: A minimal set of order fields
	//
	// * default: Return all order fields (also used when the output parameter is
	// omitted)
	Output string `json:"output,omitempty"`

	// A text filter string.
	//
	// Matches against the contact name and contact details.
	Searchterm string `json:"searchterm,omitempty"`
}

// A single contact title.
//
// More info: see the get operation
// (https://www.ticketmatic.com/docs/api/settings/system/contacttitles/get) and the
// contact titles endpoint
// (https://www.ticketmatic.com/docs/api/settings/system/contacttitles).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/ContactTitle).
type ContactTitle struct {
	// Unique ID
	//
	// Note: Ignored when creating a new contact title.
	//
	// Note: Ignored when updating an existing contact title.
	Id int64 `json:"id"`

	// Title name
	Name string `json:"name"`

	// Restricts this title from showing up on the websales pages
	Isinternal bool `json:"isinternal"`

	// Language for this title
	Languagecode string `json:"languagecode"`

	// Gender associated with this title
	Sex string `json:"sex"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new contact title.
	//
	// Note: Ignored when updating an existing contact title.
	Isarchived bool `json:"isarchived"`

	// Created timestamp
	//
	// Note: Ignored when creating a new contact title.
	//
	// Note: Ignored when updating an existing contact title.
	Createdts Time `json:"createdts"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new contact title.
	//
	// Note: Ignored when updating an existing contact title.
	Lastupdatets Time `json:"lastupdatets"`
}

// Set of parameters used to filter contact titles.
//
// More info: see contact title
// (https://www.ticketmatic.com/docs/api/types/ContactTitle), the getlist operation
// (https://www.ticketmatic.com/docs/api/settings/system/contacttitles/getlist) and
// the contact titles endpoint
// (https://www.ticketmatic.com/docs/api/settings/system/contacttitles).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/ContactTitleQuery).
type ContactTitleQuery struct {
	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`

	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`
}

// Required data for creating an order
// (https://www.ticketmatic.com/docs/api/types/Order).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/CreateOrder).
type CreateOrder struct {
	// Event IDs that might end up in this order, used to hint the rate limiter
	// (https://www.ticketmatic.com/docs/api/ratelimiting) of what might come.
	Events []int64 `json:"events"`

	// Product IDs that might end up in this order, used to hint the rate limiter
	// (https://www.ticketmatic.com/docs/api/ratelimiting) of what might come.
	Products []int64 `json:"products"`

	// Sales channel in which this order is created
	Saleschannelid int64 `json:"saleschannelid"`
}

// Info for adding a product
// (https://www.ticketmatic.com/docs/api/orders/addproducts) to an order
// (https://www.ticketmatic.com/docs/api/types/Order).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/CreateProduct).
type CreateProduct struct {
	// The id for the product you want to add.
	Productid int64 `json:"productid"`

	// The property values for the product.
	Properties map[string]string `json:"properties,omitempty"`
}

// Info for adding a ticket
// (https://www.ticketmatic.com/docs/api/orders/addtickets) to an order
// (https://www.ticketmatic.com/docs/api/types/Order).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/CreateTicket).
type CreateTicket struct {
	// The id for the optionbundle you want to add a new ticket to. Either
	// tickettypepriceid or optionbundleid should be specified, not both.
	Optionbundleid int64 `json:"optionbundleid,omitempty"`

	// Manually select a specific ticket.
	Ticketid int64 `json:"ticketid,omitempty"`

	// Should only be specified when optionbundleid is specified. The tickettypeid for
	// the ticket you want to add to the optionbundle.
	Tickettypeid int64 `json:"tickettypeid,omitempty"`

	// The ticket type price ID for the new ticket. Either tickettypepriceid or
	// optionbundleid should be specified, not both.
	Tickettypepriceid int64 `json:"tickettypepriceid,omitempty"`

	// Voucher code to use (if any)
	Vouchercode string `json:"vouchercode,omitempty"`
}

// A single custom field.
//
// More info: see the get operation
// (https://www.ticketmatic.com/docs/api/settings/system/customfields/get) and the
// custom fields endpoint
// (https://www.ticketmatic.com/docs/api/settings/system/customfields).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/CustomField).
type CustomField struct {
	// Unique ID
	//
	// Note: Ignored when creating a new custom field.
	//
	// Note: Ignored when updating an existing custom field.
	Id int64 `json:"id"`

	// Type ID
	//
	// Note: Ignored when updating an existing custom field.
	Typeid int64 `json:"typeid"`

	// Rules that define in what conditions this custom field is available when edit
	// type is checkout
	//
	// Note: Not set when retrieving a list of custom fields.
	Availability *CustomfieldAvailability `json:"availability,omitempty"`

	// Human-readable name for the custom field
	Caption string `json:"caption"`

	// Human-readable description for the custom field. Will be visible for end-users
	// when edittype checkout is used
	//
	// Note: Not set when retrieving a list of custom fields.
	Description string `json:"description"`

	// Type of editing that is allowed for the custom field. Links to systemtype
	// category 22xxx
	Edittypeid int64 `json:"edittypeid"`

	// Type of the custom field. Links to systemtype category 12xxx
	Fieldtypeid int64 `json:"fieldtypeid"`

	// The identifier for the custom field. Should contain only alphanumeric characters
	// and no whitespace, max length is 20 characters. The custom field will be
	// available in the api and the public data model as c_
	Key string `json:"key"`

	// Indicates whether the field is required
	Required bool `json:"required"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new custom field.
	//
	// Note: Ignored when updating an existing custom field.
	Isarchived bool `json:"isarchived"`

	// Created timestamp
	//
	// Note: Ignored when creating a new custom field.
	//
	// Note: Ignored when updating an existing custom field.
	Createdts Time `json:"createdts"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new custom field.
	//
	// Note: Ignored when updating an existing custom field.
	Lastupdatets Time `json:"lastupdatets"`
}

// Set of parameters used to filter custom fields.
//
// More info: see custom field
// (https://www.ticketmatic.com/docs/api/types/CustomField), the getlist operation
// (https://www.ticketmatic.com/docs/api/settings/system/customfields/getlist) and
// the custom fields endpoint
// (https://www.ticketmatic.com/docs/api/settings/system/customfields).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/CustomFieldQuery).
type CustomFieldQuery struct {
	// Only return items with the given typeid.
	Typeid int64 `json:"typeid"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`

	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`
}

// A single custom field value.
//
// More info: see the get operation
// (https://www.ticketmatic.com/docs/api/settings/system/customfieldvalues/get) and
// the custom field values endpoint
// (https://www.ticketmatic.com/docs/api/settings/system/customfieldvalues).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/CustomFieldValue).
type CustomFieldValue struct {
	// Unique ID
	//
	// Note: Ignored when creating a new custom field value.
	//
	// Note: Ignored when updating an existing custom field value.
	Id int64 `json:"id"`

	// Type ID
	//
	// Note: Ignored when updating an existing custom field value.
	Typeid int64 `json:"typeid"`

	// Human-readable name for the value
	Caption string `json:"caption"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new custom field value.
	//
	// Note: Ignored when updating an existing custom field value.
	Isarchived bool `json:"isarchived"`

	// Created timestamp
	//
	// Note: Ignored when creating a new custom field value.
	//
	// Note: Ignored when updating an existing custom field value.
	Createdts Time `json:"createdts"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new custom field value.
	//
	// Note: Ignored when updating an existing custom field value.
	Lastupdatets Time `json:"lastupdatets"`
}

// Set of parameters used to filter custom field values.
//
// More info: see custom field value
// (https://www.ticketmatic.com/docs/api/types/CustomFieldValue), the getlist
// operation
// (https://www.ticketmatic.com/docs/api/settings/system/customfieldvalues/getlist)
// and the custom field values endpoint
// (https://www.ticketmatic.com/docs/api/settings/system/customfieldvalues).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/CustomFieldValueQuery).
type CustomFieldValueQuery struct {
	// Only return items with the given typeid.
	Typeid int64 `json:"typeid,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`

	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`
}

// A CustomfieldAvailability configures in what saleschannels a custom field is
// available (during the checkout).
//
// It can also further refine the availability based on a script written in
// JavaScript.
//
// More information about writing order scripts can be found here
// (https://www.ticketmatic.com/docs/tickets/configure_ticket_sales/orderfees).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/CustomfieldAvailability).
type CustomfieldAvailability struct {
	// The custom field will be available for these saleschannels. It this is empty the
	// custom field will not be available.
	Saleschannels []int64 `json:"saleschannels"`

	// A Javascript that needs to return a boolean. It has the current order available.
	Script string `json:"script"`

	// Indicates if the script will be used.
	Usescript bool `json:"usescript"`
}

// Request data used to delete products
// (https://www.ticketmatic.com/docs/api/orders/deleteproducts) from an order
// (https://www.ticketmatic.com/docs/api/types/Order).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/DeleteProducts).
type DeleteProducts struct {
	// Product IDs
	Products []int64 `json:"products"`
}

// Request data used to delete tickets
// (https://www.ticketmatic.com/docs/api/orders/deletetickets) from an order
// (https://www.ticketmatic.com/docs/api/types/Order).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/DeleteTickets).
type DeleteTickets struct {
	// Ticket IDs
	Tickets []int64 `json:"tickets"`
}

// A single delivery scenario.
//
// More info: see the get operation
// (https://www.ticketmatic.com/docs/api/settings/ticketsales/deliveryscenarios/get)
// and the delivery scenarios endpoint
// (https://www.ticketmatic.com/docs/api/settings/ticketsales/deliveryscenarios).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/DeliveryScenario).
type DeliveryScenario struct {
	// Unique ID
	//
	// Note: Ignored when creating a new delivery scenario.
	//
	// Note: Ignored when updating an existing delivery scenario.
	Id int64 `json:"id"`

	// The type of this delivery scenario, defines when this delivery scenario is
	// triggered. The available values for this field can be found on the delivery
	// scenario overview
	// (https://www.ticketmatic.com/docs/api/settings/ticketsales/deliveryscenarios)
	// page.
	Typeid int64 `json:"typeid"`

	// Name of the delivery scenario
	Name string `json:"name"`

	// Are e-tickets allowed with this delivery scenario?
	Allowetickets int64 `json:"allowetickets"`

	// The rules that define when this scenario is available. See the delivery scenario
	// overview
	// (https://www.ticketmatic.com/docs/api/settings/ticketsales/deliveryscenarios)
	// page for a description of this field
	//
	// Note: Not set when retrieving a list of delivery scenarios.
	Availability *DeliveryscenarioAvailability `json:"availability,omitempty"`

	// An internal description field. Will not be shown to customers.
	Internalremark string `json:"internalremark,omitempty"`

	// A physical address is required
	Needsaddress bool `json:"needsaddress"`

	// The ID of the order mail template that will be used for sending out this
	// delivery scenario. Can be 0 to indicate that no mail should be sent
	OrdermailtemplateidDelivery int64 `json:"ordermailtemplateid_delivery"`

	// A short description of the deilvery scenario. Will be shown to customers.
	Shortdescription string `json:"shortdescription"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new delivery scenario.
	//
	// Note: Ignored when updating an existing delivery scenario.
	Isarchived bool `json:"isarchived"`

	// Created timestamp
	//
	// Note: Ignored when creating a new delivery scenario.
	//
	// Note: Ignored when updating an existing delivery scenario.
	Createdts Time `json:"createdts"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new delivery scenario.
	//
	// Note: Ignored when updating an existing delivery scenario.
	Lastupdatets Time `json:"lastupdatets"`
}

// Set of parameters used to filter delivery scenarios.
//
// More info: see delivery scenario
// (https://www.ticketmatic.com/docs/api/types/DeliveryScenario), the getlist
// operation
// (https://www.ticketmatic.com/docs/api/settings/ticketsales/deliveryscenarios/getlist)
// and the delivery scenarios endpoint
// (https://www.ticketmatic.com/docs/api/settings/ticketsales/deliveryscenarios).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/DeliveryScenarioQuery).
type DeliveryScenarioQuery struct {
	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`

	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`
}

// A DeliveryscenarioAvailability defines when a delivery scenario
// (https://www.ticketmatic.com/docs/api/types/DeliveryScenario) is available.
//
// This can be done in two ways:
//
// * By specifying a set of sales channels (required)
//
// * By writing a script (optional)
//
// In its simplest form, a DeliveryscenarioAvailability looks like this:
//
//
//    {
//        "saleschannels": [1, 2]
//    }
//
//
//
// This defines that the delivery scenario is available when used in the context of
// saleschannel 1 or 2.
//
// More complex logic can be specified by writing a small piece of JavaScript
// (http://en.wikipedia.org/wiki/JavaScript). To do so, you need to add a usescript
// and script field to the availability:
//
//
//    {
//        "saleschannels": [1, 2],
//        "usescript": true,
//        "script": "// script here"
//    }
//
//
//
// Note that the list of sales channel IDs is still required: the script can only
// restrict this set further.
//
// A simple example of a delivery scenario script:
//
//
//    return order.tickets.length < 3 && saleschannel.typeid == 3002;
//
//
//
// This script states that the current delivery scenario is only available if the
// amount of tickets in the order is less than 3 and the current sales channel is a
// web sales channel.
//
// With this script the DeliveryscenarioAvailability would look like this:
//
//
//    {
//        "saleschannels": [1, 2],
//        "usescript": true,
//        "script": "return order.tickets.length < 3 && saleschannel.typeid == 3002;"
//    }
//
//
//
// The following variables are available in the script:
//
// * order
//
// * saleschannel
//
// You can use any valid JavaScript syntax (including conditionals and loops). Note
// that each script has a strict time limit.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/DeliveryscenarioAvailability).
type DeliveryscenarioAvailability struct {
	// An array of sales channel IDs for which this delivery scenario can be used
	Saleschannels []int64 `json:"saleschannels"`

	// Script used to determine availability of the delivery scenario
	Script string `json:"script"`

	// Use a script to refine the set of sales channels?
	Usescript bool `json:"usescript"`
}

// A single document.
//
// More info: see the get operation
// (https://www.ticketmatic.com/docs/api/settings/communicationanddesign/documents/get)
// and the documents endpoint
// (https://www.ticketmatic.com/docs/api/settings/communicationanddesign/documents).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/Document).
type Document struct {
	// Unique ID
	//
	// Note: Ignored when creating a new document.
	//
	// Note: Ignored when updating an existing document.
	Id int64 `json:"id"`

	// Type ID
	//
	// Note: Ignored when updating an existing document.
	Typeid int64 `json:"typeid"`

	// Name of the document
	Name string `json:"name"`

	// Css content for the document template
	//
	// Note: Not set when retrieving a list of documents.
	Css string `json:"css"`

	// Description of the document
	Description string `json:"description"`

	// HTML content for the document template
	//
	// Note: Not set when retrieving a list of documents.
	Htmltemplate string `json:"htmltemplate"`

	// Translations for the document template
	//
	// Note: Not set when retrieving a list of documents.
	Translations map[string]string `json:"translations,omitempty"`

	// Created timestamp
	//
	// Note: Ignored when creating a new document.
	//
	// Note: Ignored when updating an existing document.
	Createdts Time `json:"createdts"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new document.
	//
	// Note: Ignored when updating an existing document.
	Lastupdatets Time `json:"lastupdatets"`
}

// Set of parameters used to filter documents.
//
// More info: see document (https://www.ticketmatic.com/docs/api/types/Document),
// the getlist operation
// (https://www.ticketmatic.com/docs/api/settings/communicationanddesign/documents/getlist)
// and the documents endpoint
// (https://www.ticketmatic.com/docs/api/settings/communicationanddesign/documents).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/DocumentQuery).
type DocumentQuery struct {
	// Only return items with the given typeid.
	Typeid int64 `json:"typeid,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`
}

// A single Event.
//
// Status
//
// The currentstatus field of an event can have any of the following values:
//
// * Draft (19001)
//
// * Active (19002)
//
// * Closed (19003)
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/Event).
type Event struct {
	// Event ID
	//
	// Note: Ignored when creating a new event.
	//
	// Note: Ignored when updating an existing event.
	Id int64 `json:"id,omitempty"`

	// Event name
	Name string `json:"name"`

	// Information on the availability of tickets per contingent. Read-only.
	//
	// Note: Ignored when creating a new event.
	//
	// Note: Ignored when updating an existing event.
	Availability []*EventContingentAvailability `json:"availability"`

	// Event code.
	//
	// Used as a unique identifier in web sales.
	//
	// Note: Ignored when creating a new event.
	//
	// Note: Ignored when updating an existing event.
	Code string `json:"code,omitempty"`

	// Information about the contingents in the Event that are not in the seatingplan
	Contingents []*EventContingent `json:"contingents"`

	// Event status
	//
	// The available values for this field can be found on the Event
	// (https://www.ticketmatic.com/docs/api/types/Event) page.
	Currentstatus int64 `json:"currentstatus,omitempty"`

	// Event end time
	Endts Time `json:"endts"`

	// External event code.
	//
	// This field is typically set when importing data from other systems.
	Externalcode string `json:"externalcode,omitempty"`

	// Event location ID
	//
	// See event locations
	// (https://www.ticketmatic.com/docs/api/settings/events/eventlocations) for more
	// information.
	Locationid int64 `json:"locationid,omitempty"`

	// Event location name
	//
	// Automatically derived using locationid.
	//
	// Note: Ignored when creating a new event.
	//
	// Note: Ignored when updating an existing event.
	Locationname string `json:"locationname,omitempty"`

	// Maximum number of tickets for this event that can be added to a basket
	Maxnbrofticketsperbasket int64 `json:"maxnbrofticketsperbasket,omitempty"`

	// Information on the available prices for the event
	//
	// Note: Ignored when creating a new event.
	//
	// Note: Ignored when updating an existing event.
	Prices *EventPrices `json:"prices,omitempty"`

	// Production ID
	Productionid int64 `json:"productionid,omitempty"`

	// Event publish time
	//
	// Note: Ignored when creating a new event.
	//
	// Note: Ignored when updating an existing event.
	Publishedts Time `json:"publishedts,omitempty"`

	// Queue ID
	//
	// See rate limiting (https://www.ticketmatic.com/docs/api/ratelimiting) for more
	// information.
	Queuetoken int64 `json:"queuetoken,omitempty"`

	// DEPRECATED
	Revenuesplitid int64 `json:"revenuesplitid,omitempty"`

	// Time of end of sales.
	//
	// Used for all sales channels for which no specific sales period has been defined.
	Saleendts Time `json:"saleendts,omitempty"`

	// Per-sales channel information about when this event is for sale.
	Saleschannels []*EventSalesChannel `json:"saleschannels"`

	// Time of start of sales.
	//
	// Used for all sales channels for which no specific sales period has been defined.
	Salestartts Time `json:"salestartts,omitempty"`

	// Information about the contingents defined in the seatingplan. Read-only.
	//
	// Note: Ignored when creating a new event.
	//
	// Note: Ignored when updating an existing event.
	Seatingplancontingents []*EventSeatingplanContingent `json:"seatingplancontingents"`

	// Event specific prices in addition to the prices defined in the field
	// seatingplanpricelistid. Prices from the pricelist and the event specific prices
	// are combined in one pricelist for the event. The optional position attribute
	// defines where the event specific prices will be positioned in the resulting
	// pricelist
	Seatingplaneventspecificprices *PricelistPrices `json:"seatingplaneventspecificprices,omitempty"`

	// Seating plan ID
	//
	// Only set for events with fixed seats.
	Seatingplanid int64 `json:"seatingplanid,omitempty"`

	// Name of the seatingplanlocktemplate to apply linking a seatingplanid to this
	// event. This is not a numeric id but the name of the lock template as specified
	// in the seatingplan's logicalplan.
	//
	// Note: Ignored when creating a new event.
	//
	// Note: Ignored when updating an existing event.
	Seatingplanlocktemplate string `json:"seatingplanlocktemplate,omitempty"`

	// Price list ID for fixed seats.
	//
	// Only set for events with fixed seats. See price lists
	// (https://www.ticketmatic.com/docs/api/settings/pricing/pricelists) for more
	// information.
	Seatingplanpricelistid int64 `json:"seatingplanpricelistid,omitempty"`

	// Event start time
	Startts Time `json:"startts"`

	// Event subtitle
	Subtitle string `json:"subtitle,omitempty"`

	// Event subtitle (2)
	Subtitle2 string `json:"subtitle2,omitempty"`

	// Ticket fee ID
	//
	// Determines which ticket fee rules are used for this event. See ticket fees
	// (https://www.ticketmatic.com/docs/api/settings/pricing/ticketfees) for more
	// information.
	Ticketfeeid int64 `json:"ticketfeeid,omitempty"`

	// Ticket layout ID
	//
	// See ticket layouts
	// (https://www.ticketmatic.com/docs/api/settings/communicationanddesign/ticketlayouts)
	// for more information.
	Ticketlayoutid int64 `json:"ticketlayoutid,omitempty"`

	// Small description that will be shown on the sales pages of this event
	Webremark string `json:"webremark,omitempty"`

	// Created timestamp
	//
	// Note: Ignored when creating a new event.
	//
	// Note: Ignored when updating an existing event.
	Createdts Time `json:"createdts"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new event.
	//
	// Note: Ignored when updating an existing event.
	Lastupdatets Time `json:"lastupdatets"`

	// Custom fields
	CustomFields map[string]interface{} `json:"-"`
}

// Custom unmarshaller with support for custom fields
func (o *Event) UnmarshalJSON(data []byte) error {
	// Alias the type, to avoid calling UnmarshalJSON. Unpack it.
	type tmp Event
	var obj tmp
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return err
	}

	*o = Event(obj)

	// Unpack it again, this time to a map, so we can pull out the custom fields.
	var raw map[string]interface{}
	err = json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	o.CustomFields = make(map[string]interface{})
	for key, val := range raw {
		if strings.HasPrefix(key, "c_") {
			o.CustomFields[key[2:]] = val
		}
	}

	// Note: We're doing a double JSON decode here, I'd love to get rid of it
	// but I'm not sure how we can do this easily. Suggestions welcome:
	// developers@ticketmatic.com!

	return nil
}

// Custom marshaller with support for custom fields
func (o *Event) MarshalJSON() ([]byte, error) {
	// Use a custom type to avoid the custom marshaller, marshal the data.
	type tmp struct {
		Id                             int64                          `json:"id,omitempty"`
		Name                           string                         `json:"name,omitempty"`
		Availability                   []*EventContingentAvailability `json:"availability,omitempty"`
		Code                           string                         `json:"code,omitempty"`
		Contingents                    []*EventContingent             `json:"contingents,omitempty"`
		Currentstatus                  int64                          `json:"currentstatus,omitempty"`
		Endts                          Time                           `json:"endts,omitempty"`
		Externalcode                   string                         `json:"externalcode,omitempty"`
		Locationid                     int64                          `json:"locationid,omitempty"`
		Locationname                   string                         `json:"locationname,omitempty"`
		Maxnbrofticketsperbasket       int64                          `json:"maxnbrofticketsperbasket,omitempty"`
		Prices                         *EventPrices                   `json:"prices,omitempty"`
		Productionid                   int64                          `json:"productionid,omitempty"`
		Publishedts                    Time                           `json:"publishedts,omitempty"`
		Queuetoken                     int64                          `json:"queuetoken,omitempty"`
		Revenuesplitid                 int64                          `json:"revenuesplitid,omitempty"`
		Saleendts                      Time                           `json:"saleendts,omitempty"`
		Saleschannels                  []*EventSalesChannel           `json:"saleschannels,omitempty"`
		Salestartts                    Time                           `json:"salestartts,omitempty"`
		Seatingplancontingents         []*EventSeatingplanContingent  `json:"seatingplancontingents,omitempty"`
		Seatingplaneventspecificprices *PricelistPrices               `json:"seatingplaneventspecificprices,omitempty"`
		Seatingplanid                  int64                          `json:"seatingplanid,omitempty"`
		Seatingplanlocktemplate        string                         `json:"seatingplanlocktemplate,omitempty"`
		Seatingplanpricelistid         int64                          `json:"seatingplanpricelistid,omitempty"`
		Startts                        Time                           `json:"startts,omitempty"`
		Subtitle                       string                         `json:"subtitle,omitempty"`
		Subtitle2                      string                         `json:"subtitle2,omitempty"`
		Ticketfeeid                    int64                          `json:"ticketfeeid,omitempty"`
		Ticketlayoutid                 int64                          `json:"ticketlayoutid,omitempty"`
		Webremark                      string                         `json:"webremark,omitempty"`
		Createdts                      Time                           `json:"createdts,omitempty"`
		Lastupdatets                   Time                           `json:"lastupdatets,omitempty"`
	}

	obj := tmp{
		Id:                       o.Id,
		Name:                     o.Name,
		Availability:             o.Availability,
		Code:                     o.Code,
		Contingents:              o.Contingents,
		Currentstatus:            o.Currentstatus,
		Endts:                    o.Endts,
		Externalcode:             o.Externalcode,
		Locationid:               o.Locationid,
		Locationname:             o.Locationname,
		Maxnbrofticketsperbasket: o.Maxnbrofticketsperbasket,
		Prices:                         o.Prices,
		Productionid:                   o.Productionid,
		Publishedts:                    o.Publishedts,
		Queuetoken:                     o.Queuetoken,
		Revenuesplitid:                 o.Revenuesplitid,
		Saleendts:                      o.Saleendts,
		Saleschannels:                  o.Saleschannels,
		Salestartts:                    o.Salestartts,
		Seatingplancontingents:         o.Seatingplancontingents,
		Seatingplaneventspecificprices: o.Seatingplaneventspecificprices,
		Seatingplanid:                  o.Seatingplanid,
		Seatingplanlocktemplate:        o.Seatingplanlocktemplate,
		Seatingplanpricelistid:         o.Seatingplanpricelistid,
		Startts:                        o.Startts,
		Subtitle:                       o.Subtitle,
		Subtitle2:                      o.Subtitle2,
		Ticketfeeid:                    o.Ticketfeeid,
		Ticketlayoutid:                 o.Ticketlayoutid,
		Webremark:                      o.Webremark,
		Createdts:                      o.Createdts,
		Lastupdatets:                   o.Lastupdatets,
	}
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	// Unpack it again, to get the wire representation
	var raw map[string]interface{}
	err = json.Unmarshal(data, &raw)
	if err != nil {
		return nil, err
	}

	for key, val := range o.CustomFields {
		raw["c_"+key] = val
	}

	// Pack it again
	return json.Marshal(raw)

	// Note: Like UnmarshalJSON, this is quite crazy. But it works beautifully.
	// Know a way to do this better? Get in touch!
}

// Used when requesting events, to restrict the event information to a specific
// context.
//
// Currently allows you to filter the event information (both the events and the
// pricing information within each event) to a specific saleschannel. This makes it
// very easy to show the correct information on a website.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/EventContext).
type EventContext struct {
	// The ID of the saleschannel used to restrict the event information
	Saleschannelid int64 `json:"saleschannelid,omitempty"`
}

// Information about a contingent for an event
// (https://www.ticketmatic.com/docs/api/types/Event).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/EventContingent).
type EventContingent struct {
	// Contingent ID
	Id int64 `json:"id,omitempty"`

	// Name of the contingent
	Name string `json:"name"`

	// Number of tickets in the contingent
	Amount int64 `json:"amount"`

	// Event ID
	Eventid int64 `json:"eventid,omitempty"`

	// Event specific prices in addition to the prices defined in the field
	// pricelistid. Prices from the pricelist and the event specific prices are
	// combined in one pricelist. The optional position attribute defines where the
	// event specific prices will be positioned in the resulting pricelist
	Eventspecificprices *PricelistPrices `json:"eventspecificprices,omitempty"`

	// Locked tickets for the contingent
	Locks []*EventContingentLock `json:"locks"`

	// Price list ID for this contingent
	Pricelistid int64 `json:"pricelistid,omitempty"`

	// Whether the barcodes for the tickets in this contingent were imported (true), or
	// were generated internally (false)
	Withimportedbarcodes bool `json:"withimportedbarcodes,omitempty"`
}

// Information about the availability of tickets for a contingent
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/EventContingentAvailability).
type EventContingentAvailability struct {
	// Number of complimentary tickets
	Complimentary int64 `json:"complimentary"`

	// Number of available tickets
	Free int64 `json:"free"`

	// Number of locked tickets with a hard lock type
	LockedHard int64 `json:"locked_hard"`

	// Number of locked tickets with a soft lock type
	LockedSoft int64 `json:"locked_soft"`

	// Number of tickets reserved in unconfirmed orders
	Reserved int64 `json:"reserved"`

	// Number of tickets in confirmed orders that are completely paid
	SoldPaid int64 `json:"sold_paid"`

	// Number of tickets in confirmed orders that are not completely paid
	SoldUnpaid int64 `json:"sold_unpaid"`

	// Contingent ID
	Tickettypeid int64 `json:"tickettypeid"`

	// Total number of tickets in the contingent
	Total int64 `json:"total"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets"`
}

// Information about locked tickets in a Contingent.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/EventContingentLock).
type EventContingentLock struct {
	// Number of tickets in the contingent
	Amount int64 `json:"amount"`

	// Lock type ID
	Locktypeid int64 `json:"locktypeid"`

	// Contingent ID
	Tickettypeid int64 `json:"tickettypeid"`
}

// Used when requesting events, to filter events.
//
// Currently allows you to filter based on the production ID.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/EventFilter).
type EventFilter struct {
	// The ID of the production
	Productionid int64 `json:"productionid,omitempty"`

	// The event status. By default, events with status Active or Closed will be
	// returned
	Status []int64 `json:"status"`
}

// A single event location.
//
// More info: see the get operation
// (https://www.ticketmatic.com/docs/api/settings/events/eventlocations/get) and
// the event locations endpoint
// (https://www.ticketmatic.com/docs/api/settings/events/eventlocations).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/EventLocation).
type EventLocation struct {
	// Unique ID
	//
	// Note: Ignored when creating a new event location.
	//
	// Note: Ignored when updating an existing event location.
	Id int64 `json:"id"`

	// Name of the location
	Name string `json:"name"`

	// City
	City string `json:"city"`

	// Country code. Should be an ISO 3166-1 alpha-2
	// (http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) two-letter code.
	Countrycode string `json:"countrycode"`

	// State
	State string `json:"state"`

	// Street name
	Street1 string `json:"street1"`

	// Nr. + Box
	Street2 string `json:"street2"`

	// Zipcode
	Zip string `json:"zip"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new event location.
	//
	// Note: Ignored when updating an existing event location.
	Isarchived bool `json:"isarchived"`

	// Created timestamp
	//
	// Note: Ignored when creating a new event location.
	//
	// Note: Ignored when updating an existing event location.
	Createdts Time `json:"createdts"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new event location.
	//
	// Note: Ignored when updating an existing event location.
	Lastupdatets Time `json:"lastupdatets"`
}

// Set of parameters used to filter event locations.
//
// More info: see event location
// (https://www.ticketmatic.com/docs/api/types/EventLocation), the getlist
// operation
// (https://www.ticketmatic.com/docs/api/settings/events/eventlocations/getlist)
// and the event locations endpoint
// (https://www.ticketmatic.com/docs/api/settings/events/eventlocations).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/EventLocationQuery).
type EventLocationQuery struct {
	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`

	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`
}

// Used when locking a set of tickets. Contains the locktypeid and the set of
// ticketids
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/EventLockTickets).
type EventLockTickets struct {
	// Id of the locktype to use for the lock.
	Locktypeid int64 `json:"locktypeid"`

	// Array of ticketids to lock.
	Ticketids []int64 `json:"ticketids"`
}

// Information about the prices for an event.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/EventPrices).
type EventPrices struct {
	// Price information for the contingents
	Contingents []*EventPricesContingent `json:"contingents"`
}

// Information about the prices for a contingent for an event.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/EventPricesContingent).
type EventPricesContingent struct {
	// Contingent ID
	Contingentid int64 `json:"contingentid"`

	// Price information for the pricetypes
	Pricetypes []*EventPricesPricetype `json:"pricetypes"`
}

// Information about costs for a price for an event.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/EventPricesCost).
type EventPricesCost struct {
	// The actual cost
	Cost float64 `json:"cost"`

	// Cost ID
	Costid int64 `json:"costid"`
}

// Information about the price for a pricetype for the specific sales channel for
// an event.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/EventPricesPricetype).
type EventPricesPricetype struct {
	// Pricetype ID
	Pricetypeid int64 `json:"pricetypeid"`

	// Price information for this pricetype for the different sales channels
	Saleschannels []*EventPricesSaleschannel `json:"saleschannels"`

	// Ticket type price ID, used to add tickets to an order
	Tickettypepriceid int64 `json:"tickettypepriceid"`
}

// Information about the price for a pricetype for the specific sales channel for
// an event.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/EventPricesSaleschannel).
type EventPricesSaleschannel struct {
	// The costs associated with this price
	Costs []*EventPricesCost `json:"costs"`

	// The actual price
	Price float64 `json:"price"`

	// Saleschannel ID
	Saleschannelid int64 `json:"saleschannelid"`

	// The actual servicecharge
	Servicecharge float64 `json:"servicecharge"`

	// Tickettypeprice ID
	Tickettypepriceid int64 `json:"tickettypepriceid"`
}

// Filter parameters to fetch a list of events
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/EventQuery).
type EventQuery struct {
	// Restrict the event information to a specific context.
	//
	// Currently allows you to filter the event information (both the events and the
	// pricing information within each event) to a specific saleschannel. This makes it
	// very easy to show the correct information on a website.
	Context *EventContext `json:"context,omitempty"`

	// A SQL query that returns event IDs
	//
	// Can be used to do arbitrary filtering. See the database documentation for event
	// (https://www.ticketmatic.com/docs/db/event) for more information.
	Filter string `json:"filter,omitempty"`

	// Only include events that have been updated since the given timestamp.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Limit results to at most the given amount of events.
	Limit int64 `json:"limit,omitempty"`

	// Skip the first X events.
	Offset int64 `json:"offset,omitempty"`

	// Order by the given field.
	//
	// Supported values: name, startts.
	Orderby string `json:"orderby,omitempty"`

	// Output format.
	//
	// Possible values:
	//
	// * ids: Only fill the ID field
	//
	// * default: Return all event fields (also used when the output parameter is
	// omitted)
	//
	// * withlookup: Returns all event fields and an additional lookup field which
	// contains all dependent objects
	Output string `json:"output,omitempty"`

	// A text filter string.
	//
	// Matches against the start of the event name, the production name or the
	// subtitle.
	Searchterm string `json:"searchterm,omitempty"`

	// Filters the events based on a given set of fields. Currently supports:
	// productionid, status and pricetypeids.
	Simplefilter *EventFilter `json:"simplefilter,omitempty"`
}

// Information about the sales period for a specific sales channel
// (https://www.ticketmatic.com/docs/api/types/SalesChannel) in an event
// (https://www.ticketmatic.com/docs/api/types/Event).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/EventSalesChannel).
type EventSalesChannel struct {
	// Event ID
	Eventid int64 `json:"eventid"`

	// Whether or not this sales channel is active for this event
	Isactive bool `json:"isactive"`

	// When the sales end
	Saleendts Time `json:"saleendts"`

	// Sales channel ID
	Saleschannelid int64 `json:"saleschannelid"`

	// When the sales start
	Salestartts Time `json:"salestartts"`
}

// Information about a contingent in the seating plan for an event
// (https://www.ticketmatic.com/docs/api/types/Event).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/EventSeatingplanContingent).
type EventSeatingplanContingent struct {
	// Contingent ID
	Id int64 `json:"id"`

	// Name of the contingent
	Name string `json:"name"`

	// Number of tickets in the contingent
	Amount int64 `json:"amount"`

	// Event ID
	Eventid int64 `json:"eventid"`

	// Seat rank ID
	Seatrankid int64 `json:"seatrankid"`
}

// A single ticket.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/EventTicket).
type EventTicket struct {
	// Ticket ID
	Id int64 `json:"id"`

	// Link to the order the ticket is contained in
	//
	// Note: Ignored in the result for updating tickets
	//
	// Note: Ignored when updating tickets
	Orderid int64 `json:"orderid,omitempty"`

	// The id of the scanner used for the last succesful entry
	//
	// Note: Ignored in the result for updating tickets
	//
	// Note: Ignored when updating tickets
	Accesscontrollastenteredscandeviceid int64 `json:"accesscontrollastenteredscandeviceid,omitempty"`

	// The timestamp of the last succesful entry with this ticket
	//
	// Note: Ignored in the result for updating tickets
	//
	// Note: Ignored when updating tickets
	Accesscontrollastenteredts Time `json:"accesscontrollastenteredts,omitempty"`

	// The id of the scanner used for the last succesful exit
	//
	// Note: Ignored in the result for updating tickets
	//
	// Note: Ignored when updating tickets
	Accesscontrollastexitscandeviceid int64 `json:"accesscontrollastexitscandeviceid,omitempty"`

	// The timestamp of the last succesful exit with this ticket
	//
	// Note: Ignored in the result for updating tickets
	//
	// Note: Ignored when updating tickets
	Accesscontrollastexitts Time `json:"accesscontrollastexitts,omitempty"`

	// The access control status for this ticket. 0 means not scanned, 1 means
	// succesful entry, 2 means succesful exit
	//
	// Note: Ignored in the result for updating tickets
	//
	// Note: Ignored when updating tickets
	Accesscontrolstatus int64 `json:"accesscontrolstatus,omitempty"`

	// Ticket barcode
	Barcode string `json:"barcode,omitempty"`

	// Link to the bundle (orderproduct) that this ticket belongs to.
	//
	// Note: Ignored in the result for updating tickets
	//
	// Note: Ignored when updating tickets
	Bundleid int64 `json:"bundleid,omitempty"`

	// Link to the locktype used for locking the ticket
	//
	// Note: Ignored in the result for updating tickets
	//
	// Note: Ignored when updating tickets
	Locktypeid int64 `json:"locktypeid,omitempty"`

	// Fee for the ticket in the order
	//
	// Note: Ignored in the result for updating tickets
	//
	// Note: Ignored when updating tickets
	Orderfee float64 `json:"orderfee,omitempty"`

	// Price for the ticket in the order (without fee)
	//
	// Note: Ignored in the result for updating tickets
	//
	// Note: Ignored when updating tickets
	Price float64 `json:"price,omitempty"`

	// The seat description for this ticket (only for seated tickets)
	//
	// Note: Ignored in the result for updating tickets
	//
	// Note: Ignored when updating tickets
	Seatdescription string `json:"seatdescription,omitempty"`

	// Seat ID (for seated tickets)
	//
	// Note: Ignored in the result for updating tickets
	//
	// Note: Ignored when updating tickets
	Seatid string `json:"seatid,omitempty"`

	// Number indicating the priority for this ticket for the best available algorithm.
	// Tickets with a higher priority will be considered first when performing a best
	// available allocation.
	//
	// Note: Ignored in the result for updating tickets
	//
	// Note: Ignored when updating tickets
	Seatpriority int64 `json:"seatpriority,omitempty"`

	// Row number for the ticket (only for seated tickets)
	//
	// Note: Ignored in the result for updating tickets
	//
	// Note: Ignored when updating tickets
	Seatrownumber string `json:"seatrownumber,omitempty"`

	// Seat number for the ticket (only for seated tickets)
	//
	// Note: Ignored in the result for updating tickets
	//
	// Note: Ignored when updating tickets
	Seatseatnumber string `json:"seatseatnumber,omitempty"`

	// Optional seatzone for the ticket
	//
	// Note: Ignored in the result for updating tickets
	//
	// Note: Ignored when updating tickets
	Seatzoneid int64 `json:"seatzoneid,omitempty"`

	// Zone name for the ticket (only for seated tickets)
	//
	// Note: Ignored in the result for updating tickets
	//
	// Note: Ignored when updating tickets
	Seatzonename string `json:"seatzonename,omitempty"`

	// Optional link to the contact that is the ticketholder for this ticket
	//
	// Note: Ignored in the result for updating tickets
	//
	// Note: Ignored when updating tickets
	Ticketholderid int64 `json:"ticketholderid,omitempty"`

	// Optional name on the ticket
	//
	// Note: Ignored in the result for updating tickets
	//
	// Note: Ignored when updating tickets
	Ticketname string `json:"ticketname,omitempty"`

	// Link to the contingent this ticket belongs to
	//
	// Note: Ignored in the result for updating tickets
	//
	// Note: Ignored when updating tickets
	Tickettypeid int64 `json:"tickettypeid,omitempty"`

	// Link to the tickettypeprice that is assigned to this ticket for the order.
	// Through the tickettypeprice you can retrieve the pricetype
	//
	// Note: Ignored in the result for updating tickets
	//
	// Note: Ignored when updating tickets
	Tickettypepriceid int64 `json:"tickettypepriceid,omitempty"`

	// Link to the vouchercode that was used to reserve this ticket.
	//
	// Note: Ignored in the result for updating tickets
	//
	// Note: Ignored when updating tickets
	Vouchercodeid int64 `json:"vouchercodeid,omitempty"`

	// Custom fields
	CustomFields map[string]interface{} `json:"-"`
}

// Custom unmarshaller with support for custom fields
func (o *EventTicket) UnmarshalJSON(data []byte) error {
	// Alias the type, to avoid calling UnmarshalJSON. Unpack it.
	type tmp EventTicket
	var obj tmp
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return err
	}

	*o = EventTicket(obj)

	// Unpack it again, this time to a map, so we can pull out the custom fields.
	var raw map[string]interface{}
	err = json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	o.CustomFields = make(map[string]interface{})
	for key, val := range raw {
		if strings.HasPrefix(key, "c_") {
			o.CustomFields[key[2:]] = val
		}
	}

	// Note: We're doing a double JSON decode here, I'd love to get rid of it
	// but I'm not sure how we can do this easily. Suggestions welcome:
	// developers@ticketmatic.com!

	return nil
}

// Custom marshaller with support for custom fields
func (o *EventTicket) MarshalJSON() ([]byte, error) {
	// Use a custom type to avoid the custom marshaller, marshal the data.
	type tmp struct {
		Id                                   int64   `json:"id,omitempty"`
		Orderid                              int64   `json:"orderid,omitempty"`
		Accesscontrollastenteredscandeviceid int64   `json:"accesscontrollastenteredscandeviceid,omitempty"`
		Accesscontrollastenteredts           Time    `json:"accesscontrollastenteredts,omitempty"`
		Accesscontrollastexitscandeviceid    int64   `json:"accesscontrollastexitscandeviceid,omitempty"`
		Accesscontrollastexitts              Time    `json:"accesscontrollastexitts,omitempty"`
		Accesscontrolstatus                  int64   `json:"accesscontrolstatus,omitempty"`
		Barcode                              string  `json:"barcode,omitempty"`
		Bundleid                             int64   `json:"bundleid,omitempty"`
		Locktypeid                           int64   `json:"locktypeid,omitempty"`
		Orderfee                             float64 `json:"orderfee,omitempty"`
		Price                                float64 `json:"price,omitempty"`
		Seatdescription                      string  `json:"seatdescription,omitempty"`
		Seatid                               string  `json:"seatid,omitempty"`
		Seatpriority                         int64   `json:"seatpriority,omitempty"`
		Seatrownumber                        string  `json:"seatrownumber,omitempty"`
		Seatseatnumber                       string  `json:"seatseatnumber,omitempty"`
		Seatzoneid                           int64   `json:"seatzoneid,omitempty"`
		Seatzonename                         string  `json:"seatzonename,omitempty"`
		Ticketholderid                       int64   `json:"ticketholderid,omitempty"`
		Ticketname                           string  `json:"ticketname,omitempty"`
		Tickettypeid                         int64   `json:"tickettypeid,omitempty"`
		Tickettypepriceid                    int64   `json:"tickettypepriceid,omitempty"`
		Vouchercodeid                        int64   `json:"vouchercodeid,omitempty"`
	}

	obj := tmp{
		Id:      o.Id,
		Orderid: o.Orderid,
		Accesscontrollastenteredscandeviceid: o.Accesscontrollastenteredscandeviceid,
		Accesscontrollastenteredts:           o.Accesscontrollastenteredts,
		Accesscontrollastexitscandeviceid:    o.Accesscontrollastexitscandeviceid,
		Accesscontrollastexitts:              o.Accesscontrollastexitts,
		Accesscontrolstatus:                  o.Accesscontrolstatus,
		Barcode:                              o.Barcode,
		Bundleid:                             o.Bundleid,
		Locktypeid:                           o.Locktypeid,
		Orderfee:                             o.Orderfee,
		Price:                                o.Price,
		Seatdescription:                      o.Seatdescription,
		Seatid:                               o.Seatid,
		Seatpriority:                         o.Seatpriority,
		Seatrownumber:                        o.Seatrownumber,
		Seatseatnumber:                       o.Seatseatnumber,
		Seatzoneid:                           o.Seatzoneid,
		Seatzonename:                         o.Seatzonename,
		Ticketholderid:                       o.Ticketholderid,
		Ticketname:                           o.Ticketname,
		Tickettypeid:                         o.Tickettypeid,
		Tickettypepriceid:                    o.Tickettypepriceid,
		Vouchercodeid:                        o.Vouchercodeid,
	}
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	// Unpack it again, to get the wire representation
	var raw map[string]interface{}
	err = json.Unmarshal(data, &raw)
	if err != nil {
		return nil, err
	}

	for key, val := range o.CustomFields {
		raw["c_"+key] = val
	}

	// Pack it again
	return json.Marshal(raw)

	// Note: Like UnmarshalJSON, this is quite crazy. But it works beautifully.
	// Know a way to do this better? Get in touch!
}

// Used when requesting tickets for an event, to filter the tickets.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/EventTicketFilter).
type EventTicketFilter struct {
	// The ID of the tickettype (contingent)
	Tickettypeid int64 `json:"tickettypeid,omitempty"`
}

// Filter parameters to fetch a list of tickets for an event
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/EventTicketQuery).
type EventTicketQuery struct {
	// Filters the tickets based on a given set of fields.
	Simplefilter *EventTicketFilter `json:"simplefilter,omitempty"`
}

// Used when unlocking a set of tickets.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/EventUnlockTickets).
type EventUnlockTickets struct {
	// Array of ticketids to unlock.
	Ticketids []int64 `json:"ticketids"`
}

// Used when updating the seat rank for a set of tickets.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/EventUpdateSeatRankForTickets).
type EventUpdateSeatRankForTickets struct {
	// The seat rank
	Seatrankid int64 `json:"seatrankid"`

	// Array of ticketids to unlock.
	Ticketids []int64 `json:"ticketids"`
}

// A single field definition.
//
// More info: see the get operation
// (https://www.ticketmatic.com/docs/api/settings/system/fielddefinitions/get) and
// the field definitions endpoint
// (https://www.ticketmatic.com/docs/api/settings/system/fielddefinitions).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/FieldDefinition).
type FieldDefinition struct {
	// Unique ID
	//
	// Note: Ignored when creating a new field definition.
	//
	// Note: Ignored when updating an existing field definition.
	Id int64 `json:"id"`

	// Type ID
	//
	// Note: Ignored when updating an existing field definition.
	Typeid int64 `json:"typeid"`

	// Alignment of the field definition, when used in a view. Values can be 'left',
	// 'right' or 'center'
	Align string `json:"align"`

	// Human-readable name for the field definition
	Description string `json:"description"`

	// Key for the field definition. Should only consist of lowercase alphanumeric
	// characters
	Key string `json:"key"`

	// The actual definition of the field definition. Contains the sql clause that will
	// retrieve the information element in the database.
	Sqlclause string `json:"sqlclause"`

	// Will decide how the field will be rendered when used in a view.
	Uitype string `json:"uitype"`

	// Indicates whether the width for the field definition can be adapted when
	// stretching a view that includes the field definition across the whole available
	// width.
	Variablewidth bool `json:"variablewidth"`

	// Width of the field definition, when used in a view
	Width int64 `json:"width"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new field definition.
	//
	// Note: Ignored when updating an existing field definition.
	Isarchived bool `json:"isarchived"`

	// Created timestamp
	//
	// Note: Ignored when creating a new field definition.
	//
	// Note: Ignored when updating an existing field definition.
	Createdts Time `json:"createdts"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new field definition.
	//
	// Note: Ignored when updating an existing field definition.
	Lastupdatets Time `json:"lastupdatets"`
}

// Set of parameters used to filter field definitions.
//
// More info: see field definition
// (https://www.ticketmatic.com/docs/api/types/FieldDefinition), the getlist
// operation
// (https://www.ticketmatic.com/docs/api/settings/system/fielddefinitions/getlist)
// and the field definitions endpoint
// (https://www.ticketmatic.com/docs/api/settings/system/fielddefinitions).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/FieldDefinitionQuery).
type FieldDefinitionQuery struct {
	// Only return items with the given typeid.
	Typeid int64 `json:"typeid,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`

	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`
}

// A single filter definition.
//
// More info: see the get operation
// (https://www.ticketmatic.com/docs/api/settings/system/filterdefinitions/get) and
// the filter definitions endpoint
// (https://www.ticketmatic.com/docs/api/settings/system/filterdefinitions).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/FilterDefinition).
type FilterDefinition struct {
	// Unique ID
	//
	// Note: Ignored when creating a new filter definition.
	//
	// Note: Ignored when updating an existing filter definition.
	Id int64 `json:"id"`

	// Type ID
	//
	// Note: Ignored when updating an existing filter definition.
	Typeid int64 `json:"typeid"`

	// For certain filter types, the user must select a value from a list. The
	// checklistquery contains the sql clause to retrieve the list of available values.
	Checklistquery string `json:"checklistquery"`

	// Name for the filter
	Description string `json:"description"`

	// The type of filter definition defines the UI and resulting parameters that will
	// be used when a user selects the filter. The possible values can be found here
	// (https://www.ticketmatic.com/docs/api/settings/system/filterdefinitions).
	Filtertype int64 `json:"filtertype"`

	// The sql clause that defines how the filter will work
	Sqlclause string `json:"sqlclause"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new filter definition.
	//
	// Note: Ignored when updating an existing filter definition.
	Isarchived bool `json:"isarchived"`

	// Created timestamp
	//
	// Note: Ignored when creating a new filter definition.
	//
	// Note: Ignored when updating an existing filter definition.
	Createdts Time `json:"createdts"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new filter definition.
	//
	// Note: Ignored when updating an existing filter definition.
	Lastupdatets Time `json:"lastupdatets"`
}

// Set of parameters used to filter filter definitions.
//
// More info: see filter definition
// (https://www.ticketmatic.com/docs/api/types/FilterDefinition), the getlist
// operation
// (https://www.ticketmatic.com/docs/api/settings/system/filterdefinitions/getlist)
// and the filter definitions endpoint
// (https://www.ticketmatic.com/docs/api/settings/system/filterdefinitions).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/FilterDefinitionQuery).
type FilterDefinitionQuery struct {
	// Only return items with the given typeid.
	Typeid int64 `json:"typeid"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`

	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`
}

// Used when importing an order with optiondbundle tickets
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/ImportBundleTicket).
type ImportBundleTicket struct {
	// Manually select a specific ticket.
	Id int64 `json:"id,omitempty"`

	// If boolean is set to true, the price field is used (even if set to 0) to
	// determine the price for this ticket
	Overrideprice bool `json:"overrideprice,omitempty"`

	// The price for this bundle ticket, if this value is greater than 0 it's always
	// used. If one of the bundletickets has a price, all bundletickets should have a
	// price. Setting this overrides the default behaviour of the configured bundle.
	Price float64 `json:"price,omitempty"`

	// Seatzone ID
	Seatzoneid int64 `json:"seatzoneid,omitempty"`

	// The tickettype ID for the ticket.
	Tickettypeid int64 `json:"tickettypeid"`

	// The tickettypeprice ID for the ticket. This field is required if bundletickets
	// are specified for a fixed bundle. When importing an optionbundle, if one of the
	// bundletickets has a tickettypepriceid, all bundletickets should have one.
	// Setting this, overrides the default behaviour of the configured bundle
	Tickettypepriceid int64 `json:"tickettypepriceid,omitempty"`
}

// Used to import an order.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/ImportOrder).
type ImportOrder struct {
	// Order ID
	Orderid int64 `json:"orderid"`

	// Order code
	//
	// Used as a unique identifier in web sales.
	Code string `json:"code,omitempty"`

	// Customer ID
	Customerid int64 `json:"customerid,omitempty"`

	// Address used when delivering physically
	Deliveryaddress *Address `json:"deliveryaddress,omitempty"`

	// See delivery scenarios
	// (https://www.ticketmatic.com/docs/api/settings/ticketsales/deliveryscenarios)
	// for more info.
	Deliveryscenarioid int64 `json:"deliveryscenarioid,omitempty"`

	// Delivery status
	//
	// Possible values:
	//
	// * 2601: Not delivered
	//
	// * 2602: Delivered
	//
	// * 2603: Changed after delivery
	Deliverystatus int64 `json:"deliverystatus,omitempty"`

	// Indicates if the expired order has been handled. If set to false when importing,
	// Ticketmatic will send our expiry mails if configured.
	Expiryhandled bool `json:"expiryhandled,omitempty"`

	// When the order will expire. If this is specified expiryhandled should also be
	// specified.
	Expiryts Time `json:"expiryts,omitempty"`

	// Order fees for the order
	Ordercosts []*ImportOrdercost `json:"ordercosts"`

	// Payments in the order
	Payments []*ImportPayment `json:"payments"`

	// See payment scenarios
	// (https://www.ticketmatic.com/docs/api/settings/ticketsales/paymentscenarios) for
	// more info.
	Paymentscenarioid int64 `json:"paymentscenarioid,omitempty"`

	// Products in the order
	Products []*ImportProduct `json:"products"`

	// Indicates if the overdue order has been handled. If set to false when importing,
	// Ticketmatic will send our reminder mails if configured.
	Rappelhandled bool `json:"rappelhandled,omitempty"`

	// When a reminder mail will be sent. If this is specified rappelhandled should
	// also be specified.
	Rappelts Time `json:"rappelts,omitempty"`

	// See sales channels
	// (https://www.ticketmatic.com/docs/api/settings/ticketsales/saleschannels) for
	// more info.
	Saleschannelid int64 `json:"saleschannelid"`

	// Tickets in the order
	Tickets []*ImportTicket `json:"tickets"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`
}

// Used when importing orders.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/ImportOrdercost).
type ImportOrdercost struct {
	// The amount for this ordercost, can only be specified with manual ordercosts
	Amount float64 `json:"amount,omitempty"`

	// Id of the service charge to use for this ordercost
	Servicechargedefinitionid int64 `json:"servicechargedefinitionid"`
}

// Used when importing an order.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/ImportPayment).
type ImportPayment struct {
	// Amount
	Amount float64 `json:"amount"`

	// Timestamp of payment
	Paidts Time `json:"paidts"`

	// Payment method id
	Paymentmethodid int64 `json:"paymentmethodid"`

	// Additional properties for the payment. Can contain a variable structure.
	Properties map[string]interface{} `json:"properties,omitempty"`

	// Voucher code that was used for this payment
	Vouchercode string `json:"vouchercode,omitempty"`

	// Voucher code id that was used for this payment
	Vouchercodeid int64 `json:"vouchercodeid,omitempty"`
}

// Used when importing orders.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/ImportProduct).
type ImportProduct struct {
	// List of tickets that belong to this bundle.
	Bundletickets []*ImportBundleTicket `json:"bundletickets"`

	// If boolean is set to true, the price field is used (even if set to 0) to
	// determine the price for this product
	Overrideprice bool `json:"overrideprice,omitempty"`

	// Product price, will always be used if larger than 0.
	Price float64 `json:"price,omitempty"`

	// Indicate which contact is the holder of this product. Currently only used with
	// bundles.
	Productholderid int64 `json:"productholderid,omitempty"`

	// The id for the product you want to add.
	Productid int64 `json:"productid"`

	// The property values for the product.
	Properties map[string]string `json:"properties,omitempty"`

	// If this product references a voucher, set the amount to reserve for this
	// voucher.
	Voucheramount float64 `json:"voucheramount,omitempty"`

	// If this product references a voucher, set the code for the voucher that will be
	// created. If not set, the code will be generated.
	Vouchercode string `json:"vouchercode,omitempty"`

	// If this product references a voucher, set the expiry timestamp for the
	// vouchercode that will be created. If not set, the default timestamp configured
	// in the voucher will be set.
	Voucherexpiryts Time `json:"voucherexpiryts,omitempty"`
}

// Used when importing order.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/ImportTicket).
type ImportTicket struct {
	// Manually select a specific ticket.
	Id int64 `json:"id,omitempty"`

	// If boolean is set to true, the price field is used (even if set to 0) to
	// determine the price for this ticket
	Overrideprice bool `json:"overrideprice,omitempty"`

	// If boolean is set to true, the servicecharge field is used (even if set to 0) to
	// determine the servicecharge for this ticket
	Overrideservicecharge bool `json:"overrideservicecharge,omitempty"`

	// Ticket price, will always be used if larger then 0.
	Price float64 `json:"price,omitempty"`

	// Seatzone ID
	Seatzoneid int64 `json:"seatzoneid,omitempty"`

	// Service charge for this ticket
	Servicecharge float64 `json:"servicecharge,omitempty"`

	// If this ticket should be linked to a contact, set the ticketholderid
	Ticketholderid int64 `json:"ticketholderid,omitempty"`

	// The tickettype ID for the ticket.
	Tickettypeid int64 `json:"tickettypeid"`

	// The ticket type price ID for the new ticket. Either tickettypepriceid or
	// optionbundleid should be specified, not both.
	Tickettypepriceid int64 `json:"tickettypepriceid,omitempty"`

	// Voucher code to use (if any)
	Vouchercode string `json:"vouchercode,omitempty"`

	// The voucher code to link to this ticket
	Vouchercodeid int64 `json:"vouchercodeid,omitempty"`
}

// Key-value item
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/KeyValueItem).
type KeyValueItem struct {
	// Key
	Key string `json:"key"`

	// Value
	Value string `json:"value"`
}

// A single lock type.
//
// More info: see the get operation
// (https://www.ticketmatic.com/docs/api/settings/ticketsales/locktypes/get) and
// the lock types endpoint
// (https://www.ticketmatic.com/docs/api/settings/ticketsales/locktypes).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/LockType).
type LockType struct {
	// Unique ID
	//
	// Note: Ignored when creating a new lock type.
	//
	// Note: Ignored when updating an existing lock type.
	Id int64 `json:"id"`

	// Name for the lock type
	Name string `json:"name"`

	// Indicates whether this lock is a hard lock (meaning that it normally never will
	// be released and does not count for the inventory) or a soft lock
	Ishardlock bool `json:"ishardlock"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new lock type.
	//
	// Note: Ignored when updating an existing lock type.
	Isarchived bool `json:"isarchived"`

	// Created timestamp
	//
	// Note: Ignored when creating a new lock type.
	//
	// Note: Ignored when updating an existing lock type.
	Createdts Time `json:"createdts"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new lock type.
	//
	// Note: Ignored when updating an existing lock type.
	Lastupdatets Time `json:"lastupdatets"`
}

// Set of parameters used to filter lock types.
//
// More info: see lock type (https://www.ticketmatic.com/docs/api/types/LockType),
// the getlist operation
// (https://www.ticketmatic.com/docs/api/settings/ticketsales/locktypes/getlist)
// and the lock types endpoint
// (https://www.ticketmatic.com/docs/api/settings/ticketsales/locktypes).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/LockTypeQuery).
type LockTypeQuery struct {
	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`

	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`
}

// Log item returned when requesting the log history of an order
// (https://www.ticketmatic.com/docs/api/types/Order).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/LogItem).
type LogItem struct {
	// Id of the log item
	Id int64 `json:"id"`

	// Order id
	Orderid int64 `json:"orderid"`

	// Log item type
	Typeid int64 `json:"typeid"`

	// Info
	Info map[string]interface{} `json:"info,omitempty"`

	// Lookup info
	Lookupinfo map[string]interface{} `json:"lookupinfo,omitempty"`

	// Model
	Model map[string]interface{} `json:"model,omitempty"`

	// Log item timestamp
	Ts Time `json:"ts"`

	// User id
	Userid int64 `json:"userid"`

	// User name
	Username string `json:"username"`
}

// The logical plan describes the structure and layout of seats in a zone.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/LogicalPlan).
type LogicalPlan struct {
	// The ID of the zone
	Id int64 `json:"id"`

	// The name of the zone
	Name string `json:"name"`

	// The rows layout
	Rows string `json:"rows"`
}

// A single Order.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/Order).
type Order struct {
	// Order ID
	Orderid int64 `json:"orderid,omitempty"`

	// Total amount paid
	//
	// Note: Ignored when importing orders.
	Amountpaid float64 `json:"amountpaid"`

	// Whether or not auto / script order fees should be calculated
	//
	// Note: Ignored when importing orders.
	CalculateOrdercosts bool `json:"calculate_ordercosts"`

	// Order code
	//
	// Used as a unique identifier in web sales.
	Code string `json:"code,omitempty"`

	// Customer ID
	Customerid int64 `json:"customerid"`

	// Information on the deferred payment scenario. Structure depends on payment
	// method
	//
	// Note: Ignored when importing orders.
	Deferredpaymentproperties map[string]interface{} `json:"deferredpaymentproperties,omitempty"`

	// Address used when delivering physically
	Deliveryaddress *Address `json:"deliveryaddress,omitempty"`

	// Delivery scenario ID
	//
	// See delivery scenarios
	// (https://www.ticketmatic.com/docs/api/settings/ticketsales/deliveryscenarios)
	// for more info.
	Deliveryscenarioid int64 `json:"deliveryscenarioid,omitempty"`

	// Delivery status
	//
	// Possible values:
	//
	// * 2601: Not delivered
	//
	// * 2602: Delivered
	//
	// * 2603: Changed after delivery
	Deliverystatus int64 `json:"deliverystatus"`

	// Whether the expired order has been handled (and optionally expiry mail has been
	// sent)
	Expiryhandled bool `json:"expiryhandled"`

	// When the order will expire
	Expiryts Time `json:"expiryts"`

	// Has customer authenticated?
	//
	// Note: Ignored when importing orders.
	Isauthenticatedcustomer bool `json:"isauthenticatedcustomer"`

	// Related objects
	//
	// See the lookup fields on the getlist operation
	// (https://www.ticketmatic.com/docs/api/orders/getlist) for a full description.
	//
	// Note: Ignored when importing orders.
	Lookup map[string]interface{} `json:"lookup,omitempty"`

	// Number of tickets in the order. Read-only
	//
	// Note: Ignored when importing orders.
	Nbroftickets int64 `json:"nbroftickets"`

	// Order fees for the order
	Ordercosts []*Ordercost `json:"ordercosts"`

	// Payments for the order
	Payments []*Payment `json:"payments"`

	// Payment scenario ID
	//
	// See payment scenarios
	// (https://www.ticketmatic.com/docs/api/settings/ticketsales/paymentscenarios) for
	// more info.
	Paymentscenarioid int64 `json:"paymentscenarioid,omitempty"`

	// Payment status
	//
	// Possible values:
	//
	// * 0: Incomplete
	//
	// * 1: Fully paid
	//
	// * 2: Overpaid
	//
	// Note: Ignored when importing orders.
	Paymentstatus int64 `json:"paymentstatus"`

	// Products in the order
	Products []*OrderProduct `json:"products"`

	// Promocodes active for the Order
	//
	// Note: Ignored when importing orders.
	Promocodes []string `json:"promocodes"`

	// Queue tokens for rate limiting
	//
	// Note: Ignored when importing orders.
	Queuetokens []int64 `json:"queuetokens"`

	// Whether the overdue order has been handled (and optionally reminder mail has
	// been sent)
	Rappelhandled bool `json:"rappelhandled"`

	// When the reminder mail will be sent
	Rappelts Time `json:"rappelts"`

	// Sales channel ID
	//
	// See sales channels
	// (https://www.ticketmatic.com/docs/api/settings/ticketsales/saleschannels) for
	// more info.
	Saleschannelid int64 `json:"saleschannelid"`

	// Order status
	//
	// Possible values:
	//
	// * 21001: Unconfirmed
	//
	// * 21002: Confirmed
	//
	// * 21003: Archived
	//
	// Note: Ignored when importing orders.
	Status int64 `json:"status"`

	// Tickets in the order
	Tickets []*OrderTicket `json:"tickets"`

	// Total order amount
	//
	// Includes all costs and fees.
	//
	// Note: Ignored when importing orders.
	Totalamount float64 `json:"totalamount"`

	// Reference to the webskin that is used for showing the orderdetail page.
	//
	// Note: Ignored when importing orders.
	Webskinid int64 `json:"webskinid"`

	// Created timestamp
	Createdts Time `json:"createdts,omitempty"`

	// Last updated timestamp
	Lastupdatets Time `json:"lastupdatets,omitempty"`

	// Custom fields
	CustomFields map[string]interface{} `json:"-"`
}

// Custom unmarshaller with support for custom fields
func (o *Order) UnmarshalJSON(data []byte) error {
	// Alias the type, to avoid calling UnmarshalJSON. Unpack it.
	type tmp Order
	var obj tmp
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return err
	}

	*o = Order(obj)

	// Unpack it again, this time to a map, so we can pull out the custom fields.
	var raw map[string]interface{}
	err = json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	o.CustomFields = make(map[string]interface{})
	for key, val := range raw {
		if strings.HasPrefix(key, "c_") {
			o.CustomFields[key[2:]] = val
		}
	}

	// Note: We're doing a double JSON decode here, I'd love to get rid of it
	// but I'm not sure how we can do this easily. Suggestions welcome:
	// developers@ticketmatic.com!

	return nil
}

// Custom marshaller with support for custom fields
func (o *Order) MarshalJSON() ([]byte, error) {
	// Use a custom type to avoid the custom marshaller, marshal the data.
	type tmp struct {
		Orderid                   int64                  `json:"orderid,omitempty"`
		Amountpaid                float64                `json:"amountpaid,omitempty"`
		CalculateOrdercosts       bool                   `json:"calculate_ordercosts,omitempty"`
		Code                      string                 `json:"code,omitempty"`
		Customerid                int64                  `json:"customerid,omitempty"`
		Deferredpaymentproperties map[string]interface{} `json:"deferredpaymentproperties,omitempty"`
		Deliveryaddress           *Address               `json:"deliveryaddress,omitempty"`
		Deliveryscenarioid        int64                  `json:"deliveryscenarioid,omitempty"`
		Deliverystatus            int64                  `json:"deliverystatus,omitempty"`
		Expiryhandled             bool                   `json:"expiryhandled,omitempty"`
		Expiryts                  Time                   `json:"expiryts,omitempty"`
		Isauthenticatedcustomer   bool                   `json:"isauthenticatedcustomer,omitempty"`
		Lookup                    map[string]interface{} `json:"lookup,omitempty"`
		Nbroftickets              int64                  `json:"nbroftickets,omitempty"`
		Ordercosts                []*Ordercost           `json:"ordercosts,omitempty"`
		Payments                  []*Payment             `json:"payments,omitempty"`
		Paymentscenarioid         int64                  `json:"paymentscenarioid,omitempty"`
		Paymentstatus             int64                  `json:"paymentstatus,omitempty"`
		Products                  []*OrderProduct        `json:"products,omitempty"`
		Promocodes                []string               `json:"promocodes,omitempty"`
		Queuetokens               []int64                `json:"queuetokens,omitempty"`
		Rappelhandled             bool                   `json:"rappelhandled,omitempty"`
		Rappelts                  Time                   `json:"rappelts,omitempty"`
		Saleschannelid            int64                  `json:"saleschannelid,omitempty"`
		Status                    int64                  `json:"status,omitempty"`
		Tickets                   []*OrderTicket         `json:"tickets,omitempty"`
		Totalamount               float64                `json:"totalamount,omitempty"`
		Webskinid                 int64                  `json:"webskinid,omitempty"`
		Createdts                 Time                   `json:"createdts,omitempty"`
		Lastupdatets              Time                   `json:"lastupdatets,omitempty"`
	}

	obj := tmp{
		Orderid:                   o.Orderid,
		Amountpaid:                o.Amountpaid,
		CalculateOrdercosts:       o.CalculateOrdercosts,
		Code:                      o.Code,
		Customerid:                o.Customerid,
		Deferredpaymentproperties: o.Deferredpaymentproperties,
		Deliveryaddress:           o.Deliveryaddress,
		Deliveryscenarioid:        o.Deliveryscenarioid,
		Deliverystatus:            o.Deliverystatus,
		Expiryhandled:             o.Expiryhandled,
		Expiryts:                  o.Expiryts,
		Isauthenticatedcustomer:   o.Isauthenticatedcustomer,
		Lookup:                    o.Lookup,
		Nbroftickets:              o.Nbroftickets,
		Ordercosts:                o.Ordercosts,
		Payments:                  o.Payments,
		Paymentscenarioid:         o.Paymentscenarioid,
		Paymentstatus:             o.Paymentstatus,
		Products:                  o.Products,
		Promocodes:                o.Promocodes,
		Queuetokens:               o.Queuetokens,
		Rappelhandled:             o.Rappelhandled,
		Rappelts:                  o.Rappelts,
		Saleschannelid:            o.Saleschannelid,
		Status:                    o.Status,
		Tickets:                   o.Tickets,
		Totalamount:               o.Totalamount,
		Webskinid:                 o.Webskinid,
		Createdts:                 o.Createdts,
		Lastupdatets:              o.Lastupdatets,
	}
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	// Unpack it again, to get the wire representation
	var raw map[string]interface{}
	err = json.Unmarshal(data, &raw)
	if err != nil {
		return nil, err
	}

	for key, val := range o.CustomFields {
		raw["c_"+key] = val
	}

	// Pack it again
	return json.Marshal(raw)

	// Note: Like UnmarshalJSON, this is quite crazy. But it works beautifully.
	// Know a way to do this better? Get in touch!
}

// A single order fee.
//
// More info: see the get operation
// (https://www.ticketmatic.com/docs/api/settings/ticketsales/orderfees/get) and
// the order fees endpoint
// (https://www.ticketmatic.com/docs/api/settings/ticketsales/orderfees).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/OrderFee).
type OrderFee struct {
	// Unique ID
	//
	// Note: Ignored when creating a new order fee definition.
	//
	// Note: Ignored when creating a new order fee.
	Id int64 `json:"id"`

	// Type of the order fee. Can be Automatic (2401), Script (2402) or Manual (2403)
	Typeid int64 `json:"typeid"`

	// Name for the order fee
	Name string `json:"name"`

	// Definition of the rule that defines when the order fee will be applied
	//
	// Note: Not set when retrieving a list of order fee definitions.
	//
	// Note: Not set when retrieving a list of order fees.
	Rule *OrderfeeRule `json:"rule,omitempty"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new order fee definition.
	//
	// Note: Ignored when creating a new order fee.
	Isarchived bool `json:"isarchived"`

	// Archived timestamp
	//
	// Note: Ignored when creating a new order fee definition.
	//
	// Note: Ignored when creating a new order fee.
	Archivedts Time `json:"archivedts"`

	// Created timestamp
	//
	// Note: Ignored when creating a new order fee definition.
	//
	// Note: Ignored when creating a new order fee.
	Createdts Time `json:"createdts"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new order fee definition.
	//
	// Note: Ignored when creating a new order fee.
	Lastupdatets Time `json:"lastupdatets"`
}

// A single order fee definition.
//
// More info: see the get operation
// (https://www.ticketmatic.com/docs/api/settings/pricing/orderfeedefinitions/get)
// and the order fee definitions endpoint
// (https://www.ticketmatic.com/docs/api/settings/pricing/orderfeedefinitions).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/OrderFeeDefinition).
type OrderFeeDefinition struct {
	// Unique ID
	//
	// Note: Ignored when creating a new order fee definition.
	//
	// Note: Ignored when creating a new order fee.
	Id int64 `json:"id"`

	// Type of the order fee. Can be Automatic (2401), Script (2402) or Manual (2403)
	Typeid int64 `json:"typeid"`

	// Name for the order fee
	Name string `json:"name"`

	// Definition of the rule that defines when the order fee will be applied
	//
	// Note: Not set when retrieving a list of order fee definitions.
	//
	// Note: Not set when retrieving a list of order fees.
	Rule *OrderfeeRule `json:"rule,omitempty"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new order fee definition.
	//
	// Note: Ignored when creating a new order fee.
	Isarchived bool `json:"isarchived"`

	// Archived timestamp
	//
	// Note: Ignored when creating a new order fee definition.
	//
	// Note: Ignored when creating a new order fee.
	Archivedts Time `json:"archivedts"`

	// Created timestamp
	//
	// Note: Ignored when creating a new order fee definition.
	//
	// Note: Ignored when creating a new order fee.
	Createdts Time `json:"createdts"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new order fee definition.
	//
	// Note: Ignored when creating a new order fee.
	Lastupdatets Time `json:"lastupdatets"`
}

// Set of parameters used to filter order fee definitions.
//
// More info: see order fee definition
// (https://www.ticketmatic.com/docs/api/types/OrderFeeDefinition), the getlist
// operation
// (https://www.ticketmatic.com/docs/api/settings/pricing/orderfeedefinitions/getlist)
// and the order fee definitions endpoint
// (https://www.ticketmatic.com/docs/api/settings/pricing/orderfeedefinitions).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/OrderFeeDefinitionQuery).
type OrderFeeDefinitionQuery struct {
	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`

	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`
}

// Set of parameters used to filter order fees.
//
// More info: see order fee (https://www.ticketmatic.com/docs/api/types/OrderFee),
// the getlist operation
// (https://www.ticketmatic.com/docs/api/settings/ticketsales/orderfees/getlist)
// and the order fees endpoint
// (https://www.ticketmatic.com/docs/api/settings/ticketsales/orderfees).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/OrderFeeQuery).
type OrderFeeQuery struct {
	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`

	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`
}

// Used when requesting orders, to filter orders.
//
// Specify any of the supported fields to filter the list of orders.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/OrderFilter).
type OrderFilter struct {
	// Only include orders older than the given timestamp
	Createdsince Time `json:"createdsince,omitempty"`

	// Filter orders based on customer
	Customerid int64 `json:"customerid,omitempty"`

	// Filter orders based on saleschannel
	Saleschannelid int64 `json:"saleschannelid,omitempty"`

	// Only include orders with a given status
	//
	// Possible values:
	//
	// * 21001: Unconfirmed orders
	//
	// * 21002: Confirmed orders
	//
	// * 21003: Archived orders
	Status int64 `json:"status,omitempty"`
}

// Order ID reservation
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/OrderIdReservation).
type OrderIdReservation struct {
	// Maximum ID to reserve
	Id int64 `json:"id"`
}

// Import status per order
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/OrderImportStatus).
type OrderImportStatus struct {
	// Order ID
	Id int64 `json:"id"`

	// Error message, if failed
	Error string `json:"error"`

	// Whether the import succeeded
	Ok bool `json:"ok"`
}

// A single order mail template.
//
// More info: see the get operation
// (https://www.ticketmatic.com/docs/api/settings/communicationanddesign/ordermails/get)
// and the order mail templates endpoint
// (https://www.ticketmatic.com/docs/api/settings/communicationanddesign/ordermails).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/OrderMailTemplate).
type OrderMailTemplate struct {
	// Unique ID
	//
	// Note: Ignored when creating a new order mail template.
	//
	// Note: Ignored when updating an existing order mail template.
	Id int64 `json:"id"`

	// The type of this order mail template, defines where this template is used. The
	// available values for this field can be found on the order mail template overview
	// (https://www.ticketmatic.com/docs/api/settings/communicationanddesign/ordermails)
	// page.
	Typeid int64 `json:"typeid"`

	// Name of the order mail template
	Name string `json:"name"`

	// Message body
	//
	// Note: Not set when retrieving a list of order mail templates.
	Body string `json:"body"`

	// Subject line for the order mail template
	//
	// Note: Not set when retrieving a list of order mail templates.
	Subject string `json:"subject"`

	// A map of language codes to gettext .po files
	// (http://en.wikipedia.org/wiki/Gettext). More info can be found on the order mail
	// template overview
	// (https://www.ticketmatic.com/docs/api/settings/communicationanddesign/ordermails)
	// page.
	//
	// Note: Not set when retrieving a list of order mail templates.
	Translations map[string]string `json:"translations,omitempty"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new order mail template.
	//
	// Note: Ignored when updating an existing order mail template.
	Isarchived bool `json:"isarchived"`

	// Created timestamp
	//
	// Note: Ignored when creating a new order mail template.
	//
	// Note: Ignored when updating an existing order mail template.
	Createdts Time `json:"createdts"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new order mail template.
	//
	// Note: Ignored when updating an existing order mail template.
	Lastupdatets Time `json:"lastupdatets"`
}

// Set of parameters used to filter order mail templates.
//
// More info: see order mail template
// (https://www.ticketmatic.com/docs/api/types/OrderMailTemplate), the getlist
// operation
// (https://www.ticketmatic.com/docs/api/settings/communicationanddesign/ordermails/getlist)
// and the order mail templates endpoint
// (https://www.ticketmatic.com/docs/api/settings/communicationanddesign/ordermails).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/OrderMailTemplateQuery).
type OrderMailTemplateQuery struct {
	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`

	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`
}

// A single product in an order.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/OrderProduct).
type OrderProduct struct {
	// Orderproduct ID
	Id int64 `json:"id"`

	// Order ID
	Orderid int64 `json:"orderid"`

	// Unique code for this orderproduct
	Code string `json:"code"`

	// Contact ID: the holder of this product
	Contactid int64 `json:"contactid"`

	// Ticket price
	Price float64 `json:"price"`

	// Product ID
	Productid int64 `json:"productid"`

	// Property values for this product
	Properties map[string]string `json:"properties,omitempty"`

	// Vouchercode ID for the voucher that is linked to this orderproduct
	Vouchercodeid int64 `json:"vouchercodeid"`
}

// Filter parameters to fetch a list of orders
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/OrderQuery).
type OrderQuery struct {
	// A SQL query that returns order IDs
	//
	// Can be used to do arbitrary filtering. See the database documentation for order
	// (https://www.ticketmatic.com/docs/db/order) for more information.
	Filter string `json:"filter,omitempty"`

	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// Only include orders that have been updated since the given timestamp.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`

	// Limit results to at most the given amount of orders.
	Limit int64 `json:"limit,omitempty"`

	// Skip the first X orders.
	Offset int64 `json:"offset,omitempty"`

	// Order by the given field.
	//
	// Supported values: createdts, lastupdatets.
	Orderby string `json:"orderby,omitempty"`

	// Output format.
	//
	// Possible values:
	//
	// * ids: Only fill the ID field
	//
	// * minimal: A minimal set of order fields
	//
	// * default: Return all order fields (also used when the output parameter is
	// omitted)
	//
	// * withlookup: Returns all order fields and an additional lookup field which
	// contains all dependent objects
	Output string `json:"output,omitempty"`

	// A text filter string.
	//
	// Matches against the order ID or the customer details..
	Searchterm string `json:"searchterm,omitempty"`

	// Filters the orders based on a given set of fields. Currently supports:
	// createdsince, saleschannelid, customerid, status.
	Simplefilter *OrderFilter `json:"simplefilter,omitempty"`
}

// A single ticket in an order.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/OrderTicket).
type OrderTicket struct {
	// Ticket ID
	Id int64 `json:"id"`

	// Order ID
	Orderid int64 `json:"orderid"`

	// The abo ticketid for tickets that belong to an abo
	Aboparentid int64 `json:"aboparentid"`

	// The id of the product this ticket is linked to
	Bundleid int64 `json:"bundleid"`

	// Event id
	Eventid int64 `json:"eventid"`

	// Ticket price
	Price float64 `json:"price"`

	// Pricetype id
	Pricetypeid int64 `json:"pricetypeid"`

	// Description of the ticket
	Seatdescription string `json:"seatdescription"`

	// Name of the seat
	Seatname string `json:"seatname"`

	// Seatzone ID
	Seatzoneid int64 `json:"seatzoneid"`

	// Service charge
	Servicecharge float64 `json:"servicecharge"`

	// Ticket holder ID
	Ticketholderid int64 `json:"ticketholderid"`

	// Name for the ticket holder
	Ticketname string `json:"ticketname"`

	// Contingent ID
	Tickettypeid int64 `json:"tickettypeid"`

	// Contingent name
	Tickettypename string `json:"tickettypename"`

	// Id for the tickettypeprice of this ticket for the order
	Tickettypepriceid int64 `json:"tickettypepriceid"`

	// The voucher code that was linked to this ticket
	Vouchercodeid int64 `json:"vouchercodeid"`
}

// Order tickettype
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/OrderTickettype).
type OrderTickettype struct {
	// Tickettype id
	Id int64 `json:"id"`

	// Tickettype name
	Name string `json:"name"`

	// Tickettype full name
	Fulltypename string `json:"fulltypename"`
}

// A single order fee for an order.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/Ordercost).
type Ordercost struct {
	// Order ID
	Orderid int64 `json:"orderid"`

	// Payment amount
	Amount float64 `json:"amount"`

	// Order fee ID
	Servicechargedefinitionid int64 `json:"servicechargedefinitionid"`
}

// More info about order fees can be found here
// (https://www.ticketmatic.com/docs/api/settings/ticketsales/orderfees).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/OrderfeeAutoRule).
type OrderfeeAutoRule struct {
	// The delivery scenarios that this order fee is applicable for. If not set it
	// defaults to 'all'. This is only needed if the order fee type is set to
	// automatic.
	Deliveryscenarioids []int64 `json:"deliveryscenarioids"`

	// The payment scenarios that this order fee is applicable for. If not set it
	// default to 'all'. This is only needed if the order fee type is set to automatic.
	Paymentscenarioids []int64 `json:"paymentscenarioids"`

	// The sales channels that this order fee is applicable for. If not set it defaults
	// to 'all'. This is only needed if the order fee type is set to automatic.
	Saleschannelids []int64 `json:"saleschannelids"`

	// Can be fixedfee or percentagefee. Defauls to fixedfee. This is only needed if
	// the order fee type is set to automatic.
	Status string `json:"status"`

	// The value (amount) that will be added to the order. Is required if the order fee
	// type is set to automatic.
	Value float64 `json:"value"`
}

// More info about order fees can be found here
// (https://www.ticketmatic.com/docs/api/settings/ticketsales/orderfees).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/OrderfeeRule).
type OrderfeeRule struct {
	// This is required if the order fee type is set to automatic. It is a set of rules
	// that define the order fee.
	Auto []*OrderfeeAutoRule `json:"auto"`

	// This can be set if the order fee type is set to script. It allows adding extra
	// information to the script environment.
	Context []*OrderfeeScriptContext `json:"context"`

	// This is required if the order fee type is set to script. The javascript needs to
	// return a value.
	Script string `json:"script"`
}

// More info about order fees can be found here
// (https://www.ticketmatic.com/docs/api/settings/ticketsales/orderfees).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/OrderfeeScriptContext).
type OrderfeeScriptContext struct {
	// If set to true the query will be cached for 60 seconds. If not set the query
	// will be executed again every time a script is executed.
	Cacheable bool `json:"cacheable"`

	// The name of the variable that will be added to the script environment.
	Key string `json:"key"`

	// The query that will be executed on the public data model. The result will be
	// available in the script environment.
	Query string `json:"query"`
}

// A single payment.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/Payment).
type Payment struct {
	// Payment ID
	Id int64 `json:"id"`

	// Order ID
	Orderid int64 `json:"orderid"`

	// Payment amount
	Amount float64 `json:"amount"`

	// Timestamp of payment
	Paidts Time `json:"paidts"`

	// Payment method ID
	Paymentmethodid int64 `json:"paymentmethodid"`

	// Additional properties for the payment. Structure depends on the payment method
	Properties map[string]interface{} `json:"properties,omitempty"`

	// Id for the original payment if this payment is a refund
	Refundpaymentid int64 `json:"refundpaymentid"`

	// Id of the vouchercode to use for this payment
	//
	// Note: Ignored when importing orders.
	Vouchercodeid int64 `json:"vouchercodeid"`
}

// A single payment method.
//
// More info: see the get operation
// (https://www.ticketmatic.com/docs/api/settings/ticketsales/paymentmethods/get)
// and the payment methods endpoint
// (https://www.ticketmatic.com/docs/api/settings/ticketsales/paymentmethods).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/PaymentMethod).
type PaymentMethod struct {
	// Unique ID
	//
	// Note: Ignored when creating a new payment method.
	//
	// Note: Ignored when updating an existing payment method.
	Id int64 `json:"id"`

	// Name of the payment method
	Name string `json:"name"`

	// Specific configuration for the payment method, content depends on the payment
	// method type.
	//
	// Note: Not set when retrieving a list of payment methods.
	Config map[string]interface{} `json:"config,omitempty"`

	// Internal remark, will not be shown to customers
	Internalremark string `json:"internalremark"`

	// Type of the paymentmethod.
	Paymentmethodtypeid int64 `json:"paymentmethodtypeid"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new payment method.
	//
	// Note: Ignored when updating an existing payment method.
	Isarchived bool `json:"isarchived"`

	// Created timestamp
	//
	// Note: Ignored when creating a new payment method.
	//
	// Note: Ignored when updating an existing payment method.
	Createdts Time `json:"createdts"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new payment method.
	//
	// Note: Ignored when updating an existing payment method.
	Lastupdatets Time `json:"lastupdatets"`
}

// Set of parameters used to filter payment methods.
//
// More info: see payment method
// (https://www.ticketmatic.com/docs/api/types/PaymentMethod), the getlist
// operation
// (https://www.ticketmatic.com/docs/api/settings/ticketsales/paymentmethods/getlist)
// and the payment methods endpoint
// (https://www.ticketmatic.com/docs/api/settings/ticketsales/paymentmethods).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/PaymentMethodQuery).
type PaymentMethodQuery struct {
	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`

	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`
}

// Info for requesting an immediate payment in an order
// (https://www.ticketmatic.com/docs/api/types/Order).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/PaymentRequest).
type PaymentRequest struct {
	// The language to be used during the payment processing
	Language string `json:"language"`

	// The returnurl that will be called after the payment request was done.
	Returnurl string `json:"returnurl"`
}

// A single payment scenario.
//
// More info: see the get operation
// (https://www.ticketmatic.com/docs/api/settings/ticketsales/paymentscenarios/get)
// and the payment scenarios endpoint
// (https://www.ticketmatic.com/docs/api/settings/ticketsales/paymentscenarios).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/PaymentScenario).
type PaymentScenario struct {
	// Unique ID
	//
	// Note: Ignored when creating a new payment scenario.
	//
	// Note: Ignored when updating an existing payment scenario.
	Id int64 `json:"id"`

	// Type for the payment scenario. Can be 'Immediate payment' (2701), 'Mollie bank
	// transfer' (2702), 'Regular bank transfer' (2703), 'Deferred online payment'
	// (2704), 'Deferred other' (2705).
	Typeid int64 `json:"typeid"`

	// Name of the payment scenario
	Name string `json:"name"`

	// Rules that define in what conditions this payment scenario is available
	//
	// Note: Not set when retrieving a list of payment scenarios.
	Availability *PaymentscenarioAvailability `json:"availability,omitempty"`

	// Beneficiary for the bank account number. Only used for type 2703 (Regular bank
	// transfer)
	Bankaccountbeneficiary string `json:"bankaccountbeneficiary,omitempty"`

	// BIC code for the bank account number. Only used for type 2703 (Regular bank
	// transfer)
	Bankaccountbic string `json:"bankaccountbic,omitempty"`

	// Bank account number to be used. Only used for type 2703 (Regular bank transfer)
	Bankaccountnumber string `json:"bankaccountnumber,omitempty"`

	// Rules that define when an order becomes expired. Not used for type 2701.
	//
	// Note: Not set when retrieving a list of payment scenarios.
	Expiryparameters *PaymentscenarioExpiryParameters `json:"expiryparameters,omitempty"`

	// An internal remark, which is never shown to customers. Can be used to
	// distinguish identically named payment scenarios.
	//
	// For example: You could have two VISA scenarios, one for the web sales and one
	// for the box office, each will have different fee configurations. Both will be
	// named VISA, this field can be used to distinguish them.
	Internalremark string `json:"internalremark,omitempty"`

	// Link to the order mail template that will be sent when the order is expired. Can
	// be 0 to indicate that no mail should be sent. Not used for type 2701.
	OrdermailtemplateidExpiry int64 `json:"ordermailtemplateid_expiry"`

	// Link to the order mail template that will be sent when the order is overdue. Can
	// be 0 to indicate that no mail should be sent. Not used for type 2701.
	OrdermailtemplateidOverdue int64 `json:"ordermailtemplateid_overdue"`

	// Link to the order mail template that will be sent as payment instruction. Can be
	// 0 to indicate that no mail should be sent. Not used for type 2701.
	OrdermailtemplateidPaymentinstruction int64 `json:"ordermailtemplateid_paymentinstruction"`

	// Rules that define when an order becomes overdue. Not used for type 2701.
	//
	// Note: Not set when retrieving a list of payment scenarios.
	Overdueparameters *PaymentscenarioOverdueParameters `json:"overdueparameters,omitempty"`

	// Set of payment methods that are linked to this payment scenario. Depending on
	// the type, this field has different usage.
	Paymentmethods []int64 `json:"paymentmethods"`

	// Short description of the payment scenario, will be shown to customers
	Shortdescription string `json:"shortdescription,omitempty"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new payment scenario.
	//
	// Note: Ignored when updating an existing payment scenario.
	Isarchived bool `json:"isarchived"`

	// Created timestamp
	//
	// Note: Ignored when creating a new payment scenario.
	//
	// Note: Ignored when updating an existing payment scenario.
	Createdts Time `json:"createdts"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new payment scenario.
	//
	// Note: Ignored when updating an existing payment scenario.
	Lastupdatets Time `json:"lastupdatets"`
}

// Set of parameters used to filter payment scenarios.
//
// More info: see payment scenario
// (https://www.ticketmatic.com/docs/api/types/PaymentScenario), the getlist
// operation
// (https://www.ticketmatic.com/docs/api/settings/ticketsales/paymentscenarios/getlist)
// and the payment scenarios endpoint
// (https://www.ticketmatic.com/docs/api/settings/ticketsales/paymentscenarios).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/PaymentScenarioQuery).
type PaymentScenarioQuery struct {
	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`

	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`
}

// A PaymentscenarioAvailability configures in what saleschannels a payment
// scenario is available.
//
// It can also further refine the availability based on a script written in
// JavaScript.
//
// More information about writing order scripts can be found here
// (https://www.ticketmatic.com/docs/tickets/configure_ticket_sales/orderfees).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/PaymentscenarioAvailability).
type PaymentscenarioAvailability struct {
	// The payment scenario will be available for these saleschannels. It this is empty
	// the payment scenario will not be available.
	Saleschannels []int64 `json:"saleschannels"`

	// A Javascript that needs to return a boolean. It has the current order and
	// saleschannel available. More info
	// (https://www.ticketmatic.com/docs/tickets/configure_ticket_sales/orderfees)
	Script string `json:"script"`

	// Indicates if the script will be used.
	Usescript bool `json:"usescript"`
}

// The PaymentscenarioExpiryParameters can only be set when the Paymentscenario is
// of a type deferred payment.
//
// It determines the moment in time when an order expires. It's calculated as
// MIN(<order creation date> + daysafterordercreation, <date of first event in
// order> - daysbeforeevent). If deleteonexpiry is set to true, the order will be
// deleted.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/PaymentscenarioExpiryParameters).
type PaymentscenarioExpiryParameters struct {
	// The amount of days after the paymentscenario was set that the order becomes
	// overdue.
	Daysaftercreation int64 `json:"daysaftercreation"`

	// DEPRECATED, use daysaftercreation. The amount of days after an order has been
	// created that the order becomes overdue.
	Daysafterordercreation int64 `json:"daysafterordercreation"`

	// DEPRECATED, use daysaftercreation. The number of days before an event that an
	// order becomes overdue.
	Daysbeforeevent int64 `json:"daysbeforeevent"`

	// Indicates is the order will be deleted when it's expired.
	Deleteonexpiry bool `json:"deleteonexpiry"`
}

// The PaymentscenarioOverdueParameters can only be set when the Paymentscenario is
// of type deferred payment.
//
// It determines the moment in time when an order becomes overdue. It's calculated
// as MIN(<order creation date> + daysafterordercreation, <date of first event in
// order> - daysbeforeevent).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/PaymentscenarioOverdueParameters).
type PaymentscenarioOverdueParameters struct {
	// The amount of days after the paymentscenario was set that the order becomes
	// overdue.
	Daysaftercreation int64 `json:"daysaftercreation"`

	// DEPRECATED, use daysaftercreation. The amount of days after an order has been
	// created that the order becomes overdue.
	Daysafterordercreation int64 `json:"daysafterordercreation"`

	// DEPRECATED, use daysaftercreation. The number of days before an event that an
	// order becomes overdue.
	Daysbeforeevent int64 `json:"daysbeforeevent"`
}

// A single phone number type.
//
// More info: see the get operation
// (https://www.ticketmatic.com/docs/api/settings/system/phonenumbertypes/get) and
// the phone number types endpoint
// (https://www.ticketmatic.com/docs/api/settings/system/phonenumbertypes).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/PhoneNumberType).
type PhoneNumberType struct {
	// Unique ID
	//
	// Note: Ignored when creating a new phone number type.
	//
	// Note: Ignored when updating an existing phone number type.
	Id int64 `json:"id"`

	// Name of the phone number type
	Name string `json:"name"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new phone number type.
	//
	// Note: Ignored when updating an existing phone number type.
	Isarchived bool `json:"isarchived"`

	// Created timestamp
	//
	// Note: Ignored when creating a new phone number type.
	//
	// Note: Ignored when updating an existing phone number type.
	Createdts Time `json:"createdts"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new phone number type.
	//
	// Note: Ignored when updating an existing phone number type.
	Lastupdatets Time `json:"lastupdatets"`
}

// Set of parameters used to filter phone number types.
//
// More info: see phone number type
// (https://www.ticketmatic.com/docs/api/types/PhoneNumberType), the getlist
// operation
// (https://www.ticketmatic.com/docs/api/settings/system/phonenumbertypes/getlist)
// and the phone number types endpoint
// (https://www.ticketmatic.com/docs/api/settings/system/phonenumbertypes).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/PhoneNumberTypeQuery).
type PhoneNumberTypeQuery struct {
	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`

	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`
}

// See contact (https://www.ticketmatic.com/docs/api/types/Contact) for more
// information.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/Phonenumber).
type Phonenumber struct {
	// Address ID
	//
	// Note: Only available when used for a contact
	// (https://www.ticketmatic.com/docs/api/types/Contact) address.
	Id int64 `json:"id"`

	// Phone number type ID
	Typeid int64 `json:"typeid"`

	// Contact this address belongs to
	//
	// Note: Only available when used for a contact
	// (https://www.ticketmatic.com/docs/api/types/Contact) address.
	Customerid int64 `json:"customerid"`

	// Phone number
	Number string `json:"number"`

	// Phone number type (based on typeid, returned as a convenience)
	Type string `json:"type"`
}

// A single price list.
//
// More info: see the get operation
// (https://www.ticketmatic.com/docs/api/settings/pricing/pricelists/get) and the
// price lists endpoint
// (https://www.ticketmatic.com/docs/api/settings/pricing/pricelists).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/PriceList).
type PriceList struct {
	// Unique ID
	//
	// Note: Ignored when creating a new price list.
	//
	// Note: Ignored when updating an existing price list.
	Id int64 `json:"id"`

	// Name for the pricelist
	Name string `json:"name"`

	// Boolean indicating whether this pricelist has ranks or not
	Hasranks bool `json:"hasranks"`

	// Definition of the actual prices and conditions for the pricelist
	//
	// Note: Not set when retrieving a list of price lists.
	Prices *PricelistPrices `json:"prices,omitempty"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new price list.
	//
	// Note: Ignored when updating an existing price list.
	Isarchived bool `json:"isarchived"`

	// Created timestamp
	//
	// Note: Ignored when creating a new price list.
	//
	// Note: Ignored when updating an existing price list.
	Createdts Time `json:"createdts"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new price list.
	//
	// Note: Ignored when updating an existing price list.
	Lastupdatets Time `json:"lastupdatets"`
}

// Set of parameters used to filter price lists.
//
// More info: see price list
// (https://www.ticketmatic.com/docs/api/types/PriceList), the getlist operation
// (https://www.ticketmatic.com/docs/api/settings/pricing/pricelists/getlist) and
// the price lists endpoint
// (https://www.ticketmatic.com/docs/api/settings/pricing/pricelists).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/PriceListQuery).
type PriceListQuery struct {
	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`

	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`
}

// A single price type.
//
// More info: see the get operation
// (https://www.ticketmatic.com/docs/api/settings/pricing/pricetypes/get) and the
// price types endpoint
// (https://www.ticketmatic.com/docs/api/settings/pricing/pricetypes).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/PriceType).
type PriceType struct {
	// Unique ID
	//
	// Note: Ignored when creating a new price type.
	//
	// Note: Ignored when updating an existing price type.
	Id int64 `json:"id"`

	// The category of this price type, defines how the price is displayed. The
	// available values for this field can be found on the price type overview
	// (https://www.ticketmatic.com/docs/api/settings/pricing/pricetypes) page.
	Typeid int64 `json:"typeid"`

	// Name of the price type
	Name string `json:"name"`

	// A remark that describes the price type. Will be shown to customers.
	Remark string `json:"remark"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new price type.
	//
	// Note: Ignored when updating an existing price type.
	Isarchived bool `json:"isarchived"`

	// Created timestamp
	//
	// Note: Ignored when creating a new price type.
	//
	// Note: Ignored when updating an existing price type.
	Createdts Time `json:"createdts"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new price type.
	//
	// Note: Ignored when updating an existing price type.
	Lastupdatets Time `json:"lastupdatets"`
}

// Set of parameters used to filter price types.
//
// More info: see price type
// (https://www.ticketmatic.com/docs/api/types/PriceType), the getlist operation
// (https://www.ticketmatic.com/docs/api/settings/pricing/pricetypes/getlist) and
// the price types endpoint
// (https://www.ticketmatic.com/docs/api/settings/pricing/pricetypes).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/PriceTypeQuery).
type PriceTypeQuery struct {
	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`

	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`
}

// You can find more information about price in the endpoint documentation
// (https://www.ticketmatic.com/docs/api/settings/pricing/pricelists).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/PricelistPrice).
type PricelistPrice struct {
	// Array of booleans indicating if the corresponding price is available for this
	// PricelistPrice. Should contain the same number of booleans as prices.
	Availabilities []bool `json:"availabilities"`

	// Extra conditions for this price. This can be a promocode, a ticketlimit per
	// order, ... .
	Conditions []*PricelistPriceCondition `json:"conditions"`

	// Optional, and only used for eventspecificprices. Indicates the position of this
	// price in the pricelist.
	Position int64 `json:"position"`

	// The (decimal) prices for this PricelistPrice. If no seatrankids has been set,
	// this should consist of 1 price. If seatrankids are set this should an equal
	// number of prices as the number of seatranks.
	Prices []float64 `json:"prices"`

	// The pricetype for this price.
	Pricetypeid int64 `json:"pricetypeid"`

	// The list of saleschannels for which this PricelistPrice is active.
	Saleschannels []int64 `json:"saleschannels"`
}

// These are the possible condition and example values:
//
// Ticketlimit
//
// There is a limited amount of tickets available for the selected pricetype.
//
//
//    {
//        "type": "ticketlimit",
//        "value": 10
//    }
//
//
//
// Date validity
//
// The price type is only available in this period.
//
// Absolute
//
//
//    {
//        "type": "date",
//        "value": {
//            "datetype": "absolute",
//            "absoluteStart": "2015-05-20",
//            "absoluteEnd": "2015-05-27"
//        }
//    }
//
//
//
// Relative
//
//
//    {
//        "type": "date",
//        "value": {
//            "datetype": "relative_eventdate",
//            "relativeStart": 10,
//            "relativeEnd": 5
//        }
//    }
//
//
//
// Promocode
//
// The price type is only available if the customer provides a promocode.
//
//
//    {
//        "type": "promocode",
//        "value": "TM"
//    }
//
//
//
// Max number of tickets per customer
//
// Limit the maximum number of tickets a customer can buy of this specific price
// type.
//
//
//    {
//        "type": "orderticketlimit",
//        "value": 2
//    }
//
//
//
// Voucherids
//
// When buying a ticket of this pricetype, a valid vouchercode with voucherid one
// of the values should be attached to the ticket.
//
//
//    {
//        "type": "voucherids",
//        "value": [1,2,3]
//    }
//
//
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/PricelistPriceCondition).
type PricelistPriceCondition struct {
	// The type of condition. Possible values:
	//
	// * ticketlimit
	//
	// * date
	//
	// * promocode
	//
	// * orderticketlimit
	//
	// * voucherids
	Type string `json:"type"`

	// The value of this condition. See type for info about what should be filled in.
	Value interface{} `json:"value,omitempty"`
}

// You can find more information about prices in the endpoint documentation
// (https://www.ticketmatic.com/docs/api/settings/pricing/pricelists).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/PricelistPrices).
type PricelistPrices struct {
	// The set of prices for this pricelist.
	Prices []*PricelistPrice `json:"prices"`

	// The seatranks for which this pricelist lists prices.
	Seatrankids []int64 `json:"seatrankids"`
}

// A single product.
//
// More info: see the get operation
// (https://www.ticketmatic.com/docs/api/settings/products/get) and the products
// endpoint (https://www.ticketmatic.com/docs/api/settings/products).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/Product).
type Product struct {
	// Unique ID
	//
	// Note: Ignored when creating a new product.
	//
	// Note: Ignored when updating an existing product.
	Id int64 `json:"id"`

	// Type ID
	//
	// Note: Ignored when updating an existing product.
	Typeid int64 `json:"typeid"`

	// Category for the product
	Categoryid int64 `json:"categoryid"`

	// Optional layout for the product. If not specified, there will be no ticket
	// generated for the product
	Layoutid int64 `json:"layoutid"`

	// Name for the product
	Name string `json:"name"`

	// Unique 12-digit for the product
	Code string `json:"code"`

	// Description for the product
	Description string `json:"description"`

	// Definition of the values for an instance of the product. These depend on the
	// properties
	//
	// Note: Not set when retrieving a list of products.
	Instancevalues *ProductInstancevalues `json:"instancevalues,omitempty"`

	// Definition of possible properties for the product
	Properties []*ProductProperty `json:"properties"`

	// Queue ID
	//
	// See rate limiting (https://www.ticketmatic.com/docs/api/ratelimiting) for more
	// info.
	//
	// Note: Not set when retrieving a list of products.
	Queuetoken int64 `json:"queuetoken"`

	// End of sales
	Saleendts Time `json:"saleendts"`

	// Sales is active for these saleschannels
	//
	// Note: Not set when retrieving a list of products.
	Saleschannels []int64 `json:"saleschannels"`

	// Start of sales
	Salestartts Time `json:"salestartts"`

	// Translations for the product properties
	//
	// Note: Not set when retrieving a list of products.
	Translations map[string]string `json:"translations,omitempty"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new product.
	//
	// Note: Ignored when updating an existing product.
	Isarchived bool `json:"isarchived"`

	// Created timestamp
	//
	// Note: Ignored when creating a new product.
	//
	// Note: Ignored when updating an existing product.
	Createdts Time `json:"createdts"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new product.
	//
	// Note: Ignored when updating an existing product.
	Lastupdatets Time `json:"lastupdatets"`
}

// A single product category.
//
// More info: see the get operation
// (https://www.ticketmatic.com/docs/api/settings/productcategories/get) and the
// product categories endpoint
// (https://www.ticketmatic.com/docs/api/settings/productcategories).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/ProductCategory).
type ProductCategory struct {
	// Unique ID
	//
	// Note: Ignored when creating a new product category.
	//
	// Note: Ignored when updating an existing product category.
	Id int64 `json:"id"`

	// Name for the product category
	Name string `json:"name"`

	// Name for the holder/owner of this product
	Contactname string `json:"contactname"`

	// Name for the holder/owner of this product in plural
	Contactnameplural string `json:"contactnameplural"`

	// Name for the product category in plural
	Nameplural string `json:"nameplural"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new product category.
	//
	// Note: Ignored when updating an existing product category.
	Isarchived bool `json:"isarchived"`

	// Created timestamp
	//
	// Note: Ignored when creating a new product category.
	//
	// Note: Ignored when updating an existing product category.
	Createdts Time `json:"createdts"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new product category.
	//
	// Note: Ignored when updating an existing product category.
	Lastupdatets Time `json:"lastupdatets"`
}

// Set of parameters used to filter product categories.
//
// More info: see product category
// (https://www.ticketmatic.com/docs/api/types/ProductCategory), the getlist
// operation
// (https://www.ticketmatic.com/docs/api/settings/productcategories/getlist) and
// the product categories endpoint
// (https://www.ticketmatic.com/docs/api/settings/productcategories).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/ProductCategoryQuery).
type ProductCategoryQuery struct {
	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`

	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`
}

// Product instancevalue exception
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/ProductInstanceException).
type ProductInstanceException struct {
	// Properties for which this exception is valid
	Properties map[string][]string `json:"properties,omitempty"`

	// Value for this exception
	Value *ProductInstanceValue `json:"value,omitempty"`
}

// Product Instance Pricetype Value
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/ProductInstancePricetypeValue).
type ProductInstancePricetypeValue struct {
	// Pricetype id
	Id int64 `json:"id"`

	// Min amount from which the pricetype will be applied
	From int64 `json:"from"`
}

// Product Instance Value
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/ProductInstanceValue).
type ProductInstanceValue struct {
	// Price
	Price float64 `json:"price"`

	// Set of pricetype values (used in optionbundle products)
	Pricetypes []*ProductInstancePricetypeValue `json:"pricetypes"`

	// Set of tickettypeprices (used in fixedbundle products)
	Tickettypeprices []int64 `json:"tickettypeprices"`

	// Set of tickettypes (used in optionbundle products)
	Tickettypes []int64 `json:"tickettypes"`

	// Voucher
	Voucher *ProductVoucherValue `json:"voucher,omitempty"`
}

// Product instancevalues
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/ProductInstancevalues).
type ProductInstancevalues struct {
	// Default values
	Default *ProductInstanceValue `json:"default,omitempty"`

	// Exceptions on the default values
	Exceptions []*ProductInstanceException `json:"exceptions"`
}

// Product property
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/ProductProperty).
type ProductProperty struct {
	// Name
	Name string `json:"name"`

	// Description
	Description string `json:"description"`

	// Key
	Key string `json:"key"`

	// Values
	Values []*KeyValueItem `json:"values"`
}

// Set of parameters used to filter products.
//
// More info: see product (https://www.ticketmatic.com/docs/api/types/Product), the
// getlist operation
// (https://www.ticketmatic.com/docs/api/settings/products/getlist) and the
// products endpoint (https://www.ticketmatic.com/docs/api/settings/products).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/ProductQuery).
type ProductQuery struct {
	// Only return items with the given typeid.
	Typeid int64 `json:"typeid,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`

	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`
}

// Product Voucher Value
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/ProductVoucherValue).
type ProductVoucherValue struct {
	// Amount (only used for vouchers of type Payment)
	Amount float64 `json:"amount"`

	// Voucher id
	Voucherid int64 `json:"voucherid"`
}

// Required data for creating a query on the public data model.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/QueryRequest).
type QueryRequest struct {
	// Optional limit for the result. Default 100
	//
	// Note: Ignored when exporting a query
	Limit int64 `json:"limit"`

	// Optional offset for the result. Default 0
	//
	// Note: Ignored when exporting a query
	Offset int64 `json:"offset"`

	// Actual query to execute
	Query string `json:"query"`
}

// Result of a query on the public data model.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/QueryResult).
type QueryResult struct {
	// The number of rows in the result
	Nbrofresults int64 `json:"nbrofresults"`

	// The actual resulting rows
	Results []map[string]interface{} `json:"results"`
}

// Rate limiting status. See rate limiting
// (https://www.ticketmatic.com/docs/api/ratelimiting) for more details on rate
// limiting.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/QueueStatus).
type QueueStatus struct {
	// Queueing ID, used to request status updates
	Id string `json:"id"`

	// The ID of the newly created order. Only returned when a throttled "create order"
	// call has finished queueing.
	Orderid int64 `json:"orderid"`

	// Number of people waiting ahead. Might not be returned when the queue hasn't
	// started yet.
	Ahead int64 `json:"ahead"`

	// Number of milliseconds to wait before requesting a new status update
	Backoff int64 `json:"backoff"`

	// Further instructions on how to handle this error
	Description string `json:"description"`

	// Optional message shown to waiting customers
	Message string `json:"message"`

	// Status code: 1: wait more, 2: ready to proceed
	Progress int64 `json:"progress"`

	// Whether the queue has started
	Started bool `json:"started"`

	// When the queue will start
	Starttime Time `json:"starttime"`
}

// A single relation type.
//
// More info: see the get operation
// (https://www.ticketmatic.com/docs/api/settings/system/relationtypes/get) and the
// relation types endpoint
// (https://www.ticketmatic.com/docs/api/settings/system/relationtypes).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/RelationType).
type RelationType struct {
	// Unique ID
	//
	// Note: Ignored when creating a new relation type.
	//
	// Note: Ignored when updating an existing relation type.
	Id int64 `json:"id"`

	// Name of the relation type
	Name string `json:"name"`

	// ID of the parent relation type.
	Parentid int64 `json:"parentid"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new relation type.
	//
	// Note: Ignored when updating an existing relation type.
	Isarchived bool `json:"isarchived"`

	// Created timestamp
	//
	// Note: Ignored when creating a new relation type.
	//
	// Note: Ignored when updating an existing relation type.
	Createdts Time `json:"createdts"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new relation type.
	//
	// Note: Ignored when updating an existing relation type.
	Lastupdatets Time `json:"lastupdatets"`
}

// Set of parameters used to filter relation types.
//
// More info: see relation type
// (https://www.ticketmatic.com/docs/api/types/RelationType), the getlist operation
// (https://www.ticketmatic.com/docs/api/settings/system/relationtypes/getlist) and
// the relation types endpoint
// (https://www.ticketmatic.com/docs/api/settings/system/relationtypes).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/RelationTypeQuery).
type RelationTypeQuery struct {
	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`

	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`
}

// A single report.
//
// More info: see the get operation
// (https://www.ticketmatic.com/docs/api/settings/system/reports/get) and the
// reports endpoint (https://www.ticketmatic.com/docs/api/settings/system/reports).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/Report).
type Report struct {
	// Unique ID
	//
	// Note: Ignored when creating a new report.
	//
	// Note: Ignored when updating an existing report.
	Id int64 `json:"id"`

	// Name of the report
	Name string `json:"name"`

	// The actual report definition, see reports
	// (https://www.ticketmatic.com/docs/tickets/reporting_and_followup/reporting/) for
	// more information.
	//
	// Note: Not set when retrieving a list of reports.
	Content [][]map[string]interface{} `json:"content"`

	// Reports can be generated as pdf or excel file. This field defines the default
	// format. Possible values are 'pdf' or 'excel'
	Defaultformat string `json:"defaultformat"`

	// Description of the report
	Description string `json:"description"`

	// List of email recipients that should receive the report in bcc, separated by ;
	Emailbcc string `json:"emailbcc"`

	// List of email recipients that should receive the report in cc, separated by ;
	Emailcc string `json:"emailcc"`

	// List of email recipients that should receive the report, separated by ;
	Emailrecipients string `json:"emailrecipients"`

	// Indicates if this report is scheduled to be sent by mail at a certain interval
	Emailschedule bool `json:"emailschedule"`

	// Day of the month the report will be sent.
	Emailscheduledayofmonth int64 `json:"emailscheduledayofmonth"`

	// Day of the week the report will be sent. 1 = monday -> 7 = sunday
	Emailscheduledayofweek int64 `json:"emailscheduledayofweek"`

	// Hour of the day the report will be sent
	Emailschedulehourofday int64 `json:"emailschedulehourofday"`

	// Key-value array of options. Can contain: pdfpagesize, excelpagewidth,
	// excelscaling, usesystemfont
	//
	// Note: Not set when retrieving a list of reports.
	Options *ReportOptions `json:"options,omitempty"`

	// The report type defines the UI and parameters that are used when generating the
	// report
	Reporttypeid int64 `json:"reporttypeid"`

	// A list of subtitles for the report
	//
	// Note: Not set when retrieving a list of reports.
	Subtitles []string `json:"subtitles"`

	// A map of language codes to gettext .po files
	// (http://en.wikipedia.org/wiki/Gettext).
	//
	// Note: Not set when retrieving a list of reports.
	Translations map[string]string `json:"translations,omitempty"`

	// Indicates where the report is being used. Possible values: 17001 (Sales), 17002
	// (External sales), 17003 (Hidden)
	Usagetypeid int64 `json:"usagetypeid"`

	// Created timestamp
	//
	// Note: Ignored when creating a new report.
	//
	// Note: Ignored when updating an existing report.
	Createdts Time `json:"createdts"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new report.
	//
	// Note: Ignored when updating an existing report.
	Lastupdatets Time `json:"lastupdatets"`
}

// Report Options
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/ReportOptions).
type ReportOptions struct {
	// The pagesize for the report when exported as Excel.
	Excelpagewidth int64 `json:"excelpagewidth"`

	// Excel-specific option for scaling the width
	Excelscaling float64 `json:"excelscaling"`

	// The pagesize for the report: A4 landscape, Letter landscape, ...
	Pdfpagesize string `json:"pdfpagesize"`

	// Indicates if a system font should be used.
	Usesystemfont bool `json:"usesystemfont"`
}

// Set of parameters used to filter reports.
//
// More info: see report (https://www.ticketmatic.com/docs/api/types/Report), the
// getlist operation
// (https://www.ticketmatic.com/docs/api/settings/system/reports/getlist) and the
// reports endpoint (https://www.ticketmatic.com/docs/api/settings/system/reports).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/ReportQuery).
type ReportQuery struct {
	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`
}

// A single sales channel.
//
// More info: see the get operation
// (https://www.ticketmatic.com/docs/api/settings/ticketsales/saleschannels/get)
// and the sales channels endpoint
// (https://www.ticketmatic.com/docs/api/settings/ticketsales/saleschannels).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/SalesChannel).
type SalesChannel struct {
	// Unique ID
	//
	// Note: Ignored when creating a new sales channel.
	//
	// Note: Ignored when updating an existing sales channel.
	Id int64 `json:"id"`

	// The type of this sales channel, defines where this sales channel will be used.
	// The available values for this field can be found on the sales channel overview
	// (https://www.ticketmatic.com/docs/api/settings/ticketsales/saleschannels) page.
	Typeid int64 `json:"typeid"`

	// Name of the sales channel
	Name string `json:"name"`

	// The ID of the order mail template to use for sending confirmations. Can be 0 to
	// indicate that no mail should be sent
	OrdermailtemplateidConfirmation int64 `json:"ordermailtemplateid_confirmation"`

	// Always send the confirmation, regardless of the payment method configuration
	OrdermailtemplateidConfirmationSendalways bool `json:"ordermailtemplateid_confirmation_sendalways"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new sales channel.
	//
	// Note: Ignored when updating an existing sales channel.
	Isarchived bool `json:"isarchived"`

	// Created timestamp
	//
	// Note: Ignored when creating a new sales channel.
	//
	// Note: Ignored when updating an existing sales channel.
	Createdts Time `json:"createdts"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new sales channel.
	//
	// Note: Ignored when updating an existing sales channel.
	Lastupdatets Time `json:"lastupdatets"`
}

// Set of parameters used to filter sales channels.
//
// More info: see sales channel
// (https://www.ticketmatic.com/docs/api/types/SalesChannel), the getlist operation
// (https://www.ticketmatic.com/docs/api/settings/ticketsales/saleschannels/getlist)
// and the sales channels endpoint
// (https://www.ticketmatic.com/docs/api/settings/ticketsales/saleschannels).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/SalesChannelQuery).
type SalesChannelQuery struct {
	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`

	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`
}

// A single seat rank.
//
// More info: see the get operation
// (https://www.ticketmatic.com/docs/api/settings/seatingplans/seatranks/get) and
// the seat ranks endpoint
// (https://www.ticketmatic.com/docs/api/settings/seatingplans/seatranks).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/SeatRank).
type SeatRank struct {
	// Unique ID
	//
	// Note: Ignored when creating a new seat rank.
	//
	// Note: Ignored when updating an existing seat rank.
	Id int64 `json:"id"`

	// Name for the seat rank
	Name string `json:"name"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new seat rank.
	//
	// Note: Ignored when updating an existing seat rank.
	Isarchived bool `json:"isarchived"`

	// Created timestamp
	//
	// Note: Ignored when creating a new seat rank.
	//
	// Note: Ignored when updating an existing seat rank.
	Createdts Time `json:"createdts"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new seat rank.
	//
	// Note: Ignored when updating an existing seat rank.
	Lastupdatets Time `json:"lastupdatets"`
}

// Set of parameters used to filter seat ranks.
//
// More info: see seat rank (https://www.ticketmatic.com/docs/api/types/SeatRank),
// the getlist operation
// (https://www.ticketmatic.com/docs/api/settings/seatingplans/seatranks/getlist)
// and the seat ranks endpoint
// (https://www.ticketmatic.com/docs/api/settings/seatingplans/seatranks).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/SeatRankQuery).
type SeatRankQuery struct {
	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`

	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`
}

// A single seating plan.
//
// More info: see the get operation
// (https://www.ticketmatic.com/docs/api/settings/seatingplans/seatingplans/get)
// and the seating plans endpoint
// (https://www.ticketmatic.com/docs/api/settings/seatingplans/seatingplans).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/SeatingPlan).
type SeatingPlan struct {
	// Unique ID
	//
	// Note: Ignored when creating a new seating plan.
	//
	// Note: Ignored when updating an existing seating plan.
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`

	// Translations for the seat description templates
	//
	// Note: Not set when retrieving a list of seating plans.
	Translations map[string]string `json:"translations,omitempty"`

	// When true: treat as a multi-zoned seatingplan
	Useszones bool `json:"useszones"`

	// IDs of the seat zones defined
	//
	// Note: Ignored when creating a new seating plan.
	//
	// Note: Ignored when updating an existing seating plan.
	//
	// Note: Not set when retrieving a list of seating plans.
	Zones []int64 `json:"zones"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new seating plan.
	//
	// Note: Ignored when updating an existing seating plan.
	Isarchived bool `json:"isarchived"`

	// Created timestamp
	//
	// Note: Ignored when creating a new seating plan.
	//
	// Note: Ignored when updating an existing seating plan.
	Createdts Time `json:"createdts"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new seating plan.
	//
	// Note: Ignored when updating an existing seating plan.
	Lastupdatets Time `json:"lastupdatets"`
}

// Set of parameters used to filter seating plans.
//
// More info: see seating plan
// (https://www.ticketmatic.com/docs/api/types/SeatingPlan), the getlist operation
// (https://www.ticketmatic.com/docs/api/settings/seatingplans/seatingplans/getlist)
// and the seating plans endpoint
// (https://www.ticketmatic.com/docs/api/settings/seatingplans/seatingplans).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/SeatingPlanQuery).
type SeatingPlanQuery struct {
	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`

	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`
}

// A newly created communication
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/SubscriberCommunication).
type SubscriberCommunication struct {
	// Name of the communication
	Name string `json:"name,omitempty"`

	// E-mail addresses to which the communication has been sent
	Addresses []string `json:"addresses"`

	// Optional description of the communication
	Remark string `json:"remark,omitempty"`

	// Timestamp for the communication
	Ts Time `json:"ts,omitempty"`
}

// A subscriber record to sync state back to Ticketmatic
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/SubscriberSync).
type SubscriberSync struct {
	// Subscriber e-mail
	Email string `json:"email,omitempty"`

	// Subscriber first name
	Firstname string `json:"firstname,omitempty"`

	// Subscriber last name
	Lastname string `json:"lastname,omitempty"`

	// Previous value of the email field, to indicate a changed e-mail address.
	//
	// Used to find the correct contact. The normal email field will be used when this
	// field is ommitted or empty.
	Oldemail string `json:"oldemail,omitempty"`

	// Whether or not the subscriber is still subscribed
	Subscribed bool `json:"subscribed,omitempty"`
}

// A single ticket fee.
//
// More info: see the get operation
// (https://www.ticketmatic.com/docs/api/settings/pricing/ticketfees/get) and the
// ticket fees endpoint
// (https://www.ticketmatic.com/docs/api/settings/pricing/ticketfees).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/TicketFee).
type TicketFee struct {
	// Unique ID
	//
	// Note: Ignored when creating a new ticket fee.
	//
	// Note: Ignored when updating an existing ticket fee.
	Id int64 `json:"id"`

	// Name for the ticket fee scheme
	Name string `json:"name"`

	// Definition of the rules that define when the ticket fee will be applied
	//
	// Note: Not set when retrieving a list of ticket fees.
	Rules *TicketfeeRules `json:"rules,omitempty"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new ticket fee.
	//
	// Note: Ignored when updating an existing ticket fee.
	Isarchived bool `json:"isarchived"`

	// Created timestamp
	//
	// Note: Ignored when creating a new ticket fee.
	//
	// Note: Ignored when updating an existing ticket fee.
	Createdts Time `json:"createdts"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new ticket fee.
	//
	// Note: Ignored when updating an existing ticket fee.
	Lastupdatets Time `json:"lastupdatets"`
}

// Set of parameters used to filter ticket fees.
//
// More info: see ticket fee
// (https://www.ticketmatic.com/docs/api/types/TicketFee), the getlist operation
// (https://www.ticketmatic.com/docs/api/settings/pricing/ticketfees/getlist) and
// the ticket fees endpoint
// (https://www.ticketmatic.com/docs/api/settings/pricing/ticketfees).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/TicketFeeQuery).
type TicketFeeQuery struct {
	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`

	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`
}

// A single ticket layout.
//
// More info: see the get operation
// (https://www.ticketmatic.com/docs/api/settings/communicationanddesign/ticketlayouts/get)
// and the ticket layouts endpoint
// (https://www.ticketmatic.com/docs/api/settings/communicationanddesign/ticketlayouts).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/TicketLayout).
type TicketLayout struct {
	// Unique ID
	//
	// Note: Ignored when creating a new ticket layout.
	//
	// Note: Ignored when updating an existing ticket layout.
	Id int64 `json:"id"`

	// Type ID
	//
	// Note: Ignored when updating an existing ticket layout.
	Typeid int64 `json:"typeid"`

	// Name for the ticket layout
	Name string `json:"name"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new ticket layout.
	//
	// Note: Ignored when updating an existing ticket layout.
	Isarchived bool `json:"isarchived"`

	// Created timestamp
	//
	// Note: Ignored when creating a new ticket layout.
	//
	// Note: Ignored when updating an existing ticket layout.
	Createdts Time `json:"createdts"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new ticket layout.
	//
	// Note: Ignored when updating an existing ticket layout.
	Lastupdatets Time `json:"lastupdatets"`
}

// Set of parameters used to filter ticket layouts.
//
// More info: see ticket layout
// (https://www.ticketmatic.com/docs/api/types/TicketLayout), the getlist operation
// (https://www.ticketmatic.com/docs/api/settings/communicationanddesign/ticketlayouts/getlist)
// and the ticket layouts endpoint
// (https://www.ticketmatic.com/docs/api/settings/communicationanddesign/ticketlayouts).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/TicketLayoutQuery).
type TicketLayoutQuery struct {
	// Only return items with the given typeid.
	Typeid int64 `json:"typeid,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`

	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`
}

// A single ticket layout template.
//
// More info: see the get operation
// (https://www.ticketmatic.com/docs/api/settings/communicationanddesign/ticketlayouttemplates/get)
// and the ticket layout templates endpoint
// (https://www.ticketmatic.com/docs/api/settings/communicationanddesign/ticketlayouttemplates).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/TicketLayoutTemplate).
type TicketLayoutTemplate struct {
	// Unique ID
	//
	// Note: Ignored when creating a new ticket layout template.
	//
	// Note: Ignored when updating an existing ticket layout template.
	Id int64 `json:"id"`

	// Type ID
	//
	// Note: Ignored when updating an existing ticket layout template.
	Typeid int64 `json:"typeid"`

	// Name for the ticket layout template
	Name string `json:"name"`

	// Css classes for the ticket layout template
	//
	// Note: Not set when retrieving a list of ticket layout templates.
	Css string `json:"css"`

	// Deliveryscenario's for which this ticket layout template will be used
	Deliveryscenarios []int64 `json:"deliveryscenarios"`

	// HTML template containing the definition for the ticket layout template
	//
	// Note: Not set when retrieving a list of ticket layout templates.
	Htmltemplate string `json:"htmltemplate"`

	// Translations for the ticket layout template
	//
	// Note: Not set when retrieving a list of ticket layout templates.
	Translations map[string]string `json:"translations,omitempty"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new ticket layout template.
	//
	// Note: Ignored when updating an existing ticket layout template.
	Isarchived bool `json:"isarchived"`

	// Created timestamp
	//
	// Note: Ignored when creating a new ticket layout template.
	//
	// Note: Ignored when updating an existing ticket layout template.
	Createdts Time `json:"createdts"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new ticket layout template.
	//
	// Note: Ignored when updating an existing ticket layout template.
	Lastupdatets Time `json:"lastupdatets"`
}

// Set of parameters used to filter ticket layout templates.
//
// More info: see ticket layout template
// (https://www.ticketmatic.com/docs/api/types/TicketLayoutTemplate), the getlist
// operation
// (https://www.ticketmatic.com/docs/api/settings/communicationanddesign/ticketlayouttemplates/getlist)
// and the ticket layout templates endpoint
// (https://www.ticketmatic.com/docs/api/settings/communicationanddesign/ticketlayouttemplates).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/TicketLayoutTemplateQuery).
type TicketLayoutTemplateQuery struct {
	// Only return items with the given typeid.
	Typeid int64 `json:"typeid,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`

	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`
}

// An exception to the default rule for a specific pricetype and a set of
// saleschannels.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/TicketfeeException).
type TicketfeeException struct {
	// The pricetype for which this exception is active.
	Pricetypeid int64 `json:"pricetypeid"`

	// The set of rules (one for each saleschannel).
	Saleschannels []*TicketfeeSaleschannelRule `json:"saleschannels"`
}

// Defines which fees are active for specific price types and sales channels. It's
// possible to define a fixed fee and a percentage based fee. The default rule (if
// none is specified for a specific sales channel) is always a fixed fee of 0.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/TicketfeeRules).
type TicketfeeRules struct {
	// The default ticket fee rule, one rule for each saleschannel.
	Default []*TicketfeeSaleschannelRule `json:"default"`

	// An array of exception rules for specific pricetypes.
	Exceptions []*TicketfeeException `json:"exceptions"`
}

// This is a rule for a specific saleschannel that indicates the fee based on a
// fixed amount or a percentage.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/TicketfeeSaleschannelRule).
type TicketfeeSaleschannelRule struct {
	// The saleschannel for which this rule is active.
	Saleschannelid int64 `json:"saleschannelid"`

	// The status sets the type of rule. Possible values:
	//
	// * fixedfee: A fixed ticket fee.
	//
	// * percentagefee: A fee thats a percentage of the ticket.
	Status string `json:"status"`

	// The value of this ticket fee. Can be an absolute amount (fixedfee) or a
	// percentage (percentagefee). In both cases only provide a decimal.
	Value float64 `json:"value"`
}

// Info for requesting a e-mail delivery for an order
// (https://www.ticketmatic.com/docs/api/types/Order).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/TicketsEmaildeliveryRequest).
type TicketsEmaildeliveryRequest struct {
	// Template id
	Templateid int64 `json:"templateid"`
}

// Info for requesting a PDF ticket for one or more tickets or vouchercodes in an
// order (https://www.ticketmatic.com/docs/api/types/Order).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/TicketsPdfRequest).
type TicketsPdfRequest struct {
	// Ticketids
	Tickets []int64 `json:"tickets"`

	// Vouchercodeids
	Vouchercodes []int64 `json:"vouchercodes"`
}

// Required data for requesting the ticketsprocessedstatistics.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/TicketsprocessedRequest).
type TicketsprocessedRequest struct {
	// End date of the period
	Endts string `json:"endts"`

	// How the results are grouped. Values can be 'day' or 'month'
	Groupby string `json:"groupby"`

	// Start date of the period
	Startts string `json:"startts"`
}

// Statistics on the number of tickets processed in a certain period.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/TicketsprocessedStatistics).
type TicketsprocessedStatistics struct {
	// The number of tickets processed
	Processed int64 `json:"processed"`

	// The number of tickets sold online
	Soldonline int64 `json:"soldonline"`

	// Start of the period
	Ts Time `json:"ts"`
}

// A timestamp returned by the diagnostic /time call
// (https://www.ticketmatic.com/docs/api/diagnostics/time).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/Timestamp).
type Timestamp struct {
	// Current system time
	Systemtime Time `json:"systemtime,omitempty"`
}

// Used to update an order. Each of the fields is optional. Omitting a field will
// leave it unchanged.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/UpdateOrder).
type UpdateOrder struct {
	// New customer ID
	Customerid int64 `json:"customerid,omitempty"`

	// Change custom field values
	Customfields map[string]interface{} `json:"customfields,omitempty"`

	// Delivery address
	Deliveryaddress *Address `json:"deliveryaddress,omitempty"`

	// New delivery scenario ID
	Deliveryscenarioid int64 `json:"deliveryscenarioid,omitempty"`

	// Expiry timestamp, as string in ISO 8601 format. Cannot be in the past.
	Expiryts string `json:"expiryts,omitempty"`

	// New payment scenario ID
	Paymentscenarioid int64 `json:"paymentscenarioid,omitempty"`

	// Rappel timestamp, as string in ISO 8601 format. Cannot be in the past.
	Rappelts string `json:"rappelts,omitempty"`
}

// Individual products can be updated. Per call you can specify any number of
// product IDs and one operation.
//
// Each operation accepts different parameters, dependent on the operation type:
//
// * Set product holders: an array of ticket holder IDs (see Contact
// (https://www.ticketmatic.com/docs/api/types/Contact)), one for each product
// (productholderids). *
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/UpdateProducts).
type UpdateProducts struct {
	// Operation to execute.
	//
	// Supported values:
	//
	// * setproductholders
	Operation string `json:"operation"`

	// Operation parameters
	Params map[string]interface{} `json:"params,omitempty"`

	// Product IDs
	Products []int64 `json:"products"`
}

// Individual tickets can be updated. Per call you can specify any number of ticket
// IDs and one operation.
//
// Each operation accepts different parameters, dependent on the operation type:
//
// * Set ticket holders: an array of ticket holder IDs (see Contact
// (https://www.ticketmatic.com/docs/api/types/Contact)), one for each ticket
// (ticketholderids).
//
// * Update price type: an array of ticket price type IDs (as can be found in the
// Event pricing (https://www.ticketmatic.com/docs/api/types/Event)), one for each
// ticket (tickettypepriceids).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/UpdateTickets).
type UpdateTickets struct {
	// Operation to execute.
	//
	// Supported values:
	//
	// * setticketholders
	//
	// * updatepricetype
	Operation string `json:"operation"`

	// Operation parameters
	Params map[string]interface{} `json:"params,omitempty"`

	// Ticket IDs
	Tickets []int64 `json:"tickets"`
}

// Url.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/Url).
type Url struct {
	// Url.
	Url string `json:"url"`
}

// A single view.
//
// More info: see the get operation
// (https://www.ticketmatic.com/docs/api/settings/system/views/get) and the views
// endpoint (https://www.ticketmatic.com/docs/api/settings/system/views).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/View).
type View struct {
	// Unique ID
	//
	// Note: Ignored when creating a new view.
	//
	// Note: Ignored when updating an existing view.
	Id int64 `json:"id"`

	// Type ID
	//
	// Note: Ignored when updating an existing view.
	Typeid int64 `json:"typeid"`

	// Name of the view
	Name string `json:"name"`

	// List of field definitions that are part of this view.
	Columns []*ViewColumn `json:"columns"`

	// The field definitions to order the results on.
	Orderby int64 `json:"orderby"`

	// Indicates whether the results should be ordered ascending or descending.
	OrderbyAsc bool `json:"orderby_asc"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new view.
	//
	// Note: Ignored when updating an existing view.
	Isarchived bool `json:"isarchived"`

	// Created timestamp
	//
	// Note: Ignored when creating a new view.
	//
	// Note: Ignored when updating an existing view.
	Createdts Time `json:"createdts"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new view.
	//
	// Note: Ignored when updating an existing view.
	Lastupdatets Time `json:"lastupdatets"`
}

// View column for a view.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/ViewColumn).
type ViewColumn struct {
	// ID of the field definition for this column.
	Id int64 `json:"id"`
}

// Set of parameters used to filter views.
//
// More info: see view (https://www.ticketmatic.com/docs/api/types/View), the
// getlist operation
// (https://www.ticketmatic.com/docs/api/settings/system/views/getlist) and the
// views endpoint (https://www.ticketmatic.com/docs/api/settings/system/views).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/ViewQuery).
type ViewQuery struct {
	// Only return items with the given typeid.
	Typeid int64 `json:"typeid"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`

	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`
}

// A single voucher.
//
// More info: see the get operation
// (https://www.ticketmatic.com/docs/api/settings/vouchers/get) and the vouchers
// endpoint (https://www.ticketmatic.com/docs/api/settings/vouchers).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/Voucher).
type Voucher struct {
	// Unique ID
	//
	// Note: Ignored when creating a new voucher.
	//
	// Note: Ignored when updating an existing voucher.
	Id int64 `json:"id"`

	// Type ID
	//
	// Note: Ignored when updating an existing voucher.
	Typeid int64 `json:"typeid"`

	// Name of the voucher
	Name string `json:"name"`

	// Format for the codes for the voucher
	Codeformatid int64 `json:"codeformatid"`

	// Description of the voucher
	Description string `json:"description"`

	// The number of codes that were created for this voucher.
	//
	// Note: Ignored when creating a new voucher.
	//
	// Note: Ignored when updating an existing voucher.
	Nbrofcodes int64 `json:"nbrofcodes"`

	// A validation script that is used for vouchers of type order. For each order with
	// a voucher of this type attached, the script will be run to validate the contents
	Ordervalidationscript string `json:"ordervalidationscript"`

	// Paymentmethod to use when creating payments for vouchers of type payment.
	Paymentmethodid int64 `json:"paymentmethodid"`

	// Ticketlayout to be used for this voucher.
	Ticketlayoutid int64 `json:"ticketlayoutid"`

	// Definition of the validity of this voucher. Depends on the typeid.
	//
	// Note: Not set when retrieving a list of vouchers.
	Validity *VoucherValidity `json:"validity,omitempty"`

	// Whether or not this item is archived
	//
	// Note: Ignored when creating a new voucher.
	//
	// Note: Ignored when updating an existing voucher.
	Isarchived bool `json:"isarchived"`

	// Created timestamp
	//
	// Note: Ignored when creating a new voucher.
	//
	// Note: Ignored when updating an existing voucher.
	Createdts Time `json:"createdts"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new voucher.
	//
	// Note: Ignored when updating an existing voucher.
	Lastupdatets Time `json:"lastupdatets"`
}

// Voucher code
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/VoucherCode).
type VoucherCode struct {
	// Code to use voucher
	Code string `json:"code"`

	// Expiry timestamp for this code
	//
	// Note: Only used when creating codes
	Expiryts Time `json:"expiryts,omitempty"`
}

// Set of parameters used to filter vouchers.
//
// More info: see voucher (https://www.ticketmatic.com/docs/api/types/Voucher), the
// getlist operation
// (https://www.ticketmatic.com/docs/api/settings/vouchers/getlist) and the
// vouchers endpoint (https://www.ticketmatic.com/docs/api/settings/vouchers).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/VoucherQuery).
type VoucherQuery struct {
	// Only return items with the given typeid.
	Typeid int64 `json:"typeid,omitempty"`

	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`

	// If this parameter is true, archived items will be returned as well.
	Includearchived bool `json:"includearchived,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`
}

// The definition of the validity of a voucher.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/VoucherValidity).
type VoucherValidity struct {
	// The fixed expiry date for a voucher
	ExpiryFixeddate Time `json:"expiry_fixeddate"`

	// The relative expiry date for a voucher: voucher code expires this number of
	// months after creation
	ExpiryMonthsaftercreation int64 `json:"expiry_monthsaftercreation"`

	// The max number of times the vouchercode can be used
	Maxusages int64 `json:"maxusages"`

	// The max number of times the vouchercode can be used for a single event
	Maxusagesperevent int64 `json:"maxusagesperevent"`
}

// A single web sales skin.
//
// More info: see the get operation
// (https://www.ticketmatic.com/docs/api/settings/communicationanddesign/webskins/get)
// and the web sales skins endpoint
// (https://www.ticketmatic.com/docs/api/settings/communicationanddesign/webskins).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/WebSalesSkin).
type WebSalesSkin struct {
	// Unique ID
	//
	// Note: Ignored when creating a new web sales skin.
	//
	// Note: Ignored when updating an existing web sales skin.
	Id int64 `json:"id"`

	// Name of the web sales skin
	Name string `json:"name"`

	// Skin configuration.
	//
	// See the WebSalesSkinConfiguration reference
	// (https://www.ticketmatic.com/docs/api/types/WebSalesSkinConfiguration) for an
	// overview of all possible options.
	//
	// Note: Not set when retrieving a list of web sales skins.
	Configuration *WebSalesSkinConfiguration `json:"configuration,omitempty"`

	// CSS style rules. Should always include the style import.
	//
	// Note: Not set when retrieving a list of web sales skins.
	Css string `json:"css"`

	// HTML template of the skin. See the web skin setup guide
	// (https://www.ticketmatic.com/docs/tickets/configure_ticket_sales/webskin) for
	// more information.
	//
	// Note: Not set when retrieving a list of web sales skins.
	Html string `json:"html"`

	// A map of language codes to gettext .po files
	// (http://en.wikipedia.org/wiki/Gettext). More info can be found on the web skin
	// overview
	// (https://www.ticketmatic.com/docs/api/settings/communicationanddesign/webskins)
	// page.
	//
	// Note: Not set when retrieving a list of web sales skins.
	Translations map[string]string `json:"translations,omitempty"`

	// Created timestamp
	//
	// Note: Ignored when creating a new web sales skin.
	//
	// Note: Ignored when updating an existing web sales skin.
	Createdts Time `json:"createdts"`

	// Last updated timestamp
	//
	// Note: Ignored when creating a new web sales skin.
	//
	// Note: Ignored when updating an existing web sales skin.
	Lastupdatets Time `json:"lastupdatets"`
}

// Configuration settings and parameters for a web sales skin
// (https://www.ticketmatic.com/docs/api/types/WebSalesSkin).
//
// Page titles
//
// The title field contains a template for the page title. The same variables as in
// the HTML of the skin itself can be used.
//
// Check the web sales skin setup guide
// (https://www.ticketmatic.com/docs/tickets/configure_ticket_sales/webskin) for
// more information.
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/WebSalesSkinConfiguration).
type WebSalesSkinConfiguration struct {
	// Asset path to favicon image.
	Favicon string `json:"favicon"`

	// Facebook app ID to use for Facebook authentication.
	//
	// The default Ticketmatic Facebook app will be used if you leave this field blank
	Fbappid string `json:"fbappid"`

	// Deprecated, use Google Tag Manager.
	Googleanalyticsid string `json:"googleanalyticsid"`

	// Google Tag Manager ID. Can be left blank.
	Googletagmanagerid string `json:"googletagmanagerid"`

	// Page title
	Title string `json:"title"`
}

// Set of parameters used to filter web sales skins.
//
// More info: see web sales skin
// (https://www.ticketmatic.com/docs/api/types/WebSalesSkin), the getlist operation
// (https://www.ticketmatic.com/docs/api/settings/communicationanddesign/webskins/getlist)
// and the web sales skins endpoint
// (https://www.ticketmatic.com/docs/api/settings/communicationanddesign/webskins).
//
// Help Center
//
// Full documentation can be found in the Ticketmatic Help Center
// (https://www.ticketmatic.com/docs/api/types/WebSalesSkinQuery).
type WebSalesSkinQuery struct {
	// Filter the returned items by specifying a query on the public datamodel that
	// returns the ids.
	Filter string `json:"filter,omitempty"`

	// All items that were updated since this timestamp will be returned. Timestamp
	// should be passed in YYYY-MM-DD hh:mm:ss format.
	Lastupdatesince Time `json:"lastupdatesince,omitempty"`
}
