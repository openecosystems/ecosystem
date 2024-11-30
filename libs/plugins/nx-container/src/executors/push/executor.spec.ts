import { PushExecutorSchema } from './schema';
import executor from './executor';

const options: PushExecutorSchema = {
  image: 'workloads/poc/reference/v1alpha/reference-1',
  version: {
    path: 'apps/workloads/poc/reference/v1alpha/reference-1/package.json',
    key: 'version',
  },
  registries: [],
};

describe('Push Executor', () => {
  it('can run', async () => {
    const output = await executor(options);
    expect(output.success).toBe(true);
  });
});
