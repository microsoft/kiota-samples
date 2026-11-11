using GitHub.ApiClient;
using Microsoft.Kiota.Abstractions.Authentication;
using Microsoft.Kiota.Http.HttpClientLibrary;

namespace GitHub;

public class GitHubClientFactory(HttpClient httpClient)
{
    private readonly IAuthenticationProvider _authenticationProvider = new AnonymousAuthenticationProvider();

    public GitHubClient GetClient() =>
        new(new HttpClientRequestAdapter(_authenticationProvider, httpClient: httpClient));
}
