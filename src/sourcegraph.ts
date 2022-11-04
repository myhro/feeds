import cheerio from 'cheerio';
import fetch from 'node-fetch';
import path from 'path';
import { Feed } from 'feed';

(async () => {
  const root = 'https://about.sourcegraph.com';
  const url = 'https://about.sourcegraph.com/blog';
  const feed = new Feed({
    id: url,
    title: 'Sourcegraph Blog',
    copyright: '',
  });

  const data = await fetch(url);
  const $ = cheerio.load(await data.text());
  const posts = $('.card').slice(0, 10);
  for (const p of posts) {
    const title = $('.tw-block', p).text();
    const href = $('.tw-block', p).attr('href') || '';
    const link = path.join(root, href);
    const time = $('time', p).text();
    const date = new Date(time);
    const description = $('.row p', p).text();

    feed.addItem({
      title: title,
      date: date,
      link: link,
      description: description,
    });
  }

  const atom = feed.atom1();
  console.log(atom);
})();
