import { NewtonEntityBase } from '../NewtonEntityBase';
import type { NewtonSDK } from '../NewtonSDK';
import type { Control } from '../types';
import type { Factor, FactorLoadMatch } from '../NewtonTypes';
declare class FactorEntity extends NewtonEntityBase<Factor> {
    constructor(client: NewtonSDK, entopts: any);
    make(this: FactorEntity): FactorEntity;
    load(this: any, reqmatch?: FactorLoadMatch, ctrl?: Control): Promise<FactorEntity>;
}
export { FactorEntity };
