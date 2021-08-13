require 'microsoft_kiota_abstractions'
require_relative '../../users'
require_relative '../messages'
require_relative './content'

module Graphrubyv4::Users::Messages::Content
    ## 
    # Builds and executes requests for operations under \users\{user-id}\messages\{message-id}\$value
    class ContentRequestBuilder
        
        ## 
        # Current path for the request
        @current_path
        ## 
        # The http core service to use to execute the requests.
        @http_core
        ## 
        # Path segment to use to build the URL for the current request builder
        @path_segment
        ## 
        ## Instantiates a new ContentRequestBuilder and sets the default values.
        ## @param currentPath Current path for the request
        ## @param httpCore The http core service to use to execute the requests.
        ## @return a void
        ## 
        def initialize(current_path, http_core) 
            @path_segment = "/$value"
            @http_core = http_core
            @current_path = current_path
        end
        ## 
        ## Get media content for the navigation property messages from users
        ## @param h Request headers
        ## @param o Request options for HTTP middlewares
        ## @return a request_info
        ## 
        def create_get_request_info(h=nil, o=nil) 
            request_info = MicrosoftKiotaAbstractions::RequestInfo.new()
            request_info.uri = @current_path + @path_segment
            request_info.http_method = :GET
            request_info.set_headers_from_raw_object(h)
            return request_info;
        end
        ## 
        ## Update media content for the navigation property messages in users
        ## @param body Binary request body
        ## @param h Request headers
        ## @param o Request options for HTTP middlewares
        ## @return a request_info
        ## 
        def create_put_request_info(body, h=nil, o=nil) 
            request_info = MicrosoftKiotaAbstractions::RequestInfo.new()
            request_info.uri = @current_path + @path_segment
            request_info.http_method = :PUT
            request_info.set_headers_from_raw_object(h)
            request_info.set_content_from_parsable(body, self.serializer_factory, "")
            return request_info;
        end
        ## 
        ## Get media content for the navigation property messages from users
        ## @param h Request headers
        ## @param o Request options for HTTP middlewares
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of binary
        ## 
        def get(h=nil, o=nil, response_handler=nil) 
            request_info = self.create_get_request_info(
                h
            )
            return @http_core.send_async(request_info, Graphrubyv4::Users::Messages::Content::Binary, response_handler)
        end
        ## 
        ## Update media content for the navigation property messages in users
        ## @param body Binary request body
        ## @param h Request headers
        ## @param o Request options for HTTP middlewares
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of void
        ## 
        def put(body, h=nil, o=nil, response_handler=nil) 
            request_info = self.create_put_request_info(
                body, h
            )
            return @http_core.send_async(request_info, nil, response_handler)
        end
    end
end
