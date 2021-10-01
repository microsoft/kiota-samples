require '../../../users'
require '../../item'
require '../messages'
require './item'
require 'microsoft_kiota_abstractions'
require_relative '../../../../models/microsoft/graph/message'
require_relative './attachments/attachments_request_builder'
require_relative './attachments/item/attachment_request_builder'
require_relative './extensions/extensions_request_builder'
require_relative './extensions/item/extension_request_builder'
require_relative './multi_value_extended_properties/item/multi_value_legacy_extended_property_request_builder'
require_relative './multi_value_extended_properties/multi_value_extended_properties_request_builder'
require_relative './single_value_extended_properties/item/single_value_legacy_extended_property_request_builder'
require_relative './single_value_extended_properties/single_value_extended_properties_request_builder'
require_relative './value/content_request_builder'

module Graphrubyv4::Users::Item::Messages::Item
    ## 
    # Builds and executes requests for operations under \users\{user-id}\messages\{message-id}
    class MessageRequestBuilder
        
        def attachments()
            return Graphrubyv4::Users::Item::Messages::Item::Attachments::AttachmentsRequestBuilder.new(@current_path + @path_segment , @request_adapter, false)
        end
        def content()
            return Graphrubyv4::Users::Item::Messages::Item::Value::ContentRequestBuilder.new(@current_path + @path_segment , @request_adapter, false)
        end
        ## 
        # Current path for the request
        @current_path
        def extensions()
            return Graphrubyv4::Users::Item::Messages::Item::Extensions::ExtensionsRequestBuilder.new(@current_path + @path_segment , @request_adapter, false)
        end
        ## 
        # Whether the current path is a raw URL
        @is_raw_url
        def multi_value_extended_properties()
            return Graphrubyv4::Users::Item::Messages::Item::MultiValueExtendedProperties::MultiValueExtendedPropertiesRequestBuilder.new(@current_path + @path_segment , @request_adapter, false)
        end
        ## 
        # Path segment to use to build the URL for the current request builder
        @path_segment
        ## 
        # The http core service to use to execute the requests.
        @request_adapter
        def single_value_extended_properties()
            return Graphrubyv4::Users::Item::Messages::Item::SingleValueExtendedProperties::SingleValueExtendedPropertiesRequestBuilder.new(@current_path + @path_segment , @request_adapter, false)
        end
        ## 
        ## Gets an item from the graphrubyv4.users.item.messages.item.attachments.item collection
        ## @param id Unique identifier of the item
        ## @return a attachment_request_builder
        ## 
        def attachments_by_id(id) 
            return Graphrubyv4::Users::Item::Messages::Item::Attachments::Item::AttachmentRequestBuilder.new(@current_path + @path_segment  + "/attachments/" + id, @request_adapter, false)
        end
        ## 
        ## Instantiates a new MessageRequestBuilder and sets the default values.
        ## @param currentPath Current path for the request
        ## @param isRawUrl Whether the current path is a raw URL
        ## @param requestAdapter The http core service to use to execute the requests.
        ## @return a void
        ## 
        def initialize(current_path, request_adapter, is_raw_url=true) 
            @path_segment = ""
            @request_adapter = request_adapter
            @current_path = current_path
            @is_raw_url = is_raw_url
        end
        ## 
        ## The messages in a mailbox or folder. Read-only. Nullable.
        ## @param h Request headers
        ## @param o Request options
        ## @return a request_information
        ## 
        def create_delete_request_information(h=nil, o=nil) 
            request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
            request_info.set_uri(@current_path, @path_segment, @is_raw_url)
            request_info.http_method = :DELETE
            request_info.set_headers_from_raw_object(h)
            return request_info;
        end
        ## 
        ## The messages in a mailbox or folder. Read-only. Nullable.
        ## @param h Request headers
        ## @param o Request options
        ## @param q Request query parameters
        ## @return a request_information
        ## 
        def create_get_request_information(q=nil, h=nil, o=nil) 
            request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
            request_info.set_uri(@current_path, @path_segment, @is_raw_url)
            request_info.http_method = :GET
            request_info.set_headers_from_raw_object(h)
            request_info.set_query_string_parameters_from_raw_object(q)
            return request_info;
        end
        ## 
        ## The messages in a mailbox or folder. Read-only. Nullable.
        ## @param body 
        ## @param h Request headers
        ## @param o Request options
        ## @return a request_information
        ## 
        def create_patch_request_information(body, h=nil, o=nil) 
            request_info = MicrosoftKiotaAbstractions::RequestInformation.new()
            request_info.set_uri(@current_path, @path_segment, @is_raw_url)
            request_info.http_method = :PATCH
            request_info.set_headers_from_raw_object(h)
            request_info.set_content_from_parsable(self.request_adapter, "application/json", body)
            return request_info;
        end
        ## 
        ## The messages in a mailbox or folder. Read-only. Nullable.
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
        ## Gets an item from the graphrubyv4.users.item.messages.item.extensions.item collection
        ## @param id Unique identifier of the item
        ## @return a extension_request_builder
        ## 
        def extensions_by_id(id) 
            return Graphrubyv4::Users::Item::Messages::Item::Extensions::Item::ExtensionRequestBuilder.new(@current_path + @path_segment  + "/extensions/" + id, @request_adapter, false)
        end
        ## 
        ## The messages in a mailbox or folder. Read-only. Nullable.
        ## @param h Request headers
        ## @param o Request options
        ## @param q Request query parameters
        ## @param responseHandler Response handler to use in place of the default response handling provided by the core service
        ## @return a CompletableFuture of message
        ## 
        def get(q=nil, h=nil, o=nil, response_handler=nil) 
            request_info = self.create_get_request_information(
                q, h
            )
            return @http_core.send_async(request_info, Graphrubyv4::Users::Item::Messages::Item::Message, response_handler)
        end
        ## 
        ## Gets an item from the graphrubyv4.users.item.messages.item.multiValueExtendedProperties.item collection
        ## @param id Unique identifier of the item
        ## @return a multi_value_legacy_extended_property_request_builder
        ## 
        def multi_value_extended_properties_by_id(id) 
            return Graphrubyv4::Users::Item::Messages::Item::MultiValueExtendedProperties::Item::MultiValueLegacyExtendedPropertyRequestBuilder.new(@current_path + @path_segment  + "/multiValueExtendedProperties/" + id, @request_adapter, false)
        end
        ## 
        ## The messages in a mailbox or folder. Read-only. Nullable.
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
        ## Gets an item from the graphrubyv4.users.item.messages.item.singleValueExtendedProperties.item collection
        ## @param id Unique identifier of the item
        ## @return a single_value_legacy_extended_property_request_builder
        ## 
        def single_value_extended_properties_by_id(id) 
            return Graphrubyv4::Users::Item::Messages::Item::SingleValueExtendedProperties::Item::SingleValueLegacyExtendedPropertyRequestBuilder.new(@current_path + @path_segment  + "/singleValueExtendedProperties/" + id, @request_adapter, false)
        end
    end
end
