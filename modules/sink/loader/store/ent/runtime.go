// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/blushft/strana/modules/sink/loader/schema"
	"github.com/blushft/strana/modules/sink/loader/store/ent/rawevent"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	raweventFields := schema.RawEvent{}.Fields()
	_ = raweventFields
	// raweventDescID is the schema descriptor for id field.
	raweventDescID := raweventFields[0].Descriptor()
	// rawevent.DefaultID holds the default value on creation for the id field.
	rawevent.DefaultID = raweventDescID.Default.(func() uuid.UUID)
}
