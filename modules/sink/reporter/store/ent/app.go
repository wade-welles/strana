// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/blushft/strana/modules/sink/reporter/store/ent/app"
	"github.com/facebook/ent/dialect/sql"
)

// App is the model entity for the App schema.
type App struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Version holds the value of the "version" field.
	Version string `json:"version,omitempty"`
	// Build holds the value of the "build" field.
	Build string `json:"build,omitempty"`
	// Namespace holds the value of the "namespace" field.
	Namespace string `json:"namespace,omitempty"`
	// Properties holds the value of the "properties" field.
	Properties map[string]interface{} `json:"properties,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AppQuery when eager-loading is set.
	Edges AppEdges `json:"edges"`
}

// AppEdges holds the relations/edges for other nodes in the graph.
type AppEdges struct {
	// Events holds the value of the events edge.
	Events []*Event
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EventsOrErr returns the Events value or an error if the edge
// was not loaded in eager-loading.
func (e AppEdges) EventsOrErr() ([]*Event, error) {
	if e.loadedTypes[0] {
		return e.Events, nil
	}
	return nil, &NotLoadedError{edge: "events"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*App) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
		&sql.NullString{}, // version
		&sql.NullString{}, // build
		&sql.NullString{}, // namespace
		&[]byte{},         // properties
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the App fields.
func (a *App) assignValues(values ...interface{}) error {
	if m, n := len(values), len(app.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	a.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		a.Name = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field version", values[1])
	} else if value.Valid {
		a.Version = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field build", values[2])
	} else if value.Valid {
		a.Build = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field namespace", values[3])
	} else if value.Valid {
		a.Namespace = value.String
	}

	if value, ok := values[4].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field properties", values[4])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &a.Properties); err != nil {
			return fmt.Errorf("unmarshal field properties: %v", err)
		}
	}
	return nil
}

// QueryEvents queries the events edge of the App.
func (a *App) QueryEvents() *EventQuery {
	return (&AppClient{config: a.config}).QueryEvents(a)
}

// Update returns a builder for updating this App.
// Note that, you need to call App.Unwrap() before calling this method, if this App
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *App) Update() *AppUpdateOne {
	return (&AppClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (a *App) Unwrap() *App {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: App is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *App) String() string {
	var builder strings.Builder
	builder.WriteString("App(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", name=")
	builder.WriteString(a.Name)
	builder.WriteString(", version=")
	builder.WriteString(a.Version)
	builder.WriteString(", build=")
	builder.WriteString(a.Build)
	builder.WriteString(", namespace=")
	builder.WriteString(a.Namespace)
	builder.WriteString(", properties=")
	builder.WriteString(fmt.Sprintf("%v", a.Properties))
	builder.WriteByte(')')
	return builder.String()
}

// Apps is a parsable slice of App.
type Apps []*App

func (a Apps) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
