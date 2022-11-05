import fetch from 'node-fetch';
import { load } from 'cheerio';
import { Feed } from 'feed';

function cleanDescription(html: string): string {
  const $ = load(html);
  $('.entry-header').remove();
  $('.entry-footer').remove();
  return $.text().trim();
}

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
  const posts = $('.entry-area').slice(0, 10);
  for (const p of posts) {
    const title = $('.entry-title', p).text();
    const link = $('.entry-title a', p).attr('href') || '';
    const dateText = $('.entry-post-date', p).text().trim();
    const date = new Date(dateText);
    const content = $('.entry-content', p).html() || '';
    const description = cleanDescription(content);

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
