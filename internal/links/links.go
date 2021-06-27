package links

import (
	"log"

	"github.com/jmoiron/sqlx"
	database "github.com/saputradharma/go-graphql-example/internal/pkg/db/mysql"
	"github.com/saputradharma/go-graphql-example/internal/users"
)

type Link struct {
	ID      string `db:"id"`
	Title   string `db:"title"`
	Address string `db:"address"`
	User    *users.User
}

func CreateLink(link Link) int64 {
	query := `INSERT INTO links (title, address, user_id) VALUES(?, ?, ?)`

	result, err := database.DB.Exec(query, link.Title, link.Address, link.User.ID)

	if err != nil {
		log.Print(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Print(err)
	}

	log.Println("Link succesfully inserted!")
	return id
}

func GetAll() []Link {
	var result []Link
	var rows *sqlx.Rows

	query := `SELECT l.id, l.title, l.address, l.user_id, u.username ` +
		`FROM links l JOIN users u WHERE l.user_id = u.id`

	rows, err := database.DB.Queryx(query)

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		var temp Link
		var userId string
		var username string

		err := rows.Scan(&temp.ID, &temp.Title, &temp.Address, &userId, &username)

		if err != nil {
			log.Print(err)
		}

		temp.User = &users.User{
			ID:       userId,
			Username: username,
		}

		result = append(result, temp)
	}

	if err = rows.Err(); err != nil {
		log.Print(err)
	}

	return result

}
