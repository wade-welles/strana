// Code generated by entc, DO NOT EDIT.

package connectivity

const (
	// Label holds the string label denoting the connectivity type in the database.
	Label = "connectivity"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBluetooth holds the string denoting the bluetooth field in the database.
	FieldBluetooth = "bluetooth"
	// FieldCellular holds the string denoting the cellular field in the database.
	FieldCellular = "cellular"
	// FieldWifi holds the string denoting the wifi field in the database.
	FieldWifi = "wifi"
	// FieldEthernet holds the string denoting the ethernet field in the database.
	FieldEthernet = "ethernet"
	// FieldCarrier holds the string denoting the carrier field in the database.
	FieldCarrier = "carrier"
	// FieldIsp holds the string denoting the isp field in the database.
	FieldIsp = "isp"

	// EdgeEvent holds the string denoting the event edge name in mutations.
	EdgeEvent = "event"

	// Table holds the table name of the connectivity in the database.
	Table = "connectivities"
	// EventTable is the table the holds the event relation/edge.
	EventTable = "connectivities"
	// EventInverseTable is the table name for the Event entity.
	// It exists in this package in order to avoid circular dependency with the "event" package.
	EventInverseTable = "events"
	// EventColumn is the table column denoting the event relation/edge.
	EventColumn = "event_connectivity"
)

// Columns holds all SQL columns for connectivity fields.
var Columns = []string{
	FieldID,
	FieldBluetooth,
	FieldCellular,
	FieldWifi,
	FieldEthernet,
	FieldCarrier,
	FieldIsp,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Connectivity type.
var ForeignKeys = []string{
	"event_connectivity",
}
