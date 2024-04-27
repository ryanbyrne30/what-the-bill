from pymongo import MongoClient
from datetime import datetime
from transform.common import BASE_DATE


class USBillAction:
    def __init__(self) -> None:
        self.date: datetime = BASE_DATE
        self.text = ""


class USBill:
    def __init__(self) -> None:
        self.id = ""
        self.bill_id = ""
        self.short_title = ""
        self.title = ""
        self.url = ""
        self.text = ""
        self.version = ""
        self.type = ""
        self.issued = BASE_DATE
        self.updated = BASE_DATE
        self.actions: list[USBillAction] = []


class USBillsTransform:
    def __init__(
        self, mongo_url: str, mongo_db: str, src_col: str, dst_col: str
    ) -> None:
        self.src_col_name = src_col
        self.dst_col_name = dst_col
        self.client = MongoClient(mongo_url)
        self.db = self.client[mongo_db]
        self.src_col = self.db[src_col]
        self.dst_col = self.db[dst_col]

    def extract(self):
        """
        Extracts all US bills from database that need to be updated
        """
        self.src_col.aggregate(
            [
                {
                    "$lookup": {
                        "from": self.dst_col_name,
                        "localField": "packageId",
                        "foreignField": "bill_id",
                        "as": "bill_data",
                    },
                    "$match": {"bills_data": {"$eq": []}},
                }
            ]
        )

        pass
