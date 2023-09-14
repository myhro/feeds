import fetch from 'node-fetch';
import { load } from 'cheerio';
import { Feed } from 'feed';

(async () => {
  const url = 'https://www.autossegredos.com.br/category/segredos/';
  const feed = new Feed({
    id: url,
    title: 'Autos Segredos - Arquivos Segredos',
    copyright: '',
    link: url,
  });

  const data = await fetch(url);
  const $ = load(await data.text());
  const posts = $('.tdb_module_loop').slice(0, 10);
  for (const p of posts) {
    const title = $('.td-module-title', p).text();
    const link = $('.td-module-title a', p).attr('href') || '';
    const dateText = $('.td-post-date time', p).attr('datetime') || '';
    const date = new Date(dateText);

    const post = await fetch(link);
    const content = await post.text();
    const description = $('.td-fix-index p:first', content).text();

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
