import {libsPartnerTypescriptNatsV2} from "@openecosystems/natsv2"
import { connect } from "@nats-io/transport-node";
import type { Subscription } from "@nats-io/nats-core";

async function app() {

  const nc = await connect({ servers: "api.dev-1.na-us-1.oeco.cloud:4222" });
  //const nc = await connect({ servers: "localhost:4222" });

  const s1 = nc.subscribe("configuration.>");

  async function printMsgs(s: Subscription) {
    const subj = s.getSubject();
    console.log(`listening for ${subj}`);
    const c = 13 - subj.length;
    const pad = "".padEnd(c);
    for await (const m of s) {
      console.log(
        `[${subj}]${pad} #${s.getProcessed()} - ${m.subject} ${
          m.data ? " " + m.string() : ""
        }`,
      );
    }
  }

  printMsgs(s1);

  await nc.closed();

}

console.log('Calling library: ' + libsPartnerTypescriptNatsV2());
app()
