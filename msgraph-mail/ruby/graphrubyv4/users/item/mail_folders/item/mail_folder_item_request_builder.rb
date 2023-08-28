require 'microsoft_kiota_abstractions'
require_relative '../../../../graphrubyv4'
require_relative '../../../../models/mail_folder'
require_relative '../../../users'
require_relative '../../item'
require_relative '../mail_folders'
require_relative './child_folders/child_folders_request_builder'
require_relative './item'
require_relative './message_rules/message_rules_request_builder'
require_relative './messages/messages_request_builder'

module Graphrubyv4
    module Users
        module Item
            module MailFolders
                module Item
                    ## 
                    # Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}
                    class MailFolderItemRequestBuilder < MicrosoftKiotaAbstractions::BaseRequestBuilder
                        
                        ## 
                        # The childFolders property
                        def child_folders()
                            return Graphrubyv4::Users::Item::MailFolders::Item::ChildFolders::ChildFoldersRequestBuilder.new(@path_parameters, @request_adapter)
                        end
                        ## 
                        # The messageRules property
                        def message_rules()
                            return Graphrubyv4::Users::Item::MailFolders::Item::MessageRules::MessageRulesRequestBuilder.new(@path_parameters, @request_adapter)
                        end
                        ## 
                        # The messages property
                        def messages()
                            return Graphrubyv4::Users::Item::MailFolders::Item::Messages::MessagesRequestBuilder.new(@path_parameters, @request_adapter)
                        end
                        ## 
                        ## Instantiates a new MailFolderItemRequestBuilder and sets the default values.
                        ## @param path_parameters Path parameters for the request
                        ## @param request_adapter The request adapter to use to execute the requests.
                        ## @return a void
                        ## 
                        def initialize(path_parameters, request_adapter)
                            super(path_parameters, request_adapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}{?includeHiddenFolders,%24select,%24expand}")
                        end
                        ## 
                        ## Delete the specified mailFolder. The folder can be a mailSearchFolder. You can specify a mail folder by its folder ID, or by its well-known folder name, if one exists.
                        ## @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
                        ## @return a Fiber of binary
                        ## 
                        def delete(request_configuration=nil)
                            request_info = self.to_delete_request_information(
                                request_configuration
                            )
                            return @request_adapter.send_async(request_info, Binary, nil)
                        end
                        ## 
                        ## The user's mail folders. Read-only. Nullable.
                        ## @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
                        ## @return a Fiber of mail_folder
                        ## 
                        def get(request_configuration=nil)
                            request_info = self.to_get_request_information(
                                request_configuration
                            )
                            return @request_adapter.send_async(request_info, lambda {|pn| Graphrubyv4::Models::MailFolder.create_from_discriminator_value(pn) }, nil)
                        end
                        ## 
                        ## Update the properties of mailfolder object.
                        ## @param body The request body
                        ## @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
                        ## @return a Fiber of mail_folder
                        ## 
                        def patch(body, request_configuration=nil)
                            raise StandardError, 'body cannot be null' if body.nil?
                            request_info = self.to_patch_request_information(
                                body, request_configuration
                            )
                            return @request_adapter.send_async(request_info, lambda {|pn| Graphrubyv4::Models::MailFolder.create_from_discriminator_value(pn) }, nil)
                        end
                        ## 
                        ## Delete the specified mailFolder. The folder can be a mailSearchFolder. You can specify a mail folder by its folder ID, or by its well-known folder name, if one exists.
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
                        ## The user's mail folders. Read-only. Nullable.
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
                        ## Update the properties of mailfolder object.
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
                            request_info.headers.add('Accept', 'application/json')
                            unless request_configuration.nil?
                                request_info.add_headers_from_raw_object(request_configuration.headers)
                                request_info.add_request_options(request_configuration.options)
                            end
                            request_info.set_content_from_parsable(@request_adapter, "application/json", body)
                            return request_info
                        end
                        ## 
                        ## Returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
                        ## @param raw_url The raw URL to use for the request builder.
                        ## @return a mail_folder_item_request_builder
                        ## 
                        def with_url(raw_url)
                            raise StandardError, 'raw_url cannot be null' if raw_url.nil?
                            return MailFolderItemRequestBuilder.new(raw_url, @request_adapter)
                        end

                        ## 
                        # The user's mail folders. Read-only. Nullable.
                        class MailFolderItemRequestBuilderGetQueryParameters
                            
                            ## 
                            # Expand related entities
                            attr_accessor :expand
                            ## 
                            # Include Hidden Folders
                            attr_accessor :include_hidden_folders
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
                                    when "include_hidden_folders"
                                        return "includeHiddenFolders"
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
