require 'microsoft_kiota_abstractions'
require_relative '../graphrubyv4'
require_relative './models'

module Graphrubyv4
    module Models
        class MailFolder < Graphrubyv4::Models::Entity
            include MicrosoftKiotaAbstractions::Parsable
            ## 
            # The number of immediate child mailFolders in the current mailFolder.
            @child_folder_count
            ## 
            # The collection of child folders in the mailFolder.
            @child_folders
            ## 
            # The mailFolder's display name.
            @display_name
            ## 
            # Indicates whether the mailFolder is hidden. This property can be set only when creating the folder. Find more information in Hidden mail folders.
            @is_hidden
            ## 
            # The collection of rules that apply to the user's Inbox folder.
            @message_rules
            ## 
            # The collection of messages in the mailFolder.
            @messages
            ## 
            # The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable.
            @multi_value_extended_properties
            ## 
            # The unique identifier for the mailFolder's parent mailFolder.
            @parent_folder_id
            ## 
            # The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable.
            @single_value_extended_properties
            ## 
            # The number of items in the mailFolder.
            @total_item_count
            ## 
            # The number of items in the mailFolder marked as unread.
            @unread_item_count
            ## 
            ## Gets the childFolderCount property value. The number of immediate child mailFolders in the current mailFolder.
            ## @return a integer
            ## 
            def child_folder_count
                return @child_folder_count
            end
            ## 
            ## Sets the childFolderCount property value. The number of immediate child mailFolders in the current mailFolder.
            ## @param value Value to set for the childFolderCount property.
            ## @return a void
            ## 
            def child_folder_count=(value)
                @child_folder_count = value
            end
            ## 
            ## Gets the childFolders property value. The collection of child folders in the mailFolder.
            ## @return a mail_folder
            ## 
            def child_folders
                return @child_folders
            end
            ## 
            ## Sets the childFolders property value. The collection of child folders in the mailFolder.
            ## @param value Value to set for the childFolders property.
            ## @return a void
            ## 
            def child_folders=(value)
                @child_folders = value
            end
            ## 
            ## Instantiates a new MailFolder and sets the default values.
            ## @return a void
            ## 
            def initialize()
                super
            end
            ## 
            ## Creates a new instance of the appropriate class based on discriminator value
            ## @param parse_node The parse node to use to read the discriminator value and create the object
            ## @return a mail_folder
            ## 
            def self.create_from_discriminator_value(parse_node)
                raise StandardError, 'parse_node cannot be null' if parse_node.nil?
                return MailFolder.new
            end
            ## 
            ## Gets the displayName property value. The mailFolder's display name.
            ## @return a string
            ## 
            def display_name
                return @display_name
            end
            ## 
            ## Sets the displayName property value. The mailFolder's display name.
            ## @param value Value to set for the displayName property.
            ## @return a void
            ## 
            def display_name=(value)
                @display_name = value
            end
            ## 
            ## The deserialization information for the current model
            ## @return a i_dictionary
            ## 
            def get_field_deserializers()
                return super.merge({
                    "childFolderCount" => lambda {|n| @child_folder_count = n.get_number_value() },
                    "childFolders" => lambda {|n| @child_folders = n.get_collection_of_object_values(lambda {|pn| Graphrubyv4::Models::MailFolder.create_from_discriminator_value(pn) }) },
                    "displayName" => lambda {|n| @display_name = n.get_string_value() },
                    "isHidden" => lambda {|n| @is_hidden = n.get_boolean_value() },
                    "messageRules" => lambda {|n| @message_rules = n.get_collection_of_object_values(lambda {|pn| Graphrubyv4::Models::MessageRule.create_from_discriminator_value(pn) }) },
                    "messages" => lambda {|n| @messages = n.get_collection_of_object_values(lambda {|pn| Graphrubyv4::Models::Message.create_from_discriminator_value(pn) }) },
                    "multiValueExtendedProperties" => lambda {|n| @multi_value_extended_properties = n.get_collection_of_object_values(lambda {|pn| Graphrubyv4::Models::MultiValueLegacyExtendedProperty.create_from_discriminator_value(pn) }) },
                    "parentFolderId" => lambda {|n| @parent_folder_id = n.get_string_value() },
                    "singleValueExtendedProperties" => lambda {|n| @single_value_extended_properties = n.get_collection_of_object_values(lambda {|pn| Graphrubyv4::Models::SingleValueLegacyExtendedProperty.create_from_discriminator_value(pn) }) },
                    "totalItemCount" => lambda {|n| @total_item_count = n.get_number_value() },
                    "unreadItemCount" => lambda {|n| @unread_item_count = n.get_number_value() },
                })
            end
            ## 
            ## Gets the isHidden property value. Indicates whether the mailFolder is hidden. This property can be set only when creating the folder. Find more information in Hidden mail folders.
            ## @return a boolean
            ## 
            def is_hidden
                return @is_hidden
            end
            ## 
            ## Sets the isHidden property value. Indicates whether the mailFolder is hidden. This property can be set only when creating the folder. Find more information in Hidden mail folders.
            ## @param value Value to set for the isHidden property.
            ## @return a void
            ## 
            def is_hidden=(value)
                @is_hidden = value
            end
            ## 
            ## Gets the messageRules property value. The collection of rules that apply to the user's Inbox folder.
            ## @return a message_rule
            ## 
            def message_rules
                return @message_rules
            end
            ## 
            ## Sets the messageRules property value. The collection of rules that apply to the user's Inbox folder.
            ## @param value Value to set for the messageRules property.
            ## @return a void
            ## 
            def message_rules=(value)
                @message_rules = value
            end
            ## 
            ## Gets the messages property value. The collection of messages in the mailFolder.
            ## @return a message
            ## 
            def messages
                return @messages
            end
            ## 
            ## Sets the messages property value. The collection of messages in the mailFolder.
            ## @param value Value to set for the messages property.
            ## @return a void
            ## 
            def messages=(value)
                @messages = value
            end
            ## 
            ## Gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable.
            ## @return a multi_value_legacy_extended_property
            ## 
            def multi_value_extended_properties
                return @multi_value_extended_properties
            end
            ## 
            ## Sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable.
            ## @param value Value to set for the multiValueExtendedProperties property.
            ## @return a void
            ## 
            def multi_value_extended_properties=(value)
                @multi_value_extended_properties = value
            end
            ## 
            ## Gets the parentFolderId property value. The unique identifier for the mailFolder's parent mailFolder.
            ## @return a string
            ## 
            def parent_folder_id
                return @parent_folder_id
            end
            ## 
            ## Sets the parentFolderId property value. The unique identifier for the mailFolder's parent mailFolder.
            ## @param value Value to set for the parentFolderId property.
            ## @return a void
            ## 
            def parent_folder_id=(value)
                @parent_folder_id = value
            end
            ## 
            ## Serializes information the current object
            ## @param writer Serialization writer to use to serialize this model
            ## @return a void
            ## 
            def serialize(writer)
                raise StandardError, 'writer cannot be null' if writer.nil?
                super
                writer.write_number_value("childFolderCount", @child_folder_count)
                writer.write_collection_of_object_values("childFolders", @child_folders)
                writer.write_string_value("displayName", @display_name)
                writer.write_boolean_value("isHidden", @is_hidden)
                writer.write_collection_of_object_values("messageRules", @message_rules)
                writer.write_collection_of_object_values("messages", @messages)
                writer.write_collection_of_object_values("multiValueExtendedProperties", @multi_value_extended_properties)
                writer.write_string_value("parentFolderId", @parent_folder_id)
                writer.write_collection_of_object_values("singleValueExtendedProperties", @single_value_extended_properties)
                writer.write_number_value("totalItemCount", @total_item_count)
                writer.write_number_value("unreadItemCount", @unread_item_count)
            end
            ## 
            ## Gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable.
            ## @return a single_value_legacy_extended_property
            ## 
            def single_value_extended_properties
                return @single_value_extended_properties
            end
            ## 
            ## Sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable.
            ## @param value Value to set for the singleValueExtendedProperties property.
            ## @return a void
            ## 
            def single_value_extended_properties=(value)
                @single_value_extended_properties = value
            end
            ## 
            ## Gets the totalItemCount property value. The number of items in the mailFolder.
            ## @return a integer
            ## 
            def total_item_count
                return @total_item_count
            end
            ## 
            ## Sets the totalItemCount property value. The number of items in the mailFolder.
            ## @param value Value to set for the totalItemCount property.
            ## @return a void
            ## 
            def total_item_count=(value)
                @total_item_count = value
            end
            ## 
            ## Gets the unreadItemCount property value. The number of items in the mailFolder marked as unread.
            ## @return a integer
            ## 
            def unread_item_count
                return @unread_item_count
            end
            ## 
            ## Sets the unreadItemCount property value. The number of items in the mailFolder marked as unread.
            ## @param value Value to set for the unreadItemCount property.
            ## @return a void
            ## 
            def unread_item_count=(value)
                @unread_item_count = value
            end
        end
    end
end
