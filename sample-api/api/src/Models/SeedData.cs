// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

namespace ToDoApi.Models
{
    public class SeedData
    {
        public static Guid[] SeedItemIds =
        {
            new Guid("101bc1a3-2195-47bf-91e3-9b28f675fd86"),
            new Guid("3e9ba537-d8b3-4932-93c7-88f75253ce57"),
            new Guid("429924d3-3752-407f-abc2-9f7573ac4216"),
            new Guid("dae33a54-073c-412b-a8bc-6277ff709267"),
            new Guid("ec8cea19-95f7-4e3f-8dd0-08b2ba9cb839"),
        };

        public static List<ToDoItem> ToDoItems = new List<ToDoItem>
        {
            new ToDoItem
            {
                Id = SeedItemIds[0],
                Name = "Buy bread",
                Type = "shopping",
                IsComplete = false,
                DueDate = DateTimeOffset.UtcNow.AddDays(1),
                Priority = 2
            },
            new ToDoItem
            {
                Id = SeedItemIds[1],
                Name = "Submit expense report",
                Type = "work",
                IsComplete = false,
                DueDate = DateTimeOffset.UtcNow.AddDays(3),
                Priority = 1
            },
            new ToDoItem
            {
                Id = SeedItemIds[2],
                Name = "Pick up prescription",
                Type = "shopping",
                IsComplete = true,
                DueDate = DateTimeOffset.UtcNow.AddDays(-1),
                Priority = 1
            },
            new ToDoItem
            {
                Id = SeedItemIds[3],
                Name = "Complete code review",
                Type = "work",
                IsComplete = false,
                DueDate = DateTimeOffset.UtcNow.AddDays(2),
                Priority = 1
            },
            new ToDoItem
            {
                Id = SeedItemIds[4],
                Name = "Buy birthday gift for Lynne",
                Type = "shopping",
                IsComplete = true,
                DueDate = DateTimeOffset.UtcNow.AddDays(5),
                Priority = 3
            }
        };
    }
}
