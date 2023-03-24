import fetch from 'node-fetch';
import { load } from 'cheerio';
import { Feed } from 'feed';

(async () => {
  const url = 'https://about.sourcegraph.com/blog';
  const feed = new Feed({
    id: url,
    title: 'Sourcegraph Blog',
    copyright: '',
    link: url,
  });

  const root = 'https://about.sourcegraph.com';
  const data = await fetch(url);
  const $ = load(await data.text());
  const posts = $('article').slice(0, 10);
  for (const p of posts) {
    const title = $('h4', p).text();
    const href = $('a', p).attr('href') || '';
    const link = new URL(href, root).href;
    const author = $('.mb-0 span', p).text().trim();
    const time = $('time', p).text();
    const date = new Date(time);
    const description = $('.md\\:col-span-2 p', p).text();

    feed.addItem({
      title: title,
      author: author != '' ? [{ name: author }] : [],
      date: date,
      link: link,
      description: description,
    });
  }

  const atom = feed.atom1();
  console.log(atom);
})();
