// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/blushft/strana/modules/sink/reporter/store/ent/campaign"
	"github.com/facebook/ent/dialect/sql"
)

// Campaign is the model entity for the Campaign schema.
type Campaign struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Source holds the value of the "source" field.
	Source string `json:"source,omitempty"`
	// Medium holds the value of the "medium" field.
	Medium string `json:"medium,omitempty"`
	// Term holds the value of the "term" field.
	Term string `json:"term,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CampaignQuery when eager-loading is set.
	Edges CampaignEdges `json:"edges"`
}

// CampaignEdges holds the relations/edges for other nodes in the graph.
type CampaignEdges struct {
	// Events holds the value of the events edge.
	Events []*Event
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EventsOrErr returns the Events value or an error if the edge
// was not loaded in eager-loading.
func (e CampaignEdges) EventsOrErr() ([]*Event, error) {
	if e.loadedTypes[0] {
		return e.Events, nil
	}
	return nil, &NotLoadedError{edge: "events"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Campaign) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
		&sql.NullString{}, // source
		&sql.NullString{}, // medium
		&sql.NullString{}, // term
		&sql.NullString{}, // content
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Campaign fields.
func (c *Campaign) assignValues(values ...interface{}) error {
	if m, n := len(values), len(campaign.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	c.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		c.Name = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field source", values[1])
	} else if value.Valid {
		c.Source = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field medium", values[2])
	} else if value.Valid {
		c.Medium = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field term", values[3])
	} else if value.Valid {
		c.Term = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field content", values[4])
	} else if value.Valid {
		c.Content = value.String
	}
	return nil
}

// QueryEvents queries the events edge of the Campaign.
func (c *Campaign) QueryEvents() *EventQuery {
	return (&CampaignClient{config: c.config}).QueryEvents(c)
}

// Update returns a builder for updating this Campaign.
// Note that, you need to call Campaign.Unwrap() before calling this method, if this Campaign
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Campaign) Update() *CampaignUpdateOne {
	return (&CampaignClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *Campaign) Unwrap() *Campaign {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Campaign is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Campaign) String() string {
	var builder strings.Builder
	builder.WriteString("Campaign(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", name=")
	builder.WriteString(c.Name)
	builder.WriteString(", source=")
	builder.WriteString(c.Source)
	builder.WriteString(", medium=")
	builder.WriteString(c.Medium)
	builder.WriteString(", term=")
	builder.WriteString(c.Term)
	builder.WriteString(", content=")
	builder.WriteString(c.Content)
	builder.WriteByte(')')
	return builder.String()
}

// Campaigns is a parsable slice of Campaign.
type Campaigns []*Campaign

func (c Campaigns) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
