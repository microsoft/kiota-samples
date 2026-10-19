# frozen_string_literal: true

require "microsoft_kiota_abstractions"

module GetUser
  module Models
    class User
      include MicrosoftKiotaAbstractions::Parsable
      include MicrosoftKiotaAbstractions::AdditionalDataHolder

      # Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
      attr_accessor :additional_data

      # The displayName property
      attr_accessor :display_name

      # The id property
      attr_accessor :id

      # Instantiates a new User and sets the default values.
      # @return a void
      def initialize
        @additional_data = {}
      end

      # Creates a new instance of the appropriate class based on discriminator value
      # @param parse_node The parse node to use to read the discriminator value and create the object
      # @return a user
      def self.create_from_discriminator_value(parse_node)
        raise StandardError, "parse_node cannot be null" if parse_node.nil?

        User.new
      end

      # The deserialization information for the current model
      # @return a i_dictionary
      def get_field_deserializers
        {
          "displayName" => ->(n) { @display_name = n.get_string_value },
          "id" => ->(n) { @id = n.get_string_value }
        }
      end

      # Serializes information the current object
      # @param writer Serialization writer to use to serialize this model
      # @return a void
      def serialize(writer)
        raise StandardError, "writer cannot be null" if writer.nil?

        writer.write_string_value("displayName", @display_name)
        writer.write_string_value("id", @id)
        writer.write_additional_data(@additional_data)
      end
    end
  end
end
