using GitHub;
var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

// ----------- Add this part to register the generated client -----------
// Add Kiota handlers to the dependency injection container
builder.Services.AddKiotaHandlers();

// Register the factory for the GitHub client
builder.Services.AddHttpClient<GitHubClientFactory>((sp, client) => {
    // Set the base address and accept header
    // or other settings on the http client
    client.BaseAddress = new Uri("https://api.github.com");
    client.DefaultRequestHeaders.Add("Accept", "application/vnd.github.v3+json");
}).AttachKiotaHandlers(); // Attach the Kiota handlers to the http client, this is to enable all the Kiota features.

// Register the GitHub client
builder.Services.AddTransient(sp => sp.GetRequiredService<GitHubClientFactory>().GetClient());
// ----------- Add this part to register the generated client end -------

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseHttpsRedirection();

var summaries = new[]
{
    "Freezing", "Bracing", "Chilly", "Cool", "Mild", "Warm", "Balmy", "Hot", "Sweltering", "Scorching"
};

app.MapGet("/weatherforecast", () =>
{
    var forecast =  Enumerable.Range(1, 5).Select(index =>
        new WeatherForecast
        (
            DateOnly.FromDateTime(DateTime.Now.AddDays(index)),
            Random.Shared.Next(-20, 55),
            summaries[Random.Shared.Next(summaries.Length)]
        ))
        .ToArray();
    return forecast;
})
.WithName("GetWeatherForecast")
.WithOpenApi();

// ----------- Add this part to create a new endpoint that uses the generated client -----------

app.MapGet("/dotnet/releases", async (GitHub.ApiClient.GitHubClient client, CancellationToken cancellationToken) =>
{
    var releases = await client.Repos["dotnet"]["runtime"].Releases["latest"].GetAsync(cancellationToken: cancellationToken);
    return releases;
})
.WithName("GetDotnetReleases")
.WithOpenApi();

// ----------- Add this part to create a new endpoint that uses the generated client end -----------

app.Run();

record WeatherForecast(DateOnly Date, int TemperatureC, string? Summary)
{
    public int TemperatureF => 32 + (int)(TemperatureC / 0.5556);
}
