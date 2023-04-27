/*
App Store Connect API

Testing InAppPurchasePriceSchedulesApiService

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

func Test_appstoreopenapi_InAppPurchasePriceSchedulesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test InAppPurchasePriceSchedulesApiService InAppPurchasePriceSchedulesAutomaticPricesGetToManyRelated", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.InAppPurchasePriceSchedulesApi.InAppPurchasePriceSchedulesAutomaticPricesGetToManyRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InAppPurchasePriceSchedulesApiService InAppPurchasePriceSchedulesBaseTerritoryGetToOneRelated", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.InAppPurchasePriceSchedulesApi.InAppPurchasePriceSchedulesBaseTerritoryGetToOneRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InAppPurchasePriceSchedulesApiService InAppPurchasePriceSchedulesCreateInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.InAppPurchasePriceSchedulesApi.InAppPurchasePriceSchedulesCreateInstance(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InAppPurchasePriceSchedulesApiService InAppPurchasePriceSchedulesGetInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.InAppPurchasePriceSchedulesApi.InAppPurchasePriceSchedulesGetInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InAppPurchasePriceSchedulesApiService InAppPurchasePriceSchedulesManualPricesGetToManyRelated", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.InAppPurchasePriceSchedulesApi.InAppPurchasePriceSchedulesManualPricesGetToManyRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
