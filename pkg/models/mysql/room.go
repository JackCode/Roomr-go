package mysql

import (
	"database/sql"

	"github.com/jackcode/roomr-go/pkg/models"
)

type RoomModel struct {
	DB *sql.DB
}

func (m *RoomModel) GetAllRooms() ([]*models.Room, error) {

	tx, err := m.DB.Begin()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query("SELECT * FROM get_all_rooms_v")
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer rows.Close()

	rooms := []*models.Room{}

	for rows.Next() {
		room := &models.Room{
			BathroomStyle: &models.BathroomStyle{},
			RoomType:      &models.RoomType{},
		}

		err := rows.Scan(
			&room.RoomNumber, &room.Floor, &room.NumberOfShowerHeads, &room.NumberOfSinks, &room.SquareFootage,
			&room.BathroomStyle.ID, &room.BathroomStyle.Title, &room.BathroomStyle.LongDescription,
			&room.RoomType.ID, &room.RoomType.ShortCode, &room.RoomType.LongDescription, &room.RoomType.NumberOfBeds)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		rooms = append(rooms, room)
	}

	if err = rows.Err(); err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return rooms, nil
}
