// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

using Azure.Identity;
using GetUserClient.ApiClient;
using Microsoft.Kiota.Authentication.Azure;
using Microsoft.Kiota.Http.HttpClientLibrary;

var clientId = "YOUR_CLIENT_ID";

// The auth provider will only authorize requests to
// the allowed hosts, in this case Microsoft Graph
var allowedHosts = new [] { "graph.microsoft.com" };
var graphScopes = new [] { "User.Read" };

var credential = new DeviceCodeCredential((code, cancellation) =>
{
    Console.WriteLine(code.Message);
    return Task.FromResult(0);
},
clientId);

var authProvider = new AzureIdentityAuthenticationProvider(credential, allowedHosts, scopes: graphScopes);
var requestAdapter = new HttpClientRequestAdapter(authProvider);
var client = new GetUserApiClient(requestAdapter);

var me = await client.Me.GetAsync();
Console.WriteLine($"Hello {me.DisplayName}, your ID is {me.Id}");
