// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

using System.Text.Json.Serialization;

namespace ToDoApi.Models
{
    /// <summary>
    /// Partial update for a <see cref="ToDoItem"/>. Properties left unset are not modified.
    /// </summary>
    public class ToDoItemUpdate
    {
        [JsonPropertyName("name")]
        public string? Name { get; set; }
        [JsonPropertyName("isComplete")]
        public bool? IsComplete { get; set; }
        [JsonPropertyName("type")]
        public string? Type { get; set; }
        [JsonPropertyName("priority")]
        public int? Priority { get; set; }
        [JsonPropertyName("dueDate")]
        public DateTimeOffset? DueDate { get; set; }

        public void ApplyTo(ToDoItem item)
        {
            ArgumentNullException.ThrowIfNull(item);
            if (Name is not null) item.Name = Name;
            if (IsComplete.HasValue) item.IsComplete = IsComplete.Value;
            if (Type is not null) item.Type = Type;
            if (Priority.HasValue) item.Priority = Priority.Value;
            if (DueDate.HasValue) item.DueDate = DueDate.Value;
        }
    }
}
