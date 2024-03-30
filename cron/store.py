from typing import Any, TypeVar
from pymongo import MongoClient
import json

T = TypeVar("T")


class Bill:
    def __init__(self, data: str = "") -> None:
        print("data", data)
        d = json.loads(data)

        self.url = self.json_value_or_default(d, "url", "")

        self.__dict__ = {
            "short_title": self.json_value_or_default(d, "short_title", ""),
            "full_title": self.json_value_or_default(d, "full_title", ""),
            "category": self.json_value_or_default(d, "category", ""),
            "url": self.url,
            "collection": self.json_value_or_default(d, "collection", ""),
            "sudoc_class_number": self.json_value_or_default(
                d, "sudoc_class_number", ""
            ),
            "congress_number": self.json_value_or_default(d, "congress_number", ""),
            "congress_session": self.json_value_or_default(d, "congress_session", ""),
            "last_action_date": self.json_value_or_default(d, "last_action_date", ""),
            "action": self.json_value_or_default(d, "action", ""),
            "actions": self.json_value_or_default(d, "actions", ""),
            "bill_number": self.json_value_or_default(d, "bill_number", ""),
            "bill_version": self.json_value_or_default(d, "bill_version", ""),
            "sponsors": self.json_value_or_default(d, "sponsors", ""),
            "cosponsors": self.json_value_or_default(d, "cosponsors", ""),
            "committees": self.json_value_or_default(d, "committees", ""),
            "us_code_reference": self.json_value_or_default(d, "us_code_reference", ""),
            "text": self.json_value_or_default(d, "text", ""),
        }

    def json_value_or_default(self, d: Any, key: str, default: T) -> T:
        if key in d:
            return d[key]
        return default


client = MongoClient("mongodb://localhost:27017/")

db = client["bills"]

col = db["us-congress-bills"]


if __name__ == "__main__":
    import sys

    for line in sys.stdin:
        bill = Bill(line.strip())

        print("Finding", bill.url)

        doc = col.find_one({"url": bill.url})
        if doc is not None:
            continue

        id = col.insert_one(bill.__dict__).inserted_id
        print("ID", id)
