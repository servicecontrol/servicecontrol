package projects

import (
	"fmt"
	"time"

	db "servicecontrol.io/servicecontrol/lib/storage/postgresql"
)

// Project represents a project
type Project struct {
	ID          int64
	Name        string
	Description string
	Created     time.Time
}

// String returns projects in readable format
func (p Project) String() string {
	return fmt.Sprintf("User<%d %s %s %v", p.ID, p.Name, p.Description, p.Created)
}

// CreateProject creates a project
func CreateProject(p Project) error {
	return db.Instance().Create(p)
}
