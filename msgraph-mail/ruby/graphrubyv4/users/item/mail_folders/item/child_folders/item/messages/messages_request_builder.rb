require 'microsoft_kiota_abstractions'
require_relative '../../../../../../../models/message'
require_relative '../../../../../../../models/message_collection_response'
require_relative '../../../../../../users'
require_relative '../../../../../item'
require_relative '../../../../mail_folders'
require_relative '../../../item'
require_relative '../../child_folders'
require_relative '../item'
require_relative './messages'

module Graphrubyv4::Users::Item::MailFolders::Item::ChildFolders::Item::Messages
    ## 
    # Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\childFolders\{mailFolder-id1}\messages
    class MessagesRequestBuilder
        
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
        ## Instantiates a new MessagesRequestBuilder and sets the default values.
        ## @param pathParameters Path parameters for the request
        ## @param requestAdapter The request adapter to use to execute the requests.
        ## @return a void
        ## 
        def initialize(path_parameters, request_adapter) 
            @url_template = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}/messages{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}"
            @request_adapter = request_adapter
            if path_parameters.is_a? String
                path_parameters = { "request-raw-url" => path_parameters }
            end
            @path_parameters = path_parameters
        end
        ## 
        ## Get all the messages in the specified user's mailbox, or those messages in a specified folder in the mailbox.
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
                request_info.add_request_options(request_configuration.options)
            end
            return request_info
        end
        ## 
        ## Use this API to create a new Message in a mailfolder.
        ## @param body The request body
        ## @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
        ## @return a request_information
        ## 
        def create_post_request_information(body, request_configuration=nil) 
            request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
            request_info.url_template = @url_template
            request_info.path_parameters = @path_parameters
            request_info.http_method = :POST
            request_info.headers['Accept'] = 'application/json'
            unless request_configuration.nil?
                request_info.set_headers_from_raw_object(request_configuration.headers)
                request_info.add_request_options(request_configuration.options)
            end
            request_info.set_content_from_parsable(self.request_adapter, "application/json", body)
            return request_info
        end
        ## 
        ## Get all the messages in the specified user's mailbox, or those messages in a specified folder in the mailbox.
        ## @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of message_collection_response
        ## 
        def get(request_configuration=nil, response_handler=nil) 
            request_info = self.create_get_request_information(
                request_configuration
            )
            return @request_adapter.send_async(request_info, Graphrubyv4::Models::MessageCollectionResponse, response_handler)
        end
        ## 
        ## Use this API to create a new Message in a mailfolder.
        ## @param body The request body
        ## @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of message
        ## 
        def post(body, request_configuration=nil, response_handler=nil) 
            request_info = self.create_post_request_information(
                body, request_configuration
            )
            return @request_adapter.send_async(request_info, Graphrubyv4::Models::Message, response_handler)
        end

        ## 
        # Get all the messages in the specified user's mailbox, or those messages in a specified folder in the mailbox.
        class MessagesRequestBuilderGetQueryParameters
            
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
            # Order items by property values
            attr_accessor :orderby
            ## 
            # Search items by search phrases
            attr_accessor :search
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
            ## @param originalName The original query parameter name in the class.
            ## @return a string
            ## 
            def get_query_parameter(original_name) 
                case original_name
                    when "count"
                        return "%24count"
                    when "expand"
                        return "%24expand"
                    when "filter"
                        return "%24filter"
                    when "orderby"
                        return "%24orderby"
                    when "search"
                        return "%24search"
                    when "select"
                        return "%24select"
                    when "skip"
                        return "%24skip"
                    when "top"
                        return "%24top"
                    else
                        return originalName
                end
            end
        end

        ## 
        # Configuration for the request such as headers, query parameters, and middleware options.
        class MessagesRequestBuilderGetRequestConfiguration
            
            ## 
            # Request headers
            attr_accessor :headers
            ## 
            # Request options
            attr_accessor :options
            ## 
            # Request query parameters
            attr_accessor :query_parameters
        end

        ## 
        # Configuration for the request such as headers, query parameters, and middleware options.
        class MessagesRequestBuilderPostRequestConfiguration
            
            ## 
            # Request headers
            attr_accessor :headers
            ## 
            # Request options
            attr_accessor :options
        end
    end
end
