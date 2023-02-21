// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

import 'node-fetch';
import { DeviceCodeCredential } from '@azure/identity';
import { AzureIdentityAuthenticationProvider } from '@microsoft/kiota-authentication-azure';
import { FetchRequestAdapter } from '@microsoft/kiota-http-fetchlibrary';

import * as settings from './settings';
import { ToDoApiClient } from './client/toDoApiClient';
import { ToDoItemsRequestBuilderGetQueryParameters } from './client/toDoItems/toDoItemsRequestBuilderGetQueryParameters';
import { ToDoItem } from './client/models/toDoApi/models';

// By default Node won't allow connection to HTTPS server
// with a self-signed certificate. Set this to allow
// connection to test API
process.env.NODE_TLS_REJECT_UNAUTHORIZED = '0';

async function main() {
  const credential = new DeviceCodeCredential({
    clientId: settings.CLIENT_ID,
    tenantId: settings.TENANT_ID,
    userPromptCallback: (info) => {
      console.log(info.message);
    },
  });

  const token = await credential.getToken([settings.API_SCOPE]);
  console.log(`Token: ${token.token}`);

  const allowedHosts = new Set<string>(['localhost:7206']);
  const authProvider = new AzureIdentityAuthenticationProvider(
    credential,
    [settings.API_SCOPE],
    undefined,
    allowedHosts
  );

  const adapter = new FetchRequestAdapter(authProvider);
  const client = new ToDoApiClient(adapter);

  try {
    // GET /ToDoItems
    const items = await client.toDoItems.get();
    console.log(`Found ${items?.value?.length} items`);
    for (const item of items?.value || []) {
      console.log(`TODO: ${item.name}, Due: ${item.dueDate}`);
    }

    // Select
    const query = new ToDoItemsRequestBuilderGetQueryParameters();
    query.select = ['name', 'priority', 'isComplete'];
    const select = await client.toDoItems.get({
      queryParameters: query,
    });
    for (const item of select?.value || []) {
      console.log(
        `${item.name} - Priority: ${item.priority} - Complete: ${item.isComplete}`
      );
    }

    // Get by ID
    const specificId = items?.value?.at(3)?.id || '';
    const specificItem = await client.toDoItemsById(specificId).get();
    console.log(specificItem?.name);

    // Create an item
    const newItem = new ToDoItem();
    newItem.name = 'Start using Kiota generators in my project';
    newItem.type = 'work';
    newItem.isComplete = false;
    // 1 day from now
    newItem.dueDate = new Date(Date.now() + 86400000);
    newItem.priority = 1;

    const createdItem = await client.toDoItems.post(newItem);
    console.log(`Item created with ID: ${createdItem?.id}`);

    // Update an item
    const updateItem = new ToDoItem();
    updateItem.isComplete = true;

    await client.toDoItemsById(createdItem?.id || '').patch(updateItem);
    console.log('Updated item');

    // Delete an item
    await client.toDoItemsById(createdItem?.id || '').delete();
  } catch (error) {
    console.log('Caught an exception');
    console.log(JSON.stringify(error));
  }
}

main();
