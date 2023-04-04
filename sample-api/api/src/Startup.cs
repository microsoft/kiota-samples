// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.OData;
using Microsoft.EntityFrameworkCore;
using Microsoft.Identity.Web;
using Microsoft.OData.ModelBuilder;
using Microsoft.OData.UriParser;
using ToDoApi.Models;

namespace ToDoApi;
public class Startup
{
    public Startup(IConfiguration configuration)
    {
        Configuration = configuration;
    }

    public IConfiguration Configuration { get; }

    // This method gets called by the runtime. Use this method to add services to the container.
    public void ConfigureServices(IServiceCollection services)
    {
        services.AddMicrosoftIdentityWebApiAuthentication(Configuration);
        services.AddControllers().AddOData((opt) =>
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
        app.UseODataRouteDebug();
        app.UseAuthentication();
        app.UseAuthorization();
        app.UseRequestDecompression();
        // Enable buffering so the payload
        // can be re-read for logging purposes
        app.Use((context, next) => {
            context.Request.EnableBuffering();
            return next();
        });
        app.UseEndpoints(endpoints =>
        {
            endpoints.MapControllers();
        });
    }

    private void ConfigureMvcOptions(MvcOptions mvcOptions)
    {
    }
}
