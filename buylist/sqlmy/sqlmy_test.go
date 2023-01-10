package sqlmy

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	"github.com/jmoiron/sqlx"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "pass"
	dbname   = "test"
)

func TestCreateTestDatabase(t *testing.T) {
	containerReq := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
		Env: map[string]string{
			"POSTGRES_DB":       dbname,
			"POSTGRES_PASSWORD": password,
			"POSTGRES_USER":     user,
		},
	}

	dbContainer, err := testcontainers.GenericContainer(
		context.Background(),
		testcontainers.GenericContainerRequest{
			ContainerRequest: containerReq,
			Started:          true,
		})
	if err != nil {
		t.Errorf("fail create container: %v", err)
	}

	host, err := dbContainer.Host(context.Background())
	if err != nil {
		t.Errorf("fail get host container: %v", err)
	}
	port, err := dbContainer.MappedPort(context.Background(), "5432")
	if err != nil {
		t.Errorf("fail get port container: %v", err)
	}

	connString := fmt.Sprintf("postgres://%s:%s@%v:%v/%s?sslmode=disable", user, password, host, port.Port(), dbname)
	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		t.Errorf("fail sql connect: %v", err)
	}
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		t.Errorf("fail get driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://../../migrations",
		"postgres", driver)
	if err != nil {
		t.Errorf("fail migrate connect: %v", err)
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		t.Errorf("fail migration Up: %v", err)
	}
	defer dbContainer.Terminate(context.Background())
}

func CreateTestDatabase() (testcontainers.Container, *sqlData, error) {
	containerReq := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
		Env: map[string]string{
			"POSTGRES_DB":       dbname,
			"POSTGRES_PASSWORD": password,
			"POSTGRES_USER":     user,
		},
	}

	dbContainer, err := testcontainers.GenericContainer(
		context.Background(),
		testcontainers.GenericContainerRequest{
			ContainerRequest: containerReq,
			Started:          true,
		})
	if err != nil {
		return nil, nil, fmt.Errorf("fail create container: %w", err)
	}

	host, err := dbContainer.Host(context.Background())
	if err != nil {
		return nil, nil, fmt.Errorf("fail get host container: %w", err)
	}
	port, err := dbContainer.MappedPort(context.Background(), "5432")
	if err != nil {
		return nil, nil, fmt.Errorf("fail get port container: %w", err)
	}

	connString := fmt.Sprintf("postgres://%s:%s@%v:%v/%s?sslmode=disable", user, password, host, port.Port(), dbname)
	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		return nil, nil, fmt.Errorf("fail sql connect: %w", err)
	}
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return nil, nil, fmt.Errorf("fail get driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://../../migrations",
		"postgres", driver)
	if err != nil {
		return nil, nil, fmt.Errorf("fail migrate connect: %w", err)
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return nil, nil, fmt.Errorf("fail migration Up: %w", err)
	}

	sql := sqlData{
		db: db,
		m:  m,
	}
	return dbContainer, &sql, nil
}

func TestMigrationDown(t *testing.T) {
	container, sql, err := CreateTestDatabase()
	if err != nil {
		t.Errorf("fail Create test database: %v", err)
	}
	defer container.Terminate(context.Background())

	err = MigrationDown(sql.m)
	if err != nil {
		t.Error(err)
	}
}

func TestAddProdukt(t *testing.T) {
	container, sql, err := CreateTestDatabase()
	if err != nil {
		t.Errorf("fail Create test database: %v", err)
	}
	defer container.Terminate(context.Background())

	pr := ProduktDB{}
	_, err = sql.AddProdukt(0, &pr)
	if err != nil {
		t.Error(err)
	}
}
