from __future__ import annotations
from kiota_abstractions.serialization import AdditionalDataHolder, Parsable, ParseNode, SerializationWriter
from typing import Any, Callable, Dict, List, Optional, Union

from . import body_type

class ItemBody(AdditionalDataHolder, Parsable):
    @property
    def additional_data(self,) -> Dict[str, Any]:
        """
        Gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        Returns: Dict[str, Any]
        """
        return self._additional_data

    @additional_data.setter
    def additional_data(self,value: Dict[str, Any]) -> None:
        """
        Sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        Args:
            value: Value to set for the AdditionalData property.
        """
        self._additional_data = value

    def __init__(self,) -> None:
        """
        Instantiates a new itemBody and sets the default values.
        """
        # The content of the item.
        self._content: Optional[str] = None

        # The contentType property
        self._content_type: Optional[body_type.BodyType] = None

        # Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        self._additional_data: Dict[str, Any] = {}


    @property
    def content(self,) -> Optional[str]:
        """
        Gets the content property value. The content of the item.
        Returns: Optional[str]
        """
        return self._content

    @content.setter
    def content(self,value: Optional[str] = None) -> None:
        """
        Sets the content property value. The content of the item.
        Args:
            value: Value to set for the content property.
        """
        self._content = value

    @property
    def content_type(self,) -> Optional[body_type.BodyType]:
        """
        Gets the contentType property value. The contentType property
        Returns: Optional[body_type.BodyType]
        """
        return self._content_type

    @content_type.setter
    def content_type(self,value: Optional[body_type.BodyType] = None) -> None:
        """
        Sets the contentType property value. The contentType property
        Args:
            value: Value to set for the contentType property.
        """
        self._content_type = value

    @staticmethod
    def create_from_discriminator_value(parse_node: Optional[ParseNode] = None) -> ItemBody:
        """
        Creates a new instance of the appropriate class based on discriminator value
        Args:
            parseNode: The parse node to use to read the discriminator value and create the object
        Returns: ItemBody
        """
        if not parse_node:
            raise Exception("parse_node cannot be undefined")
        return ItemBody()

    def get_field_deserializers(self,) -> Dict[str, Callable[[ParseNode], None]]:
        """
        The deserialization information for the current model
        Returns: Dict[str, Callable[[ParseNode], None]]
        """
        fields = {
            "content": lambda n : setattr(self, 'content', n.get_str_value()),
            "content_type": lambda n : setattr(self, 'content_type', n.get_enum_value(body_type.BodyType)),
        }
        return fields

    def serialize(self,writer: SerializationWriter) -> None:
        """
        Serializes information the current object
        Args:
            writer: Serialization writer to use to serialize this model
        """
        if not writer:
            raise Exception("writer cannot be undefined")
        writer.write_str_value("content", self.content)
        writer.write_enum_value("contentType", self.content_type)
        writer.write_additional_data_value(self.additional_data)


