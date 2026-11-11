// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

using System.Text.Json;
using Microsoft.AspNetCore.Authorization.Policy;
using Microsoft.AspNetCore.Mvc.Testing;
using Microsoft.AspNetCore.TestHost;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.DependencyInjection.Extensions;
using ToDoApi.Models;
using Xunit;

namespace ToDoApi.Tests;

public class ToDoItemsControllerTests : IClassFixture<WebApplicationFactory<Program>>
{
    private WebApplicationFactory<Program> _factory;
    private ToDoContext _todoContext;
    private HttpClient _client;

    public ToDoItemsControllerTests(WebApplicationFactory<Program> factory)
    {
        _factory = factory;
        _todoContext = factory.Services.GetRequiredService<ToDoContext>();
        _client = _factory.WithWebHostBuilder(builder =>
        {
            builder.ConfigureTestServices(services =>
            {
                // Disable authentication
                services.RemoveAll<IPolicyEvaluator>();
                services.AddSingleton<IPolicyEvaluator, DisableAuthenticationPolicyEvaluator>();
            });
        })
        .CreateClient();
    }

    [Fact]
    public async Task ShouldGenerateOpenApiDescription()
    {
        var response = await _client.GetAsync("/openapi/openapi.yaml");
        response.EnsureSuccessStatusCode();
        var description = await response.Content.ReadAsStringAsync();

        Assert.Equal("text/plain+yaml", response.Content.Headers.ContentType?.MediaType);
        Assert.Matches(@"^openapi: '?3\.1\.", description);
        Assert.Contains("/ToDoItems:", description);
        Assert.Contains("'/ToDoItems/{id}':", description);
        Assert.Contains("name: isComplete", description);
    }

    [Fact]
    public async Task ShouldListToDoItems()
    {
        // Arrange
        _todoContext.ReSeedDatabase();
        var expectedReturnCount = SeedData.ToDoItems.Count;

        // Act
        var response = await _client.GetAsync("/ToDoItems");
        response.EnsureSuccessStatusCode();
        var items = await DeserializeResponseAsync<List<ToDoItem>>(response);

        // Assert
        // GET should return 200
        Assert.Equal(System.Net.HttpStatusCode.OK, response.StatusCode);

        // GET with no params should return all initial items
        Assert.NotNull(items);
        Assert.Equal(expectedReturnCount, items!.Count);
    }

    [Theory]
    [InlineData(true)]
    [InlineData(false)]
    public async Task ShouldFilterByCompletionStatus(bool isComplete)
    {
        // Arrange
        _todoContext.ReSeedDatabase();
        var expectedReturnCount = SeedData.ToDoItems.Count(i => i.IsComplete == isComplete);

        // Act
        var response = await _client.GetAsync($"/ToDoItems?isComplete={isComplete.ToString().ToLowerInvariant()}");
        response.EnsureSuccessStatusCode();
        var items = await DeserializeResponseAsync<List<ToDoItem>>(response);

        // Assert
        // GET should return 200
        Assert.Equal(System.Net.HttpStatusCode.OK, response.StatusCode);

        // GET with filter should return only those that match the filter
        Assert.NotNull(items);
        Assert.Equal(expectedReturnCount, items!.Count);
        Assert.All(items, item => Assert.Equal(isComplete, item.IsComplete));
    }

    [Fact]
    public async Task ShouldGetItemById()
    {
        // Arrange
        _todoContext.ReSeedDatabase();
        var itemId = SeedData.SeedItemIds[3];

        // Act
        var response = await _client.GetAsync($"/ToDoItems/{itemId}");
        response.EnsureSuccessStatusCode();
        var item = await DeserializeResponseAsync<ToDoItem>(response);

        // Assert
        // GET should return 200
        Assert.Equal(System.Net.HttpStatusCode.OK, response.StatusCode);

        // GET by ID should return an item with matching ID
        Assert.NotNull(item);
        Assert.Equal(itemId, item!.Id);

        // The item returned should match the seeded item
        Assert.Equal(SeedData.ToDoItems[3].Name, item.Name);
    }

    [Fact]
    public async Task ShouldCreateNewItem()
    {
        // Arrange
        _todoContext.ReSeedDatabase();
        var expectedName = "Test item";
        var expectedIsComplete = true;
        var expectedType = "test";
        var expectedPriority = 2;
        var expectedDueDate = new DateTimeOffset(2022, 1, 20, 20, 0, 0, 0, new System.Globalization.GregorianCalendar(), new TimeSpan(0, 0, 0));

        var jsonItem = $"{{ \"name\": \"{expectedName}\", \"isComplete\": {expectedIsComplete.ToString().ToLower()}, \"type\": \"{expectedType}\", \"priority\": {expectedPriority}, \"dueDate\": \"{expectedDueDate.ToString("O")}\" }}";
        var content = new StringContent(jsonItem, System.Text.Encoding.UTF8, "application/json");

        // Act
        var response = await _client.PostAsync("/ToDoItems", content);
        response.EnsureSuccessStatusCode();
        var item = await DeserializeResponseAsync<ToDoItem>(response);

        // Assert
        // POST should return 201 with the location of the new item
        Assert.Equal(System.Net.HttpStatusCode.Created, response.StatusCode);
        Assert.NotNull(response.Headers.Location);

        // The item should be returned
        Assert.NotNull(item);

        // The item values should match
        Assert.NotEqual(Guid.Empty, item!.Id);
        Assert.EndsWith($"/ToDoItems/{item.Id}", response.Headers.Location!.ToString());
        Assert.Equal(expectedName, item.Name);
        Assert.Equal(expectedIsComplete, item.IsComplete);
        Assert.Equal(expectedType, item.Type);
        Assert.Equal(expectedPriority, item.Priority);
        Assert.Equal(expectedDueDate, item.DueDate);
    }

    [Fact]
    public async Task ShouldUpdateExistingItem()
    {
        // Arrange
        _todoContext.ReSeedDatabase();
        var itemIdIndex = 2;
        var itemToUpdateId = SeedData.SeedItemIds[itemIdIndex];
        var expectedName = "UPDATED NAME";
        var expectedIsComplete = SeedData.ToDoItems[itemIdIndex].IsComplete;
        var expectedType = SeedData.ToDoItems[itemIdIndex].Type;
        var expectedPriority = SeedData.ToDoItems[itemIdIndex].Priority;
        var expectedDueDate = SeedData.ToDoItems[itemIdIndex].DueDate;

        var jsonItem = $"{{ \"name\": \"{expectedName}\" }}";
        var content = new StringContent(jsonItem, System.Text.Encoding.UTF8, "application/json");

        // Act
        var updateResponse = await _client.PatchAsync($"/ToDoItems/{itemToUpdateId}", content);
        var getResponse = await _client.GetAsync($"/ToDoItems/{itemToUpdateId}");
        updateResponse.EnsureSuccessStatusCode();
        getResponse.EnsureSuccessStatusCode();
        var item = await DeserializeResponseAsync<ToDoItem>(getResponse);

        // Assert
        // PATCH should return 204
        Assert.Equal(System.Net.HttpStatusCode.NoContent, updateResponse.StatusCode);
        // GET should return 200
        Assert.Equal(System.Net.HttpStatusCode.OK, getResponse.StatusCode);

        // The item values should match
        Assert.Equal(expectedName, item!.Name);
        Assert.Equal(expectedIsComplete, item.IsComplete);
        Assert.Equal(expectedType, item.Type);
        Assert.Equal(expectedPriority, item.Priority);
        Assert.Equal(expectedDueDate, item.DueDate);
    }

    [Fact]
    public async Task ShouldDeleteExistingItem()
    {
        // Arrange
        _todoContext.ReSeedDatabase();
        var itemToDeleteId = SeedData.SeedItemIds[3];

        // Act
        var deleteResponse = await _client.DeleteAsync($"/ToDoItems/{itemToDeleteId}");
        var getResponse = await _client.GetAsync($"/ToDoItems/{itemToDeleteId}");

        // Assert
        // DELETE should return 204
        Assert.Equal(System.Net.HttpStatusCode.NoContent, deleteResponse.StatusCode);

        // GET should return 404
        Assert.Equal(System.Net.HttpStatusCode.NotFound, getResponse.StatusCode);
    }

    private async Task<T?> DeserializeResponseAsync<T>(HttpResponseMessage response)
    {
        var body = await response.Content.ReadAsStringAsync();

        try
        {
            return JsonSerializer.Deserialize<T>(body);
        }
        catch (JsonException)
        {
            return default(T);
        }
    }
}
