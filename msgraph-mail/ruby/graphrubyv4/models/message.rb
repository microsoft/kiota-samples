require 'date'
require 'microsoft_kiota_abstractions'
require_relative './models'
require_relative './outlook_item'

module Graphrubyv4::Models
    class Message < OutlookItem
        include MicrosoftKiotaAbstractions::Parsable
        ## 
        # The fileAttachment and itemAttachment attachments for the message.
        @attachments
        ## 
        # The Bcc: recipients for the message.
        @bcc_recipients
        ## 
        # The body property
        @body
        ## 
        # The first 255 characters of the message body. It is in text format.
        @body_preview
        ## 
        # The Cc: recipients for the message.
        @cc_recipients
        ## 
        # The ID of the conversation the email belongs to.
        @conversation_id
        ## 
        # Indicates the position of the message within the conversation.
        @conversation_index
        ## 
        # The collection of open extensions defined for the message. Nullable.
        @extensions
        ## 
        # The flag property
        @flag
        ## 
        # The from property
        @from
        ## 
        # Indicates whether the message has attachments. This property doesn't include inline attachments, so if a message contains only inline attachments, this property is false. To verify the existence of inline attachments, parse the body property to look for a src attribute, such as <IMG src='cid:image001.jpg@01D26CD8.6C05F070'>.
        @has_attachments
        ## 
        # The importance property
        @importance
        ## 
        # The inferenceClassification property
        @inference_classification
        ## 
        # The internetMessageHeaders property
        @internet_message_headers
        ## 
        # The internetMessageId property
        @internet_message_id
        ## 
        # The isDeliveryReceiptRequested property
        @is_delivery_receipt_requested
        ## 
        # The isDraft property
        @is_draft
        ## 
        # The isRead property
        @is_read
        ## 
        # The isReadReceiptRequested property
        @is_read_receipt_requested
        ## 
        # The collection of multi-value extended properties defined for the message. Nullable.
        @multi_value_extended_properties
        ## 
        # The parentFolderId property
        @parent_folder_id
        ## 
        # The receivedDateTime property
        @received_date_time
        ## 
        # The replyTo property
        @reply_to
        ## 
        # The sender property
        @sender
        ## 
        # The sentDateTime property
        @sent_date_time
        ## 
        # The collection of single-value extended properties defined for the message. Nullable.
        @single_value_extended_properties
        ## 
        # The subject property
        @subject
        ## 
        # The toRecipients property
        @to_recipients
        ## 
        # The uniqueBody property
        @unique_body
        ## 
        # The webLink property
        @web_link
        ## 
        ## Gets the attachments property value. The fileAttachment and itemAttachment attachments for the message.
        ## @return a attachment
        ## 
        def  attachments
            return @attachments
        end
        ## 
        ## Sets the attachments property value. The fileAttachment and itemAttachment attachments for the message.
        ## @param value Value to set for the attachments property.
        ## @return a void
        ## 
        def  attachments=(attachments)
            @attachments = attachments
        end
        ## 
        ## Gets the bccRecipients property value. The Bcc: recipients for the message.
        ## @return a recipient
        ## 
        def  bcc_recipients
            return @bcc_recipients
        end
        ## 
        ## Sets the bccRecipients property value. The Bcc: recipients for the message.
        ## @param value Value to set for the bccRecipients property.
        ## @return a void
        ## 
        def  bcc_recipients=(bccRecipients)
            @bcc_recipients = bccRecipients
        end
        ## 
        ## Gets the body property value. The body property
        ## @return a item_body
        ## 
        def  body
            return @body
        end
        ## 
        ## Sets the body property value. The body property
        ## @param value Value to set for the body property.
        ## @return a void
        ## 
        def  body=(body)
            @body = body
        end
        ## 
        ## Gets the bodyPreview property value. The first 255 characters of the message body. It is in text format.
        ## @return a string
        ## 
        def  body_preview
            return @body_preview
        end
        ## 
        ## Sets the bodyPreview property value. The first 255 characters of the message body. It is in text format.
        ## @param value Value to set for the bodyPreview property.
        ## @return a void
        ## 
        def  body_preview=(bodyPreview)
            @body_preview = bodyPreview
        end
        ## 
        ## Gets the ccRecipients property value. The Cc: recipients for the message.
        ## @return a recipient
        ## 
        def  cc_recipients
            return @cc_recipients
        end
        ## 
        ## Sets the ccRecipients property value. The Cc: recipients for the message.
        ## @param value Value to set for the ccRecipients property.
        ## @return a void
        ## 
        def  cc_recipients=(ccRecipients)
            @cc_recipients = ccRecipients
        end
        ## 
        ## Gets the conversationId property value. The ID of the conversation the email belongs to.
        ## @return a string
        ## 
        def  conversation_id
            return @conversation_id
        end
        ## 
        ## Sets the conversationId property value. The ID of the conversation the email belongs to.
        ## @param value Value to set for the conversationId property.
        ## @return a void
        ## 
        def  conversation_id=(conversationId)
            @conversation_id = conversationId
        end
        ## 
        ## Gets the conversationIndex property value. Indicates the position of the message within the conversation.
        ## @return a binary
        ## 
        def  conversation_index
            return @conversation_index
        end
        ## 
        ## Sets the conversationIndex property value. Indicates the position of the message within the conversation.
        ## @param value Value to set for the conversationIndex property.
        ## @return a void
        ## 
        def  conversation_index=(conversationIndex)
            @conversation_index = conversationIndex
        end
        ## 
        ## Creates a new instance of the appropriate class based on discriminator value
        ## @param parseNode The parse node to use to read the discriminator value and create the object
        ## @return a message
        ## 
        def create_from_discriminator_value(parse_node) 
            return nil;
        end
        ## 
        ## Gets the extensions property value. The collection of open extensions defined for the message. Nullable.
        ## @return a extension
        ## 
        def  extensions
            return @extensions
        end
        ## 
        ## Sets the extensions property value. The collection of open extensions defined for the message. Nullable.
        ## @param value Value to set for the extensions property.
        ## @return a void
        ## 
        def  extensions=(extensions)
            @extensions = extensions
        end
        ## 
        ## Gets the flag property value. The flag property
        ## @return a followup_flag
        ## 
        def  flag
            return @flag
        end
        ## 
        ## Sets the flag property value. The flag property
        ## @param value Value to set for the flag property.
        ## @return a void
        ## 
        def  flag=(flag)
            @flag = flag
        end
        ## 
        ## Gets the from property value. The from property
        ## @return a recipient
        ## 
        def  from
            return @from
        end
        ## 
        ## Sets the from property value. The from property
        ## @param value Value to set for the from property.
        ## @return a void
        ## 
        def  from=(from)
            @from = from
        end
        ## 
        ## The deserialization information for the current model
        ## @return a i_dictionary
        ## 
        def get_field_deserializers() 
            return super.merge({
                "attachments" => lambda {|n| @attachments = n.get_collection_of_object_values(Graphrubyv4::Models::Attachment) },
                "bccRecipients" => lambda {|n| @bcc_recipients = n.get_collection_of_object_values(Graphrubyv4::Models::Recipient) },
                "body" => lambda {|n| @body = n.get_object_value(Graphrubyv4::Models::ItemBody) },
                "bodyPreview" => lambda {|n| @body_preview = n.get_string_value() },
                "ccRecipients" => lambda {|n| @cc_recipients = n.get_collection_of_object_values(Graphrubyv4::Models::Recipient) },
                "conversationId" => lambda {|n| @conversation_id = n.get_string_value() },
                "conversationIndex" => lambda {|n| @conversation_index = n.get_object_value(::Binary) },
                "extensions" => lambda {|n| @extensions = n.get_collection_of_object_values(Graphrubyv4::Models::Extension) },
                "flag" => lambda {|n| @flag = n.get_object_value(Graphrubyv4::Models::FollowupFlag) },
                "from" => lambda {|n| @from = n.get_object_value(Graphrubyv4::Models::Recipient) },
                "hasAttachments" => lambda {|n| @has_attachments = n.get_boolean_value() },
                "importance" => lambda {|n| @importance = n.get_enum_value(Graphrubyv4::Models::Importance) },
                "inferenceClassification" => lambda {|n| @inference_classification = n.get_enum_value(Graphrubyv4::Models::InferenceClassificationType) },
                "internetMessageHeaders" => lambda {|n| @internet_message_headers = n.get_collection_of_object_values(Graphrubyv4::Models::InternetMessageHeader) },
                "internetMessageId" => lambda {|n| @internet_message_id = n.get_string_value() },
                "isDeliveryReceiptRequested" => lambda {|n| @is_delivery_receipt_requested = n.get_boolean_value() },
                "isDraft" => lambda {|n| @is_draft = n.get_boolean_value() },
                "isRead" => lambda {|n| @is_read = n.get_boolean_value() },
                "isReadReceiptRequested" => lambda {|n| @is_read_receipt_requested = n.get_boolean_value() },
                "multiValueExtendedProperties" => lambda {|n| @multi_value_extended_properties = n.get_collection_of_object_values(Graphrubyv4::Models::MultiValueLegacyExtendedProperty) },
                "parentFolderId" => lambda {|n| @parent_folder_id = n.get_string_value() },
                "receivedDateTime" => lambda {|n| @received_date_time = n.get_date_time_value() },
                "replyTo" => lambda {|n| @reply_to = n.get_collection_of_object_values(Graphrubyv4::Models::Recipient) },
                "sender" => lambda {|n| @sender = n.get_object_value(Graphrubyv4::Models::Recipient) },
                "sentDateTime" => lambda {|n| @sent_date_time = n.get_date_time_value() },
                "singleValueExtendedProperties" => lambda {|n| @single_value_extended_properties = n.get_collection_of_object_values(Graphrubyv4::Models::SingleValueLegacyExtendedProperty) },
                "subject" => lambda {|n| @subject = n.get_string_value() },
                "toRecipients" => lambda {|n| @to_recipients = n.get_collection_of_object_values(Graphrubyv4::Models::Recipient) },
                "uniqueBody" => lambda {|n| @unique_body = n.get_object_value(Graphrubyv4::Models::ItemBody) },
                "webLink" => lambda {|n| @web_link = n.get_string_value() },
            })
        end
        ## 
        ## Gets the hasAttachments property value. Indicates whether the message has attachments. This property doesn't include inline attachments, so if a message contains only inline attachments, this property is false. To verify the existence of inline attachments, parse the body property to look for a src attribute, such as <IMG src='cid:image001.jpg@01D26CD8.6C05F070'>.
        ## @return a boolean
        ## 
        def  has_attachments
            return @has_attachments
        end
        ## 
        ## Sets the hasAttachments property value. Indicates whether the message has attachments. This property doesn't include inline attachments, so if a message contains only inline attachments, this property is false. To verify the existence of inline attachments, parse the body property to look for a src attribute, such as <IMG src='cid:image001.jpg@01D26CD8.6C05F070'>.
        ## @param value Value to set for the hasAttachments property.
        ## @return a void
        ## 
        def  has_attachments=(hasAttachments)
            @has_attachments = hasAttachments
        end
        ## 
        ## Gets the importance property value. The importance property
        ## @return a importance
        ## 
        def  importance
            return @importance
        end
        ## 
        ## Sets the importance property value. The importance property
        ## @param value Value to set for the importance property.
        ## @return a void
        ## 
        def  importance=(importance)
            @importance = importance
        end
        ## 
        ## Gets the inferenceClassification property value. The inferenceClassification property
        ## @return a inference_classification_type
        ## 
        def  inference_classification
            return @inference_classification
        end
        ## 
        ## Sets the inferenceClassification property value. The inferenceClassification property
        ## @param value Value to set for the inferenceClassification property.
        ## @return a void
        ## 
        def  inference_classification=(inferenceClassification)
            @inference_classification = inferenceClassification
        end
        ## 
        ## Gets the internetMessageHeaders property value. The internetMessageHeaders property
        ## @return a internet_message_header
        ## 
        def  internet_message_headers
            return @internet_message_headers
        end
        ## 
        ## Sets the internetMessageHeaders property value. The internetMessageHeaders property
        ## @param value Value to set for the internetMessageHeaders property.
        ## @return a void
        ## 
        def  internet_message_headers=(internetMessageHeaders)
            @internet_message_headers = internetMessageHeaders
        end
        ## 
        ## Gets the internetMessageId property value. The internetMessageId property
        ## @return a string
        ## 
        def  internet_message_id
            return @internet_message_id
        end
        ## 
        ## Sets the internetMessageId property value. The internetMessageId property
        ## @param value Value to set for the internetMessageId property.
        ## @return a void
        ## 
        def  internet_message_id=(internetMessageId)
            @internet_message_id = internetMessageId
        end
        ## 
        ## Gets the isDeliveryReceiptRequested property value. The isDeliveryReceiptRequested property
        ## @return a boolean
        ## 
        def  is_delivery_receipt_requested
            return @is_delivery_receipt_requested
        end
        ## 
        ## Sets the isDeliveryReceiptRequested property value. The isDeliveryReceiptRequested property
        ## @param value Value to set for the isDeliveryReceiptRequested property.
        ## @return a void
        ## 
        def  is_delivery_receipt_requested=(isDeliveryReceiptRequested)
            @is_delivery_receipt_requested = isDeliveryReceiptRequested
        end
        ## 
        ## Gets the isDraft property value. The isDraft property
        ## @return a boolean
        ## 
        def  is_draft
            return @is_draft
        end
        ## 
        ## Sets the isDraft property value. The isDraft property
        ## @param value Value to set for the isDraft property.
        ## @return a void
        ## 
        def  is_draft=(isDraft)
            @is_draft = isDraft
        end
        ## 
        ## Gets the isRead property value. The isRead property
        ## @return a boolean
        ## 
        def  is_read
            return @is_read
        end
        ## 
        ## Sets the isRead property value. The isRead property
        ## @param value Value to set for the isRead property.
        ## @return a void
        ## 
        def  is_read=(isRead)
            @is_read = isRead
        end
        ## 
        ## Gets the isReadReceiptRequested property value. The isReadReceiptRequested property
        ## @return a boolean
        ## 
        def  is_read_receipt_requested
            return @is_read_receipt_requested
        end
        ## 
        ## Sets the isReadReceiptRequested property value. The isReadReceiptRequested property
        ## @param value Value to set for the isReadReceiptRequested property.
        ## @return a void
        ## 
        def  is_read_receipt_requested=(isReadReceiptRequested)
            @is_read_receipt_requested = isReadReceiptRequested
        end
        ## 
        ## Gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the message. Nullable.
        ## @return a multi_value_legacy_extended_property
        ## 
        def  multi_value_extended_properties
            return @multi_value_extended_properties
        end
        ## 
        ## Sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the message. Nullable.
        ## @param value Value to set for the multiValueExtendedProperties property.
        ## @return a void
        ## 
        def  multi_value_extended_properties=(multiValueExtendedProperties)
            @multi_value_extended_properties = multiValueExtendedProperties
        end
        ## 
        ## Gets the parentFolderId property value. The parentFolderId property
        ## @return a string
        ## 
        def  parent_folder_id
            return @parent_folder_id
        end
        ## 
        ## Sets the parentFolderId property value. The parentFolderId property
        ## @param value Value to set for the parentFolderId property.
        ## @return a void
        ## 
        def  parent_folder_id=(parentFolderId)
            @parent_folder_id = parentFolderId
        end
        ## 
        ## Gets the receivedDateTime property value. The receivedDateTime property
        ## @return a date_time
        ## 
        def  received_date_time
            return @received_date_time
        end
        ## 
        ## Sets the receivedDateTime property value. The receivedDateTime property
        ## @param value Value to set for the receivedDateTime property.
        ## @return a void
        ## 
        def  received_date_time=(receivedDateTime)
            @received_date_time = receivedDateTime
        end
        ## 
        ## Gets the replyTo property value. The replyTo property
        ## @return a recipient
        ## 
        def  reply_to
            return @reply_to
        end
        ## 
        ## Sets the replyTo property value. The replyTo property
        ## @param value Value to set for the replyTo property.
        ## @return a void
        ## 
        def  reply_to=(replyTo)
            @reply_to = replyTo
        end
        ## 
        ## Gets the sender property value. The sender property
        ## @return a recipient
        ## 
        def  sender
            return @sender
        end
        ## 
        ## Sets the sender property value. The sender property
        ## @param value Value to set for the sender property.
        ## @return a void
        ## 
        def  sender=(sender)
            @sender = sender
        end
        ## 
        ## Gets the sentDateTime property value. The sentDateTime property
        ## @return a date_time
        ## 
        def  sent_date_time
            return @sent_date_time
        end
        ## 
        ## Sets the sentDateTime property value. The sentDateTime property
        ## @param value Value to set for the sentDateTime property.
        ## @return a void
        ## 
        def  sent_date_time=(sentDateTime)
            @sent_date_time = sentDateTime
        end
        ## 
        ## Serializes information the current object
        ## @param writer Serialization writer to use to serialize this model
        ## @return a void
        ## 
        def serialize(writer) 
            super
            writer.write_collection_of_object_values("attachments", @attachments)
            writer.write_collection_of_object_values("bccRecipients", @bcc_recipients)
            writer.write_object_value("body", @body)
            writer.write_string_value("bodyPreview", @body_preview)
            writer.write_collection_of_object_values("ccRecipients", @cc_recipients)
            writer.write_string_value("conversationId", @conversation_id)
            writer.write_object_value("conversationIndex", @conversation_index)
            writer.write_collection_of_object_values("extensions", @extensions)
            writer.write_object_value("flag", @flag)
            writer.write_object_value("from", @from)
            writer.write_boolean_value("hasAttachments", @has_attachments)
            writer.write_enum_value("importance", @importance)
            writer.write_enum_value("inferenceClassification", @inference_classification)
            writer.write_collection_of_object_values("internetMessageHeaders", @internet_message_headers)
            writer.write_string_value("internetMessageId", @internet_message_id)
            writer.write_boolean_value("isDeliveryReceiptRequested", @is_delivery_receipt_requested)
            writer.write_boolean_value("isDraft", @is_draft)
            writer.write_boolean_value("isRead", @is_read)
            writer.write_boolean_value("isReadReceiptRequested", @is_read_receipt_requested)
            writer.write_collection_of_object_values("multiValueExtendedProperties", @multi_value_extended_properties)
            writer.write_string_value("parentFolderId", @parent_folder_id)
            writer.write_date_time_value("receivedDateTime", @received_date_time)
            writer.write_collection_of_object_values("replyTo", @reply_to)
            writer.write_object_value("sender", @sender)
            writer.write_date_time_value("sentDateTime", @sent_date_time)
            writer.write_collection_of_object_values("singleValueExtendedProperties", @single_value_extended_properties)
            writer.write_string_value("subject", @subject)
            writer.write_collection_of_object_values("toRecipients", @to_recipients)
            writer.write_object_value("uniqueBody", @unique_body)
            writer.write_string_value("webLink", @web_link)
        end
        ## 
        ## Gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the message. Nullable.
        ## @return a single_value_legacy_extended_property
        ## 
        def  single_value_extended_properties
            return @single_value_extended_properties
        end
        ## 
        ## Sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the message. Nullable.
        ## @param value Value to set for the singleValueExtendedProperties property.
        ## @return a void
        ## 
        def  single_value_extended_properties=(singleValueExtendedProperties)
            @single_value_extended_properties = singleValueExtendedProperties
        end
        ## 
        ## Gets the subject property value. The subject property
        ## @return a string
        ## 
        def  subject
            return @subject
        end
        ## 
        ## Sets the subject property value. The subject property
        ## @param value Value to set for the subject property.
        ## @return a void
        ## 
        def  subject=(subject)
            @subject = subject
        end
        ## 
        ## Gets the toRecipients property value. The toRecipients property
        ## @return a recipient
        ## 
        def  to_recipients
            return @to_recipients
        end
        ## 
        ## Sets the toRecipients property value. The toRecipients property
        ## @param value Value to set for the toRecipients property.
        ## @return a void
        ## 
        def  to_recipients=(toRecipients)
            @to_recipients = toRecipients
        end
        ## 
        ## Gets the uniqueBody property value. The uniqueBody property
        ## @return a item_body
        ## 
        def  unique_body
            return @unique_body
        end
        ## 
        ## Sets the uniqueBody property value. The uniqueBody property
        ## @param value Value to set for the uniqueBody property.
        ## @return a void
        ## 
        def  unique_body=(uniqueBody)
            @unique_body = uniqueBody
        end
        ## 
        ## Gets the webLink property value. The webLink property
        ## @return a string
        ## 
        def  web_link
            return @web_link
        end
        ## 
        ## Sets the webLink property value. The webLink property
        ## @param value Value to set for the webLink property.
        ## @return a void
        ## 
        def  web_link=(webLink)
            @web_link = webLink
        end
    end
end
