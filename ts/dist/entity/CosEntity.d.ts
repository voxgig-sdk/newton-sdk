import { NewtonEntityBase } from '../NewtonEntityBase';
import type { NewtonSDK } from '../NewtonSDK';
import type { Control } from '../types';
import type { Cos, CosLoadMatch } from '../NewtonTypes';
declare class CosEntity extends NewtonEntityBase<Cos> {
    constructor(client: NewtonSDK, entopts: any);
    make(this: CosEntity): CosEntity;
    load(this: any, reqmatch?: CosLoadMatch, ctrl?: Control): Promise<CosEntity>;
}
export { CosEntity };
