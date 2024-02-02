// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// <ProgramSnippet>
import { DeviceCodeCredential } from '@azure/identity';
import { AzureIdentityAuthenticationProvider } from '@microsoft/kiota-authentication-azure';
import { FetchRequestAdapter } from '@microsoft/kiota-http-fetchlibrary';
import { createGetUserApiClient } from './client/getUserApiClient';

const clientId = 'YOUR_CLIENT_ID';

// The auth provider will only authorize requests to
// the allowed hosts, in this case Microsoft Graph
const allowedHosts = new Set<string>([ 'graph.microsoft.com' ]);
const graphScopes = [ 'User.Read' ];

const credential = new DeviceCodeCredential({
  clientId: clientId,
  userPromptCallback: (deviceCodeInfo) => {
    console.log(deviceCodeInfo.message);
  }
});

const authProvider =
  new AzureIdentityAuthenticationProvider(credential, graphScopes, undefined, allowedHosts);
const adapter = new FetchRequestAdapter(authProvider);

const client = createGetUserApiClient(adapter);

async function GetUser(): Promise<void> {
  try {
    const me = await client.me.get();
    console.log(`Hello ${me?.displayName}, your ID is ${me?.id}`);
  } catch (err) {
    console.log(err);
  }
}

GetUser();
// </ProgramSnippet>
