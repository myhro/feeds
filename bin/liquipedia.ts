import fetch from 'node-fetch';
import { load } from 'cheerio';
import { Feed } from 'feed';
import { description, entryID } from '../src/liquipedia'

(async () => {
  const url = 'https://liquipedia.net/dota2/Portal:Transfers';
  const feed = new Feed({
    id: url,
    title: 'Liquipedia - Player Transfers',
    copyright: '',
    link: url,
  });

  const data = await fetch(url);
  const $ = load(await data.text());
  const posts = $('.divRow').slice(0, 10);
  for (const p of posts) {
    const title = $('.Name', p).text().trim();
    const dateText = $('.Date', p).text();
    const date = new Date(dateText);
    const content = $(p).html() || '';
    const desc = description(content);
    const id = await entryID(content);

    feed.addItem({
      id: id,
      title: title,
      date: date,
      link: url,
      description: desc,
    });
  }

  const atom = feed.atom1();
  console.log(atom);
})();
