// Code generated by entc, DO NOT EDIT.

package viewport

import (
	"github.com/blushft/strana/modules/sink/reporter/store/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Viewport {
	return predicate.Viewport(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Viewport {
	return predicate.Viewport(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Viewport {
	return predicate.Viewport(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Viewport {
	return predicate.Viewport(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Viewport {
	return predicate.Viewport(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Viewport {
	return predicate.Viewport(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Viewport {
	return predicate.Viewport(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Viewport {
	return predicate.Viewport(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Viewport {
	return predicate.Viewport(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Density applies equality check predicate on the "density" field. It's identical to DensityEQ.
func Density(v int) predicate.Viewport {
	return predicate.Viewport(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDensity), v))
	})
}

// Width applies equality check predicate on the "width" field. It's identical to WidthEQ.
func Width(v int) predicate.Viewport {
	return predicate.Viewport(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWidth), v))
	})
}

// Height applies equality check predicate on the "height" field. It's identical to HeightEQ.
func Height(v int) predicate.Viewport {
	return predicate.Viewport(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHeight), v))
	})
}

// DensityEQ applies the EQ predicate on the "density" field.
func DensityEQ(v int) predicate.Viewport {
	return predicate.Viewport(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDensity), v))
	})
}

// DensityNEQ applies the NEQ predicate on the "density" field.
func DensityNEQ(v int) predicate.Viewport {
	return predicate.Viewport(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDensity), v))
	})
}

// DensityIn applies the In predicate on the "density" field.
func DensityIn(vs ...int) predicate.Viewport {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Viewport(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDensity), v...))
	})
}

// DensityNotIn applies the NotIn predicate on the "density" field.
func DensityNotIn(vs ...int) predicate.Viewport {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Viewport(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDensity), v...))
	})
}

// DensityGT applies the GT predicate on the "density" field.
func DensityGT(v int) predicate.Viewport {
	return predicate.Viewport(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDensity), v))
	})
}

// DensityGTE applies the GTE predicate on the "density" field.
func DensityGTE(v int) predicate.Viewport {
	return predicate.Viewport(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDensity), v))
	})
}

// DensityLT applies the LT predicate on the "density" field.
func DensityLT(v int) predicate.Viewport {
	return predicate.Viewport(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDensity), v))
	})
}

// DensityLTE applies the LTE predicate on the "density" field.
func DensityLTE(v int) predicate.Viewport {
	return predicate.Viewport(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDensity), v))
	})
}

// WidthEQ applies the EQ predicate on the "width" field.
func WidthEQ(v int) predicate.Viewport {
	return predicate.Viewport(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWidth), v))
	})
}

// WidthNEQ applies the NEQ predicate on the "width" field.
func WidthNEQ(v int) predicate.Viewport {
	return predicate.Viewport(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWidth), v))
	})
}

// WidthIn applies the In predicate on the "width" field.
func WidthIn(vs ...int) predicate.Viewport {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Viewport(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldWidth), v...))
	})
}

// WidthNotIn applies the NotIn predicate on the "width" field.
func WidthNotIn(vs ...int) predicate.Viewport {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Viewport(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldWidth), v...))
	})
}

// WidthGT applies the GT predicate on the "width" field.
func WidthGT(v int) predicate.Viewport {
	return predicate.Viewport(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWidth), v))
	})
}

// WidthGTE applies the GTE predicate on the "width" field.
func WidthGTE(v int) predicate.Viewport {
	return predicate.Viewport(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWidth), v))
	})
}

// WidthLT applies the LT predicate on the "width" field.
func WidthLT(v int) predicate.Viewport {
	return predicate.Viewport(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWidth), v))
	})
}

// WidthLTE applies the LTE predicate on the "width" field.
func WidthLTE(v int) predicate.Viewport {
	return predicate.Viewport(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWidth), v))
	})
}

// HeightEQ applies the EQ predicate on the "height" field.
func HeightEQ(v int) predicate.Viewport {
	return predicate.Viewport(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHeight), v))
	})
}

// HeightNEQ applies the NEQ predicate on the "height" field.
func HeightNEQ(v int) predicate.Viewport {
	return predicate.Viewport(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHeight), v))
	})
}

// HeightIn applies the In predicate on the "height" field.
func HeightIn(vs ...int) predicate.Viewport {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Viewport(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHeight), v...))
	})
}

// HeightNotIn applies the NotIn predicate on the "height" field.
func HeightNotIn(vs ...int) predicate.Viewport {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Viewport(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHeight), v...))
	})
}

// HeightGT applies the GT predicate on the "height" field.
func HeightGT(v int) predicate.Viewport {
	return predicate.Viewport(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHeight), v))
	})
}

// HeightGTE applies the GTE predicate on the "height" field.
func HeightGTE(v int) predicate.Viewport {
	return predicate.Viewport(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHeight), v))
	})
}

// HeightLT applies the LT predicate on the "height" field.
func HeightLT(v int) predicate.Viewport {
	return predicate.Viewport(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHeight), v))
	})
}

// HeightLTE applies the LTE predicate on the "height" field.
func HeightLTE(v int) predicate.Viewport {
	return predicate.Viewport(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHeight), v))
	})
}

// HasEvents applies the HasEdge predicate on the "events" edge.
func HasEvents() predicate.Viewport {
	return predicate.Viewport(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EventsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, EventsTable, EventsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEventsWith applies the HasEdge predicate on the "events" edge with a given conditions (other predicates).
func HasEventsWith(preds ...predicate.Event) predicate.Viewport {
	return predicate.Viewport(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EventsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, EventsTable, EventsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Viewport) predicate.Viewport {
	return predicate.Viewport(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Viewport) predicate.Viewport {
	return predicate.Viewport(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Viewport) predicate.Viewport {
	return predicate.Viewport(func(s *sql.Selector) {
		p(s.Not())
	})
}