from urllib.request import urlopen
from urllib.parse import urlparse
from bs4 import BeautifulSoup
import re
import datetime
import random

pages = set()
random.seed(1)


def get_internal_links(bs: BeautifulSoup, include_url):
    include_url = '{}://{}'.format(urlparse(include_url).scheme, urlparse(include_url).netloc)
    internal_links = []

    for link in bs.find_all('a', href=re.compile('^(/|.*'+include_url+')')):
        if link.attrs['href'] is not None:
            if link.attrs['href'] not in internal_links:
                if link.attrs['href'].startswith('/'):
                    internal_links.append(include_url + link.attrs['href'])
                else:
                    internal_links.append(link.attrs['href'])
    return internal_links


def get_external_links(bs: BeautifulSoup, exclude_url):
    external_links = []

    for link in bs.find_all('a'):
        href = re.compile('^(http|www)((?!' + exclude_url + ').)*$')
        if link.attrs['href'] is not None:
            if link.attrs['href'] not in external_links:
                external_links.append(link.attrs['href'])
    return external_links


def get_random_external_link(start_page):
    html = urlopen(start_page)
    bs = BeautifulSoup(html, 'html.parser')
    external_links = get_external_links(bs, urlparse(start_page).netloc)
    if len(external_links) == 0:
        print('No external links.')
        domain = '{}://{}'.format(urlparse(start_page).scheme, urlparse(start_page).netloc)
        internal_links = get_internal_links(bs, domain)
        return get_random_external_link(internal_links[random.randint(0, len(internal_links)-1)])
    else:
        return external_links[random.randint(0, len(external_links)-1)]
