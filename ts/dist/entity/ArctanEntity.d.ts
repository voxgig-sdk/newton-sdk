import { NewtonEntityBase } from '../NewtonEntityBase';
import type { NewtonSDK } from '../NewtonSDK';
import type { Control } from '../types';
import type { Arctan, ArctanLoadMatch } from '../NewtonTypes';
declare class ArctanEntity extends NewtonEntityBase<Arctan> {
    constructor(client: NewtonSDK, entopts: any);
    make(this: ArctanEntity): ArctanEntity;
    load(this: any, reqmatch?: ArctanLoadMatch, ctrl?: Control): Promise<ArctanEntity>;
}
export { ArctanEntity };
