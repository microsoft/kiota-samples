require 'microsoft_kiota_abstractions'
require_relative '../../../users'
require_relative '../../mail_folders'
require_relative '../messages'
require_relative './item/single_value_legacy_extended_property_request_builder'
require_relative './single_value_extended_properties'

module Graphrubyv4::Users::MailFolders::Messages::SingleValueExtendedProperties
    ## 
    # Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\messages\{message-id}\singleValueExtendedProperties
    class SingleValueExtendedPropertiesRequestBuilder
        
        ## 
        # Current path for the request
        @current_path
        ## 
        # The http core service to use to execute the requests.
        @http_core
        ## 
        # Path segment to use to build the URL for the current request builder
        @path_segment
        ## 
        # Gets an item from the graphrubyv4.users.mailFolders.messages.singleValueExtendedProperties collection
        def [](position)
            return SingleValueLegacyExtendedPropertyRequestBuilder.new(@current_path + @path_segment  + "/" + position, @http_core)
        end
        ## 
        ## Instantiates a new SingleValueExtendedPropertiesRequestBuilder and sets the default values.
        ## @param currentPath Current path for the request
        ## @param httpCore The http core service to use to execute the requests.
        ## @return a void
        ## 
        def initialize(current_path, http_core) 
            @path_segment = "/singleValueExtendedProperties"
            @http_core = http_core
            @current_path = current_path
        end
        ## 
        ## The collection of single-value extended properties defined for the message. Nullable.
        ## @param h Request headers
        ## @param o Request options for HTTP middlewares
        ## @param q Request query parameters
        ## @return a request_info
        ## 
        def create_get_request_info(q, h, o) 
            request_info = RequestInfo.new()
            request_info.URI = current_path + path_segment
            request_info.http_method = :GET
            request_info.set_headers_from_raw_object(h)
            request_info.set_query_string_parameters_from_raw_object(q)
            return request_info;
        end
        ## 
        ## The collection of single-value extended properties defined for the message. Nullable.
        ## @param body 
        ## @param h Request headers
        ## @param o Request options for HTTP middlewares
        ## @return a request_info
        ## 
        def create_post_request_info(body, h, o) 
            request_info = RequestInfo.new()
            request_info.URI = current_path + path_segment
            request_info.http_method = :POST
            request_info.set_headers_from_raw_object(h)
            request_info.set_content_from_parsable(self.serializer_factory, "application/json", body)
            return request_info;
        end
        ## 
        ## The collection of single-value extended properties defined for the message. Nullable.
        ## @param h Request headers
        ## @param o Request options for HTTP middlewares
        ## @param q Request query parameters
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of single_value_extended_properties_response
        ## 
        def get(q, h, o, response_handler) 
            request_info = self.create_get_request_info(
                q, h
            )
            return self.http_core.send_async(request_info, response_handler)
        end
        ## 
        ## The collection of single-value extended properties defined for the message. Nullable.
        ## @param body 
        ## @param h Request headers
        ## @param o Request options for HTTP middlewares
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of single_value_legacy_extended_property
        ## 
        def post(body, h, o, response_handler) 
            request_info = self.create_post_request_info(
                body, h
            )
            return self.http_core.send_async(request_info, response_handler)
        end
    end
end
