package model

type Buku struct {
	IDBuku    uint   `gorm:"primaryKey;column:id_buku" json:"id_buku"`
	Judul     string `gorm:"column:judul" json:"judul"`
	Pengarang string `gorm:"column:pengarang" json:"pengarang"`
	Stok      int    `gorm:"column:stok" json:"stok"`
}

func (Buku) TableName() string {
	return "buku"
}
