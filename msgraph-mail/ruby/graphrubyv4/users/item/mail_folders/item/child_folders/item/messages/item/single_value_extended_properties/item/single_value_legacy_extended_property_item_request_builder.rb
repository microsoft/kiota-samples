require 'microsoft_kiota_abstractions'
require_relative '../../../../../../../../../../graphrubyv4'
require_relative '../../../../../../../../../../models/single_value_legacy_extended_property'
require_relative '../../../../../../../../../users'
require_relative '../../../../../../../../item'
require_relative '../../../../../../../mail_folders'
require_relative '../../../../../../item'
require_relative '../../../../../child_folders'
require_relative '../../../../item'
require_relative '../../../messages'
require_relative '../../item'
require_relative '../single_value_extended_properties'
require_relative './item'

module Graphrubyv4
    module Users
        module Item
            module MailFolders
                module Item
                    module ChildFolders
                        module Item
                            module Messages
                                module Item
                                    module SingleValueExtendedProperties
                                        module Item
                                            ## 
                                            # Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\childFolders\{mailFolder-id1}\messages\{message-id}\singleValueExtendedProperties\{singleValueLegacyExtendedProperty-id}
                                            class SingleValueLegacyExtendedPropertyItemRequestBuilder < MicrosoftKiotaAbstractions::BaseRequestBuilder
                                                
                                                ## 
                                                ## Instantiates a new SingleValueLegacyExtendedPropertyItemRequestBuilder and sets the default values.
                                                ## @param path_parameters Path parameters for the request
                                                ## @param request_adapter The request adapter to use to execute the requests.
                                                ## @return a void
                                                ## 
                                                def initialize(path_parameters, request_adapter)
                                                    super(path_parameters, request_adapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}/messages/{message%2Did}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty%2Did}{?%24select,%24expand}")
                                                end
                                                ## 
                                                ## Delete navigation property singleValueExtendedProperties for users
                                                ## @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
                                                ## @return a Fiber of void
                                                ## 
                                                def delete(request_configuration=nil)
                                                    request_info = self.to_delete_request_information(
                                                        request_configuration
                                                    )
                                                    return @request_adapter.send_async(request_info, nil, nil)
                                                end
                                                ## 
                                                ## The collection of single-value extended properties defined for the message. Nullable.
                                                ## @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
                                                ## @return a Fiber of single_value_legacy_extended_property
                                                ## 
                                                def get(request_configuration=nil)
                                                    request_info = self.to_get_request_information(
                                                        request_configuration
                                                    )
                                                    return @request_adapter.send_async(request_info, lambda {|pn| Graphrubyv4::Models::SingleValueLegacyExtendedProperty.create_from_discriminator_value(pn) }, nil)
                                                end
                                                ## 
                                                ## Update the navigation property singleValueExtendedProperties in users
                                                ## @param body The request body
                                                ## @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
                                                ## @return a Fiber of void
                                                ## 
                                                def patch(body, request_configuration=nil)
                                                    raise StandardError, 'body cannot be null' if body.nil?
                                                    request_info = self.to_patch_request_information(
                                                        body, request_configuration
                                                    )
                                                    return @request_adapter.send_async(request_info, nil, nil)
                                                end
                                                ## 
                                                ## Delete navigation property singleValueExtendedProperties for users
                                                ## @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
                                                ## @return a request_information
                                                ## 
                                                def to_delete_request_information(request_configuration=nil)
                                                    request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
                                                    request_info.url_template = @url_template
                                                    request_info.path_parameters = @path_parameters
                                                    request_info.http_method = :DELETE
                                                    unless request_configuration.nil?
                                                        request_info.add_headers_from_raw_object(request_configuration.headers)
                                                        request_info.add_request_options(request_configuration.options)
                                                    end
                                                    return request_info
                                                end
                                                ## 
                                                ## The collection of single-value extended properties defined for the message. Nullable.
                                                ## @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
                                                ## @return a request_information
                                                ## 
                                                def to_get_request_information(request_configuration=nil)
                                                    request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
                                                    request_info.url_template = @url_template
                                                    request_info.path_parameters = @path_parameters
                                                    request_info.http_method = :GET
                                                    request_info.headers.add('Accept', 'application/json')
                                                    unless request_configuration.nil?
                                                        request_info.add_headers_from_raw_object(request_configuration.headers)
                                                        request_info.set_query_string_parameters_from_raw_object(request_configuration.query_parameters)
                                                        request_info.add_request_options(request_configuration.options)
                                                    end
                                                    return request_info
                                                end
                                                ## 
                                                ## Update the navigation property singleValueExtendedProperties in users
                                                ## @param body The request body
                                                ## @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
                                                ## @return a request_information
                                                ## 
                                                def to_patch_request_information(body, request_configuration=nil)
                                                    raise StandardError, 'body cannot be null' if body.nil?
                                                    request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
                                                    request_info.url_template = @url_template
                                                    request_info.path_parameters = @path_parameters
                                                    request_info.http_method = :PATCH
                                                    unless request_configuration.nil?
                                                        request_info.add_headers_from_raw_object(request_configuration.headers)
                                                        request_info.add_request_options(request_configuration.options)
                                                    end
                                                    request_info.set_content_from_parsable(self.request_adapter, "application/json", body)
                                                    return request_info
                                                end

                                                ## 
                                                # The collection of single-value extended properties defined for the message. Nullable.
                                                class SingleValueLegacyExtendedPropertyItemRequestBuilderGetQueryParameters
                                                    
                                                    ## 
                                                    # Expand related entities
                                                    attr_accessor :expand
                                                    ## 
                                                    # Select properties to be returned
                                                    attr_accessor :select
                                                    ## 
                                                    ## Maps the query parameters names to their encoded names for the URI template parsing.
                                                    ## @param original_name The original query parameter name in the class.
                                                    ## @return a string
                                                    ## 
                                                    def get_query_parameter(original_name)
                                                        raise StandardError, 'original_name cannot be null' if original_name.nil?
                                                        case original_name
                                                            when "expand"
                                                                return "%24expand"
                                                            when "select"
                                                                return "%24select"
                                                            else
                                                                return original_name
                                                        end
                                                    end
                                                end
                                            end
                                        end
                                    end
                                end
                            end
                        end
                    end
                end
            end
        end
    end
end
