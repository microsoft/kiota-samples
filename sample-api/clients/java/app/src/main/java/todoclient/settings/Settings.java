// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package todoclient.settings;

import java.io.IOException;
import java.util.Properties;

public class Settings {
    private static final String PROPERTIES_FILE = "azure.properties";
    private static final String CLIENT_ID_KEY = "azure.clientId";
    private static final String CLIENT_ID_PLACEHOLDER = "YOUR_CLIENT_ID";
    private static final String TENANT_ID_KEY = "azure.tenantId";
    private static final String TENANT_ID_PLACEHOLDER = "YOUR_TENANT_ID";
    private static final String SCOPES_KEY = "azure.scopes";
    private static final String SCOPES_PLACEHOLDER = "api://YOUR_TODO_API_CLIENT_ID/ToDoItem.ReadWrite";

    private String _clientId;
    private String _tenantId;
    private String[] _apiScopes;

    public String getClientId() {
        return this._clientId;
    }

    public String getTenantId() {
        return this._tenantId;
    }

    public String[] getApiScopes() {
        return this._apiScopes;
    }

    public static Settings loadSettings() throws IOException {
        final Properties azureProperties = new Properties();

        azureProperties.load(Settings.class.getResourceAsStream(PROPERTIES_FILE));

        return new Settings(azureProperties);
    }

    public boolean Validate()
    {
        return !IsNullOrBlankOrPlaceholder(this._clientId, CLIENT_ID_PLACEHOLDER) &&
               !IsNullOrBlankOrPlaceholder(this._tenantId, TENANT_ID_PLACEHOLDER) &&
               !IsNullOrBlankOrPlaceholder(this._apiScopes, SCOPES_PLACEHOLDER);
    }

    private boolean IsNullOrBlankOrPlaceholder(String value, String placeholder)
    {
        return value == null ||
               value.isBlank() ||
               value.contains(placeholder);
    }

    private boolean IsNullOrBlankOrPlaceholder(String[] value, String placeholder)
    {
        if (value == null || value.length <= 0) return true;

        for (String string : value) {
            if (string.contains(placeholder)) return true;
        }

        return false;
    }

    private Settings(Properties properties)
    {
        this._clientId = properties.getProperty(CLIENT_ID_KEY);
        this._tenantId = properties.getProperty(TENANT_ID_KEY);
        this._apiScopes = properties.getProperty(SCOPES_KEY).split(",");
    }
}
