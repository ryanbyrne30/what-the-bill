from .fetch import Fetch
from bs4 import BeautifulSoup
from typing import Any


class Bill:
    def __init__(self) -> None:
        self.url: str = ""
        self.category: str = ""
        self.collection: str = ""
        self.sudoc_class_number: str = ""
        self.congress_number: str = ""
        self.congress_session: str = ""
        self.last_action_date: str = ""
        self.action: str = ""
        self.actions: str = ""
        self.bill_number: str = ""
        self.bill_version: str = ""
        self.short_title: str = ""
        self.full_title: str = ""
        self.sponsors: str = ""
        self.cosponsors: str = ""
        self.committees: str = ""
        self.us_code_reference: str = ""
        self.text: str = ""
        self.soup: BeautifulSoup | None = None
        self.__dict__ = {
            "short_title": self.short_title,
            "full_title": self.full_title,
            "category": self.category,
            "url": self.url,
            "collection": self.collection,
            "sudoc_class_number": self.sudoc_class_number,
            "congress_number": self.congress_number,
            "congress_session": self.congress_session,
            "last_action_date": self.last_action_date,
            "action": self.action,
            "actions": self.actions,
            "bill_number": self.bill_number,
            "bill_version": self.bill_version,
            "sponsors": self.sponsors,
            "cosponsors": self.cosponsors,
            "committees": self.committees,
            "us_code_reference": self.us_code_reference,
            "text": self.text,
        }

    def __str__(self) -> str:
        return f"Bill(category: {self.category}, bill_number: {self.bill_number}, short_title: {self.short_title}, text_length: {len(self.text)})"


class BillScraper:
    def __init__(self, fetch: Fetch) -> None:
        self.fetch = fetch

    def scrape(self, url: str) -> Bill:
        soup = self.__soupify(url)
        bill = Bill()
        bill.url = url
        bill.category = self.__scrape_summary_value(soup, "Category")
        bill.collection = self.__scrape_summary_value(soup, "Collection")
        bill.sudoc_class_number = self.__scrape_summary_value(
            soup, "SuDoc Class Number"
        )
        bill.congress_number = self.__scrape_summary_value(soup, "Congress Number")
        bill.congress_session = self.__scrape_summary_value(soup, "Congress Session")
        bill.last_action_date = self.__scrape_summary_value(
            soup, "Last Action Date Listed"
        )
        bill.action = self.__scrape_summary_value(soup, "Action")
        bill.action = self.__scrape_summary_value(soup, "Actions")
        bill.bill_number = self.__scrape_summary_value(soup, "Bill Number")
        bill.bill_version = self.__scrape_summary_value(soup, "Bill Version")
        bill.short_title = self.__scrape_summary_value(soup, "Short Title")
        bill.full_title = self.__scrape_summary_value(soup, "Full Title")
        bill.sponsors = self.__scrape_summary_value(soup, "Sponsors")
        bill.cosponsors = self.__scrape_summary_value(soup, "Cosponsors")
        bill.committees = self.__scrape_summary_value(soup, "Committees")
        bill.us_code_reference = self.__scrape_summary_value(
            soup, "United States Code Reference"
        )
        bill.text = self.__scrape_text(soup)
        return bill

    def __soupify(self, url: str) -> BeautifulSoup:
        html = self.fetch.dynamic_request(url, "div#accMetadata div.row")
        soup = BeautifulSoup(html, "html.parser")
        return soup

    def __scrape_text(self, soup: BeautifulSoup) -> str:
        if soup is None:
            return ""
        text_el = soup.select_one("div.panel a#text")
        if text_el is None:
            return ""
        text_link = "https:" + text_el.attrs["href"]
        html = self.fetch.dynamic_request(text_link, wait_for="body")
        soup = BeautifulSoup(html, "html.parser")
        text = soup.text
        return text

    def __scrape_summary_value(self, soup: BeautifulSoup, data_id: str) -> str:
        label = soup.find("div", attrs={"data-id": data_id})
        if label is None:
            return ""

        parent = label.parent
        if parent is None:
            return ""

        p_tag = parent.find("p")
        if p_tag is None:
            return ""

        return p_tag.text.strip()
