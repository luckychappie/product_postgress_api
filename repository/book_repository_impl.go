package repository

import (
	"context"
	"database/sql"
	"golang/product_postgress_api/helper"
	"golang/product_postgress_api/model"
)

type BookRepositoryImpl struct {
	Db *sql.DB
}

func NewBookRepository(Db *sql.DB) BookRepository {
	return &BookRepositoryImpl{Db: Db}
}

func (b *BookRepositoryImpl) Delete(ctx context.Context, bookId int) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer
}

func (*BookRepositoryImpl) FindAll(ctx context.Context) []model.Book {
	panic("unimplemented")
}

func (*BookRepositoryImpl) FindBy(ctx context.Context, bookId int) (model.Book, error) {
	panic("unimplemented")
}

func (*BookRepositoryImpl) Save(ctx context.Context, book model.Book) {
	panic("unimplemented")
}
