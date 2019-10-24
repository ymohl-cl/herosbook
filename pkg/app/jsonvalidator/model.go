package jsonvalidator

// Model force validate method implementationt
type Model interface {
	Validate() error
}
