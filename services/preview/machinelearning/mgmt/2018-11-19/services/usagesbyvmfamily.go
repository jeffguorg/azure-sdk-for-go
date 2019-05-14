package services

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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// UsagesByVMFamilyClient is the these APIs allow end users to operate on Azure Machine Learning Workspace resources.
type UsagesByVMFamilyClient struct {
	BaseClient
}

// NewUsagesByVMFamilyClient creates an instance of the UsagesByVMFamilyClient client.
func NewUsagesByVMFamilyClient(subscriptionID string) UsagesByVMFamilyClient {
	return NewUsagesByVMFamilyClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewUsagesByVMFamilyClientWithBaseURI creates an instance of the UsagesByVMFamilyClient client.
func NewUsagesByVMFamilyClientWithBaseURI(baseURI string, subscriptionID string) UsagesByVMFamilyClient {
	return UsagesByVMFamilyClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List gets the current usage information in a detailed format as well as limits for Batch AI resources for given
// subscription, by VM family, workspace and cluster hierarchy.
// Parameters:
// location - the location for which resource usage is queried.
func (client UsagesByVMFamilyClient) List(ctx context.Context, location string) (result ListUsagesByVMFamilyResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsagesByVMFamilyClient.List")
		defer func() {
			sc := -1
			if result.lubvfr.Response.Response != nil {
				sc = result.lubvfr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: location,
			Constraints: []validation.Constraint{{Target: "location", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("services.UsagesByVMFamilyClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, location)
	if err != nil {
		err = autorest.NewErrorWithError(err, "services.UsagesByVMFamilyClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.lubvfr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "services.UsagesByVMFamilyClient", "List", resp, "Failure sending request")
		return
	}

	result.lubvfr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "services.UsagesByVMFamilyClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client UsagesByVMFamilyClient) ListPreparer(ctx context.Context, location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-11-19"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearningServices/locations/{location}/usagesByVMFamily", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client UsagesByVMFamilyClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client UsagesByVMFamilyClient) ListResponder(resp *http.Response) (result ListUsagesByVMFamilyResult, err error) {
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
func (client UsagesByVMFamilyClient) listNextResults(ctx context.Context, lastResults ListUsagesByVMFamilyResult) (result ListUsagesByVMFamilyResult, err error) {
	req, err := lastResults.listUsagesByVMFamilyResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "services.UsagesByVMFamilyClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "services.UsagesByVMFamilyClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "services.UsagesByVMFamilyClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client UsagesByVMFamilyClient) ListComplete(ctx context.Context, location string) (result ListUsagesByVMFamilyResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsagesByVMFamilyClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, location)
	return
}
