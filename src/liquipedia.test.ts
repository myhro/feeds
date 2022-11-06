import assert from 'assert';
import path from 'path';
import { readFileSync } from 'fs';
import { load } from 'cheerio';
import { description, descriptionWithoutRef, entryID } from './liquipedia';

const Brame = 'Brame';
const BrameDiffRef = 'BrameDiffRef';
const Creepwave = 'Creepwave';
const ThePrimeAndArmyGeniuses = 'ThePrimeAndArmyGeniuses';

type TeamData = {
  [key: string]: string;
};

async function loadTestData(): Promise<TeamData> {
  const data = {
    Brame: await readFile('brame.html'),
    BrameDiffRef: await readFile('brame-diff-ref.html'),
    Creepwave: await readFile('creepwave.html'),
    ThePrimeAndArmyGeniuses: await readFile('the-prime-and-army-geniuses.html'),
  };
  return data;
}

async function readFile(name: string): Promise<string> {
  const file = path.join(__dirname, './testdata/', name);
  const content = readFileSync(file);
  return content.toString();
}

it('description contains a single link for each team', async function () {
  const data = await loadTestData();
  const tests = [
    {
      in: Brame,
      href: '/dota2/Brame',
    },
    {
      in: Creepwave,
      href: '/dota2/Creepwave',
    },
    {
      in: ThePrimeAndArmyGeniuses,
      href: '/dota2/The_Prime',
    },
    {
      in: ThePrimeAndArmyGeniuses,
      href: '/dota2/Army_Geniuses',
    },
  ];

  for (const t of tests) {
    const desc = description(data[t.in]);
    const $ = load(desc);
    const filter = `a[href*="${t.href}"]`;
    assert.equal($(filter).length, 1);
  }
});

it('description has no flags', async function () {
  const data = await loadTestData();
  const tests = [
    {
      in: Brame,
      flags: ['Ukraine'],
    },
    {
      in: Creepwave,
      flags: ['Belarus', 'Bulgaria', 'Jordan', 'Netherlands', 'Russia'],
    },
    {
      in: ThePrimeAndArmyGeniuses,
      flags: ['Indonesia'],
    },
  ];

  for (const t of tests) {
    const desc = description(data[t.in]);
    for (const f of t.flags) {
      assert.equal(desc.includes(f), false);
    }
  }
});

it('description has no images', async function () {
  const data = await loadTestData();
  const tests = [Brame, Creepwave, ThePrimeAndArmyGeniuses];

  for (const t of tests) {
    const desc = description(data[t]);
    const $ = load(desc);
    assert.equal($('img').length, 0);
  }
});

it('description used for hashing has no ref', async function () {
  const data = await loadTestData();
  const tests = [Brame, Creepwave, ThePrimeAndArmyGeniuses];

  for (const t of tests) {
    const desc = descriptionWithoutRef(data[t]);
    const $ = load(desc);
    assert.equal($('.Ref').length, 0);
  }
});

it('entry id contains hash', async function () {
  const data = await loadTestData();
  const tests = [
    {
      in: Brame,
      id: 'tag:liquipedia.net,2021-10-21:zHB8tgNDBQKQa6gbfhps9LY6zThmlZJN6pon3T9Xp04=',
    },
    {
      in: Creepwave,
      id: 'tag:liquipedia.net,2021-10-22:loV/ga28Pof6XJqRv28KcrjD+NXW+uKkPC6aQFW0eck=',
    },
    {
      in: ThePrimeAndArmyGeniuses,
      id: 'tag:liquipedia.net,2021-10-22:cOPmYZxq/B+/s/nAipZ4zPguIhbiDrzvVkUfX8Svhy8=',
    },
  ];

  for (const t of tests) {
    const id = await entryID(data[t.in]);
    assert.equal(t.id, id);
  }
});

it('entry id does not depend on ref', async function () {
  const brameID =
    'tag:liquipedia.net,2021-10-21:zHB8tgNDBQKQa6gbfhps9LY6zThmlZJN6pon3T9Xp04=';
  const data = await loadTestData();
  const tests = [Brame, BrameDiffRef];

  for (const t of tests) {
    const id = await entryID(data[t]);
    assert.equal(brameID, id);
  }
});
