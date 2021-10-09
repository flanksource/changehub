package v1

import "time"

type Identifier struct {
	Id, Name   string
	Attributes map[string]string
}

type Changes struct {
	FirstTimestamp, LastTimestamp time.Time
	Count                         int
	Scope                         []Scope
	Affects                       []interface{}
	Changes                       []Change
}

type Change struct {
	Type string
	// e.g. Succeeded, Failed, Slow, Flapping, Abnormal, Updated, Reverted
	Category    string
	Description string
	Icon        string
	Data        interface{}
	Identifier
}

type User struct {
	Type string
	Ref  string
	Identifier
}

type Scope struct {
	Type string
	Identifier
}

type Network struct {
	Identifier
	CIDR string
}

type MaintenanceWindow struct {
	From, To time.Time
}

type Scaling struct {
	Old, New int
}

type Diff struct {
	Identifier
}

type GitRepository struct {
	Identifier
}

type Check struct {
	Identifier
}

// Encompasses all
type Health struct {
	HealthType  string // e.g. canary, test, alert, slo
	Description string
	Identifier
}
