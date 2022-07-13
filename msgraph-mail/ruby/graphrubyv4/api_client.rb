require './graphrubyv4'
require 'microsoft_kiota_abstractions'
require 'microsoft_kiota_serialization'
require_relative './users/item/user_item_request_builder'
require_relative './users/users_request_builder'

module Graphrubyv4
    ## 
    # The main entry point of the SDK, exposes the configuration and the fluent API.
    class ApiClient
        
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
        # The users property
        def users()
            return Graphrubyv4::Users::UsersRequestBuilder.new(@path_parameters, @request_adapter)
        end
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
        ## 
        ## Gets an item from the graphrubyv4.users.item collection
        ## @param id Unique identifier of the item
        ## @return a user_item_request_builder
        ## 
        def users_by_id(id) 
            url_tpl_params = @path_parameters.clone
            url_tpl_params["user%2Did"] = id
            return Graphrubyv4::Users::Item::UserItemRequestBuilder.new(url_tpl_params, @request_adapter)
        end
    end
end
