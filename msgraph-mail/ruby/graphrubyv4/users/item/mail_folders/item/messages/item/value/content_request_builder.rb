require '../../../../../../users'
require '../../../../../item'
require '../../../../mail_folders'
require '../../../item'
require '../../messages'
require '../item'
require './value'
require 'microsoft_kiota_abstractions'

module Graphrubyv4::Users::Item::MailFolders::Item::Messages::Item::Value
    ## 
    # Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\messages\{message-id}\$value
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
            @url_template = "{+baseurl}/users/{user_id}/mailFolders/{mailFolder_id}/messages/{message_id}/$value"
            @request_adapter = request_adapter
            if path_parameters.is_a? String
                path_parameters = { "request-raw-url" => path_parameters }
            end
            @path_parameters = path_parameters
        end
        ## 
        ## Get media content for the navigation property messages from users
        ## @param h Request headers
        ## @param o Request options
        ## @return a request_information
        ## 
        def create_get_request_information(h=nil, o=nil) 
            request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
            request_info.url_template = @url_template
            request_info.path_parameters = @path_parameters
            request_info.http_method = :GET
            request_info.set_headers_from_raw_object(h)
            return request_info;
        end
        ## 
        ## Update media content for the navigation property messages in users
        ## @param body Binary request body
        ## @param h Request headers
        ## @param o Request options
        ## @return a request_information
        ## 
        def create_put_request_information(body, h=nil, o=nil) 
            request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
            request_info.url_template = @url_template
            request_info.path_parameters = @path_parameters
            request_info.http_method = :PUT
            request_info.set_headers_from_raw_object(h)
            request_info.set_content_from_parsable(self.request_adapter, "", body)
            return request_info;
        end
        ## 
        ## Get media content for the navigation property messages from users
        ## @param h Request headers
        ## @param o Request options
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of binary
        ## 
        def get(h=nil, o=nil, response_handler=nil) 
            request_info = self.create_get_request_information(
                h
            )
            return @http_core.send_async(request_info, Graphrubyv4::Users::Item::MailFolders::Item::Messages::Item::Value::Binary, response_handler)
        end
        ## 
        ## Update media content for the navigation property messages in users
        ## @param body Binary request body
        ## @param h Request headers
        ## @param o Request options
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of void
        ## 
        def put(body, h=nil, o=nil, response_handler=nil) 
            request_info = self.create_put_request_information(
                body, h
            )
            return @http_core.send_async(request_info, nil, response_handler)
        end
    end
end
