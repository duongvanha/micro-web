package movie

import (
	"database/sql"
	model "github.com/duongvanha/micro-web/proto"
	_ "github.com/lib/pq"
	"github.com/russross/meddler"
	"os"
)

func init() {
	meddler.Default = meddler.PostgreSQL
	meddler.Mapper = meddler.LowerCase
}

type Repository struct {
}

func (r *Repository) GetByPage(page int) (movies []*model.Movie, err error) {

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return
	}

	err = meddler.QueryAll(db, &movies, "SELECT * FROM movies LIMIT 5")

	return
}
