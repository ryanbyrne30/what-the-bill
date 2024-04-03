from datetime import datetime, timedelta
import logging
import os
from bills import GovInfoBills, Fetch, GovInfoBill, Mongo
from typing import Any


if __name__ == "__main__":
    logging.root.setLevel(logging.INFO)

    api_key = os.environ["GOVINFO_API_KEY"]
    mongo_url = os.environ["MONGODB_URL"]
    mongo_db = os.environ["MONGODB_DATABASE"]
    mongo_col = os.environ["MONGODB_COLLECTION"]
    mongo = Mongo(mongo_url, mongo_db, mongo_col)

    fetch = Fetch()
    bills_scraper = GovInfoBills(api_key=api_key, fetch=fetch)
    bill_scraper = GovInfoBill(api_key=api_key, fetch=fetch)
    response = bills_scraper.fetch_bills(datetime.now() - timedelta(days=30), 4)
    packages = response.packages
    packages = list(
        filter(lambda p: not mongo.bill_exists(p.id, p.last_modified), packages)
    )

    print(f"Found {response.count} bills")

    for bill in packages:
        b = bill_scraper.fetch_bill(bill.link)
        print(f"{bill.last_modified} - {bill.title}")
        mongo.upsert_bill(b)
