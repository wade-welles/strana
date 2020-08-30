// Code generated by entc, DO NOT EDIT.

package referrer

const (
	// Label holds the string label denoting the referrer type in the database.
	Label = "referrer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldHostname holds the string denoting the hostname field in the database.
	FieldHostname = "hostname"
	// FieldLink holds the string denoting the link field in the database.
	FieldLink = "link"

	// EdgeEvents holds the string denoting the events edge name in mutations.
	EdgeEvents = "events"

	// Table holds the table name of the referrer in the database.
	Table = "referrers"
	// EventsTable is the table the holds the events relation/edge.
	EventsTable = "events"
	// EventsInverseTable is the table name for the Event entity.
	// It exists in this package in order to avoid circular dependency with the "event" package.
	EventsInverseTable = "events"
	// EventsColumn is the table column denoting the events relation/edge.
	EventsColumn = "event_referrer"
)

// Columns holds all SQL columns for referrer fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldType,
	FieldHostname,
	FieldLink,
}
