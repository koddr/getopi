package stores

import (
	"github.com/jmoiron/sqlx"
	"github.com/koddr/getopi/models"
)

// ProjectStore ...
type ProjectStore struct {
	*sqlx.DB
}

// FindProjectByAlias ...
//
// TODO: Add description
//
func (s *ProjectStore) FindProjectByAlias(alias string) (models.Project, error) {
	var project models.Project
	if err := s.Get(&project, `SELECT * FROM projects WHERE alias = $1`, alias); err != nil {
		return models.Project{}, err
	}
	return project, nil
}

// GetProjects ...
//
// TODO: Add description
//
func (s *ProjectStore) GetProjects() ([]models.Project, error) {
	var projects []models.Project
	if err := s.Select(
		&projects,
		`SELECT * FROM projects WHERE project_status = 1 AND project_attrs ->> 'is_private' = 'false'`,
	); err != nil {
		return []models.Project{}, err
	}
	return projects, nil
}
