require 'microsoft_kiota_abstractions'
require_relative '../../../../../../models/multi_value_legacy_extended_property'
require_relative '../../../../../users'
require_relative '../../../../item'
require_relative '../../../mail_folders'
require_relative '../../item'
require_relative '../multi_value_extended_properties'
require_relative './item'

module Graphrubyv4::Users::Item::MailFolders::Item::MultiValueExtendedProperties::Item
    ## 
    # Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\multiValueExtendedProperties\{multiValueLegacyExtendedProperty-id}
    class MultiValueLegacyExtendedPropertyItemRequestBuilder
        
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
        ## Instantiates a new MultiValueLegacyExtendedPropertyItemRequestBuilder and sets the default values.
        ## @param pathParameters Path parameters for the request
        ## @param requestAdapter The request adapter to use to execute the requests.
        ## @return a void
        ## 
        def initialize(path_parameters, request_adapter) 
            @url_template = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty%2Did}{?%24select,%24expand}"
            @request_adapter = request_adapter
            if path_parameters.is_a? String
                path_parameters = { "request-raw-url" => path_parameters }
            end
            @path_parameters = path_parameters
        end
        ## 
        ## Delete navigation property multiValueExtendedProperties for users
        ## @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
        ## @return a request_information
        ## 
        def create_delete_request_information(request_configuration=nil) 
            request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
            request_info.url_template = @url_template
            request_info.path_parameters = @path_parameters
            request_info.http_method = :DELETE
            unless request_configuration.nil?
                request_info.set_headers_from_raw_object(request_configuration.headers)
            end
            return request_info
        end
        ## 
        ## The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable.
        ## @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
        ## @return a request_information
        ## 
        def create_get_request_information(request_configuration=nil) 
            request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
            request_info.url_template = @url_template
            request_info.path_parameters = @path_parameters
            request_info.http_method = :GET
            request_info.headers['Accept'] = 'application/json'
            unless request_configuration.nil?
                request_info.set_headers_from_raw_object(request_configuration.headers)
                request_info.set_query_string_parameters_from_raw_object(request_configuration.query_parameters)
            end
            return request_info
        end
        ## 
        ## Update the navigation property multiValueExtendedProperties in users
        ## @param body 
        ## @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
        ## @return a request_information
        ## 
        def create_patch_request_information(body, request_configuration=nil) 
            request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
            request_info.url_template = @url_template
            request_info.path_parameters = @path_parameters
            request_info.http_method = :PATCH
            unless request_configuration.nil?
                request_info.set_headers_from_raw_object(request_configuration.headers)
            end
            request_info.set_content_from_parsable(self.request_adapter, "application/json", body)
            return request_info
        end
        ## 
        ## Delete navigation property multiValueExtendedProperties for users
        ## @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of void
        ## 
        def delete(request_configuration=nil, response_handler=nil) 
            request_info = self.create_delete_request_information(
                request_configuration
            )
            return @request_adapter.send_async(request_info, nil, response_handler)
        end
        ## 
        ## The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable.
        ## @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of multi_value_legacy_extended_property
        ## 
        def get(request_configuration=nil, response_handler=nil) 
            request_info = self.create_get_request_information(
                request_configuration
            )
            return @request_adapter.send_async(request_info, Graphrubyv4::Models::MultiValueLegacyExtendedProperty, response_handler)
        end
        ## 
        ## Update the navigation property multiValueExtendedProperties in users
        ## @param body 
        ## @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of void
        ## 
        def patch(body, request_configuration=nil, response_handler=nil) 
            request_info = self.create_patch_request_information(
                body, request_configuration
            )
            return @request_adapter.send_async(request_info, nil, response_handler)
        end

        ## 
        # Configuration for the request such as headers, query parameters, and middleware options.
        class MultiValueLegacyExtendedPropertyItemRequestBuilderDeleteRequestConfiguration
            
            ## 
            # Request headers
            @headers
            ## 
            # Request options
            @options
        end

        ## 
        # The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable.
        class MultiValueLegacyExtendedPropertyItemRequestBuilderGetQueryParameters
            
            ## 
            # Expand related entities
            @expand
            ## 
            # Select properties to be returned
            @select
        end

        ## 
        # Configuration for the request such as headers, query parameters, and middleware options.
        class MultiValueLegacyExtendedPropertyItemRequestBuilderGetRequestConfiguration
            
            ## 
            # Request headers
            @headers
            ## 
            # Request options
            @options
            ## 
            # Request query parameters
            @query_parameters
        end

        ## 
        # Configuration for the request such as headers, query parameters, and middleware options.
        class MultiValueLegacyExtendedPropertyItemRequestBuilderPatchRequestConfiguration
            
            ## 
            # Request headers
            @headers
            ## 
            # Request options
            @options
        end
    end
end
