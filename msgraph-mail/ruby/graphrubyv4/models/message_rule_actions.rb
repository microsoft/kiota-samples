require 'microsoft_kiota_abstractions'
require_relative '../graphrubyv4'
require_relative './models'

module Graphrubyv4
    module Models
        class MessageRuleActions
            include MicrosoftKiotaAbstractions::AdditionalDataHolder, MicrosoftKiotaAbstractions::Parsable
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
            ## 
            # The markImportance property
            @mark_importance
            ## 
            # The ID of the folder that a message will be moved to.
            @move_to_folder
            ## 
            # Indicates whether a message should be permanently deleted and not saved to the Deleted Items folder.
            @permanent_delete
            ## 
            # The email addresses to which a message should be redirected.
            @redirect_to
            ## 
            # Indicates whether subsequent rules should be evaluated.
            @stop_processing_rules
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
            ## Gets the assignCategories property value. A list of categories to be assigned to a message.
            ## @return a string
            ## 
            def assign_categories
                return @assign_categories
            end
            ## 
            ## Sets the assignCategories property value. A list of categories to be assigned to a message.
            ## @param value Value to set for the assignCategories property.
            ## @return a void
            ## 
            def assign_categories=(value)
                @assign_categories = value
            end
            ## 
            ## Instantiates a new MessageRuleActions and sets the default values.
            ## @return a void
            ## 
            def initialize()
                @additional_data = Hash.new
            end
            ## 
            ## Gets the copyToFolder property value. The ID of a folder that a message is to be copied to.
            ## @return a string
            ## 
            def copy_to_folder
                return @copy_to_folder
            end
            ## 
            ## Sets the copyToFolder property value. The ID of a folder that a message is to be copied to.
            ## @param value Value to set for the copyToFolder property.
            ## @return a void
            ## 
            def copy_to_folder=(value)
                @copy_to_folder = value
            end
            ## 
            ## Creates a new instance of the appropriate class based on discriminator value
            ## @param parse_node The parse node to use to read the discriminator value and create the object
            ## @return a message_rule_actions
            ## 
            def self.create_from_discriminator_value(parse_node)
                raise StandardError, 'parse_node cannot be null' if parse_node.nil?
                return MessageRuleActions.new
            end
            ## 
            ## Gets the delete property value. Indicates whether a message should be moved to the Deleted Items folder.
            ## @return a boolean
            ## 
            def delete
                return @delete
            end
            ## 
            ## Sets the delete property value. Indicates whether a message should be moved to the Deleted Items folder.
            ## @param value Value to set for the delete property.
            ## @return a void
            ## 
            def delete=(value)
                @delete = value
            end
            ## 
            ## Gets the forwardAsAttachmentTo property value. The email addresses of the recipients to which a message should be forwarded as an attachment.
            ## @return a recipient
            ## 
            def forward_as_attachment_to
                return @forward_as_attachment_to
            end
            ## 
            ## Sets the forwardAsAttachmentTo property value. The email addresses of the recipients to which a message should be forwarded as an attachment.
            ## @param value Value to set for the forwardAsAttachmentTo property.
            ## @return a void
            ## 
            def forward_as_attachment_to=(value)
                @forward_as_attachment_to = value
            end
            ## 
            ## Gets the forwardTo property value. The email addresses of the recipients to which a message should be forwarded.
            ## @return a recipient
            ## 
            def forward_to
                return @forward_to
            end
            ## 
            ## Sets the forwardTo property value. The email addresses of the recipients to which a message should be forwarded.
            ## @param value Value to set for the forwardTo property.
            ## @return a void
            ## 
            def forward_to=(value)
                @forward_to = value
            end
            ## 
            ## The deserialization information for the current model
            ## @return a i_dictionary
            ## 
            def get_field_deserializers()
                return {
                    "assignCategories" => lambda {|n| @assign_categories = n.get_collection_of_primitive_values(String) },
                    "copyToFolder" => lambda {|n| @copy_to_folder = n.get_string_value() },
                    "delete" => lambda {|n| @delete = n.get_boolean_value() },
                    "forwardAsAttachmentTo" => lambda {|n| @forward_as_attachment_to = n.get_collection_of_object_values(lambda {|pn| Graphrubyv4::Models::Recipient.create_from_discriminator_value(pn) }) },
                    "forwardTo" => lambda {|n| @forward_to = n.get_collection_of_object_values(lambda {|pn| Graphrubyv4::Models::Recipient.create_from_discriminator_value(pn) }) },
                    "markAsRead" => lambda {|n| @mark_as_read = n.get_boolean_value() },
                    "markImportance" => lambda {|n| @mark_importance = n.get_enum_value(Graphrubyv4::Models::Importance) },
                    "moveToFolder" => lambda {|n| @move_to_folder = n.get_string_value() },
                    "permanentDelete" => lambda {|n| @permanent_delete = n.get_boolean_value() },
                    "redirectTo" => lambda {|n| @redirect_to = n.get_collection_of_object_values(lambda {|pn| Graphrubyv4::Models::Recipient.create_from_discriminator_value(pn) }) },
                    "stopProcessingRules" => lambda {|n| @stop_processing_rules = n.get_boolean_value() },
                }
            end
            ## 
            ## Gets the markAsRead property value. Indicates whether a message should be marked as read.
            ## @return a boolean
            ## 
            def mark_as_read
                return @mark_as_read
            end
            ## 
            ## Sets the markAsRead property value. Indicates whether a message should be marked as read.
            ## @param value Value to set for the markAsRead property.
            ## @return a void
            ## 
            def mark_as_read=(value)
                @mark_as_read = value
            end
            ## 
            ## Gets the markImportance property value. The markImportance property
            ## @return a importance
            ## 
            def mark_importance
                return @mark_importance
            end
            ## 
            ## Sets the markImportance property value. The markImportance property
            ## @param value Value to set for the markImportance property.
            ## @return a void
            ## 
            def mark_importance=(value)
                @mark_importance = value
            end
            ## 
            ## Gets the moveToFolder property value. The ID of the folder that a message will be moved to.
            ## @return a string
            ## 
            def move_to_folder
                return @move_to_folder
            end
            ## 
            ## Sets the moveToFolder property value. The ID of the folder that a message will be moved to.
            ## @param value Value to set for the moveToFolder property.
            ## @return a void
            ## 
            def move_to_folder=(value)
                @move_to_folder = value
            end
            ## 
            ## Gets the permanentDelete property value. Indicates whether a message should be permanently deleted and not saved to the Deleted Items folder.
            ## @return a boolean
            ## 
            def permanent_delete
                return @permanent_delete
            end
            ## 
            ## Sets the permanentDelete property value. Indicates whether a message should be permanently deleted and not saved to the Deleted Items folder.
            ## @param value Value to set for the permanentDelete property.
            ## @return a void
            ## 
            def permanent_delete=(value)
                @permanent_delete = value
            end
            ## 
            ## Gets the redirectTo property value. The email addresses to which a message should be redirected.
            ## @return a recipient
            ## 
            def redirect_to
                return @redirect_to
            end
            ## 
            ## Sets the redirectTo property value. The email addresses to which a message should be redirected.
            ## @param value Value to set for the redirectTo property.
            ## @return a void
            ## 
            def redirect_to=(value)
                @redirect_to = value
            end
            ## 
            ## Serializes information the current object
            ## @param writer Serialization writer to use to serialize this model
            ## @return a void
            ## 
            def serialize(writer)
                raise StandardError, 'writer cannot be null' if writer.nil?
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
            ## Gets the stopProcessingRules property value. Indicates whether subsequent rules should be evaluated.
            ## @return a boolean
            ## 
            def stop_processing_rules
                return @stop_processing_rules
            end
            ## 
            ## Sets the stopProcessingRules property value. Indicates whether subsequent rules should be evaluated.
            ## @param value Value to set for the stopProcessingRules property.
            ## @return a void
            ## 
            def stop_processing_rules=(value)
                @stop_processing_rules = value
            end
        end
    end
end
