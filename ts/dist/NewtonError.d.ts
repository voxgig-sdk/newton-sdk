import { Context } from './Context';
declare class NewtonError extends Error {
    isNewtonError: boolean;
    sdk: string;
    code: string;
    ctx: Context;
    status: number;
    get notFound(): boolean;
    constructor(code: string, msg: string, ctx: Context);
}
export { NewtonError };
