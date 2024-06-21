using GitHub.ApiClient;
using Microsoft.Kiota.Abstractions.Authentication;
using Microsoft.Kiota.Http.HttpClientLibrary;

namespace GitHub;

public class GitHubClientFactory
{
    private readonly IAuthenticationProvider _authenticationProvider;
    private readonly HttpClient _httpClient;

    public GitHubClientFactory(HttpClient httpClient)
    {
        _authenticationProvider = new AnonymousAuthenticationProvider();
        _httpClient = httpClient;
    }

    public GitHubClient GetClient() {
      return new GitHubClient(new HttpClientRequestAdapter(_authenticationProvider, httpClient: _httpClient));
    }
}