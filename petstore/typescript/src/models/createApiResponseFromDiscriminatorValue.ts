import {ApiResponse} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createApiResponseFromDiscriminatorValue(parseNode: ParseNode | undefined) : ApiResponse {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new ApiResponse();
}
