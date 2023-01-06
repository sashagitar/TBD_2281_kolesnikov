package sqlmy

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestWithRedis(t *testing.T) {
	_, connPool := GetTestDatabase(t)

	// conn := "postgres://postgres:pass@localhost:5432/test"
	// connPool, err := Connect(conn)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	tf, err := time.Parse("2006-01-02 15:04", "2023-01-10")
	if err != nil {
		return
	}
	pr := ProduktDB{
		Id:          0,
		Id_user:     123,
		Name:        "Test",
		Weight:      1,
		Bought:      false,
		Used:        false,
		Thrown_out:  false,
		Date_start:  time.Now(),
		Date_finish: tf,
	}

	_, err = connPool.AddProdukt(123, &pr)
	if err != nil {

		t.Fatal(err)
	}

	t.Errorf("Success %s ", err)

}

func GetTestDatabase(t *testing.T) (testcontainers.Container, *sqlData) {
	// 1. Create PostgreSQL container request
	ctx := context.Background()
	// portt := "5432"
	user := "postgres"
	password := "pass"
	database := "test"
	req := testcontainers.ContainerRequest{
		Image:        "postgres:14.1-alpine",
		ExposedPorts: []string{"5432/tcp"},
		AutoRemove:   true,
		// WaitingFor:   wait.ForListeningPort("5432/tcp"),
		Env: map[string]string{
			"POSTGRES_USER":     user,
			"POSTGRES_PASSWORD": password,
			"POSTGRES_DB":       database,
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
		// testcontainers.ContainerRequest{
		// Image:        "postgres:14.1-alpine",
		// ExposedPorts: []string{portt},
		// Cmd:          []string{"postgres", "-c", "fsync=off"},
		// Env: map[string]string{
		// 	"POSTGRES_USER":     user,
		// 	"POSTGRES_PASSWORD": password,
		// 	"POSTGRES_DB":       database,
		// },
		// // WaitingFor: wait.ForLog("database system is ready to accept connections"),
		// wait.ForSQL(nat.Port(port), "postgres", dbURL).
		// 	WithStartupTimeout(time.Second * 5).
		// 	WithQuery("SELECT 10"),
	}
	redisC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(t, err)
	// return &TestDatabase{
	// 	instance: postgres,
	// }

	// // 2. Start PostgreSQL container
	// redisC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
	// 	ContainerRequest: req,
	// 	Started:          true,
	// })
	// if err != nil {
	// 	t.Error(err)
	// }
	// defer func() {
	// 	if err := redisC.Terminate(ctx); err != nil {
	// 		t.Fatalf("failed to terminate container: %s", err.Error())
	// 	}
	// }()
	// 3.1 Get host and port of PostgreSQL container
	// host, _ := redisC.Host(context.Background())
	// port, _ := redisC.MappedPort(context.Background(), "5432")

	// 3.2 Create db connection string and connect

	dbURI := "postgres://postgres:pass@127.0.0.1:5432/test"

	connPool, err := Connect(dbURI)
	err = fmt.Errorf("dbURI : %s, %w", dbURI, err)
	if err != nil {
		t.Fatal(err)
	}

	return redisC, connPool
}

func TestDatabase(t *testing.T) {
	// 1. Create PostgreSQL container request
	ctx := context.Background()
	portt := "5432"
	user := "postgres"
	password := "123"
	database := "test"
	req := testcontainers.ContainerRequest{
		Image:        "postgres:14.1-alpine",
		ExposedPorts: []string{portt},
		Cmd:          []string{"postgres", "-c", "fsync=off"},
		Env: map[string]string{
			"POSTGRES_USER":     user,
			"POSTGRES_PASSWORD": password,
			"POSTGRES_DB":       database,
		},
		WaitingFor: wait.ForLog("database system is ready to accept connections"),
		// wait.ForSQL(nat.Port(port), "postgres", dbURL).
		// 	WithStartupTimeout(time.Second * 5).
		// 	WithQuery("SELECT 10"),
	}

	// 2. Start PostgreSQL container
	redisC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := redisC.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err.Error())
		}
	}()

	// 3.1 Get host and port of PostgreSQL container
	host, _ := redisC.Host(context.Background())
	port, _ := redisC.MappedPort(context.Background(), "5432")

	// 3.2 Create db connection string and connect

	dbURI := fmt.Sprintf("host=%s port=%s user=postgres password=123 dbname=test sslmode=disable", host, port.Port())

	// connPool := ""
	_, err = Connect(dbURI)
	t.Fatal(err)

	// return redisC, connPool
}

// func TestAdd(t *testing.T) {

// }
