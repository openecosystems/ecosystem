export default {
    displayName: 'libs-public-typescript-connector-v2alpha',
    preset: '../../../../../jest.preset.js',
    testEnvironment: 'node',
    transform: {
        '^.+\\.[tj]s$': ['ts-jest', { tsconfig: '<rootDir>/tsconfig.spec.json' }],
    },
    moduleFileExtensions: ['ts', 'js', 'html'],
    coverageDirectory: '../../../../../coverage/libs/public/typescript/connector/v2alpha',
};
