// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// <ProgramSnippet>
using System.CommandLine.Builder;
using System.CommandLine.Parsing;
using Azure.Identity;
using GetUserClient.ApiClient;
using Microsoft.Kiota.Authentication.Azure;
using Microsoft.Kiota.Cli.Commons.Extensions;
using Microsoft.Kiota.Http.HttpClientLibrary;

var rootCommand = new GetUserApiClient().BuildRootCommand();
rootCommand.Description = "CLI description";

// Set up services
var builder = new CommandLineBuilder(rootCommand)
    .UseDefaults()
    .UseRequestAdapter(ic =>
    {
        // The auth provider will only authorize requests to
        // the allowed hosts, in this case Microsoft Graph
        var allowedHosts = new [] { "graph.microsoft.com" };
        var graphScopes = new [] { "User.Read" };
        var options = new DeviceCodeCredentialOptions
        {
            ClientId = "15efe9a3-d292-48e9-afa0-6747c61b925c",
            DeviceCodeCallback = (code, cancellation) =>
            {
                Console.WriteLine(code.Message);
                return Task.FromResult(0);
            },
        };
        var credential = new DeviceCodeCredential(options);

        var authProvider = new AzureIdentityAuthenticationProvider(credential, allowedHosts, scopes: graphScopes);
        var adapter = new HttpClientRequestAdapter(authProvider);
        adapter.BaseUrl = "https://graph.microsoft.com/v1.0";
        return adapter;
    }).RegisterCommonServices();

return await builder.Build().InvokeAsync(args);
// </ProgramSnippet>
