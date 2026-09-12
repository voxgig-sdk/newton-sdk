import { NewtonEntityBase } from '../NewtonEntityBase';
import type { NewtonSDK } from '../NewtonSDK';
import type { Control } from '../types';
import type { Derive, DeriveLoadMatch } from '../NewtonTypes';
declare class DeriveEntity extends NewtonEntityBase<Derive> {
    constructor(client: NewtonSDK, entopts: any);
    make(this: DeriveEntity): DeriveEntity;
    load(this: any, reqmatch?: DeriveLoadMatch, ctrl?: Control): Promise<DeriveEntity>;
}
export { DeriveEntity };
