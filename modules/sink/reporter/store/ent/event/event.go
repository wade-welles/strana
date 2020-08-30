// Code generated by entc, DO NOT EDIT.

package event

import (
	"fmt"
)

const (
	// Label holds the string label denoting the event type in the database.
	Label = "event"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTrackingID holds the string denoting the tracking_id field in the database.
	FieldTrackingID = "tracking_id"
	// FieldEvent holds the string denoting the event field in the database.
	FieldEvent = "event"
	// FieldNonInteractive holds the string denoting the non_interactive field in the database.
	FieldNonInteractive = "non_interactive"
	// FieldChannel holds the string denoting the channel field in the database.
	FieldChannel = "channel"
	// FieldPlatform holds the string denoting the platform field in the database.
	FieldPlatform = "platform"
	// FieldProperties holds the string denoting the properties field in the database.
	FieldProperties = "properties"
	// FieldTimestamp holds the string denoting the timestamp field in the database.
	FieldTimestamp = "timestamp"

	// EdgeAction holds the string denoting the action edge name in mutations.
	EdgeAction = "action"
	// EdgeAlias holds the string denoting the alias edge name in mutations.
	EdgeAlias = "alias"
	// EdgeApp holds the string denoting the app edge name in mutations.
	EdgeApp = "app"
	// EdgeBrowser holds the string denoting the browser edge name in mutations.
	EdgeBrowser = "browser"
	// EdgeCampaign holds the string denoting the campaign edge name in mutations.
	EdgeCampaign = "campaign"
	// EdgeConnectivity holds the string denoting the connectivity edge name in mutations.
	EdgeConnectivity = "connectivity"
	// EdgeDevice holds the string denoting the device edge name in mutations.
	EdgeDevice = "device"
	// EdgeExtra holds the string denoting the extra edge name in mutations.
	EdgeExtra = "extra"
	// EdgeGroup holds the string denoting the group edge name in mutations.
	EdgeGroup = "group"
	// EdgeLibrary holds the string denoting the library edge name in mutations.
	EdgeLibrary = "library"
	// EdgeLocation holds the string denoting the location edge name in mutations.
	EdgeLocation = "location"
	// EdgeNetwork holds the string denoting the network edge name in mutations.
	EdgeNetwork = "network"
	// EdgeOs holds the string denoting the os edge name in mutations.
	EdgeOs = "os"
	// EdgePage holds the string denoting the page edge name in mutations.
	EdgePage = "page"
	// EdgeReferrer holds the string denoting the referrer edge name in mutations.
	EdgeReferrer = "referrer"
	// EdgeScreen holds the string denoting the screen edge name in mutations.
	EdgeScreen = "screen"
	// EdgeSession holds the string denoting the session edge name in mutations.
	EdgeSession = "session"
	// EdgeTiming holds the string denoting the timing edge name in mutations.
	EdgeTiming = "timing"
	// EdgeViewport holds the string denoting the viewport edge name in mutations.
	EdgeViewport = "viewport"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"

	// Table holds the table name of the event in the database.
	Table = "events"
	// ActionTable is the table the holds the action relation/edge.
	ActionTable = "actions"
	// ActionInverseTable is the table name for the Action entity.
	// It exists in this package in order to avoid circular dependency with the "action" package.
	ActionInverseTable = "actions"
	// ActionColumn is the table column denoting the action relation/edge.
	ActionColumn = "event_action"
	// AliasTable is the table the holds the alias relation/edge.
	AliasTable = "alias"
	// AliasInverseTable is the table name for the Alias entity.
	// It exists in this package in order to avoid circular dependency with the "alias" package.
	AliasInverseTable = "alias"
	// AliasColumn is the table column denoting the alias relation/edge.
	AliasColumn = "event_alias"
	// AppTable is the table the holds the app relation/edge.
	AppTable = "events"
	// AppInverseTable is the table name for the App entity.
	// It exists in this package in order to avoid circular dependency with the "app" package.
	AppInverseTable = "apps"
	// AppColumn is the table column denoting the app relation/edge.
	AppColumn = "event_app"
	// BrowserTable is the table the holds the browser relation/edge.
	BrowserTable = "events"
	// BrowserInverseTable is the table name for the Browser entity.
	// It exists in this package in order to avoid circular dependency with the "browser" package.
	BrowserInverseTable = "browsers"
	// BrowserColumn is the table column denoting the browser relation/edge.
	BrowserColumn = "event_browser"
	// CampaignTable is the table the holds the campaign relation/edge.
	CampaignTable = "events"
	// CampaignInverseTable is the table name for the Campaign entity.
	// It exists in this package in order to avoid circular dependency with the "campaign" package.
	CampaignInverseTable = "campaigns"
	// CampaignColumn is the table column denoting the campaign relation/edge.
	CampaignColumn = "event_campaign"
	// ConnectivityTable is the table the holds the connectivity relation/edge.
	ConnectivityTable = "events"
	// ConnectivityInverseTable is the table name for the Connectivity entity.
	// It exists in this package in order to avoid circular dependency with the "connectivity" package.
	ConnectivityInverseTable = "connectivities"
	// ConnectivityColumn is the table column denoting the connectivity relation/edge.
	ConnectivityColumn = "event_connectivity"
	// DeviceTable is the table the holds the device relation/edge.
	DeviceTable = "events"
	// DeviceInverseTable is the table name for the Device entity.
	// It exists in this package in order to avoid circular dependency with the "device" package.
	DeviceInverseTable = "devices"
	// DeviceColumn is the table column denoting the device relation/edge.
	DeviceColumn = "event_device"
	// ExtraTable is the table the holds the extra relation/edge.
	ExtraTable = "events"
	// ExtraInverseTable is the table name for the Extra entity.
	// It exists in this package in order to avoid circular dependency with the "extra" package.
	ExtraInverseTable = "extras"
	// ExtraColumn is the table column denoting the extra relation/edge.
	ExtraColumn = "event_extra"
	// GroupTable is the table the holds the group relation/edge.
	GroupTable = "events"
	// GroupInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupInverseTable = "groups"
	// GroupColumn is the table column denoting the group relation/edge.
	GroupColumn = "event_group"
	// LibraryTable is the table the holds the library relation/edge.
	LibraryTable = "events"
	// LibraryInverseTable is the table name for the Library entity.
	// It exists in this package in order to avoid circular dependency with the "library" package.
	LibraryInverseTable = "libraries"
	// LibraryColumn is the table column denoting the library relation/edge.
	LibraryColumn = "event_library"
	// LocationTable is the table the holds the location relation/edge.
	LocationTable = "events"
	// LocationInverseTable is the table name for the Location entity.
	// It exists in this package in order to avoid circular dependency with the "location" package.
	LocationInverseTable = "locations"
	// LocationColumn is the table column denoting the location relation/edge.
	LocationColumn = "event_location"
	// NetworkTable is the table the holds the network relation/edge.
	NetworkTable = "events"
	// NetworkInverseTable is the table name for the Network entity.
	// It exists in this package in order to avoid circular dependency with the "network" package.
	NetworkInverseTable = "networks"
	// NetworkColumn is the table column denoting the network relation/edge.
	NetworkColumn = "event_network"
	// OsTable is the table the holds the os relation/edge.
	OsTable = "events"
	// OsInverseTable is the table name for the OSContext entity.
	// It exists in this package in order to avoid circular dependency with the "oscontext" package.
	OsInverseTable = "os"
	// OsColumn is the table column denoting the os relation/edge.
	OsColumn = "event_os"
	// PageTable is the table the holds the page relation/edge.
	PageTable = "events"
	// PageInverseTable is the table name for the Page entity.
	// It exists in this package in order to avoid circular dependency with the "page" package.
	PageInverseTable = "pages"
	// PageColumn is the table column denoting the page relation/edge.
	PageColumn = "event_page"
	// ReferrerTable is the table the holds the referrer relation/edge.
	ReferrerTable = "events"
	// ReferrerInverseTable is the table name for the Referrer entity.
	// It exists in this package in order to avoid circular dependency with the "referrer" package.
	ReferrerInverseTable = "referrers"
	// ReferrerColumn is the table column denoting the referrer relation/edge.
	ReferrerColumn = "event_referrer"
	// ScreenTable is the table the holds the screen relation/edge.
	ScreenTable = "events"
	// ScreenInverseTable is the table name for the Screen entity.
	// It exists in this package in order to avoid circular dependency with the "screen" package.
	ScreenInverseTable = "screens"
	// ScreenColumn is the table column denoting the screen relation/edge.
	ScreenColumn = "event_screen"
	// SessionTable is the table the holds the session relation/edge.
	SessionTable = "events"
	// SessionInverseTable is the table name for the Session entity.
	// It exists in this package in order to avoid circular dependency with the "session" package.
	SessionInverseTable = "sessions"
	// SessionColumn is the table column denoting the session relation/edge.
	SessionColumn = "event_session"
	// TimingTable is the table the holds the timing relation/edge.
	TimingTable = "events"
	// TimingInverseTable is the table name for the Timing entity.
	// It exists in this package in order to avoid circular dependency with the "timing" package.
	TimingInverseTable = "timings"
	// TimingColumn is the table column denoting the timing relation/edge.
	TimingColumn = "event_timing"
	// ViewportTable is the table the holds the viewport relation/edge.
	ViewportTable = "events"
	// ViewportInverseTable is the table name for the Viewport entity.
	// It exists in this package in order to avoid circular dependency with the "viewport" package.
	ViewportInverseTable = "viewports"
	// ViewportColumn is the table column denoting the viewport relation/edge.
	ViewportColumn = "event_viewport"
	// UserTable is the table the holds the user relation/edge.
	UserTable = "events"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "event_user"
)

// Columns holds all SQL columns for event fields.
var Columns = []string{
	FieldID,
	FieldTrackingID,
	FieldEvent,
	FieldNonInteractive,
	FieldChannel,
	FieldPlatform,
	FieldProperties,
	FieldTimestamp,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Event type.
var ForeignKeys = []string{
	"event_app",
	"event_browser",
	"event_campaign",
	"event_connectivity",
	"event_device",
	"event_extra",
	"event_group",
	"event_library",
	"event_location",
	"event_network",
	"event_os",
	"event_page",
	"event_referrer",
	"event_screen",
	"event_session",
	"event_timing",
	"event_viewport",
	"event_user",
}

// Event defines the type for the event enum field.
type Event string

// Event values.
const (
	EventAction      Event = "action"
	EventAlias       Event = "alias"
	EventGroup       Event = "group"
	EventIdentify    Event = "identify"
	EventPageview    Event = "pageview"
	EventScreenview  Event = "screenview"
	EventSession     Event = "session"
	EventTiming      Event = "timing"
	EventTransaction Event = "transaction"
)

func (e Event) String() string {
	return string(e)
}

// EventValidator is a validator for the "event" field enum values. It is called by the builders before save.
func EventValidator(e Event) error {
	switch e {
	case EventAction, EventAlias, EventGroup, EventIdentify, EventPageview, EventScreenview, EventSession, EventTiming, EventTransaction:
		return nil
	default:
		return fmt.Errorf("event: invalid enum value for event field: %q", e)
	}
}