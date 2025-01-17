import type { NatsConnection, Subscription } from '@nats-io/nats-core/lib/mod';
import { Stream } from './stream';
import { connect } from '@nats-io/transport-node';
import { GetMultiplexedRequestSubjectName } from './registry';

export function getQueueGroupName(scope: Stream, entityName: string): string {
  return `req.${scope}-${entityName}`;
}

export async function Connector(scope: Stream, subject: string, entity: string): Promise<NatsConnection[]> {
  const connections: NatsConnection[] = [];

  const queue = getQueueGroupName(scope, entity)

  const nc = await connect({ servers: "localhost:4222", name: `${subject}` });
  nc.closed().then((err) => {
    if (err) {
      console.error(`service ${subject} exited because of error: ${err.message}`);
    }
  });

  const s = GetMultiplexedRequestSubjectName(scope, subject)
  const subscription = nc.subscribe(s, { queue: queue });
  const _ = handleRequest(subject, subscription);
  console.log(`${subject} is listening for ${s} requests...`);
  connections.push(nc);

  return connections;
}

// simple handler for service requests
async function handleRequest(subject: string, s: Subscription) {
  const p = 12 - subject.length;
  const pad = "".padEnd(p);
  for await (const m of s) {
    // respond returns true if the message had a reply subject, thus it could respond
    if (m.respond(m.data)) {
      console.log(
        `[${subject}]:${pad} #${s.getProcessed()} echoed ${m.string()}`,
      );
    } else {
      console.log(
        `[${subject}]:${pad} #${s.getProcessed()} ignoring request - no reply subject`,
      );
    }
  }
}




//
// async function app() {
//
//   const nc = await connect({ servers: "localhost:4222" });
//
//   //const nc = await connect({ servers: "localhost:4222" });
//
//   const s1 = nc.subscribe("n.>");
//
//   async function printMsgs(s: Subscription) {
//     const subj = s.getSubject();
//     console.log(`listening for ${subj}`);
//     const c = 13 - subj.length;
//     const pad = "".padEnd(c);
//     for await (const m of s) {
//       console.log(
//         `[${subj}]${pad} #${s.getProcessed()} - ${m.subject} ${
//           m.data ? " " + m.string() : ""
//         }`,
//       );
//     }
//   }
//
//   printMsgs(s1);
//
//   await nc.closed();
//
// }
//
// app()
