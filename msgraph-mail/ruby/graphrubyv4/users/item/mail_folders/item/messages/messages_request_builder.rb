require '../../../../users'
require '../../../item'
require '../../mail_folders'
require '../item'
require './messages'
require 'microsoft_kiota_abstractions'
require_relative '../../../../../models/message'

module Graphrubyv4::Users::Item::MailFolders::Item::Messages
    ## 
    # Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\messages
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
            @url_template = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/messages{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}"
            @request_adapter = request_adapter
            if path_parameters.is_a? String
                path_parameters = { "request-raw-url" => path_parameters }
            end
            @path_parameters = path_parameters
        end
        ## 
        ## Get messages from users
        ## @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
        ## @return a request_information
        ## 
        def create_get_request_information(request_configuration=nil) 
            request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
            request_info.url_template = @url_template
            request_info.path_parameters = @path_parameters
            request_info.http_method = :GET
            request_info.set_headers_from_raw_object(requestConfiguration.headers)
            request_info.set_query_string_parameters_from_raw_object(requestConfiguration.query_parameters)
            return request_info;
        end
        ## 
        ## Create new navigation property to messages for users
        ## @param body 
        ## @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
        ## @return a request_information
        ## 
        def create_post_request_information(body, request_configuration=nil) 
            request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
            request_info.url_template = @url_template
            request_info.path_parameters = @path_parameters
            request_info.http_method = :POST
            request_info.set_headers_from_raw_object(requestConfiguration.headers)
            request_info.set_content_from_parsable(self.request_adapter, "application/json", body)
            return request_info;
        end
        ## 
        ## Get messages from users
        ## @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of message_collection_response
        ## 
        def get(request_configuration=nil, response_handler=nil) 
            request_info = self.create_get_request_information(
                request_configuration
            )
            return @http_core.send_async(request_info, Graphrubyv4::Users::Item::MailFolders::Item::Messages::MessageCollectionResponse, response_handler)
        end
        ## 
        ## Create new navigation property to messages for users
        ## @param body 
        ## @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of message
        ## 
        def post(body, request_configuration=nil, response_handler=nil) 
            request_info = self.create_post_request_information(
                body, request_configuration
            )
            return @http_core.send_async(request_info, Graphrubyv4::Users::Item::MailFolders::Item::Messages::Message, response_handler)
        end
        require '../../../../users'
        require '../../../item'
        require '../../mail_folders'
        require '../item'
        require './messages'

        ## 
        # Get messages from users
        class MessagesRequestBuilderGetQueryParameters
            
            ## 
            # Include count of items
            @count
            ## 
            # Expand related entities
            @expand
            ## 
            # Filter items by property values
            @filter
            ## 
            # Order items by property values
            @orderby
            ## 
            # Search items by search phrases
            @search
            ## 
            # Select properties to be returned
            @select
            ## 
            # Skip the first n items
            @skip
            ## 
            # Show only the first n items
            @top
        end
        require '../../../../users'
        require '../../../item'
        require '../../mail_folders'
        require '../item'
        require './messages'

        ## 
        # Configuration for the request such as headers, query parameters, and middleware options.
        class MessagesRequestBuilderGetRequestConfiguration
            
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
        require '../../../../users'
        require '../../../item'
        require '../../mail_folders'
        require '../item'
        require './messages'

        ## 
        # Configuration for the request such as headers, query parameters, and middleware options.
        class MessagesRequestBuilderPostRequestConfiguration
            
            ## 
            # Request headers
            @headers
            ## 
            # Request options
            @options
        end
    end
end
