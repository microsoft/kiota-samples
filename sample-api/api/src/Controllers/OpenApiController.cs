// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

using Microsoft.AspNetCore.Mvc;
using Microsoft.OData.ModelBuilder;
using Microsoft.OpenApi;
using Microsoft.OpenApi.OData;
using ToDoApi.Models;
using System.IO;

namespace ToDoApi.Controllers
{
    [Route("[controller]")]
    public class OpenApiController : ControllerBase
    {
        [HttpGet("openapi.yaml")]
        public async Task<ActionResult> GetOpenApiDoc()
        {
            var odataBuilder = new ODataConventionModelBuilder();
            odataBuilder.EnableLowerCamelCase();
            odataBuilder.EntitySet<ToDoItem>("ToDoItems");
            var model = odataBuilder.GetEdmModel();

            var openApi = model.ConvertToOpenApi(new OpenApiConvertSettings
            {
                ServiceRoot = new Uri($"{Request.Scheme}://{Request.Host}/"),
                EnableKeyAsSegment = true,
                PrefixEntityTypeNameBeforeKey =  true
            });


            // The asp.net infrastructure will dispose the stream, so we need to leave it open.
            var ms = new MemoryStream();
            await using var streamWriter = new StreamWriter(ms, leaveOpen: true);
            var yamlWriter = new OpenApiYamlWriter(streamWriter);
            openApi.SerializeAsV31(yamlWriter);
            await streamWriter.FlushAsync();
            ms.Position = 0;

            return File(ms, "text/yaml");
        }
    }
}
