from datetime import datetime, timedelta
import logging
import os
from bills import GovInfoBills, Fetch, GovInfoBill
import json
from typing import Any


if __name__ == "__main__":
    logging.root.setLevel(logging.INFO)

    api_key = os.environ["GOVINFO_API_KEY"]

    fetch = Fetch()
    bills_scraper = GovInfoBills(api_key=api_key, fetch=fetch)
    bill_scraper = GovInfoBill(api_key=api_key, fetch=fetch)
    response = bills_scraper.fetch_bills(datetime.now() - timedelta(days=30), 2)

    print(f"Found {response.count} bills")

    for bill in response.packages:
        b = bill_scraper.fetch_bill(bill.link)
        print(f"{bill.last_modified} - {bill.title}")

        d = {
            "id": b.bill_id,
            "title": b.title,
            "url": b.url,
            "issued": b.issued.strftime("%Y-%m-%d"),
            "updated": b.updated.strftime("%Y-%m-%d"),
            "actions": [
                {"action": a.text, "date": a.date.strftime("%Y-%m-%d")}
                for a in b.actions
            ],
            "text": b.text,
        }

        with open(f"output/{b.bill_id}", "w") as f:
            json.dump(d, f, indent=2)
