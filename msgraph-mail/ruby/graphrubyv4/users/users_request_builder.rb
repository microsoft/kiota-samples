require 'microsoft_kiota_abstractions'
require_relative './users'

module Graphrubyv4::Users
    ## 
    # Builds and executes requests for operations under \users
    class UsersRequestBuilder
        
        ## 
        # Current path for the request
        @current_path
        ## 
        # The http core service to use to execute the requests.
        @http_core
        ## 
        # Whether the current path is a raw URL
        @is_raw_url
        ## 
        # Path segment to use to build the URL for the current request builder
        @path_segment
        ## 
        ## Instantiates a new UsersRequestBuilder and sets the default values.
        ## @param currentPath Current path for the request
        ## @param httpCore The http core service to use to execute the requests.
        ## @param isRawUrl Whether the current path is a raw URL
        ## @return a void
        ## 
        def initialize(current_path, http_core, is_raw_url=true) 
            @path_segment = "/users"
            @http_core = http_core
            @current_path = current_path
            @is_raw_url = is_raw_url
        end
    end
end
