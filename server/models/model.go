package models

import (
	"time"
)

type Model struct {
	Id        *uint64    `json:",omitempty"`
	UpdatedAt *time.Time `json:",omitempty"`
	CreatedAt *time.Time `json:",omitempty"`
}

/*
	用于清理模型固定数据
	防止在更新时带入
*/
func (m *Model) BeforeUpdate() {
	m.Id = nil
	m.UpdatedAt = nil
	m.CreatedAt = nil
}
