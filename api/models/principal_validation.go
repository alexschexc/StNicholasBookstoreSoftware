package models

import (
	"fmt"
	"strings"
)

// Scoped returns whether Principal has scope.
func (p *Principal) Scoped(scope string) error {
	if !strings.Contains(p.Scopes, scope) {
		return fmt.Errorf("scope %q is not found in %q", scope, p.Scopes)
	}

	return nil
}
