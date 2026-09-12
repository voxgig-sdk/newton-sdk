import { NewtonEntityBase } from '../NewtonEntityBase';
import type { NewtonSDK } from '../NewtonSDK';
import type { Control } from '../types';
import type { Arcco, ArccoLoadMatch } from '../NewtonTypes';
declare class ArccoEntity extends NewtonEntityBase<Arcco> {
    constructor(client: NewtonSDK, entopts: any);
    make(this: ArccoEntity): ArccoEntity;
    load(this: any, reqmatch?: ArccoLoadMatch, ctrl?: Control): Promise<ArccoEntity>;
}
export { ArccoEntity };
