// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/blushft/strana/modules/sink/reporter/store/ent/session"
	"github.com/facebook/ent/dialect/sql"
	"github.com/google/uuid"
)

// Session is the model entity for the Session schema.
type Session struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// NewUser holds the value of the "new_user" field.
	NewUser bool `json:"new_user,omitempty"`
	// IsUnique holds the value of the "is_unique" field.
	IsUnique bool `json:"is_unique,omitempty"`
	// IsBounce holds the value of the "is_bounce" field.
	IsBounce bool `json:"is_bounce,omitempty"`
	// IsFinished holds the value of the "is_finished" field.
	IsFinished bool `json:"is_finished,omitempty"`
	// Duration holds the value of the "duration" field.
	Duration int `json:"duration,omitempty"`
	// StartedAt holds the value of the "started_at" field.
	StartedAt time.Time `json:"started_at,omitempty"`
	// FinishedAt holds the value of the "finished_at" field.
	FinishedAt *time.Time `json:"finished_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SessionQuery when eager-loading is set.
	Edges SessionEdges `json:"edges"`
}

// SessionEdges holds the relations/edges for other nodes in the graph.
type SessionEdges struct {
	// Events holds the value of the events edge.
	Events []*Event
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EventsOrErr returns the Events value or an error if the edge
// was not loaded in eager-loading.
func (e SessionEdges) EventsOrErr() ([]*Event, error) {
	if e.loadedTypes[0] {
		return e.Events, nil
	}
	return nil, &NotLoadedError{edge: "events"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Session) scanValues() []interface{} {
	return []interface{}{
		&uuid.UUID{},     // id
		&sql.NullBool{},  // new_user
		&sql.NullBool{},  // is_unique
		&sql.NullBool{},  // is_bounce
		&sql.NullBool{},  // is_finished
		&sql.NullInt64{}, // duration
		&sql.NullTime{},  // started_at
		&sql.NullTime{},  // finished_at
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Session fields.
func (s *Session) assignValues(values ...interface{}) error {
	if m, n := len(values), len(session.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value != nil {
		s.ID = *value
	}
	values = values[1:]
	if value, ok := values[0].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field new_user", values[0])
	} else if value.Valid {
		s.NewUser = value.Bool
	}
	if value, ok := values[1].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field is_unique", values[1])
	} else if value.Valid {
		s.IsUnique = value.Bool
	}
	if value, ok := values[2].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field is_bounce", values[2])
	} else if value.Valid {
		s.IsBounce = value.Bool
	}
	if value, ok := values[3].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field is_finished", values[3])
	} else if value.Valid {
		s.IsFinished = value.Bool
	}
	if value, ok := values[4].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field duration", values[4])
	} else if value.Valid {
		s.Duration = int(value.Int64)
	}
	if value, ok := values[5].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field started_at", values[5])
	} else if value.Valid {
		s.StartedAt = value.Time
	}
	if value, ok := values[6].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field finished_at", values[6])
	} else if value.Valid {
		s.FinishedAt = new(time.Time)
		*s.FinishedAt = value.Time
	}
	return nil
}

// QueryEvents queries the events edge of the Session.
func (s *Session) QueryEvents() *EventQuery {
	return (&SessionClient{config: s.config}).QueryEvents(s)
}

// Update returns a builder for updating this Session.
// Note that, you need to call Session.Unwrap() before calling this method, if this Session
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Session) Update() *SessionUpdateOne {
	return (&SessionClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (s *Session) Unwrap() *Session {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Session is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Session) String() string {
	var builder strings.Builder
	builder.WriteString("Session(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", new_user=")
	builder.WriteString(fmt.Sprintf("%v", s.NewUser))
	builder.WriteString(", is_unique=")
	builder.WriteString(fmt.Sprintf("%v", s.IsUnique))
	builder.WriteString(", is_bounce=")
	builder.WriteString(fmt.Sprintf("%v", s.IsBounce))
	builder.WriteString(", is_finished=")
	builder.WriteString(fmt.Sprintf("%v", s.IsFinished))
	builder.WriteString(", duration=")
	builder.WriteString(fmt.Sprintf("%v", s.Duration))
	builder.WriteString(", started_at=")
	builder.WriteString(s.StartedAt.Format(time.ANSIC))
	if v := s.FinishedAt; v != nil {
		builder.WriteString(", finished_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Sessions is a parsable slice of Session.
type Sessions []*Session

func (s Sessions) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
