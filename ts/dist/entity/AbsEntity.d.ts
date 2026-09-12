import { NewtonEntityBase } from '../NewtonEntityBase';
import type { NewtonSDK } from '../NewtonSDK';
import type { Control } from '../types';
import type { Abs, AbsLoadMatch } from '../NewtonTypes';
declare class AbsEntity extends NewtonEntityBase<Abs> {
    constructor(client: NewtonSDK, entopts: any);
    make(this: AbsEntity): AbsEntity;
    load(this: any, reqmatch?: AbsLoadMatch, ctrl?: Control): Promise<AbsEntity>;
}
export { AbsEntity };
