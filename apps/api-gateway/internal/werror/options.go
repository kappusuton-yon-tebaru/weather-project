package werror

type WErrorOption func(werr WError) WError

func WithExternalSource() WErrorOption {
	return func(werr WError) WError {
		werr.Source = EXTERNAL
		return werr
	}
}

func WithCode(code int) WErrorOption {
	return func(werr WError) WError {
		werr.code = &code
		return werr
	}
}
