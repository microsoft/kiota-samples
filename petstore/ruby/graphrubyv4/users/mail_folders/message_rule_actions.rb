require 'microsoft_kiota_abstractions'
require_relative '../users'
require_relative './mail_folders'

module Graphrubyv4::Users::MailFolders
    class MessageRuleActions
        include MicrosoftKiotaAbstractions::Parsable
        ## 
        # Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        @additional_data
        ## 
        # A list of categories to be assigned to a message.
        @assign_categories
        ## 
        # The ID of a folder that a message is to be copied to.
        @copy_to_folder
        ## 
        # Indicates whether a message should be moved to the Deleted Items folder.
        @delete
        ## 
        # The email addresses of the recipients to which a message should be forwarded as an attachment.
        @forward_as_attachment_to
        ## 
        # The email addresses of the recipients to which a message should be forwarded.
        @forward_to
        ## 
        # Indicates whether a message should be marked as read.
        @mark_as_read
        @mark_importance
        ## 
        # The ID of the folder that a message will be moved to.
        @move_to_folder
        ## 
        # Indicates whether a message should be permanently deleted and not saved to the Deleted Items folder.
        @permanent_delete
        ## 
        # The email address to which a message should be redirected.
        @redirect_to
        ## 
        # Indicates whether subsequent rules should be evaluated.
        @stop_processing_rules
        ## 
        ## Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        ## @return a i_dictionary
        ## 
        def  additional_data
            return @additional_data
        end
        ## 
        ## Gets the assignCategories property value. A list of categories to be assigned to a message.
        ## @return a string
        ## 
        def  assign_categories
            return @assign_categories
        end
        ## 
        ## Gets the copyToFolder property value. The ID of a folder that a message is to be copied to.
        ## @return a string
        ## 
        def  copy_to_folder
            return @copy_to_folder
        end
        ## 
        ## Gets the delete property value. Indicates whether a message should be moved to the Deleted Items folder.
        ## @return a boolean
        ## 
        def  delete
            return @delete
        end
        ## 
        ## Gets the forwardAsAttachmentTo property value. The email addresses of the recipients to which a message should be forwarded as an attachment.
        ## @return a recipient
        ## 
        def  forward_as_attachment_to
            return @forward_as_attachment_to
        end
        ## 
        ## Gets the forwardTo property value. The email addresses of the recipients to which a message should be forwarded.
        ## @return a recipient
        ## 
        def  forward_to
            return @forward_to
        end
        ## 
        ## Gets the markAsRead property value. Indicates whether a message should be marked as read.
        ## @return a boolean
        ## 
        def  mark_as_read
            return @mark_as_read
        end
        ## 
        ## Gets the markImportance property value. 
        ## @return a importance
        ## 
        def  mark_importance
            return @mark_importance
        end
        ## 
        ## Gets the moveToFolder property value. The ID of the folder that a message will be moved to.
        ## @return a string
        ## 
        def  move_to_folder
            return @move_to_folder
        end
        ## 
        ## Gets the permanentDelete property value. Indicates whether a message should be permanently deleted and not saved to the Deleted Items folder.
        ## @return a boolean
        ## 
        def  permanent_delete
            return @permanent_delete
        end
        ## 
        ## Gets the redirectTo property value. The email address to which a message should be redirected.
        ## @return a recipient
        ## 
        def  redirect_to
            return @redirect_to
        end
        ## 
        ## Gets the stopProcessingRules property value. Indicates whether subsequent rules should be evaluated.
        ## @return a boolean
        ## 
        def  stop_processing_rules
            return @stop_processing_rules
        end
        ## 
        ## The deserialization information for the current model
        ## @return a i_dictionary
        ## 
        def get_field_deserializers() 
            return {
                "assignCategories" => lambda {|o, n| o.assign_categories = n.get_collection_of_primitive_values(String) },
                "copyToFolder" => lambda {|o, n| o.copy_to_folder = n.get_string_value() },
                "delete" => lambda {|o, n| o.delete = n.get_boolean_value() },
                "forwardAsAttachmentTo" => lambda {|o, n| o.forward_as_attachment_to = n.get_collection_of_object_values(Graphrubyv4::Users::Recipient) },
                "forwardTo" => lambda {|o, n| o.forward_to = n.get_collection_of_object_values(Graphrubyv4::Users::Recipient) },
                "markAsRead" => lambda {|o, n| o.mark_as_read = n.get_boolean_value() },
                "markImportance" => lambda {|o, n| o.mark_importance = n.get_enum_value(Graphrubyv4::Users::Importance) },
                "moveToFolder" => lambda {|o, n| o.move_to_folder = n.get_string_value() },
                "permanentDelete" => lambda {|o, n| o.permanent_delete = n.get_boolean_value() },
                "redirectTo" => lambda {|o, n| o.redirect_to = n.get_collection_of_object_values(Graphrubyv4::Users::Recipient) },
                "stopProcessingRules" => lambda {|o, n| o.stop_processing_rules = n.get_boolean_value() },
            }
        end
        ## 
        ## Serializes information the current object
        ## @param writer Serialization writer to use to serialize this model
        ## @return a void
        ## 
        def serialize(writer) 
            writer.write_collection_of_primitive_values("assignCategories", @assign_categories)
            writer.write_string_value("copyToFolder", @copy_to_folder)
            writer.write_boolean_value("delete", @delete)
            writer.write_collection_of_object_values("forwardAsAttachmentTo", @forward_as_attachment_to)
            writer.write_collection_of_object_values("forwardTo", @forward_to)
            writer.write_boolean_value("markAsRead", @mark_as_read)
            writer.write_enum_value("markImportance", @mark_importance)
            writer.write_string_value("moveToFolder", @move_to_folder)
            writer.write_boolean_value("permanentDelete", @permanent_delete)
            writer.write_collection_of_object_values("redirectTo", @redirect_to)
            writer.write_boolean_value("stopProcessingRules", @stop_processing_rules)
            writer.write_additional_data(@additional_data)
        end
        ## 
        ## Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        ## @param value Value to set for the AdditionalData property.
        ## @return a void
        ## 
        def  additional_data=(additionalData)
            @additional_data = additionalData
        end
        ## 
        ## Sets the assignCategories property value. A list of categories to be assigned to a message.
        ## @param value Value to set for the assignCategories property.
        ## @return a void
        ## 
        def  assign_categories=(assignCategories)
            @assign_categories = assignCategories
        end
        ## 
        ## Sets the copyToFolder property value. The ID of a folder that a message is to be copied to.
        ## @param value Value to set for the copyToFolder property.
        ## @return a void
        ## 
        def  copy_to_folder=(copyToFolder)
            @copy_to_folder = copyToFolder
        end
        ## 
        ## Sets the delete property value. Indicates whether a message should be moved to the Deleted Items folder.
        ## @param value Value to set for the delete property.
        ## @return a void
        ## 
        def  delete=(delete)
            @delete = delete
        end
        ## 
        ## Sets the forwardAsAttachmentTo property value. The email addresses of the recipients to which a message should be forwarded as an attachment.
        ## @param value Value to set for the forwardAsAttachmentTo property.
        ## @return a void
        ## 
        def  forward_as_attachment_to=(forwardAsAttachmentTo)
            @forward_as_attachment_to = forwardAsAttachmentTo
        end
        ## 
        ## Sets the forwardTo property value. The email addresses of the recipients to which a message should be forwarded.
        ## @param value Value to set for the forwardTo property.
        ## @return a void
        ## 
        def  forward_to=(forwardTo)
            @forward_to = forwardTo
        end
        ## 
        ## Sets the markAsRead property value. Indicates whether a message should be marked as read.
        ## @param value Value to set for the markAsRead property.
        ## @return a void
        ## 
        def  mark_as_read=(markAsRead)
            @mark_as_read = markAsRead
        end
        ## 
        ## Sets the markImportance property value. 
        ## @param value Value to set for the markImportance property.
        ## @return a void
        ## 
        def  mark_importance=(markImportance)
            @mark_importance = markImportance
        end
        ## 
        ## Sets the moveToFolder property value. The ID of the folder that a message will be moved to.
        ## @param value Value to set for the moveToFolder property.
        ## @return a void
        ## 
        def  move_to_folder=(moveToFolder)
            @move_to_folder = moveToFolder
        end
        ## 
        ## Sets the permanentDelete property value. Indicates whether a message should be permanently deleted and not saved to the Deleted Items folder.
        ## @param value Value to set for the permanentDelete property.
        ## @return a void
        ## 
        def  permanent_delete=(permanentDelete)
            @permanent_delete = permanentDelete
        end
        ## 
        ## Sets the redirectTo property value. The email address to which a message should be redirected.
        ## @param value Value to set for the redirectTo property.
        ## @return a void
        ## 
        def  redirect_to=(redirectTo)
            @redirect_to = redirectTo
        end
        ## 
        ## Sets the stopProcessingRules property value. Indicates whether subsequent rules should be evaluated.
        ## @param value Value to set for the stopProcessingRules property.
        ## @return a void
        ## 
        def  stop_processing_rules=(stopProcessingRules)
            @stop_processing_rules = stopProcessingRules
        end
    end
end
