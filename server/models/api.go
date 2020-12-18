package models

import "time"

type Api struct {
	Id        *uint64
	Method    *string
	Path      *string
	UpdatedAt *time.Time
	CreatedAt *time.Time
}

func (m *Api) GetMethod() string {
	return *m.Method
}
func (m *Api) GetPath() string {
	return *m.Path
}
