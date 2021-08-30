package main_test

import (
	"fmt"
	"net/http"
)

func (suite *AcceptanceTestSuite) TestGetItems() {
	suite.Run("Items listed successfully", func() {

		r, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/", suite.serverAddress), nil)
		suite.NoError(err)

		resp, err := suite.httpClient.Do(r)
		suite.NoError(err)

		suite.Equal(http.StatusOK, resp.StatusCode)
	})
}
func (suite *AcceptanceTestSuite) TestCreateItems() {
	suite.Run("Items listed successfully", func() {

		r, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/", suite.serverAddress), nil)
		suite.NoError(err)

		resp, err := suite.httpClient.Do(r)
		suite.NoError(err)

		suite.Equal(http.StatusOK, resp.StatusCode)
	})
}
