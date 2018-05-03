package mysql

type SaimaGame struct {
	ID        int64 `gorm:"primary_key" json:"id"`
	WeddingId int64 `gorm:"not null; default:0; type:int; index" json:"wedding_id"`
	UserId    int64 `gorm:"not null; defalut:0; type:int; index" json:"user_id"`
	Seconds   int64 `gorm:"not null; default:0; type:int" json:"seconds"`
	Status    int64 `gorm:"not null; default:0; type:int" json:"status"`
	CreateAt  int64 `gorm:"not null; default:0; type:int" json:"create_at"`
	UpdateAt  int64 `gorm:"not null; default:0; type:int" json:"update_at"`
}

// todo: 赛马表
func (SaimaGame) TableName() string {
	return "saima_game"
}