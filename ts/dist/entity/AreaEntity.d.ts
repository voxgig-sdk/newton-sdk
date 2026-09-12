import { NewtonEntityBase } from '../NewtonEntityBase';
import type { NewtonSDK } from '../NewtonSDK';
import type { Control } from '../types';
import type { Area, AreaLoadMatch } from '../NewtonTypes';
declare class AreaEntity extends NewtonEntityBase<Area> {
    constructor(client: NewtonSDK, entopts: any);
    make(this: AreaEntity): AreaEntity;
    load(this: any, reqmatch?: AreaLoadMatch, ctrl?: Control): Promise<AreaEntity>;
}
export { AreaEntity };
