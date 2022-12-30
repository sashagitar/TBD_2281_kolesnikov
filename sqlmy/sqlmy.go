package sqlmy

import (
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type ProduktDB struct {
	Id          int       `db:"id"`
	Id_user     int       `db:"id_user"`
	Name        string    `db:"name"`
	Weight      float64   `db:"weight"`
	Bought      bool      `db:"bought"`
	Used        bool      `db:"used"`
	Thrown_out  bool      `db:"trown_out"`
	Date_start  time.Time `db:"date_start"`
	Date_finish time.Time `db:"date_finish"`
}

type Users struct {
	Id    int    `db:"id"`
	Id_tg string `db:"id_tg"`
}

type Notification struct {
	Id       int    `db:"id"`
	Interval int    `db:"days"`
	Time     string `db:"time"`
}

type store struct {
	db *sqlx.DB
}

func Connect(connUri string) *store {
	db, err := sqlx.Connect("pgx", connUri)
	if err != nil {
		return nil
	}
	s := store{db: db}
	return &s

}

func AddProdukt(id_user int, p *ProduktDB) (bool, error) {
	sql := "insert into produkts (id_user, name, weight, bought, used, trown_out, date_start, date_finish) values (:id_user, :name, :weight, :bought, :used, :trown_out, :date_start, :date_finish)"
	_, err := DB.NamedExec(sql, *p)
	if err != nil {
		return false, err
	}
	return true, err
}

func Trash(id int, id_user int) (bool, error) {
	sql := "update produkts set trown_out = true where id = $2 and id_user = $3 and bought = true"
	_, err := DB.Exec(sql, id, id_user)
	if err != nil {
		return false, err
	}
	return true, err
}

func Used(id int, id_user int) (bool, error) {
	sql := "update produkts set used = true date_finish = NOW() where id = $2 and id_user = $3 and bought = true"
	_, err := DB.Exec(sql, id, id_user)
	if err != nil {
		return false, err
	}
	return true, err
}

func SetFinish(id_user int, id int, date time.Time) (bool, error) {
	sql := "update produkts set date_finish = $1 where id = $2 and id_user = $3"
	_, err := DB.Exec(sql, date, id, id_user)
	if err != nil {
		return false, err
	}
	return true, err
}

func MoveInHolodos(id_user int, id int) (bool, error) {
	sql := "update produkts set bought = true, date_start = NOW() where id = $1 and id_user = $2"
	_, err := DB.Exec(sql, id, id_user)
	if err != nil {
		return false, err
	}
	return true, err
}

func GetList(id_user int, bought bool, sort bool) (*[]ProduktDB, error) {
	prdkts := make([]ProduktDB, 0)
	sql := "SELECT * FROM produkts where id_user = $1 and bought = $2 order by "
	v1 := "name"
	v2 := "date_finish, name"
	if sort {
		sql += v1
	} else {
		sql += v2
	}
	if err := DB.Select(&prdkts, sql, id_user, bought); err == nil {
		return nil, err
	}
	if len(prdkts) == 0 {
		return nil, nil
	}
	return &prdkts, nil
}

func GetUseList(id_user int) (*[]ProduktDB, error) {
	sql := "select * from produkts where id_user = $1 and used = true"
	prdkts := make([]ProduktDB, 0)
	if err := DB.Select(&prdkts, sql, id_user); err == nil {
		return nil, err
	}
	return &prdkts, nil
}

func GetStats(id_user int, ts time.Time, tf time.Time) (int, int, error) {
	var use int
	var drop int
	var err error
	if err = DB.Select(&use, "SELECT count(*) FROM produkts where id_user = $1 and bought = false and  used = true date_finish >= $2 adn date_finish <= $3 ", id_user, ts, tf); err == nil {
		if err = DB.Select(&drop, "SELECT count(*) FROM produkts where id_user = $1 and bought = false and  trown_out = true", id_user); err == nil {
			return use, drop, err
		}
		return -1, -1, err
	}
	return -1, -1, err
}
