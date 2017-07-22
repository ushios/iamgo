package golicy

type (
	Policy interface {
		PolicyScheme() string
	}

	PolicyDocument struct {
		Version   string
		Statement []StatementEntry
	}

	StatementEntry struct {
		Effect   string
		Action   []string
		Resource string
	}
)
