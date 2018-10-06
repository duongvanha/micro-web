package movie

import (
	model "../../proto"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/russross/meddler"
	"os"
)

func init() {
	meddler.Default = meddler.PostgreSQL
}

type Repository struct {
}

func (r *Repository) GetByPage(page int) (movies []*model.Movie, err error) {

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return
	}

	err = meddler.QueryAll(db, &model.Movie{}, "SELECT * FROM movies LIMIT 5")

	return
}
