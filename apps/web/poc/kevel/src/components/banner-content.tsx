import { Content } from '@adzerk/decision-sdk';

export function BannerContent(content: Content) {
    const height: number = (content?.data as any)?.height ?? 50;
    const width: number = (content?.data as any)?.width ?? 100;

    return (
        content?.body && (
            <div
                style={{
                    height: height,
                    width: width,
                }}
                dangerouslySetInnerHTML={{ __html: content.body }}
            ></div>
        )
    );
}
