// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/blushft/strana/modules/sink/reporter/store/ent/event"
	"github.com/blushft/strana/modules/sink/reporter/store/ent/predicate"
	"github.com/blushft/strana/modules/sink/reporter/store/ent/session"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// SessionUpdate is the builder for updating Session entities.
type SessionUpdate struct {
	config
	hooks      []Hook
	mutation   *SessionMutation
	predicates []predicate.Session
}

// Where adds a new predicate for the builder.
func (su *SessionUpdate) Where(ps ...predicate.Session) *SessionUpdate {
	su.predicates = append(su.predicates, ps...)
	return su
}

// SetNewUser sets the new_user field.
func (su *SessionUpdate) SetNewUser(b bool) *SessionUpdate {
	su.mutation.SetNewUser(b)
	return su
}

// SetIsUnique sets the is_unique field.
func (su *SessionUpdate) SetIsUnique(b bool) *SessionUpdate {
	su.mutation.SetIsUnique(b)
	return su
}

// SetIsBounce sets the is_bounce field.
func (su *SessionUpdate) SetIsBounce(b bool) *SessionUpdate {
	su.mutation.SetIsBounce(b)
	return su
}

// SetIsFinished sets the is_finished field.
func (su *SessionUpdate) SetIsFinished(b bool) *SessionUpdate {
	su.mutation.SetIsFinished(b)
	return su
}

// SetDuration sets the duration field.
func (su *SessionUpdate) SetDuration(i int) *SessionUpdate {
	su.mutation.ResetDuration()
	su.mutation.SetDuration(i)
	return su
}

// SetNillableDuration sets the duration field if the given value is not nil.
func (su *SessionUpdate) SetNillableDuration(i *int) *SessionUpdate {
	if i != nil {
		su.SetDuration(*i)
	}
	return su
}

// AddDuration adds i to duration.
func (su *SessionUpdate) AddDuration(i int) *SessionUpdate {
	su.mutation.AddDuration(i)
	return su
}

// ClearDuration clears the value of duration.
func (su *SessionUpdate) ClearDuration() *SessionUpdate {
	su.mutation.ClearDuration()
	return su
}

// SetStartedAt sets the started_at field.
func (su *SessionUpdate) SetStartedAt(t time.Time) *SessionUpdate {
	su.mutation.SetStartedAt(t)
	return su
}

// SetFinishedAt sets the finished_at field.
func (su *SessionUpdate) SetFinishedAt(t time.Time) *SessionUpdate {
	su.mutation.SetFinishedAt(t)
	return su
}

// SetNillableFinishedAt sets the finished_at field if the given value is not nil.
func (su *SessionUpdate) SetNillableFinishedAt(t *time.Time) *SessionUpdate {
	if t != nil {
		su.SetFinishedAt(*t)
	}
	return su
}

// ClearFinishedAt clears the value of finished_at.
func (su *SessionUpdate) ClearFinishedAt() *SessionUpdate {
	su.mutation.ClearFinishedAt()
	return su
}

// AddEventIDs adds the events edge to Event by ids.
func (su *SessionUpdate) AddEventIDs(ids ...uuid.UUID) *SessionUpdate {
	su.mutation.AddEventIDs(ids...)
	return su
}

// AddEvents adds the events edges to Event.
func (su *SessionUpdate) AddEvents(e ...*Event) *SessionUpdate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return su.AddEventIDs(ids...)
}

// Mutation returns the SessionMutation object of the builder.
func (su *SessionUpdate) Mutation() *SessionMutation {
	return su.mutation
}

// RemoveEventIDs removes the events edge to Event by ids.
func (su *SessionUpdate) RemoveEventIDs(ids ...uuid.UUID) *SessionUpdate {
	su.mutation.RemoveEventIDs(ids...)
	return su
}

// RemoveEvents removes events edges to Event.
func (su *SessionUpdate) RemoveEvents(e ...*Event) *SessionUpdate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return su.RemoveEventIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (su *SessionUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SessionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SessionUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SessionUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SessionUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SessionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   session.Table,
			Columns: session.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: session.FieldID,
			},
		},
	}
	if ps := su.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.NewUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: session.FieldNewUser,
		})
	}
	if value, ok := su.mutation.IsUnique(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: session.FieldIsUnique,
		})
	}
	if value, ok := su.mutation.IsBounce(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: session.FieldIsBounce,
		})
	}
	if value, ok := su.mutation.IsFinished(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: session.FieldIsFinished,
		})
	}
	if value, ok := su.mutation.Duration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: session.FieldDuration,
		})
	}
	if value, ok := su.mutation.AddedDuration(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: session.FieldDuration,
		})
	}
	if su.mutation.DurationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: session.FieldDuration,
		})
	}
	if value, ok := su.mutation.StartedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: session.FieldStartedAt,
		})
	}
	if value, ok := su.mutation.FinishedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: session.FieldFinishedAt,
		})
	}
	if su.mutation.FinishedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: session.FieldFinishedAt,
		})
	}
	if nodes := su.mutation.RemovedEventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   session.EventsTable,
			Columns: []string{session.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: event.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   session.EventsTable,
			Columns: []string{session.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: event.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// SessionUpdateOne is the builder for updating a single Session entity.
type SessionUpdateOne struct {
	config
	hooks    []Hook
	mutation *SessionMutation
}

// SetNewUser sets the new_user field.
func (suo *SessionUpdateOne) SetNewUser(b bool) *SessionUpdateOne {
	suo.mutation.SetNewUser(b)
	return suo
}

// SetIsUnique sets the is_unique field.
func (suo *SessionUpdateOne) SetIsUnique(b bool) *SessionUpdateOne {
	suo.mutation.SetIsUnique(b)
	return suo
}

// SetIsBounce sets the is_bounce field.
func (suo *SessionUpdateOne) SetIsBounce(b bool) *SessionUpdateOne {
	suo.mutation.SetIsBounce(b)
	return suo
}

// SetIsFinished sets the is_finished field.
func (suo *SessionUpdateOne) SetIsFinished(b bool) *SessionUpdateOne {
	suo.mutation.SetIsFinished(b)
	return suo
}

// SetDuration sets the duration field.
func (suo *SessionUpdateOne) SetDuration(i int) *SessionUpdateOne {
	suo.mutation.ResetDuration()
	suo.mutation.SetDuration(i)
	return suo
}

// SetNillableDuration sets the duration field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableDuration(i *int) *SessionUpdateOne {
	if i != nil {
		suo.SetDuration(*i)
	}
	return suo
}

// AddDuration adds i to duration.
func (suo *SessionUpdateOne) AddDuration(i int) *SessionUpdateOne {
	suo.mutation.AddDuration(i)
	return suo
}

// ClearDuration clears the value of duration.
func (suo *SessionUpdateOne) ClearDuration() *SessionUpdateOne {
	suo.mutation.ClearDuration()
	return suo
}

// SetStartedAt sets the started_at field.
func (suo *SessionUpdateOne) SetStartedAt(t time.Time) *SessionUpdateOne {
	suo.mutation.SetStartedAt(t)
	return suo
}

// SetFinishedAt sets the finished_at field.
func (suo *SessionUpdateOne) SetFinishedAt(t time.Time) *SessionUpdateOne {
	suo.mutation.SetFinishedAt(t)
	return suo
}

// SetNillableFinishedAt sets the finished_at field if the given value is not nil.
func (suo *SessionUpdateOne) SetNillableFinishedAt(t *time.Time) *SessionUpdateOne {
	if t != nil {
		suo.SetFinishedAt(*t)
	}
	return suo
}

// ClearFinishedAt clears the value of finished_at.
func (suo *SessionUpdateOne) ClearFinishedAt() *SessionUpdateOne {
	suo.mutation.ClearFinishedAt()
	return suo
}

// AddEventIDs adds the events edge to Event by ids.
func (suo *SessionUpdateOne) AddEventIDs(ids ...uuid.UUID) *SessionUpdateOne {
	suo.mutation.AddEventIDs(ids...)
	return suo
}

// AddEvents adds the events edges to Event.
func (suo *SessionUpdateOne) AddEvents(e ...*Event) *SessionUpdateOne {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return suo.AddEventIDs(ids...)
}

// Mutation returns the SessionMutation object of the builder.
func (suo *SessionUpdateOne) Mutation() *SessionMutation {
	return suo.mutation
}

// RemoveEventIDs removes the events edge to Event by ids.
func (suo *SessionUpdateOne) RemoveEventIDs(ids ...uuid.UUID) *SessionUpdateOne {
	suo.mutation.RemoveEventIDs(ids...)
	return suo
}

// RemoveEvents removes events edges to Event.
func (suo *SessionUpdateOne) RemoveEvents(e ...*Event) *SessionUpdateOne {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return suo.RemoveEventIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (suo *SessionUpdateOne) Save(ctx context.Context) (*Session, error) {

	var (
		err  error
		node *Session
	)
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SessionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SessionUpdateOne) SaveX(ctx context.Context) *Session {
	s, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return s
}

// Exec executes the query on the entity.
func (suo *SessionUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SessionUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SessionUpdateOne) sqlSave(ctx context.Context) (s *Session, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   session.Table,
			Columns: session.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: session.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Session.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := suo.mutation.NewUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: session.FieldNewUser,
		})
	}
	if value, ok := suo.mutation.IsUnique(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: session.FieldIsUnique,
		})
	}
	if value, ok := suo.mutation.IsBounce(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: session.FieldIsBounce,
		})
	}
	if value, ok := suo.mutation.IsFinished(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: session.FieldIsFinished,
		})
	}
	if value, ok := suo.mutation.Duration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: session.FieldDuration,
		})
	}
	if value, ok := suo.mutation.AddedDuration(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: session.FieldDuration,
		})
	}
	if suo.mutation.DurationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: session.FieldDuration,
		})
	}
	if value, ok := suo.mutation.StartedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: session.FieldStartedAt,
		})
	}
	if value, ok := suo.mutation.FinishedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: session.FieldFinishedAt,
		})
	}
	if suo.mutation.FinishedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: session.FieldFinishedAt,
		})
	}
	if nodes := suo.mutation.RemovedEventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   session.EventsTable,
			Columns: []string{session.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: event.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   session.EventsTable,
			Columns: []string{session.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: event.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	s = &Session{config: suo.config}
	_spec.Assign = s.assignValues
	_spec.ScanValues = s.scanValues()
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return s, nil
}
