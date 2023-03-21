using GetUserClient.ApiClient.Me;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using Microsoft.Kiota.Abstractions;
using Microsoft.Kiota.Abstractions.Extensions;
using Microsoft.Kiota.Cli.Commons.IO;
using Microsoft.Kiota.Serialization.Form;
using Microsoft.Kiota.Serialization.Json;
using Microsoft.Kiota.Serialization.Text;
using System;
using System.Collections.Generic;
using System.CommandLine;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
namespace GetUserClient.ApiClient {
    /// <summary>
    /// The main entry point of the SDK, exposes the configuration and the fluent API.
    /// </summary>
    public class GetUserApiClient {
        /// <summary>Path parameters for the request</summary>
        private Dictionary<string, object> PathParameters { get; set; }
        /// <summary>Url template to use to build the URL for the current request builder</summary>
        private string UrlTemplate { get; set; }
        /// <summary>
        /// The me property
        /// </summary>
        public Command BuildMeCommand() {
            var command = new Command("me");
            command.Description = "The me property";
            var builder = new MeRequestBuilder(PathParameters);
            command.AddCommand(builder.BuildGetCommand());
            return command;
        }
        /// <summary>
        /// Instantiates a new GetUserApiClient and sets the default values.
        /// </summary>
        public Command BuildRootCommand() {
            var command = new RootCommand();
            command.Description = "Instantiates a new GetUserApiClient and sets the default values.";
            command.AddCommand(BuildMeCommand());
            return command;
        }
        /// <summary>
        /// Instantiates a new GetUserApiClient and sets the default values.
        /// </summary>
        public GetUserApiClient() {
            PathParameters = new Dictionary<string, object>();
            UrlTemplate = "{+baseurl}";
        }
    }
}
