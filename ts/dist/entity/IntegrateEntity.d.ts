import { NewtonEntityBase } from '../NewtonEntityBase';
import type { NewtonSDK } from '../NewtonSDK';
import type { Control } from '../types';
import type { Integrate, IntegrateLoadMatch } from '../NewtonTypes';
declare class IntegrateEntity extends NewtonEntityBase<Integrate> {
    constructor(client: NewtonSDK, entopts: any);
    make(this: IntegrateEntity): IntegrateEntity;
    load(this: any, reqmatch?: IntegrateLoadMatch, ctrl?: Control): Promise<IntegrateEntity>;
}
export { IntegrateEntity };
