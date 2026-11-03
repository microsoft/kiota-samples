# frozen_string_literal: true

require "microsoft_kiota_abstractions"

module KiotaPosts
  module Models
    class Post
      include MicrosoftKiotaAbstractions::Parsable
      include MicrosoftKiotaAbstractions::AdditionalDataHolder

      # Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
      attr_accessor :additional_data

      # The body property
      attr_accessor :body

      # The id property
      attr_accessor :id

      # The title property
      attr_accessor :title

      # The userId property
      attr_accessor :user_id

      # Instantiates a new Post and sets the default values.
      # @return a void
      def initialize
        @additional_data = {}
      end

      # Creates a new instance of the appropriate class based on discriminator value
      # @param parse_node The parse node to use to read the discriminator value and create the object
      # @return a post
      def self.create_from_discriminator_value(parse_node)
        raise StandardError, "parse_node cannot be null" if parse_node.nil?

        Post.new
      end

      # The deserialization information for the current model
      # @return a i_dictionary
      def get_field_deserializers
        {
          "body" => ->(n) { @body = n.get_string_value },
          "id" => ->(n) { @id = n.get_number_value },
          "title" => ->(n) { @title = n.get_string_value },
          "userId" => ->(n) { @user_id = n.get_number_value }
        }
      end

      # Serializes information the current object
      # @param writer Serialization writer to use to serialize this model
      # @return a void
      def serialize(writer)
        raise StandardError, "writer cannot be null" if writer.nil?

        writer.write_string_value("body", @body)
        writer.write_number_value("id", @id)
        writer.write_string_value("title", @title)
        writer.write_number_value("userId", @user_id)
        writer.write_additional_data(@additional_data)
      end
    end
  end
end
