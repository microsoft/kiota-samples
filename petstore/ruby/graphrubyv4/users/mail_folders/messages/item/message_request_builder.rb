require 'microsoft_kiota_abstractions'
require_relative '../../../users'
require_relative '../../mail_folders'
require_relative '../attachments/attachments_request_builder'
require_relative '../attachments/item/attachment_request_builder'
require_relative '../content/content_request_builder'
require_relative '../extensions/extensions_request_builder'
require_relative '../extensions/item/extension_request_builder'
require_relative '../messages'
require_relative '../multi_value_extended_properties/item/multi_value_legacy_extended_property_request_builder'
require_relative '../multi_value_extended_properties/multi_value_extended_properties_request_builder'
require_relative '../single_value_extended_properties/item/single_value_legacy_extended_property_request_builder'
require_relative '../single_value_extended_properties/single_value_extended_properties_request_builder'
require_relative './item'

module Graphrubyv4::Users::MailFolders::Messages::Item
    ## 
    # Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\messages\{message-id}
    class MessageRequestBuilder
        
        def attachments()
            return Graphrubyv4::Users::MailFolders::Messages::Attachments::AttachmentsRequestBuilder.new(@current_path + @path_segment , @http_core)
        end
        def content()
            return Graphrubyv4::Users::MailFolders::Messages::Content::ContentRequestBuilder.new(@current_path + @path_segment , @http_core)
        end
        ## 
        # Current path for the request
        @current_path
        def extensions()
            return Graphrubyv4::Users::MailFolders::Messages::Extensions::ExtensionsRequestBuilder.new(@current_path + @path_segment , @http_core)
        end
        ## 
        # The http core service to use to execute the requests.
        @http_core
        def multi_value_extended_properties()
            return Graphrubyv4::Users::MailFolders::Messages::MultiValueExtendedProperties::MultiValueExtendedPropertiesRequestBuilder.new(@current_path + @path_segment , @http_core)
        end
        ## 
        # Path segment to use to build the URL for the current request builder
        @path_segment
        def single_value_extended_properties()
            return Graphrubyv4::Users::MailFolders::Messages::SingleValueExtendedProperties::SingleValueExtendedPropertiesRequestBuilder.new(@current_path + @path_segment , @http_core)
        end
        ## 
        ## Gets an item from the graphrubyv4.users.mailFolders.messages.attachments collection
        ## @param id Unique identifier of the item
        ## @return a attachment_request_builder
        ## 
        def attachments_by_id(id) 
            return Graphrubyv4::Users::MailFolders::Messages::Attachments::Item::AttachmentRequestBuilder.new(@current_path + @path_segment  + "/attachments/" + id, @http_core)
        end
        ## 
        ## Instantiates a new MessageRequestBuilder and sets the default values.
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
        ## Delete navigation property messages for users
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
        ## Get messages from users
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
        ## Update the navigation property messages in users
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
        ## Delete navigation property messages for users
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
        ## Gets an item from the graphrubyv4.users.mailFolders.messages.extensions collection
        ## @param id Unique identifier of the item
        ## @return a extension_request_builder
        ## 
        def extensions_by_id(id) 
            return Graphrubyv4::Users::MailFolders::Messages::Extensions::Item::ExtensionRequestBuilder.new(@current_path + @path_segment  + "/extensions/" + id, @http_core)
        end
        ## 
        ## Get messages from users
        ## @param h Request headers
        ## @param o Request options for HTTP middlewares
        ## @param q Request query parameters
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of message
        ## 
        def get(q=nil, h=nil, o=nil, response_handler=nil) 
            request_info = self.create_get_request_info(
                q, h
            )
            return @http_core.send_async(request_info, Graphrubyv4::Users::MailFolders::Messages::Item::Message, response_handler)
        end
        ## 
        ## Gets an item from the graphrubyv4.users.mailFolders.messages.multiValueExtendedProperties collection
        ## @param id Unique identifier of the item
        ## @return a multi_value_legacy_extended_property_request_builder
        ## 
        def multi_value_extended_properties_by_id(id) 
            return Graphrubyv4::Users::MailFolders::Messages::MultiValueExtendedProperties::Item::MultiValueLegacyExtendedPropertyRequestBuilder.new(@current_path + @path_segment  + "/multiValueExtendedProperties/" + id, @http_core)
        end
        ## 
        ## Update the navigation property messages in users
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
        ## Gets an item from the graphrubyv4.users.mailFolders.messages.singleValueExtendedProperties collection
        ## @param id Unique identifier of the item
        ## @return a single_value_legacy_extended_property_request_builder
        ## 
        def single_value_extended_properties_by_id(id) 
            return Graphrubyv4::Users::MailFolders::Messages::SingleValueExtendedProperties::Item::SingleValueLegacyExtendedPropertyRequestBuilder.new(@current_path + @path_segment  + "/singleValueExtendedProperties/" + id, @http_core)
        end
    end
end
