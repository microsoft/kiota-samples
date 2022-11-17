from __future__ import annotations
from kiota_abstractions.serialization import Parsable, ParseNode, SerializationWriter
from typing import Any, Callable, Dict, List, Optional, Union

from . import entity

class MultiValueLegacyExtendedProperty(entity.Entity):
    def __init__(self,) -> None:
        """
        Instantiates a new MultiValueLegacyExtendedProperty and sets the default values.
        """
        super().__init__()
        # A collection of property values.
        self._value: Optional[List[str]] = None


    @staticmethod
    def create_from_discriminator_value(parse_node: Optional[ParseNode] = None) -> MultiValueLegacyExtendedProperty:
        """
        Creates a new instance of the appropriate class based on discriminator value
        Args:
            parseNode: The parse node to use to read the discriminator value and create the object
        Returns: MultiValueLegacyExtendedProperty
        """
        if not parse_node:
            raise Exception("parse_node cannot be undefined")
        return MultiValueLegacyExtendedProperty()

    def get_field_deserializers(self,) -> Dict[str, Callable[[ParseNode], None]]:
        """
        The deserialization information for the current model
        Returns: Dict[str, Callable[[ParseNode], None]]
        """
        fields = {
            "value": lambda n : setattr(self, 'value', n.get_collection_of_primitive_values(str)),
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
        writer.write_collection_of_primitive_values("value", self.value)

    @property
    def value(self,) -> Optional[List[str]]:
        """
        Gets the value property value. A collection of property values.
        Returns: Optional[List[str]]
        """
        return self._value

    @value.setter
    def value(self,value: Optional[List[str]] = None) -> None:
        """
        Sets the value property value. A collection of property values.
        Args:
            value: Value to set for the value property.
        """
        self._value = value


