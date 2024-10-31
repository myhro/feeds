import fetch from 'node-fetch';
import { load } from 'cheerio';
import { Feed } from 'feed';

(async () => {
  const url = 'https://devblogs.microsoft.com/oldnewthing/';
  const feed = new Feed({
    id: url,
    title: 'The Old New Thing',
    copyright: '',
    link: url,
  });

  const data = await fetch(url);
  const $ = load(await data.text());
  const posts = $('.masonry-card').slice(0, 10);
  for (const p of posts) {
    const title = $('h3', p).text().trim();
    const link = $('h3 a', p).attr('href') || '';
    const dateElem = $('.justify-content-between > div:first-child', p);
    const date = new Date(dateElem.text().trim());
    const content = $('p.mb-24', p).html() || '';
    const description = content.trim();

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
