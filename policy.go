package iamgo

import "encoding/json"

type (
	Policy interface {
		PolicyScheme() ([]byte, error)
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

func NewLatestPolicy(st []StatementEntry) *PolicyDocument {
	return &PolicyDocument{
		Version:   Version20121017,
		Statement: st,
	}
}

// PolicyScheme for Policy interface
func (p *PolicyDocument) PolicyScheme() ([]byte, error) {
	b, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return b, nil
}
