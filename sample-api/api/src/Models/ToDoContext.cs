// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

using Microsoft.EntityFrameworkCore;

namespace ToDoApi.Models
{
    public class ToDoContext : DbContext
    {
        public ToDoContext(DbContextOptions<ToDoContext> options)
            : base(options)
        {
            Database.EnsureCreated();
        }

        public DbSet<ToDoItem> ToDoItems { get; set; } = null!;

        protected override void OnModelCreating(ModelBuilder builder)
        {
            builder.Entity<ToDoItem>().HasData(SeedData.ToDoItems);
        }

        // Used for unit tests to reset database to known values
        internal void ReSeedDatabase()
        {
            Database.EnsureDeleted();
            SaveChanges();
            Database.EnsureCreated();
        }
    }
}
