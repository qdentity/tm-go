package tools

import (
	"io"

	"github.com/ticketmatic/tm-go/ticketmatic"
)

type QueryStream struct {
	stream *ticketmatic.Stream
}

func (s *QueryStream) Next() (map[string]interface{}, error) {
	var obj map[string]interface{}
	err := s.stream.Next(&obj)
	if err != nil {
		if err == io.EOF {
			err = nil
		}
		return nil, err
	}
	return obj, nil
}

func (s *QueryStream) Close() {
	s.stream.Close()

}

// Execute a query on the public data model
//
// Use this method to execute random (read-only) queries on the public data model.
// Note that this is not meant for long-running queries or for returning large
// resultsets. If the query executes too long or uses too much memory, an exception
// will be returned.
func Queries(client *ticketmatic.Client, data *ticketmatic.QueryRequest) (*ticketmatic.QueryResult, error) {
	r := client.NewRequest("POST", "/{accountname}/tools/queries", "")
	r.Body(data)

	var obj *ticketmatic.QueryResult
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// Export a query on the public data model
//
// Executes a query against the public data model and exports the results as a
// stream of JSON lines (http://jsonlines.org/): each line contains a JSON object
// which holds one row of the query result.
func Export(client *ticketmatic.Client, data *ticketmatic.QueryRequest) (*QueryStream, error) {
	r := client.NewRequest("POST", "/{accountname}/tools/queries/export", "")
	r.Body(data)

	stream, err := r.Stream()
	if err != nil {
		return nil, err
	}
	return &QueryStream{stream}, nil
}

// Get statistics on the tickets processed during a certain period
//
// Use this method to retrieve the statistics on the number of tickets processed
// and sold online during a certain period. The results can be grouped by day or
// month. These statistics are often used as basis for invoicing or reporting.
func Ticketsprocessedstatistics(client *ticketmatic.Client, params *ticketmatic.TicketsprocessedRequest) ([]*ticketmatic.TicketsprocessedStatistics, error) {
	r := client.NewRequest("GET", "/{accountname}/tools/ticketsprocessedstatistics", "")
	if params != nil {
		r.AddParameter("endts", params.Endts)
		r.AddParameter("groupby", params.Groupby)
		r.AddParameter("startts", params.Startts)
	}

	var obj []*ticketmatic.TicketsprocessedStatistics
	err := r.Run(&obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
