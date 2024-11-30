import {
  mkdtemp,
  mkdir,
  rm,
  access,
  constants,
  readFile,
} from 'node:fs/promises';
import { join } from 'node:path';
import { tmpdir } from 'node:os';
import { readJsonFile } from '@nx/devkit';
import fetch from 'node-fetch';
import { downloadFile } from './fs';
import { readYamlFile } from './yaml';

const { Response } = jest.requireActual('node-fetch');

jest.mock('node-fetch');

describe('downloadFile', () => {
  const relativeDownloadPath = 'src';
  let dir: string;
  let downloadPath: string;

  beforeEach(async () => {
    dir = await mkdtemp(join(tmpdir(), 'downloadFile-'));
    downloadPath = join(dir, relativeDownloadPath);
    await mkdir(downloadPath);
  });

  afterEach(async () => {
    try {
      await access(dir, constants.F_OK);
      await rm(dir, { recursive: true, force: true }); // clean up
    } catch (e) {
      console.log('error during cleanup: ' + e);
    }
  });

  it('should download JSON response to the specified location', async () => {
    const expectedResponse = { test: 'test' };
    (fetch as jest.MockedFunction<typeof fetch>).mockResolvedValueOnce(
      new Response(JSON.stringify(expectedResponse), {
        headers: { 'content-type': 'application/json' },
      })
    );

    await downloadFile('http://fake.com/file.json', downloadPath);
    expect(fetch).toHaveBeenCalledWith('http://fake.com/file.json');
    await expect(access(join(downloadPath, 'file.json'), constants.F_OK))
      .resolves;
    expect(readJsonFile(join(downloadPath, 'file.json'))).toEqual(
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

    await downloadFile('http://fake.com/test.yaml', downloadPath);
    expect(fetch).toHaveBeenCalledWith('http://fake.com/test.yaml');
    await expect(access(join(downloadPath, 'test.yaml'), constants.F_OK))
      .resolves;
    await expect(
      readYamlFile(join(downloadPath, 'test.yaml'))
    ).resolves.toEqual(expectedResponse);
  });

  it('should download allow explicit setting of fileName', async () => {
    const expectedResponse = { test: 'test' };
    (fetch as jest.MockedFunction<typeof fetch>).mockResolvedValueOnce(
      new Response(JSON.stringify(expectedResponse), {
        headers: { 'content-type': 'text/yaml' },
      })
    );

    await downloadFile('http://fake.com/test.yaml', downloadPath, 'myfile');
    expect(fetch).toHaveBeenCalledWith('http://fake.com/test.yaml');
    await expect(access(join(downloadPath, 'myfile.yaml'), constants.F_OK))
      .resolves;
    await expect(
      readYamlFile(join(downloadPath, 'myfile.yaml'))
    ).resolves.toEqual(expectedResponse);
  });

  it('should download allow explicit setting of fileExtension', async () => {
    const expectedResponse = { test: 'test' };
    (fetch as jest.MockedFunction<typeof fetch>).mockResolvedValueOnce(
      new Response(JSON.stringify(expectedResponse), {
        headers: { 'content-type': 'text/yaml' },
      })
    );

    await downloadFile(
      'http://fake.com/test.yaml',
      downloadPath,
      'myfile',
      'ext'
    );
    expect(fetch).toHaveBeenCalledWith('http://fake.com/test.yaml');
    await expect(access(join(downloadPath, 'myfile.ext'), constants.F_OK))
      .resolves;
  });

  it('should not have an extension when content-type is unrecognized', async () => {
    const expectedResponse = 'somecontent';
    (fetch as jest.MockedFunction<typeof fetch>).mockResolvedValueOnce(
      new Response(expectedResponse, {
        headers: { 'content-type': 'application/some-weird-thing' },
      })
    );

    await downloadFile('http://fake.com/test.thing', downloadPath);
    expect(fetch).toHaveBeenCalledWith('http://fake.com/test.thing');
    await expect(access(join(downloadPath, 'test'), constants.F_OK)).resolves;
    await expect(
      readFile(join(downloadPath, 'test'), { encoding: 'utf-8' })
    ).resolves.toEqual(expectedResponse);
  });
});
