import { Connector, InboundStream } from '@openecosystems/natsv2';
import type { NatsConnection } from '@nats-io/nats-core';
import { DecisionV1 as DecisionV12 } from '@openecosystems/model-partner';
import {getDecisions} from './get_decisions';


async function app() {

    const connections: NatsConnection[] = [];
    connections.push(...await Connector(InboundStream, DecisionV12.CommandDataDecisionTopic, "decision", getDecisions));
    //connections.push(...await Connector("other-echo", 2));

    const a: Promise<void | Error>[] = [];
    connections.forEach((c) => {
        a.push(c.closed());
    });
    await Promise.all(a);
}

app()

