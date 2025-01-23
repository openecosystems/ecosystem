import { Connector, InboundStream } from '@openecosystems/natsv2';
import type { NatsConnection, Subscription } from "@nats-io/nats-core";
import { DecisionV1 } from "@openecosystems/protobuf-partner"
import { DecisionV1 as DecisionV12 } from "@openecosystems/model-partner"
import { create, toBinary } from '@bufbuild/protobuf';


async function app() {

    const data: Partial<DecisionV1.GetDecisionsResponse> = {
        specContext: {
            $typeName: 'platform.spec.v2.SpecResponseContext',
            responseValidation: {
                $typeName: 'platform.type.v2.ResponseValidation',
                validateOnly: false,
            },
            organizationSlug:"",
            workspaceSlug: "",
            workspaceJan: null,
            routineId: "",
        },
        decisions: {
            $typeName: 'kevel.advertisement.v1.Decisions',

        },
    };

    const response = create(DecisionV1.GetDecisionsResponseSchema, data as DecisionV1.GetDecisionsResponse);
    const binary = toBinary(DecisionV1.GetDecisionsResponseSchema, response);

    const connections: NatsConnection[] = [];
    //connections.push(...await Connector(InboundStream, null, "decision", null));
    connections.push(...await Connector(InboundStream, DecisionV12.CommandDataDecisionTopic, "decision", binary));
    //connections.push(...await Connector("other-echo", 2));

    const a: Promise<void | Error>[] = [];
    connections.forEach((c) => {
    a.push(c.closed());
    });
    await Promise.all(a);

}

app()

