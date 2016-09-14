package project

import (
	"fmt"
	"time"

	db "servicecontrol.io/servicecontrol/lib/storage/postgresql"
)

// Project represents a project
type Project struct {
	TableName   struct{} `sql:"projects,alias:project"`
	ID          int64
	Name        string
	Description string
	Created     time.Time
}

// String returns projects in readable format
func (p Project) String() string {
	return fmt.Sprintf("Project<%d %s %s %v>", p.ID, p.Name, p.Description, p.Created)
}

// CreateProject creates a project
func CreateProject(p *Project) error {
	return db.Instance().Create(p)
}
