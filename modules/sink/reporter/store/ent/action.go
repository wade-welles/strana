// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/blushft/strana/modules/sink/reporter/store/ent/action"
	"github.com/blushft/strana/modules/sink/reporter/store/ent/event"
	"github.com/facebook/ent/dialect/sql"
	"github.com/google/uuid"
)

// Action is the model entity for the Action schema.
type Action struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Action holds the value of the "action" field.
	Action string `json:"action,omitempty"`
	// ActionLabel holds the value of the "action_label" field.
	ActionLabel string `json:"action_label,omitempty"`
	// Property holds the value of the "property" field.
	Property string `json:"property,omitempty"`
	// Value holds the value of the "value" field.
	Value []byte `json:"value,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ActionQuery when eager-loading is set.
	Edges        ActionEdges `json:"edges"`
	event_action *uuid.UUID
}

// ActionEdges holds the relations/edges for other nodes in the graph.
type ActionEdges struct {
	// Event holds the value of the event edge.
	Event *Event
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EventOrErr returns the Event value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ActionEdges) EventOrErr() (*Event, error) {
	if e.loadedTypes[0] {
		if e.Event == nil {
			// The edge event was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: event.Label}
		}
		return e.Event, nil
	}
	return nil, &NotLoadedError{edge: "event"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Action) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // action
		&sql.NullString{}, // action_label
		&sql.NullString{}, // property
		&[]byte{},         // value
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Action) fkValues() []interface{} {
	return []interface{}{
		&uuid.UUID{}, // event_action
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Action fields.
func (a *Action) assignValues(values ...interface{}) error {
	if m, n := len(values), len(action.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	a.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field action", values[0])
	} else if value.Valid {
		a.Action = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field action_label", values[1])
	} else if value.Valid {
		a.ActionLabel = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field property", values[2])
	} else if value.Valid {
		a.Property = value.String
	}
	if value, ok := values[3].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field value", values[3])
	} else if value != nil {
		a.Value = *value
	}
	values = values[4:]
	if len(values) == len(action.ForeignKeys) {
		if value, ok := values[0].(*uuid.UUID); !ok {
			return fmt.Errorf("unexpected type %T for field event_action", values[0])
		} else if value != nil {
			a.event_action = value
		}
	}
	return nil
}

// QueryEvent queries the event edge of the Action.
func (a *Action) QueryEvent() *EventQuery {
	return (&ActionClient{config: a.config}).QueryEvent(a)
}

// Update returns a builder for updating this Action.
// Note that, you need to call Action.Unwrap() before calling this method, if this Action
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Action) Update() *ActionUpdateOne {
	return (&ActionClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (a *Action) Unwrap() *Action {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Action is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Action) String() string {
	var builder strings.Builder
	builder.WriteString("Action(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", action=")
	builder.WriteString(a.Action)
	builder.WriteString(", action_label=")
	builder.WriteString(a.ActionLabel)
	builder.WriteString(", property=")
	builder.WriteString(a.Property)
	builder.WriteString(", value=")
	builder.WriteString(fmt.Sprintf("%v", a.Value))
	builder.WriteByte(')')
	return builder.String()
}

// Actions is a parsable slice of Action.
type Actions []*Action

func (a Actions) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
