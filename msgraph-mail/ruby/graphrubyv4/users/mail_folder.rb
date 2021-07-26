require 'microsoft_kiota_abstractions'
require_relative './entity'
require_relative './users'

module Graphrubyv4::Users
    class MailFolder < Graphrubyv4::Users::Entity
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
        def  child_folder_count
            return @child_folder_count
        end
        ## 
        ## Gets the childFolders property value. The collection of child folders in the mailFolder.
        ## @return a mail_folder
        ## 
        def  child_folders
            return @child_folders
        end
        ## 
        ## Gets the displayName property value. The mailFolder's display name.
        ## @return a string
        ## 
        def  display_name
            return @display_name
        end
        ## 
        ## Gets the isHidden property value. Indicates whether the mailFolder is hidden. This property can be set only when creating the folder. Find more information in Hidden mail folders.
        ## @return a boolean
        ## 
        def  is_hidden
            return @is_hidden
        end
        ## 
        ## Gets the messageRules property value. The collection of rules that apply to the user's Inbox folder.
        ## @return a message_rule
        ## 
        def  message_rules
            return @message_rules
        end
        ## 
        ## Gets the messages property value. The collection of messages in the mailFolder.
        ## @return a message
        ## 
        def  messages
            return @messages
        end
        ## 
        ## Gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable.
        ## @return a multi_value_legacy_extended_property
        ## 
        def  multi_value_extended_properties
            return @multi_value_extended_properties
        end
        ## 
        ## Gets the parentFolderId property value. The unique identifier for the mailFolder's parent mailFolder.
        ## @return a string
        ## 
        def  parent_folder_id
            return @parent_folder_id
        end
        ## 
        ## Gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable.
        ## @return a single_value_legacy_extended_property
        ## 
        def  single_value_extended_properties
            return @single_value_extended_properties
        end
        ## 
        ## Gets the totalItemCount property value. The number of items in the mailFolder.
        ## @return a integer
        ## 
        def  total_item_count
            return @total_item_count
        end
        ## 
        ## Gets the unreadItemCount property value. The number of items in the mailFolder marked as unread.
        ## @return a integer
        ## 
        def  unread_item_count
            return @unread_item_count
        end
        ## 
        ## The deserialization information for the current model
        ## @return a i_dictionary
        ## 
        def get_field_deserializers() 
            return {
                "child_folder_count" => lambda {|o, n| o.child_folder_count = n.get_object_value(integer) },
                "child_folders" => lambda {|o, n| o.child_folders = n.get_collection_of_object_values(mail_folder) },
                "display_name" => lambda {|o, n| o.display_name = n.get_object_value(string) },
                "is_hidden" => lambda {|o, n| o.is_hidden = n.get_object_value(boolean) },
                "message_rules" => lambda {|o, n| o.message_rules = n.get_collection_of_object_values(message_rule) },
                "messages" => lambda {|o, n| o.messages = n.get_collection_of_object_values(message) },
                "multi_value_extended_properties" => lambda {|o, n| o.multi_value_extended_properties = n.get_collection_of_object_values(multi_value_legacy_extended_property) },
                "parent_folder_id" => lambda {|o, n| o.parent_folder_id = n.get_object_value(string) },
                "single_value_extended_properties" => lambda {|o, n| o.single_value_extended_properties = n.get_collection_of_object_values(single_value_legacy_extended_property) },
                "total_item_count" => lambda {|o, n| o.total_item_count = n.get_object_value(integer) },
                "unread_item_count" => lambda {|o, n| o.unread_item_count = n.get_object_value(integer) },
            }
        end
        ## 
        ## Serializes information the current object
        ## @param writer Serialization writer to use to serialize this model
        ## @return a void
        ## 
        def serialize(writer) 
            super.serialize(writer)
            writer.write_object_value("child_folder_count", self.child_folder_count)
            writer.write_collection_of_object_values("child_folders", self.child_folders)
            writer.write_object_value("display_name", self.display_name)
            writer.write_object_value("is_hidden", self.is_hidden)
            writer.write_collection_of_object_values("message_rules", self.message_rules)
            writer.write_collection_of_object_values("messages", self.messages)
            writer.write_collection_of_object_values("multi_value_extended_properties", self.multi_value_extended_properties)
            writer.write_object_value("parent_folder_id", self.parent_folder_id)
            writer.write_collection_of_object_values("single_value_extended_properties", self.single_value_extended_properties)
            writer.write_object_value("total_item_count", self.total_item_count)
            writer.write_object_value("unread_item_count", self.unread_item_count)
        end
        ## 
        ## Sets the childFolderCount property value. The number of immediate child mailFolders in the current mailFolder.
        ## @param value Value to set for the childFolderCount property.
        ## @return a void
        ## 
        def  child_folder_count=(childFolderCount)
            @child_folder_count = (childFolderCount)
        end
        ## 
        ## Sets the childFolders property value. The collection of child folders in the mailFolder.
        ## @param value Value to set for the childFolders property.
        ## @return a void
        ## 
        def  child_folders=(childFolders)
            @child_folders = (childFolders)
        end
        ## 
        ## Sets the displayName property value. The mailFolder's display name.
        ## @param value Value to set for the displayName property.
        ## @return a void
        ## 
        def  display_name=(displayName)
            @display_name = (displayName)
        end
        ## 
        ## Sets the isHidden property value. Indicates whether the mailFolder is hidden. This property can be set only when creating the folder. Find more information in Hidden mail folders.
        ## @param value Value to set for the isHidden property.
        ## @return a void
        ## 
        def  is_hidden=(isHidden)
            @is_hidden = (isHidden)
        end
        ## 
        ## Sets the messageRules property value. The collection of rules that apply to the user's Inbox folder.
        ## @param value Value to set for the messageRules property.
        ## @return a void
        ## 
        def  message_rules=(messageRules)
            @message_rules = (messageRules)
        end
        ## 
        ## Sets the messages property value. The collection of messages in the mailFolder.
        ## @param value Value to set for the messages property.
        ## @return a void
        ## 
        def  messages=(messages)
            @messages = (messages)
        end
        ## 
        ## Sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable.
        ## @param value Value to set for the multiValueExtendedProperties property.
        ## @return a void
        ## 
        def  multi_value_extended_properties=(multiValueExtendedProperties)
            @multi_value_extended_properties = (multiValueExtendedProperties)
        end
        ## 
        ## Sets the parentFolderId property value. The unique identifier for the mailFolder's parent mailFolder.
        ## @param value Value to set for the parentFolderId property.
        ## @return a void
        ## 
        def  parent_folder_id=(parentFolderId)
            @parent_folder_id = (parentFolderId)
        end
        ## 
        ## Sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable.
        ## @param value Value to set for the singleValueExtendedProperties property.
        ## @return a void
        ## 
        def  single_value_extended_properties=(singleValueExtendedProperties)
            @single_value_extended_properties = (singleValueExtendedProperties)
        end
        ## 
        ## Sets the totalItemCount property value. The number of items in the mailFolder.
        ## @param value Value to set for the totalItemCount property.
        ## @return a void
        ## 
        def  total_item_count=(totalItemCount)
            @total_item_count = (totalItemCount)
        end
        ## 
        ## Sets the unreadItemCount property value. The number of items in the mailFolder marked as unread.
        ## @param value Value to set for the unreadItemCount property.
        ## @return a void
        ## 
        def  unread_item_count=(unreadItemCount)
            @unread_item_count = (unreadItemCount)
        end
    end
end
