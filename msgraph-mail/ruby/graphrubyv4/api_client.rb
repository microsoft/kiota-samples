require 'microsoft_kiota_abstractions'
require 'microsoft_kiota_serialization_json'
require_relative './graphrubyv4'
require_relative './users/item/user_item_request_builder'
require_relative './users/users_request_builder'

module Graphrubyv4
    ## 
    # The main entry point of the SDK, exposes the configuration and the fluent API.
    class ApiClient < MicrosoftKiotaAbstractions::BaseRequestBuilder
        
        ## 
        # The users property
        def users()
            return Graphrubyv4::Users::UsersRequestBuilder.new(@path_parameters, @request_adapter)
        end
        ## 
        ## Instantiates a new ApiClient and sets the default values.
        ## @param request_adapter The request adapter to use to execute the requests.
        ## @return a void
        ## 
        def initialize(request_adapter)
            super(Hash.new, request_adapter, "{+baseurl}")
            MicrosoftKiotaAbstractions::ApiClientBuilder.register_default_serializer(MicrosoftKiotaSerializationJson::JsonSerializationWriterFactory)
            MicrosoftKiotaAbstractions::ApiClientBuilder.register_default_deserializer(MicrosoftKiotaSerializationJson::JsonParseNodeFactory)
            if @request_adapter.get_base_url.nil? || @request_adapter.get_base_url.empty?
                @request_adapter.set_base_url('https://graph.microsoft.com/v1.0')
            end
            @path_parameters['baseurl'] = @request_adapter.get_base_url
        end
        ## 
        ## Gets an item from the graphrubyv4.users.item collection
        ## @param id Unique identifier of the item
        ## @return a user_item_request_builder
        ## 
        def users_by_id(id)
            raise StandardError, 'id cannot be null' if id.nil?
            url_tpl_params = @path_parameters.clone
            url_tpl_params["user%2Did"] = id
            return Graphrubyv4::Users::Item::UserItemRequestBuilder.new(url_tpl_params, @request_adapter)
        end
    end
end
