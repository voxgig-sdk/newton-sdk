import { NewtonEntityBase } from '../NewtonEntityBase';
import type { NewtonSDK } from '../NewtonSDK';
import type { Control } from '../types';
import type { Sin, SinLoadMatch } from '../NewtonTypes';
declare class SinEntity extends NewtonEntityBase<Sin> {
    constructor(client: NewtonSDK, entopts: any);
    make(this: SinEntity): SinEntity;
    load(this: any, reqmatch?: SinLoadMatch, ctrl?: Control): Promise<SinEntity>;
}
export { SinEntity };
