// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

using Azure.Core;
using Azure.Identity;
using Microsoft.Kiota.Authentication.Azure;
using Microsoft.Kiota.Http.HttpClientLibrary;
using ToDoClient.ApiClient;
using ToDoClient.ApiClient.Models.ToDoApi.Models;

var settings = Settings.LoadSettings();
if (!settings.Validate())
{
    Console.WriteLine("Missing or invalid configuration values in appsettings.json.");
    Console.WriteLine("See README.md for more information.");
    return;
}

var scopes = new[] { settings.ApiScope ?? "" };
var credential = new DeviceCodeCredential((code, cancellation) =>
{
    Console.WriteLine(code.Message);
    return Task.FromResult(0);
},
settings.TenantId,
settings.ClientId);

var context = new TokenRequestContext(scopes);
var response = await credential.GetTokenAsync(context);

Console.WriteLine($"Token: {response.Token}");

var authProvider = new AzureIdentityAuthenticationProvider(credential, scopes);
var requestAdapter = new HttpClientRequestAdapter(authProvider);

var client = new ToDoApiClient(requestAdapter);

try
{
    // GET /ToDoItems
    var items = await client.ToDoItems.GetAsync();

    Console.WriteLine($"Found {items.Value.Count} items");
    foreach (var item in items.Value)
    {
        Console.WriteLine($"TODO: {item.Name}, Due: {item.DueDate.ToString()}");
    }

    // Select
    var select = await client.ToDoItems.GetAsync(c =>
    {
        c.QueryParameters.Select = new [] { "name", "priority", "isComplete" };
    });
    foreach (var item in select.Value)
    {
        // Name, priority, and isComplete should have values
        Console.WriteLine($"{item.Name} - Priority: {item.Priority} - Complete: {((item.IsComplete ?? false) ? "yes" : "no")}");
    }

    // Get by ID
    var specificId = items.Value[3].Id;
    var specificItem = await client.ToDoItems[specificId].GetAsync();
    Console.WriteLine(specificItem.Name ?? "No name");

    // Create an item
    var newItem = new ToDoItem
    {
        Name = "Start using Kiota generators in my project",
        Type = "work",
        IsComplete = false,
        DueDate = DateTimeOffset.UtcNow.AddDays(7),
        Priority = 1
    };

    var createdItem = await client.ToDoItems.PostAsync(newItem);
    Console.WriteLine($"Item created with ID: {createdItem.Id}");

    // Update an item
    var updateItem = new ToDoItem
    {
        IsComplete = true
    };

    await client.ToDoItems[$"{createdItem.Id}"].PatchAsync(updateItem);
    Console.WriteLine("Updated item");

    // Delete an item
    await client.ToDoItems[$"{createdItem.Id}"].DeleteAsync();
    Console.WriteLine("Deleted item");

    // Confirm item is deleted
    var itemsAfterDelete = await client.ToDoItems.GetAsync();
    if (itemsAfterDelete.Value.Any(i => i.Id == createdItem.Id))
    {
        Console.WriteLine("Unexpected: item still exists!");
    }
    else
    {
        Console.WriteLine("Item not present in API response");
    }
}
catch (Exception ex)
{
    Console.WriteLine($"Caught exception: {ex.Message}");
    Console.WriteLine(ex.ToString());
}
