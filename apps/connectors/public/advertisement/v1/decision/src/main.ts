import {libsPartnerTypescriptNatsV2, Connector, GetMultiplexedRequestSubjectName} from "@openecosystems/natsv2"
import type { NatsConnection, Subscription } from "@nats-io/nats-core";


async function app() {

  const subject = GetMultiplexedRequestSubjectName("", "")

  const connections: NatsConnection[] = [];
  connections.push(...await Connector(subject, 3, "echo"));
  connections.push(...await Connector("other-echo", 2));


  const a: Promise<void | Error>[] = [];
  connections.forEach((c) => {
    a.push(c.closed());
  });
  await Promise.all(a);

  console.log('Calling library: ' + libsPartnerTypescriptNatsV2());

}

app()

