require 'microsoft_kiota_abstractions'
require_relative './users/users_request_builder'

module Graphrubyv4::Utilities
    ## 
    # The main entry point of the SDK, exposes the configuration and the fluent API.
    class ApiClient
        
        ## 
        # The http core service to use to execute the requests.
        @http_core
        ## 
        # Path segment to use to build the URL for the current request builder
        @path_segment
        def users()
            UsersRequestBuilder.new(@path_segment , @http_core)
        end
        ## 
        ## Instantiates a new ApiClient and sets the default values.
        ## @param httpCore The http core service to use to execute the requests.
        ## @return a void
        ## 
        def initialize(http_core) 
            @path_segment = "https://graph.microsoft.com/v1.0"
            @http_core = httpCore
        end
    end
end
