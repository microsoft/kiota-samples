# Swagger Petstore samples

The following section showcases client SDK generated using the [OpenApi 'Swagger Petstore'](https://petstore.swagger.io/) using Kiota.

NOTE: Kiota version 0.51 was used to generate these SDK's.

## Table of content

| Language | Command line |
| -------- | ------------ |
| [dotnet](./dotnet) | `.\kiota.exe --openapi https://petstore.swagger.io/v2/swagger.json --additional-data false --language csharp -o PetstoreClient -c PetstoreApiClient --log-level information --clean-output true --namespace-name KiotaSamples.PetStoreSdk` |
| [java](./java) | |
| [typescript](./typescript) | |


## Example of data retrieved

This is a dotnet example that shows how the data looks, after generation AND the target api is called:

![image](https://user-images.githubusercontent.com/899878/189605921-14f1d963-594c-4608-9d4d-eccbbfeea246.png)


## How to generate an SDK against the Swagger Petstore

These steps are all based on the official docs for [generating dotnet Kiota SDK's](https://microsoft.github.io/kiota/get-started/dotnet.html):

- Create a new project or solution.  
  e.g.
  ```
  dotnet new console -o petstore
  cd petstore
  ``` 
- Add the nuget dependencies, which we will need for the SDK generated code to work:  
  e.g.
  ```
  dotnet add package Microsoft.Kiota.Abstractions --prerelease
  dotnet add package Microsoft.Kiota.Http.HttpClientLibrary --prerelease
  dotnet add package Microsoft.Kiota.Serialization.Json --prerelease
  dotnet add package Microsoft.Kiota.Serialization.Text --prerelease
  ```
- Create the Petstore SDK using Kiota and [the custom params](https://microsoft.github.io/kiota/using.html):  
  e.g.  
  ```
  .\kiota.exe \
      --openapi https://petstore.swagger.io/v2/swagger.json \
      --additional-data false \
      --language csharp \
      --output src\PetstoreClient \ 
      --class-name PetstoreApiClient \
      --log-level information \ 
      --clean-output true \
      --namespace-name KiotaSamples.PetStoreSdk
  ```
- And that's it! you're ready to go! So, lets grab some pets:  
  ```
  var pet = await client.Pet["69"].GetAsync();
  ```

-- end of readme --