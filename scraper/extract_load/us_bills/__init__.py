from .api import API
from .mongo import Mongo
from extract_load.fetch import Fetch
from datetime import datetime


class USBills:
    def __init__(
        self, api_key: str, mongo_url: str, mongo_db: str, mongo_col: str
    ) -> None:
        fetch = Fetch()
        self.api = API(api_key=api_key, fetch=fetch)
        self.mongo = Mongo(mongo_url, mongo_db, mongo_col)

    def run(self, since: datetime, limit: int = 1000, offset: int = 0):
        data = self.api.fetch_bills(from_date=since, limit=limit, offset=offset)
        return self.mongo.upsert_bills(data)
