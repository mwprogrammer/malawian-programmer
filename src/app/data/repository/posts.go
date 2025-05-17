package repository

import (
	"github.com/mwprogrammer/malawian-programmer/src/app/data/database"
	"github.com/mwprogrammer/malawian-programmer/src/app/models"
)

func AddPost(post models.Post) error {

	err := database.AddRecord(post)

	if err != nil {
		return err
	}

	return nil

}