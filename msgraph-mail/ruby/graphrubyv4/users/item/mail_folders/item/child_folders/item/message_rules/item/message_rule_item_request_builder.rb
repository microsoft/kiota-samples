require '../../../../../../../users'
require '../../../../../../item'
require '../../../../../mail_folders'
require '../../../../item'
require '../../../child_folders'
require '../../item'
require '../message_rules'
require './item'
require 'microsoft_kiota_abstractions'
require_relative '../../../../../../../../models/message_rule'

module GraphC::Users::Item::MailFolders::Item::ChildFolders::Item::MessageRules::Item
    ## 
    # Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\childFolders\{mailFolder-id1}\messageRules\{messageRule-id}
    class MessageRuleItemRequestBuilder
        
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
        ## Instantiates a new MessageRuleItemRequestBuilder and sets the default values.
        ## @param pathParameters Path parameters for the request
        ## @param requestAdapter The request adapter to use to execute the requests.
        ## @return a void
        ## 
        def initialize(path_parameters, request_adapter) 
            @url_template = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}/messageRules/{messageRule%2Did}{?%24select}"
            @request_adapter = request_adapter
            if path_parameters.is_a? String
                path_parameters = { "request-raw-url" => path_parameters }
            end
            @path_parameters = path_parameters
        end
        ## 
        ## Delete navigation property messageRules for users
        ## @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
        ## @return a request_information
        ## 
        def create_delete_request_information(request_configuration=nil) 
            request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
            request_info.url_template = @url_template
            request_info.path_parameters = @path_parameters
            request_info.http_method = :DELETE
            request_info.set_headers_from_raw_object(requestConfiguration.headers)
            return request_info;
        end
        ## 
        ## Get messageRules from users
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
        ## Update the navigation property messageRules in users
        ## @param body 
        ## @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
        ## @return a request_information
        ## 
        def create_patch_request_information(body, request_configuration=nil) 
            request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
            request_info.url_template = @url_template
            request_info.path_parameters = @path_parameters
            request_info.http_method = :PATCH
            request_info.set_headers_from_raw_object(requestConfiguration.headers)
            request_info.set_content_from_parsable(self.request_adapter, "application/json", body)
            return request_info;
        end
        ## 
        ## Delete navigation property messageRules for users
        ## @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of void
        ## 
        def delete(request_configuration=nil, response_handler=nil) 
            request_info = self.create_delete_request_information(
                request_configuration
            )
            return @http_core.send_async(request_info, nil, response_handler)
        end
        ## 
        ## Get messageRules from users
        ## @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of message_rule
        ## 
        def get(request_configuration=nil, response_handler=nil) 
            request_info = self.create_get_request_information(
                request_configuration
            )
            return @http_core.send_async(request_info, GraphC::Users::Item::MailFolders::Item::ChildFolders::Item::MessageRules::Item::MessageRule, response_handler)
        end
        ## 
        ## Update the navigation property messageRules in users
        ## @param body 
        ## @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of void
        ## 
        def patch(body, request_configuration=nil, response_handler=nil) 
            request_info = self.create_patch_request_information(
                body, request_configuration
            )
            return @http_core.send_async(request_info, nil, response_handler)
        end
        require '../../../../../../../users'
        require '../../../../../../item'
        require '../../../../../mail_folders'
        require '../../../../item'
        require '../../../child_folders'
        require '../../item'
        require '../message_rules'
        require './item'

        ## 
        # Configuration for the request such as headers, query parameters, and middleware options.
        class MessageRuleItemRequestBuilderDeleteRequestConfiguration
            
            ## 
            # Request headers
            @headers
            ## 
            # Request options
            @options
        end
        require '../../../../../../../users'
        require '../../../../../../item'
        require '../../../../../mail_folders'
        require '../../../../item'
        require '../../../child_folders'
        require '../../item'
        require '../message_rules'
        require './item'

        ## 
        # Get messageRules from users
        class MessageRuleItemRequestBuilderGetQueryParameters
            
            ## 
            # Select properties to be returned
            @select
        end
        require '../../../../../../../users'
        require '../../../../../../item'
        require '../../../../../mail_folders'
        require '../../../../item'
        require '../../../child_folders'
        require '../../item'
        require '../message_rules'
        require './item'

        ## 
        # Configuration for the request such as headers, query parameters, and middleware options.
        class MessageRuleItemRequestBuilderGetRequestConfiguration
            
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
        require '../../../../../../../users'
        require '../../../../../../item'
        require '../../../../../mail_folders'
        require '../../../../item'
        require '../../../child_folders'
        require '../../item'
        require '../message_rules'
        require './item'

        ## 
        # Configuration for the request such as headers, query parameters, and middleware options.
        class MessageRuleItemRequestBuilderPatchRequestConfiguration
            
            ## 
            # Request headers
            @headers
            ## 
            # Request options
            @options
        end
    end
end
