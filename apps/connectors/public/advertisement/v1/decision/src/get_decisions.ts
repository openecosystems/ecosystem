import { DecisionV1, Spec } from '@openecosystems/protobuf-partner';
import type { Msg } from '@nats-io/nats-core/lib/mod';
import { Client, Content } from '@adzerk/decision-sdk';
import { create, toBinary } from '@bufbuild/protobuf';

export async function getDecisions(spec: Spec, m: Msg): Promise<Uint8Array> {

    let client = new Client({ networkId: 11603, siteId: 1301620 });

    let request = {
        placements: [{ adTypes: [3] }],
        user: { key: "local.developer" },
        keywords: ["keyword1", "keyword2"]
    };

    let binary : Uint8Array
    await client.decisions.get(request).then(response => {

        const content: Partial<Content> = response.decisions['div0'][0].contents[0]

        let cdata: DecisionV1.Data
        if (content.data && typeof content.data === "object") {
            // Narrowing: Access properties only after narrowing
            cdata = (content.data as DecisionV1.Data)
        }

        const contentData: DecisionV1.Data = {
            $typeName: 'kevel.advertisement.v1.Data',
            fileName: cdata.fileName,
            height: cdata.height,
            imageUrl: cdata.imageUrl,
            width: cdata.width,
        }

        const data: Partial<DecisionV1.GetDecisionsResponse> = {
            $typeName: 'kevel.advertisement.v1.GetDecisionsResponse',
            specContext: {
                $typeName: 'platform.spec.v2.SpecResponseContext',
                responseValidation: {
                    $typeName: 'platform.type.v2.ResponseValidation',
                    validateOnly: spec.context.validation.validateOnly,
                },
                organizationSlug: spec.context.organizationSlug,
                workspaceSlug: spec.context.workspaceSlug,
                workspaceJan: spec.context.workspaceJan,
                routineId: "",
            },
            user: {
                $typeName: 'kevel.advertisement.v1.User',
                key: response.user.key,
            },
            decision: {
                $typeName: 'kevel.advertisement.v1.Decision',
                adId: BigInt(response.decisions['div0'][0].adId),
                advertiserId: BigInt(response.decisions['div0'][0].advertiserId),
                creativeId: BigInt(response.decisions['div0'][0].creativeId),
                flightId: BigInt(response.decisions['div0'][0].flightId),
                campaignId: BigInt(response.decisions['div0'][0].campaignId),
                priorityId: BigInt(response.decisions['div0'][0].priorityId),
                clickUrl: response.decisions['div0'][0].clickUrl,
                contents: [{
                    $typeName: 'kevel.advertisement.v1.Content',
                    type: content.type,
                    template: content.template,
                    body: content.body,
                    customTemplate: content.customTemplate,
                    data: contentData,
                }],
                impressionUrl: response.decisions['div0'][0].impressionUrl,
                events: [],
                matchedPoints: [],
                // pricing: {
                //     $typeName: 'kevel.advertisement.v1.PricingData',
                // }
            }
        };

        const res = create(DecisionV1.GetDecisionsResponseSchema, data as DecisionV1.GetDecisionsResponse);
        binary = toBinary(DecisionV1.GetDecisionsResponseSchema, res)

        //console.log(response['div0']);
        //console.dir(response, { depth: null });
    });

    return binary

}
