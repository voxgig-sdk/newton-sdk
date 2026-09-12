import { NewtonEntityBase } from '../NewtonEntityBase';
import type { NewtonSDK } from '../NewtonSDK';
import type { Control } from '../types';
import type { Tan, TanLoadMatch } from '../NewtonTypes';
declare class TanEntity extends NewtonEntityBase<Tan> {
    constructor(client: NewtonSDK, entopts: any);
    make(this: TanEntity): TanEntity;
    load(this: any, reqmatch?: TanLoadMatch, ctrl?: Control): Promise<TanEntity>;
}
export { TanEntity };
