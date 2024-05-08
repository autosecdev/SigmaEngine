package sigma

import "reflect"

// Event is a struct that contains information about the event
type Event struct {
	// Image is the image that is being run
	Image string
	// CommandLine is the command line that is being run
	CommandLine string
	// ParentImage is the parent image of the image that is being run
	ParentImage string
	// Matched is a boolean that indicates if the event matched the rule
	Matched bool
	// RuleTitle is the title of the rule that was matched
	RuleTitle string
}

// CheckEventFieldName checks if the field name exists in the Event struct
func (event *Event) CheckEventFieldName(fieldName string) bool {

	_, ok := reflect.TypeOf(event).Elem().FieldByName(fieldName)

	return ok
}

// FieldByName returns the value of the field name in the Event struct
func (event *Event) FieldByName(fieldName string) string {

	return reflect.ValueOf(event).Elem().FieldByName(fieldName).String()
}