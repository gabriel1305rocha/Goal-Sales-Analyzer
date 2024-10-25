package models

import (
	"gorm.io/gorm"
)

// Db é a conexão com o banco de dados que será compartilhada entre os modelos.
var Db *gorm.DB