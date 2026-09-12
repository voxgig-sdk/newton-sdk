import { NewtonEntityBase } from '../NewtonEntityBase';
import type { NewtonSDK } from '../NewtonSDK';
import type { Control } from '../types';
import type { Zero, ZeroLoadMatch } from '../NewtonTypes';
declare class ZeroEntity extends NewtonEntityBase<Zero> {
    constructor(client: NewtonSDK, entopts: any);
    make(this: ZeroEntity): ZeroEntity;
    load(this: any, reqmatch?: ZeroLoadMatch, ctrl?: Control): Promise<ZeroEntity>;
}
export { ZeroEntity };
