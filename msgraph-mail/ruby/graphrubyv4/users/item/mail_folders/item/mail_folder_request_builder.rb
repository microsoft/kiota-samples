require '../../../users'
require '../../item'
require '../mail_folders'
require './item'
require 'microsoft_kiota_abstractions'
require_relative '../../../../models/microsoft/graph/mail_folder'
require_relative './child_folders/child_folders_request_builder'
require_relative './child_folders/item/mail_folder_request_builder'
require_relative './message_rules/item/message_rule_request_builder'
require_relative './message_rules/message_rules_request_builder'
require_relative './messages/item/message_request_builder'
require_relative './messages/messages_request_builder'
require_relative './multi_value_extended_properties/item/multi_value_legacy_extended_property_request_builder'
require_relative './multi_value_extended_properties/multi_value_extended_properties_request_builder'
require_relative './single_value_extended_properties/item/single_value_legacy_extended_property_request_builder'
require_relative './single_value_extended_properties/single_value_extended_properties_request_builder'

module Graphrubyv4::Users::Item::MailFolders::Item
    ## 
    # Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}
    class MailFolderRequestBuilder
        
        def child_folders()
            return Graphrubyv4::Users::Item::MailFolders::Item::ChildFolders::ChildFoldersRequestBuilder.new(@path_parameters, @request_adapter)
        end
        def message_rules()
            return Graphrubyv4::Users::Item::MailFolders::Item::MessageRules::MessageRulesRequestBuilder.new(@path_parameters, @request_adapter)
        end
        def messages()
            return Graphrubyv4::Users::Item::MailFolders::Item::Messages::MessagesRequestBuilder.new(@path_parameters, @request_adapter)
        end
        def multi_value_extended_properties()
            return Graphrubyv4::Users::Item::MailFolders::Item::MultiValueExtendedProperties::MultiValueExtendedPropertiesRequestBuilder.new(@path_parameters, @request_adapter)
        end
        ## 
        # Path parameters for the request
        @path_parameters
        ## 
        # The request adapter to use to execute the requests.
        @request_adapter
        def single_value_extended_properties()
            return Graphrubyv4::Users::Item::MailFolders::Item::SingleValueExtendedProperties::SingleValueExtendedPropertiesRequestBuilder.new(@path_parameters, @request_adapter)
        end
        ## 
        # Url template to use to build the URL for the current request builder
        @url_template
        ## 
        ## Gets an item from the graphrubyv4.users.item.mailFolders.item.childFolders.item collection
        ## @param id Unique identifier of the item
        ## @return a mail_folder_request_builder
        ## 
        def child_folders_by_id(id) 
            url_tpl_params = @path_parameters.clone
            url_tpl_params["mailFolder_id1"] = id
            return Graphrubyv4::Users::Item::MailFolders::Item::ChildFolders::Item::MailFolderRequestBuilder.new(url_tpl_params, @request_adapter)
        end
        ## 
        ## Instantiates a new MailFolderRequestBuilder and sets the default values.
        ## @param pathParameters Path parameters for the request
        ## @param requestAdapter The request adapter to use to execute the requests.
        ## @return a void
        ## 
        def initialize(path_parameters, request_adapter) 
            @url_template = "{+baseurl}/users/{user_id}/mailFolders/{mailFolder_id}{?select}"
            @request_adapter = request_adapter
            if path_parameters.is_a? String
                path_parameters = { "request-raw-url" => path_parameters }
            end
            @path_parameters = path_parameters
        end
        ## 
        ## The user's mail folders. Read-only. Nullable.
        ## @param h Request headers
        ## @param o Request options
        ## @return a request_information
        ## 
        def create_delete_request_information(h=nil, o=nil) 
            request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
            request_info.url_template = @url_template
            request_info.path_parameters = @path_parameters
            request_info.http_method = :DELETE
            request_info.set_headers_from_raw_object(h)
            return request_info;
        end
        ## 
        ## The user's mail folders. Read-only. Nullable.
        ## @param h Request headers
        ## @param o Request options
        ## @param q Request query parameters
        ## @return a request_information
        ## 
        def create_get_request_information(q=nil, h=nil, o=nil) 
            request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
            request_info.url_template = @url_template
            request_info.path_parameters = @path_parameters
            request_info.http_method = :GET
            request_info.set_headers_from_raw_object(h)
            request_info.set_query_string_parameters_from_raw_object(q)
            return request_info;
        end
        ## 
        ## The user's mail folders. Read-only. Nullable.
        ## @param body 
        ## @param h Request headers
        ## @param o Request options
        ## @return a request_information
        ## 
        def create_patch_request_information(body, h=nil, o=nil) 
            request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
            request_info.url_template = @url_template
            request_info.path_parameters = @path_parameters
            request_info.http_method = :PATCH
            request_info.set_headers_from_raw_object(h)
            request_info.set_content_from_parsable(self.request_adapter, "application/json", body)
            return request_info;
        end
        ## 
        ## The user's mail folders. Read-only. Nullable.
        ## @param h Request headers
        ## @param o Request options
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of void
        ## 
        def delete(h=nil, o=nil, response_handler=nil) 
            request_info = self.create_delete_request_information(
                h
            )
            return @http_core.send_async(request_info, nil, response_handler)
        end
        ## 
        ## The user's mail folders. Read-only. Nullable.
        ## @param h Request headers
        ## @param o Request options
        ## @param q Request query parameters
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of mail_folder
        ## 
        def get(q=nil, h=nil, o=nil, response_handler=nil) 
            request_info = self.create_get_request_information(
                q, h
            )
            return @http_core.send_async(request_info, Graphrubyv4::Users::Item::MailFolders::Item::MailFolder, response_handler)
        end
        ## 
        ## Gets an item from the graphrubyv4.users.item.mailFolders.item.messageRules.item collection
        ## @param id Unique identifier of the item
        ## @return a message_rule_request_builder
        ## 
        def message_rules_by_id(id) 
            url_tpl_params = @path_parameters.clone
            url_tpl_params["messageRule_id"] = id
            return Graphrubyv4::Users::Item::MailFolders::Item::MessageRules::Item::MessageRuleRequestBuilder.new(url_tpl_params, @request_adapter)
        end
        ## 
        ## Gets an item from the graphrubyv4.users.item.mailFolders.item.messages.item collection
        ## @param id Unique identifier of the item
        ## @return a message_request_builder
        ## 
        def messages_by_id(id) 
            url_tpl_params = @path_parameters.clone
            url_tpl_params["message_id"] = id
            return Graphrubyv4::Users::Item::MailFolders::Item::Messages::Item::MessageRequestBuilder.new(url_tpl_params, @request_adapter)
        end
        ## 
        ## Gets an item from the graphrubyv4.users.item.mailFolders.item.multiValueExtendedProperties.item collection
        ## @param id Unique identifier of the item
        ## @return a multi_value_legacy_extended_property_request_builder
        ## 
        def multi_value_extended_properties_by_id(id) 
            url_tpl_params = @path_parameters.clone
            url_tpl_params["multiValueLegacyExtendedProperty_id"] = id
            return Graphrubyv4::Users::Item::MailFolders::Item::MultiValueExtendedProperties::Item::MultiValueLegacyExtendedPropertyRequestBuilder.new(url_tpl_params, @request_adapter)
        end
        ## 
        ## The user's mail folders. Read-only. Nullable.
        ## @param body 
        ## @param h Request headers
        ## @param o Request options
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of void
        ## 
        def patch(body, h=nil, o=nil, response_handler=nil) 
            request_info = self.create_patch_request_information(
                body, h
            )
            return @http_core.send_async(request_info, nil, response_handler)
        end
        ## 
        ## Gets an item from the graphrubyv4.users.item.mailFolders.item.singleValueExtendedProperties.item collection
        ## @param id Unique identifier of the item
        ## @return a single_value_legacy_extended_property_request_builder
        ## 
        def single_value_extended_properties_by_id(id) 
            url_tpl_params = @path_parameters.clone
            url_tpl_params["singleValueLegacyExtendedProperty_id"] = id
            return Graphrubyv4::Users::Item::MailFolders::Item::SingleValueExtendedProperties::Item::SingleValueLegacyExtendedPropertyRequestBuilder.new(url_tpl_params, @request_adapter)
        end
    end
end
