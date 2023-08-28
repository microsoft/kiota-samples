require 'microsoft_kiota_abstractions'
require 'microsoft_kiota_serialization_json'
require_relative './graphrubyv4'
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
    end
end
