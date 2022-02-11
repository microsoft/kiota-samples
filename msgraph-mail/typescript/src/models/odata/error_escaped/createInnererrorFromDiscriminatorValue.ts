import {Innererror} from './innererror';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createInnererrorFromDiscriminatorValue(parseNode: ParseNode | undefined) : Innererror {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new Innererror();
}
