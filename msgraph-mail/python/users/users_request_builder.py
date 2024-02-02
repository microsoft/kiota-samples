from __future__ import annotations
from kiota_abstractions.base_request_builder import BaseRequestBuilder
from kiota_abstractions.get_path_parameters import get_path_parameters
from kiota_abstractions.request_adapter import RequestAdapter
from typing import Any, Callable, Dict, List, Optional, TYPE_CHECKING, Union

if TYPE_CHECKING:
    from .item.user_item_request_builder import UserItemRequestBuilder

class UsersRequestBuilder(BaseRequestBuilder):
    """
    Builds and executes requests for operations under /users
    """
    def __init__(self,request_adapter: RequestAdapter, path_parameters: Union[str, Dict[str, Any]]) -> None:
        """
        Instantiates a new UsersRequestBuilder and sets the default values.
        param path_parameters: The raw url or the url-template parameters for the request.
        param request_adapter: The request adapter to use to execute the requests.
        Returns: None
        """
        super().__init__(request_adapter, "{+baseurl}/users", path_parameters)
    
    def by_user_id(self,user_id: str) -> UserItemRequestBuilder:
        """
        Gets an item from the GraphPythonv1.users.item collection
        param user_id: The unique identifier of user
        Returns: UserItemRequestBuilder
        """
        if not user_id:
            raise TypeError("user_id cannot be null.")
        from .item.user_item_request_builder import UserItemRequestBuilder

        url_tpl_params = get_path_parameters(self.path_parameters)
        url_tpl_params["user%2Did"] = user_id
        return UserItemRequestBuilder(self.request_adapter, url_tpl_params)
    

