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
        _client = _factory.WithWebHostBuilder(builder => {
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
    public async Task ShouldListToDoItems()
    {
        // Arrange
        _todoContext.ReSeedDatabase();
        var expectedReturnCount = SeedData.ToDoItems.Count;

        // Act
        var response = await _client.GetAsync("/ToDoItems");
        response.EnsureSuccessStatusCode();
        var items = await DeserializeResponseAsync<ODataCollection<ToDoItem>>(response);

        // Assert
        // GET should return 200
        Assert.Equal(System.Net.HttpStatusCode.OK, response.StatusCode);

        // GET with no params should return all initial items
        Assert.NotNull(items?.Value);
        Assert.Equal(expectedReturnCount, items!.Value.Count);
    }

    [Fact]
    public async Task ShouldReturnOnlySelectedFields()
    {
        // Arrange
        _todoContext.ReSeedDatabase();
        var expectedReturnCount = SeedData.ToDoItems.Count;
        // Default values from class
        var expectedPriority = -1;
        var expectedDueDate = default(DateTimeOffset);

        // Act
        var response = await _client.GetAsync("/ToDoItems?select=name,isComplete");
        response.EnsureSuccessStatusCode();
        var items = await DeserializeResponseAsync<ODataCollection<ToDoItem>>(response);

        // Assert
        // GET should return 200
        Assert.Equal(System.Net.HttpStatusCode.OK, response.StatusCode);

        // GET with no filter or top should return all initial items
        Assert.NotNull(items?.Value);
        Assert.Equal(expectedReturnCount, items!.Value.Count);

        // Each item should only have Name and IsComplete
        Assert.All(items.Value, item => {
            Assert.NotNull(item.Name);
            Assert.Equal(expectedPriority, item.Priority);
            Assert.Null(item.Type);
            Assert.Equal(expectedDueDate, item.DueDate);
        });
    }

    [Fact]
    public async Task ShouldFilterCorrectly()
    {
        // Arrange
        _todoContext.ReSeedDatabase();
        var type = "work";
        var expectedReturnCount = SeedData.ToDoItems.Where(i => i.Type == type).Count();

        // Act
        var response = await _client.GetAsync($"/ToDoItems?filter=type eq '{type}'");
        response.EnsureSuccessStatusCode();
        var items = await DeserializeResponseAsync<ODataCollection<ToDoItem>>(response);

        // Assert
        // GET should return 200
        Assert.Equal(System.Net.HttpStatusCode.OK, response.StatusCode);

        // GET with filter should return only those that match the filter
        Assert.NotNull(items?.Value);
        Assert.Equal(expectedReturnCount, items!.Value.Count);

        // Each item should have the expected Type
        Assert.All(items.Value, item => {
            Assert.Equal(type, item.Type);
        });
    }

    [Fact]
    public async Task SkipTokenShouldSkipCorrectly()
    {
        // Arrange
        _todoContext.ReSeedDatabase();
        var skipToken = 2;
        var expectedReturnCount = SeedData.ToDoItems.Count - skipToken;
        var expectedName = SeedData.ToDoItems[skipToken].Name;

        // Act
        var response = await _client.GetAsync($"/ToDoItems?skip={skipToken}");
        response.EnsureSuccessStatusCode();
        var items = await DeserializeResponseAsync<ODataCollection<ToDoItem>>(response);

        // Assert
        // GET should return 200
        Assert.Equal(System.Net.HttpStatusCode.OK, response.StatusCode);

        // GET with skip should return all items less the skip number
        Assert.NotNull(items?.Value);
        Assert.Equal(expectedReturnCount, items!.Value.Count);

        // The first item should be the item at the skip index
        Assert.Equal(expectedName, items.Value.First().Name);
    }

    [Fact]
    public async Task OrderByShouldSortCorrectly()
    {
        // Arrange
        _todoContext.ReSeedDatabase();
        var expectedReturnCount = SeedData.ToDoItems.Count;
        var sortedList = SeedData.ToDoItems.ToList();
        sortedList.Sort((x, y) =>
        {
            if (string.IsNullOrEmpty(x?.Name) && string.IsNullOrEmpty(y?.Name)) return 0;
            if (string.IsNullOrEmpty(x?.Name)) return -1;
            if (string.IsNullOrEmpty(y?.Name)) return 1;
            return x.Name.CompareTo(y.Name);
        });

        // Act
        var response = await _client.GetAsync($"/ToDoItems?orderby=name");
        response.EnsureSuccessStatusCode();
        var items = await DeserializeResponseAsync<ODataCollection<ToDoItem>>(response);

        // Assert
        // GET should return 200
        Assert.Equal(System.Net.HttpStatusCode.OK, response.StatusCode);

        // GET with orderby should return all items
        Assert.NotNull(items?.Value);
        Assert.Equal(expectedReturnCount, items!.Value.Count);

        // The items should be sorted
        for (int i = 0; i < items.Value.Count; i++)
        {
            Assert.Equal(sortedList[i].Name, items.Value[i].Name);
        }
    }

    [Fact]
    public async Task TopShouldLimitReturnCorrectly()
    {
        // Arrange
        _todoContext.ReSeedDatabase();
        var top = 2;
        var expectedReturnCount = top;

        // Act
        var response = await _client.GetAsync($"/ToDoItems?top={top}");
        response.EnsureSuccessStatusCode();
        var items = await DeserializeResponseAsync<ODataCollection<ToDoItem>>(response);

        // Assert
        // GET should return 200
        Assert.Equal(System.Net.HttpStatusCode.OK, response.StatusCode);

        // GET with top should only return the specified number of items
        Assert.NotNull(items?.Value);
        Assert.Equal(expectedReturnCount, items!.Value.Count);
    }

    [Fact]
    public async Task ShouldGetItemById()
    {
        // Arrange
        _todoContext.ReSeedDatabase();
        var itemId = SeedData.SeedItemIds[3];

        // Act
        var response = await _client.GetAsync($"ToDoItems/{itemId}?select=name,id");
        response.EnsureSuccessStatusCode();
        var item = await DeserializeResponseAsync<ToDoItem>(response);

        // Assert
        // GET should return 200
        Assert.Equal(System.Net.HttpStatusCode.OK, response.StatusCode);

        // GET by ID should return an item with matching ID
        Assert.NotNull(item);
        Assert.Equal(itemId, item!.Id);

        // The item returned should match the item at the corresponding index
        //Assert.Equal(SeedData.ToDoItems[(int)itemId - 1].Name, item.Name);
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
        var expectedDueDate = new DateTimeOffset(2022, 1, 20, 20, 0, 0, 0, new System.Globalization.GregorianCalendar(), new TimeSpan(0,0,0));

        var jsonItem = $"{{ \"name\": \"{expectedName}\", \"isComplete\": {expectedIsComplete.ToString().ToLower()}, \"type\": \"{expectedType}\", \"priority\": {expectedPriority}, \"dueDate\": \"{expectedDueDate.ToString("O")}\" }}";
        var content = new StringContent(jsonItem, System.Text.Encoding.UTF8, "application/json");

        // Act
        var response = await _client.PostAsync("/ToDoItems", content);
        response.EnsureSuccessStatusCode();
        var item = await DeserializeResponseAsync<ToDoItem>(response);

        // Assert
        // POST should return 201
        Assert.Equal(System.Net.HttpStatusCode.Created, response.StatusCode);

        // The item should be returned
        Assert.NotNull(item);

        // The item values should match
        Assert.Equal(expectedName, item!.Name);
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
