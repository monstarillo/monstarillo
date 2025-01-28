package engine

import "github.com/monstarillo/monstarillo/models"

type MonstarilloOrmContext struct {
	Models       []models.OrmModel
	Tags         []models.Tag
	CurrentModel models.OrmModel
}
