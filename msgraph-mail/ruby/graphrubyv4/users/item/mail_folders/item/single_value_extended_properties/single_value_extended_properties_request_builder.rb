require '../../../../users'
require '../../../item'
require '../../mail_folders'
require '../item'
require './single_value_extended_properties'
require 'microsoft_kiota_abstractions'
require_relative '../../../../../models/microsoft/graph/single_value_legacy_extended_property'

module Graphrubyv4::Users::Item::MailFolders::Item::SingleValueExtendedProperties
    ## 
    # Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\singleValueExtendedProperties
    class SingleValueExtendedPropertiesRequestBuilder
        
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
        ## Instantiates a new SingleValueExtendedPropertiesRequestBuilder and sets the default values.
        ## @param currentPath Current path for the request
        ## @param isRawUrl Whether the current path is a raw URL
        ## @param requestAdapter The http core service to use to execute the requests.
        ## @return a void
        ## 
        def initialize(current_path, request_adapter, is_raw_url=true) 
            @path_segment = "/singleValueExtendedProperties"
            @request_adapter = request_adapter
            @current_path = current_path
            @is_raw_url = is_raw_url
        end
        ## 
        ## The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable.
        ## @param h Request headers
        ## @param o Request options
        ## @param q Request query parameters
        ## @return a request_information
        ## 
        def create_get_request_information(q=nil, h=nil, o=nil) 
            request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
            request_info.set_uri(@current_path, @path_segment, @is_raw_url)
            request_info.http_method = :GET
            request_info.set_headers_from_raw_object(h)
            request_info.set_query_string_parameters_from_raw_object(q)
            return request_info;
        end
        ## 
        ## The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable.
        ## @param body 
        ## @param h Request headers
        ## @param o Request options
        ## @return a request_information
        ## 
        def create_post_request_information(body, h=nil, o=nil) 
            request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
            request_info.set_uri(@current_path, @path_segment, @is_raw_url)
            request_info.http_method = :POST
            request_info.set_headers_from_raw_object(h)
            request_info.set_content_from_parsable(self.request_adapter, "application/json", body)
            return request_info;
        end
        ## 
        ## The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable.
        ## @param h Request headers
        ## @param o Request options
        ## @param q Request query parameters
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of single_value_extended_properties_response
        ## 
        def get(q=nil, h=nil, o=nil, response_handler=nil) 
            request_info = self.create_get_request_information(
                q, h
            )
            return @http_core.send_async(request_info, Graphrubyv4::Users::Item::MailFolders::Item::SingleValueExtendedProperties::SingleValueExtendedPropertiesResponse, response_handler)
        end
        ## 
        ## The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable.
        ## @param body 
        ## @param h Request headers
        ## @param o Request options
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of single_value_legacy_extended_property
        ## 
        def post(body, h=nil, o=nil, response_handler=nil) 
            request_info = self.create_post_request_information(
                body, h
            )
            return @http_core.send_async(request_info, Graphrubyv4::Users::Item::MailFolders::Item::SingleValueExtendedProperties::SingleValueLegacyExtendedProperty, response_handler)
        end
    end
end
