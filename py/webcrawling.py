import datetime
import random
from urllib.request import urlopen
from bs4 import BeautifulSoup
import re

pages = set()
max_depth = 10


def get_links(url, depth):
    global pages

    if depth > max_depth:
        return

    html = urlopen('http://en.wikipedia.org{}'.format(url))
    bs = BeautifulSoup(html, 'html.parser')

    try:
        print(bs.h1.get_text())
        print(bs.find(id='mw-content-text').find_all('p')[0])
        print(bs.find(id='ca-edit').find('span').find('a').attrs['href'])
    except AttributeError:
        print("Missing something")

    for link in bs.find_all('a', href=re.compile('^(/wiki/)')):
        if 'href' in link.attrs:
            if link.attrs['href'] not in pages:
                new_page = link.attrs['href']
                print('-'*20)
                print(new_page)
                pages.add(new_page)
                get_links(new_page, depth + 1)


def main():
    get_links('/wiki/Kevin_Bacon', 0)


if __name__ == "__main__":
    # random.seed(datetime.datetime.now())
    main()
