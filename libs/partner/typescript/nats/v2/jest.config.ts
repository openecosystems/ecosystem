export default {
    displayName: 'libs-partner-typescript-nats-v2',
    preset: '../../../../../jest.preset.js',
    testEnvironment: 'node',
    transform: {
        '^.+\\.[tj]s$': ['ts-jest', { tsconfig: '<rootDir>/tsconfig.spec.json' }],
    },
    moduleFileExtensions: ['ts', 'js', 'html'],
    coverageDirectory: '../../../../../coverage/libs/partner/typescript/nats/v2',
};
