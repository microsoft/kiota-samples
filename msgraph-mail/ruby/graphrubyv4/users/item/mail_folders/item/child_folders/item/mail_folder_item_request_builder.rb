require '../../../../../users'
require '../../../../item'
require '../../../mail_folders'
require '../../item'
require '../child_folders'
require './item'
require 'microsoft_kiota_abstractions'
require_relative '../../../../../../models/mail_folder'
require_relative './message_rules/item/message_rule_item_request_builder'
require_relative './message_rules/message_rules_request_builder'
require_relative './messages/item/message_item_request_builder'
require_relative './messages/messages_request_builder'
require_relative './multi_value_extended_properties/item/multi_value_legacy_extended_property_item_request_builder'
require_relative './multi_value_extended_properties/multi_value_extended_properties_request_builder'
require_relative './single_value_extended_properties/item/single_value_legacy_extended_property_item_request_builder'
require_relative './single_value_extended_properties/single_value_extended_properties_request_builder'

module Graphrubyv4::Users::Item::MailFolders::Item::ChildFolders::Item
    ## 
    # Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\childFolders\{mailFolder-id1}
    class MailFolderItemRequestBuilder
        
        ## 
        # The messageRules property
        def message_rules()
            return Graphrubyv4::Users::Item::MailFolders::Item::ChildFolders::Item::MessageRules::MessageRulesRequestBuilder.new(@path_parameters, @request_adapter)
        end
        ## 
        # The messages property
        def messages()
            return Graphrubyv4::Users::Item::MailFolders::Item::ChildFolders::Item::Messages::MessagesRequestBuilder.new(@path_parameters, @request_adapter)
        end
        ## 
        # The multiValueExtendedProperties property
        def multi_value_extended_properties()
            return Graphrubyv4::Users::Item::MailFolders::Item::ChildFolders::Item::MultiValueExtendedProperties::MultiValueExtendedPropertiesRequestBuilder.new(@path_parameters, @request_adapter)
        end
        ## 
        # Path parameters for the request
        @path_parameters
        ## 
        # The request adapter to use to execute the requests.
        @request_adapter
        ## 
        # The singleValueExtendedProperties property
        def single_value_extended_properties()
            return Graphrubyv4::Users::Item::MailFolders::Item::ChildFolders::Item::SingleValueExtendedProperties::SingleValueExtendedPropertiesRequestBuilder.new(@path_parameters, @request_adapter)
        end
        ## 
        # Url template to use to build the URL for the current request builder
        @url_template
        ## 
        ## Instantiates a new MailFolderItemRequestBuilder and sets the default values.
        ## @param pathParameters Path parameters for the request
        ## @param requestAdapter The request adapter to use to execute the requests.
        ## @return a void
        ## 
        def initialize(path_parameters, request_adapter) 
            @url_template = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}{?%24select,%24expand}"
            @request_adapter = request_adapter
            if path_parameters.is_a? String
                path_parameters = { "request-raw-url" => path_parameters }
            end
            @path_parameters = path_parameters
        end
        ## 
        ## Delete navigation property childFolders for users
        ## @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
        ## @return a request_information
        ## 
        def create_delete_request_information(request_configuration=nil) 
            request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
            request_info.url_template = @url_template
            request_info.path_parameters = @path_parameters
            request_info.http_method = :DELETE
            request_info.set_headers_from_raw_object(request_configuration.headers)
            return request_info;
        end
        ## 
        ## Get childFolders from users
        ## @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
        ## @return a request_information
        ## 
        def create_get_request_information(request_configuration=nil) 
            request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
            request_info.url_template = @url_template
            request_info.path_parameters = @path_parameters
            request_info.http_method = :GET
            request_info.set_headers_from_raw_object(request_configuration.headers)
            request_info.set_query_string_parameters_from_raw_object(request_configuration.query_parameters)
            return request_info;
        end
        ## 
        ## Update the navigation property childFolders in users
        ## @param body 
        ## @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
        ## @return a request_information
        ## 
        def create_patch_request_information(body, request_configuration=nil) 
            request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
            request_info.url_template = @url_template
            request_info.path_parameters = @path_parameters
            request_info.http_method = :PATCH
            request_info.set_headers_from_raw_object(request_configuration.headers)
            request_info.set_content_from_parsable(self.request_adapter, "application/json", body)
            return request_info;
        end
        ## 
        ## Delete navigation property childFolders for users
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
        ## Get childFolders from users
        ## @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of mail_folder
        ## 
        def get(request_configuration=nil, response_handler=nil) 
            request_info = self.create_get_request_information(
                request_configuration
            )
            return @http_core.send_async(request_info, Graphrubyv4::Users::Item::MailFolders::Item::ChildFolders::Item::MailFolder, response_handler)
        end
        ## 
        ## Gets an item from the graphrubyv4.users.item.mailFolders.item.childFolders.item.messageRules.item collection
        ## @param id Unique identifier of the item
        ## @return a message_rule_item_request_builder
        ## 
        def message_rules_by_id(id) 
            url_tpl_params = @path_parameters.clone
            url_tpl_params["messageRule%2Did"] = id
            return Graphrubyv4::Users::Item::MailFolders::Item::ChildFolders::Item::MessageRules::Item::MessageRuleItemRequestBuilder.new(url_tpl_params, @request_adapter)
        end
        ## 
        ## Gets an item from the graphrubyv4.users.item.mailFolders.item.childFolders.item.messages.item collection
        ## @param id Unique identifier of the item
        ## @return a message_item_request_builder
        ## 
        def messages_by_id(id) 
            url_tpl_params = @path_parameters.clone
            url_tpl_params["message%2Did"] = id
            return Graphrubyv4::Users::Item::MailFolders::Item::ChildFolders::Item::Messages::Item::MessageItemRequestBuilder.new(url_tpl_params, @request_adapter)
        end
        ## 
        ## Gets an item from the graphrubyv4.users.item.mailFolders.item.childFolders.item.multiValueExtendedProperties.item collection
        ## @param id Unique identifier of the item
        ## @return a multi_value_legacy_extended_property_item_request_builder
        ## 
        def multi_value_extended_properties_by_id(id) 
            url_tpl_params = @path_parameters.clone
            url_tpl_params["multiValueLegacyExtendedProperty%2Did"] = id
            return Graphrubyv4::Users::Item::MailFolders::Item::ChildFolders::Item::MultiValueExtendedProperties::Item::MultiValueLegacyExtendedPropertyItemRequestBuilder.new(url_tpl_params, @request_adapter)
        end
        ## 
        ## Update the navigation property childFolders in users
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
        ## 
        ## Gets an item from the graphrubyv4.users.item.mailFolders.item.childFolders.item.singleValueExtendedProperties.item collection
        ## @param id Unique identifier of the item
        ## @return a single_value_legacy_extended_property_item_request_builder
        ## 
        def single_value_extended_properties_by_id(id) 
            url_tpl_params = @path_parameters.clone
            url_tpl_params["singleValueLegacyExtendedProperty%2Did"] = id
            return Graphrubyv4::Users::Item::MailFolders::Item::ChildFolders::Item::SingleValueExtendedProperties::Item::SingleValueLegacyExtendedPropertyItemRequestBuilder.new(url_tpl_params, @request_adapter)
        end

        ## 
        # Configuration for the request such as headers, query parameters, and middleware options.
        class MailFolderItemRequestBuilderDeleteRequestConfiguration
            
            ## 
            # Request headers
            @headers
            ## 
            # Request options
            @options
        end

        ## 
        # Get childFolders from users
        class MailFolderItemRequestBuilderGetQueryParameters
            
            ## 
            # Expand related entities
            @expand
            ## 
            # Select properties to be returned
            @select
        end

        ## 
        # Configuration for the request such as headers, query parameters, and middleware options.
        class MailFolderItemRequestBuilderGetRequestConfiguration
            
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

        ## 
        # Configuration for the request such as headers, query parameters, and middleware options.
        class MailFolderItemRequestBuilderPatchRequestConfiguration
            
            ## 
            # Request headers
            @headers
            ## 
            # Request options
            @options
        end
    end
end
