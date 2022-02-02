// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

using Microsoft.AspNetCore.OData;
using Microsoft.EntityFrameworkCore;
using Microsoft.Identity.Web;
using Microsoft.OData.ModelBuilder;
using Microsoft.OData.UriParser;
using ToDoApi.Models;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
builder.Services
    .AddMicrosoftIdentityWebApiAuthentication(builder.Configuration);

builder.Services.AddControllers().AddOData((opt) =>
{
    opt.Select().Filter().OrderBy().SetMaxTop(100).Count();
    opt.TimeZone = TimeZoneInfo.Utc;

    var odataBuilder = new ODataConventionModelBuilder();
    odataBuilder.EnableLowerCamelCase();
    odataBuilder.EntitySet<ToDoItem>("ToDoItems");
    opt.AddRouteComponents("", odataBuilder.GetEdmModel(),
        s => s.AddSingleton<ODataUriResolver>(
            sp => new UnqualifiedODataUriResolver
            {
                EnableCaseInsensitive = true,
                EnableNoDollarQueryOptions = true
            }
        )
    );
});

builder.Services.AddDbContext<ToDoContext>(opt =>
    opt.UseInMemoryDatabase("ToDoList"),
    contextLifetime: ServiceLifetime.Singleton);

var app = builder.Build();

// Configure the HTTP request pipeline.
app.UseHttpsRedirection();

app.UseRouting();

app.UseODataRouteDebug();

app.UseAuthentication();
app.UseAuthorization();

app.MapControllers();

// Enable buffering so the payload
// can be re-read for logging purposes
app.Use((context, next) => {
    context.Request.EnableBuffering();
    return next();
});

app.Run();

// Here so test project can access the Program class
public partial class Program { }
