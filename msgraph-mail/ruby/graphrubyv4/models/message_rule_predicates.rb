require 'microsoft_kiota_abstractions'
require_relative '../graphrubyv4'
require_relative './models'

module Graphrubyv4
    module Models
        class MessageRulePredicates
            include MicrosoftKiotaAbstractions::AdditionalDataHolder, MicrosoftKiotaAbstractions::Parsable
            ## 
            # Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
            @additional_data
            ## 
            # Represents the strings that should appear in the body of an incoming message in order for the condition or exception to apply.
            @body_contains
            ## 
            # Represents the strings that should appear in the body or subject of an incoming message in order for the condition or exception to apply.
            @body_or_subject_contains
            ## 
            # Represents the categories that an incoming message should be labeled with in order for the condition or exception to apply.
            @categories
            ## 
            # Represents the specific sender email addresses of an incoming message in order for the condition or exception to apply.
            @from_addresses
            ## 
            # Indicates whether an incoming message must have attachments in order for the condition or exception to apply.
            @has_attachments
            ## 
            # Represents the strings that appear in the headers of an incoming message in order for the condition or exception to apply.
            @header_contains
            ## 
            # The importance property
            @importance
            ## 
            # Indicates whether an incoming message must be an approval request in order for the condition or exception to apply.
            @is_approval_request
            ## 
            # Indicates whether an incoming message must be automatically forwarded in order for the condition or exception to apply.
            @is_automatic_forward
            ## 
            # Indicates whether an incoming message must be an auto reply in order for the condition or exception to apply.
            @is_automatic_reply
            ## 
            # Indicates whether an incoming message must be encrypted in order for the condition or exception to apply.
            @is_encrypted
            ## 
            # Indicates whether an incoming message must be a meeting request in order for the condition or exception to apply.
            @is_meeting_request
            ## 
            # Indicates whether an incoming message must be a meeting response in order for the condition or exception to apply.
            @is_meeting_response
            ## 
            # Indicates whether an incoming message must be a non-delivery report in order for the condition or exception to apply.
            @is_non_delivery_report
            ## 
            # Indicates whether an incoming message must be permission controlled (RMS-protected) in order for the condition or exception to apply.
            @is_permission_controlled
            ## 
            # Indicates whether an incoming message must be a read receipt in order for the condition or exception to apply.
            @is_read_receipt
            ## 
            # Indicates whether an incoming message must be S/MIME-signed in order for the condition or exception to apply.
            @is_signed
            ## 
            # Indicates whether an incoming message must be a voice mail in order for the condition or exception to apply.
            @is_voicemail
            ## 
            # The messageActionFlag property
            @message_action_flag
            ## 
            # Indicates whether the owner of the mailbox must not be a recipient of an incoming message in order for the condition or exception to apply.
            @not_sent_to_me
            ## 
            # Represents the strings that appear in either the toRecipients or ccRecipients properties of an incoming message in order for the condition or exception to apply.
            @recipient_contains
            ## 
            # Represents the strings that appear in the from property of an incoming message in order for the condition or exception to apply.
            @sender_contains
            ## 
            # The sensitivity property
            @sensitivity
            ## 
            # Indicates whether the owner of the mailbox must be in the ccRecipients property of an incoming message in order for the condition or exception to apply.
            @sent_cc_me
            ## 
            # Indicates whether the owner of the mailbox must be the only recipient in an incoming message in order for the condition or exception to apply.
            @sent_only_to_me
            ## 
            # Represents the email addresses that an incoming message must have been sent to in order for the condition or exception to apply.
            @sent_to_addresses
            ## 
            # Indicates whether the owner of the mailbox must be in the toRecipients property of an incoming message in order for the condition or exception to apply.
            @sent_to_me
            ## 
            # Indicates whether the owner of the mailbox must be in either a toRecipients or ccRecipients property of an incoming message in order for the condition or exception to apply.
            @sent_to_or_cc_me
            ## 
            # Represents the strings that appear in the subject of an incoming message in order for the condition or exception to apply.
            @subject_contains
            ## 
            # The withinSizeRange property
            @within_size_range
            ## 
            ## Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
            ## @return a i_dictionary
            ## 
            def additional_data
                return @additional_data
            end
            ## 
            ## Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
            ## @param value Value to set for the AdditionalData property.
            ## @return a void
            ## 
            def additional_data=(value)
                @additional_data = value
            end
            ## 
            ## Gets the bodyContains property value. Represents the strings that should appear in the body of an incoming message in order for the condition or exception to apply.
            ## @return a string
            ## 
            def body_contains
                return @body_contains
            end
            ## 
            ## Sets the bodyContains property value. Represents the strings that should appear in the body of an incoming message in order for the condition or exception to apply.
            ## @param value Value to set for the bodyContains property.
            ## @return a void
            ## 
            def body_contains=(value)
                @body_contains = value
            end
            ## 
            ## Gets the bodyOrSubjectContains property value. Represents the strings that should appear in the body or subject of an incoming message in order for the condition or exception to apply.
            ## @return a string
            ## 
            def body_or_subject_contains
                return @body_or_subject_contains
            end
            ## 
            ## Sets the bodyOrSubjectContains property value. Represents the strings that should appear in the body or subject of an incoming message in order for the condition or exception to apply.
            ## @param value Value to set for the bodyOrSubjectContains property.
            ## @return a void
            ## 
            def body_or_subject_contains=(value)
                @body_or_subject_contains = value
            end
            ## 
            ## Gets the categories property value. Represents the categories that an incoming message should be labeled with in order for the condition or exception to apply.
            ## @return a string
            ## 
            def categories
                return @categories
            end
            ## 
            ## Sets the categories property value. Represents the categories that an incoming message should be labeled with in order for the condition or exception to apply.
            ## @param value Value to set for the categories property.
            ## @return a void
            ## 
            def categories=(value)
                @categories = value
            end
            ## 
            ## Instantiates a new MessageRulePredicates and sets the default values.
            ## @return a void
            ## 
            def initialize()
                @additional_data = Hash.new
            end
            ## 
            ## Creates a new instance of the appropriate class based on discriminator value
            ## @param parse_node The parse node to use to read the discriminator value and create the object
            ## @return a message_rule_predicates
            ## 
            def self.create_from_discriminator_value(parse_node)
                raise StandardError, 'parse_node cannot be null' if parse_node.nil?
                return MessageRulePredicates.new
            end
            ## 
            ## Gets the fromAddresses property value. Represents the specific sender email addresses of an incoming message in order for the condition or exception to apply.
            ## @return a recipient
            ## 
            def from_addresses
                return @from_addresses
            end
            ## 
            ## Sets the fromAddresses property value. Represents the specific sender email addresses of an incoming message in order for the condition or exception to apply.
            ## @param value Value to set for the fromAddresses property.
            ## @return a void
            ## 
            def from_addresses=(value)
                @from_addresses = value
            end
            ## 
            ## The deserialization information for the current model
            ## @return a i_dictionary
            ## 
            def get_field_deserializers()
                return {
                    "bodyContains" => lambda {|n| @body_contains = n.get_collection_of_primitive_values(String) },
                    "bodyOrSubjectContains" => lambda {|n| @body_or_subject_contains = n.get_collection_of_primitive_values(String) },
                    "categories" => lambda {|n| @categories = n.get_collection_of_primitive_values(String) },
                    "fromAddresses" => lambda {|n| @from_addresses = n.get_collection_of_object_values(lambda {|pn| Graphrubyv4::Models::Recipient.create_from_discriminator_value(pn) }) },
                    "hasAttachments" => lambda {|n| @has_attachments = n.get_boolean_value() },
                    "headerContains" => lambda {|n| @header_contains = n.get_collection_of_primitive_values(String) },
                    "importance" => lambda {|n| @importance = n.get_enum_value(Graphrubyv4::Models::Importance) },
                    "isApprovalRequest" => lambda {|n| @is_approval_request = n.get_boolean_value() },
                    "isAutomaticForward" => lambda {|n| @is_automatic_forward = n.get_boolean_value() },
                    "isAutomaticReply" => lambda {|n| @is_automatic_reply = n.get_boolean_value() },
                    "isEncrypted" => lambda {|n| @is_encrypted = n.get_boolean_value() },
                    "isMeetingRequest" => lambda {|n| @is_meeting_request = n.get_boolean_value() },
                    "isMeetingResponse" => lambda {|n| @is_meeting_response = n.get_boolean_value() },
                    "isNonDeliveryReport" => lambda {|n| @is_non_delivery_report = n.get_boolean_value() },
                    "isPermissionControlled" => lambda {|n| @is_permission_controlled = n.get_boolean_value() },
                    "isReadReceipt" => lambda {|n| @is_read_receipt = n.get_boolean_value() },
                    "isSigned" => lambda {|n| @is_signed = n.get_boolean_value() },
                    "isVoicemail" => lambda {|n| @is_voicemail = n.get_boolean_value() },
                    "messageActionFlag" => lambda {|n| @message_action_flag = n.get_enum_value(Graphrubyv4::Models::MessageActionFlag) },
                    "notSentToMe" => lambda {|n| @not_sent_to_me = n.get_boolean_value() },
                    "recipientContains" => lambda {|n| @recipient_contains = n.get_collection_of_primitive_values(String) },
                    "senderContains" => lambda {|n| @sender_contains = n.get_collection_of_primitive_values(String) },
                    "sensitivity" => lambda {|n| @sensitivity = n.get_enum_value(Graphrubyv4::Models::Sensitivity) },
                    "sentCcMe" => lambda {|n| @sent_cc_me = n.get_boolean_value() },
                    "sentOnlyToMe" => lambda {|n| @sent_only_to_me = n.get_boolean_value() },
                    "sentToAddresses" => lambda {|n| @sent_to_addresses = n.get_collection_of_object_values(lambda {|pn| Graphrubyv4::Models::Recipient.create_from_discriminator_value(pn) }) },
                    "sentToMe" => lambda {|n| @sent_to_me = n.get_boolean_value() },
                    "sentToOrCcMe" => lambda {|n| @sent_to_or_cc_me = n.get_boolean_value() },
                    "subjectContains" => lambda {|n| @subject_contains = n.get_collection_of_primitive_values(String) },
                    "withinSizeRange" => lambda {|n| @within_size_range = n.get_object_value(lambda {|pn| Graphrubyv4::Models::SizeRange.create_from_discriminator_value(pn) }) },
                }
            end
            ## 
            ## Gets the hasAttachments property value. Indicates whether an incoming message must have attachments in order for the condition or exception to apply.
            ## @return a boolean
            ## 
            def has_attachments
                return @has_attachments
            end
            ## 
            ## Sets the hasAttachments property value. Indicates whether an incoming message must have attachments in order for the condition or exception to apply.
            ## @param value Value to set for the hasAttachments property.
            ## @return a void
            ## 
            def has_attachments=(value)
                @has_attachments = value
            end
            ## 
            ## Gets the headerContains property value. Represents the strings that appear in the headers of an incoming message in order for the condition or exception to apply.
            ## @return a string
            ## 
            def header_contains
                return @header_contains
            end
            ## 
            ## Sets the headerContains property value. Represents the strings that appear in the headers of an incoming message in order for the condition or exception to apply.
            ## @param value Value to set for the headerContains property.
            ## @return a void
            ## 
            def header_contains=(value)
                @header_contains = value
            end
            ## 
            ## Gets the importance property value. The importance property
            ## @return a importance
            ## 
            def importance
                return @importance
            end
            ## 
            ## Sets the importance property value. The importance property
            ## @param value Value to set for the importance property.
            ## @return a void
            ## 
            def importance=(value)
                @importance = value
            end
            ## 
            ## Gets the isApprovalRequest property value. Indicates whether an incoming message must be an approval request in order for the condition or exception to apply.
            ## @return a boolean
            ## 
            def is_approval_request
                return @is_approval_request
            end
            ## 
            ## Sets the isApprovalRequest property value. Indicates whether an incoming message must be an approval request in order for the condition or exception to apply.
            ## @param value Value to set for the isApprovalRequest property.
            ## @return a void
            ## 
            def is_approval_request=(value)
                @is_approval_request = value
            end
            ## 
            ## Gets the isAutomaticForward property value. Indicates whether an incoming message must be automatically forwarded in order for the condition or exception to apply.
            ## @return a boolean
            ## 
            def is_automatic_forward
                return @is_automatic_forward
            end
            ## 
            ## Sets the isAutomaticForward property value. Indicates whether an incoming message must be automatically forwarded in order for the condition or exception to apply.
            ## @param value Value to set for the isAutomaticForward property.
            ## @return a void
            ## 
            def is_automatic_forward=(value)
                @is_automatic_forward = value
            end
            ## 
            ## Gets the isAutomaticReply property value. Indicates whether an incoming message must be an auto reply in order for the condition or exception to apply.
            ## @return a boolean
            ## 
            def is_automatic_reply
                return @is_automatic_reply
            end
            ## 
            ## Sets the isAutomaticReply property value. Indicates whether an incoming message must be an auto reply in order for the condition or exception to apply.
            ## @param value Value to set for the isAutomaticReply property.
            ## @return a void
            ## 
            def is_automatic_reply=(value)
                @is_automatic_reply = value
            end
            ## 
            ## Gets the isEncrypted property value. Indicates whether an incoming message must be encrypted in order for the condition or exception to apply.
            ## @return a boolean
            ## 
            def is_encrypted
                return @is_encrypted
            end
            ## 
            ## Sets the isEncrypted property value. Indicates whether an incoming message must be encrypted in order for the condition or exception to apply.
            ## @param value Value to set for the isEncrypted property.
            ## @return a void
            ## 
            def is_encrypted=(value)
                @is_encrypted = value
            end
            ## 
            ## Gets the isMeetingRequest property value. Indicates whether an incoming message must be a meeting request in order for the condition or exception to apply.
            ## @return a boolean
            ## 
            def is_meeting_request
                return @is_meeting_request
            end
            ## 
            ## Sets the isMeetingRequest property value. Indicates whether an incoming message must be a meeting request in order for the condition or exception to apply.
            ## @param value Value to set for the isMeetingRequest property.
            ## @return a void
            ## 
            def is_meeting_request=(value)
                @is_meeting_request = value
            end
            ## 
            ## Gets the isMeetingResponse property value. Indicates whether an incoming message must be a meeting response in order for the condition or exception to apply.
            ## @return a boolean
            ## 
            def is_meeting_response
                return @is_meeting_response
            end
            ## 
            ## Sets the isMeetingResponse property value. Indicates whether an incoming message must be a meeting response in order for the condition or exception to apply.
            ## @param value Value to set for the isMeetingResponse property.
            ## @return a void
            ## 
            def is_meeting_response=(value)
                @is_meeting_response = value
            end
            ## 
            ## Gets the isNonDeliveryReport property value. Indicates whether an incoming message must be a non-delivery report in order for the condition or exception to apply.
            ## @return a boolean
            ## 
            def is_non_delivery_report
                return @is_non_delivery_report
            end
            ## 
            ## Sets the isNonDeliveryReport property value. Indicates whether an incoming message must be a non-delivery report in order for the condition or exception to apply.
            ## @param value Value to set for the isNonDeliveryReport property.
            ## @return a void
            ## 
            def is_non_delivery_report=(value)
                @is_non_delivery_report = value
            end
            ## 
            ## Gets the isPermissionControlled property value. Indicates whether an incoming message must be permission controlled (RMS-protected) in order for the condition or exception to apply.
            ## @return a boolean
            ## 
            def is_permission_controlled
                return @is_permission_controlled
            end
            ## 
            ## Sets the isPermissionControlled property value. Indicates whether an incoming message must be permission controlled (RMS-protected) in order for the condition or exception to apply.
            ## @param value Value to set for the isPermissionControlled property.
            ## @return a void
            ## 
            def is_permission_controlled=(value)
                @is_permission_controlled = value
            end
            ## 
            ## Gets the isReadReceipt property value. Indicates whether an incoming message must be a read receipt in order for the condition or exception to apply.
            ## @return a boolean
            ## 
            def is_read_receipt
                return @is_read_receipt
            end
            ## 
            ## Sets the isReadReceipt property value. Indicates whether an incoming message must be a read receipt in order for the condition or exception to apply.
            ## @param value Value to set for the isReadReceipt property.
            ## @return a void
            ## 
            def is_read_receipt=(value)
                @is_read_receipt = value
            end
            ## 
            ## Gets the isSigned property value. Indicates whether an incoming message must be S/MIME-signed in order for the condition or exception to apply.
            ## @return a boolean
            ## 
            def is_signed
                return @is_signed
            end
            ## 
            ## Sets the isSigned property value. Indicates whether an incoming message must be S/MIME-signed in order for the condition or exception to apply.
            ## @param value Value to set for the isSigned property.
            ## @return a void
            ## 
            def is_signed=(value)
                @is_signed = value
            end
            ## 
            ## Gets the isVoicemail property value. Indicates whether an incoming message must be a voice mail in order for the condition or exception to apply.
            ## @return a boolean
            ## 
            def is_voicemail
                return @is_voicemail
            end
            ## 
            ## Sets the isVoicemail property value. Indicates whether an incoming message must be a voice mail in order for the condition or exception to apply.
            ## @param value Value to set for the isVoicemail property.
            ## @return a void
            ## 
            def is_voicemail=(value)
                @is_voicemail = value
            end
            ## 
            ## Gets the messageActionFlag property value. The messageActionFlag property
            ## @return a message_action_flag
            ## 
            def message_action_flag
                return @message_action_flag
            end
            ## 
            ## Sets the messageActionFlag property value. The messageActionFlag property
            ## @param value Value to set for the messageActionFlag property.
            ## @return a void
            ## 
            def message_action_flag=(value)
                @message_action_flag = value
            end
            ## 
            ## Gets the notSentToMe property value. Indicates whether the owner of the mailbox must not be a recipient of an incoming message in order for the condition or exception to apply.
            ## @return a boolean
            ## 
            def not_sent_to_me
                return @not_sent_to_me
            end
            ## 
            ## Sets the notSentToMe property value. Indicates whether the owner of the mailbox must not be a recipient of an incoming message in order for the condition or exception to apply.
            ## @param value Value to set for the notSentToMe property.
            ## @return a void
            ## 
            def not_sent_to_me=(value)
                @not_sent_to_me = value
            end
            ## 
            ## Gets the recipientContains property value. Represents the strings that appear in either the toRecipients or ccRecipients properties of an incoming message in order for the condition or exception to apply.
            ## @return a string
            ## 
            def recipient_contains
                return @recipient_contains
            end
            ## 
            ## Sets the recipientContains property value. Represents the strings that appear in either the toRecipients or ccRecipients properties of an incoming message in order for the condition or exception to apply.
            ## @param value Value to set for the recipientContains property.
            ## @return a void
            ## 
            def recipient_contains=(value)
                @recipient_contains = value
            end
            ## 
            ## Gets the senderContains property value. Represents the strings that appear in the from property of an incoming message in order for the condition or exception to apply.
            ## @return a string
            ## 
            def sender_contains
                return @sender_contains
            end
            ## 
            ## Sets the senderContains property value. Represents the strings that appear in the from property of an incoming message in order for the condition or exception to apply.
            ## @param value Value to set for the senderContains property.
            ## @return a void
            ## 
            def sender_contains=(value)
                @sender_contains = value
            end
            ## 
            ## Gets the sensitivity property value. The sensitivity property
            ## @return a sensitivity
            ## 
            def sensitivity
                return @sensitivity
            end
            ## 
            ## Sets the sensitivity property value. The sensitivity property
            ## @param value Value to set for the sensitivity property.
            ## @return a void
            ## 
            def sensitivity=(value)
                @sensitivity = value
            end
            ## 
            ## Gets the sentCcMe property value. Indicates whether the owner of the mailbox must be in the ccRecipients property of an incoming message in order for the condition or exception to apply.
            ## @return a boolean
            ## 
            def sent_cc_me
                return @sent_cc_me
            end
            ## 
            ## Sets the sentCcMe property value. Indicates whether the owner of the mailbox must be in the ccRecipients property of an incoming message in order for the condition or exception to apply.
            ## @param value Value to set for the sentCcMe property.
            ## @return a void
            ## 
            def sent_cc_me=(value)
                @sent_cc_me = value
            end
            ## 
            ## Gets the sentOnlyToMe property value. Indicates whether the owner of the mailbox must be the only recipient in an incoming message in order for the condition or exception to apply.
            ## @return a boolean
            ## 
            def sent_only_to_me
                return @sent_only_to_me
            end
            ## 
            ## Sets the sentOnlyToMe property value. Indicates whether the owner of the mailbox must be the only recipient in an incoming message in order for the condition or exception to apply.
            ## @param value Value to set for the sentOnlyToMe property.
            ## @return a void
            ## 
            def sent_only_to_me=(value)
                @sent_only_to_me = value
            end
            ## 
            ## Gets the sentToAddresses property value. Represents the email addresses that an incoming message must have been sent to in order for the condition or exception to apply.
            ## @return a recipient
            ## 
            def sent_to_addresses
                return @sent_to_addresses
            end
            ## 
            ## Sets the sentToAddresses property value. Represents the email addresses that an incoming message must have been sent to in order for the condition or exception to apply.
            ## @param value Value to set for the sentToAddresses property.
            ## @return a void
            ## 
            def sent_to_addresses=(value)
                @sent_to_addresses = value
            end
            ## 
            ## Gets the sentToMe property value. Indicates whether the owner of the mailbox must be in the toRecipients property of an incoming message in order for the condition or exception to apply.
            ## @return a boolean
            ## 
            def sent_to_me
                return @sent_to_me
            end
            ## 
            ## Sets the sentToMe property value. Indicates whether the owner of the mailbox must be in the toRecipients property of an incoming message in order for the condition or exception to apply.
            ## @param value Value to set for the sentToMe property.
            ## @return a void
            ## 
            def sent_to_me=(value)
                @sent_to_me = value
            end
            ## 
            ## Gets the sentToOrCcMe property value. Indicates whether the owner of the mailbox must be in either a toRecipients or ccRecipients property of an incoming message in order for the condition or exception to apply.
            ## @return a boolean
            ## 
            def sent_to_or_cc_me
                return @sent_to_or_cc_me
            end
            ## 
            ## Sets the sentToOrCcMe property value. Indicates whether the owner of the mailbox must be in either a toRecipients or ccRecipients property of an incoming message in order for the condition or exception to apply.
            ## @param value Value to set for the sentToOrCcMe property.
            ## @return a void
            ## 
            def sent_to_or_cc_me=(value)
                @sent_to_or_cc_me = value
            end
            ## 
            ## Serializes information the current object
            ## @param writer Serialization writer to use to serialize this model
            ## @return a void
            ## 
            def serialize(writer)
                raise StandardError, 'writer cannot be null' if writer.nil?
                writer.write_collection_of_primitive_values("bodyContains", @body_contains)
                writer.write_collection_of_primitive_values("bodyOrSubjectContains", @body_or_subject_contains)
                writer.write_collection_of_primitive_values("categories", @categories)
                writer.write_collection_of_object_values("fromAddresses", @from_addresses)
                writer.write_boolean_value("hasAttachments", @has_attachments)
                writer.write_collection_of_primitive_values("headerContains", @header_contains)
                writer.write_enum_value("importance", @importance)
                writer.write_boolean_value("isApprovalRequest", @is_approval_request)
                writer.write_boolean_value("isAutomaticForward", @is_automatic_forward)
                writer.write_boolean_value("isAutomaticReply", @is_automatic_reply)
                writer.write_boolean_value("isEncrypted", @is_encrypted)
                writer.write_boolean_value("isMeetingRequest", @is_meeting_request)
                writer.write_boolean_value("isMeetingResponse", @is_meeting_response)
                writer.write_boolean_value("isNonDeliveryReport", @is_non_delivery_report)
                writer.write_boolean_value("isPermissionControlled", @is_permission_controlled)
                writer.write_boolean_value("isReadReceipt", @is_read_receipt)
                writer.write_boolean_value("isSigned", @is_signed)
                writer.write_boolean_value("isVoicemail", @is_voicemail)
                writer.write_enum_value("messageActionFlag", @message_action_flag)
                writer.write_boolean_value("notSentToMe", @not_sent_to_me)
                writer.write_collection_of_primitive_values("recipientContains", @recipient_contains)
                writer.write_collection_of_primitive_values("senderContains", @sender_contains)
                writer.write_enum_value("sensitivity", @sensitivity)
                writer.write_boolean_value("sentCcMe", @sent_cc_me)
                writer.write_boolean_value("sentOnlyToMe", @sent_only_to_me)
                writer.write_collection_of_object_values("sentToAddresses", @sent_to_addresses)
                writer.write_boolean_value("sentToMe", @sent_to_me)
                writer.write_boolean_value("sentToOrCcMe", @sent_to_or_cc_me)
                writer.write_collection_of_primitive_values("subjectContains", @subject_contains)
                writer.write_object_value("withinSizeRange", @within_size_range)
                writer.write_additional_data(@additional_data)
            end
            ## 
            ## Gets the subjectContains property value. Represents the strings that appear in the subject of an incoming message in order for the condition or exception to apply.
            ## @return a string
            ## 
            def subject_contains
                return @subject_contains
            end
            ## 
            ## Sets the subjectContains property value. Represents the strings that appear in the subject of an incoming message in order for the condition or exception to apply.
            ## @param value Value to set for the subjectContains property.
            ## @return a void
            ## 
            def subject_contains=(value)
                @subject_contains = value
            end
            ## 
            ## Gets the withinSizeRange property value. The withinSizeRange property
            ## @return a size_range
            ## 
            def within_size_range
                return @within_size_range
            end
            ## 
            ## Sets the withinSizeRange property value. The withinSizeRange property
            ## @param value Value to set for the withinSizeRange property.
            ## @return a void
            ## 
            def within_size_range=(value)
                @within_size_range = value
            end
        end
    end
end
