package iamgo

type (
	Policy interface {
		PolicyScheme() string
	}

	PolicyDocument struct {
		Version   Version
		Statement []StatementEntry
	}

	StatementEntry struct {
		Effect   Effect
		Action   []string
		Resource string
	}
)
