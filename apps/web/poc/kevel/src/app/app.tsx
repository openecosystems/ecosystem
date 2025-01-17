import './app.module.scss';

import React, { useState } from 'react';
import { createClient } from '@connectrpc/connect';
import { createConnectTransport } from '@connectrpc/connect-web';
import { create } from '@bufbuild/protobuf';

// eslint-disable-next-line @nx/enforce-module-boundaries
import {
    ConfigurationService,
    CreateConfigurationRequest,
    CreateConfigurationRequestSchema,
    GetConfigurationRequest,
    GetConfigurationRequestSchema,
} from '../../../../../../libs/public/typescript/protobuf/gen/platform/configuration/v2alpha/configuration_pb';

// HACK to allow BigInt to be serialized to JSON
(BigInt.prototype as any)['toJSON'] = function () {
    return this.toString();
};

enum Command {
    GetConfiguration = 'getConfig',
    CreateConfiguration = 'createConfig',
}

interface Response {
    text: string;
    sender: 'kevel' | 'user';
}

interface UserCommand {
    id: Command;
    name: string;
}

const commands: UserCommand[] = [
    {
        id: Command.GetConfiguration,
        name: 'Get configuration',
    },
    {
        id: Command.CreateConfiguration,
        name: 'Create configuration',
    },
];

export function App() {
    const [selectedCommand, setSelectedCommand] = useState<Command>(Command.GetConfiguration);
    const [introFinished, setIntroFinished] = useState<boolean>(false);
    const [responses, setResponses] = useState<Response[]>([
        {
            text: 'Select and send command',
            sender: 'kevel',
        },
    ]);

    const headers = new Headers();

    headers.set('x-spec-workspace-slug', 'workspace123');
    headers.set('x-spec-organization-slug', 'organization123');

    const client = createClient(
        ConfigurationService,
        createConnectTransport({
            //baseUrl: 'http://api.dev-1.na-us-1.oeco.cloud:6477',
            baseUrl: 'http://localhost:6477',
        })
    );

    const send = async (command: Command) => {
        setResponses((resp) => [...resp, { text: command, sender: 'user' }]);

        if (command === Command.CreateConfiguration) {
            await createConfigurationCommand();
        } else if (command === Command.GetConfiguration) {
            await getConfigurationCommand();
        }
    };

    /**
     * Get configuration command
     */
    const getConfigurationCommand = async () => {
        const data: Partial<GetConfigurationRequest> = {
            id: '1',
        };

        if (introFinished) {
            const response = await client
                .getConfiguration(data as GetConfigurationRequest, {
                    headers: headers,
                })
                .catch((error) => {
                    setResponses((resp) => [...resp, { text: error?.stack, sender: 'kevel' }]);
                    console.error(error);
                });
            if (response?.configuration) {
                setResponses((resp) => [...resp, { text: JSON.stringify(response.configuration), sender: 'kevel' }]);
            }
        } else {
            const request = create(GetConfigurationRequestSchema, data as GetConfigurationRequest);

            // Handle error
            const response = await client
                .getConfiguration(request, {
                    headers: headers,
                })
                .catch((error) => {
                    setResponses((resp) => [...resp, { text: error?.stack, sender: 'kevel' }]);
                    console.error(error);
                });

            if (response?.configuration) {
                setResponses((resp) => [...resp, { text: JSON.stringify(response.configuration), sender: 'kevel' }]);
            }

            setIntroFinished(true);
        }
    };

    /**
     * Create configuration command
     */
    const createConfigurationCommand = async () => {
        const data: Partial<CreateConfigurationRequest> = {
            type: 1,
            parentId: '123',
        };

        if (introFinished) {
            const response = await client
                .createConfiguration(data as CreateConfigurationRequest, {
                    headers: headers,
                })
                .catch((error) => {
                    setResponses((resp) => [...resp, { text: error?.stack, sender: 'kevel' }]);
                    console.error(error);
                });
            if (response?.configuration) {
                setResponses((resp) => [...resp, { text: JSON.stringify(response.configuration), sender: 'kevel' }]);
            }
        } else {
            const request = create(CreateConfigurationRequestSchema, data as CreateConfigurationRequest);

            // Handle error
            const response = await client
                .createConfiguration(request, {
                    headers: headers,
                })
                .catch((error) => {
                    setResponses((resp) => [...resp, { text: error?.stack, sender: 'kevel' }]);
                    console.error(error);
                });

            if (response?.configuration) {
                setResponses((resp) => [...resp, { text: JSON.stringify(response.configuration), sender: 'kevel' }]);
            }

            setIntroFinished(true);
        }
    };

    /**
     * Handle the send button click event
     */
    const handleSend = () => {
        send(selectedCommand);
    };

    /**
     * Process the key press event
     */
    const handleKeyPress = (event: any) => {
        if (event.key === 'Enter') {
            handleSend();
        }
    };

    const handleCommandChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
        setSelectedCommand(event.target.value as Command);
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
                    <select className="text-input" onChange={handleCommandChange} onKeyPress={handleKeyPress}>
                        {commands.map((command) => (
                            <option key={command.id} value={command.id}>
                                {command.name}
                            </option>
                        ))}
                    </select>

                    <button onClick={handleSend}>Send</button>
                </div>
            </div>
        </div>
    );
}

export default App;
