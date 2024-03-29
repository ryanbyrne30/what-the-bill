from .fetch import Fetch
from bs4 import BeautifulSoup

class Bill:
    def __init__(self) -> None:
        self.category: str = ""
        self.collection: str = ""
        self.sudoc_class_number: str = ""
        self.congress_number: str = ""
        self.congress_session: str = ""
        self.last_action_date: str = ""
        self.action: str = ""
        self.bill_number: str = ""
        self.bill_version: str = ""
        self.short_title: str = ""
        self.full_title: str = ""
        self.sponsors: str = ""
        self.committees: str = ""
        self.us_code_reference: str = ""
        self.text: str = "" 
        self.soup: BeautifulSoup = None
    
    def __str__(self) -> str:
        return f"Bill(category: {self.category}, bill_number: {self.bill_number}, short_title: {self.short_title}, text_length: {len(self.text)})"
    
        
class BillScraper:
    def __init__(self, fetch: Fetch) -> None:
        self.fetch = fetch

    def scrape(self, url: str) -> Bill:
        soup = self.__soupify(url)
        bill = Bill()
        bill.category = self.__scrape_summary_value(soup, "Category")
        bill.collection = self.__scrape_summary_value(soup, "Collection")
        bill.sudoc_class_number = self.__scrape_summary_value(soup, "SuDoc Class Number")
        bill.congress_number = self.__scrape_summary_value(soup, "Congress Number")
        bill.congress_session = self.__scrape_summary_value(soup, "Congress Session")
        bill.last_action_date = self.__scrape_summary_value(soup, "Last Action Date Listed")
        bill.action = self.__scrape_summary_value(soup, "Action")
        bill.bill_number = self.__scrape_summary_value(soup, "Bill Number")
        bill.bill_version = self.__scrape_summary_value(soup, "Bill Version")
        bill.short_title = self.__scrape_summary_value(soup, "Short Title")
        bill.full_title = self.__scrape_summary_value(soup, "Full Title")
        bill.sponsors = self.__scrape_summary_value(soup, "Sponsors")
        bill.committees = self.__scrape_summary_value(soup, "Committees")
        bill.us_code_reference = self.__scrape_summary_value(soup, "United States Code Reference")
        bill.text = self.__scrape_text(soup)
        return bill

    def __soupify(self, url: str) -> BeautifulSoup:
        html = self.fetch.dynamic_request(url, "div#accMetadata div.row")
        soup = BeautifulSoup(html, "html.parser")
        return soup

    def __scrape_text(self, soup: BeautifulSoup) -> str:
        if soup is None:
            return ""
        text_link = "https:" + soup.select_one("div.panel a#text").attrs["href"]
        html = self.fetch.simple_request(text_link)
        soup = BeautifulSoup(html, "html.parser")
        text = soup.text 
        return text
    
    def __scrape_summary_value(self, soup: BeautifulSoup, data_id: str) -> str:
        label = soup.find('div', attrs={'data-id': data_id})
        return label.parent.find("p").text.strip()



    


    
        
