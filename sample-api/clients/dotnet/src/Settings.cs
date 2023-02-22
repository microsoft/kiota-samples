// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

using Microsoft.Extensions.Configuration;

public class Settings
{
    private static string ClientIdPlaceholder = "YOUR_CLIENT_ID";
    private static string TenantIdPlaceholder = "YOUR_TENANT_ID";
    private static string ApiScopePlaceholder = "api://YOUR_TODO_API_CLIENT_ID/ToDoItem.ReadWrite";

    public string? ClientId { get; set; }
    public string? TenantId { get; set; }
    public string? ApiScope { get; set; }

    public static Settings? LoadSettings()
    {
        IConfiguration config = new ConfigurationBuilder()
            .AddJsonFile("appsettings.json", optional: false)
            .AddJsonFile($"appsettings.local.json", optional: true)
            .Build();

        return config.GetRequiredSection("Settings").Get<Settings>();
    }

    public bool Validate()
    {
        return !string.IsNullOrEmpty(ClientId) &&
               string.Compare(ClientIdPlaceholder, ClientId) != 0 &&
               !string.IsNullOrEmpty(TenantId) &&
               string.Compare(TenantIdPlaceholder, TenantId) != 0 &&
               !string.IsNullOrEmpty(ApiScope) &&
               string.Compare(ApiScopePlaceholder, ApiScope) != 0;
    }
}
