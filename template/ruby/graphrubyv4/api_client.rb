require 'microsoft_kiota_abstractions'
require_relative './graphrubyv4'
require_relative './users/item/user_request_builder'
require_relative './users/users_request_builder'

module Graphrubyv4
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
            return Graphrubyv4::Users::UsersRequestBuilder.new(@path_segment , @http_core)
        end
        ## 
        ## Instantiates a new ApiClient and sets the default values.
        ## @param httpCore The http core service to use to execute the requests.
        ## @return a void
        ## 
        def initialize(http_core) 
            @path_segment = "https://graph.microsoft.com/v1.0"
            @http_core = http_core
        end
        ## 
        ## Gets an item from the graphrubyv4.users collection
        ## @param id Unique identifier of the item
        ## @return a user_request_builder
        ## 
        def users_by_id(id) 
            return Graphrubyv4::Users::Item::UserRequestBuilder.new(@path_segment  + "/users/" + id, @http_core)
        end
    end
end
