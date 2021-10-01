require './users'
require 'microsoft_kiota_abstractions'

module Graphrubyv4::Users
    ## 
    # Builds and executes requests for operations under \users
    class UsersRequestBuilder
        
        ## 
        # Current path for the request
        @current_path
        ## 
        # Whether the current path is a raw URL
        @is_raw_url
        ## 
        # Path segment to use to build the URL for the current request builder
        @path_segment
        ## 
        # The http core service to use to execute the requests.
        @request_adapter
        ## 
        ## Instantiates a new UsersRequestBuilder and sets the default values.
        ## @param currentPath Current path for the request
        ## @param isRawUrl Whether the current path is a raw URL
        ## @param requestAdapter The http core service to use to execute the requests.
        ## @return a void
        ## 
        def initialize(current_path, request_adapter, is_raw_url=true) 
            @path_segment = "/users"
            @request_adapter = request_adapter
            @current_path = current_path
            @is_raw_url = is_raw_url
        end
    end
end
