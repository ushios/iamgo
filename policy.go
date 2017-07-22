package golicy

type (
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
