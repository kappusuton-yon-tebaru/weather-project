package werror

type ErrorSource string

const (
	INTERNAL ErrorSource = "internal"
	EXTERNAL ErrorSource = "external"
)

type WError struct {
	Source  ErrorSource `json:"source"`
	code    *int        `json:"-"`
	Message string      `json:"message"`
}

func NewFromError(err error, opts ...WErrorOption) WError {
	return New(err.Error(), opts...)
}

func New(msg string, opts ...WErrorOption) WError {
	werr := WError{
		Source:  INTERNAL,
		Message: msg,
	}

	for _, opt := range opts {
		werr = opt(werr)
	}

	return werr
}

func Must(err error) WError {
	werr, ok := err.(WError)
	if !ok {
		werr = NewFromError(err)
	}
	return werr
}

func (err WError) Error() string {
	return err.Message
}

func (err WError) Code(fallback int) int {
	if err.code == nil {
		return fallback
	}
	return *err.code
}
