# \SalesReportsApi

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SalesReportsGetCollection**](SalesReportsApi.md#SalesReportsGetCollection) | **Get** /v1/salesReports | 



## SalesReportsGetCollection

> *os.File SalesReportsGetCollection(ctx).FilterFrequency(filterFrequency).FilterReportSubType(filterReportSubType).FilterReportType(filterReportType).FilterVendorNumber(filterVendorNumber).FilterReportDate(filterReportDate).FilterVersion(filterVersion).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/storm8studios/appstoreconnect-openapi-go/generated"
)

func main() {
    filterFrequency := []string{"FilterFrequency_example"} // []string | filter by attribute 'frequency'
    filterReportSubType := []string{"FilterReportSubType_example"} // []string | filter by attribute 'reportSubType'
    filterReportType := []string{"FilterReportType_example"} // []string | filter by attribute 'reportType'
    filterVendorNumber := []string{"Inner_example"} // []string | filter by attribute 'vendorNumber'
    filterReportDate := []string{"Inner_example"} // []string | filter by attribute 'reportDate' (optional)
    filterVersion := []string{"Inner_example"} // []string | filter by attribute 'version' (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SalesReportsApi.SalesReportsGetCollection(context.Background()).FilterFrequency(filterFrequency).FilterReportSubType(filterReportSubType).FilterReportType(filterReportType).FilterVendorNumber(filterVendorNumber).FilterReportDate(filterReportDate).FilterVersion(filterVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SalesReportsApi.SalesReportsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SalesReportsGetCollection`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SalesReportsApi.SalesReportsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSalesReportsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterFrequency** | **[]string** | filter by attribute &#39;frequency&#39; | 
 **filterReportSubType** | **[]string** | filter by attribute &#39;reportSubType&#39; | 
 **filterReportType** | **[]string** | filter by attribute &#39;reportType&#39; | 
 **filterVendorNumber** | **[]string** | filter by attribute &#39;vendorNumber&#39; | 
 **filterReportDate** | **[]string** | filter by attribute &#39;reportDate&#39; | 
 **filterVersion** | **[]string** | filter by attribute &#39;version&#39; | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/a-gzip, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

