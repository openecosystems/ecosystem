import type { Msg, NatsConnection, Subscription } from '@nats-io/nats-core/lib/mod';
import { Stream, GetQueueGroupName, GetMultiplexedRequestSubjectName } from './stream';
import { connect } from '@nats-io/transport-node';
import { fromBinary } from '@bufbuild/protobuf';
import { SpecSchema } from '@openecosystems/protobuf-partner';

export async function Connector(scope: Stream, subject: string, entity: string, process: (spec: any, m: Msg) => Promise<Uint8Array>): Promise<NatsConnection[]> {
  const connections: NatsConnection[] = [];

  const queue = GetQueueGroupName(scope, entity)
  const s = GetMultiplexedRequestSubjectName(scope, subject)

  const nc = await connect({ servers: "localhost:4222", name: `${s}` });
  nc.closed().then((err) => {
    if (err) {
      console.error(`service ${subject} exited because of error: ${err.message}`);
    }
  });

  const subscription = nc.subscribe(s, { queue: queue });

    const _ = handleRequest(s, subscription, process);
    console.log(`${subject} is listening for ${s} requests...`);
    connections.push(nc);

  return connections;
}

// simple handler for service requests
async function handleRequest(subject: string, s: Subscription, process: (spec: any, m: Msg) => Promise<Uint8Array>) {
  const p = 12 - subject.length;
  const pad = "".padEnd(p);
  for await (const m of s) {

      const spec = fromBinary(SpecSchema, m.data)
      const binary = await process(spec, m)

      // respond returns true if the message had a reply subject, thus it could respond
      if (m.respond(binary)) {
          console.log(
              `[${subject}]:${pad} #${s.getProcessed()}`,
          );
      } else {
          console.log(
              `[${subject}]:${pad} #${s.getProcessed()} ignoring request - no reply subject`,
          );
      }
  }
}
