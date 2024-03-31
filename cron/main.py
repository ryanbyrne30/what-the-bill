import math
import time
from scrape import (
    Proxy,
    Fetch,
    BillsSitemapScraper,
    BillScraper,
    BillSitemap,
    Bill as DataBill,
)
from store import Store, Bill as StoreBill
import logging
import threading
import os

URL = "https://www.govinfo.gov/sitemap/BILLS_2024_sitemap.xml"


def create_proxies(n: int) -> list[Proxy]:
    proxies = [Proxy(i, 5) for i in range(n)]
    for proxy in proxies:
        proxy.run()
    time.sleep(30)
    return proxies


def kill_proxies(proxies: list[Proxy]):
    for proxy in proxies:
        proxy.kill()


def to_store_bill(b: DataBill) -> StoreBill:
    sb = StoreBill()

    sb.short_title = b.short_title
    sb.full_title = b.full_title
    sb.category = b.category
    sb.url = b.url
    sb.collection = b.collection
    sb.sudoc_class_number = b.sudoc_class_number
    sb.congress_number = b.congress_number
    sb.congress_session = b.congress_session
    sb.last_action_date = b.last_action_date
    sb.action = b.action
    sb.actions = b.actions
    sb.bill_number = b.bill_number
    sb.bill_version = b.bill_version
    sb.sponsors = b.sponsors
    sb.cosponsors = b.cosponsors
    sb.committees = b.committees
    sb.us_code_reference = b.us_code_reference
    sb.text = b.text
    sb.congress_updated_at = b.congress_updated_at

    return sb


def scrape_bills(
    id: int, bills_to_scrape: list[BillSitemap], scraper: BillScraper, store: Store
):
    num_bills = len(bills_to_scrape)
    for i, b in enumerate(bills_to_scrape):
        pct_complete = round(i / num_bills * 100, 0)
        logging.info(
            f"Thread[{id}]: Scraping {i} of {num_bills} ({pct_complete}% complete)"
        )
        try:
            bill = scraper.scrape(b.url)
            bill.congress_updated_at = b.updated_at
            store_bill = to_store_bill(bill)
            store.upsert_bill(store_bill)
        except Exception as e:
            logging.error(f"Could not scrape bill: {b.url}")
            logging.error(e)


if __name__ == "__main__":
    logging.basicConfig()
    logging.root.setLevel(logging.INFO)

    # get env vars
    mongo_url = os.environ["MONGODB_URL"]
    mongo_db = os.environ["MONGODB_DATABASE"]
    mongo_col = os.environ["MONGODB_COLLECTION"]

    # setup resources for threads
    thread_count = 10
    logging.info("Setting up proxies...")
    proxies = create_proxies(thread_count)

    # setup main resources
    fetch = Fetch(proxies[0])
    sitemap_scraper = BillsSitemapScraper(fetch)
    store = Store(mongo_url, mongo_db, mongo_col)

    # fetch and parse bill metadata
    logging.info("Fetching bill metadata")
    bills_meta = sitemap_scraper.get_bills(URL)
    logging.info(f"Metadata found for {len(bills_meta)} bills")
    bills_to_scrape = list(filter(lambda b: store.should_scrape(b.url), bills_meta))
    bills_to_scrape = bills_to_scrape[:10]
    logging.info(f"{len(bills_to_scrape)} bills to scrape")

    # prep bills for threads
    chunk_size = math.ceil(len(bills_to_scrape) / thread_count)
    chunks = [
        bills_to_scrape[i * chunk_size : (i + 1) * chunk_size]
        for i in range(thread_count)
    ]

    # create threads
    scrapers = [BillScraper(Fetch(proxies[i])) for i in range(thread_count)]
    stores = [
        Store("mongodb://localhost:27017/", "bills", "us-congress-bills")
        for _ in range(thread_count)
    ]
    threads = [
        threading.Thread(
            target=scrape_bills,
            args=(i, chunks[i], scrapers[i], stores[i]),
        )
        for i in range(thread_count)
    ]

    logging.info(f"Starting {thread_count} threads...")
    for t in threads:
        t.start()

    logging.info(f"Done. Awaiting process to finish")
    for t in threads:
        t.join()

    kill_proxies(proxies)
