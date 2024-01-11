package render

// InvalidCode is the reason why validation failed.
type InvalidCode string

const (
	// CodeMissing means a resource does not exist
	CodeMissing InvalidCode = "missing"
	// CodeMissingField means a required field on a resource has not been set.
	CodeMissingField InvalidCode = "missing_field"
	// CodeInvalid means the formatting of a field is invalid
	CodeInvalid InvalidCode = "invalid"
	// CodeAlreadyExists means another resource has the same value as this field.
	CodeAlreadyExists InvalidCode = "already_exists"
)
