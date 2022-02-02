// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

using System.Text.Json.Serialization;

// Used to deserialize OData collections for testing
public class ODataCollection<T>
{
    [JsonPropertyName("@odata.context")]
    public string ODataContext { get; set; } = "";
    [JsonPropertyName("value")]
    public List<T> Value { get; set; } = new List<T>();
    [JsonPropertyName("@odata.nextLink")]
    public string? NextLink { get; set; }
}
