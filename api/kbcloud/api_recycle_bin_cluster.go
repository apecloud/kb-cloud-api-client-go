// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package kbcloud

import (
	"bytes"
	_context "context"
	_fmt "fmt"
	_io "io"
	_log "log"
	_nethttp "net/http"
	_neturl "net/url"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// RecycleBinClusterApi service type
type RecycleBinClusterApi common.Service

// DeleteRecycleBinCluster Delete cluster from the Recycle Bin of the Org.
func (a *RecycleBinClusterApi) DeleteRecycleBinCluster(ctx _context.Context, orgName string, clusterName string, isDeleteBackup bool) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
	)

    

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".RecycleBinClusterApi.DeleteRecycleBinCluster")
	if err != nil {
		return nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/recycleBin/clusters/{clusterName}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarQueryParams.Add("isDeleteBackup", common.ParameterToString(isDeleteBackup, ""))
	localVarHeaderParams["Accept"] = "application/json"

	
        common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"BearerToken", "authorization"},
	)
	req, err := a.Client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := common.ReadBody(localVarHTTPResponse)
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := common.GenericOpenAPIError{
			ErrorBody:  localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if
		localVarHTTPResponse.StatusCode == 401||localVarHTTPResponse.StatusCode == 403||localVarHTTPResponse.StatusCode == 404{
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

// GetRecycleBinCluster Get cluster in the Recycle Bin of the Org.
func (a *RecycleBinClusterApi) GetRecycleBinCluster(ctx _context.Context, orgName string, clusterName string) (RecycleBinCluster, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarReturnValue  RecycleBinCluster
	)

    

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".RecycleBinClusterApi.GetRecycleBinCluster")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/recycleBin/clusters/{clusterName}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "application/json"

	
        common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"BearerToken", "authorization"},
	)
	req, err := a.Client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := common.ReadBody(localVarHTTPResponse)
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := common.GenericOpenAPIError{
			ErrorBody:  localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if
		localVarHTTPResponse.StatusCode == 401||localVarHTTPResponse.StatusCode == 403||localVarHTTPResponse.StatusCode == 404{
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := common.GenericOpenAPIError{
			ErrorBody:  localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// ListRecycleBinCluster List clusters in the Recycle Bin of the Org.
func (a *RecycleBinClusterApi) ListRecycleBinCluster(ctx _context.Context, orgName string) (RecycleBinClusterList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarReturnValue  RecycleBinClusterList
	)

    

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".RecycleBinClusterApi.ListRecycleBinCluster")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/recycleBin/clusters"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "application/json"

	
        common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"BearerToken", "authorization"},
	)
	req, err := a.Client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := common.ReadBody(localVarHTTPResponse)
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := common.GenericOpenAPIError{
			ErrorBody:  localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if
		localVarHTTPResponse.StatusCode == 401||localVarHTTPResponse.StatusCode == 403||localVarHTTPResponse.StatusCode == 404{
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := common.GenericOpenAPIError{
			ErrorBody:  localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// RestoreRecycleBinCluster Restore cluster from the Recycle Bin of the Org.
func (a *RecycleBinClusterApi) RestoreRecycleBinCluster(ctx _context.Context, orgName string, clusterName string) (RecycleBinCluster, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarReturnValue  RecycleBinCluster
	)

    

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(ctx, ".RecycleBinClusterApi.RestoreRecycleBinCluster")
	if err != nil {
		return localVarReturnValue, nil, common.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/organizations/{orgName}/recycleBin/clusters/{clusterName}/restore"
	localVarPath = strings.Replace(localVarPath, "{"+"orgName"+"}", _neturl.PathEscape(common.ParameterToString(orgName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.PathEscape(common.ParameterToString(clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	localVarHeaderParams["Accept"] = "application/json"

	
        common.SetAuthKeys(
		ctx,
		&localVarHeaderParams,
		[2]string{"BearerToken", "authorization"},
	)
	req, err := a.Client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := common.ReadBody(localVarHTTPResponse)
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := common.GenericOpenAPIError{
			ErrorBody:  localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if
		localVarHTTPResponse.StatusCode == 401||localVarHTTPResponse.StatusCode == 403||localVarHTTPResponse.StatusCode == 404{
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := common.GenericOpenAPIError{
			ErrorBody:  localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// NewRecycleBinClusterApi Returns NewRecycleBinClusterApi.
func NewRecycleBinClusterApi(client *common.APIClient) *RecycleBinClusterApi {
	return &RecycleBinClusterApi{
		Client: client,
	}
}
