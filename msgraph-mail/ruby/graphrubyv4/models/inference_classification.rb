require 'microsoft_kiota_abstractions'
require_relative '../graphrubyv4'
require_relative './models'

module Graphrubyv4
    module Models
        class InferenceClassification < Graphrubyv4::Models::Entity
            include MicrosoftKiotaAbstractions::Parsable
            ## 
            # A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable.
            @overrides
            ## 
            ## Instantiates a new InferenceClassification and sets the default values.
            ## @return a void
            ## 
            def initialize()
                super
            end
            ## 
            ## Creates a new instance of the appropriate class based on discriminator value
            ## @param parse_node The parse node to use to read the discriminator value and create the object
            ## @return a inference_classification
            ## 
            def self.create_from_discriminator_value(parse_node)
                raise StandardError, 'parse_node cannot be null' if parse_node.nil?
                return InferenceClassification.new
            end
            ## 
            ## The deserialization information for the current model
            ## @return a i_dictionary
            ## 
            def get_field_deserializers()
                return super.merge({
                    "overrides" => lambda {|n| @overrides = n.get_collection_of_object_values(lambda {|pn| Graphrubyv4::Models::InferenceClassificationOverride.create_from_discriminator_value(pn) }) },
                })
            end
            ## 
            ## Gets the overrides property value. A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable.
            ## @return a inference_classification_override
            ## 
            def overrides
                return @overrides
            end
            ## 
            ## Sets the overrides property value. A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable.
            ## @param value Value to set for the overrides property.
            ## @return a void
            ## 
            def overrides=(value)
                @overrides = value
            end
            ## 
            ## Serializes information the current object
            ## @param writer Serialization writer to use to serialize this model
            ## @return a void
            ## 
            def serialize(writer)
                raise StandardError, 'writer cannot be null' if writer.nil?
                super
                writer.write_collection_of_object_values("overrides", @overrides)
            end
        end
    end
end
