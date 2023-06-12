# Copyright (c) Microsoft Corporation.
# Licensed under the MIT License.

# <ProgramSnippet>
import asyncio
from azure.identity import DeviceCodeCredential
from kiota_authentication_azure.azure_identity_authentication_provider import (
    AzureIdentityAuthenticationProvider)
from kiota_http.httpx_request_adapter import HttpxRequestAdapter
from client.get_user_api_client import GetUserApiClient

async def main():
    # You may need this if your're using asyncio on windows
    # See: https://stackoverflow.com/questions/63860576
    # asyncio.set_event_loop_policy(asyncio.WindowsSelectorEventLoopPolicy())

    client_id = 'YOUR_CLIENT_ID'
    graph_scopes = ['User.Read']

    credential = DeviceCodeCredential(client_id)
    auth_provider = AzureIdentityAuthenticationProvider(credential, scopes=graph_scopes)

    request_adapter = HttpxRequestAdapter(auth_provider)
    client = GetUserApiClient(request_adapter)

    user_me = await client.me.get()
    print(f"Hello {user_me.display_name}, your ID is {user_me.id}")

# Run main
asyncio.run(main())
# </ProgramSnippet>
