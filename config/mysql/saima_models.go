package mysql

import (
	"time"
	"log"
)

func CreateSaimaGame(game *SaimaGame) error {
	game.CreateAt = time.Now().Unix()
	game.UpdateAt = time.Now().Unix()
	if err := db.Create(&game).Error; err != nil {
		return err
	}
	return nil
}

func GetNowPlaySaimaGame(weddingId int64) (*SaimaGame, bool) {
	info := &SaimaGame{}
	err := db.Where("wedding_id = ? and status in (1, 2)", weddingId).First(&info).Error
	if err != nil {
		log.Printf("get now game play info err: [%v]", err)
	}
	if info.ID == 0 {
		return nil, false
	}
	return info, true
}