/*
App Store Connect API

Testing BetaBuildLocalizationsApiService

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

func Test_appstoreopenapi_BetaBuildLocalizationsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BetaBuildLocalizationsApiService BetaBuildLocalizationsBuildGetToOneRelated", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.BetaBuildLocalizationsApi.BetaBuildLocalizationsBuildGetToOneRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BetaBuildLocalizationsApiService BetaBuildLocalizationsCreateInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BetaBuildLocalizationsApi.BetaBuildLocalizationsCreateInstance(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BetaBuildLocalizationsApiService BetaBuildLocalizationsDeleteInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.BetaBuildLocalizationsApi.BetaBuildLocalizationsDeleteInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BetaBuildLocalizationsApiService BetaBuildLocalizationsGetCollection", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BetaBuildLocalizationsApi.BetaBuildLocalizationsGetCollection(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BetaBuildLocalizationsApiService BetaBuildLocalizationsGetInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.BetaBuildLocalizationsApi.BetaBuildLocalizationsGetInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BetaBuildLocalizationsApiService BetaBuildLocalizationsUpdateInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.BetaBuildLocalizationsApi.BetaBuildLocalizationsUpdateInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
