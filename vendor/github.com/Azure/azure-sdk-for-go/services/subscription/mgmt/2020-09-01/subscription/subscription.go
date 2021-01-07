package subscription

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// Client is the the subscription client
type Client struct {
	BaseClient
}

// NewClient creates an instance of the Client client.
func NewClient() Client {
	return NewClientWithBaseURI(DefaultBaseURI)
}

// NewClientWithBaseURI creates an instance of the Client client using a custom endpoint.  Use this when interacting
// with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewClientWithBaseURI(baseURI string) Client {
	return Client{NewWithBaseURI(baseURI)}
}

// Cancel the operation to cancel a subscription
// Parameters:
// subscriptionID - subscription Id.
func (client Client) Cancel(ctx context.Context, subscriptionID string) (result CanceledSubscriptionID, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Cancel")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CancelPreparer(ctx, subscriptionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.Client", "Cancel", nil, "Failure preparing request")
		return
	}

	resp, err := client.CancelSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "subscription.Client", "Cancel", resp, "Failure sending request")
		return
	}

	result, err = client.CancelResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.Client", "Cancel", resp, "Failure responding to request")
	}

	return
}

// CancelPreparer prepares the Cancel request.
func (client Client) CancelPreparer(ctx context.Context, subscriptionID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2020-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Subscription/cancel", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CancelSender sends the Cancel request. The method will close the
// http.Response Body if it receives an error.
func (client Client) CancelSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CancelResponder handles the response to the Cancel request. The method always
// closes the http.Response Body.
func (client Client) CancelResponder(resp *http.Response) (result CanceledSubscriptionID, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Enable the operation to enable a subscription
// Parameters:
// subscriptionID - subscription Id.
func (client Client) Enable(ctx context.Context, subscriptionID string) (result EnabledSubscriptionID, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Enable")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.EnablePreparer(ctx, subscriptionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.Client", "Enable", nil, "Failure preparing request")
		return
	}

	resp, err := client.EnableSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "subscription.Client", "Enable", resp, "Failure sending request")
		return
	}

	result, err = client.EnableResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.Client", "Enable", resp, "Failure responding to request")
	}

	return
}

// EnablePreparer prepares the Enable request.
func (client Client) EnablePreparer(ctx context.Context, subscriptionID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2020-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Subscription/enable", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// EnableSender sends the Enable request. The method will close the
// http.Response Body if it receives an error.
func (client Client) EnableSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// EnableResponder handles the response to the Enable request. The method always
// closes the http.Response Body.
func (client Client) EnableResponder(resp *http.Response) (result EnabledSubscriptionID, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Rename the operation to rename a subscription
// Parameters:
// subscriptionID - subscription Id.
// body - subscription Name
func (client Client) Rename(ctx context.Context, subscriptionID string, body Name) (result RenamedSubscriptionID, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Rename")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RenamePreparer(ctx, subscriptionID, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.Client", "Rename", nil, "Failure preparing request")
		return
	}

	resp, err := client.RenameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "subscription.Client", "Rename", resp, "Failure sending request")
		return
	}

	result, err = client.RenameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscription.Client", "Rename", resp, "Failure responding to request")
	}

	return
}

// RenamePreparer prepares the Rename request.
func (client Client) RenamePreparer(ctx context.Context, subscriptionID string, body Name) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2020-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Subscription/rename", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RenameSender sends the Rename request. The method will close the
// http.Response Body if it receives an error.
func (client Client) RenameSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// RenameResponder handles the response to the Rename request. The method always
// closes the http.Response Body.
func (client Client) RenameResponder(resp *http.Response) (result RenamedSubscriptionID, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}