// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// <ProgramSnippet>
using Azure.Identity;
using GetUserClient.ApiClient;
using Microsoft.Kiota.Authentication.Azure;
using Microsoft.Kiota.Bundle;

// The auth provider will only authorize requests to
// the allowed hosts, in this case Microsoft Graph
var allowedHosts = new[] { "graph.microsoft.com" };
var graphScopes = new[] { "User.Read" };

var options = new DeviceCodeCredentialOptions
{
    ClientId = "YOUR_CLIENT_ID",
    DeviceCodeCallback = (code, cancellation) =>
    {
        Console.WriteLine(code.Message);
        return Task.FromResult(0);
    },
};

var credential = new DeviceCodeCredential(options);

var authProvider = new AzureIdentityAuthenticationProvider(credential, allowedHosts, scopes: graphScopes);
var requestAdapter = new DefaultRequestAdapter(authProvider);
var client = new GetUserApiClient(requestAdapter);

var me = await client.Me.GetAsync();
Console.WriteLine($"Hello {me.DisplayName}, your ID is {me.Id}");
// </ProgramSnippet>
