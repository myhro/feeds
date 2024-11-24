import fetch from 'node-fetch';
import { load } from 'cheerio';
import { Feed } from 'feed';
import { description, entryID } from '../src/liquipedia';

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
  const posts = $('.divRow');
  let validPosts = 0;
  for (const p of posts) {
    if (validPosts == 10) break;

    const dateText = $('.Date', p).text();
    const date = new Date(dateText);
    if (isNaN(date.getTime())) {
      continue;
    }

    const title = $('.Name', p).text().trim();
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

    validPosts++;
  }

  const atom = feed.atom1();
  console.log(atom);
})();
