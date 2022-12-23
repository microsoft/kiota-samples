require 'microsoft_kiota_abstractions'
require_relative '../models'
require_relative './o_data_errors'

module Graphrubyv4::Models::ODataErrors
    class ODataError < MicrosoftKiotaAbstractions::ApiError
        include MicrosoftKiotaAbstractions::AdditionalDataHolder, MicrosoftKiotaAbstractions::Parsable
        ## 
        # Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        @additional_data
        ## 
        # The error property
        @error
        ## 
        ## Gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        ## @return a i_dictionary
        ## 
        def additional_data
            return @additional_data
        end
        ## 
        ## Sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        ## @param value Value to set for the AdditionalData property.
        ## @return a void
        ## 
        def additional_data=(value)
            @additional_data = value
        end
        ## 
        ## Instantiates a new ODataError and sets the default values.
        ## @return a void
        ## 
        def initialize()
            super
            @additional_data = Hash.new
        end
        ## 
        ## Creates a new instance of the appropriate class based on discriminator value
        ## @param parseNode The parse node to use to read the discriminator value and create the object
        ## @return a o_data_error
        ## 
        def self.create_from_discriminator_value(parse_node)
            raise StandardError, 'parse_node cannot be null' if parse_node.nil?
            return ODataError.new
        end
        ## 
        ## Gets the error property value. The error property
        ## @return a main_error
        ## 
        def error
            return @error
        end
        ## 
        ## Sets the error property value. The error property
        ## @param value Value to set for the error property.
        ## @return a void
        ## 
        def error=(value)
            @error = value
        end
        ## 
        ## The deserialization information for the current model
        ## @return a i_dictionary
        ## 
        def get_field_deserializers()
            return super.merge({
                "error" => lambda {|n| @error = n.get_object_value(lambda {|pn| Graphrubyv4::Models::ODataErrors::MainError.create_from_discriminator_value(pn) }) },
            })
        end
        ## 
        ## Serializes information the current object
        ## @param writer Serialization writer to use to serialize this model
        ## @return a void
        ## 
        def serialize(writer)
            raise StandardError, 'writer cannot be null' if writer.nil?
            super
            writer.write_object_value("error", @error)
            writer.write_additional_data(@additional_data)
        end
    end
end
