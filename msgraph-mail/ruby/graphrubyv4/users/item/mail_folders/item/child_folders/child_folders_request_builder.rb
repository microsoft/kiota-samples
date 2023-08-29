require 'microsoft_kiota_abstractions'
require_relative '../../../../../graphrubyv4'
require_relative '../../../../../models/mail_folder'
require_relative '../../../../../models/mail_folder_collection_response'
require_relative '../../../../users'
require_relative '../../../item'
require_relative '../../mail_folders'
require_relative '../item'
require_relative './child_folders'
require_relative './count/count_request_builder'
require_relative './item/mail_folder_item_request_builder'

module Graphrubyv4
    module Users
        module Item
            module MailFolders
                module Item
                    module ChildFolders
                        ## 
                        # Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\childFolders
                        class ChildFoldersRequestBuilder < MicrosoftKiotaAbstractions::BaseRequestBuilder
                            
                            ## 
                            # The Count property
                            def count()
                                return Graphrubyv4::Users::Item::MailFolders::Item::ChildFolders::Count::CountRequestBuilder.new(@path_parameters, @request_adapter)
                            end
                            ## 
                            ## Gets an item from the graphrubyv4.users.item.mailFolders.item.childFolders.item collection
                            ## @param mail_folder_id1 The unique identifier of mailFolder
                            ## @return a mail_folder_item_request_builder
                            ## 
                            def by_mail_folder_id1(mail_folder_id1)
                                raise StandardError, 'mail_folder_id1 cannot be null' if mail_folder_id1.nil?
                                url_tpl_params = @path_parameters.clone
                                url_tpl_params["mailFolder%2Did1"] = mail_folder_id1
                                return Graphrubyv4::Users::Item::MailFolders::Item::ChildFolders::Item::MailFolderItemRequestBuilder.new(url_tpl_params, @request_adapter)
                            end
                            ## 
                            ## Instantiates a new ChildFoldersRequestBuilder and sets the default values.
                            ## @param path_parameters Path parameters for the request
                            ## @param request_adapter The request adapter to use to execute the requests.
                            ## @return a void
                            ## 
                            def initialize(path_parameters, request_adapter)
                                super(path_parameters, request_adapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders{?includeHiddenFolders,%24top,%24skip,%24filter,%24count,%24orderby,%24select,%24expand}")
                            end
                            ## 
                            ## The collection of child folders in the mailFolder.
                            ## @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
                            ## @return a Fiber of mail_folder_collection_response
                            ## 
                            def get(request_configuration=nil)
                                request_info = self.to_get_request_information(
                                    request_configuration
                                )
                                return @request_adapter.send_async(request_info, lambda {|pn| Graphrubyv4::Models::MailFolderCollectionResponse.create_from_discriminator_value(pn) }, nil)
                            end
                            ## 
                            ## Use this API to create a new child mailFolder. If you intend a new folder to be hidden, you must set the isHidden property to true on creation.
                            ## @param body The request body
                            ## @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
                            ## @return a Fiber of mail_folder
                            ## 
                            def post(body, request_configuration=nil)
                                raise StandardError, 'body cannot be null' if body.nil?
                                request_info = self.to_post_request_information(
                                    body, request_configuration
                                )
                                return @request_adapter.send_async(request_info, lambda {|pn| Graphrubyv4::Models::MailFolder.create_from_discriminator_value(pn) }, nil)
                            end
                            ## 
                            ## The collection of child folders in the mailFolder.
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
                            ## Use this API to create a new child mailFolder. If you intend a new folder to be hidden, you must set the isHidden property to true on creation.
                            ## @param body The request body
                            ## @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
                            ## @return a request_information
                            ## 
                            def to_post_request_information(body, request_configuration=nil)
                                raise StandardError, 'body cannot be null' if body.nil?
                                request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
                                request_info.url_template = @url_template
                                request_info.path_parameters = @path_parameters
                                request_info.http_method = :POST
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
                            ## @return a child_folders_request_builder
                            ## 
                            def with_url(raw_url)
                                raise StandardError, 'raw_url cannot be null' if raw_url.nil?
                                return ChildFoldersRequestBuilder.new(raw_url, @request_adapter)
                            end

                            ## 
                            # The collection of child folders in the mailFolder.
                            class ChildFoldersRequestBuilderGetQueryParameters
                                
                                ## 
                                # Include count of items
                                attr_accessor :count
                                ## 
                                # Expand related entities
                                attr_accessor :expand
                                ## 
                                # Filter items by property values
                                attr_accessor :filter
                                ## 
                                # Include Hidden Folders
                                attr_accessor :include_hidden_folders
                                ## 
                                # Order items by property values
                                attr_accessor :orderby
                                ## 
                                # Select properties to be returned
                                attr_accessor :select
                                ## 
                                # Skip the first n items
                                attr_accessor :skip
                                ## 
                                # Show only the first n items
                                attr_accessor :top
                                ## 
                                ## Maps the query parameters names to their encoded names for the URI template parsing.
                                ## @param original_name The original query parameter name in the class.
                                ## @return a string
                                ## 
                                def get_query_parameter(original_name)
                                    raise StandardError, 'original_name cannot be null' if original_name.nil?
                                    case original_name
                                        when "count"
                                            return "%24count"
                                        when "expand"
                                            return "%24expand"
                                        when "filter"
                                            return "%24filter"
                                        when "include_hidden_folders"
                                            return "includeHiddenFolders"
                                        when "orderby"
                                            return "%24orderby"
                                        when "select"
                                            return "%24select"
                                        when "skip"
                                            return "%24skip"
                                        when "top"
                                            return "%24top"
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
