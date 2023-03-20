// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package getuserclient;

import com.azure.identity.DeviceCodeCredential;
import com.azure.identity.DeviceCodeCredentialBuilder;
import com.microsoft.kiota.authentication.AzureIdentityAuthenticationProvider;
import com.microsoft.kiota.http.OkHttpRequestAdapter;

import getuserclient.apiclient.GetUserApiClient;

public class App {

    public static void main(String[] args) {
        final String clientId = "YOUR_CLIENT_ID";

        // The auth provider will only authorize requests to
        // the allowed hosts, in this case Microsoft Graph
        final String[] allowedHosts = new String[] { "graph.microsoft.com" };
        final String[] graphScopes = new String[] { "User.Read" };

        final DeviceCodeCredential credential = new DeviceCodeCredentialBuilder()
            .clientId(clientId)
            .challengeConsumer(challenge -> System.out.println(challenge.getMessage()))
            .build();


        final AzureIdentityAuthenticationProvider authProvider =
            new AzureIdentityAuthenticationProvider(credential, allowedHosts, graphScopes);
        final OkHttpRequestAdapter adapter = new OkHttpRequestAdapter(authProvider);

        final GetUserApiClient client = new GetUserApiClient(adapter);

        client.me()
            .get()
            .thenAccept(me -> {
                System.out.printf("Hello %s, your ID is %s%n",
                    me.getDisplayName(), me.getId());
            })
            .exceptionally(err -> {
                System.out.printf("Error: %s%n", err.getMessage());
                return null;
            })
            .join();
    }
}
