import {RequestOption} from '@microsoft/kiota-abstractions';

<<<<<<< HEAD
export interface ExtensionsRequestBuilderPostRequestConfiguration {
    /** Request headers */
    headers?: Record<string, string> | undefined;
    /** Request options */
    options?: RequestOption[] | undefined;
=======
/** Configuration for the request such as headers, query parameters, and middleware options. */
export class ExtensionsRequestBuilderPostRequestConfiguration {
    /** Request headers */
    public headers?: Record<string, string> | undefined;
    /** Request options */
    public options?: RequestOption[] | undefined;
>>>>>>> main
}
