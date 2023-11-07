package models

import (
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Name     string `json:"name" gorm:"unique" validate:"required"`
	Location string `json:"location" validate:"required"`
}

type Jobs struct {
	gorm.Model
	Company         Company           `json:"-" gorm:"ForeignKey:cid"`
	Cid             uint              `json:"cid"`
	MinNP           int               `json:"min_np"`
	MaxNP           int               `json:"max_np"`
	Budget          int               `json:"budget"`
	Locations       []Locations       `gorm:"many2many:jobs_locations;"`
	TechnologyStack []TechnologyStack `gorm:"many2many:jobs_technologyStack;"`
	WorkMode        []WorkMode        `gorm:"many2many:jobs_workMode;"`
	Description     string            `json:"description"`
	MinExp          int               `json:"min_exp"`
	MaxExp          int               `json:"max_exp"`
	Qualification   []Qualification   `gorm:"many2many:jobs_qualification;"`
	Shift           []Shift           `gorm:"many2many:jobs_shift;"`
	JobType         []JobType         `gorm:"many2many:jobs_jobType;"`
}

type Locations struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}

type TechnologyStack struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}

type WorkMode struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}

type Qualification struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}

type Shift struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}

type JobType struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}
