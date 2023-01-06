package sqlmy

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-migrate/migrate"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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

type sqlData struct {
	db     *sqlx.DB
	People *Users
}

var connDB = "postgres://postgres:123@127.0.0.1:5432/test?sslmode=require"

func MigrateDb(connUri string) error {
	// make migration

	// Read migrations from /home/migrations and connect to a local postgres database.
	m, err := migrate.New("file://migrations", connDB)
	if err != nil {
		panic(err)
	}

	// Migrate all the way up ...
	if err := m.Up(); err != nil {
		panic(err)
	}
	err = m.Force(1)
	if err := m.Down(); err != nil {
		return fmt.Errorf("sqlmy MigrateDb: %w", err)
	}
	log.Println("migration is done")
	return nil
	//11 is migrations version number, you may use your latest version

}

func MigrationDown(connUri string) error {
	m, err := migrate.New("file://migrations", connDB)
	if err != nil {
		panic(err)
	}
	err = m.Force(1)
	if err := m.Down(); err != nil {
		return fmt.Errorf("sqlmy MigrationDown: %w", err)
	}
	log.Println("migration is done")

	//11 is migrations version number, you may use your latest version
	if err != nil {
		return fmt.Errorf("sqlmy MigrationDown: %w", err)
	}
	return nil
}

func Connect(connUri string) (*sqlData, error) {
	db, err := sqlx.Connect("postgres", connUri)
	if err != nil {
		return nil, fmt.Errorf("sqlmy Connect: %w", err)
	}
	u := Users{
		Id:    0,
		Id_tg: "123",
	}
	s := sqlData{
		db:     db,
		People: &u,
	}
	err = MigrateDb(connUri)
	return &s, nil

}

func (s *sqlData) AddProdukt(id_user int, p *ProduktDB) (bool, error) {
	sql := "insert into produkts (id_user, name, weight, bought, used, trown_out, date_start, date_finish) values (:id_user, :name, :weight, :bought, :used, :trown_out, :date_start, :date_finish)"
	_, err := s.db.NamedExec(sql, *p)
	if err != nil {
		return false, fmt.Errorf("sqlmy AddProdukt: %w", err)
	}
	return true, err
}

func (s *sqlData) Trash(id int, id_user int) (bool, error) {
	sql := "update produkts set trown_out = true where id = $2 and id_user = $3 and bought = true"
	_, err := s.db.Exec(sql, id, id_user)
	if err != nil {
		return false, fmt.Errorf("sqlmy Teash: %w", err)
	}
	return true, err
}

func (s *sqlData) Used(id int, id_user int) (bool, error) {
	sql := "update produkts set used = true date_finish = NOW() where id = $2 and id_user = $3 and bought = true"
	_, err := s.db.Exec(sql, id, id_user)
	if err != nil {
		return false, fmt.Errorf("sqlmy Used: %w", err)
	}
	return true, err
}

func (s *sqlData) SetFinish(id_user int, id int, date time.Time) (bool, error) {
	sql := "update produkts set date_finish = $1 where id = $2 and id_user = $3"
	_, err := s.db.Exec(sql, date, id, id_user)
	if err != nil {
		return false, fmt.Errorf("sqlmy SetFinish: %w", err)
	}
	return true, err
}

func (s *sqlData) MoveInHolodos(id_user int, id int) (bool, error) {
	sql := "update produkts set bought = true, date_start = NOW() where id = $1 and id_user = $2"
	_, err := s.db.Exec(sql, id, id_user)
	if err != nil {
		return false, fmt.Errorf("sqlmy MoveInHolodos: %w", err)
	}
	return true, err
}

func (s *sqlData) GetList(id_user int, bought bool, sort bool) (*[]ProduktDB, error) {
	prdkts := make([]ProduktDB, 0)
	sql := "SELECT * FROM produkts where id_user = $1 and bought = $2 order by "
	v1 := "name"
	v2 := "date_finish, name"
	if sort {
		sql += v1
	} else {
		sql += v2
	}
	if err := s.db.Select(&prdkts, sql, id_user, bought); err == nil {
		return nil, fmt.Errorf("sqlmy GetList: %w", err)
	}
	if len(prdkts) == 0 {
		return nil, nil
	}
	return &prdkts, nil
}

func (s *sqlData) GetUseList(id_user int) (*[]ProduktDB, error) {
	sql := "select * from produkts where id_user = $1 and used = true"
	prdkts := make([]ProduktDB, 0)
	if err := s.db.Select(&prdkts, sql, id_user); err == nil {
		return nil, fmt.Errorf("sqlmy GetUseList: %w", err)
	}
	return &prdkts, nil
}

func (s *sqlData) GetStats(id_user int, ts time.Time, tf time.Time) (int, int, error) {
	var use int
	var drop int
	var err error
	if err = s.db.Select(&use, "SELECT count(*) FROM produkts where id_user = $1 and bought = false and  used = true date_finish >= $2 adn date_finish <= $3 ", id_user, ts, tf); err == nil {
		if err = s.db.Select(&drop, "SELECT count(*) FROM produkts where id_user = $1 and bought = false and  trown_out = true", id_user); err == nil {
			return use, drop, fmt.Errorf("sqlmy GetStats drop: %w", err)
		}
		return -1, -1, fmt.Errorf("sqlmy GetStats use: %w", err)
	}
	return -1, -1, err
}
