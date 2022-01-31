// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

using Microsoft.AspNetCore.Mvc;
using Microsoft.OData.ModelBuilder;
using Microsoft.OpenApi;
using Microsoft.OpenApi.Extensions;
using Microsoft.OpenApi.OData;
using ToDoApi.Models;

namespace ToDoApi.Controllers
{
    [Route("[controller]")]
    public class OpenApiController : ControllerBase
    {
        [HttpGet("openapi.yaml")]
        public ActionResult GetOpenApiDoc()
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

            var yaml = openApi.SerializeAsYaml(OpenApiSpecVersion.OpenApi3_0);
            return File(System.Text.Encoding.UTF8.GetBytes(yaml), "text/yaml");
        }
    }
}
