import { Connector, InboundStream } from '@openecosystems/natsv2';
import type { NatsConnection, Subscription } from "@nats-io/nats-core";
import { DecisionV1 } from "@openecosystems/model-partner"


async function app() {

  const connections: NatsConnection[] = [];
  connections.push(...await Connector(InboundStream, DecisionV1.CommandDataDecisionTopic, "decision"));
  //connections.push(...await Connector("other-echo", 2));


  const a: Promise<void | Error>[] = [];
  connections.forEach((c) => {
    a.push(c.closed());
  });
  await Promise.all(a);

}

app()

