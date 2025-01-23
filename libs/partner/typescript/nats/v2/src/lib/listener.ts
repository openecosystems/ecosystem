import type { NatsConnection, Subscription } from '@nats-io/nats-core/lib/mod';
import { Stream, GetQueueGroupName, GetMultiplexedRequestSubjectName } from './stream';
import { connect } from '@nats-io/transport-node';
import { fromBinary } from '@bufbuild/protobuf';
import { SpecSchema } from '@openecosystems/protobuf-partner';
import { Client } from "@adzerk/decision-sdk";


export async function Connector(scope: Stream, subject: string, entity: string, binary: Uint8Array): Promise<NatsConnection[]> {
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
  const _ = handleRequest(s, subscription, binary);
  console.log(`${subject} is listening for ${s} requests...`);
  connections.push(nc);

  return connections;
}

// simple handler for service requests
async function handleRequest(subject: string, s: Subscription, binary: Uint8Array) {
  const p = 12 - subject.length;
  const pad = "".padEnd(p);
  for await (const m of s) {

      const spec = fromBinary(SpecSchema, m.data)
      console.log(`Received ${p}: ${spec}`);

      let client = new Client({ networkId: 11603, siteId: 1301620 });

      let request = {
          placements: [{ adTypes: [3] }],
          user: { key: "abc" },
          keywords: ["keyword1", "keyword2"]
      };

      client.decisions.get(request).then(response => {
          console.dir(response, { depth: null });
      });


    // respond returns true if the message had a reply subject, thus it could respond
    //if (m.respond(m.data)) {
    if (m.respond(binary)) {
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
