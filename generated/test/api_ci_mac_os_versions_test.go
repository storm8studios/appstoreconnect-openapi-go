/*
App Store Connect API

Testing CiMacOsVersionsApiService

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

func Test_appstoreopenapi_CiMacOsVersionsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CiMacOsVersionsApiService CiMacOsVersionsGetCollection", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CiMacOsVersionsApi.CiMacOsVersionsGetCollection(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CiMacOsVersionsApiService CiMacOsVersionsGetInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.CiMacOsVersionsApi.CiMacOsVersionsGetInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CiMacOsVersionsApiService CiMacOsVersionsXcodeVersionsGetToManyRelated", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.CiMacOsVersionsApi.CiMacOsVersionsXcodeVersionsGetToManyRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
