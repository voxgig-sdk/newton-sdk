import { NewtonEntityBase } from '../NewtonEntityBase';
import type { NewtonSDK } from '../NewtonSDK';
import type { Control } from '../types';
import type { Simplify, SimplifyLoadMatch } from '../NewtonTypes';
declare class SimplifyEntity extends NewtonEntityBase<Simplify> {
    constructor(client: NewtonSDK, entopts: any);
    make(this: SimplifyEntity): SimplifyEntity;
    load(this: any, reqmatch?: SimplifyLoadMatch, ctrl?: Control): Promise<SimplifyEntity>;
}
export { SimplifyEntity };
