package tabularasa

import (
	"testing"
)

func Test_allowedTabs(t *testing.T) {
	tests := []struct {
		name         string
		givenTabs    []Tab
		givenRoles   []string
		expectedTabs []Tab
	}{
		{
			name:         "no givenRoles, no tabs",
			givenTabs:    []Tab{},
			givenRoles:   []string{},
			expectedTabs: []Tab{},
		},
		{
			name:         "only one tab visible",
			givenTabs:    []Tab{{"tab", "esi", "state", []string{"role"}}, {"tab2", "esi", "state", []string{"no_role"}}},
			givenRoles:   []string{"role"},
			expectedTabs: []Tab{{"tab", "esi", "state", []string{"role"}}},
		},
		{
			name:         "all tabs allowed by multiple givenRoles",
			givenTabs:    []Tab{{"tab", "esi", "state", []string{"role"}}, {"tab2", "esi", "state", []string{"role"}}},
			givenRoles:   []string{"role", "role2"},
			expectedTabs: []Tab{{"tab", "esi", "state", []string{"role"}}, {"tab2", "esi", "state", []string{"role"}}},
		},
		{
			name:         "all tabs filtered by empty string role",
			givenTabs:    []Tab{{"tab", "esi", "state", []string{"role"}}, {"tab2", "esi", "state", []string{"no_role"}}},
			givenRoles:   []string{""},
			expectedTabs: []Tab{},
		},
		{
			name:         "all tabs filtered by no givenRoles at all",
			givenTabs:    []Tab{{"tab", "esi", "state", []string{"role"}}, {"tab2", "esi", "state", []string{"no_role"}}},
			givenRoles:   []string{""},
			expectedTabs: []Tab{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllowedTabs(tt.givenTabs, tt.givenRoles); !equal(got, tt.expectedTabs) {
				t.Errorf("allowedTabs() = %v, expectedTabs %v", got, tt.expectedTabs)
			}
		})
	}
}

func equal(a, b []Tab) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !v.Equal(b[i]) {
			return false
		}
	}
	return true
}
