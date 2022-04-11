package model

import (
	"github.com/moaabid/golang-bitly/database"
)

func GetAllBitlies() ([]database.Bitly, error) {
	var bitlies []database.Bitly

	tx := database.Db.Find(&bitlies)

	if tx.Error != nil {
		return []database.Bitly{}, tx.Error
	}
	return bitlies, nil
}

func GetBilty(id uint64) (database.Bitly, error) {
	var bitly database.Bitly

	tx := database.Db.Where("id = ?", id).First(&bitly)

	if tx.Error != nil {
		return database.Bitly{}, tx.Error
	}
	return bitly, nil
}

func CreateBitly(bitly database.Bitly) error {
	tx := database.Db.Create(&bitly)
	return tx.Error
}

func UpdateBitly(bitly database.Bitly) error {
	tx := database.Db.Save(&bitly)
	return tx.Error
}

func DeleteBitly(id uint64) error {
	tx := database.Db.Unscoped().Delete(&database.Bitly{}, id)
	return tx.Error
}

func FindByBitlyUrl(url string) (database.Bitly, error) {
	var bitly database.Bitly

	tx := database.Db.Where("bitly= ?", url).First(&bitly)

	return bitly, tx.Error
}
