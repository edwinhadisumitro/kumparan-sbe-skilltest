package news

import "database/sql"

type NewsSQL struct {
	ID      sql.NullInt64  `db:"id"`
	Author  sql.NullString `db:"author"`
	Body    sql.NullString `db:"body"`
	Created sql.NullString `db:"created"`
}

type News struct {
	ID      int    `json:"id"`
	Author  string `json:"author"`
	Body    string `json:"body"`
	Created string `json:"created"`
}

func (news *NewsSQL) ConvertToJSON() News {
	data := News{}

	if news.ID.Valid {
		data.ID = int(news.ID.Int64)
	}

	if news.Author.Valid {
		data.Author = news.Author.String
	}

	if news.Body.Valid {
		data.Body = news.Body.String
	}

	if news.Created.Valid {
		data.Created = news.Created.String
	}

	return data
}
