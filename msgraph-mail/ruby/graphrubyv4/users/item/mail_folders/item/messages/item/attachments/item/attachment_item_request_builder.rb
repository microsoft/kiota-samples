require 'microsoft_kiota_abstractions'
require_relative '../../../../../../../../graphrubyv4'
require_relative '../../../../../../../../models/attachment'
require_relative '../../../../../../../../models/o_data_errors_o_data_error'
require_relative '../../../../../../../users'
require_relative '../../../../../../item'
require_relative '../../../../../mail_folders'
require_relative '../../../../item'
require_relative '../../../messages'
require_relative '../../item'
require_relative '../attachments'
require_relative './item'

module Graphrubyv4
    module Users
        module Item
            module MailFolders
                module Item
                    module Messages
                        module Item
                            module Attachments
                                module Item
                                    ## 
                                    # Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\messages\{message-id}\attachments\{attachment-id}
                                    class AttachmentItemRequestBuilder < MicrosoftKiotaAbstractions::BaseRequestBuilder
                                        
                                        ## 
                                        ## Instantiates a new AttachmentItemRequestBuilder and sets the default values.
                                        ## @param path_parameters Path parameters for the request
                                        ## @param request_adapter The request adapter to use to execute the requests.
                                        ## @return a void
                                        ## 
                                        def initialize(path_parameters, request_adapter)
                                            super(path_parameters, request_adapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/messages/{message%2Did}/attachments/{attachment%2Did}{?%24expand,%24select}")
                                        end
                                        ## 
                                        ## Delete navigation property attachments for users
                                        ## @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
                                        ## @return a Fiber of binary
                                        ## 
                                        def delete(request_configuration=nil)
                                            request_info = self.to_delete_request_information(
                                                request_configuration
                                            )
                                            error_mapping = Hash.new
                                            error_mapping["4XX"] = lambda {|pn| Graphrubyv4::Models::ODataErrorsODataError.create_from_discriminator_value(pn) }
                                            error_mapping["5XX"] = lambda {|pn| Graphrubyv4::Models::ODataErrorsODataError.create_from_discriminator_value(pn) }
                                            return @request_adapter.send_async(request_info, Binary, error_mapping)
                                        end
                                        ## 
                                        ## Read the properties, relationships, or raw contents of an attachment that is attached to a user event, message, or group post. An attachment can be one of the following types: All these types of attachments are derived from the attachment resource.
                                        ## @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
                                        ## @return a Fiber of attachment
                                        ## 
                                        def get(request_configuration=nil)
                                            request_info = self.to_get_request_information(
                                                request_configuration
                                            )
                                            error_mapping = Hash.new
                                            error_mapping["4XX"] = lambda {|pn| Graphrubyv4::Models::ODataErrorsODataError.create_from_discriminator_value(pn) }
                                            error_mapping["5XX"] = lambda {|pn| Graphrubyv4::Models::ODataErrorsODataError.create_from_discriminator_value(pn) }
                                            return @request_adapter.send_async(request_info, lambda {|pn| Graphrubyv4::Models::Attachment.create_from_discriminator_value(pn) }, error_mapping)
                                        end
                                        ## 
                                        ## Delete navigation property attachments for users
                                        ## @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
                                        ## @return a request_information
                                        ## 
                                        def to_delete_request_information(request_configuration=nil)
                                            request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
                                            unless request_configuration.nil?
                                                request_info.add_headers_from_raw_object(request_configuration.headers)
                                                request_info.add_request_options(request_configuration.options)
                                            end
                                            request_info.url_template = '{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/messages/{message%2Did}/attachments/{attachment%2Did}'
                                            request_info.path_parameters = @path_parameters
                                            request_info.http_method = :DELETE
                                            request_info.headers.try_add('Accept', 'application/json')
                                            return request_info
                                        end
                                        ## 
                                        ## Read the properties, relationships, or raw contents of an attachment that is attached to a user event, message, or group post. An attachment can be one of the following types: All these types of attachments are derived from the attachment resource.
                                        ## @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
                                        ## @return a request_information
                                        ## 
                                        def to_get_request_information(request_configuration=nil)
                                            request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
                                            unless request_configuration.nil?
                                                request_info.add_headers_from_raw_object(request_configuration.headers)
                                                request_info.set_query_string_parameters_from_raw_object(request_configuration.query_parameters)
                                                request_info.add_request_options(request_configuration.options)
                                            end
                                            request_info.url_template = @url_template
                                            request_info.path_parameters = @path_parameters
                                            request_info.http_method = :GET
                                            request_info.headers.try_add('Accept', 'application/json')
                                            return request_info
                                        end
                                        ## 
                                        ## Returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
                                        ## @param raw_url The raw URL to use for the request builder.
                                        ## @return a attachment_item_request_builder
                                        ## 
                                        def with_url(raw_url)
                                            raise StandardError, 'raw_url cannot be null' if raw_url.nil?
                                            return AttachmentItemRequestBuilder.new(raw_url, @request_adapter)
                                        end

                                        ## 
                                        # Read the properties, relationships, or raw contents of an attachment that is attached to a user event, message, or group post. An attachment can be one of the following types: All these types of attachments are derived from the attachment resource.
                                        class AttachmentItemRequestBuilderGetQueryParameters
                                            
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
