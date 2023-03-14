require 'microsoft_kiota_abstractions'
require_relative '../../../../../../../../../graphrubyv4'
require_relative '../../../../../../../../users'
require_relative '../../../../../../../item'
require_relative '../../../../../../mail_folders'
require_relative '../../../../../item'
require_relative '../../../../child_folders'
require_relative '../../../item'
require_relative '../../messages'
require_relative '../item'
require_relative './value'

module Graphrubyv4
    module Users
        module Item
            module MailFolders
                module Item
                    module ChildFolders
                        module Item
                            module Messages
                                module Item
                                    module Value
                                        ## 
                                        # Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\childFolders\{mailFolder-id1}\messages\{message-id}\$value
                                        class ContentRequestBuilder < MicrosoftKiotaAbstractions::BaseRequestBuilder
                                            
                                            ## 
                                            ## Instantiates a new ContentRequestBuilder and sets the default values.
                                            ## @param path_parameters Path parameters for the request
                                            ## @param request_adapter The request adapter to use to execute the requests.
                                            ## @return a void
                                            ## 
                                            def initialize(path_parameters, request_adapter)
                                                super(path_parameters, request_adapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}/messages/{message%2Did}/$value")
                                            end
                                            ## 
                                            ## Get media content for the navigation property messages from users
                                            ## @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
                                            ## @return a Fiber of binary
                                            ## 
                                            def get(request_configuration=nil)
                                                request_info = self.to_get_request_information(
                                                    request_configuration
                                                )
                                                return @request_adapter.send_async(request_info, Binary, nil)
                                            end
                                            ## 
                                            ## Update media content for the navigation property messages in users
                                            ## @param body Binary request body
                                            ## @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
                                            ## @return a Fiber of void
                                            ## 
                                            def put(body, request_configuration=nil)
                                                raise StandardError, 'body cannot be null' if body.nil?
                                                request_info = self.to_put_request_information(
                                                    body, request_configuration
                                                )
                                                return @request_adapter.send_async(request_info, nil, nil)
                                            end
                                            ## 
                                            ## Get media content for the navigation property messages from users
                                            ## @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
                                            ## @return a request_information
                                            ## 
                                            def to_get_request_information(request_configuration=nil)
                                                request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
                                                request_info.url_template = @url_template
                                                request_info.path_parameters = @path_parameters
                                                request_info.http_method = :GET
                                                unless request_configuration.nil?
                                                    request_info.add_headers_from_raw_object(request_configuration.headers)
                                                    request_info.add_request_options(request_configuration.options)
                                                end
                                                return request_info
                                            end
                                            ## 
                                            ## Update media content for the navigation property messages in users
                                            ## @param body Binary request body
                                            ## @param request_configuration Configuration for the request such as headers, query parameters, and middleware options.
                                            ## @return a request_information
                                            ## 
                                            def to_put_request_information(body, request_configuration=nil)
                                                raise StandardError, 'body cannot be null' if body.nil?
                                                request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
                                                request_info.url_template = @url_template
                                                request_info.path_parameters = @path_parameters
                                                request_info.http_method = :PUT
                                                unless request_configuration.nil?
                                                    request_info.add_headers_from_raw_object(request_configuration.headers)
                                                    request_info.add_request_options(request_configuration.options)
                                                end
                                                request_info.set_content_from_parsable(self.request_adapter, "", body)
                                                return request_info
                                            end

                                            ## 
                                            # Configuration for the request such as headers, query parameters, and middleware options.
                                            class ContentRequestBuilderGetRequestConfiguration
                                                
                                                ## 
                                                # Request headers
                                                attr_accessor :headers
                                                ## 
                                                # Request options
                                                attr_accessor :options
                                            end

                                            ## 
                                            # Configuration for the request such as headers, query parameters, and middleware options.
                                            class ContentRequestBuilderPutRequestConfiguration
                                                
                                                ## 
                                                # Request headers
                                                attr_accessor :headers
                                                ## 
                                                # Request options
                                                attr_accessor :options
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
