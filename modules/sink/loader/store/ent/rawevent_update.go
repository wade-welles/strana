// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/blushft/strana/modules/sink/loader/store/ent/predicate"
	"github.com/blushft/strana/modules/sink/loader/store/ent/rawevent"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// RawEventUpdate is the builder for updating RawEvent entities.
type RawEventUpdate struct {
	config
	hooks      []Hook
	mutation   *RawEventMutation
	predicates []predicate.RawEvent
}

// Where adds a new predicate for the builder.
func (reu *RawEventUpdate) Where(ps ...predicate.RawEvent) *RawEventUpdate {
	reu.predicates = append(reu.predicates, ps...)
	return reu
}

// SetTrackingID sets the tracking_id field.
func (reu *RawEventUpdate) SetTrackingID(u uuid.UUID) *RawEventUpdate {
	reu.mutation.SetTrackingID(u)
	return reu
}

// SetUserID sets the user_id field.
func (reu *RawEventUpdate) SetUserID(s string) *RawEventUpdate {
	reu.mutation.SetUserID(s)
	return reu
}

// SetAnonymous sets the anonymous field.
func (reu *RawEventUpdate) SetAnonymous(b bool) *RawEventUpdate {
	reu.mutation.SetAnonymous(b)
	return reu
}

// SetGroupID sets the group_id field.
func (reu *RawEventUpdate) SetGroupID(s string) *RawEventUpdate {
	reu.mutation.SetGroupID(s)
	return reu
}

// SetNillableGroupID sets the group_id field if the given value is not nil.
func (reu *RawEventUpdate) SetNillableGroupID(s *string) *RawEventUpdate {
	if s != nil {
		reu.SetGroupID(*s)
	}
	return reu
}

// ClearGroupID clears the value of group_id.
func (reu *RawEventUpdate) ClearGroupID() *RawEventUpdate {
	reu.mutation.ClearGroupID()
	return reu
}

// SetSessionID sets the session_id field.
func (reu *RawEventUpdate) SetSessionID(s string) *RawEventUpdate {
	reu.mutation.SetSessionID(s)
	return reu
}

// SetNillableSessionID sets the session_id field if the given value is not nil.
func (reu *RawEventUpdate) SetNillableSessionID(s *string) *RawEventUpdate {
	if s != nil {
		reu.SetSessionID(*s)
	}
	return reu
}

// ClearSessionID clears the value of session_id.
func (reu *RawEventUpdate) ClearSessionID() *RawEventUpdate {
	reu.mutation.ClearSessionID()
	return reu
}

// SetDeviceID sets the device_id field.
func (reu *RawEventUpdate) SetDeviceID(s string) *RawEventUpdate {
	reu.mutation.SetDeviceID(s)
	return reu
}

// SetNillableDeviceID sets the device_id field if the given value is not nil.
func (reu *RawEventUpdate) SetNillableDeviceID(s *string) *RawEventUpdate {
	if s != nil {
		reu.SetDeviceID(*s)
	}
	return reu
}

// ClearDeviceID clears the value of device_id.
func (reu *RawEventUpdate) ClearDeviceID() *RawEventUpdate {
	reu.mutation.ClearDeviceID()
	return reu
}

// SetEvent sets the event field.
func (reu *RawEventUpdate) SetEvent(s string) *RawEventUpdate {
	reu.mutation.SetEvent(s)
	return reu
}

// SetNonInteractive sets the non_interactive field.
func (reu *RawEventUpdate) SetNonInteractive(b bool) *RawEventUpdate {
	reu.mutation.SetNonInteractive(b)
	return reu
}

// SetChannel sets the channel field.
func (reu *RawEventUpdate) SetChannel(s string) *RawEventUpdate {
	reu.mutation.SetChannel(s)
	return reu
}

// SetNillableChannel sets the channel field if the given value is not nil.
func (reu *RawEventUpdate) SetNillableChannel(s *string) *RawEventUpdate {
	if s != nil {
		reu.SetChannel(*s)
	}
	return reu
}

// ClearChannel clears the value of channel.
func (reu *RawEventUpdate) ClearChannel() *RawEventUpdate {
	reu.mutation.ClearChannel()
	return reu
}

// SetPlatform sets the platform field.
func (reu *RawEventUpdate) SetPlatform(s string) *RawEventUpdate {
	reu.mutation.SetPlatform(s)
	return reu
}

// SetNillablePlatform sets the platform field if the given value is not nil.
func (reu *RawEventUpdate) SetNillablePlatform(s *string) *RawEventUpdate {
	if s != nil {
		reu.SetPlatform(*s)
	}
	return reu
}

// ClearPlatform clears the value of platform.
func (reu *RawEventUpdate) ClearPlatform() *RawEventUpdate {
	reu.mutation.ClearPlatform()
	return reu
}

// SetTimestamp sets the timestamp field.
func (reu *RawEventUpdate) SetTimestamp(t time.Time) *RawEventUpdate {
	reu.mutation.SetTimestamp(t)
	return reu
}

// SetContext sets the context field.
func (reu *RawEventUpdate) SetContext(m map[string]interface{}) *RawEventUpdate {
	reu.mutation.SetContext(m)
	return reu
}

// Mutation returns the RawEventMutation object of the builder.
func (reu *RawEventUpdate) Mutation() *RawEventMutation {
	return reu.mutation
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (reu *RawEventUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(reu.hooks) == 0 {
		affected, err = reu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RawEventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			reu.mutation = mutation
			affected, err = reu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(reu.hooks) - 1; i >= 0; i-- {
			mut = reu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, reu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (reu *RawEventUpdate) SaveX(ctx context.Context) int {
	affected, err := reu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (reu *RawEventUpdate) Exec(ctx context.Context) error {
	_, err := reu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (reu *RawEventUpdate) ExecX(ctx context.Context) {
	if err := reu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (reu *RawEventUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   rawevent.Table,
			Columns: rawevent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: rawevent.FieldID,
			},
		},
	}
	if ps := reu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := reu.mutation.TrackingID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: rawevent.FieldTrackingID,
		})
	}
	if value, ok := reu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rawevent.FieldUserID,
		})
	}
	if value, ok := reu.mutation.Anonymous(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: rawevent.FieldAnonymous,
		})
	}
	if value, ok := reu.mutation.GroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rawevent.FieldGroupID,
		})
	}
	if reu.mutation.GroupIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: rawevent.FieldGroupID,
		})
	}
	if value, ok := reu.mutation.SessionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rawevent.FieldSessionID,
		})
	}
	if reu.mutation.SessionIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: rawevent.FieldSessionID,
		})
	}
	if value, ok := reu.mutation.DeviceID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rawevent.FieldDeviceID,
		})
	}
	if reu.mutation.DeviceIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: rawevent.FieldDeviceID,
		})
	}
	if value, ok := reu.mutation.Event(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rawevent.FieldEvent,
		})
	}
	if value, ok := reu.mutation.NonInteractive(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: rawevent.FieldNonInteractive,
		})
	}
	if value, ok := reu.mutation.Channel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rawevent.FieldChannel,
		})
	}
	if reu.mutation.ChannelCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: rawevent.FieldChannel,
		})
	}
	if value, ok := reu.mutation.Platform(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rawevent.FieldPlatform,
		})
	}
	if reu.mutation.PlatformCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: rawevent.FieldPlatform,
		})
	}
	if value, ok := reu.mutation.Timestamp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: rawevent.FieldTimestamp,
		})
	}
	if value, ok := reu.mutation.Context(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: rawevent.FieldContext,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, reu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rawevent.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// RawEventUpdateOne is the builder for updating a single RawEvent entity.
type RawEventUpdateOne struct {
	config
	hooks    []Hook
	mutation *RawEventMutation
}

// SetTrackingID sets the tracking_id field.
func (reuo *RawEventUpdateOne) SetTrackingID(u uuid.UUID) *RawEventUpdateOne {
	reuo.mutation.SetTrackingID(u)
	return reuo
}

// SetUserID sets the user_id field.
func (reuo *RawEventUpdateOne) SetUserID(s string) *RawEventUpdateOne {
	reuo.mutation.SetUserID(s)
	return reuo
}

// SetAnonymous sets the anonymous field.
func (reuo *RawEventUpdateOne) SetAnonymous(b bool) *RawEventUpdateOne {
	reuo.mutation.SetAnonymous(b)
	return reuo
}

// SetGroupID sets the group_id field.
func (reuo *RawEventUpdateOne) SetGroupID(s string) *RawEventUpdateOne {
	reuo.mutation.SetGroupID(s)
	return reuo
}

// SetNillableGroupID sets the group_id field if the given value is not nil.
func (reuo *RawEventUpdateOne) SetNillableGroupID(s *string) *RawEventUpdateOne {
	if s != nil {
		reuo.SetGroupID(*s)
	}
	return reuo
}

// ClearGroupID clears the value of group_id.
func (reuo *RawEventUpdateOne) ClearGroupID() *RawEventUpdateOne {
	reuo.mutation.ClearGroupID()
	return reuo
}

// SetSessionID sets the session_id field.
func (reuo *RawEventUpdateOne) SetSessionID(s string) *RawEventUpdateOne {
	reuo.mutation.SetSessionID(s)
	return reuo
}

// SetNillableSessionID sets the session_id field if the given value is not nil.
func (reuo *RawEventUpdateOne) SetNillableSessionID(s *string) *RawEventUpdateOne {
	if s != nil {
		reuo.SetSessionID(*s)
	}
	return reuo
}

// ClearSessionID clears the value of session_id.
func (reuo *RawEventUpdateOne) ClearSessionID() *RawEventUpdateOne {
	reuo.mutation.ClearSessionID()
	return reuo
}

// SetDeviceID sets the device_id field.
func (reuo *RawEventUpdateOne) SetDeviceID(s string) *RawEventUpdateOne {
	reuo.mutation.SetDeviceID(s)
	return reuo
}

// SetNillableDeviceID sets the device_id field if the given value is not nil.
func (reuo *RawEventUpdateOne) SetNillableDeviceID(s *string) *RawEventUpdateOne {
	if s != nil {
		reuo.SetDeviceID(*s)
	}
	return reuo
}

// ClearDeviceID clears the value of device_id.
func (reuo *RawEventUpdateOne) ClearDeviceID() *RawEventUpdateOne {
	reuo.mutation.ClearDeviceID()
	return reuo
}

// SetEvent sets the event field.
func (reuo *RawEventUpdateOne) SetEvent(s string) *RawEventUpdateOne {
	reuo.mutation.SetEvent(s)
	return reuo
}

// SetNonInteractive sets the non_interactive field.
func (reuo *RawEventUpdateOne) SetNonInteractive(b bool) *RawEventUpdateOne {
	reuo.mutation.SetNonInteractive(b)
	return reuo
}

// SetChannel sets the channel field.
func (reuo *RawEventUpdateOne) SetChannel(s string) *RawEventUpdateOne {
	reuo.mutation.SetChannel(s)
	return reuo
}

// SetNillableChannel sets the channel field if the given value is not nil.
func (reuo *RawEventUpdateOne) SetNillableChannel(s *string) *RawEventUpdateOne {
	if s != nil {
		reuo.SetChannel(*s)
	}
	return reuo
}

// ClearChannel clears the value of channel.
func (reuo *RawEventUpdateOne) ClearChannel() *RawEventUpdateOne {
	reuo.mutation.ClearChannel()
	return reuo
}

// SetPlatform sets the platform field.
func (reuo *RawEventUpdateOne) SetPlatform(s string) *RawEventUpdateOne {
	reuo.mutation.SetPlatform(s)
	return reuo
}

// SetNillablePlatform sets the platform field if the given value is not nil.
func (reuo *RawEventUpdateOne) SetNillablePlatform(s *string) *RawEventUpdateOne {
	if s != nil {
		reuo.SetPlatform(*s)
	}
	return reuo
}

// ClearPlatform clears the value of platform.
func (reuo *RawEventUpdateOne) ClearPlatform() *RawEventUpdateOne {
	reuo.mutation.ClearPlatform()
	return reuo
}

// SetTimestamp sets the timestamp field.
func (reuo *RawEventUpdateOne) SetTimestamp(t time.Time) *RawEventUpdateOne {
	reuo.mutation.SetTimestamp(t)
	return reuo
}

// SetContext sets the context field.
func (reuo *RawEventUpdateOne) SetContext(m map[string]interface{}) *RawEventUpdateOne {
	reuo.mutation.SetContext(m)
	return reuo
}

// Mutation returns the RawEventMutation object of the builder.
func (reuo *RawEventUpdateOne) Mutation() *RawEventMutation {
	return reuo.mutation
}

// Save executes the query and returns the updated entity.
func (reuo *RawEventUpdateOne) Save(ctx context.Context) (*RawEvent, error) {
	var (
		err  error
		node *RawEvent
	)
	if len(reuo.hooks) == 0 {
		node, err = reuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RawEventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			reuo.mutation = mutation
			node, err = reuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(reuo.hooks) - 1; i >= 0; i-- {
			mut = reuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, reuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (reuo *RawEventUpdateOne) SaveX(ctx context.Context) *RawEvent {
	re, err := reuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return re
}

// Exec executes the query on the entity.
func (reuo *RawEventUpdateOne) Exec(ctx context.Context) error {
	_, err := reuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (reuo *RawEventUpdateOne) ExecX(ctx context.Context) {
	if err := reuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (reuo *RawEventUpdateOne) sqlSave(ctx context.Context) (re *RawEvent, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   rawevent.Table,
			Columns: rawevent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: rawevent.FieldID,
			},
		},
	}
	id, ok := reuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing RawEvent.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := reuo.mutation.TrackingID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: rawevent.FieldTrackingID,
		})
	}
	if value, ok := reuo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rawevent.FieldUserID,
		})
	}
	if value, ok := reuo.mutation.Anonymous(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: rawevent.FieldAnonymous,
		})
	}
	if value, ok := reuo.mutation.GroupID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rawevent.FieldGroupID,
		})
	}
	if reuo.mutation.GroupIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: rawevent.FieldGroupID,
		})
	}
	if value, ok := reuo.mutation.SessionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rawevent.FieldSessionID,
		})
	}
	if reuo.mutation.SessionIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: rawevent.FieldSessionID,
		})
	}
	if value, ok := reuo.mutation.DeviceID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rawevent.FieldDeviceID,
		})
	}
	if reuo.mutation.DeviceIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: rawevent.FieldDeviceID,
		})
	}
	if value, ok := reuo.mutation.Event(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rawevent.FieldEvent,
		})
	}
	if value, ok := reuo.mutation.NonInteractive(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: rawevent.FieldNonInteractive,
		})
	}
	if value, ok := reuo.mutation.Channel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rawevent.FieldChannel,
		})
	}
	if reuo.mutation.ChannelCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: rawevent.FieldChannel,
		})
	}
	if value, ok := reuo.mutation.Platform(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rawevent.FieldPlatform,
		})
	}
	if reuo.mutation.PlatformCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: rawevent.FieldPlatform,
		})
	}
	if value, ok := reuo.mutation.Timestamp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: rawevent.FieldTimestamp,
		})
	}
	if value, ok := reuo.mutation.Context(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: rawevent.FieldContext,
		})
	}
	re = &RawEvent{config: reuo.config}
	_spec.Assign = re.assignValues
	_spec.ScanValues = re.scanValues()
	if err = sqlgraph.UpdateNode(ctx, reuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rawevent.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return re, nil
}
