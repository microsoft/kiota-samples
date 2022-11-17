import {ExtensionItemRequestBuilderGetQueryParameters} from './extensionItemRequestBuilderGetQueryParameters';
import {RequestOption} from '@microsoft/kiota-abstractions';

<<<<<<< HEAD
export interface ExtensionItemRequestBuilderGetRequestConfiguration {
    /** Request headers */
    headers?: Record<string, string> | undefined;
    /** Request options */
    options?: RequestOption[] | undefined;
    /** Request query parameters */
    queryParameters?: ExtensionItemRequestBuilderGetQueryParameters | undefined;
=======
/** Configuration for the request such as headers, query parameters, and middleware options. */
export class ExtensionItemRequestBuilderGetRequestConfiguration {
    /** Request headers */
    public headers?: Record<string, string> | undefined;
    /** Request options */
    public options?: RequestOption[] | undefined;
    /** Request query parameters */
    public queryParameters?: ExtensionItemRequestBuilderGetQueryParameters | undefined;
>>>>>>> main
}
