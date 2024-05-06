package sigma

// Rule is a struct that contains information about the Sigma rule
type Rule struct {
	Title          string
	ID             string
	Status         string
	Description    string
	References     []string
	Author         string
	Date           string
	Modified       string
	Tags           []string
	LogSource      LogSource
	Detection      map[string]interface{}
	FalsePositives []string
	Level          string
}

// LogSource is a struct that contains information about the log source
type LogSource struct {
	Product  string
	Category string
}

// Match is a function that matches the Sigma rule
func (r *Rule) Match(event Event) bool {

	return false
}