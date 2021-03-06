// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/blushft/strana/modules/sink/reporter/store/ent/device"
	"github.com/facebook/ent/dialect/sql"
)

// Device is the model entity for the Device schema.
type Device struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Manufacturer holds the value of the "manufacturer" field.
	Manufacturer string `json:"manufacturer,omitempty"`
	// Model holds the value of the "model" field.
	Model string `json:"model,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Version holds the value of the "version" field.
	Version string `json:"version,omitempty"`
	// Mobile holds the value of the "mobile" field.
	Mobile bool `json:"mobile,omitempty"`
	// Tablet holds the value of the "tablet" field.
	Tablet bool `json:"tablet,omitempty"`
	// Desktop holds the value of the "desktop" field.
	Desktop bool `json:"desktop,omitempty"`
	// Properties holds the value of the "properties" field.
	Properties map[string]interface{} `json:"properties,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DeviceQuery when eager-loading is set.
	Edges DeviceEdges `json:"edges"`
}

// DeviceEdges holds the relations/edges for other nodes in the graph.
type DeviceEdges struct {
	// Events holds the value of the events edge.
	Events []*Event
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EventsOrErr returns the Events value or an error if the edge
// was not loaded in eager-loading.
func (e DeviceEdges) EventsOrErr() ([]*Event, error) {
	if e.loadedTypes[0] {
		return e.Events, nil
	}
	return nil, &NotLoadedError{edge: "events"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Device) scanValues() []interface{} {
	return []interface{}{
		&sql.NullString{}, // id
		&sql.NullString{}, // manufacturer
		&sql.NullString{}, // model
		&sql.NullString{}, // name
		&sql.NullString{}, // type
		&sql.NullString{}, // version
		&sql.NullBool{},   // mobile
		&sql.NullBool{},   // tablet
		&sql.NullBool{},   // desktop
		&[]byte{},         // properties
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Device fields.
func (d *Device) assignValues(values ...interface{}) error {
	if m, n := len(values), len(device.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value.Valid {
		d.ID = value.String
	}
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field manufacturer", values[0])
	} else if value.Valid {
		d.Manufacturer = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field model", values[1])
	} else if value.Valid {
		d.Model = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[2])
	} else if value.Valid {
		d.Name = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field type", values[3])
	} else if value.Valid {
		d.Type = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field version", values[4])
	} else if value.Valid {
		d.Version = value.String
	}
	if value, ok := values[5].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field mobile", values[5])
	} else if value.Valid {
		d.Mobile = value.Bool
	}
	if value, ok := values[6].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field tablet", values[6])
	} else if value.Valid {
		d.Tablet = value.Bool
	}
	if value, ok := values[7].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field desktop", values[7])
	} else if value.Valid {
		d.Desktop = value.Bool
	}

	if value, ok := values[8].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field properties", values[8])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &d.Properties); err != nil {
			return fmt.Errorf("unmarshal field properties: %v", err)
		}
	}
	return nil
}

// QueryEvents queries the events edge of the Device.
func (d *Device) QueryEvents() *EventQuery {
	return (&DeviceClient{config: d.config}).QueryEvents(d)
}

// Update returns a builder for updating this Device.
// Note that, you need to call Device.Unwrap() before calling this method, if this Device
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Device) Update() *DeviceUpdateOne {
	return (&DeviceClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (d *Device) Unwrap() *Device {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Device is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Device) String() string {
	var builder strings.Builder
	builder.WriteString("Device(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteString(", manufacturer=")
	builder.WriteString(d.Manufacturer)
	builder.WriteString(", model=")
	builder.WriteString(d.Model)
	builder.WriteString(", name=")
	builder.WriteString(d.Name)
	builder.WriteString(", type=")
	builder.WriteString(d.Type)
	builder.WriteString(", version=")
	builder.WriteString(d.Version)
	builder.WriteString(", mobile=")
	builder.WriteString(fmt.Sprintf("%v", d.Mobile))
	builder.WriteString(", tablet=")
	builder.WriteString(fmt.Sprintf("%v", d.Tablet))
	builder.WriteString(", desktop=")
	builder.WriteString(fmt.Sprintf("%v", d.Desktop))
	builder.WriteString(", properties=")
	builder.WriteString(fmt.Sprintf("%v", d.Properties))
	builder.WriteByte(')')
	return builder.String()
}

// Devices is a parsable slice of Device.
type Devices []*Device

func (d Devices) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
