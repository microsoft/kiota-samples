import {RequestOption} from '@microsoft/kiota-abstractions';

<<<<<<< HEAD
export interface ExtensionItemRequestBuilderDeleteRequestConfiguration {
    /** Request headers */
    headers?: Record<string, string> | undefined;
    /** Request options */
    options?: RequestOption[] | undefined;
=======
/** Configuration for the request such as headers, query parameters, and middleware options. */
export class ExtensionItemRequestBuilderDeleteRequestConfiguration {
    /** Request headers */
    public headers?: Record<string, string> | undefined;
    /** Request options */
    public options?: RequestOption[] | undefined;
>>>>>>> main
}
