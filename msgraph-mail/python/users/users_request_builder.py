from __future__ import annotations
from kiota_abstractions.get_path_parameters import get_path_parameters
from kiota_abstractions.request_adapter import RequestAdapter
from typing import Any, Callable, Dict, List, Optional, TYPE_CHECKING, Union

if TYPE_CHECKING:
    from .item import user_item_request_builder

class UsersRequestBuilder():
    """
    Builds and executes requests for operations under /users
    """
    def __init__(self,request_adapter: RequestAdapter, path_parameters: Optional[Union[Dict[str, Any], str]] = None) -> None:
        """
        Instantiates a new UsersRequestBuilder and sets the default values.
        Args:
            pathParameters: The raw url or the Url template parameters for the request.
            requestAdapter: The request adapter to use to execute the requests.
        """
        if path_parameters is None:
            raise Exception("path_parameters cannot be undefined")
        if request_adapter is None:
            raise Exception("request_adapter cannot be undefined")
        # Url template to use to build the URL for the current request builder
        self.url_template: str = "{+baseurl}/users"

        url_tpl_params = get_path_parameters(path_parameters)
        self.path_parameters = url_tpl_params
        self.request_adapter = request_adapter
    
    def by_user_id(self,user_id: str) -> user_item_request_builder.UserItemRequestBuilder:
        """
        Gets an item from the GraphPythonv1.users.item collection
        Args:
            user_id: Unique identifier of the item
        Returns: user_item_request_builder.UserItemRequestBuilder
        """
        if user_id is None:
            raise Exception("user_id cannot be undefined")
        from .item import user_item_request_builder

        url_tpl_params = get_path_parameters(self.path_parameters)
        url_tpl_params["user%2Did"] = user_id
        return user_item_request_builder.UserItemRequestBuilder(self.request_adapter, url_tpl_params)
    

