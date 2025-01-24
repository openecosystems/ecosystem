import { useEffect, useState } from 'react';
import { Content, Decision, DecisionResponse } from '@adzerk/decision-sdk';
import { BannerContent } from './banner-content';

const API_URL = 'assets/kevel-response.json';

export function Banner() {
    const [decisionData, setDecisionData] = useState<any>(null);

    useEffect(() => {
        const fetchData = async () => {
            const response = await fetch(API_URL);
            const result: DecisionResponse = await response.json();

            const decision: Decision | undefined = result?.decisions?.['div0']?.[0];
            setDecisionData(decision);
        };
        fetchData();
    }, []);

    return (
        decisionData?.contents &&
        decisionData.contents.map((content: Content, index: number) => (
            <BannerContent key={index} {...content}></BannerContent>
        ))
    );
}
