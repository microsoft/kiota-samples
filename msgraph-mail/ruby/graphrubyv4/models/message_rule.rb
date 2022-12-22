require 'microsoft_kiota_abstractions'
require_relative './entity'
require_relative './models'

module Graphrubyv4::Models
    class MessageRule < Entity
        include MicrosoftKiotaAbstractions::Parsable
        ## 
        # The actions property
        @actions
        ## 
        # The conditions property
        @conditions
        ## 
        # The display name of the rule.
        @display_name
        ## 
        # The exceptions property
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
        ## Gets the actions property value. The actions property
        ## @return a message_rule_actions
        ## 
        def actions
            return @actions
        end
        ## 
        ## Sets the actions property value. The actions property
        ## @param value Value to set for the actions property.
        ## @return a void
        ## 
        def actions=(actions)
            @actions = actions
        end
        ## 
        ## Gets the conditions property value. The conditions property
        ## @return a message_rule_predicates
        ## 
        def conditions
            return @conditions
        end
        ## 
        ## Sets the conditions property value. The conditions property
        ## @param value Value to set for the conditions property.
        ## @return a void
        ## 
        def conditions=(conditions)
            @conditions = conditions
        end
        ## 
        ## Instantiates a new messageRule and sets the default values.
        ## @return a void
        ## 
        def initialize() 
            super
        end
        ## 
        ## Creates a new instance of the appropriate class based on discriminator value
        ## @param parseNode The parse node to use to read the discriminator value and create the object
        ## @return a message_rule
        ## 
        def create_from_discriminator_value(parse_node) 
            return nil;
        end
        ## 
        ## Gets the displayName property value. The display name of the rule.
        ## @return a string
        ## 
        def display_name
            return @display_name
        end
        ## 
        ## Sets the displayName property value. The display name of the rule.
        ## @param value Value to set for the displayName property.
        ## @return a void
        ## 
        def display_name=(displayName)
            @display_name = displayName
        end
        ## 
        ## Gets the exceptions property value. The exceptions property
        ## @return a message_rule_predicates
        ## 
        def exceptions
            return @exceptions
        end
        ## 
        ## Sets the exceptions property value. The exceptions property
        ## @param value Value to set for the exceptions property.
        ## @return a void
        ## 
        def exceptions=(exceptions)
            @exceptions = exceptions
        end
        ## 
        ## The deserialization information for the current model
        ## @return a i_dictionary
        ## 
        def get_field_deserializers() 
            return super.merge({
                "actions" => lambda {|n| @actions = n.get_object_value(Graphrubyv4::Models::MessageRuleActions) },
                "conditions" => lambda {|n| @conditions = n.get_object_value(Graphrubyv4::Models::MessageRulePredicates) },
                "displayName" => lambda {|n| @display_name = n.get_string_value() },
                "exceptions" => lambda {|n| @exceptions = n.get_object_value(Graphrubyv4::Models::MessageRulePredicates) },
                "hasError" => lambda {|n| @has_error = n.get_boolean_value() },
                "isEnabled" => lambda {|n| @is_enabled = n.get_boolean_value() },
                "isReadOnly" => lambda {|n| @is_read_only = n.get_boolean_value() },
                "sequence" => lambda {|n| @sequence = n.get_number_value() },
            })
        end
        ## 
        ## Gets the hasError property value. Indicates whether the rule is in an error condition. Read-only.
        ## @return a boolean
        ## 
        def has_error
            return @has_error
        end
        ## 
        ## Sets the hasError property value. Indicates whether the rule is in an error condition. Read-only.
        ## @param value Value to set for the hasError property.
        ## @return a void
        ## 
        def has_error=(hasError)
            @has_error = hasError
        end
        ## 
        ## Gets the isEnabled property value. Indicates whether the rule is enabled to be applied to messages.
        ## @return a boolean
        ## 
        def is_enabled
            return @is_enabled
        end
        ## 
        ## Sets the isEnabled property value. Indicates whether the rule is enabled to be applied to messages.
        ## @param value Value to set for the isEnabled property.
        ## @return a void
        ## 
        def is_enabled=(isEnabled)
            @is_enabled = isEnabled
        end
        ## 
        ## Gets the isReadOnly property value. Indicates if the rule is read-only and cannot be modified or deleted by the rules REST API.
        ## @return a boolean
        ## 
        def is_read_only
            return @is_read_only
        end
        ## 
        ## Sets the isReadOnly property value. Indicates if the rule is read-only and cannot be modified or deleted by the rules REST API.
        ## @param value Value to set for the isReadOnly property.
        ## @return a void
        ## 
        def is_read_only=(isReadOnly)
            @is_read_only = isReadOnly
        end
        ## 
        ## Gets the sequence property value. Indicates the order in which the rule is executed, among other rules.
        ## @return a integer
        ## 
        def sequence
            return @sequence
        end
        ## 
        ## Sets the sequence property value. Indicates the order in which the rule is executed, among other rules.
        ## @param value Value to set for the sequence property.
        ## @return a void
        ## 
        def sequence=(sequence)
            @sequence = sequence
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
    end
end
