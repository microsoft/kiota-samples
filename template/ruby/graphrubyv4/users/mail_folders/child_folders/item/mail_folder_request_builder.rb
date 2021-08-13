require 'microsoft_kiota_abstractions'
require_relative '../../../users'
require_relative '../../mail_folders'
require_relative '../child_folders'
require_relative './item'

module Graphrubyv4::Users::MailFolders::ChildFolders::Item
    ## 
    # Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\childFolders\{mailFolder-id1}
    class MailFolderRequestBuilder
        
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
        ## Instantiates a new MailFolderRequestBuilder and sets the default values.
        ## @param currentPath Current path for the request
        ## @param httpCore The http core service to use to execute the requests.
        ## @return a void
        ## 
        def initialize(current_path, http_core) 
            @path_segment = ""
            @http_core = http_core
            @current_path = current_path
        end
        ## 
        ## Delete navigation property childFolders for users
        ## @param h Request headers
        ## @param o Request options for HTTP middlewares
        ## @return a request_info
        ## 
        def create_delete_request_info(h=nil, o=nil) 
            request_info = MicrosoftKiotaAbstractions::RequestInfo.new()
            request_info.uri = @current_path + @path_segment
            request_info.http_method = :DELETE
            request_info.set_headers_from_raw_object(h)
            return request_info;
        end
        ## 
        ## Get childFolders from users
        ## @param h Request headers
        ## @param o Request options for HTTP middlewares
        ## @param q Request query parameters
        ## @return a request_info
        ## 
        def create_get_request_info(q=nil, h=nil, o=nil) 
            request_info = MicrosoftKiotaAbstractions::RequestInfo.new()
            request_info.uri = @current_path + @path_segment
            request_info.http_method = :GET
            request_info.set_headers_from_raw_object(h)
            request_info.set_query_string_parameters_from_raw_object(q)
            return request_info;
        end
        ## 
        ## Update the navigation property childFolders in users
        ## @param body 
        ## @param h Request headers
        ## @param o Request options for HTTP middlewares
        ## @return a request_info
        ## 
        def create_patch_request_info(body, h=nil, o=nil) 
            request_info = MicrosoftKiotaAbstractions::RequestInfo.new()
            request_info.uri = @current_path + @path_segment
            request_info.http_method = :PATCH
            request_info.set_headers_from_raw_object(h)
            request_info.set_content_from_parsable(body, self.serializer_factory, "application/json")
            return request_info;
        end
        ## 
        ## Delete navigation property childFolders for users
        ## @param h Request headers
        ## @param o Request options for HTTP middlewares
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of void
        ## 
        def delete(h=nil, o=nil, response_handler=nil) 
            request_info = self.create_delete_request_info(
                h
            )
            return @http_core.send_async(request_info, nil, response_handler)
        end
        ## 
        ## Get childFolders from users
        ## @param h Request headers
        ## @param o Request options for HTTP middlewares
        ## @param q Request query parameters
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of mail_folder
        ## 
        def get(q=nil, h=nil, o=nil, response_handler=nil) 
            request_info = self.create_get_request_info(
                q, h
            )
            return @http_core.send_async(request_info, Graphrubyv4::Users::MailFolders::ChildFolders::Item::MailFolder, response_handler)
        end
        ## 
        ## Update the navigation property childFolders in users
        ## @param body 
        ## @param h Request headers
        ## @param o Request options for HTTP middlewares
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of void
        ## 
        def patch(body, h=nil, o=nil, response_handler=nil) 
            request_info = self.create_patch_request_info(
                body, h
            )
            return @http_core.send_async(request_info, nil, response_handler)
        end
    end
end
