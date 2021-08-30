package main_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"testing"
	"time"

	main "github.com/adolsalamanca/go-clean-boilerplate/cmd"
	"github.com/adolsalamanca/go-clean-boilerplate/internal/infrastructure/config"
	"github.com/adolsalamanca/go-clean-boilerplate/internal/infrastructure/environment"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/suite"
)

const (
	testDbPort = "5432"
	testDbHost = "localhost"
	testDbUser = "adol"
	testDbName = "database_name"
)

type AcceptanceTestSuite struct {
	suite.Suite
	serverAddress string
	dbConn        *sql.DB
	httpClient    *http.Client
}

func TestAcceptanceTestSuite(t *testing.T) {
	suite.Run(t, &AcceptanceTestSuite{})
}

func (suite *AcceptanceTestSuite) SetupSuite() {
	tcpPort := getRandomTCPPort(suite.T())
	hostPort := fmt.Sprintf("localhost:%s", tcpPort)

	suite.serverAddress = fmt.Sprintf("http://%s", hostPort)
	fmt.Printf("Server address: %s \n", suite.serverAddress)
	suite.httpClient = &http.Client{
		Timeout: time.Second * 5,
	}

	os.Setenv("SERVER_PORT", tcpPort)
	os.Setenv("DB_PORT", testDbPort)
	os.Setenv("DB_HOST", testDbHost)
	os.Setenv("DB_USER", testDbUser)
	os.Setenv("DB_NAME", testDbName)

	psqlConnectString := fmt.Sprintf("postgres://%s:@%s:%s/%s", testDbUser, testDbHost, testDbPort, testDbName)
	fmt.Printf("connection string: %v \n", psqlConnectString)

	conn, err := pgx.Connect(context.Background(), psqlConnectString)
	suite.NoError(err)
	_, err = conn.Query(context.Background(), "SELECT 1")
	suite.NoError(err)

	cfg := config.LoadConfigProvider()
	err = environment.Verify(cfg)
	if err != nil {
		log.Fatalf("could not initialize app: %v", err)
	}

	go main.Run(cfg)

	waitFor(hostPort)
}

func getRandomTCPPort(t *testing.T) string {
	t.Helper()
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Errorf("could not pick a random port, %v", err)
	}

	portNumber := listener.Addr().(*net.TCPAddr).Port
	listener.Close()
	return fmt.Sprintf("%d", portNumber)
}

func waitFor(address string) {
	timeout := 100 * time.Millisecond

	for i := 0; i <= 50; i++ {
		_, err := net.DialTimeout("tcp", address, timeout)
		if err == nil {
			return
		}

		if i%25 == 0 {
			log.Printf("Waiting for the service: %v\n", address)
		}
		time.Sleep(timeout)
	}
}
