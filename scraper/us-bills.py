import os
from scraper.extract_load import USBills
from datetime import datetime, timedelta
import logging

if __name__ == "__main__":
    from dotenv import load_dotenv

    load_dotenv()

    logging.basicConfig()
    logging.getLogger().setLevel(logging.INFO)

    api_key = os.environ["GOVINFO_API_KEY"]
    mongo_url = os.environ["MONGODB_URL"]
    mongo_db = "legislature"
    mongo_col = "us-bills"

    us_bills = USBills(
        api_key=api_key, mongo_url=mongo_url, mongo_db=mongo_db, mongo_col=mongo_col
    )

    since = datetime.now() - timedelta(days=1)

    us_bills.run(since=since, limit=1000, offset=0)
