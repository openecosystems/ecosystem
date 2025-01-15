import {libsPartnerTypescriptNatsV2} from "@openecosystems/natsv2"
console.log('Hello World ' + libsPartnerTypescriptNatsV2());
/*
 * Copyright 2020-2021 The NATS Authors
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { connect } from "@nats-io/transport-node";
import type { Subscription } from "@nats-io/nats-core";

async function app() {

  const nc = await connect({ servers: "api.dev-1.na-us-1.oeco.cloud:4222" });

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

app()
