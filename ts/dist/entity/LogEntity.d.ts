import { NewtonEntityBase } from '../NewtonEntityBase';
import type { NewtonSDK } from '../NewtonSDK';
import type { Control } from '../types';
import type { Log, LogLoadMatch } from '../NewtonTypes';
declare class LogEntity extends NewtonEntityBase<Log> {
    constructor(client: NewtonSDK, entopts: any);
    make(this: LogEntity): LogEntity;
    load(this: any, reqmatch?: LogLoadMatch, ctrl?: Control): Promise<LogEntity>;
}
export { LogEntity };
