import crypto from 'crypto';
import { load } from 'cheerio';

export function description(html: string): string {
  const $ = load(html);
  $('.flag').remove();
  $('.team-template-darkmode').remove();

  $('img').each((i, el) => {
    const alt = el.attribs['alt'];
    $(el).replaceWith(alt);
  });

  return $.html();
}

export function descriptionWithoutRef(html: string): string {
  const $ = load(html);
  $('.Ref').remove();
  return $.html();
}

export async function entryID(html: string): Promise<string> {
  const $ = load(html);
  const date = $('.Date').text();
  const desc = descriptionWithoutRef(html);
  const sha256 = crypto.createHash('sha256');
  const hash = sha256.update(desc).digest('base64');
  const id = `tag:liquipedia.net,${date}:${hash}`;
  return id;
}
