# Kiota with .NET Dependency Injection

This project is the finished product of tutorial [Kiota with .NET Dependency Injection](https://learn.microsoft.com/openapi/kiota/tutorials/dotnet-dependency-injection?tabs=portal).

## Parts of the sample

- `KiotaServiceCollectionExtensions.cs` - Extension methods to add Kiota services to the service collection.
- `GitHub/GitHubClientFactory.cs` - Factory class to create GitHub clients.
- `Program.cs` - Registring the kiota client with all its handlers.

## Generate the client

To generate the client, run the following command:

```bash
kiota generate -l csharp -d github-releases.yml -c GitHubClient -n GitHub.ApiClient -o ./GitHub
```

## Run the sample

To run the sample, run the following command:

```bash
dotnet run
```

And navigate to `https://localhost:{port}/swagger` in your browser.
