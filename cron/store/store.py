from pymongo import MongoClient
from datetime import datetime
from .bill import Bill
from . import config


class Store:
    def __init__(self, conn_str: str, db: str, col: str) -> None:
        self.client = MongoClient(conn_str)
        self.db = self.client[db]
        self.col = self.db[col]

    def __stored_str_to_datetime(self, d: str) -> datetime:
        return datetime.strptime(d, config.DATE_FMT)

    def __find_by_url(self, url: str) -> Bill | None:
        data = self.col.find_one({"url": url})
        if data is None:
            return None

        bill = Bill()

        if "short_title" in data:
            bill.short_title = data["short_title"]
        if "full_title" in data:
            bill.full_title = data["full_title"]
        if "category" in data:
            bill.category = data["category"]
        if "url" in data:
            bill.url = data["url"]
        if "collection" in data:
            bill.collection = data["collection"]
        if "sudoc_class_number" in data:
            bill.sudoc_class_number = data["sudoc_class_number"]
        if "congress_number" in data:
            bill.congress_number = data["congress_number"]
        if "congress_session" in data:
            bill.congress_session = data["congress_session"]
        if "last_action_date" in data:
            bill.last_action_date = data["last_action_date"]
        if "action" in data:
            bill.action = data["action"]
        if "actions" in data:
            bill.actions = data["actions"]
        if "bill_number" in data:
            bill.bill_number = data["bill_number"]
        if "bill_version" in data:
            bill.bill_version = data["bill_version"]
        if "sponsors" in data:
            bill.sponsors = data["sponsors"]
        if "cosponsors" in data:
            bill.cosponsors = data["cosponsors"]
        if "committees" in data:
            bill.committees = data["committees"]
        if "us_code_reference" in data:
            bill.us_code_reference = data["us_code_reference"]
        if "text" in data:
            bill.text = data["text"]
        if "updated_at" in data:
            bill.updated_at = data["updated_at"]
        if "congress_updated_at" in data:
            bill.congress_updated_at = data["congress_updated_at"]

        return bill

    def should_scrape(self, bill_url: str) -> bool:
        bill = self.__find_by_url(bill_url)
        if bill is None:
            return True
        return bill.congress_updated_at > bill.updated_at

    def upsert_bill(self, bill: Bill):
        bill.updated_at = datetime.now()

        query = {"url": bill.url}
        update = {"$set": bill.__dict__}

        self.col.update_one(query, update, upsert=True)
