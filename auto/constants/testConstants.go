package constants

type Action string

const (
	CLICK    Action = "click"
	FILL     Action = "fill"
	NAVIGATE Action = "navigate"
	WAIT     Action = "wait"
	SNAP     Action = "snap"
)

type SelectorType string

const (
	Goto        SelectorType = "goto"
	Role        SelectorType = "role"
	Default     SelectorType = "default"
	Snap        SelectorType = "snap"
	Text        SelectorType = "text"
	Label       SelectorType = "label"
	Placeholder SelectorType = "placeholder"
	AltText     SelectorType = "altText"
	Title       SelectorType = "title"
	TestId      SelectorType = "testId"
	Wait        SelectorType = "wait"
)
