package models

import "strconv"

type FirstFormOptionEnum int

const (
	OpenApp FirstFormOptionEnum = iota
	OpenGroupEnum
	ManageGroupsEnum
	RefetchAppsEnum
	QuitEnum
)

func (e FirstFormOptionEnum) String() string {
	switch e {
	case OpenApp:
		return "Open Apps"
	case OpenGroupEnum:
		return "Open Group"
	case ManageGroupsEnum:
		return "Manage Groups"
	case RefetchAppsEnum:
		return "Re-fetch Apps"
	case QuitEnum:
		return "Quit"
	default:
		return "Unknown"
	}
}
func (e FirstFormOptionEnum) StringIota() string {
	return strconv.Itoa(int(e))
}

// Manage Groups Form
type ManageFormOptionEnum int

const (
	CreateGroupEnum ManageFormOptionEnum = iota
	EditGroupEnum
	DeleteGroupEnum
)

func (e ManageFormOptionEnum) String() string {
	switch e {
	case CreateGroupEnum:
		return "Create Group"
	case EditGroupEnum:
		return "Edit Group"
	case DeleteGroupEnum:
		return "Delete Groups"
	default:
		return "Unknown"
	}
}

func (e ManageFormOptionEnum) StringIota() string {
	return strconv.Itoa(int(e))
}
