from .fetch import Fetch
from bs4 import BeautifulSoup

class BillsSitemapScraper:
    def __init__(self, fetch: Fetch) -> None:
        self.fetch = fetch

    def get_bill_urls(self, url: str) -> list[str]:
        content = self.fetch.simple_request(url)
        soup = BeautifulSoup(content, 'xml')
        urls = soup.find_all("url")
        return [
            url.find('loc').text
            for url in urls
        ]