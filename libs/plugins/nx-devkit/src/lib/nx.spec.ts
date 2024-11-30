import { mkdtemp, mkdir, rm } from 'node:fs/promises';
import { join } from 'node:path';
import { tmpdir } from 'node:os';
import { readJson } from '@nx/devkit';
import { FsTree } from 'nx/src/generators/tree';
import fetch from 'node-fetch';
import { download } from './nx';
import { readYaml } from './yaml';

const { Response } = jest.requireActual('node-fetch');

jest.mock('node-fetch');

describe('downloadFile', () => {
  const downloadPath = 'src';
  let tree: FsTree;

  beforeEach(async () => {
    const tmpDir = await mkdtemp(join(tmpdir(), 'download-'));
    tree = new FsTree(tmpDir, false);

    await mkdir(join(tmpDir, downloadPath));
  });

  afterEach(async () => {
    if (tree) {
      await rm(tree.root, { recursive: true, force: true }); // clean up
    }
  });

  it('should download JSON response to the specified location', async () => {
    const expectedResponse = { test: 'test' };
    (fetch as jest.MockedFunction<typeof fetch>).mockResolvedValueOnce(
      new Response(JSON.stringify(expectedResponse), {
        headers: { 'content-type': 'application/json' },
      })
    );

    await download(tree, 'http://fake.com/file.json', downloadPath);
    expect(fetch).toHaveBeenCalledWith('http://fake.com/file.json');
    expect(tree.exists(join(downloadPath, 'file.json'))).toBeTruthy();
    expect(readJson(tree, join(downloadPath, 'file.json'))).toEqual(
      expectedResponse
    );
  });

  it('should download YAML response to the specified location', async () => {
    const expectedResponse = { test: 'test' };
    (fetch as jest.MockedFunction<typeof fetch>).mockResolvedValueOnce(
      new Response(JSON.stringify(expectedResponse), {
        headers: { 'content-type': 'text/yaml' },
      })
    );

    await download(tree, 'http://fake.com/test.yaml', downloadPath);
    expect(fetch).toHaveBeenCalledWith('http://fake.com/test.yaml');
    expect(tree.exists(join(downloadPath, 'test.yaml'))).toBeTruthy();
    expect(readYaml(tree, join(downloadPath, 'test.yaml'))).toEqual(
      expectedResponse
    );
  });

  it('should download allow explicit setting of fileName', async () => {
    const expectedResponse = { test: 'test' };
    (fetch as jest.MockedFunction<typeof fetch>).mockResolvedValueOnce(
      new Response(JSON.stringify(expectedResponse), {
        headers: { 'content-type': 'text/yaml' },
      })
    );

    await download(tree, 'http://fake.com/test.yaml', downloadPath, 'myfile');
    expect(fetch).toHaveBeenCalledWith('http://fake.com/test.yaml');
    expect(tree.exists(join(downloadPath, 'myfile.yaml'))).toBeTruthy();
    expect(readYaml(tree, join(downloadPath, 'myfile.yaml'))).toEqual(
      expectedResponse
    );
  });

  it('should download allow explicit setting of fileExtension', async () => {
    const expectedResponse = { test: 'test' };
    (fetch as jest.MockedFunction<typeof fetch>).mockResolvedValueOnce(
      new Response(JSON.stringify(expectedResponse), {
        headers: { 'content-type': 'text/yaml' },
      })
    );

    await download(
      tree,
      'http://fake.com/test.yaml',
      downloadPath,
      'myfile',
      'ext'
    );
    expect(fetch).toHaveBeenCalledWith('http://fake.com/test.yaml');
    expect(tree.exists(join(downloadPath, 'myfile.ext'))).toBeTruthy();
    expect(readYaml(tree, join(downloadPath, 'myfile.ext'))).toEqual(
      expectedResponse
    );
  });

  it('should not have an extension when content-type is unrecognized', async () => {
    const expectedResponse = 'somecontent';
    (fetch as jest.MockedFunction<typeof fetch>).mockResolvedValueOnce(
      new Response(expectedResponse, {
        headers: { 'content-type': 'application/some-weird-thing' },
      })
    );

    await download(tree, 'http://fake.com/test.thing', downloadPath);
    expect(fetch).toHaveBeenCalledWith('http://fake.com/test.thing');
    expect(tree.exists(join(downloadPath, 'test'))).toBeTruthy();
    expect(tree.read(join(downloadPath, 'test')).toString()).toEqual(
      expectedResponse
    );
  });
});
