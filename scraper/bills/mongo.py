from datetime import datetime
from pymongo import MongoClient
from bills.bill import Bill


class Mongo:
    def __init__(self, conn_str: str, db: str, col: str) -> None:
        self.client = MongoClient(conn_str)
        self.db = self.client[db]
        self.col = self.db[col]

    def bill_exists(self, bill_id: str, last_updated: datetime) -> bool:
        bill = self.col.find_one({"bill_id": bill_id, "updated": last_updated})
        return bill is not None

    def upsert_bill(self, bill: Bill):
        query = {"bill_id": bill.bill_id}
        update = {
            "$set": {
                "bill_id": bill.bill_id,
                "title": bill.title,
                "short_title": bill.short_title,
                "url": bill.url,
                "text": bill.text,
                "version": bill.version,
                "type": bill.type,
                "issued": bill.issued,
                "updated": bill.updated,
                "actions": [{"text": a.text, "date": a.date} for a in bill.actions],
            }
        }
        self.col.update_one(query, update, upsert=True)
