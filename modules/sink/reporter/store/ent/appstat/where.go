// Code generated by entc, DO NOT EDIT.

package appstat

import (
	"time"

	"github.com/blushft/strana/modules/sink/reporter/store/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
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
func IDGT(id int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Pageviews applies equality check predicate on the "pageviews" field. It's identical to PageviewsEQ.
func Pageviews(v int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPageviews), v))
	})
}

// Visitors applies equality check predicate on the "visitors" field. It's identical to VisitorsEQ.
func Visitors(v int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVisitors), v))
	})
}

// Sessions applies equality check predicate on the "sessions" field. It's identical to SessionsEQ.
func Sessions(v int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSessions), v))
	})
}

// BouceRate applies equality check predicate on the "bouce_rate" field. It's identical to BouceRateEQ.
func BouceRate(v float64) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBouceRate), v))
	})
}

// KnownDurations applies equality check predicate on the "known_durations" field. It's identical to KnownDurationsEQ.
func KnownDurations(v int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKnownDurations), v))
	})
}

// AvgDuration applies equality check predicate on the "avg_duration" field. It's identical to AvgDurationEQ.
func AvgDuration(v float64) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAvgDuration), v))
	})
}

// Date applies equality check predicate on the "date" field. It's identical to DateEQ.
func Date(v time.Time) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDate), v))
	})
}

// PageviewsEQ applies the EQ predicate on the "pageviews" field.
func PageviewsEQ(v int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPageviews), v))
	})
}

// PageviewsNEQ applies the NEQ predicate on the "pageviews" field.
func PageviewsNEQ(v int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPageviews), v))
	})
}

// PageviewsIn applies the In predicate on the "pageviews" field.
func PageviewsIn(vs ...int) predicate.AppStat {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppStat(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPageviews), v...))
	})
}

// PageviewsNotIn applies the NotIn predicate on the "pageviews" field.
func PageviewsNotIn(vs ...int) predicate.AppStat {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppStat(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPageviews), v...))
	})
}

// PageviewsGT applies the GT predicate on the "pageviews" field.
func PageviewsGT(v int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPageviews), v))
	})
}

// PageviewsGTE applies the GTE predicate on the "pageviews" field.
func PageviewsGTE(v int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPageviews), v))
	})
}

// PageviewsLT applies the LT predicate on the "pageviews" field.
func PageviewsLT(v int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPageviews), v))
	})
}

// PageviewsLTE applies the LTE predicate on the "pageviews" field.
func PageviewsLTE(v int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPageviews), v))
	})
}

// VisitorsEQ applies the EQ predicate on the "visitors" field.
func VisitorsEQ(v int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVisitors), v))
	})
}

// VisitorsNEQ applies the NEQ predicate on the "visitors" field.
func VisitorsNEQ(v int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldVisitors), v))
	})
}

// VisitorsIn applies the In predicate on the "visitors" field.
func VisitorsIn(vs ...int) predicate.AppStat {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppStat(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldVisitors), v...))
	})
}

// VisitorsNotIn applies the NotIn predicate on the "visitors" field.
func VisitorsNotIn(vs ...int) predicate.AppStat {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppStat(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldVisitors), v...))
	})
}

// VisitorsGT applies the GT predicate on the "visitors" field.
func VisitorsGT(v int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldVisitors), v))
	})
}

// VisitorsGTE applies the GTE predicate on the "visitors" field.
func VisitorsGTE(v int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldVisitors), v))
	})
}

// VisitorsLT applies the LT predicate on the "visitors" field.
func VisitorsLT(v int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldVisitors), v))
	})
}

// VisitorsLTE applies the LTE predicate on the "visitors" field.
func VisitorsLTE(v int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldVisitors), v))
	})
}

// SessionsEQ applies the EQ predicate on the "sessions" field.
func SessionsEQ(v int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSessions), v))
	})
}

// SessionsNEQ applies the NEQ predicate on the "sessions" field.
func SessionsNEQ(v int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSessions), v))
	})
}

// SessionsIn applies the In predicate on the "sessions" field.
func SessionsIn(vs ...int) predicate.AppStat {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppStat(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSessions), v...))
	})
}

// SessionsNotIn applies the NotIn predicate on the "sessions" field.
func SessionsNotIn(vs ...int) predicate.AppStat {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppStat(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSessions), v...))
	})
}

// SessionsGT applies the GT predicate on the "sessions" field.
func SessionsGT(v int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSessions), v))
	})
}

// SessionsGTE applies the GTE predicate on the "sessions" field.
func SessionsGTE(v int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSessions), v))
	})
}

// SessionsLT applies the LT predicate on the "sessions" field.
func SessionsLT(v int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSessions), v))
	})
}

// SessionsLTE applies the LTE predicate on the "sessions" field.
func SessionsLTE(v int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSessions), v))
	})
}

// BouceRateEQ applies the EQ predicate on the "bouce_rate" field.
func BouceRateEQ(v float64) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBouceRate), v))
	})
}

// BouceRateNEQ applies the NEQ predicate on the "bouce_rate" field.
func BouceRateNEQ(v float64) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBouceRate), v))
	})
}

// BouceRateIn applies the In predicate on the "bouce_rate" field.
func BouceRateIn(vs ...float64) predicate.AppStat {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppStat(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBouceRate), v...))
	})
}

// BouceRateNotIn applies the NotIn predicate on the "bouce_rate" field.
func BouceRateNotIn(vs ...float64) predicate.AppStat {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppStat(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBouceRate), v...))
	})
}

// BouceRateGT applies the GT predicate on the "bouce_rate" field.
func BouceRateGT(v float64) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBouceRate), v))
	})
}

// BouceRateGTE applies the GTE predicate on the "bouce_rate" field.
func BouceRateGTE(v float64) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBouceRate), v))
	})
}

// BouceRateLT applies the LT predicate on the "bouce_rate" field.
func BouceRateLT(v float64) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBouceRate), v))
	})
}

// BouceRateLTE applies the LTE predicate on the "bouce_rate" field.
func BouceRateLTE(v float64) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBouceRate), v))
	})
}

// KnownDurationsEQ applies the EQ predicate on the "known_durations" field.
func KnownDurationsEQ(v int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKnownDurations), v))
	})
}

// KnownDurationsNEQ applies the NEQ predicate on the "known_durations" field.
func KnownDurationsNEQ(v int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldKnownDurations), v))
	})
}

// KnownDurationsIn applies the In predicate on the "known_durations" field.
func KnownDurationsIn(vs ...int) predicate.AppStat {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppStat(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldKnownDurations), v...))
	})
}

// KnownDurationsNotIn applies the NotIn predicate on the "known_durations" field.
func KnownDurationsNotIn(vs ...int) predicate.AppStat {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppStat(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldKnownDurations), v...))
	})
}

// KnownDurationsGT applies the GT predicate on the "known_durations" field.
func KnownDurationsGT(v int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldKnownDurations), v))
	})
}

// KnownDurationsGTE applies the GTE predicate on the "known_durations" field.
func KnownDurationsGTE(v int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldKnownDurations), v))
	})
}

// KnownDurationsLT applies the LT predicate on the "known_durations" field.
func KnownDurationsLT(v int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldKnownDurations), v))
	})
}

// KnownDurationsLTE applies the LTE predicate on the "known_durations" field.
func KnownDurationsLTE(v int) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldKnownDurations), v))
	})
}

// AvgDurationEQ applies the EQ predicate on the "avg_duration" field.
func AvgDurationEQ(v float64) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAvgDuration), v))
	})
}

// AvgDurationNEQ applies the NEQ predicate on the "avg_duration" field.
func AvgDurationNEQ(v float64) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAvgDuration), v))
	})
}

// AvgDurationIn applies the In predicate on the "avg_duration" field.
func AvgDurationIn(vs ...float64) predicate.AppStat {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppStat(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAvgDuration), v...))
	})
}

// AvgDurationNotIn applies the NotIn predicate on the "avg_duration" field.
func AvgDurationNotIn(vs ...float64) predicate.AppStat {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppStat(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAvgDuration), v...))
	})
}

// AvgDurationGT applies the GT predicate on the "avg_duration" field.
func AvgDurationGT(v float64) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAvgDuration), v))
	})
}

// AvgDurationGTE applies the GTE predicate on the "avg_duration" field.
func AvgDurationGTE(v float64) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAvgDuration), v))
	})
}

// AvgDurationLT applies the LT predicate on the "avg_duration" field.
func AvgDurationLT(v float64) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAvgDuration), v))
	})
}

// AvgDurationLTE applies the LTE predicate on the "avg_duration" field.
func AvgDurationLTE(v float64) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAvgDuration), v))
	})
}

// DateEQ applies the EQ predicate on the "date" field.
func DateEQ(v time.Time) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDate), v))
	})
}

// DateNEQ applies the NEQ predicate on the "date" field.
func DateNEQ(v time.Time) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDate), v))
	})
}

// DateIn applies the In predicate on the "date" field.
func DateIn(vs ...time.Time) predicate.AppStat {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppStat(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDate), v...))
	})
}

// DateNotIn applies the NotIn predicate on the "date" field.
func DateNotIn(vs ...time.Time) predicate.AppStat {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppStat(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDate), v...))
	})
}

// DateGT applies the GT predicate on the "date" field.
func DateGT(v time.Time) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDate), v))
	})
}

// DateGTE applies the GTE predicate on the "date" field.
func DateGTE(v time.Time) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDate), v))
	})
}

// DateLT applies the LT predicate on the "date" field.
func DateLT(v time.Time) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDate), v))
	})
}

// DateLTE applies the LTE predicate on the "date" field.
func DateLTE(v time.Time) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDate), v))
	})
}

// HasApp applies the HasEdge predicate on the "app" edge.
func HasApp() predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AppTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AppTable, AppColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAppWith applies the HasEdge predicate on the "app" edge with a given conditions (other predicates).
func HasAppWith(preds ...predicate.App) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AppInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AppTable, AppColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.AppStat) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.AppStat) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
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
func Not(p predicate.AppStat) predicate.AppStat {
	return predicate.AppStat(func(s *sql.Selector) {
		p(s.Not())
	})
}