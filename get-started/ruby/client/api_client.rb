require 'microsoft_kiota_abstractions'
require 'microsoft_kiota_serialization'
require_relative './graph'
require_relative './me/me_request_builder'

module Graph
    ## 
    # The main entry point of the SDK, exposes the configuration and the fluent API.
    class ApiClient
        
        ## 
        # The me property
        def me()
            return Graph::Me::MeRequestBuilder.new(@path_parameters, @request_adapter)
        end
        ## 
        # Path parameters for the request
        @path_parameters
        ## 
        # The request adapter to use to execute the requests.
        @request_adapter
        ## 
        # Url template to use to build the URL for the current request builder
        @url_template
        ## 
        ## Instantiates a new ApiClient and sets the default values.
        ## @param requestAdapter The request adapter to use to execute the requests.
        ## @return a void
        ## 
        def initialize(request_adapter) 
            @path_parameters = Hash.new
            @url_template = "{+baseurl}"
            @request_adapter = request_adapter
            request_adapter.set_base_url('https://graph.microsoft.com/v1.0')
        end
    end
end
