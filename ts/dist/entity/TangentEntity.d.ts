import { NewtonEntityBase } from '../NewtonEntityBase';
import type { NewtonSDK } from '../NewtonSDK';
import type { Control } from '../types';
import type { Tangent, TangentLoadMatch } from '../NewtonTypes';
declare class TangentEntity extends NewtonEntityBase<Tangent> {
    constructor(client: NewtonSDK, entopts: any);
    make(this: TangentEntity): TangentEntity;
    load(this: any, reqmatch?: TangentLoadMatch, ctrl?: Control): Promise<TangentEntity>;
}
export { TangentEntity };
