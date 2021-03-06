// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/blushft/strana/modules/sink/reporter/store/ent/event"
	"github.com/blushft/strana/modules/sink/reporter/store/ent/network"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// NetworkCreate is the builder for creating a Network entity.
type NetworkCreate struct {
	config
	mutation *NetworkMutation
	hooks    []Hook
}

// SetIP sets the ip field.
func (nc *NetworkCreate) SetIP(s string) *NetworkCreate {
	nc.mutation.SetIP(s)
	return nc
}

// SetUseragent sets the useragent field.
func (nc *NetworkCreate) SetUseragent(s string) *NetworkCreate {
	nc.mutation.SetUseragent(s)
	return nc
}

// SetNillableUseragent sets the useragent field if the given value is not nil.
func (nc *NetworkCreate) SetNillableUseragent(s *string) *NetworkCreate {
	if s != nil {
		nc.SetUseragent(*s)
	}
	return nc
}

// AddEventIDs adds the events edge to Event by ids.
func (nc *NetworkCreate) AddEventIDs(ids ...uuid.UUID) *NetworkCreate {
	nc.mutation.AddEventIDs(ids...)
	return nc
}

// AddEvents adds the events edges to Event.
func (nc *NetworkCreate) AddEvents(e ...*Event) *NetworkCreate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return nc.AddEventIDs(ids...)
}

// Mutation returns the NetworkMutation object of the builder.
func (nc *NetworkCreate) Mutation() *NetworkMutation {
	return nc.mutation
}

// Save creates the Network in the database.
func (nc *NetworkCreate) Save(ctx context.Context) (*Network, error) {
	if err := nc.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *Network
	)
	if len(nc.hooks) == 0 {
		node, err = nc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NetworkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nc.mutation = mutation
			node, err = nc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(nc.hooks) - 1; i >= 0; i-- {
			mut = nc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NetworkCreate) SaveX(ctx context.Context) *Network {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (nc *NetworkCreate) preSave() error {
	if _, ok := nc.mutation.IP(); !ok {
		return &ValidationError{Name: "ip", err: errors.New("ent: missing required field \"ip\"")}
	}
	return nil
}

func (nc *NetworkCreate) sqlSave(ctx context.Context) (*Network, error) {
	n, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	n.ID = int(id)
	return n, nil
}

func (nc *NetworkCreate) createSpec() (*Network, *sqlgraph.CreateSpec) {
	var (
		n     = &Network{config: nc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: network.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: network.FieldID,
			},
		}
	)
	if value, ok := nc.mutation.IP(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: network.FieldIP,
		})
		n.IP = value
	}
	if value, ok := nc.mutation.Useragent(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: network.FieldUseragent,
		})
		n.Useragent = value
	}
	if nodes := nc.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   network.EventsTable,
			Columns: []string{network.EventsColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return n, _spec
}

// NetworkCreateBulk is the builder for creating a bulk of Network entities.
type NetworkCreateBulk struct {
	config
	builders []*NetworkCreate
}

// Save creates the Network entities in the database.
func (ncb *NetworkCreateBulk) Save(ctx context.Context) ([]*Network, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ncb.builders))
	nodes := make([]*Network, len(ncb.builders))
	mutators := make([]Mutator, len(ncb.builders))
	for i := range ncb.builders {
		func(i int, root context.Context) {
			builder := ncb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*NetworkMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ncb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ncb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (ncb *NetworkCreateBulk) SaveX(ctx context.Context) []*Network {
	v, err := ncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
