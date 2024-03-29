from .fetch import Fetch
from .bill import BillScraper
from .bills import BillsSitemapScraper

fetch = Fetch()
bill_scraper = BillScraper(fetch=fetch)
bill_sitemap_scraper = BillsSitemapScraper(fetch)