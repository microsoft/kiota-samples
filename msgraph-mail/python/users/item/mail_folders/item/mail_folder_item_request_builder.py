from __future__ import annotations
from dataclasses import dataclass, field
from kiota_abstractions.get_path_parameters import get_path_parameters
from kiota_abstractions.method import Method
from kiota_abstractions.request_adapter import RequestAdapter
from kiota_abstractions.request_information import RequestInformation
from kiota_abstractions.request_option import RequestOption
from kiota_abstractions.response_handler import ResponseHandler
from kiota_abstractions.serialization import Parsable, ParsableFactory
from typing import Any, Callable, Dict, List, Optional, TYPE_CHECKING, Union

if TYPE_CHECKING:
    from .....models.mail_folder import MailFolder
    from .child_folders.child_folders_request_builder import ChildFoldersRequestBuilder
    from .message_rules.message_rules_request_builder import MessageRulesRequestBuilder
    from .messages.messages_request_builder import MessagesRequestBuilder
    from .multi_value_extended_properties.multi_value_extended_properties_request_builder import MultiValueExtendedPropertiesRequestBuilder
    from .single_value_extended_properties.single_value_extended_properties_request_builder import SingleValueExtendedPropertiesRequestBuilder

class MailFolderItemRequestBuilder():
    """
    Builds and executes requests for operations under /users/{user-id}/mailFolders/{mailFolder-id}
    """
    def __init__(self,request_adapter: RequestAdapter, path_parameters: Optional[Union[Dict[str, Any], str]] = None) -> None:
        """
        Instantiates a new MailFolderItemRequestBuilder and sets the default values.
        Args:
            pathParameters: The raw url or the Url template parameters for the request.
            requestAdapter: The request adapter to use to execute the requests.
        """
        if not path_parameters:
            raise TypeError("path_parameters cannot be null.")
        if not request_adapter:
            raise TypeError("request_adapter cannot be null.")
        # Url template to use to build the URL for the current request builder
        self.url_template: str = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}{?includeHiddenFolders,%24select,%24expand}"

        url_tpl_params = get_path_parameters(path_parameters)
        self.path_parameters = url_tpl_params
        self.request_adapter = request_adapter
    
    async def delete(self,request_configuration: Optional[MailFolderItemRequestBuilderDeleteRequestConfiguration] = None) -> None:
        """
        Delete navigation property mailFolders for users
        Args:
            requestConfiguration: Configuration for the request such as headers, query parameters, and middleware options.
        """
        request_info = self.to_delete_request_information(
            request_configuration
        )
        if not self.request_adapter:
            raise Exception("Http core is null") 
        return await self.request_adapter.send_no_response_content_async(request_info, None)
    
    async def get(self,request_configuration: Optional[MailFolderItemRequestBuilderGetRequestConfiguration] = None) -> Optional[MailFolder]:
        """
        The user's mail folders. Read-only. Nullable.
        Args:
            requestConfiguration: Configuration for the request such as headers, query parameters, and middleware options.
        Returns: Optional[MailFolder]
        """
        request_info = self.to_get_request_information(
            request_configuration
        )
        if not self.request_adapter:
            raise Exception("Http core is null") 
        from .....models.mail_folder import MailFolder

        return await self.request_adapter.send_async(request_info, MailFolder, None)
    
    async def patch(self,body: Optional[MailFolder] = None, request_configuration: Optional[MailFolderItemRequestBuilderPatchRequestConfiguration] = None) -> None:
        """
        Update the navigation property mailFolders in users
        Args:
            body: The request body
            requestConfiguration: Configuration for the request such as headers, query parameters, and middleware options.
        """
        if not body:
            raise TypeError("body cannot be null.")
        request_info = self.to_patch_request_information(
            body, request_configuration
        )
        if not self.request_adapter:
            raise Exception("Http core is null") 
        return await self.request_adapter.send_no_response_content_async(request_info, None)
    
    def to_delete_request_information(self,request_configuration: Optional[MailFolderItemRequestBuilderDeleteRequestConfiguration] = None) -> RequestInformation:
        """
        Delete navigation property mailFolders for users
        Args:
            requestConfiguration: Configuration for the request such as headers, query parameters, and middleware options.
        Returns: RequestInformation
        """
        request_info = RequestInformation()
        request_info.url_template = self.url_template
        request_info.path_parameters = self.path_parameters
        request_info.http_method = Method.DELETE
        if request_configuration:
            request_info.add_request_headers(request_configuration.headers)
            request_info.add_request_options(request_configuration.options)
        return request_info
    
    def to_get_request_information(self,request_configuration: Optional[MailFolderItemRequestBuilderGetRequestConfiguration] = None) -> RequestInformation:
        """
        The user's mail folders. Read-only. Nullable.
        Args:
            requestConfiguration: Configuration for the request such as headers, query parameters, and middleware options.
        Returns: RequestInformation
        """
        request_info = RequestInformation()
        request_info.url_template = self.url_template
        request_info.path_parameters = self.path_parameters
        request_info.http_method = Method.GET
        request_info.headers["Accept"] = ["application/json"]
        if request_configuration:
            request_info.add_request_headers(request_configuration.headers)
            request_info.set_query_string_parameters_from_raw_object(request_configuration.query_parameters)
            request_info.add_request_options(request_configuration.options)
        return request_info
    
    def to_patch_request_information(self,body: Optional[MailFolder] = None, request_configuration: Optional[MailFolderItemRequestBuilderPatchRequestConfiguration] = None) -> RequestInformation:
        """
        Update the navigation property mailFolders in users
        Args:
            body: The request body
            requestConfiguration: Configuration for the request such as headers, query parameters, and middleware options.
        Returns: RequestInformation
        """
        if not body:
            raise TypeError("body cannot be null.")
        request_info = RequestInformation()
        request_info.url_template = self.url_template
        request_info.path_parameters = self.path_parameters
        request_info.http_method = Method.PATCH
        if request_configuration:
            request_info.add_request_headers(request_configuration.headers)
            request_info.add_request_options(request_configuration.options)
        request_info.set_content_from_parsable(self.request_adapter, "application/json", body)
        return request_info
    
    @property
    def child_folders(self) -> ChildFoldersRequestBuilder:
        """
        The childFolders property
        """
        from .child_folders.child_folders_request_builder import ChildFoldersRequestBuilder

        return ChildFoldersRequestBuilder(self.request_adapter, self.path_parameters)
    
    @property
    def message_rules(self) -> MessageRulesRequestBuilder:
        """
        The messageRules property
        """
        from .message_rules.message_rules_request_builder import MessageRulesRequestBuilder

        return MessageRulesRequestBuilder(self.request_adapter, self.path_parameters)
    
    @property
    def messages(self) -> MessagesRequestBuilder:
        """
        The messages property
        """
        from .messages.messages_request_builder import MessagesRequestBuilder

        return MessagesRequestBuilder(self.request_adapter, self.path_parameters)
    
    @property
    def multi_value_extended_properties(self) -> MultiValueExtendedPropertiesRequestBuilder:
        """
        The multiValueExtendedProperties property
        """
        from .multi_value_extended_properties.multi_value_extended_properties_request_builder import MultiValueExtendedPropertiesRequestBuilder

        return MultiValueExtendedPropertiesRequestBuilder(self.request_adapter, self.path_parameters)
    
    @property
    def single_value_extended_properties(self) -> SingleValueExtendedPropertiesRequestBuilder:
        """
        The singleValueExtendedProperties property
        """
        from .single_value_extended_properties.single_value_extended_properties_request_builder import SingleValueExtendedPropertiesRequestBuilder

        return SingleValueExtendedPropertiesRequestBuilder(self.request_adapter, self.path_parameters)
    
    @dataclass
    class MailFolderItemRequestBuilderDeleteRequestConfiguration():
        """
        Configuration for the request such as headers, query parameters, and middleware options.
        """
        # Request headers
        headers: Optional[Dict[str, Union[str, List[str]]]] = None

        # Request options
        options: Optional[List[RequestOption]] = None

    
    @dataclass
    class MailFolderItemRequestBuilderGetQueryParameters():
        """
        The user's mail folders. Read-only. Nullable.
        """
        def get_query_parameter(self,original_name: Optional[str] = None) -> str:
            """
            Maps the query parameters names to their encoded names for the URI template parsing.
            Args:
                originalName: The original query parameter name in the class.
            Returns: str
            """
            if not original_name:
                raise TypeError("original_name cannot be null.")
            if original_name == "expand":
                return "%24expand"
            if original_name == "select":
                return "%24select"
            if original_name == "include_hidden_folders":
                return "includeHiddenFolders"
            return original_name
        
        # Expand related entities
        expand: Optional[List[str]] = None

        # Include Hidden Folders
        include_hidden_folders: Optional[str] = None

        # Select properties to be returned
        select: Optional[List[str]] = None

    
    @dataclass
    class MailFolderItemRequestBuilderGetRequestConfiguration():
        """
        Configuration for the request such as headers, query parameters, and middleware options.
        """
        # Request headers
        headers: Optional[Dict[str, Union[str, List[str]]]] = None

        # Request options
        options: Optional[List[RequestOption]] = None

        # Request query parameters
        query_parameters: Optional[MailFolderItemRequestBuilder.MailFolderItemRequestBuilderGetQueryParameters] = None

    
    @dataclass
    class MailFolderItemRequestBuilderPatchRequestConfiguration():
        """
        Configuration for the request such as headers, query parameters, and middleware options.
        """
        # Request headers
        headers: Optional[Dict[str, Union[str, List[str]]]] = None

        # Request options
        options: Optional[List[RequestOption]] = None

    

