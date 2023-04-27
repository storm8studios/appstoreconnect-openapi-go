/*
App Store Connect API

Testing CiTestResultsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package appstoreopenapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/storm8studios/appstoreconnect-openapi-go/generated"
)

func Test_appstoreopenapi_CiTestResultsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CiTestResultsApiService CiTestResultsGetInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.CiTestResultsApi.CiTestResultsGetInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
