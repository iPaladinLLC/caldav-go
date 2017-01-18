package properties

import "strings"

type PropertyName string

const (
	UIDPropertyName                 PropertyName = "UID"
	CommentPropertyName                          = "COMMENT"
	OrganizerPropertyName                        = "ORGANIZER"
	AttendeePropertyName                         = "ATTENDEE"
	ExceptionDateTimesPropertyName               = "EXDATE"
	RecurrenceDateTimesPropertyName              = "RDATE"
	RecurrenceRulePropertyName                   = "RRULE"
	LocationPropertyName                         = "LOCATION"
	EmailPropertyName                            = "EMAIL"
	NamePropertyName                             = "N"
)

type ParameterName string

const (
	CanonicalNameParameterName  ParameterName = "CN"
	TimeZoneIdPropertyName                    = "TZID"
	ValuePropertyName                         = "VALUE"
	AlternateRepresentationName               = "ALTREP"
)

type Param struct {
	Name  ParameterName
	Value string
}

type Params []Param

func (p PropertyName) Equals(test string) bool {
	return strings.EqualFold(string(p), test)
}
