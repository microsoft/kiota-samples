require 'microsoft_kiota_abstractions'
require_relative '../graphrubyv4'
require_relative './item/user_item_request_builder'
require_relative './users'

module Graphrubyv4
    module Users
        ## 
        # Builds and executes requests for operations under \users
        class UsersRequestBuilder < MicrosoftKiotaAbstractions::BaseRequestBuilder
            
            ## 
            ## Gets an item from the graphrubyv4.users.item collection
            ## @param user_id The unique identifier of user
            ## @return a user_item_request_builder
            ## 
            def by_user_id(user_id)
                raise StandardError, 'user_id cannot be null' if user_id.nil?
                url_tpl_params = @path_parameters.clone
                url_tpl_params["user%2Did"] = user_id
                return Graphrubyv4::Users::Item::UserItemRequestBuilder.new(url_tpl_params, @request_adapter)
            end
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
