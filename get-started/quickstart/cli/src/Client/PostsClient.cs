using KiotaPostsCLI.Client.Posts;
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
namespace KiotaPostsCLI.Client {
    /// <summary>
    /// The main entry point of the SDK, exposes the configuration and the fluent API.
    /// </summary>
    public class PostsClient {
        /// <summary>Path parameters for the request</summary>
        private Dictionary<string, object> PathParameters { get; set; }
        /// <summary>Url template to use to build the URL for the current request builder</summary>
        private string UrlTemplate { get; set; }
        /// <summary>
        /// The posts property
        /// </summary>
        public Command BuildPostsCommand() {
            var command = new Command("posts");
            command.Description = "The posts property";
            var builder = new PostsRequestBuilder(PathParameters);
            command.AddCommand(builder.BuildCommand());
            command.AddCommand(builder.BuildCreateCommand());
            command.AddCommand(builder.BuildListCommand());
            return command;
        }
        /// <summary>
        /// Instantiates a new PostsClient and sets the default values.
        /// </summary>
        public Command BuildRootCommand() {
            var command = new RootCommand();
            command.Description = "Instantiates a new PostsClient and sets the default values.";
            command.AddCommand(BuildPostsCommand());
            return command;
        }
        /// <summary>
        /// Instantiates a new PostsClient and sets the default values.
        /// </summary>
        public PostsClient() {
            PathParameters = new Dictionary<string, object>();
            UrlTemplate = "{+baseurl}";
        }
    }
}
