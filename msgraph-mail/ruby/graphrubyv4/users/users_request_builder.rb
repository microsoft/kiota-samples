require 'microsoft_kiota_abstractions'
require_relative '../graphrubyv4'
require_relative './users'

module Graphrubyv4
    module Users
        ## 
        # Builds and executes requests for operations under \users
        class UsersRequestBuilder < MicrosoftKiotaAbstractions::BaseRequestBuilder
            
            ## 
            ## Instantiates a new UsersRequestBuilder and sets the default values.
            ## @param path_parameters Path parameters for the request
            ## @param request_adapter The request adapter to use to execute the requests.
            ## @return a void
            ## 
            def initialize(path_parameters, request_adapter)
                super(path_parameters, request_adapter, "{+baseurl}/users")
            end
        end
    end
end
