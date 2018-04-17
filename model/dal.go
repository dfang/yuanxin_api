package model

type DataAccessLayer interface {
	Insert(db XODB) error
	Delete(db XODB) error
	Update(db XODB) error
	Save(db XODB) error
}
