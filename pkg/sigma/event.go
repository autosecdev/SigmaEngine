package sigma

// Event is a struct that contains information about the event
type Event struct {
	// Image is the image that is being run
	Image string
	// CommandLine is the command line that is being run
	CommandLine string
	// ParentImage is the parent image of the image that is being run
	ParentImage string
}

