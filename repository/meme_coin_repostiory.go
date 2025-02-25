package repository

import (
	"github.com/jacobintern/meme_coin/pkg/errors"
	"github.com/jacobintern/meme_coin/pkg/mysqlx"
	"github.com/jacobintern/meme_coin/service/adapter_model.go"
	"gorm.io/gorm"
)

type IMemeCoinRepository interface {
	Create(m adapter_model.MemeCoinInfo) error
	Get(id int64) (adapter_model.MemeCoinInfo, error)
	Update(entity adapter_model.MemeCoinInfo) error
	Delete(id int64) error
	Poke(id int64) error
}

type memeCoinRepository struct {
	mysql *mysqlx.Client
}

func NewMemeCoinRepository(client *mysqlx.Client) IMemeCoinRepository {
	return &memeCoinRepository{
		mysql: client,
	}
}

func (repo *memeCoinRepository) Create(e adapter_model.MemeCoinInfo) error {
	model := newMemeCoin().initModel(e)
	res := repo.mysql.DB.Create(&model)
	return res.Error
}

func (repo *memeCoinRepository) Get(id int64) (adapter_model.MemeCoinInfo, error) {
	model := newMemeCoin()
	res := repo.mysql.DB.FirstOrInit(&model, id)
	if res.Error != nil {
		return adapter_model.NewMemeCoinInfo(), res.Error
	}

	if model == newMemeCoin() {
		return adapter_model.NewMemeCoinInfo(), errors.MemeCoinNotFound
	}

	return model.toEntity(), nil
}

func (repo *memeCoinRepository) Update(entity adapter_model.MemeCoinInfo) error {
	model := newMemeCoin().toModel(entity)
	res := repo.mysql.DB.Updates(&model)
	return res.Error
}
func (repo *memeCoinRepository) Delete(id int64) error {
	model := newMemeCoin()
	res := repo.mysql.DB.Delete(&model, id)
	return res.Error
}
func (repo *memeCoinRepository) Poke(id int64) error {
	err := repo.mysql.DB.Transaction(func(tx *gorm.DB) error {
		model := newMemeCoin()
		r := tx.Model(&model).
			Where("id = ?", id).
			UpdateColumn("popularity_score", gorm.Expr("popularity_score + ?", 1))

		return r.Error
	})
	return err
}
