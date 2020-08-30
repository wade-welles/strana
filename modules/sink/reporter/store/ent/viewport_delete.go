// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/blushft/strana/modules/sink/reporter/store/ent/predicate"
	"github.com/blushft/strana/modules/sink/reporter/store/ent/viewport"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// ViewportDelete is the builder for deleting a Viewport entity.
type ViewportDelete struct {
	config
	hooks      []Hook
	mutation   *ViewportMutation
	predicates []predicate.Viewport
}

// Where adds a new predicate to the delete builder.
func (vd *ViewportDelete) Where(ps ...predicate.Viewport) *ViewportDelete {
	vd.predicates = append(vd.predicates, ps...)
	return vd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vd *ViewportDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(vd.hooks) == 0 {
		affected, err = vd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ViewportMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			vd.mutation = mutation
			affected, err = vd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vd.hooks) - 1; i >= 0; i-- {
			mut = vd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (vd *ViewportDelete) ExecX(ctx context.Context) int {
	n, err := vd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vd *ViewportDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: viewport.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: viewport.FieldID,
			},
		},
	}
	if ps := vd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, vd.driver, _spec)
}

// ViewportDeleteOne is the builder for deleting a single Viewport entity.
type ViewportDeleteOne struct {
	vd *ViewportDelete
}

// Exec executes the deletion query.
func (vdo *ViewportDeleteOne) Exec(ctx context.Context) error {
	n, err := vdo.vd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{viewport.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vdo *ViewportDeleteOne) ExecX(ctx context.Context) {
	vdo.vd.ExecX(ctx)
}
