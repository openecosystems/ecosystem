import { Configurable } from './configurer';

export interface Bindings {
    registered: Map<string, Binding>;
    registeredListenableChannels: Map<string, Listenable>;
}

export interface Binding {
    name(): string;

    validate(ctx: Context, bindings: Bindings): Promise<void>;

    bind(ctx: Context, bindings: Bindings): Bindings;

    getBinding(): unknown;

    close(): Promise<void>;
}

export interface Listenable {
    listen(ctx: Context, listenerErr: (err: SpecListenableErr) => void): void;
}

export class SpecListenableErr {
    constructor(public error: Error) {}
}

export let Bounds: Bindings | undefined;

export async function registerBindings(ctx: Context, bounds: Binding[]): Promise<Bindings> {
    const registered = new Map<string, Binding>();
    const registeredListenableChannels = new Map<string, Listenable>();
    const bindingsInstance: Bindings = {
        registered,
        registeredListenableChannels,
    };

    const errs: Error[] = [];

    for (const binding of bounds) {
        if ('resolveConfiguration' in binding && 'validateConfiguration' in binding) {
            // const configurable = binding as Configurable;
            // configurable.resolveConfiguration();
            // try {
            //   await configurable.validateConfiguration();
            // } catch (err) {
            //   errs.push(err as Error);
            // }
        }

        try {
            await binding.validate(ctx, bindingsInstance);
        } catch (err) {
            console.error('validate error:', err);
            errs.push(err as Error);
        }

        if (errs.length > 0) {
            console.error('binding errors:', errs);
            throw errs;
        }

        binding.bind(ctx, bindingsInstance);
    }

    Bounds = bindingsInstance;

    return bindingsInstance;
}

export async function shutdownBindings(bindings: Bindings): Promise<void> {
    if (bindings.registered) {
        for (const binding of bindings.registered.values()) {
            try {
                await binding.close();
            } catch (err) {
                console.error(err);
            }
        }
    }
}

export interface Context {
    // A placeholder for context management, similar to Go's context.Context
    cancel?: () => void;
    deadline?: Date;
}
