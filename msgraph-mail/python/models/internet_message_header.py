from __future__ import annotations
from kiota_abstractions.serialization import AdditionalDataHolder, Parsable, ParseNode, SerializationWriter
from typing import Any, Callable, Dict, List, Optional, Union

class InternetMessageHeader(AdditionalDataHolder, Parsable):
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
        Instantiates a new internetMessageHeader and sets the default values.
        """
        # Represents the key in a key-value pair.
        self._name: Optional[str] = None

        # The value in a key-value pair.
        self._value: Optional[str] = None

        # Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        self._additional_data: Dict[str, Any] = {}


    @staticmethod
    def create_from_discriminator_value(parse_node: Optional[ParseNode] = None) -> InternetMessageHeader:
        """
        Creates a new instance of the appropriate class based on discriminator value
        Args:
            parseNode: The parse node to use to read the discriminator value and create the object
        Returns: InternetMessageHeader
        """
        if not parse_node:
            raise Exception("parse_node cannot be undefined")
        return InternetMessageHeader()

    def get_field_deserializers(self,) -> Dict[str, Callable[[ParseNode], None]]:
        """
        The deserialization information for the current model
        Returns: Dict[str, Callable[[ParseNode], None]]
        """
        fields = {
            "name": lambda n : setattr(self, 'name', n.get_str_value()),
            "value": lambda n : setattr(self, 'value', n.get_str_value()),
        }
        return fields

    @property
    def name(self,) -> Optional[str]:
        """
        Gets the name property value. Represents the key in a key-value pair.
        Returns: Optional[str]
        """
        return self._name

    @name.setter
    def name(self,value: Optional[str] = None) -> None:
        """
        Sets the name property value. Represents the key in a key-value pair.
        Args:
            value: Value to set for the name property.
        """
        self._name = value

    def serialize(self,writer: SerializationWriter) -> None:
        """
        Serializes information the current object
        Args:
            writer: Serialization writer to use to serialize this model
        """
        if not writer:
            raise Exception("writer cannot be undefined")
        writer.write_str_value("name", self.name)
        writer.write_str_value("value", self.value)
        writer.write_additional_data_value(self.additional_data)

    @property
    def value(self,) -> Optional[str]:
        """
        Gets the value property value. The value in a key-value pair.
        Returns: Optional[str]
        """
        return self._value

    @value.setter
    def value(self,value: Optional[str] = None) -> None:
        """
        Sets the value property value. The value in a key-value pair.
        Args:
            value: Value to set for the value property.
        """
        self._value = value


