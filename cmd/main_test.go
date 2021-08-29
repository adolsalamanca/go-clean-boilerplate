package main_test

import (
	"fmt"
	"net"
	"os"
	"testing"
)

func LoadTestEnv(tcpPort string) {
	os.Setenv("SERVER_PORT", tcpPort)
	os.Setenv("STATSD_ADDRESS", "")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_NAME", "postgres")
	os.Setenv("DB_PASS", "")
}

func TestGetItems(t *testing.T) {
	/*
		cfg := config.LoadConfigProvider()
		port := getRandomTCPPort(t)
	*/
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

/*var (
	router  *mux.Router
	client  *http.Client
	address string
)

var _ = BeforeSuite(func() {

	client = &http.Client{Timeout: 5 * time.Second}
	server := muxApp.Server{}
	deps := &muxApp.Deps{}

	router = server.Start(deps)
	address = fmt.Sprintf("localhost:%s", getRandomTCPPort(GinkgoT()))

	go func() {
		if err := http.ListenAndServe(address, router); err != nil {
			fmt.Printf("unexpected error, %v", err.Error())
		}
	}()

})

var _ = Describe("Mux tests", func() {
	Describe("Perform request to the server", func() {
		Context("And root path without any body", func() {
			It("should return a 200 ok", func() {
				req, err := http.NewRequest(http.MethodGet, "http://"+address, nil)
				Expect(err).ShouldNot(HaveOccurred())

				resp, err := client.Do(req)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(resp.StatusCode).To(BeEquivalentTo(http.StatusOK))
			})
		})

		Context("To any unknown path", func() {
			It("should return a 404 not found", func() {
				req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s%s", "http://", address, "/fakePath"), nil)
				Expect(err).ShouldNot(HaveOccurred())

				resp, err := client.Do(req)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(resp.StatusCode).To(BeEquivalentTo(http.StatusNotFound))
			})
		})
	})

})
*/
