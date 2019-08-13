package tabularasa

import "reflect"

type Tab struct {
	Title        string   `json:"title"`
	EsiLink      string   `json:"esi-link"`
	InitialState string   `json:"initial-state"`
	Roles        []string `json:"roles"`
}

type Tabs struct {
	Tabs []Tab `json:"tabs"`
}

func (tab *Tab) isAllowed(roles []string) bool {
	if len(tab.Roles) == 0 {
		return true
	} else if len(intersection(tab.Roles, roles)) > 0 {
		return true
	} else {
		return false
	}
}

func (tab *Tab) Equal(anotherTab Tab) bool {
	return tab.EsiLink == anotherTab.EsiLink &&
		tab.InitialState == anotherTab.InitialState &&
		reflect.DeepEqual(tab.Roles, anotherTab.Roles) &&
		tab.Title == anotherTab.Title
}

func AllowedTabs(tabs []Tab, roles []string) []Tab {
	var allowedTabs []Tab
	for _, tab := range tabs {
		if tab.isAllowed(roles) {
			allowedTabs = append(allowedTabs, tab)
		}
	}
	return allowedTabs
}

func intersection(a, b []string) (c []string) {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}
