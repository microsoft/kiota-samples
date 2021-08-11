require 'microsoft_kiota_abstractions'
require_relative './entity'
require_relative './mail_folders/message_rule_predicates'
require_relative './users'

module Graphrubyv4::Users
    class MessageRule < Graphrubyv4::Users::Entity
        include MicrosoftKiotaAbstractions::Parsable
        @actions
        @conditions
        ## 
        # The display name of the rule.
        @display_name
        @exceptions
        ## 
        # Indicates whether the rule is in an error condition. Read-only.
        @has_error
        ## 
        # Indicates whether the rule is enabled to be applied to messages.
        @is_enabled
        ## 
        # Indicates if the rule is read-only and cannot be modified or deleted by the rules REST API.
        @is_read_only
        ## 
        # Indicates the order in which the rule is executed, among other rules.
        @sequence
        ## 
        ## Gets the actions property value. 
        ## @return a message_rule_actions
        ## 
        def  actions
            return @actions
        end
        ## 
        ## Gets the conditions property value. 
        ## @return a message_rule_predicates
        ## 
        def  conditions
            return @conditions
        end
        ## 
        ## Gets the displayName property value. The display name of the rule.
        ## @return a string
        ## 
        def  display_name
            return @display_name
        end
        ## 
        ## Gets the exceptions property value. 
        ## @return a message_rule_predicates
        ## 
        def  exceptions
            return @exceptions
        end
        ## 
        ## Gets the hasError property value. Indicates whether the rule is in an error condition. Read-only.
        ## @return a boolean
        ## 
        def  has_error
            return @has_error
        end
        ## 
        ## Gets the isEnabled property value. Indicates whether the rule is enabled to be applied to messages.
        ## @return a boolean
        ## 
        def  is_enabled
            return @is_enabled
        end
        ## 
        ## Gets the isReadOnly property value. Indicates if the rule is read-only and cannot be modified or deleted by the rules REST API.
        ## @return a boolean
        ## 
        def  is_read_only
            return @is_read_only
        end
        ## 
        ## Gets the sequence property value. Indicates the order in which the rule is executed, among other rules.
        ## @return a integer
        ## 
        def  sequence
            return @sequence
        end
        ## 
        ## The deserialization information for the current model
        ## @return a i_dictionary
        ## 
        def get_field_deserializers() 
            return super.merge({
                "actions" => lambda {|o, n| o.actions = n.get_object_value(Graphrubyv4::Users::MailFolders::MessageRuleActions) },
                "conditions" => lambda {|o, n| o.conditions = n.get_object_value(Graphrubyv4::Users::MailFolders::MessageRulePredicates) },
                "displayName" => lambda {|o, n| o.display_name = n.get_string_value() },
                "exceptions" => lambda {|o, n| o.exceptions = n.get_object_value(Graphrubyv4::Users::MailFolders::MessageRulePredicates) },
                "hasError" => lambda {|o, n| o.has_error = n.get_boolean_value() },
                "isEnabled" => lambda {|o, n| o.is_enabled = n.get_boolean_value() },
                "isReadOnly" => lambda {|o, n| o.is_read_only = n.get_boolean_value() },
                "sequence" => lambda {|o, n| o.sequence = n.get_number_value() },
            })
        end
        ## 
        ## Serializes information the current object
        ## @param writer Serialization writer to use to serialize this model
        ## @return a void
        ## 
        def serialize(writer) 
            super
            writer.write_object_value("actions", @actions)
            writer.write_object_value("conditions", @conditions)
            writer.write_string_value("displayName", @display_name)
            writer.write_object_value("exceptions", @exceptions)
            writer.write_boolean_value("hasError", @has_error)
            writer.write_boolean_value("isEnabled", @is_enabled)
            writer.write_boolean_value("isReadOnly", @is_read_only)
            writer.write_number_value("sequence", @sequence)
        end
        ## 
        ## Sets the actions property value. 
        ## @param value Value to set for the actions property.
        ## @return a void
        ## 
        def  actions=(actions)
            @actions = actions
        end
        ## 
        ## Sets the conditions property value. 
        ## @param value Value to set for the conditions property.
        ## @return a void
        ## 
        def  conditions=(conditions)
            @conditions = conditions
        end
        ## 
        ## Sets the displayName property value. The display name of the rule.
        ## @param value Value to set for the displayName property.
        ## @return a void
        ## 
        def  display_name=(displayName)
            @display_name = displayName
        end
        ## 
        ## Sets the exceptions property value. 
        ## @param value Value to set for the exceptions property.
        ## @return a void
        ## 
        def  exceptions=(exceptions)
            @exceptions = exceptions
        end
        ## 
        ## Sets the hasError property value. Indicates whether the rule is in an error condition. Read-only.
        ## @param value Value to set for the hasError property.
        ## @return a void
        ## 
        def  has_error=(hasError)
            @has_error = hasError
        end
        ## 
        ## Sets the isEnabled property value. Indicates whether the rule is enabled to be applied to messages.
        ## @param value Value to set for the isEnabled property.
        ## @return a void
        ## 
        def  is_enabled=(isEnabled)
            @is_enabled = isEnabled
        end
        ## 
        ## Sets the isReadOnly property value. Indicates if the rule is read-only and cannot be modified or deleted by the rules REST API.
        ## @param value Value to set for the isReadOnly property.
        ## @return a void
        ## 
        def  is_read_only=(isReadOnly)
            @is_read_only = isReadOnly
        end
        ## 
        ## Sets the sequence property value. Indicates the order in which the rule is executed, among other rules.
        ## @param value Value to set for the sequence property.
        ## @return a void
        ## 
        def  sequence=(sequence)
            @sequence = sequence
        end
    end
end
