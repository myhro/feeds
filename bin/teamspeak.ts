import fetch from 'node-fetch';
import { Feed } from 'feed';

(async () => {
  const url = 'https://hub.docker.com/_/teamspeak/tags';
  const feed = new Feed({
    id: url,
    title: 'Docker Hub - Teamspeak',
    copyright: '',
    link: url,
  });

  const data = await fetch(
    'https://hub.docker.com/v2/namespaces/library/repositories/teamspeak/tags'
  );
  const json = await data.json();
  const results = json.results.slice(0, 10);
  for (const r of results) {
    const title = 'teamspeak:' + r.name;
    const date = new Date(r.last_updated);
    const image = r.images.shift();
    const hash = image.digest.replace(':', '-');
    const link = `https://hub.docker.com/layers/library/teamspeak/${r.name}/images/${hash}`;

    feed.addItem({
      title: title,
      date: date,
      link: link,
    });
  }

  const atom = feed.atom1();
  console.log(atom);
})();
