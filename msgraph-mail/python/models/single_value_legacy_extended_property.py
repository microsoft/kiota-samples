from __future__ import annotations
from kiota_abstractions.serialization import Parsable, ParseNode, SerializationWriter
from typing import Any, Callable, Dict, List, Optional, Union

from . import entity

class SingleValueLegacyExtendedProperty(entity.Entity):
    def __init__(self,) -> None:
        """
        Instantiates a new SingleValueLegacyExtendedProperty and sets the default values.
        """
        super().__init__()
        # A property value.
        self._value: Optional[str] = None


    @staticmethod
    def create_from_discriminator_value(parse_node: Optional[ParseNode] = None) -> SingleValueLegacyExtendedProperty:
        """
        Creates a new instance of the appropriate class based on discriminator value
        Args:
            parseNode: The parse node to use to read the discriminator value and create the object
        Returns: SingleValueLegacyExtendedProperty
        """
        if not parse_node:
            raise Exception("parse_node cannot be undefined")
        return SingleValueLegacyExtendedProperty()

    def get_field_deserializers(self,) -> Dict[str, Callable[[ParseNode], None]]:
        """
        The deserialization information for the current model
        Returns: Dict[str, Callable[[ParseNode], None]]
        """
        fields = {
            "value": lambda n : setattr(self, 'value', n.get_str_value()),
        }
        super_fields = super().get_field_deserializers()
        fields.update(super_fields)
        return fields

    def serialize(self,writer: SerializationWriter) -> None:
        """
        Serializes information the current object
        Args:
            writer: Serialization writer to use to serialize this model
        """
        if not writer:
            raise Exception("writer cannot be undefined")
        super().serialize(writer)
        writer.write_str_value("value", self.value)

    @property
    def value(self,) -> Optional[str]:
        """
        Gets the value property value. A property value.
        Returns: Optional[str]
        """
        return self._value

    @value.setter
    def value(self,value: Optional[str] = None) -> None:
        """
        Sets the value property value. A property value.
        Args:
            value: Value to set for the value property.
        """
        self._value = value


