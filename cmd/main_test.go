package main_test

import (
	"fmt"
	"net/http"
	"strings"
)

func (suite *AcceptanceTestSuite) TestGetItems() {
	suite.Run("Items listed successfully", func() {

		r, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/items", suite.serverAddress), nil)
		suite.NoError(err)

		resp, err := suite.httpClient.Do(r)
		suite.NoError(err)

		suite.Equal(http.StatusOK, resp.StatusCode)
	})
}
func (suite *AcceptanceTestSuite) TestCreateItems() {
	suite.Run("Items listed successfully", func() {

		r, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/items", suite.serverAddress), strings.NewReader(`{"name":"car", "price": 6000}`))
		suite.NoError(err)

		resp, err := suite.httpClient.Do(r)
		suite.NoError(err)

		suite.Equal(http.StatusOK, resp.StatusCode)
	})
}
