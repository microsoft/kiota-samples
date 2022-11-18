// This Startup file is based on ASP.NET Core new project templates and is included
// as a starting point for DI registration and HTTP request processing pipeline configuration.
// This file will need updated according to the specific scenario of the application being upgraded.
// For more information on ASP.NET Core startup files, see https://docs.microsoft.com/aspnet/core/fundamentals/startup

using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.NewtonsoftJson;
using Microsoft.AspNetCore.Routing;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
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
        // services.AddControllersWithViews(ConfigureMvcOptions)
        //     // Newtonsoft.Json is added for compatibility reasons
        //     // The recommended approach is to use System.Text.Json for serialization
        //     // Visit the following link for more guidance about moving away from Newtonsoft.Json to System.Text.Json
        //     // https://docs.microsoft.com/dotnet/standard/serialization/system-text-json-migrate-from-newtonsoft-how-to
        //     .AddNewtonsoftJson(options =>
        //     {
        //         options.UseMemberCasing();
        //     });

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
        // app.MapControllers();
        // Enable buffering so the payload
        // can be re-read for logging purposes
        app.Use((context, next) => {
            context.Request.EnableBuffering();
            return next();
        });
        app.UseEndpoints(endpoints =>
        {
            endpoints.MapControllers();
            // endpoints.MapControllerRoute(
            //     name: "default",
            //     pattern: "{controller=Home}/{action=Index}/{id?}");
        });
    }

    private void ConfigureMvcOptions(MvcOptions mvcOptions)
    { 
    }
}
