import assert from 'assert';
import fs from 'fs';
import path from 'path';
import { parseStringPromise } from 'xml2js';

async function countEntries(filePath: string): Promise<number> {
  try {
    const xmlContent: string = fs.readFileSync(filePath, 'utf8');
    const result = await parseStringPromise(xmlContent);

    if (!result || !result.feed || !result.feed.entry) {
      return 0;
    }

    const entries = result.feed.entry;
    return entries.length;
  } catch (error) {
    return 0;
  }
}

it('generated feeds should have entries', async function () {
  const folder = path.join(__dirname, '../dist');
  const files = fs
    .readdirSync(folder)
    .filter((f) => f.endsWith('.xml'))
    .map((f) => path.join(folder, f));

  for (const f of files) {
    const count = await countEntries(f);
    assert.ok(count > 5, f);
  }
});
