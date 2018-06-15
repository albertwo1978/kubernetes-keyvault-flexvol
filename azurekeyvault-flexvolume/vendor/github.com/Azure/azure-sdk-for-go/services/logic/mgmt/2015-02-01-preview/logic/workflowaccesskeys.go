package logic

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-02-01-preview/logic instead.
// WorkflowAccessKeysClient is the REST API for Azure Logic Apps.
type WorkflowAccessKeysClient struct {
	BaseClient
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-02-01-preview/logic instead.
// NewWorkflowAccessKeysClient creates an instance of the WorkflowAccessKeysClient client.
func NewWorkflowAccessKeysClient(subscriptionID string) WorkflowAccessKeysClient {
	return NewWorkflowAccessKeysClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-02-01-preview/logic instead.
// NewWorkflowAccessKeysClientWithBaseURI creates an instance of the WorkflowAccessKeysClient client.
func NewWorkflowAccessKeysClientWithBaseURI(baseURI string, subscriptionID string) WorkflowAccessKeysClient {
	return WorkflowAccessKeysClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-02-01-preview/logic instead.
// CreateOrUpdate creates or updates a workflow access key.
//
// resourceGroupName is the resource group name. workflowName is the workflow name. accessKeyName is the workflow
// access key name. workflowAccesskey is the workflow access key.
func (client WorkflowAccessKeysClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workflowName string, accessKeyName string, workflowAccesskey WorkflowAccessKey) (result WorkflowAccessKey, err error) {
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, workflowName, accessKeyName, workflowAccesskey)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-02-01-preview/logic instead.
// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client WorkflowAccessKeysClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, workflowName string, accessKeyName string, workflowAccesskey WorkflowAccessKey) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accessKeyName":     autorest.Encode("path", accessKeyName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workflowName":      autorest.Encode("path", workflowName),
	}

	const APIVersion = "2015-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/accessKeys/{accessKeyName}", pathParameters),
		autorest.WithJSON(workflowAccesskey),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-02-01-preview/logic instead.
// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowAccessKeysClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-02-01-preview/logic instead.
// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client WorkflowAccessKeysClient) CreateOrUpdateResponder(resp *http.Response) (result WorkflowAccessKey, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-02-01-preview/logic instead.
// Delete deletes a workflow access key.
//
// resourceGroupName is the resource group name. workflowName is the workflow name. accessKeyName is the workflow
// access key name.
func (client WorkflowAccessKeysClient) Delete(ctx context.Context, resourceGroupName string, workflowName string, accessKeyName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(ctx, resourceGroupName, workflowName, accessKeyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-02-01-preview/logic instead.
// DeletePreparer prepares the Delete request.
func (client WorkflowAccessKeysClient) DeletePreparer(ctx context.Context, resourceGroupName string, workflowName string, accessKeyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accessKeyName":     autorest.Encode("path", accessKeyName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workflowName":      autorest.Encode("path", workflowName),
	}

	const APIVersion = "2015-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/accessKeys/{accessKeyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-02-01-preview/logic instead.
// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowAccessKeysClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-02-01-preview/logic instead.
// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client WorkflowAccessKeysClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-02-01-preview/logic instead.
// Get gets a workflow access key.
//
// resourceGroupName is the resource group name. workflowName is the workflow name. accessKeyName is the workflow
// access key name.
func (client WorkflowAccessKeysClient) Get(ctx context.Context, resourceGroupName string, workflowName string, accessKeyName string) (result WorkflowAccessKey, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, workflowName, accessKeyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "Get", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-02-01-preview/logic instead.
// GetPreparer prepares the Get request.
func (client WorkflowAccessKeysClient) GetPreparer(ctx context.Context, resourceGroupName string, workflowName string, accessKeyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accessKeyName":     autorest.Encode("path", accessKeyName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workflowName":      autorest.Encode("path", workflowName),
	}

	const APIVersion = "2015-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/accessKeys/{accessKeyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-02-01-preview/logic instead.
// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowAccessKeysClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-02-01-preview/logic instead.
// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client WorkflowAccessKeysClient) GetResponder(resp *http.Response) (result WorkflowAccessKey, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-02-01-preview/logic instead.
// List gets a list of workflow access keys.
//
// resourceGroupName is the resource group name. workflowName is the workflow name. top is the number of items to
// be included in the result.
func (client WorkflowAccessKeysClient) List(ctx context.Context, resourceGroupName string, workflowName string, top *int32) (result WorkflowAccessKeyListResultPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, workflowName, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.waklr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "List", resp, "Failure sending request")
		return
	}

	result.waklr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "List", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-02-01-preview/logic instead.
// ListPreparer prepares the List request.
func (client WorkflowAccessKeysClient) ListPreparer(ctx context.Context, resourceGroupName string, workflowName string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workflowName":      autorest.Encode("path", workflowName),
	}

	const APIVersion = "2015-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/accessKeys", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-02-01-preview/logic instead.
// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowAccessKeysClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-02-01-preview/logic instead.
// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client WorkflowAccessKeysClient) ListResponder(resp *http.Response) (result WorkflowAccessKeyListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client WorkflowAccessKeysClient) listNextResults(lastResults WorkflowAccessKeyListResult) (result WorkflowAccessKeyListResult, err error) {
	req, err := lastResults.workflowAccessKeyListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-02-01-preview/logic instead.
// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client WorkflowAccessKeysClient) ListComplete(ctx context.Context, resourceGroupName string, workflowName string, top *int32) (result WorkflowAccessKeyListResultIterator, err error) {
	result.page, err = client.List(ctx, resourceGroupName, workflowName, top)
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-02-01-preview/logic instead.
// ListSecretKeys lists secret keys.
//
// resourceGroupName is the resource group name. workflowName is the workflow name. accessKeyName is the workflow
// access key name.
func (client WorkflowAccessKeysClient) ListSecretKeys(ctx context.Context, resourceGroupName string, workflowName string, accessKeyName string) (result WorkflowSecretKeys, err error) {
	req, err := client.ListSecretKeysPreparer(ctx, resourceGroupName, workflowName, accessKeyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "ListSecretKeys", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSecretKeysSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "ListSecretKeys", resp, "Failure sending request")
		return
	}

	result, err = client.ListSecretKeysResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "ListSecretKeys", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-02-01-preview/logic instead.
// ListSecretKeysPreparer prepares the ListSecretKeys request.
func (client WorkflowAccessKeysClient) ListSecretKeysPreparer(ctx context.Context, resourceGroupName string, workflowName string, accessKeyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accessKeyName":     autorest.Encode("path", accessKeyName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workflowName":      autorest.Encode("path", workflowName),
	}

	const APIVersion = "2015-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/accessKeys/{accessKeyName}/list", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-02-01-preview/logic instead.
// ListSecretKeysSender sends the ListSecretKeys request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowAccessKeysClient) ListSecretKeysSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-02-01-preview/logic instead.
// ListSecretKeysResponder handles the response to the ListSecretKeys request. The method always
// closes the http.Response Body.
func (client WorkflowAccessKeysClient) ListSecretKeysResponder(resp *http.Response) (result WorkflowSecretKeys, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-02-01-preview/logic instead.
// RegenerateSecretKey regenerates secret key.
//
// resourceGroupName is the resource group name. workflowName is the workflow name. accessKeyName is the workflow
// access key name. parameters is the parameters.
func (client WorkflowAccessKeysClient) RegenerateSecretKey(ctx context.Context, resourceGroupName string, workflowName string, accessKeyName string, parameters RegenerateSecretKeyParameters) (result WorkflowSecretKeys, err error) {
	req, err := client.RegenerateSecretKeyPreparer(ctx, resourceGroupName, workflowName, accessKeyName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "RegenerateSecretKey", nil, "Failure preparing request")
		return
	}

	resp, err := client.RegenerateSecretKeySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "RegenerateSecretKey", resp, "Failure sending request")
		return
	}

	result, err = client.RegenerateSecretKeyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.WorkflowAccessKeysClient", "RegenerateSecretKey", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-02-01-preview/logic instead.
// RegenerateSecretKeyPreparer prepares the RegenerateSecretKey request.
func (client WorkflowAccessKeysClient) RegenerateSecretKeyPreparer(ctx context.Context, resourceGroupName string, workflowName string, accessKeyName string, parameters RegenerateSecretKeyParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accessKeyName":     autorest.Encode("path", accessKeyName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workflowName":      autorest.Encode("path", workflowName),
	}

	const APIVersion = "2015-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/accessKeys/{accessKeyName}/regenerate", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-02-01-preview/logic instead.
// RegenerateSecretKeySender sends the RegenerateSecretKey request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowAccessKeysClient) RegenerateSecretKeySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/logic/mgmt/2015-02-01-preview/logic instead.
// RegenerateSecretKeyResponder handles the response to the RegenerateSecretKey request. The method always
// closes the http.Response Body.
func (client WorkflowAccessKeysClient) RegenerateSecretKeyResponder(resp *http.Response) (result WorkflowSecretKeys, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
