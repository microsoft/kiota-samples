


// ** Swagger Petstore sample

// - Setup how we will render the json to the console. hint: we want it to render nicely!
// - Create an instance of our API Client, which will end up talking to the Swagger Petstore api
//   and deserialize any of the returned data into nice POCO's. 


Console.WriteLine(Environment.NewLine);
Console.WriteLine("Welcome to the Swagger Petstore example!");
Console.WriteLine(Environment.NewLine);
Console.WriteLine(Environment.NewLine);
Console.WriteLine("------------------------------------------------------");
Console.WriteLine(Environment.NewLine);


// Lets make the json look pretty, on the screen 😻
var jsonSerializerOptions = new JsonSerializerOptions
{
    WriteIndented = true
};
jsonSerializerOptions.Converters.Add(new JsonStringEnumConverter());




// Create our Api Client.

// The Swagger Petstore OpenApi endpoints don't require any authentication.
var authProvider = new Microsoft.Kiota.Abstractions.Authentication.AnonymousAuthenticationProvider();

// When need a Kiota-specific HTTP Client to do our work. As such, we'll setup our 'Adapter'
// to use the anonymous authentication provider because we don't need to authenticate when making any API requests.
var requestAdapter = new Microsoft.Kiota.Http.HttpClientLibrary.HttpClientRequestAdapter(authProvider);

// Finally -> our main api client (strongly typed, of course!)
var petstoreApiClient = new PetstoreApiClient(requestAdapter);


// Now lets start calling the API and getting some data! 🚀

Console.WriteLine("  => Calling api endpoint: pet/findByStatus");
var firstPetId = await FindPetsByStatus(petstoreApiClient, jsonSerializerOptions);

Console.WriteLine(Environment.NewLine);
Console.WriteLine(Environment.NewLine);
Console.WriteLine("  => Calling api endpoint: store/inventory");
await GetPetInfoAsync(firstPetId, petstoreApiClient, jsonSerializerOptions);

Console.WriteLine(Environment.NewLine);
Console.WriteLine(Environment.NewLine);
Console.WriteLine(" -- end of sample --");
Console.WriteLine(Environment.NewLine);
Console.WriteLine(Environment.NewLine);




/**************************************************************/

// Special notes:
// This endpoint accepts -multiple- querystring params with the same keyword: 'status'.
// e.g. status=pending and
//      status=sold
// This is translated to an array, with the generated SDK. ⚡
static async Task<long> FindPetsByStatus(PetstoreApiClient client, JsonSerializerOptions jsonSerializerOptions)
{
    var statuses = new[] { "pending", "sold" }; // Mutliple statuses we are going to query on.

    // Sample curl:
    // curl -X 'GET' 'https://petstore.swagger.io/v2/pet/findByStatus?status=available&status=pending&status=sold' -H 'accept: application/json'

    var results = (await client.Pet.FindByStatus.GetAsync(x => x.QueryParameters.Status = statuses))?.ToList();
    if (results is null)
    {
        Console.WriteLine("No results found.");
        return -1;
    }

    var numberOfItems = results.Count > 5 ? 5 : results.Count;

    Console.WriteLine($"Found {results.Count} results.");
    Console.WriteLine($"    .. First 1-{numberOfItems} results (where applicable): ");

    for (int i = 0; i < numberOfItems; i++)
    {
        var json = JsonSerializer.Serialize(results[i], jsonSerializerOptions);

        Console.Write($"    {i}) ----------------------------------------------------------------------------- ");
        Console.WriteLine(json);
    }

    var firstPetId = results.First().Id!.Value;
    return firstPetId;
}


static async Task GetPetInfoAsync(long firstPetId, PetstoreApiClient client, JsonSerializerOptions jsonSerializerOptions)
{
    // Sample curl:
    // curl -X 'GET' 'https://petstore.swagger.io/v2/store/inventory' -H 'accept: application/json'

    var result = await client.Pet[firstPetId].GetAsync();


    var json = JsonSerializer.Serialize(result, jsonSerializerOptions);
    Console.WriteLine(json);
}
