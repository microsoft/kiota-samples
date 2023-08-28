import {MailFolderItemRequestBuilderGetQueryParameters} from './mailFolderItemRequestBuilderGetQueryParameters';
import type {RequestOption} from '@microsoft/kiota-abstractions';

export interface MailFolderItemRequestBuilderGetRequestConfiguration {
    /**
     * Request headers
     */
    headers?: Record<string, string[]> | undefined;
    /**
     * Request options
     */
    options?: RequestOption[] | undefined;
    /**
     * Request query parameters
     */
    queryParameters?: MailFolderItemRequestBuilderGetQueryParameters | undefined;
}
