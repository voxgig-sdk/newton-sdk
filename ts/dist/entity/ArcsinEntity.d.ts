import { NewtonEntityBase } from '../NewtonEntityBase';
import type { NewtonSDK } from '../NewtonSDK';
import type { Control } from '../types';
import type { Arcsin, ArcsinLoadMatch } from '../NewtonTypes';
declare class ArcsinEntity extends NewtonEntityBase<Arcsin> {
    constructor(client: NewtonSDK, entopts: any);
    make(this: ArcsinEntity): ArcsinEntity;
    load(this: any, reqmatch?: ArcsinLoadMatch, ctrl?: Control): Promise<ArcsinEntity>;
}
export { ArcsinEntity };
