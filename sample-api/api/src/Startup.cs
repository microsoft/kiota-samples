// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

using System.Text.Json.Serialization;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using Microsoft.Identity.Web;
using Microsoft.OpenApi;
using ToDoApi.Models;

namespace ToDoApi;

public class Startup
{
    private const string OpenApiDocumentName = "openapi";

    public Startup(IConfiguration configuration)
    {
        Configuration = configuration;
    }

    public IConfiguration Configuration { get; }

    // This method gets called by the runtime. Use this method to add services to the container.
    public void ConfigureServices(IServiceCollection services)
    {
        services.AddMicrosoftIdentityWebApiAuthentication(Configuration);
        services.AddControllers()
            .AddJsonOptions(options => options.JsonSerializerOptions.NumberHandling = JsonNumberHandling.Strict);
        services.ConfigureHttpJsonOptions(options => options.SerializerOptions.NumberHandling = JsonNumberHandling.Strict);
        services.AddOpenApi(OpenApiDocumentName, options =>
        {
            options.OpenApiVersion = OpenApiSpecVersion.OpenApi3_1;
            options.AddDocumentTransformer((document, context, cancellationToken) =>
            {
                document.Info.Title = "ToDo API";
                document.Info.Version = "v1";
                return Task.CompletedTask;
            });
        });
        services.AddDbContext<ToDoContext>(opt =>
            opt.UseInMemoryDatabase("ToDoList"),
            contextLifetime: ServiceLifetime.Singleton);

        services.AddRequestDecompression();
    }

    // This method gets called by the runtime. Use this method to configure the HTTP request pipeline.
    public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
    {
        if (env.IsDevelopment())
        {
            app.UseDeveloperExceptionPage();
        }

        app.UseHttpsRedirection();
        app.UseRouting();
        app.UseAuthentication();
        app.UseAuthorization();
        app.UseRequestDecompression();
        // Enable buffering so the payload
        // can be re-read for logging purposes
        app.Use((context, next) =>
        {
            context.Request.EnableBuffering();
            return next();
        });
        app.UseEndpoints(endpoints =>
        {
            endpoints.MapControllers();
            endpoints.MapOpenApi("/openapi/{documentName}.yaml");
        });
    }

    private void ConfigureMvcOptions(MvcOptions mvcOptions)
    {
    }
}
