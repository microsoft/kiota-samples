require '../../../../../../../../users'
require '../../../../../../../item'
require '../../../../../../mail_folders'
require '../../../../../item'
require '../../../../child_folders'
require '../../../item'
require '../../messages'
require '../item'
require './value'
require 'microsoft_kiota_abstractions'

module Graphrubyv4::Users::Item::MailFolders::Item::ChildFolders::Item::Messages::Item::Value
    ## 
    # Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\childFolders\{mailFolder-id1}\messages\{message-id}\$value
    class ContentRequestBuilder
        
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
        ## Instantiates a new ContentRequestBuilder and sets the default values.
        ## @param pathParameters Path parameters for the request
        ## @param requestAdapter The request adapter to use to execute the requests.
        ## @return a void
        ## 
        def initialize(path_parameters, request_adapter) 
            @url_template = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}/messages/{message%2Did}/$value"
            @request_adapter = request_adapter
            if path_parameters.is_a? String
                path_parameters = { "request-raw-url" => path_parameters }
            end
            @path_parameters = path_parameters
        end
        ## 
        ## Get media content for the navigation property messages from users
        ## @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
        ## @return a request_information
        ## 
        def create_get_request_information(request_configuration=nil) 
            request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
            request_info.url_template = @url_template
            request_info.path_parameters = @path_parameters
            request_info.http_method = :GET
            request_info.set_headers_from_raw_object(request_configuration.headers)
            return request_info;
        end
        ## 
        ## Update media content for the navigation property messages in users
        ## @param body Binary request body
        ## @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
        ## @return a request_information
        ## 
        def create_put_request_information(body, request_configuration=nil) 
            request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
            request_info.url_template = @url_template
            request_info.path_parameters = @path_parameters
            request_info.http_method = :PUT
            request_info.set_headers_from_raw_object(request_configuration.headers)
            request_info.set_content_from_parsable(self.request_adapter, "", body)
            return request_info;
        end
        ## 
        ## Get media content for the navigation property messages from users
        ## @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of binary
        ## 
        def get(request_configuration=nil, response_handler=nil) 
            request_info = self.create_get_request_information(
                request_configuration
            )
            return @http_core.send_async(request_info, Graphrubyv4::Users::Item::MailFolders::Item::ChildFolders::Item::Messages::Item::Value::Binary, response_handler)
        end
        ## 
        ## Update media content for the navigation property messages in users
        ## @param body Binary request body
        ## @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of void
        ## 
        def put(body, request_configuration=nil, response_handler=nil) 
            request_info = self.create_put_request_information(
                body, request_configuration
            )
            return @http_core.send_async(request_info, nil, response_handler)
        end

        ## 
        # Configuration for the request such as headers, query parameters, and middleware options.
        class ContentRequestBuilderGetRequestConfiguration
            
            ## 
            # Request headers
            @headers
            ## 
            # Request options
            @options
        end

        ## 
        # Configuration for the request such as headers, query parameters, and middleware options.
        class ContentRequestBuilderPutRequestConfiguration
            
            ## 
            # Request headers
            @headers
            ## 
            # Request options
            @options
        end
    end
end
