require '../../../../../users'
require '../../../../item'
require '../../../mail_folders'
require '../../item'
require '../messages'
require './item'
require 'microsoft_kiota_abstractions'
require_relative '../../../../../../models/microsoft/graph/message'
require_relative './attachments/attachments_request_builder'
require_relative './attachments/item/attachment_request_builder'
require_relative './extensions/extensions_request_builder'
require_relative './extensions/item/extension_request_builder'
require_relative './multi_value_extended_properties/item/multi_value_legacy_extended_property_request_builder'
require_relative './multi_value_extended_properties/multi_value_extended_properties_request_builder'
require_relative './single_value_extended_properties/item/single_value_legacy_extended_property_request_builder'
require_relative './single_value_extended_properties/single_value_extended_properties_request_builder'
require_relative './value/content_request_builder'

module Graphrubyv4::Users::Item::MailFolders::Item::Messages::Item
    ## 
    # Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\messages\{message-id}
    class MessageRequestBuilder
        
        def attachments()
            return Graphrubyv4::Users::Item::MailFolders::Item::Messages::Item::Attachments::AttachmentsRequestBuilder.new(@path_parameters, @request_adapter)
        end
        def content()
            return Graphrubyv4::Users::Item::MailFolders::Item::Messages::Item::Value::ContentRequestBuilder.new(@path_parameters, @request_adapter)
        end
        def extensions()
            return Graphrubyv4::Users::Item::MailFolders::Item::Messages::Item::Extensions::ExtensionsRequestBuilder.new(@path_parameters, @request_adapter)
        end
        def multi_value_extended_properties()
            return Graphrubyv4::Users::Item::MailFolders::Item::Messages::Item::MultiValueExtendedProperties::MultiValueExtendedPropertiesRequestBuilder.new(@path_parameters, @request_adapter)
        end
        ## 
        # Path parameters for the request
        @path_parameters
        ## 
        # The request adapter to use to execute the requests.
        @request_adapter
        def single_value_extended_properties()
            return Graphrubyv4::Users::Item::MailFolders::Item::Messages::Item::SingleValueExtendedProperties::SingleValueExtendedPropertiesRequestBuilder.new(@path_parameters, @request_adapter)
        end
        ## 
        # Url template to use to build the URL for the current request builder
        @url_template
        ## 
        ## Gets an item from the graphrubyv4.users.item.mailFolders.item.messages.item.attachments.item collection
        ## @param id Unique identifier of the item
        ## @return a attachment_request_builder
        ## 
        def attachments_by_id(id) 
            url_tpl_params = @path_parameters.clone
            url_tpl_params["attachment_id"] = id
            return Graphrubyv4::Users::Item::MailFolders::Item::Messages::Item::Attachments::Item::AttachmentRequestBuilder.new(url_tpl_params, @request_adapter)
        end
        ## 
        ## Instantiates a new MessageRequestBuilder and sets the default values.
        ## @param pathParameters Path parameters for the request
        ## @param requestAdapter The request adapter to use to execute the requests.
        ## @return a void
        ## 
        def initialize(path_parameters, request_adapter) 
            @url_template = "{+baseurl}/users/{user_id}/mailFolders/{mailFolder_id}/messages/{message_id}{?select,expand}"
            @request_adapter = request_adapter
            if path_parameters.is_a? String
                path_parameters = { "request-raw-url" => path_parameters }
            end
            @path_parameters = path_parameters
        end
        ## 
        ## The collection of messages in the mailFolder.
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
        ## The collection of messages in the mailFolder.
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
        ## The collection of messages in the mailFolder.
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
        ## The collection of messages in the mailFolder.
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
        ## Gets an item from the graphrubyv4.users.item.mailFolders.item.messages.item.extensions.item collection
        ## @param id Unique identifier of the item
        ## @return a extension_request_builder
        ## 
        def extensions_by_id(id) 
            url_tpl_params = @path_parameters.clone
            url_tpl_params["extension_id"] = id
            return Graphrubyv4::Users::Item::MailFolders::Item::Messages::Item::Extensions::Item::ExtensionRequestBuilder.new(url_tpl_params, @request_adapter)
        end
        ## 
        ## The collection of messages in the mailFolder.
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
            return @http_core.send_async(request_info, Graphrubyv4::Users::Item::MailFolders::Item::Messages::Item::Message, response_handler)
        end
        ## 
        ## Gets an item from the graphrubyv4.users.item.mailFolders.item.messages.item.multiValueExtendedProperties.item collection
        ## @param id Unique identifier of the item
        ## @return a multi_value_legacy_extended_property_request_builder
        ## 
        def multi_value_extended_properties_by_id(id) 
            url_tpl_params = @path_parameters.clone
            url_tpl_params["multiValueLegacyExtendedProperty_id"] = id
            return Graphrubyv4::Users::Item::MailFolders::Item::Messages::Item::MultiValueExtendedProperties::Item::MultiValueLegacyExtendedPropertyRequestBuilder.new(url_tpl_params, @request_adapter)
        end
        ## 
        ## The collection of messages in the mailFolder.
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
        ## Gets an item from the graphrubyv4.users.item.mailFolders.item.messages.item.singleValueExtendedProperties.item collection
        ## @param id Unique identifier of the item
        ## @return a single_value_legacy_extended_property_request_builder
        ## 
        def single_value_extended_properties_by_id(id) 
            url_tpl_params = @path_parameters.clone
            url_tpl_params["singleValueLegacyExtendedProperty_id"] = id
            return Graphrubyv4::Users::Item::MailFolders::Item::Messages::Item::SingleValueExtendedProperties::Item::SingleValueLegacyExtendedPropertyRequestBuilder.new(url_tpl_params, @request_adapter)
        end
    end
end
