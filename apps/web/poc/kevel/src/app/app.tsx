import './app.module.scss';

import React, { useState } from 'react';
import { createClient } from '@connectrpc/connect';
import { createConnectTransport } from '@connectrpc/connect-web';
import { create } from '@bufbuild/protobuf';

// eslint-disable-next-line @nx/enforce-module-boundaries
import {
    CertificateAuthorityService,
    CreateCertificateAuthorityRequestSchema,
} from '../../../../../../libs/public/typescript/protobuf/gen/platform/cryptography/v2alpha/certificate_authority_pb';

interface Response {
    text: string;
    sender: 'kevel' | 'user';
}

export function App() {
    const [statement, setStatement] = useState<string>('');
    const [introFinished, setIntroFinished] = useState<boolean>(false);
    const [responses, setResponses] = useState<Response[]>([
        {
            text: 'Input command',
            sender: 'kevel',
        },
    ]);

    const client = createClient(
        CertificateAuthorityService,
        createConnectTransport({
            baseUrl: 'http://144.202.125.179:6477/v2alpha/cryptography/ca/create',
        })
    );

    const send = async (sentence: string) => {
        setResponses((resp) => [...resp, { text: sentence, sender: 'user' }]);
        setStatement('');

        if (introFinished) {
            const res = await client.createCertificateAuthority({
                name: sentence,
            });
        } else {
            const request = create(CreateCertificateAuthorityRequestSchema, {
                name: sentence,
            });

            // Handle error
            await client.createCertificateAuthority(request).catch((error) => {
                setResponses((resp) => [...resp, { text: error?.stack, sender: 'kevel' }]);
                console.error(error);
            });

            setIntroFinished(true);
        }
    };

  /**
   * Handle the statement change event
   */
  const handleStatementChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setStatement(event.target.value);
    };

  /**
   * Handle the send button click event
   */
  const handleSend = () => {
        send(statement);
    };

  /**
   * Process the key press event
   */
  const handleKeyPress = (event: any) => {
        if (event.key === 'Enter') {
            handleSend();
        }
    };

    return (
        <div>
            <header className="app-header">
                <h1>Kevel</h1>
            </header>
            <div className="container">
                {responses.map((resp, i) => {
                    return (
                        <div
                            key={`resp${i}`}
                            className={resp.sender === 'kevel' ? 'kevel-resp-container' : 'user-resp-container'}
                        >
                            <p className="resp-text">{resp.text}</p>
                        </div>
                    );
                })}
                <div>
                    <input
                        type="text"
                        className="text-input"
                        value={statement}
                        onChange={handleStatementChange}
                        onKeyPress={handleKeyPress}
                    />
                    <button onClick={handleSend}>Send</button>
                </div>
            </div>
        </div>
    );
}

export default App;
