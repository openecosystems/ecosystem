import './app.module.scss';

import React, { useState } from 'react';
import { createClient } from '@connectrpc/connect';
import { createConnectTransport } from '@connectrpc/connect-web';
import { create } from '@bufbuild/protobuf';

import { Response } from '../models/Response';
import { UserCommand } from '../models/UserCommand';

import { ConfigurationV2Alpha } from '@openecosystems/protobuf-public';
import { Command } from '../models/Command';
import { Banner } from '../components/banner';

// HACK to allow BigInt to be serialized to JSON
(BigInt.prototype as any)['toJSON'] = function () {
    return this.toString();
};

/**
 * List of commands
 */
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

/**
 * Headers required for the API
 */
const headers = new Headers({
    'x-spec-workspace-slug': 'workspace123',
    'x-spec-organization-slug': 'organization123',
});

export function App() {
    const [selectedCommand, setSelectedCommand] = useState<Command>(Command.GetConfiguration);
    const [introFinished, setIntroFinished] = useState<boolean>(false);
    const [responses, setResponses] = useState<Response[]>([
        {
            text: 'Select and send command',
            sender: 'kevel',
        },
    ]);

    const client = createClient(
        ConfigurationV2Alpha.ConfigurationService,
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
        // Data to send
        const data: Partial<ConfigurationV2Alpha.GetConfigurationRequest> = {
            id: '1',
        };

        if (introFinished) {
            const response = await client
                .getConfiguration(data as ConfigurationV2Alpha.GetConfigurationRequest, {
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
            const request = create(
                ConfigurationV2Alpha.GetConfigurationRequestSchema,
                data as ConfigurationV2Alpha.GetConfigurationRequest
            );

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
        // Data to send
        const data: Partial<ConfigurationV2Alpha.CreateConfigurationRequest> = {
            type: 1,
            parentId: '123',
        };

        if (introFinished) {
            const response = await client
                .createConfiguration(data as ConfigurationV2Alpha.CreateConfigurationRequest, {
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
            const request = create(
                ConfigurationV2Alpha.CreateConfigurationRequestSchema,
                data as ConfigurationV2Alpha.CreateConfigurationRequest
            );

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
                <h4>React/esbuild</h4>
            </header>
            <div className="banner">
                <Banner></Banner>
            </div>
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
