import {MessageItemRequestBuilderGetQueryParameters} from './messageItemRequestBuilderGetQueryParameters';
import {RequestOption} from '@microsoft/kiota-abstractions';

<<<<<<< HEAD
export interface MessageItemRequestBuilderGetRequestConfiguration {
    /** Request headers */
    headers?: Record<string, string> | undefined;
    /** Request options */
    options?: RequestOption[] | undefined;
    /** Request query parameters */
    queryParameters?: MessageItemRequestBuilderGetQueryParameters | undefined;
=======
/** Configuration for the request such as headers, query parameters, and middleware options. */
export class MessageItemRequestBuilderGetRequestConfiguration {
    /** Request headers */
    public headers?: Record<string, string> | undefined;
    /** Request options */
    public options?: RequestOption[] | undefined;
    /** Request query parameters */
    public queryParameters?: MessageItemRequestBuilderGetQueryParameters | undefined;
>>>>>>> main
}
