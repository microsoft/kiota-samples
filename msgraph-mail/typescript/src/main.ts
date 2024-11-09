import { Logger } from "./app/common/logger.js";
import { DeviceCodeCredential } from "@azure/identity";
import { AzureIdentityAuthenticationProvider } from "@microsoft/kiota-authentication-azure";
import {
  FetchRequestAdapter,
  HttpClient,
} from "@microsoft/kiota-http-fetchlibrary";
import { createApiClient } from "./apiClient.js";
import {
  JsonParseNodeFactory,
  JsonSerializationWriterFactory,
} from "@microsoft/kiota-serialization-json";

async function main(): Promise<void> {
  const cred = new DeviceCodeCredential({
    tenantId: "YOUR TENANT ID",
    clientId: "YOUR CLIENT ID",
    userPromptCallback: (deviceCodeInfo) => {
      // eslint-disable-next-line no-console
      console.log(deviceCodeInfo.message);
    },
  });
  const authProvider = new AzureIdentityAuthenticationProvider(cred, [
    "Mail.Read",
  ]);
  const requestAdapter = new FetchRequestAdapter(
    authProvider,
    new JsonParseNodeFactory(),
    new JsonSerializationWriterFactory(),
    new HttpClient(),
    {
      includeEUIIAttributes: true,
    }
  );
  const client = createApiClient(requestAdapter);
  const result = await client.users
    .byUserId("TARGET USER EMAIL")
    .messages.get();
  for (const message of result?.value!) {
    Logger.logTask(`${message.subject}`);
  }
  Logger.logTask("hello world");
}
main();
