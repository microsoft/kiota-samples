from __future__ import annotations
from dataclasses import dataclass, field
from kiota_abstractions.base_request_builder import BaseRequestBuilder
from kiota_abstractions.base_request_configuration import RequestConfiguration
from kiota_abstractions.get_path_parameters import get_path_parameters
from kiota_abstractions.method import Method
from kiota_abstractions.request_adapter import RequestAdapter
from kiota_abstractions.request_information import RequestInformation
from kiota_abstractions.request_option import RequestOption
from kiota_abstractions.serialization import Parsable, ParsableFactory
from typing import Any, Callable, Dict, List, Optional, TYPE_CHECKING, Union

if TYPE_CHECKING:
    from .....models.message import Message
    from .....models.o_data_errors.o_data_error import ODataError
    from .attachments.attachments_request_builder import AttachmentsRequestBuilder
    from .extensions.extensions_request_builder import ExtensionsRequestBuilder
    from .value.content_request_builder import ContentRequestBuilder

class MessageItemRequestBuilder(BaseRequestBuilder):
    """
    Builds and executes requests for operations under /users/{user-id}/messages/{message-id}
    """
    def __init__(self,request_adapter: RequestAdapter, path_parameters: Union[str, Dict[str, Any]]) -> None:
        """
        Instantiates a new MessageItemRequestBuilder and sets the default values.
        param path_parameters: The raw url or the url-template parameters for the request.
        param request_adapter: The request adapter to use to execute the requests.
        Returns: None
        """
        super().__init__(request_adapter, "{+baseurl}/users/{user%2Did}/messages/{message%2Did}{?%24expand,%24select,includeHiddenMessages}", path_parameters)
    
    async def delete(self,request_configuration: Optional[RequestConfiguration] = None) -> bytes:
        """
        Delete a message in the specified user's mailbox, or delete a relationship of the message.
        param request_configuration: Configuration for the request such as headers, query parameters, and middleware options.
        Returns: bytes
        Find more info here: https://learn.microsoft.com/graph/api/message-delete?view=graph-rest-1.0
        """
        request_info = self.to_delete_request_information(
            request_configuration
        )
        from .....models.o_data_errors.o_data_error import ODataError

        error_mapping: Dict[str, ParsableFactory] = {
            "4XX": ODataError,
            "5XX": ODataError,
        }
        if not self.request_adapter:
            raise Exception("Http core is null") 
        return await self.request_adapter.send_primitive_async(request_info, "bytes", error_mapping)
    
    async def get(self,request_configuration: Optional[RequestConfiguration] = None) -> Optional[Message]:
        """
        The messages in a mailbox or folder. Read-only. Nullable.
        param request_configuration: Configuration for the request such as headers, query parameters, and middleware options.
        Returns: Optional[Message]
        Find more info here: https://learn.microsoft.com/graph/api/eventmessage-get?view=graph-rest-1.0
        """
        request_info = self.to_get_request_information(
            request_configuration
        )
        from .....models.o_data_errors.o_data_error import ODataError

        error_mapping: Dict[str, ParsableFactory] = {
            "4XX": ODataError,
            "5XX": ODataError,
        }
        if not self.request_adapter:
            raise Exception("Http core is null") 
        from .....models.message import Message

        return await self.request_adapter.send_async(request_info, Message, error_mapping)
    
    async def patch(self,body: Optional[Message] = None, request_configuration: Optional[RequestConfiguration] = None) -> Optional[Message]:
        """
        Update the properties of an eventMessage object.
        param body: The request body
        param request_configuration: Configuration for the request such as headers, query parameters, and middleware options.
        Returns: Optional[Message]
        Find more info here: https://learn.microsoft.com/graph/api/eventmessage-update?view=graph-rest-1.0
        """
        if not body:
            raise TypeError("body cannot be null.")
        request_info = self.to_patch_request_information(
            body, request_configuration
        )
        from .....models.o_data_errors.o_data_error import ODataError

        error_mapping: Dict[str, ParsableFactory] = {
            "4XX": ODataError,
            "5XX": ODataError,
        }
        if not self.request_adapter:
            raise Exception("Http core is null") 
        from .....models.message import Message

        return await self.request_adapter.send_async(request_info, Message, error_mapping)
    
    def to_delete_request_information(self,request_configuration: Optional[RequestConfiguration] = None) -> RequestInformation:
        """
        Delete a message in the specified user's mailbox, or delete a relationship of the message.
        param request_configuration: Configuration for the request such as headers, query parameters, and middleware options.
        Returns: RequestInformation
        """
        request_info = RequestInformation(Method.DELETE, '{+baseurl}/users/{user%2Did}/messages/{message%2Did}', self.path_parameters)
        request_info.configure(request_configuration)
        request_info.headers.try_add("Accept", "application/json")
        return request_info
    
    def to_get_request_information(self,request_configuration: Optional[RequestConfiguration] = None) -> RequestInformation:
        """
        The messages in a mailbox or folder. Read-only. Nullable.
        param request_configuration: Configuration for the request such as headers, query parameters, and middleware options.
        Returns: RequestInformation
        """
        request_info = RequestInformation(Method.GET, self.url_template, self.path_parameters)
        request_info.configure(request_configuration)
        request_info.headers.try_add("Accept", "application/json")
        return request_info
    
    def to_patch_request_information(self,body: Optional[Message] = None, request_configuration: Optional[RequestConfiguration] = None) -> RequestInformation:
        """
        Update the properties of an eventMessage object.
        param body: The request body
        param request_configuration: Configuration for the request such as headers, query parameters, and middleware options.
        Returns: RequestInformation
        """
        if not body:
            raise TypeError("body cannot be null.")
        request_info = RequestInformation(Method.PATCH, '{+baseurl}/users/{user%2Did}/messages/{message%2Did}', self.path_parameters)
        request_info.configure(request_configuration)
        request_info.headers.try_add("Accept", "application/json")
        request_info.set_content_from_parsable(self.request_adapter, "application/json", body)
        return request_info
    
    def with_url(self,raw_url: Optional[str] = None) -> MessageItemRequestBuilder:
        """
        Returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
        param raw_url: The raw URL to use for the request builder.
        Returns: MessageItemRequestBuilder
        """
        if not raw_url:
            raise TypeError("raw_url cannot be null.")
        return MessageItemRequestBuilder(self.request_adapter, raw_url)
    
    @property
    def attachments(self) -> AttachmentsRequestBuilder:
        """
        The attachments property
        """
        from .attachments.attachments_request_builder import AttachmentsRequestBuilder

        return AttachmentsRequestBuilder(self.request_adapter, self.path_parameters)
    
    @property
    def content(self) -> ContentRequestBuilder:
        """
        The Content property
        """
        from .value.content_request_builder import ContentRequestBuilder

        return ContentRequestBuilder(self.request_adapter, self.path_parameters)
    
    @property
    def extensions(self) -> ExtensionsRequestBuilder:
        """
        The extensions property
        """
        from .extensions.extensions_request_builder import ExtensionsRequestBuilder

        return ExtensionsRequestBuilder(self.request_adapter, self.path_parameters)
    
    @dataclass
    class MessageItemRequestBuilderGetQueryParameters():
        """
        The messages in a mailbox or folder. Read-only. Nullable.
        """
        def get_query_parameter(self,original_name: Optional[str] = None) -> str:
            """
            Maps the query parameters names to their encoded names for the URI template parsing.
            param original_name: The original query parameter name in the class.
            Returns: str
            """
            if not original_name:
                raise TypeError("original_name cannot be null.")
            if original_name == "expand":
                return "%24expand"
            if original_name == "include_hidden_messages":
                return "includeHiddenMessages"
            if original_name == "select":
                return "%24select"
            return original_name
        
        # Expand related entities
        expand: Optional[List[str]] = None

        # Include Hidden Messages
        include_hidden_messages: Optional[str] = None

        # Select properties to be returned
        select: Optional[List[str]] = None

    

