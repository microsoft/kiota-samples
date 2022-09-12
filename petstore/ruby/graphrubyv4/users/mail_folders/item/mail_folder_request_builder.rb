require 'microsoft_kiota_abstractions'
require_relative '../../users'
require_relative '../child_folders/child_folders_request_builder'
require_relative '../mail_folders'
require_relative '../message_rules/item/message_rule_request_builder'
require_relative '../message_rules/message_rules_request_builder'
require_relative '../messages/item/message_request_builder'
require_relative '../messages/messages_request_builder'
require_relative '../multi_value_extended_properties/item/multi_value_legacy_extended_property_request_builder'
require_relative '../multi_value_extended_properties/multi_value_extended_properties_request_builder'
require_relative '../single_value_extended_properties/item/single_value_legacy_extended_property_request_builder'
require_relative '../single_value_extended_properties/single_value_extended_properties_request_builder'
require_relative './item'

module Graphrubyv4::Users::MailFolders::Item
    ## 
    # Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}
    class MailFolderRequestBuilder
        
        def child_folders()
            return Graphrubyv4::Users::MailFolders::ChildFolders::ChildFoldersRequestBuilder.new(@current_path + @path_segment , @http_core)
        end
        ## 
        # Current path for the request
        @current_path
        ## 
        # The http core service to use to execute the requests.
        @http_core
        def message_rules()
            return Graphrubyv4::Users::MailFolders::MessageRules::MessageRulesRequestBuilder.new(@current_path + @path_segment , @http_core)
        end
        def messages()
            return Graphrubyv4::Users::MailFolders::Messages::MessagesRequestBuilder.new(@current_path + @path_segment , @http_core)
        end
        def multi_value_extended_properties()
            return Graphrubyv4::Users::MailFolders::MultiValueExtendedProperties::MultiValueExtendedPropertiesRequestBuilder.new(@current_path + @path_segment , @http_core)
        end
        ## 
        # Path segment to use to build the URL for the current request builder
        @path_segment
        def single_value_extended_properties()
            return Graphrubyv4::Users::MailFolders::SingleValueExtendedProperties::SingleValueExtendedPropertiesRequestBuilder.new(@current_path + @path_segment , @http_core)
        end
        ## 
        ## Gets an item from the graphrubyv4.users.mailFolders.childFolders collection
        ## @param id Unique identifier of the item
        ## @return a mail_folder_request_builder
        ## 
        def child_folders_by_id(id) 
            return Graphrubyv4::Users::MailFolders::ChildFolders::Item::MailFolderRequestBuilder.new(@current_path + @path_segment  + "/childFolders/" + id, @http_core)
        end
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
        ## Delete navigation property mailFolders for users
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
        ## Get mailFolders from users
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
        ## Update the navigation property mailFolders in users
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
        ## Delete navigation property mailFolders for users
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
        ## Get mailFolders from users
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
            return @http_core.send_async(request_info, Graphrubyv4::Users::MailFolders::Item::MailFolder, response_handler)
        end
        ## 
        ## Gets an item from the graphrubyv4.users.mailFolders.messageRules collection
        ## @param id Unique identifier of the item
        ## @return a message_rule_request_builder
        ## 
        def message_rules_by_id(id) 
            return Graphrubyv4::Users::MailFolders::MessageRules::Item::MessageRuleRequestBuilder.new(@current_path + @path_segment  + "/messageRules/" + id, @http_core)
        end
        ## 
        ## Gets an item from the graphrubyv4.users.mailFolders.messages collection
        ## @param id Unique identifier of the item
        ## @return a message_request_builder
        ## 
        def messages_by_id(id) 
            return Graphrubyv4::Users::MailFolders::Messages::Item::MessageRequestBuilder.new(@current_path + @path_segment  + "/messages/" + id, @http_core)
        end
        ## 
        ## Gets an item from the graphrubyv4.users.mailFolders.multiValueExtendedProperties collection
        ## @param id Unique identifier of the item
        ## @return a multi_value_legacy_extended_property_request_builder
        ## 
        def multi_value_extended_properties_by_id(id) 
            return Graphrubyv4::Users::MailFolders::MultiValueExtendedProperties::Item::MultiValueLegacyExtendedPropertyRequestBuilder.new(@current_path + @path_segment  + "/multiValueExtendedProperties/" + id, @http_core)
        end
        ## 
        ## Update the navigation property mailFolders in users
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
        ## 
        ## Gets an item from the graphrubyv4.users.mailFolders.singleValueExtendedProperties collection
        ## @param id Unique identifier of the item
        ## @return a single_value_legacy_extended_property_request_builder
        ## 
        def single_value_extended_properties_by_id(id) 
            return Graphrubyv4::Users::MailFolders::SingleValueExtendedProperties::Item::SingleValueLegacyExtendedPropertyRequestBuilder.new(@current_path + @path_segment  + "/singleValueExtendedProperties/" + id, @http_core)
        end
    end
end
