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

-- end of readme --