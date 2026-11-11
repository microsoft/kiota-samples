// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

import { DeviceCodeCredential } from '@azure/identity';
import { AzureIdentityAuthenticationProvider } from '@microsoft/kiota-authentication-azure';
import { DefaultRequestAdapter } from '@microsoft/kiota-bundle';

import * as settings from './settings.js';
import { createToDoApiClient } from './client/toDoApiClient.js';
import type { ToDoItem, ToDoItemUpdate } from './client/models/index.js';

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
    allowedHosts,
  );

  const adapter = new DefaultRequestAdapter(authProvider);
  const client = createToDoApiClient(adapter);

  try {
    // GET /ToDoItems
    const items = await client.toDoItems.get();
    console.log(`Found ${items?.length} items`);
    for (const item of items || []) {
      console.log(`TODO: ${item.name}, Due: ${item.dueDate}`);
    }

    // Filter
    const incompleteItems = await client.toDoItems.get({
      queryParameters: { isComplete: false },
    });
    console.log(`Found ${incompleteItems?.length} incomplete items`);
    for (const item of incompleteItems || []) {
      console.log(
        `${item.name} - Priority: ${item.priority} - Complete: ${item.isComplete}`,
      );
    }

    // Get by ID
    const specificId = items?.at(3)?.id;
    if (!specificId) {
      throw new Error('Item ID is empty!');
    }
    const specificItem = await client.toDoItems.byId(specificId).get();
    console.log(specificItem?.name);

    // Create an item
    const newItem: ToDoItem = {
      name: 'Start using Kiota generators in my project',
      type: 'work',
      isComplete: false,
      // 1 day from now
      dueDate: new Date(Date.now() + 86400000),
      priority: 1,
    };

    const createdItem = await client.toDoItems.post(newItem);
    console.log(`Item created with ID: ${createdItem?.id}`);

    // Update an item
    const createdItemId = createdItem?.id;
    if (!createdItemId) {
      throw new Error('Item ID is empty!');
    }
    const updateItem: ToDoItemUpdate = {
      isComplete: true,
    };

    await client.toDoItems.byId(createdItemId).patch(updateItem);
    console.log('Updated item');

    // Delete an item
    await client.toDoItems.byId(createdItemId).delete();
    console.log('Deleted item');
  } catch (error) {
    console.log('Caught an exception');
    console.log(JSON.stringify(error));
  }
}

main();
