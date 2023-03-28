from __future__ import annotations
from kiota_abstractions.serialization import AdditionalDataHolder, Parsable, ParseNode, SerializationWriter
from kiota_abstractions.utils import lazy_import
from typing import Any, Callable, Dict, List, Optional, Union

class Post(AdditionalDataHolder, Parsable):
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
    
    @property
    def body(self,) -> Optional[str]:
        """
        Gets the body property value. The body property
        Returns: Optional[str]
        """
        return self._body
    
    @body.setter
    def body(self,value: Optional[str] = None) -> None:
        """
        Sets the body property value. The body property
        Args:
            value: Value to set for the body property.
        """
        self._body = value
    
    def __init__(self,) -> None:
        """
        Instantiates a new Post and sets the default values.
        """
        # Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
        self._additional_data: Dict[str, Any] = {}

        # The body property
        self._body: Optional[str] = None
        # The id property
        self._id: Optional[int] = None
        # The title property
        self._title: Optional[str] = None
        # The userId property
        self._user_id: Optional[int] = None
    
    @staticmethod
    def create_from_discriminator_value(parse_node: Optional[ParseNode] = None) -> Post:
        """
        Creates a new instance of the appropriate class based on discriminator value
        Args:
            parseNode: The parse node to use to read the discriminator value and create the object
        Returns: Post
        """
        if parse_node is None:
            raise Exception("parse_node cannot be undefined")
        return Post()
    
    def get_field_deserializers(self,) -> Dict[str, Callable[[ParseNode], None]]:
        """
        The deserialization information for the current model
        Returns: Dict[str, Callable[[ParseNode], None]]
        """
        fields = {
            "body": lambda n : setattr(self, 'body', n.get_str_value()),
            "id": lambda n : setattr(self, 'id', n.get_int_value()),
            "title": lambda n : setattr(self, 'title', n.get_str_value()),
            "userId": lambda n : setattr(self, 'user_id', n.get_int_value()),
        }
        return fields
    
    @property
    def id(self,) -> Optional[int]:
        """
        Gets the id property value. The id property
        Returns: Optional[int]
        """
        return self._id
    
    @id.setter
    def id(self,value: Optional[int] = None) -> None:
        """
        Sets the id property value. The id property
        Args:
            value: Value to set for the id property.
        """
        self._id = value
    
    def serialize(self,writer: SerializationWriter) -> None:
        """
        Serializes information the current object
        Args:
            writer: Serialization writer to use to serialize this model
        """
        if writer is None:
            raise Exception("writer cannot be undefined")
        writer.write_str_value("body", self.body)
        writer.write_int_value("id", self.id)
        writer.write_str_value("title", self.title)
        writer.write_int_value("userId", self.user_id)
        writer.write_additional_data_value(self.additional_data)
    
    @property
    def title(self,) -> Optional[str]:
        """
        Gets the title property value. The title property
        Returns: Optional[str]
        """
        return self._title
    
    @title.setter
    def title(self,value: Optional[str] = None) -> None:
        """
        Sets the title property value. The title property
        Args:
            value: Value to set for the title property.
        """
        self._title = value
    
    @property
    def user_id(self,) -> Optional[int]:
        """
        Gets the userId property value. The userId property
        Returns: Optional[int]
        """
        return self._user_id
    
    @user_id.setter
    def user_id(self,value: Optional[int] = None) -> None:
        """
        Sets the userId property value. The userId property
        Args:
            value: Value to set for the user_id property.
        """
        self._user_id = value
    

