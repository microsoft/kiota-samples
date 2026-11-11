// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package todoclient.app;

import java.io.IOException;
import java.time.OffsetDateTime;
import java.util.List;
import java.util.UUID;

import com.azure.core.credential.AccessToken;
import com.azure.core.credential.TokenRequestContext;
import com.azure.identity.DeviceCodeCredential;
import com.azure.identity.DeviceCodeCredentialBuilder;
import com.microsoft.kiota.authentication.AzureIdentityAuthenticationProvider;
import com.microsoft.kiota.http.OkHttpRequestAdapter;

import okhttp3.OkHttpClient;
import todoclient.http.Http;
import todoclient.settings.Settings;
import todoclient.utilities.ToDoApiClient;
import todoclient.utilities.models.ToDoItem;
import todoclient.utilities.models.ToDoItemUpdate;

public class App {
    private static final String[] API_HOSTS = { "localhost" };
    public static void main(String[] args) {
        Settings settings;
        try {
            settings = Settings.loadSettings();
        } catch (IOException e) {
            System.out.println("Missing or invalid configuration values in azure.properties.");
            System.out.println("See README.md for more information.");
            return;
        }

        final DeviceCodeCredential credential = new DeviceCodeCredentialBuilder()
            .clientId(settings.getClientId())
            .tenantId(settings.getTenantId())
            .challengeConsumer(challenge -> System.out.println(challenge.getMessage()))
            .build();

        final TokenRequestContext context = new TokenRequestContext().addScopes(settings.getApiScopes());
        final AccessToken token = credential.getToken(context).block();
        System.out.printf("Token: %s%n", token.getToken());

        final AzureIdentityAuthenticationProvider authProvider =
            new AzureIdentityAuthenticationProvider(credential, API_HOSTS, settings.getApiScopes());

        // This is required because the sample API project running on localhost
        // uses a self-signed certificate, and OkHttpClient by default doesn't
        // send requests to servers with self-signed certs.
        // If you're calling an API on a "real" server, this isn't necessary and you can just
        // do final OkHttpRequestAdapter adapter = new OkHttpRequestAdapter(authProvider);
        final OkHttpClient httpClient = Http.getHttpClientWithDevCert();
        final OkHttpRequestAdapter adapter = new OkHttpRequestAdapter(authProvider, null, null, httpClient);

        final ToDoApiClient client = new ToDoApiClient(adapter);

        try {
            // GET /ToDoItems
            final List<ToDoItem> items = client.toDoItems().get();
            System.out.printf("Found %d items%n", items.size());

            for (ToDoItem item: items) {
                System.out.printf("TODO: %s, Due: %s%n", item.getName(), item.getDueDate().toLocalDateTime().toString());
            }

            // Filter
            final List<ToDoItem> incompleteItems = client.toDoItems()
                .get(q -> q.queryParameters.isComplete = false);
            System.out.printf("Found %d incomplete items%n", incompleteItems.size());

            for (ToDoItem item: incompleteItems) {
                System.out.printf("%s - Priority: %d - Complete: %b%n", item.getName(), item.getPriority(), item.getIsComplete());
            }

            // Get by ID
            final UUID specificId = items.get(3).getId();
            final ToDoItem specificItem = client.toDoItems().byId(specificId).get();
            System.out.printf("Item with ID %s: %s%n", specificId, specificItem.getName());

            // Create an item
            final ToDoItem newItem = new ToDoItem();
            newItem.setName("Start using Kiota generators in my project");
            newItem.setType("work");
            newItem.setIsComplete(false);
            newItem.setDueDate(OffsetDateTime.now().plusDays(7));
            newItem.setPriority(1);

            final ToDoItem createdItem = client.toDoItems().post(newItem);
            final UUID itemId = createdItem.getId();
            System.out.printf("Item created with ID: %s%n", itemId);

            // Update an item
            final ToDoItemUpdate updateItem = new ToDoItemUpdate();
            updateItem.setIsComplete(true);

            client.toDoItems().byId(itemId).patch(updateItem);
            System.out.println("Updated item");

            // Delete an item
            client.toDoItems().byId(itemId).delete();
            System.out.println("Deleted item");

            System.out.println("All done.");
        } catch (Exception e) {
            System.out.println(e.getMessage());
        }
    }
}
