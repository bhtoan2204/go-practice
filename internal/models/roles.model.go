package models

type Role struct {
	BaseModel
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewRole(name, description string) *Role {
	return &Role{
		Name:        name,
		Description: description,
	}
}

func (r *Role) GetName() string {
	return r.Name
}

func (r *Role) SetName(name string) {
	r.Name = name
}
