from .fetch import Fetch
from bs4 import BeautifulSoup, ResultSet, Tag
from datetime import datetime
from .config import BASE_DATE


class BillSitemap:
    def __init__(self) -> None:
        self.url = ""
        self.updated_at: datetime = BASE_DATE


class BillsSitemapScraper:
    def __init__(self, fetch: Fetch) -> None:
        self.fetch = fetch

    def get_bills(self, sitemap_url: str) -> list[BillSitemap]:
        content = self.fetch.simple_request(sitemap_url)
        soup = BeautifulSoup(content, "xml")
        blocks: ResultSet[Tag] = soup.find_all("url")
        bills: list[BillSitemap] = []

        for block in blocks:
            bill = BillSitemap()

            url = block.find("loc")
            if url is None:
                continue
            url = url.text.strip()

            updated_at = block.find("lastmod")
            if updated_at is None:
                continue
            updated_at = updated_at.text.split(".")[0]
            # 2024-01-03T04:15:00.400Z
            updated_at = datetime.strptime(updated_at, "%Y-%m-%dT%H:%M:%S")

            bill.url = url
            bill.updated_at = updated_at
            bills.append(bill)

        return bills

    def get_bill_urls(self, url: str) -> list[str]:
        content = self.fetch.simple_request(url)
        soup = BeautifulSoup(content, "xml")
        urls = soup.find_all("url")
        return [url.find("loc").text for url in urls]
