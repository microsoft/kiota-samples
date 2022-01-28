// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

using System.Text.Json.Serialization;

namespace ToDoApi.Models
{
    public class ToDoItem
    {
        [JsonPropertyName("id")]
        public Guid Id { get; set; }
        [JsonPropertyName("name")]
        public string? Name { get; set; }
        [JsonPropertyName("isComplete")]
        public bool IsComplete { get; set; }
        [JsonPropertyName("type")]
        public string? Type { get; set; }
        [JsonPropertyName("priority")]
        public int Priority { get; set; } = -1;
        [JsonPropertyName("dueDate")]
        public DateTimeOffset DueDate { get; set; }
    }
}
